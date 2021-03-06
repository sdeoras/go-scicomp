package nnl4

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/sdeoras/comp/nn"

	"github.com/sdeoras/comp/image"
)

func TestOp_Step(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	height := 128
	width := 128
	numIterations := 100
	miniBatchSize := 100
	numClasses := 2
	learningRate := 0.0003

	op, err := New([]int{height * width, 100, 50, 10, numClasses}, learningRate, "image-classification", "/tmp/train", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer op.Close()

	fmt.Println("run >>> tensorboard --logdir=/tmp/train")

	trainingImages := make([]image.Image, miniBatchSize)
	trainingLabels := make([][]float32, miniBatchSize)
	validationImages := make([]image.Image, miniBatchSize)
	validationLabels := make([][]float32, miniBatchSize)

	imageOp, err := image.NewOperator(nil)
	if err != nil {
		t.Fatal(err)
	}

	for k := 0; k < numIterations; k++ {
		for i := range trainingImages {
			img := image.NewBlack(height, width)
			trainingLabels[i] = make([]float32, 2)
			if rand.Intn(2) > 0 {
				//img.RandBox()
				trainingLabels[i][0] = 1
			} else {
				img.Circle(64, 64, 30)
				trainingLabels[i][1] = 1
			}
			trainingImages[i] = img
		}

		for i := range validationImages {
			img := image.NewBlack(height, width)
			validationLabels[i] = make([]float32, 2)
			if rand.Intn(2) > 0 {
				//img.RandBox()
				validationLabels[i][0] = 1
			} else {
				img.Circle(64, 64, 30)
				validationLabels[i][1] = 1
			}
			validationImages[i] = img
		}

		trainingData, err := imageOp.MakeBatch(height, width, trainingImages...)
		if err != nil {
			t.Fatal(err)
		}

		validationData, err := imageOp.MakeBatch(height, width, validationImages...)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := op.Step(
			&nn.Data{
				X: trainingData,
				Y: trainingLabels,
			},
			&nn.Data{
				X: validationData,
				Y: validationLabels,
			}); err != nil {
			t.Fatal(err)
		}
	}

	// now save checkpoint
	cp, err := op.Save()
	if err != nil {
		t.Fatal(err)
	}

	// load checkpoint
	if err := op.Load(cp); err != nil {
		t.Fatal(err)
	}

	// now check prediction
	img := image.NewBlack(height, width)
	labels := make([]float32, 2)
	if rand.Intn(2) > 0 {
		//img.RandBox()
		labels[0] = 1
	} else {
		img.Circle(64, 64, 30)
		labels[1] = 1
	}

	data, err := imageOp.MakeBatch(height, width, img)
	if err != nil {
		t.Fatal(err)
	}

	out, err := op.Predict(&nn.Data{X: data, Y: [][]float32{labels}})
	if err != nil {
		t.Fatal(err)
	}

	if jb, err := json.MarshalIndent(out, "", "  "); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(string(jb))
	}

	if out.Accuracy != 1 {
		t.Fatal("test failed to predict with 100% accuracy, got:", out.Accuracy)
	}
}

func TestOp_FizzBuzz(t *testing.T) {
	numIterations := 10000
	miniBatchSize := 1000
	numClasses := 4 // fizz, buzz, fizz-buzz, .
	learningRate := 0.0001

	primes := []int{2, 3, 5, 7, 11, 13}
	numFeatures := len(primes) + 1

	op, err := New([]int{numFeatures, 5, 4, 3, numClasses}, learningRate, "fizz-buzz", "/tmp/train", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer op.Close()

	fmt.Println("run >>> tensorboard --logdir=/tmp/train")

	// training and validation data derived from 100 random numbers
	trainingData := make([][]float32, miniBatchSize)
	validationData := make([][]float32, miniBatchSize)

	trainingLabels := make([][]float32, miniBatchSize)
	validationLabels := make([][]float32, miniBatchSize)

	for k := 0; k < numIterations; k++ {
		for i := range trainingData {
			n := rand.Int()
			trainingData[i] = make([]float32, numFeatures)
			trainingLabels[i] = make([]float32, numClasses)

			// build input data
			assigned1 := false
			for j := range primes {
				if n%primes[j] == 0 {
					trainingData[i][j] = 1
					assigned1 = true
				}
			}

			if !assigned1 {
				trainingData[i][numFeatures-1] = 1
			}

			// now assign classes
			if n%15 == 0 {
				trainingLabels[i][0] = 1
			} else if n%5 == 0 {
				trainingLabels[i][1] = 1
			} else if n%3 == 0 {
				trainingLabels[i][2] = 1
			} else {
				trainingLabels[i][3] = 1
			}
		}

		for i := range validationData {
			n := rand.Int()
			validationData[i] = make([]float32, numFeatures)
			validationLabels[i] = make([]float32, numClasses)

			// build input data
			assigned1 := false
			for j := range primes {
				if n%primes[j] == 0 {
					validationData[i][j] = 1
					assigned1 = true
				}
			}

			if !assigned1 {
				validationData[i][numFeatures-1] = 1
			}

			// now assign classes
			if n%15 == 0 {
				validationLabels[i][0] = 1
			} else if n%5 == 0 {
				validationLabels[i][1] = 1
			} else if n%3 == 0 {
				validationLabels[i][2] = 1
			} else {
				validationLabels[i][3] = 1
			}
		}

		if _, err := op.Step(
			&nn.Data{
				X: trainingData,
				Y: trainingLabels,
			},
			&nn.Data{
				X: validationData,
				Y: validationLabels,
			}); err != nil {
			t.Fatal(err)
		}
	}

	// now save checkpoint
	cp, err := op.Save()
	if err != nil {
		t.Fatal(err)
	}

	// load checkpoint
	if err := op.Load(cp); err != nil {
		t.Fatal(err)
	}
}

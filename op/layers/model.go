package layers

const (
	// model is base64 representation of protocol buffer TF model.
	// this go file was auto-generated using following command:
	//      go get github.com/sdeoras/tensorflow/cmd/mo
	//      mo u -h
	modelPB = "" +
		"CjUKB3ZlcnNpb24SBUNvbnN0KhYKBXZhbHVlEg1CCwgHEgBCBTAuMS4wKgsKBWR0eXBlEgIwBwpSCgdt" +
		"eUlucHV0EgtQbGFjZWhvbGRlciotCgVzaGFwZRIkOiISAggBEgsI////////////ARILCP//////////" +
		"/wESAggDKgsKBWR0eXBlEgIwBApTCgxteUlucHV0U2xpY2USC1BsYWNlaG9sZGVyKgsKBWR0eXBlEgIw" +
		"BCopCgVzaGFwZRIgOh4SAggBEgsI////////////ARILCP///////////wEKOAoNZmxhdHRlbi9TaGFw" +
		"ZRIFU2hhcGUaB215SW5wdXQqBwoBVBICMAQqDgoIb3V0X3R5cGUSAjADCkkKG2ZsYXR0ZW4vc3RyaWRl" +
		"ZF9zbGljZS9zdGFjaxIFQ29uc3QqFgoFdmFsdWUSDUILCAMSBBICCAE6AQAqCwoFZHR5cGUSAjADCksK" +
		"HWZsYXR0ZW4vc3RyaWRlZF9zbGljZS9zdGFja18xEgVDb25zdCoWCgV2YWx1ZRINQgsIAxIEEgIIAToB" +
		"ASoLCgVkdHlwZRICMAMKSwodZmxhdHRlbi9zdHJpZGVkX3NsaWNlL3N0YWNrXzISBUNvbnN0KgsKBWR0" +
		"eXBlEgIwAyoWCgV2YWx1ZRINQgsIAxIEEgIIAToBAQqJAgoVZmxhdHRlbi9zdHJpZGVkX3NsaWNlEgxT" +
		"dHJpZGVkU2xpY2UaDWZsYXR0ZW4vU2hhcGUaG2ZsYXR0ZW4vc3RyaWRlZF9zbGljZS9zdGFjaxodZmxh" +
		"dHRlbi9zdHJpZGVkX3NsaWNlL3N0YWNrXzEaHWZsYXR0ZW4vc3RyaWRlZF9zbGljZS9zdGFja18yKhAK" +
		"CmJlZ2luX21hc2sSAhgAKhMKDWVsbGlwc2lzX21hc2sSAhgAKhMKDW5ld19heGlzX21hc2sSAhgAKg4K" +
		"CGVuZF9tYXNrEgIYACoHCgFUEgIwAyoLCgVJbmRleBICMAMqFgoQc2hyaW5rX2F4aXNfbWFzaxICGAEK" +
		"SgoXZmxhdHRlbi9SZXNoYXBlL3NoYXBlLzESBUNvbnN0KhsKBXZhbHVlEhJCEAgDEgA6Cv//////////" +
		"/wEqCwoFZHR5cGUSAjADCmsKFWZsYXR0ZW4vUmVzaGFwZS9zaGFwZRIEUGFjaxoVZmxhdHRlbi9zdHJp" +
		"ZGVkX3NsaWNlGhdmbGF0dGVuL1Jlc2hhcGUvc2hhcGUvMSoHCgFUEgIwAyoKCgRheGlzEgIYACoHCgFO" +
		"EgIYAgpRCg9mbGF0dGVuL1Jlc2hhcGUSB1Jlc2hhcGUaB215SW5wdXQaFWZsYXR0ZW4vUmVzaGFwZS9z" +
		"aGFwZSoHCgFUEgIwBCoMCgZUc2hhcGUSAjADCjIKDGZsYXR0ZW5JbWFnZRIISWRlbnRpdHkaD2ZsYXR0" +
		"ZW4vUmVzaGFwZSoHCgFUEgIwBAo/Cg9mbGF0dGVuXzEvU2hhcGUSBVNoYXBlGgxteUlucHV0U2xpY2Uq" +
		"BwoBVBICMAQqDgoIb3V0X3R5cGUSAjADCksKHWZsYXR0ZW5fMS9zdHJpZGVkX3NsaWNlL3N0YWNrEgVD" +
		"b25zdCoLCgVkdHlwZRICMAMqFgoFdmFsdWUSDUILCAMSBBICCAE6AQAKTQofZmxhdHRlbl8xL3N0cmlk" +
		"ZWRfc2xpY2Uvc3RhY2tfMRIFQ29uc3QqFgoFdmFsdWUSDUILCAMSBBICCAE6AQEqCwoFZHR5cGUSAjAD" +
		"Ck0KH2ZsYXR0ZW5fMS9zdHJpZGVkX3NsaWNlL3N0YWNrXzISBUNvbnN0KgsKBWR0eXBlEgIwAyoWCgV2" +
		"YWx1ZRINQgsIAxIEEgIIAToBAQqTAgoXZmxhdHRlbl8xL3N0cmlkZWRfc2xpY2USDFN0cmlkZWRTbGlj" +
		"ZRoPZmxhdHRlbl8xL1NoYXBlGh1mbGF0dGVuXzEvc3RyaWRlZF9zbGljZS9zdGFjaxofZmxhdHRlbl8x" +
		"L3N0cmlkZWRfc2xpY2Uvc3RhY2tfMRofZmxhdHRlbl8xL3N0cmlkZWRfc2xpY2Uvc3RhY2tfMioQCgpi" +
		"ZWdpbl9tYXNrEgIYACoTCg1lbGxpcHNpc19tYXNrEgIYACoTCg1uZXdfYXhpc19tYXNrEgIYACoOCghl" +
		"bmRfbWFzaxICGAAqBwoBVBICMAMqCwoFSW5kZXgSAjADKhYKEHNocmlua19heGlzX21hc2sSAhgBCkwK" +
		"GWZsYXR0ZW5fMS9SZXNoYXBlL3NoYXBlLzESBUNvbnN0KhsKBXZhbHVlEhJCEAgDEgA6Cv//////////" +
		"/wEqCwoFZHR5cGUSAjADCnEKF2ZsYXR0ZW5fMS9SZXNoYXBlL3NoYXBlEgRQYWNrGhdmbGF0dGVuXzEv" +
		"c3RyaWRlZF9zbGljZRoZZmxhdHRlbl8xL1Jlc2hhcGUvc2hhcGUvMSoHCgFUEgIwAyoKCgRheGlzEgIY" +
		"ACoHCgFOEgIYAgpaChFmbGF0dGVuXzEvUmVzaGFwZRIHUmVzaGFwZRoMbXlJbnB1dFNsaWNlGhdmbGF0" +
		"dGVuXzEvUmVzaGFwZS9zaGFwZSoHCgFUEgIwBCoMCgZUc2hhcGUSAjADCjQKDGZsYXR0ZW5TbGljZRII" +
		"SWRlbnRpdHkaEWZsYXR0ZW5fMS9SZXNoYXBlKgcKAVQSAjAEIgIIGw=="
)

package mat

const (
	// model is base64 representation of protocol buffer TF model.
	// this go file was auto-generated using following command:
	//      go get github.com/sdeoras/tensorflow/cmd/mo
	//      mo u -h
	modelPB = "" +
		"CjUKB3ZlcnNpb24SBUNvbnN0KhYKBXZhbHVlEg1CCwgHEgBCBTAuMS4wKgsKBWR0eXBlEgIwBwowCgVi" +
		"dWZmMRILUGxhY2Vob2xkZXIqCwoFZHR5cGUSAjACKg0KBXNoYXBlEgQ6AhgBCjEKBnNoYXBlMRILUGxh" +
		"Y2Vob2xkZXIqCwoFZHR5cGUSAjAJKg0KBXNoYXBlEgQ6AhgBCjAKBWJ1ZmYyEgtQbGFjZWhvbGRlcioN" +
		"CgVzaGFwZRIEOgIYASoLCgVkdHlwZRICMAIKMQoGc2hhcGUyEgtQbGFjZWhvbGRlcioLCgVkdHlwZRIC" +
		"MAkqDQoFc2hhcGUSBDoCGAEKNQoKc2hhcGVCZWdpbhILUGxhY2Vob2xkZXIqCwoFZHR5cGUSAjAJKg0K" +
		"BXNoYXBlEgQ6AhgBCjQKCXNoYXBlU2l6ZRILUGxhY2Vob2xkZXIqCwoFZHR5cGUSAjAJKg0KBXNoYXBl" +
		"EgQ6AhgBCjgKB1Jlc2hhcGUSB1Jlc2hhcGUaBWJ1ZmYxGgZzaGFwZTEqBwoBVBICMAIqDAoGVHNoYXBl" +
		"EgIwCQo/Cg1NYXRyaXhJbnZlcnNlEg1NYXRyaXhJbnZlcnNlGgdSZXNoYXBlKgcKAVQSAjACKg0KB2Fk" +
		"am9pbnQSAigACkAKCWludi9zaGFwZRIFQ29uc3QqCwoFZHR5cGUSAjADKh8KBXZhbHVlEhZCFAgDEgQS" +
		"AggBOgr///////////8BCj8KA2ludhIHUmVzaGFwZRoNTWF0cml4SW52ZXJzZRoJaW52L3NoYXBlKgcK" +
		"AVQSAjACKgwKBlRzaGFwZRICMAMKOgoJUmVzaGFwZV8xEgdSZXNoYXBlGgVidWZmMRoGc2hhcGUxKgcK" +
		"AVQSAjACKgwKBlRzaGFwZRICMAkKMQoVbWF0cml4X3RyYW5zcG9zZS9SYW5rEgRSYW5rGglSZXNoYXBl" +
		"XzEqBwoBVBICMAIKQAoWbWF0cml4X3RyYW5zcG9zZS9zdWIveRIFQ29uc3QqEgoFdmFsdWUSCUIHCAMS" +
		"ADoBAioLCgVkdHlwZRICMAMKUwoUbWF0cml4X3RyYW5zcG9zZS9zdWISA1N1YhoVbWF0cml4X3RyYW5z" +
		"cG9zZS9SYW5rGhZtYXRyaXhfdHJhbnNwb3NlL3N1Yi95KgcKAVQSAjADCkYKHG1hdHJpeF90cmFuc3Bv" +
		"c2UvUmFuZ2Uvc3RhcnQSBUNvbnN0KhIKBXZhbHVlEglCBwgDEgA6AQAqCwoFZHR5cGUSAjADCkYKHG1h" +
		"dHJpeF90cmFuc3Bvc2UvUmFuZ2UvZGVsdGESBUNvbnN0KhIKBXZhbHVlEglCBwgDEgA6AQEqCwoFZHR5" +
		"cGUSAjADCn0KFm1hdHJpeF90cmFuc3Bvc2UvUmFuZ2USBVJhbmdlGhxtYXRyaXhfdHJhbnNwb3NlL1Jh" +
		"bmdlL3N0YXJ0GhRtYXRyaXhfdHJhbnNwb3NlL3N1YhocbWF0cml4X3RyYW5zcG9zZS9SYW5nZS9kZWx0" +
		"YSoKCgRUaWR4EgIwAwpCChhtYXRyaXhfdHJhbnNwb3NlL3N1Yl8xL3kSBUNvbnN0KhIKBXZhbHVlEglC" +
		"BwgDEgA6AQEqCwoFZHR5cGUSAjADClcKFm1hdHJpeF90cmFuc3Bvc2Uvc3ViXzESA1N1YhoVbWF0cml4" +
		"X3RyYW5zcG9zZS9SYW5rGhhtYXRyaXhfdHJhbnNwb3NlL3N1Yl8xL3kqBwoBVBICMAMKQgoYbWF0cml4" +
		"X3RyYW5zcG9zZS9zdWJfMi95EgVDb25zdCoLCgVkdHlwZRICMAMqEgoFdmFsdWUSCUIHCAMSADoBAgpX" +
		"ChZtYXRyaXhfdHJhbnNwb3NlL3N1Yl8yEgNTdWIaFW1hdHJpeF90cmFuc3Bvc2UvUmFuaxoYbWF0cml4" +
		"X3RyYW5zcG9zZS9zdWJfMi95KgcKAVQSAjADCnYKIG1hdHJpeF90cmFuc3Bvc2UvY29uY2F0L3ZhbHVl" +
		"c18xEgRQYWNrGhZtYXRyaXhfdHJhbnNwb3NlL3N1Yl8xGhZtYXRyaXhfdHJhbnNwb3NlL3N1Yl8yKgcK" +
		"AU4SAhgCKgcKAVQSAjADKgoKBGF4aXMSAhgACkYKHG1hdHJpeF90cmFuc3Bvc2UvY29uY2F0L2F4aXMS" +
		"BUNvbnN0KhIKBXZhbHVlEglCBwgDEgA6AQAqCwoFZHR5cGUSAjADCpkBChdtYXRyaXhfdHJhbnNwb3Nl" +
		"L2NvbmNhdBIIQ29uY2F0VjIaFm1hdHJpeF90cmFuc3Bvc2UvUmFuZ2UaIG1hdHJpeF90cmFuc3Bvc2Uv" +
		"Y29uY2F0L3ZhbHVlc18xGhxtYXRyaXhfdHJhbnNwb3NlL2NvbmNhdC9heGlzKgcKAU4SAhgCKgoKBFRp" +
		"ZHgSAjADKgcKAVQSAjADCmEKGm1hdHJpeF90cmFuc3Bvc2UvdHJhbnNwb3NlEglUcmFuc3Bvc2UaCVJl" +
		"c2hhcGVfMRoXbWF0cml4X3RyYW5zcG9zZS9jb25jYXQqCwoFVHBlcm0SAjADKgcKAVQSAjACCkgKEXRy" +
		"YW5zcG9zZU9wL3NoYXBlEgVDb25zdCofCgV2YWx1ZRIWQhQIAxIEEgIIAToK////////////ASoLCgVk" +
		"dHlwZRICMAMKXAoLdHJhbnNwb3NlT3ASB1Jlc2hhcGUaGm1hdHJpeF90cmFuc3Bvc2UvdHJhbnNwb3Nl" +
		"GhF0cmFuc3Bvc2VPcC9zaGFwZSoHCgFUEgIwAioMCgZUc2hhcGUSAjADCjoKCVJlc2hhcGVfMhIHUmVz" +
		"aGFwZRoFYnVmZjEaBnNoYXBlMSoHCgFUEgIwAioMCgZUc2hhcGUSAjAJCjEKAlFyEgJRchoJUmVzaGFw" +
		"ZV8yKhMKDWZ1bGxfbWF0cmljZXMSAigAKgcKAVQSAjACCkYKD1Jlc2hhcGVfMy9zaGFwZRIFQ29uc3Qq" +
		"CwoFZHR5cGUSAjADKh8KBXZhbHVlEhZCFAgDEgQSAggBOgr///////////8BCkAKCVJlc2hhcGVfMxIH" +
		"UmVzaGFwZRoCUXIaD1Jlc2hhcGVfMy9zaGFwZSoHCgFUEgIwAioMCgZUc2hhcGUSAjADCkYKD1Jlc2hh" +
		"cGVfNC9zaGFwZRIFQ29uc3QqHwoFdmFsdWUSFkIUCAMSBBICCAE6Cv///////////wEqCwoFZHR5cGUS" +
		"AjADCkIKCVJlc2hhcGVfNBIHUmVzaGFwZRoEUXI6MRoPUmVzaGFwZV80L3NoYXBlKgcKAVQSAjACKgwK" +
		"BlRzaGFwZRICMAMKKgoKcXJkZWNvbXBfcRIISWRlbnRpdHkaCVJlc2hhcGVfMyoHCgFUEgIwAgoqCgpx" +
		"cmRlY29tcF9yEghJZGVudGl0eRoJUmVzaGFwZV80KgcKAVQSAjACCkoKE3plcm9zL1Jlc2hhcGUvc2hh" +
		"cGUSBUNvbnN0KgsKBWR0eXBlEgIwAyofCgV2YWx1ZRIWQhQIAxIEEgIIAToK////////////AQpMCg16" +
		"ZXJvcy9SZXNoYXBlEgdSZXNoYXBlGgZzaGFwZTEaE3plcm9zL1Jlc2hhcGUvc2hhcGUqBwoBVBICMAkq" +
		"DAoGVHNoYXBlEgIwAwo8Cgt6ZXJvcy9Db25zdBIFQ29uc3QqGQoFdmFsdWUSEEIOCAISADIIAAAAAAAA" +
		"AAAqCwoFZHR5cGUSAjACCkQKBXplcm9zEgRGaWxsGg16ZXJvcy9SZXNoYXBlGgt6ZXJvcy9Db25zdCoH" +
		"CgFUEgIwAioQCgppbmRleF90eXBlEgIwCQpECg16ZXJvc18xL3NoYXBlEgVDb25zdCofCgV2YWx1ZRIW" +
		"QhQIAxIEEgIIAToK////////////ASoLCgVkdHlwZRICMAMKPwoHemVyb3NfMRIHUmVzaGFwZRoFemVy" +
		"b3MaDXplcm9zXzEvc2hhcGUqBwoBVBICMAIqDAoGVHNoYXBlEgIwAwpJChJvbmVzL1Jlc2hhcGUvc2hh" +
		"cGUSBUNvbnN0Kh8KBXZhbHVlEhZCFAgDEgQSAggBOgr///////////8BKgsKBWR0eXBlEgIwAwpKCgxv" +
		"bmVzL1Jlc2hhcGUSB1Jlc2hhcGUaBnNoYXBlMRoSb25lcy9SZXNoYXBlL3NoYXBlKgcKAVQSAjAJKgwK" +
		"BlRzaGFwZRICMAMKOwoKb25lcy9Db25zdBIFQ29uc3QqCwoFZHR5cGUSAjACKhkKBXZhbHVlEhBCDggC" +
		"EgAyCAAAAAAAAPA/CkEKBG9uZXMSBEZpbGwaDG9uZXMvUmVzaGFwZRoKb25lcy9Db25zdCoHCgFUEgIw" +
		"AioQCgppbmRleF90eXBlEgIwCQpDCgxvbmVzXzEvc2hhcGUSBUNvbnN0KgsKBWR0eXBlEgIwAyofCgV2" +
		"YWx1ZRIWQhQIAxIEEgIIAToK////////////AQo8CgZvbmVzXzESB1Jlc2hhcGUaBG9uZXMaDG9uZXNf" +
		"MS9zaGFwZSoHCgFUEgIwAioMCgZUc2hhcGUSAjADCkMKEnJhbmRvbV91bmlmb3JtL21pbhIFQ29uc3Qq" +
		"GQoFdmFsdWUSEEIOCAISADIIAAAAAAAAAAAqCwoFZHR5cGUSAjACCkMKEnJhbmRvbV91bmlmb3JtL21h" +
		"eBIFQ29uc3QqGQoFdmFsdWUSEEIOCAISADIIAAAAAAAA8D8qCwoFZHR5cGUSAjACCmQKHHJhbmRvbV91" +
		"bmlmb3JtL1JhbmRvbVVuaWZvcm0SDVJhbmRvbVVuaWZvcm0aBnNoYXBlMSoHCgFUEgIwCSoLCgVkdHlw" +
		"ZRICMAIqCwoFc2VlZDISAhgAKgoKBHNlZWQSAhgACkoKEnJhbmRvbV91bmlmb3JtL3N1YhIDU3ViGhJy" +
		"YW5kb21fdW5pZm9ybS9tYXgaEnJhbmRvbV91bmlmb3JtL21pbioHCgFUEgIwAgpUChJyYW5kb21fdW5p" +
		"Zm9ybS9tdWwSA011bBoccmFuZG9tX3VuaWZvcm0vUmFuZG9tVW5pZm9ybRoScmFuZG9tX3VuaWZvcm0v" +
		"c3ViKgcKAVQSAjACCkYKDnJhbmRvbV91bmlmb3JtEgNBZGQaEnJhbmRvbV91bmlmb3JtL211bBoScmFu" +
		"ZG9tX3VuaWZvcm0vbWluKgcKAVQSAjACCkEKCnJhbmQvc2hhcGUSBUNvbnN0Kh8KBXZhbHVlEhZCFAgD" +
		"EgQSAggBOgr///////////8BKgsKBWR0eXBlEgIwAwpCCgRyYW5kEgdSZXNoYXBlGg5yYW5kb21fdW5p" +
		"Zm9ybRoKcmFuZC9zaGFwZSoHCgFUEgIwAioMCgZUc2hhcGUSAjADCkMKEnJhbmRvbV9ub3JtYWwvbWVh" +
		"bhIFQ29uc3QqGQoFdmFsdWUSEEIOCAISADIIAAAAAAAAAAAqCwoFZHR5cGUSAjACCkUKFHJhbmRvbV9u" +
		"b3JtYWwvc3RkZGV2EgVDb25zdCoZCgV2YWx1ZRIQQg4IAhIAMggAAAAAAADwPyoLCgVkdHlwZRICMAIK" +
		"cQoicmFuZG9tX25vcm1hbC9SYW5kb21TdGFuZGFyZE5vcm1hbBIUUmFuZG9tU3RhbmRhcmROb3JtYWwa" +
		"BnNoYXBlMSoHCgFUEgIwCSoLCgVkdHlwZRICMAIqCwoFc2VlZDISAhgAKgoKBHNlZWQSAhgAClsKEXJh" +
		"bmRvbV9ub3JtYWwvbXVsEgNNdWwaInJhbmRvbV9ub3JtYWwvUmFuZG9tU3RhbmRhcmROb3JtYWwaFHJh" +
		"bmRvbV9ub3JtYWwvc3RkZGV2KgcKAVQSAjACCkQKDXJhbmRvbV9ub3JtYWwSA0FkZBoRcmFuZG9tX25v" +
		"cm1hbC9tdWwaEnJhbmRvbV9ub3JtYWwvbWVhbioHCgFUEgIwAgpCCgtyYW5kbi9zaGFwZRIFQ29uc3Qq" +
		"CwoFZHR5cGUSAjADKh8KBXZhbHVlEhZCFAgDEgQSAggBOgr///////////8BCkQKBXJhbmRuEgdSZXNo" +
		"YXBlGg5yYW5kb21fdW5pZm9ybRoLcmFuZG4vc2hhcGUqBwoBVBICMAIqDAoGVHNoYXBlEgIwAwo6CglS" +
		"ZXNoYXBlXzUSB1Jlc2hhcGUaBWJ1ZmYxGgZzaGFwZTEqBwoBVBICMAIqDAoGVHNoYXBlEgIwCQo6CglS" +
		"ZXNoYXBlXzYSB1Jlc2hhcGUaBWJ1ZmYyGgZzaGFwZTIqBwoBVBICMAIqDAoGVHNoYXBlEgIwCQpOCgZN" +
		"YXRNdWwSC0JhdGNoTWF0TXVsGglSZXNoYXBlXzUaCVJlc2hhcGVfNioLCgVhZGpfeBICKAAqCwoFYWRq" +
		"X3kSAigAKgcKAVQSAjACCjIKCG11bFNoYXBlEgVTaGFwZRoGTWF0TXVsKgcKAVQSAjACKg4KCG91dF90" +
		"eXBlEgIwCQpACgltdWwvc2hhcGUSBUNvbnN0Kh8KBXZhbHVlEhZCFAgDEgQSAggBOgr///////////8B" +
		"KgsKBWR0eXBlEgIwAwo4CgNtdWwSB1Jlc2hhcGUaBk1hdE11bBoJbXVsL3NoYXBlKgcKAVQSAjACKgwK" +
		"BlRzaGFwZRICMAMKOgoJUmVzaGFwZV83EgdSZXNoYXBlGgVidWZmMRoGc2hhcGUxKgcKAVQSAjACKgwK" +
		"BlRzaGFwZRICMAkKRgoFU2xpY2USBVNsaWNlGglSZXNoYXBlXzcaCnNoYXBlQmVnaW4aCXNoYXBlU2l6" +
		"ZSoHCgFUEgIwAioLCgVJbmRleBICMAkKLgoFU2hhcGUSBVNoYXBlGgVTbGljZSoHCgFUEgIwAioOCghv" +
		"dXRfdHlwZRICMAkKKAoMc2xpY2VTaGFwZU9wEghJZGVudGl0eRoFU2hhcGUqBwoBVBICMAkKRgoPUmVz" +
		"aGFwZV84L3NoYXBlEgVDb25zdCofCgV2YWx1ZRIWQhQIAxIEEgIIAToK////////////ASoLCgVkdHlw" +
		"ZRICMAMKQwoJUmVzaGFwZV84EgdSZXNoYXBlGgVTbGljZRoPUmVzaGFwZV84L3NoYXBlKgcKAVQSAjAC" +
		"KgwKBlRzaGFwZRICMAMKJwoHc2xpY2VPcBIISWRlbnRpdHkaCVJlc2hhcGVfOCoHCgFUEgIwAgo6CglS" +
		"ZXNoYXBlXzkSB1Jlc2hhcGUaBWJ1ZmYxGgZzaGFwZTEqBwoBVBICMAIqDAoGVHNoYXBlEgIwCQo/CgpS" +
		"ZXNoYXBlXzEwEgdSZXNoYXBlGglSZXNoYXBlXzkaBnNoYXBlMioHCgFUEgIwAioMCgZUc2hhcGUSAjAJ" +
		"CkcKEFJlc2hhcGVfMTEvc2hhcGUSBUNvbnN0Kh8KBXZhbHVlEhZCFAgDEgQSAggBOgr///////////8B" +
		"KgsKBWR0eXBlEgIwAwpKCgpSZXNoYXBlXzExEgdSZXNoYXBlGgpSZXNoYXBlXzEwGhBSZXNoYXBlXzEx" +
		"L3NoYXBlKgcKAVQSAjACKgwKBlRzaGFwZRICMAMKKgoJcmVzaGFwZU9wEghJZGVudGl0eRoKUmVzaGFw" +
		"ZV8xMSoHCgFUEgIwAgo7CgpSZXNoYXBlXzEyEgdSZXNoYXBlGgVidWZmMRoGc2hhcGUxKgcKAVQSAjAC" +
		"KgwKBlRzaGFwZRICMAkKOwoEVGlsZRIEVGlsZRoKUmVzaGFwZV8xMhoGc2hhcGUyKhAKClRtdWx0aXBs" +
		"ZXMSAjAJKgcKAVQSAjACCi8KB1NoYXBlXzESBVNoYXBlGgRUaWxlKgcKAVQSAjACKg4KCG91dF90eXBl" +
		"EgIwCQopCgt0aWxlU2hhcGVPcBIISWRlbnRpdHkaB1NoYXBlXzEqBwoBVBICMAkKRwoQUmVzaGFwZV8x" +
		"My9zaGFwZRIFQ29uc3QqHwoFdmFsdWUSFkIUCAMSBBICCAE6Cv///////////wEqCwoFZHR5cGUSAjAD" +
		"CkQKClJlc2hhcGVfMTMSB1Jlc2hhcGUaBFRpbGUaEFJlc2hhcGVfMTMvc2hhcGUqBwoBVBICMAIqDAoG" +
		"VHNoYXBlEgIwAwonCgZ0aWxlT3ASCElkZW50aXR5GgpSZXNoYXBlXzEzKgcKAVQSAjACIgIIGw=="
)

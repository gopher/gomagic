package gomagic

import (
	"io/ioutil"
	"testing"

	"github.com/gopher/gomagic"
)

func Test_magic_textual_output(t *testing.T) {

	m, err := gomagic.New(gomagic.NoneFlag)
	if err != nil {
		t.Error("Failed to initialize magic library")
	}

	// test decoding of raw text file via ExamineFile API
	fileName := "test_input/test_file.txt"
	result, err := m.ExamineFile(fileName)
	if err != nil {
		t.Errorf("gomagic.ExamineFile: Failed to parse %s with %s", fileName, m.Error())
	}
	expectedResult := "ASCII text"
	if result != expectedResult {
		t.Errorf("gomagic.ExamineFile: Failed to properly decode %s. Expected \"%s\" but go \"%s\"", fileName, expectedResult, result)
	}

	// test decoding of pdf file via ExamineFile API
	fileName = "test_input/test_file.pdf"
	result, err = m.ExamineFile(fileName)
	if err != nil {
		t.Errorf("gomagic.ExamineFile: Failed to parse %s with %s", fileName, m.Error())
	}
	expectedResult = "PDF document, version 1.3"
	if result != expectedResult {
		t.Errorf("gomagic.ExamineFile: Failed to properly decode %s. Expected \"%s\" but go \"%s\"", fileName, expectedResult, result)
	}

	// test decoding of raw text file via ExamineBuffer API
	fileName = "test_input/test_file.txt"
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Errorf("ioutil.ReadFile: Failed to parse file %s with %s", fileName, m.Error())
	}
	result, err = m.ExamineBuffer(buf)
	if err != nil {
		t.Errorf("gomagic.ExamineBuffer: Failed to parse %s buffer with %s", fileName, m.Error())
	}
	expectedResult = "ASCII text"
	if result != expectedResult {
		t.Errorf("gomagic.ExamineBuffer: Failed to properly decode %s. Expected \"%s\" but go \"%s\"", fileName, expectedResult, result)
	}

	// test decoding of raw text file via ExamineBuffer API
	fileName = "test_input/test_file.pdf"
	buf, err = ioutil.ReadFile(fileName)
	if err != nil {
		t.Errorf("ioutil.ReadFile: Failed to parse file %s with %s", fileName, m.Error())
	}
	result, err = m.ExamineBuffer(buf)
	if err != nil {
		t.Errorf("gomagic.ExamineBuffer: Failed to parse %s buffer with %s", fileName, m.Error())
	}
	expectedResult = "PDF document, version 1.3"
	if result != expectedResult {
		t.Errorf("gomagic.ExamineBuffer: Failed to properly decode %s. Expected \"%s\" but go \"%s\"", fileName, expectedResult, result)
	}
}

func Test_magic_mime_output(t *testing.T) {

	m, err := gomagic.New(gomagic.NoneFlag)
	if err != nil {
		t.Error("Failed to initialize magic library")
	}

	err = m.SetFlags(gomagic.NodescFlag)
	if err != nil {
		t.Errorf("gomagic.SetFlags: Failed to set magic.NodescFlag")
	}

	// test decoding of raw text file via ExamineFile API
	fileName := "test_input/test_file.txt"
	result, err := m.ExamineFile(fileName)
	if err != nil {
		t.Errorf("gomagic.ExamineFile: Failed to parse %s with %s", fileName, m.Error())
	}
	expectedResult := "application/octet-stream; charset=us-ascii"
	if result != expectedResult {
		t.Errorf("gomagic.ExamineFile: Failed to properly decode %s. Expected \"%s\" but go \"%s\"", fileName, expectedResult, result)
	}

	// test decoding of pdf file via ExamineFile API
	fileName = "test_input/test_file.pdf"
	result, err = m.ExamineFile(fileName)
	if err != nil {
		t.Errorf("gomagic.ExamineFile: Failed to parse %s with %s", fileName, m.Error())
	}
	expectedResult = "application/pdf; charset=binary"
	if result != expectedResult {
		t.Errorf("gomagic.ExamineFile: Failed to properly decode %s. Expected \"%s\" but go \"%s\"", fileName, expectedResult, result)
	}

	// test decoding of raw text file via ExamineBuffer API
	fileName = "test_input/test_file.txt"
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Errorf("ioutil.ReadFile: Failed to parse file %s with %s", fileName, m.Error())
	}
	result, err = m.ExamineBuffer(buf)
	if err != nil {
		t.Errorf("gomagic.ExamineBuffer: Failed to parse %s buffer with %s", fileName, m.Error())
	}
	expectedResult = "application/octet-stream; charset=us-ascii"
	if result != expectedResult {
		t.Errorf("gomagic.ExamineBuffer: Failed to properly decode %s. Expected \"%s\" but go \"%s\"", fileName, expectedResult, result)
	}

	// test decoding of pdf file via ExamineBuffer API
	fileName = "test_input/test_file.pdf"
	buf, err = ioutil.ReadFile(fileName)
	if err != nil {
		t.Errorf("ioutil.ReadFile: Failed to parse file %s with %s", fileName, m.Error())
	}
	result, err = m.ExamineBuffer(buf)
	if err != nil {
		t.Errorf("gomagic.ExamineBuffer: Failed to parse %s buffer with %s", fileName, m.Error())
	}
	expectedResult = "application/pdf; charset=binary"
	if result != expectedResult {
		t.Errorf("gomagic.ExamineBuffer: Failed to properly decode %s. Expected \"%s\" but go \"%s\"", fileName, expectedResult, result)
	}
}

func Test_magic_settings(t *testing.T) {

	m, err := gomagic.New(gomagic.NoneFlag)
	if err != nil {
		t.Error("Failed to initialize magic library")
	}

	fileName := "test_input/test_db.mgc"
	err = m.Load(fileName)
	if err != nil {
		t.Errorf("gomagic.Load: Failed to load database file %s with %s", fileName, m.Error())
	}

	bytesMax := 50
	err = m.Setparam(gomagic.ParamBytesMax, bytesMax)
	if err != nil {
		t.Errorf("gomagic.Setparam: Failed to set parameter magic.ParamBytesMax to %d with %s", bytesMax, m.Error())
	}

	val, err := m.Getparam(gomagic.ParamBytesMax)
	if err != nil {
		t.Errorf("gomagic.Getparam: Failed to get parameter magic.ParamBytesMax with %s", m.Error())
	}
	if val != bytesMax {
		t.Errorf("gomagic.Getparam: Retrieved value for magic.ParamBytesMax is incorrect. Expected \"%d\" but go \"%d\"", bytesMax, val)
	}
}

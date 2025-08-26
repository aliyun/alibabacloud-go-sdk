// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeVINCodeAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeVINCodeAdvanceRequest
	GetImageURLObject() io.Reader
}

type RecognizeVINCodeAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeVINCode/vin1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVINCodeAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVINCodeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeVINCodeAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeVINCodeAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeVINCodeAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

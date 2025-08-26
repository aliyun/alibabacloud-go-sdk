// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeLicensePlateAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeLicensePlateAdvanceRequest
	GetImageURLObject() io.Reader
}

type RecognizeLicensePlateAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeLicensePlate/cpsb1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeLicensePlateAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeLicensePlateAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeLicensePlateAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeLicensePlateAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeLicensePlateAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

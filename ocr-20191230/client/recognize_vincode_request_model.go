// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVINCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeVINCodeRequest
	GetImageURL() *string
}

type RecognizeVINCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeVINCode/vin1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVINCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVINCodeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeVINCodeRequest) SetImageURL(v string) *RecognizeVINCodeRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeVINCodeRequest) Validate() error {
	return dara.Validate(s)
}

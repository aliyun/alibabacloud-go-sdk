// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeLicensePlateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeLicensePlateRequest
	GetImageURL() *string
}

type RecognizeLicensePlateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeLicensePlate/cpsb1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeLicensePlateRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeLicensePlateRequest) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeLicensePlateRequest) SetImageURL(v string) *RecognizeLicensePlateRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeLicensePlateRequest) Validate() error {
	return dara.Validate(s)
}

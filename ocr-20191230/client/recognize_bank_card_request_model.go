// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeBankCardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeBankCardRequest
	GetImageURL() *string
}

type RecognizeBankCardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeBankCard/yhk3.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeBankCardRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBankCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeBankCardRequest) SetImageURL(v string) *RecognizeBankCardRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeBankCardRequest) Validate() error {
	return dara.Validate(s)
}

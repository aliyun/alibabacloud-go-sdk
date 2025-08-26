// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeIdentityCardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeIdentityCardRequest
	GetImageURL() *string
	SetSide(v string) *RecognizeIdentityCardRequest
	GetSide() *string
}

type RecognizeIdentityCardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeIdentityCard/sfz1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// face
	Side *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeIdentityCardRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeIdentityCardRequest) GetSide() *string {
	return s.Side
}

func (s *RecognizeIdentityCardRequest) SetImageURL(v string) *RecognizeIdentityCardRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeIdentityCardRequest) SetSide(v string) *RecognizeIdentityCardRequest {
	s.Side = &v
	return s
}

func (s *RecognizeIdentityCardRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeIdentityCardAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeIdentityCardAdvanceRequest
	GetImageURLObject() io.Reader
	SetSide(v string) *RecognizeIdentityCardAdvanceRequest
	GetSide() *string
}

type RecognizeIdentityCardAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeIdentityCard/sfz1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// face
	Side *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeIdentityCardAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeIdentityCardAdvanceRequest) GetSide() *string {
	return s.Side
}

func (s *RecognizeIdentityCardAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeIdentityCardAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeIdentityCardAdvanceRequest) SetSide(v string) *RecognizeIdentityCardAdvanceRequest {
	s.Side = &v
	return s
}

func (s *RecognizeIdentityCardAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

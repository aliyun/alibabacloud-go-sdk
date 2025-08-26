// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeDriverLicenseAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeDriverLicenseAdvanceRequest
	GetImageURLObject() io.Reader
	SetSide(v string) *RecognizeDriverLicenseAdvanceRequest
	GetSide() *string
}

type RecognizeDriverLicenseAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeDriverLicense/jsz2.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// face
	Side *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeDriverLicenseAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDriverLicenseAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeDriverLicenseAdvanceRequest) GetSide() *string {
	return s.Side
}

func (s *RecognizeDriverLicenseAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeDriverLicenseAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeDriverLicenseAdvanceRequest) SetSide(v string) *RecognizeDriverLicenseAdvanceRequest {
	s.Side = &v
	return s
}

func (s *RecognizeDriverLicenseAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

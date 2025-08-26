// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeDriverLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeDriverLicenseRequest
	GetImageURL() *string
	SetSide(v string) *RecognizeDriverLicenseRequest
	GetSide() *string
}

type RecognizeDriverLicenseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeDriverLicense/jsz2.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// face
	Side *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeDriverLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDriverLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeDriverLicenseRequest) GetSide() *string {
	return s.Side
}

func (s *RecognizeDriverLicenseRequest) SetImageURL(v string) *RecognizeDriverLicenseRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeDriverLicenseRequest) SetSide(v string) *RecognizeDriverLicenseRequest {
	s.Side = &v
	return s
}

func (s *RecognizeDriverLicenseRequest) Validate() error {
	return dara.Validate(s)
}

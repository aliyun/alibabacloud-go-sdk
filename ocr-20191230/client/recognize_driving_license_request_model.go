// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeDrivingLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeDrivingLicenseRequest
	GetImageURL() *string
	SetSide(v string) *RecognizeDrivingLicenseRequest
	GetSide() *string
}

type RecognizeDrivingLicenseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeDrivingLicense/xsz2.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// face
	Side *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeDrivingLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDrivingLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeDrivingLicenseRequest) GetSide() *string {
	return s.Side
}

func (s *RecognizeDrivingLicenseRequest) SetImageURL(v string) *RecognizeDrivingLicenseRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeDrivingLicenseRequest) SetSide(v string) *RecognizeDrivingLicenseRequest {
	s.Side = &v
	return s
}

func (s *RecognizeDrivingLicenseRequest) Validate() error {
	return dara.Validate(s)
}

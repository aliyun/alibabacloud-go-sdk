// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeBusinessLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeBusinessLicenseRequest
	GetImageURL() *string
}

type RecognizeBusinessLicenseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeBusinessLicense/RecognizeBusinessLicense1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeBusinessLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBusinessLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeBusinessLicenseRequest) SetImageURL(v string) *RecognizeBusinessLicenseRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeBusinessLicenseRequest) Validate() error {
	return dara.Validate(s)
}

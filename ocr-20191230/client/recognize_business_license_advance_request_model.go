// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeBusinessLicenseAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeBusinessLicenseAdvanceRequest
	GetImageURLObject() io.Reader
}

type RecognizeBusinessLicenseAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeBusinessLicense/RecognizeBusinessLicense1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeBusinessLicenseAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBusinessLicenseAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeBusinessLicenseAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeBusinessLicenseAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeBusinessLicenseAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

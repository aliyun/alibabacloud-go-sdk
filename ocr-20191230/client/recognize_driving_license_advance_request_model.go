// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeDrivingLicenseAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeDrivingLicenseAdvanceRequest
	GetImageURLObject() io.Reader
	SetSide(v string) *RecognizeDrivingLicenseAdvanceRequest
	GetSide() *string
}

type RecognizeDrivingLicenseAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeDrivingLicense/xsz2.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// face
	Side *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeDrivingLicenseAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDrivingLicenseAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeDrivingLicenseAdvanceRequest) GetSide() *string {
	return s.Side
}

func (s *RecognizeDrivingLicenseAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeDrivingLicenseAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeDrivingLicenseAdvanceRequest) SetSide(v string) *RecognizeDrivingLicenseAdvanceRequest {
	s.Side = &v
	return s
}

func (s *RecognizeDrivingLicenseAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

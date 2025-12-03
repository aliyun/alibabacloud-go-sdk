// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ActivateLicenseRequest
	GetBizId() *string
	SetBizType(v string) *ActivateLicenseRequest
	GetBizType() *string
	SetLicenseCode(v string) *ActivateLicenseRequest
	GetLicenseCode() *string
	SetLicenseNo(v string) *ActivateLicenseRequest
	GetLicenseNo() *string
	SetLicensePublisher(v string) *ActivateLicenseRequest
	GetLicensePublisher() *string
}

type ActivateLicenseRequest struct {
	// example:
	//
	// P20211118170645000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esp.bookkeeping_course
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// B09YICKLVHNR7ZQR
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	LicenseNo   *string `json:"LicenseNo,omitempty" xml:"LicenseNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yunMarket
	LicensePublisher *string `json:"LicensePublisher,omitempty" xml:"LicensePublisher,omitempty"`
}

func (s ActivateLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateLicenseRequest) GoString() string {
	return s.String()
}

func (s *ActivateLicenseRequest) GetBizId() *string {
	return s.BizId
}

func (s *ActivateLicenseRequest) GetBizType() *string {
	return s.BizType
}

func (s *ActivateLicenseRequest) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *ActivateLicenseRequest) GetLicenseNo() *string {
	return s.LicenseNo
}

func (s *ActivateLicenseRequest) GetLicensePublisher() *string {
	return s.LicensePublisher
}

func (s *ActivateLicenseRequest) SetBizId(v string) *ActivateLicenseRequest {
	s.BizId = &v
	return s
}

func (s *ActivateLicenseRequest) SetBizType(v string) *ActivateLicenseRequest {
	s.BizType = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicenseCode(v string) *ActivateLicenseRequest {
	s.LicenseCode = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicenseNo(v string) *ActivateLicenseRequest {
	s.LicenseNo = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicensePublisher(v string) *ActivateLicenseRequest {
	s.LicensePublisher = &v
	return s
}

func (s *ActivateLicenseRequest) Validate() error {
	return dara.Validate(s)
}

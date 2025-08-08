// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentification(v string) *ActivateLicenseRequest
	GetIdentification() *string
	SetLicenseCode(v string) *ActivateLicenseRequest
	GetLicenseCode() *string
}

type ActivateLicenseRequest struct {
	// example:
	//
	// 129****1111
	Identification *string `json:"Identification,omitempty" xml:"Identification,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// APSEDH8TA5CSGK-********_6CNTACBH9EQPOATFXJQL4B2COE7M43VVQ7GUGKAA
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
}

func (s ActivateLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateLicenseRequest) GoString() string {
	return s.String()
}

func (s *ActivateLicenseRequest) GetIdentification() *string {
	return s.Identification
}

func (s *ActivateLicenseRequest) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *ActivateLicenseRequest) SetIdentification(v string) *ActivateLicenseRequest {
	s.Identification = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicenseCode(v string) *ActivateLicenseRequest {
	s.LicenseCode = &v
	return s
}

func (s *ActivateLicenseRequest) Validate() error {
	return dara.Validate(s)
}

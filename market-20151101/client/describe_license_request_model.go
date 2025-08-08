// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseCode(v string) *DescribeLicenseRequest
	GetLicenseCode() *string
}

type DescribeLicenseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
}

func (s DescribeLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseRequest) GoString() string {
	return s.String()
}

func (s *DescribeLicenseRequest) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *DescribeLicenseRequest) SetLicenseCode(v string) *DescribeLicenseRequest {
	s.LicenseCode = &v
	return s
}

func (s *DescribeLicenseRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicenseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseId(v string) *GetLicenseInfoRequest
	GetLicenseId() *string
}

type GetLicenseInfoRequest struct {
	// This parameter is required.
	LicenseId *string `json:"LicenseId,omitempty" xml:"LicenseId,omitempty"`
}

func (s GetLicenseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseInfoRequest) GoString() string {
	return s.String()
}

func (s *GetLicenseInfoRequest) GetLicenseId() *string {
	return s.LicenseId
}

func (s *GetLicenseInfoRequest) SetLicenseId(v string) *GetLicenseInfoRequest {
	s.LicenseId = &v
	return s
}

func (s *GetLicenseInfoRequest) Validate() error {
	return dara.Validate(s)
}

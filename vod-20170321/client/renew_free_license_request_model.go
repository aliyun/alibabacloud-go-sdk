// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewFreeLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppItemId(v string) *RenewFreeLicenseRequest
	GetAppItemId() *string
	SetLicenseItemId(v string) *RenewFreeLicenseRequest
	GetLicenseItemId() *string
	SetValidPeriod(v int32) *RenewFreeLicenseRequest
	GetValidPeriod() *int32
}

type RenewFreeLicenseRequest struct {
	// This parameter is required.
	AppItemId *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
	// This parameter is required.
	LicenseItemId *string `json:"LicenseItemId,omitempty" xml:"LicenseItemId,omitempty"`
	// This parameter is required.
	ValidPeriod *int32 `json:"ValidPeriod,omitempty" xml:"ValidPeriod,omitempty"`
}

func (s RenewFreeLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewFreeLicenseRequest) GoString() string {
	return s.String()
}

func (s *RenewFreeLicenseRequest) GetAppItemId() *string {
	return s.AppItemId
}

func (s *RenewFreeLicenseRequest) GetLicenseItemId() *string {
	return s.LicenseItemId
}

func (s *RenewFreeLicenseRequest) GetValidPeriod() *int32 {
	return s.ValidPeriod
}

func (s *RenewFreeLicenseRequest) SetAppItemId(v string) *RenewFreeLicenseRequest {
	s.AppItemId = &v
	return s
}

func (s *RenewFreeLicenseRequest) SetLicenseItemId(v string) *RenewFreeLicenseRequest {
	s.LicenseItemId = &v
	return s
}

func (s *RenewFreeLicenseRequest) SetValidPeriod(v int32) *RenewFreeLicenseRequest {
	s.ValidPeriod = &v
	return s
}

func (s *RenewFreeLicenseRequest) Validate() error {
	return dara.Validate(s)
}

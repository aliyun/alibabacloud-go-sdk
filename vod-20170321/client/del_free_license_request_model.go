// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelFreeLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppItemId(v string) *DelFreeLicenseRequest
	GetAppItemId() *string
	SetLicenseItemId(v string) *DelFreeLicenseRequest
	GetLicenseItemId() *string
}

type DelFreeLicenseRequest struct {
	// This parameter is required.
	AppItemId *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
	// This parameter is required.
	LicenseItemId *string `json:"LicenseItemId,omitempty" xml:"LicenseItemId,omitempty"`
}

func (s DelFreeLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s DelFreeLicenseRequest) GoString() string {
	return s.String()
}

func (s *DelFreeLicenseRequest) GetAppItemId() *string {
	return s.AppItemId
}

func (s *DelFreeLicenseRequest) GetLicenseItemId() *string {
	return s.LicenseItemId
}

func (s *DelFreeLicenseRequest) SetAppItemId(v string) *DelFreeLicenseRequest {
	s.AppItemId = &v
	return s
}

func (s *DelFreeLicenseRequest) SetLicenseItemId(v string) *DelFreeLicenseRequest {
	s.LicenseItemId = &v
	return s
}

func (s *DelFreeLicenseRequest) Validate() error {
	return dara.Validate(s)
}

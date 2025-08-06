// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFreeLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppItemId(v string) *DeleteFreeLicenseRequest
	GetAppItemId() *string
	SetLicenseItemId(v string) *DeleteFreeLicenseRequest
	GetLicenseItemId() *string
}

type DeleteFreeLicenseRequest struct {
	// This parameter is required.
	AppItemId *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
	// This parameter is required.
	LicenseItemId *string `json:"LicenseItemId,omitempty" xml:"LicenseItemId,omitempty"`
}

func (s DeleteFreeLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFreeLicenseRequest) GoString() string {
	return s.String()
}

func (s *DeleteFreeLicenseRequest) GetAppItemId() *string {
	return s.AppItemId
}

func (s *DeleteFreeLicenseRequest) GetLicenseItemId() *string {
	return s.LicenseItemId
}

func (s *DeleteFreeLicenseRequest) SetAppItemId(v string) *DeleteFreeLicenseRequest {
	s.AppItemId = &v
	return s
}

func (s *DeleteFreeLicenseRequest) SetLicenseItemId(v string) *DeleteFreeLicenseRequest {
	s.LicenseItemId = &v
	return s
}

func (s *DeleteFreeLicenseRequest) Validate() error {
	return dara.Validate(s)
}

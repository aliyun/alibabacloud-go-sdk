// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppItemId(v string) *DeleteAppLicenseRequest
	GetAppItemId() *string
	SetLicenseItemIds(v string) *DeleteAppLicenseRequest
	GetLicenseItemIds() *string
}

type DeleteAppLicenseRequest struct {
	AppItemId      *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
	LicenseItemIds *string `json:"LicenseItemIds,omitempty" xml:"LicenseItemIds,omitempty"`
}

func (s DeleteAppLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppLicenseRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppLicenseRequest) GetAppItemId() *string {
	return s.AppItemId
}

func (s *DeleteAppLicenseRequest) GetLicenseItemIds() *string {
	return s.LicenseItemIds
}

func (s *DeleteAppLicenseRequest) SetAppItemId(v string) *DeleteAppLicenseRequest {
	s.AppItemId = &v
	return s
}

func (s *DeleteAppLicenseRequest) SetLicenseItemIds(v string) *DeleteAppLicenseRequest {
	s.LicenseItemIds = &v
	return s
}

func (s *DeleteAppLicenseRequest) Validate() error {
	return dara.Validate(s)
}

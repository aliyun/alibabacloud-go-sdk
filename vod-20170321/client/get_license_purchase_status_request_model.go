// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicensePurchaseStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseItemIds(v string) *GetLicensePurchaseStatusRequest
	GetLicenseItemIds() *string
}

type GetLicensePurchaseStatusRequest struct {
	LicenseItemIds *string `json:"LicenseItemIds,omitempty" xml:"LicenseItemIds,omitempty"`
}

func (s GetLicensePurchaseStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLicensePurchaseStatusRequest) GoString() string {
	return s.String()
}

func (s *GetLicensePurchaseStatusRequest) GetLicenseItemIds() *string {
	return s.LicenseItemIds
}

func (s *GetLicensePurchaseStatusRequest) SetLicenseItemIds(v string) *GetLicensePurchaseStatusRequest {
	s.LicenseItemIds = &v
	return s
}

func (s *GetLicensePurchaseStatusRequest) Validate() error {
	return dara.Validate(s)
}

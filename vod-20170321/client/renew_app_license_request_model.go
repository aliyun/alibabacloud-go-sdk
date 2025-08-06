// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLicenseItemIds(v string) *RenewAppLicenseRequest
	GetLicenseItemIds() *string
	SetOrderIds(v string) *RenewAppLicenseRequest
	GetOrderIds() *string
	SetPurchaseMethod(v string) *RenewAppLicenseRequest
	GetPurchaseMethod() *string
}

type RenewAppLicenseRequest struct {
	LicenseItemIds *string `json:"LicenseItemIds,omitempty" xml:"LicenseItemIds,omitempty"`
	OrderIds       *string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty"`
	PurchaseMethod *string `json:"PurchaseMethod,omitempty" xml:"PurchaseMethod,omitempty"`
}

func (s RenewAppLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewAppLicenseRequest) GoString() string {
	return s.String()
}

func (s *RenewAppLicenseRequest) GetLicenseItemIds() *string {
	return s.LicenseItemIds
}

func (s *RenewAppLicenseRequest) GetOrderIds() *string {
	return s.OrderIds
}

func (s *RenewAppLicenseRequest) GetPurchaseMethod() *string {
	return s.PurchaseMethod
}

func (s *RenewAppLicenseRequest) SetLicenseItemIds(v string) *RenewAppLicenseRequest {
	s.LicenseItemIds = &v
	return s
}

func (s *RenewAppLicenseRequest) SetOrderIds(v string) *RenewAppLicenseRequest {
	s.OrderIds = &v
	return s
}

func (s *RenewAppLicenseRequest) SetPurchaseMethod(v string) *RenewAppLicenseRequest {
	s.PurchaseMethod = &v
	return s
}

func (s *RenewAppLicenseRequest) Validate() error {
	return dara.Validate(s)
}

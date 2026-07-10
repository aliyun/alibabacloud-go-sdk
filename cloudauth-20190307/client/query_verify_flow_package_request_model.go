// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVerifyFlowPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *QueryVerifyFlowPackageRequest
	GetProductType() *string
}

type QueryVerifyFlowPackageRequest struct {
	// The product type. Valid values:
	//
	// - **FINANCE_VERIFY**: Financial-Grade ID Verification
	//
	// - **SMART_VERIFY**: Enhanced ID Verification (discontinued)
	//
	// - **FACE_VERIFY**: ID Verification (discontinued).
	//
	// This parameter is required.
	//
	// example:
	//
	// FINANCE_VERIFY
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s QueryVerifyFlowPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyFlowPackageRequest) GoString() string {
	return s.String()
}

func (s *QueryVerifyFlowPackageRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryVerifyFlowPackageRequest) SetProductType(v string) *QueryVerifyFlowPackageRequest {
	s.ProductType = &v
	return s
}

func (s *QueryVerifyFlowPackageRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonBillingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMonth(v string) *QueryCommonBillingRequest
	GetMonth() *string
	SetOwnerId(v int64) *QueryCommonBillingRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QueryCommonBillingRequest
	GetProdCode() *string
	SetProductType(v int32) *QueryCommonBillingRequest
	GetProductType() *int32
	SetResourceOwnerAccount(v string) *QueryCommonBillingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryCommonBillingRequest
	GetResourceOwnerId() *int64
}

type QueryCommonBillingRequest struct {
	// This parameter is required.
	Month    *string `json:"Month,omitempty" xml:"Month,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// This parameter is required.
	ProductType          *int32  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCommonBillingRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonBillingRequest) GoString() string {
	return s.String()
}

func (s *QueryCommonBillingRequest) GetMonth() *string {
	return s.Month
}

func (s *QueryCommonBillingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryCommonBillingRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryCommonBillingRequest) GetProductType() *int32 {
	return s.ProductType
}

func (s *QueryCommonBillingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryCommonBillingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryCommonBillingRequest) SetMonth(v string) *QueryCommonBillingRequest {
	s.Month = &v
	return s
}

func (s *QueryCommonBillingRequest) SetOwnerId(v int64) *QueryCommonBillingRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCommonBillingRequest) SetProdCode(v string) *QueryCommonBillingRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryCommonBillingRequest) SetProductType(v int32) *QueryCommonBillingRequest {
	s.ProductType = &v
	return s
}

func (s *QueryCommonBillingRequest) SetResourceOwnerAccount(v string) *QueryCommonBillingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCommonBillingRequest) SetResourceOwnerId(v int64) *QueryCommonBillingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCommonBillingRequest) Validate() error {
	return dara.Validate(s)
}

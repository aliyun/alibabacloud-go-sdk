// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCommonProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *OpenCommonProductRequest
	GetOwnerId() *int64
	SetProdCode(v string) *OpenCommonProductRequest
	GetProdCode() *string
	SetProductType(v int32) *OpenCommonProductRequest
	GetProductType() *int32
	SetResourceOwnerAccount(v string) *OpenCommonProductRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenCommonProductRequest
	GetResourceOwnerId() *int64
}

type OpenCommonProductRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// This parameter is required.
	ProductType          *int32  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OpenCommonProductRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenCommonProductRequest) GoString() string {
	return s.String()
}

func (s *OpenCommonProductRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenCommonProductRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *OpenCommonProductRequest) GetProductType() *int32 {
	return s.ProductType
}

func (s *OpenCommonProductRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenCommonProductRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenCommonProductRequest) SetOwnerId(v int64) *OpenCommonProductRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenCommonProductRequest) SetProdCode(v string) *OpenCommonProductRequest {
	s.ProdCode = &v
	return s
}

func (s *OpenCommonProductRequest) SetProductType(v int32) *OpenCommonProductRequest {
	s.ProductType = &v
	return s
}

func (s *OpenCommonProductRequest) SetResourceOwnerAccount(v string) *OpenCommonProductRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenCommonProductRequest) SetResourceOwnerId(v int64) *OpenCommonProductRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenCommonProductRequest) Validate() error {
	return dara.Validate(s)
}

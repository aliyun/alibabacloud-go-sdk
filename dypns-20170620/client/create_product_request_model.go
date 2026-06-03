// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CreateProductRequest
	GetOwnerId() *int64
	SetProdCode(v string) *CreateProductRequest
	GetProdCode() *string
	SetProdId(v string) *CreateProductRequest
	GetProdId() *string
	SetResourceOwnerAccount(v string) *CreateProductRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateProductRequest
	GetResourceOwnerId() *int64
}

type CreateProductRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// This parameter is required.
	ProdId               *string `json:"ProdId,omitempty" xml:"ProdId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateProductRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProductRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateProductRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *CreateProductRequest) GetProdId() *string {
	return s.ProdId
}

func (s *CreateProductRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateProductRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateProductRequest) SetOwnerId(v int64) *CreateProductRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateProductRequest) SetProdCode(v string) *CreateProductRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateProductRequest) SetProdId(v string) *CreateProductRequest {
	s.ProdId = &v
	return s
}

func (s *CreateProductRequest) SetResourceOwnerAccount(v string) *CreateProductRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateProductRequest) SetResourceOwnerId(v int64) *CreateProductRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateProductRequest) Validate() error {
	return dara.Validate(s)
}

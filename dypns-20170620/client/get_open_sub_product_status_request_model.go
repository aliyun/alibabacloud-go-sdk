// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenSubProductStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetOpenSubProductStatusRequest
	GetOwnerId() *int64
	SetProductCode(v string) *GetOpenSubProductStatusRequest
	GetProductCode() *string
	SetResourceOwnerAccount(v string) *GetOpenSubProductStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetOpenSubProductStatusRequest
	GetResourceOwnerId() *int64
}

type GetOpenSubProductStatusRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	ProductCode          *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetOpenSubProductStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpenSubProductStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOpenSubProductStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetOpenSubProductStatusRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetOpenSubProductStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetOpenSubProductStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetOpenSubProductStatusRequest) SetOwnerId(v int64) *GetOpenSubProductStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetOpenSubProductStatusRequest) SetProductCode(v string) *GetOpenSubProductStatusRequest {
	s.ProductCode = &v
	return s
}

func (s *GetOpenSubProductStatusRequest) SetResourceOwnerAccount(v string) *GetOpenSubProductStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetOpenSubProductStatusRequest) SetResourceOwnerId(v int64) *GetOpenSubProductStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetOpenSubProductStatusRequest) Validate() error {
	return dara.Validate(s)
}

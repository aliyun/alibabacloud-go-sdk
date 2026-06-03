// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttrKey(v string) *QueryTagStatusRequest
	GetAttrKey() *string
	SetBizType(v int32) *QueryTagStatusRequest
	GetBizType() *int32
	SetOwnerId(v int64) *QueryTagStatusRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QueryTagStatusRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *QueryTagStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryTagStatusRequest
	GetResourceOwnerId() *int64
}

type QueryTagStatusRequest struct {
	// This parameter is required.
	AttrKey *string `json:"AttrKey,omitempty" xml:"AttrKey,omitempty"`
	// This parameter is required.
	BizType              *int32  `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryTagStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTagStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryTagStatusRequest) GetAttrKey() *string {
	return s.AttrKey
}

func (s *QueryTagStatusRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *QueryTagStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryTagStatusRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryTagStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryTagStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryTagStatusRequest) SetAttrKey(v string) *QueryTagStatusRequest {
	s.AttrKey = &v
	return s
}

func (s *QueryTagStatusRequest) SetBizType(v int32) *QueryTagStatusRequest {
	s.BizType = &v
	return s
}

func (s *QueryTagStatusRequest) SetOwnerId(v int64) *QueryTagStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTagStatusRequest) SetProdCode(v string) *QueryTagStatusRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryTagStatusRequest) SetResourceOwnerAccount(v string) *QueryTagStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTagStatusRequest) SetResourceOwnerId(v int64) *QueryTagStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTagStatusRequest) Validate() error {
	return dara.Validate(s)
}

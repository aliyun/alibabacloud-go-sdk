// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeCdnTypesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCdnTypesRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnTypesRequest
	GetSecurityToken() *string
}

type DescribeCdnTypesRequest struct {
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnTypesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCdnTypesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnTypesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnTypesRequest) SetOwnerAccount(v string) *DescribeCdnTypesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCdnTypesRequest) SetOwnerId(v int64) *DescribeCdnTypesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnTypesRequest) SetSecurityToken(v string) *DescribeCdnTypesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnTypesRequest) Validate() error {
	return dara.Validate(s)
}

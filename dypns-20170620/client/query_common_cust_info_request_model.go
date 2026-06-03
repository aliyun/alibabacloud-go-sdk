// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonCustInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryCommonCustInfoRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QueryCommonCustInfoRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *QueryCommonCustInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryCommonCustInfoRequest
	GetResourceOwnerId() *int64
}

type QueryCommonCustInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCommonCustInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonCustInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCommonCustInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryCommonCustInfoRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryCommonCustInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryCommonCustInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryCommonCustInfoRequest) SetOwnerId(v int64) *QueryCommonCustInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCommonCustInfoRequest) SetProdCode(v string) *QueryCommonCustInfoRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryCommonCustInfoRequest) SetResourceOwnerAccount(v string) *QueryCommonCustInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCommonCustInfoRequest) SetResourceOwnerId(v int64) *QueryCommonCustInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCommonCustInfoRequest) Validate() error {
	return dara.Validate(s)
}

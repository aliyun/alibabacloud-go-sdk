// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryPnsConfigRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QueryPnsConfigRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *QueryPnsConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPnsConfigRequest
	GetResourceOwnerId() *int64
}

type QueryPnsConfigRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPnsConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryPnsConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPnsConfigRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryPnsConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPnsConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPnsConfigRequest) SetOwnerId(v int64) *QueryPnsConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPnsConfigRequest) SetProdCode(v string) *QueryPnsConfigRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPnsConfigRequest) SetResourceOwnerAccount(v string) *QueryPnsConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPnsConfigRequest) SetResourceOwnerId(v int64) *QueryPnsConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPnsConfigRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmsSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerId(v int64) *ListSmsSignRequest
	GetCustomerId() *int64
	SetOwnerId(v int64) *ListSmsSignRequest
	GetOwnerId() *int64
	SetProdCode(v string) *ListSmsSignRequest
	GetProdCode() *string
	SetQuerySmsSign(v string) *ListSmsSignRequest
	GetQuerySmsSign() *string
	SetResourceOwnerAccount(v string) *ListSmsSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListSmsSignRequest
	GetResourceOwnerId() *int64
}

type ListSmsSignRequest struct {
	CustomerId           *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	QuerySmsSign         *string `json:"QuerySmsSign,omitempty" xml:"QuerySmsSign,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListSmsSignRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSmsSignRequest) GoString() string {
	return s.String()
}

func (s *ListSmsSignRequest) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *ListSmsSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListSmsSignRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *ListSmsSignRequest) GetQuerySmsSign() *string {
	return s.QuerySmsSign
}

func (s *ListSmsSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListSmsSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListSmsSignRequest) SetCustomerId(v int64) *ListSmsSignRequest {
	s.CustomerId = &v
	return s
}

func (s *ListSmsSignRequest) SetOwnerId(v int64) *ListSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSmsSignRequest) SetProdCode(v string) *ListSmsSignRequest {
	s.ProdCode = &v
	return s
}

func (s *ListSmsSignRequest) SetQuerySmsSign(v string) *ListSmsSignRequest {
	s.QuerySmsSign = &v
	return s
}

func (s *ListSmsSignRequest) SetResourceOwnerAccount(v string) *ListSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListSmsSignRequest) SetResourceOwnerId(v int64) *ListSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListSmsSignRequest) Validate() error {
	return dara.Validate(s)
}

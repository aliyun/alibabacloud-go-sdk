// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySdkVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QuerySdkVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySdkVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySdkVersionRequest
	GetResourceOwnerId() *int64
	SetCustomerId(v int64) *QuerySdkVersionRequest
	GetCustomerId() *int64
	SetProdCode(v string) *QuerySdkVersionRequest
	GetProdCode() *string
}

type QuerySdkVersionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CustomerId           *int64  `json:"customerId,omitempty" xml:"customerId,omitempty"`
	ProdCode             *string `json:"prodCode,omitempty" xml:"prodCode,omitempty"`
}

func (s QuerySdkVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySdkVersionRequest) GoString() string {
	return s.String()
}

func (s *QuerySdkVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySdkVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySdkVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySdkVersionRequest) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *QuerySdkVersionRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QuerySdkVersionRequest) SetOwnerId(v int64) *QuerySdkVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySdkVersionRequest) SetResourceOwnerAccount(v string) *QuerySdkVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySdkVersionRequest) SetResourceOwnerId(v int64) *QuerySdkVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySdkVersionRequest) SetCustomerId(v int64) *QuerySdkVersionRequest {
	s.CustomerId = &v
	return s
}

func (s *QuerySdkVersionRequest) SetProdCode(v string) *QuerySdkVersionRequest {
	s.ProdCode = &v
	return s
}

func (s *QuerySdkVersionRequest) Validate() error {
	return dara.Validate(s)
}

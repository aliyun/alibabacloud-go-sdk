// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPackageTypeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryPackageTypeInfoRequest
	GetOwnerId() *int64
	SetProductName(v string) *QueryPackageTypeInfoRequest
	GetProductName() *string
	SetResourceOwnerAccount(v string) *QueryPackageTypeInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPackageTypeInfoRequest
	GetResourceOwnerId() *int64
}

type QueryPackageTypeInfoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 产品名称，如dysms
	//
	// example:
	//
	// dytns
	ProductName          *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPackageTypeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPackageTypeInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryPackageTypeInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPackageTypeInfoRequest) GetProductName() *string {
	return s.ProductName
}

func (s *QueryPackageTypeInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPackageTypeInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPackageTypeInfoRequest) SetOwnerId(v int64) *QueryPackageTypeInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPackageTypeInfoRequest) SetProductName(v string) *QueryPackageTypeInfoRequest {
	s.ProductName = &v
	return s
}

func (s *QueryPackageTypeInfoRequest) SetResourceOwnerAccount(v string) *QueryPackageTypeInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPackageTypeInfoRequest) SetResourceOwnerId(v int64) *QueryPackageTypeInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPackageTypeInfoRequest) Validate() error {
	return dara.Validate(s)
}

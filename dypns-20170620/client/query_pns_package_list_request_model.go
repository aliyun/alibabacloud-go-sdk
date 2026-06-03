// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsPackageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillCycle(v string) *QueryPnsPackageListRequest
	GetBillCycle() *string
	SetOwnerId(v int64) *QueryPnsPackageListRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QueryPnsPackageListRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *QueryPnsPackageListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPnsPackageListRequest
	GetResourceOwnerId() *int64
}

type QueryPnsPackageListRequest struct {
	BillCycle            *string `json:"BillCycle,omitempty" xml:"BillCycle,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPnsPackageListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsPackageListRequest) GoString() string {
	return s.String()
}

func (s *QueryPnsPackageListRequest) GetBillCycle() *string {
	return s.BillCycle
}

func (s *QueryPnsPackageListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPnsPackageListRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryPnsPackageListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPnsPackageListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPnsPackageListRequest) SetBillCycle(v string) *QueryPnsPackageListRequest {
	s.BillCycle = &v
	return s
}

func (s *QueryPnsPackageListRequest) SetOwnerId(v int64) *QueryPnsPackageListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPnsPackageListRequest) SetProdCode(v string) *QueryPnsPackageListRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPnsPackageListRequest) SetResourceOwnerAccount(v string) *QueryPnsPackageListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPnsPackageListRequest) SetResourceOwnerId(v int64) *QueryPnsPackageListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPnsPackageListRequest) Validate() error {
	return dara.Validate(s)
}

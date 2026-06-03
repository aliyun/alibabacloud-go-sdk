// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsPackageDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryPnsPackageDetailRequest
	GetOwnerId() *int64
	SetPageNo(v string) *QueryPnsPackageDetailRequest
	GetPageNo() *string
	SetPageSize(v string) *QueryPnsPackageDetailRequest
	GetPageSize() *string
	SetProdCode(v string) *QueryPnsPackageDetailRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *QueryPnsPackageDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPnsPackageDetailRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *QueryPnsPackageDetailRequest
	GetStatus() *string
}

type QueryPnsPackageDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryPnsPackageDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsPackageDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryPnsPackageDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPnsPackageDetailRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *QueryPnsPackageDetailRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *QueryPnsPackageDetailRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryPnsPackageDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPnsPackageDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPnsPackageDetailRequest) GetStatus() *string {
	return s.Status
}

func (s *QueryPnsPackageDetailRequest) SetOwnerId(v int64) *QueryPnsPackageDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPnsPackageDetailRequest) SetPageNo(v string) *QueryPnsPackageDetailRequest {
	s.PageNo = &v
	return s
}

func (s *QueryPnsPackageDetailRequest) SetPageSize(v string) *QueryPnsPackageDetailRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPnsPackageDetailRequest) SetProdCode(v string) *QueryPnsPackageDetailRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPnsPackageDetailRequest) SetResourceOwnerAccount(v string) *QueryPnsPackageDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPnsPackageDetailRequest) SetResourceOwnerId(v int64) *QueryPnsPackageDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPnsPackageDetailRequest) SetStatus(v string) *QueryPnsPackageDetailRequest {
	s.Status = &v
	return s
}

func (s *QueryPnsPackageDetailRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsPackageStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryPnsPackageStatisticsRequest
	GetOwnerId() *int64
	SetProdCode(v string) *QueryPnsPackageStatisticsRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *QueryPnsPackageStatisticsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPnsPackageStatisticsRequest
	GetResourceOwnerId() *int64
}

type QueryPnsPackageStatisticsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPnsPackageStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsPackageStatisticsRequest) GoString() string {
	return s.String()
}

func (s *QueryPnsPackageStatisticsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPnsPackageStatisticsRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryPnsPackageStatisticsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPnsPackageStatisticsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPnsPackageStatisticsRequest) SetOwnerId(v int64) *QueryPnsPackageStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPnsPackageStatisticsRequest) SetProdCode(v string) *QueryPnsPackageStatisticsRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryPnsPackageStatisticsRequest) SetResourceOwnerAccount(v string) *QueryPnsPackageStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPnsPackageStatisticsRequest) SetResourceOwnerId(v int64) *QueryPnsPackageStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPnsPackageStatisticsRequest) Validate() error {
	return dara.Validate(s)
}

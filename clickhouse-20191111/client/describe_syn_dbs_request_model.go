// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynDbsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbClusterId(v string) *DescribeSynDbsRequest
	GetDbClusterId() *string
	SetOwnerAccount(v string) *DescribeSynDbsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSynDbsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSynDbsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSynDbsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeSynDbsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSynDbsRequest
	GetResourceOwnerId() *int64
}

type DescribeSynDbsRequest struct {
	// The ID of the ApsaraDB for ClickHouse cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1ab22b80814****
	DbClusterId  *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSynDbsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynDbsRequest) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *DescribeSynDbsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSynDbsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSynDbsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSynDbsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSynDbsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSynDbsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSynDbsRequest) SetDbClusterId(v string) *DescribeSynDbsRequest {
	s.DbClusterId = &v
	return s
}

func (s *DescribeSynDbsRequest) SetOwnerAccount(v string) *DescribeSynDbsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSynDbsRequest) SetOwnerId(v int64) *DescribeSynDbsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynDbsRequest) SetPageNumber(v int32) *DescribeSynDbsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSynDbsRequest) SetPageSize(v int32) *DescribeSynDbsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSynDbsRequest) SetResourceOwnerAccount(v string) *DescribeSynDbsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSynDbsRequest) SetResourceOwnerId(v int64) *DescribeSynDbsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSynDbsRequest) Validate() error {
	return dara.Validate(s)
}

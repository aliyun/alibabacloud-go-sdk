// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDistributeCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalInstanceId(v string) *DescribeGlobalDistributeCacheRequest
	GetGlobalInstanceId() *string
	SetOwnerAccount(v string) *DescribeGlobalDistributeCacheRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGlobalDistributeCacheRequest
	GetOwnerId() *int64
	SetPageNumber(v string) *DescribeGlobalDistributeCacheRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeGlobalDistributeCacheRequest
	GetPageSize() *string
	SetResourceOwnerAccount(v string) *DescribeGlobalDistributeCacheRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGlobalDistributeCacheRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeGlobalDistributeCacheRequest
	GetSecurityToken() *string
	SetSubInstanceId(v string) *DescribeGlobalDistributeCacheRequest
	GetSubInstanceId() *string
}

type DescribeGlobalDistributeCacheRequest struct {
	// The ID of the distributed instance.
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return each page.
	//
	// example:
	//
	// 20
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the child instance that is attached to the distributed instance.
	//
	// example:
	//
	// gr-bp1zcjlobkyrq****
	SubInstanceId *string `json:"SubInstanceId,omitempty" xml:"SubInstanceId,omitempty"`
}

func (s DescribeGlobalDistributeCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDistributeCacheRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDistributeCacheRequest) GetGlobalInstanceId() *string {
	return s.GlobalInstanceId
}

func (s *DescribeGlobalDistributeCacheRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGlobalDistributeCacheRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGlobalDistributeCacheRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeGlobalDistributeCacheRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeGlobalDistributeCacheRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGlobalDistributeCacheRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGlobalDistributeCacheRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeGlobalDistributeCacheRequest) GetSubInstanceId() *string {
	return s.SubInstanceId
}

func (s *DescribeGlobalDistributeCacheRequest) SetGlobalInstanceId(v string) *DescribeGlobalDistributeCacheRequest {
	s.GlobalInstanceId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetOwnerAccount(v string) *DescribeGlobalDistributeCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetOwnerId(v int64) *DescribeGlobalDistributeCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetPageNumber(v string) *DescribeGlobalDistributeCacheRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetPageSize(v string) *DescribeGlobalDistributeCacheRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetResourceOwnerAccount(v string) *DescribeGlobalDistributeCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetResourceOwnerId(v int64) *DescribeGlobalDistributeCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetSecurityToken(v string) *DescribeGlobalDistributeCacheRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) SetSubInstanceId(v string) *DescribeGlobalDistributeCacheRequest {
	s.SubInstanceId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheRequest) Validate() error {
	return dara.Validate(s)
}

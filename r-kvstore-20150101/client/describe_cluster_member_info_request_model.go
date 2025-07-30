// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterMemberInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeClusterMemberInfoRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeClusterMemberInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeClusterMemberInfoRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeClusterMemberInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClusterMemberInfoRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeClusterMemberInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClusterMemberInfoRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeClusterMemberInfoRequest
	GetSecurityToken() *string
}

type DescribeClusterMemberInfoRequest struct {
	// The ID of the Tair (Redis OSS-compatible) instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeClusterMemberInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterMemberInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterMemberInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeClusterMemberInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeClusterMemberInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClusterMemberInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterMemberInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterMemberInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClusterMemberInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClusterMemberInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeClusterMemberInfoRequest) SetInstanceId(v string) *DescribeClusterMemberInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetOwnerAccount(v string) *DescribeClusterMemberInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetOwnerId(v int64) *DescribeClusterMemberInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetPageNumber(v int32) *DescribeClusterMemberInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetPageSize(v int32) *DescribeClusterMemberInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetResourceOwnerAccount(v string) *DescribeClusterMemberInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetResourceOwnerId(v int64) *DescribeClusterMemberInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) SetSecurityToken(v string) *DescribeClusterMemberInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeClusterMemberInfoRequest) Validate() error {
	return dara.Validate(s)
}

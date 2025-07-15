// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouterInterfaceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeRouterInterfaceAttributeRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DescribeRouterInterfaceAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeRouterInterfaceAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeRouterInterfaceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRouterInterfaceAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeRouterInterfaceAttributeRequest struct {
	// The ID of the router interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// ri-m5egfc10sednwk2yt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the router interface belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRouterInterfaceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfaceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfaceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRouterInterfaceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRouterInterfaceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouterInterfaceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRouterInterfaceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRouterInterfaceAttributeRequest) SetInstanceId(v string) *DescribeRouterInterfaceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeRequest) SetOwnerId(v int64) *DescribeRouterInterfaceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeRequest) SetRegionId(v string) *DescribeRouterInterfaceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeRequest) SetResourceOwnerAccount(v string) *DescribeRouterInterfaceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeRequest) SetResourceOwnerId(v int64) *DescribeRouterInterfaceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeRequest) Validate() error {
	return dara.Validate(s)
}

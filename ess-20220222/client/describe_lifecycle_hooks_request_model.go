// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecycleHooksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLifecycleHookIds(v []*string) *DescribeLifecycleHooksRequest
	GetLifecycleHookIds() []*string
	SetLifecycleHookName(v string) *DescribeLifecycleHooksRequest
	GetLifecycleHookName() *string
	SetOwnerAccount(v string) *DescribeLifecycleHooksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLifecycleHooksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLifecycleHooksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLifecycleHooksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLifecycleHooksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLifecycleHooksRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DescribeLifecycleHooksRequest
	GetScalingGroupId() *string
}

type DescribeLifecycleHooksRequest struct {
	// The IDs of the lifecycle hooks that you want to query.
	LifecycleHookIds []*string `json:"LifecycleHookIds,omitempty" xml:"LifecycleHookIds,omitempty" type:"Repeated"`
	// The name of the lifecycle hook.
	//
	// example:
	//
	// lifecyclehook****
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 50.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeLifecycleHooksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecycleHooksRequest) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksRequest) GetLifecycleHookIds() []*string {
	return s.LifecycleHookIds
}

func (s *DescribeLifecycleHooksRequest) GetLifecycleHookName() *string {
	return s.LifecycleHookName
}

func (s *DescribeLifecycleHooksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLifecycleHooksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLifecycleHooksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLifecycleHooksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLifecycleHooksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLifecycleHooksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLifecycleHooksRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeLifecycleHooksRequest) SetLifecycleHookIds(v []*string) *DescribeLifecycleHooksRequest {
	s.LifecycleHookIds = v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetLifecycleHookName(v string) *DescribeLifecycleHooksRequest {
	s.LifecycleHookName = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetOwnerAccount(v string) *DescribeLifecycleHooksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetOwnerId(v int64) *DescribeLifecycleHooksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetPageNumber(v int32) *DescribeLifecycleHooksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetPageSize(v int32) *DescribeLifecycleHooksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetRegionId(v string) *DescribeLifecycleHooksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetResourceOwnerAccount(v string) *DescribeLifecycleHooksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) SetScalingGroupId(v string) *DescribeLifecycleHooksRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeLifecycleHooksRequest) Validate() error {
	return dara.Validate(s)
}

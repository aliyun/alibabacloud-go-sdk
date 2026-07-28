// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToVbrShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrantType(v string) *GrantInstanceToVbrShrinkRequest
	GetGrantType() *string
	SetInstanceId(v string) *GrantInstanceToVbrShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *GrantInstanceToVbrShrinkRequest
	GetRegionId() *string
	SetVbrInstanceIdsShrink(v string) *GrantInstanceToVbrShrinkRequest
	GetVbrInstanceIdsShrink() *string
	SetVbrOwnerUid(v int64) *GrantInstanceToVbrShrinkRequest
	GetVbrOwnerUid() *int64
	SetVbrRegionNo(v string) *GrantInstanceToVbrShrinkRequest
	GetVbrRegionNo() *string
}

type GrantInstanceToVbrShrinkRequest struct {
	// The scope of VBR instances to be granted authorization. Valid values:
	//
	// - **All**: Grants authorization of the VPC-connected instance to all VBR instances in the specified region under the specified Alibaba Cloud account. In this case, the **VbrInstanceIds*	- parameter can be left empty.
	//
	// - **Specify**: Grants authorization of the VPC-connected instance to the specified VBR instances. In this case, the **VbrInstanceIds*	- parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// All
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// The ID of the VPC-connected instance for which to grant authorization.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1lqhq93q8evjpky****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the VPC-connected instance for which to grant authorization.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of VBR instances to be granted authorization.
	//
	// if can be null:
	// true
	VbrInstanceIdsShrink *string `json:"VbrInstanceIds,omitempty" xml:"VbrInstanceIds,omitempty"`
	// The Alibaba Cloud account ID that owns the VBR instance to be granted authorization.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1210123456123456
	VbrOwnerUid *int64 `json:"VbrOwnerUid,omitempty" xml:"VbrOwnerUid,omitempty"`
	// The region ID of the VBR instance to be granted authorization.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VbrRegionNo *string `json:"VbrRegionNo,omitempty" xml:"VbrRegionNo,omitempty"`
}

func (s GrantInstanceToVbrShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToVbrShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantInstanceToVbrShrinkRequest) GetGrantType() *string {
	return s.GrantType
}

func (s *GrantInstanceToVbrShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GrantInstanceToVbrShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantInstanceToVbrShrinkRequest) GetVbrInstanceIdsShrink() *string {
	return s.VbrInstanceIdsShrink
}

func (s *GrantInstanceToVbrShrinkRequest) GetVbrOwnerUid() *int64 {
	return s.VbrOwnerUid
}

func (s *GrantInstanceToVbrShrinkRequest) GetVbrRegionNo() *string {
	return s.VbrRegionNo
}

func (s *GrantInstanceToVbrShrinkRequest) SetGrantType(v string) *GrantInstanceToVbrShrinkRequest {
	s.GrantType = &v
	return s
}

func (s *GrantInstanceToVbrShrinkRequest) SetInstanceId(v string) *GrantInstanceToVbrShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantInstanceToVbrShrinkRequest) SetRegionId(v string) *GrantInstanceToVbrShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GrantInstanceToVbrShrinkRequest) SetVbrInstanceIdsShrink(v string) *GrantInstanceToVbrShrinkRequest {
	s.VbrInstanceIdsShrink = &v
	return s
}

func (s *GrantInstanceToVbrShrinkRequest) SetVbrOwnerUid(v int64) *GrantInstanceToVbrShrinkRequest {
	s.VbrOwnerUid = &v
	return s
}

func (s *GrantInstanceToVbrShrinkRequest) SetVbrRegionNo(v string) *GrantInstanceToVbrShrinkRequest {
	s.VbrRegionNo = &v
	return s
}

func (s *GrantInstanceToVbrShrinkRequest) Validate() error {
	return dara.Validate(s)
}

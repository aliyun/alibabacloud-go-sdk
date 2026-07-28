// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromVbrShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrantType(v string) *RevokeInstanceFromVbrShrinkRequest
	GetGrantType() *string
	SetInstanceId(v string) *RevokeInstanceFromVbrShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *RevokeInstanceFromVbrShrinkRequest
	GetRegionId() *string
	SetVbrInstanceIdsShrink(v string) *RevokeInstanceFromVbrShrinkRequest
	GetVbrInstanceIdsShrink() *string
	SetVbrOwnerUid(v string) *RevokeInstanceFromVbrShrinkRequest
	GetVbrOwnerUid() *string
	SetVbrRegionNo(v string) *RevokeInstanceFromVbrShrinkRequest
	GetVbrRegionNo() *string
}

type RevokeInstanceFromVbrShrinkRequest struct {
	// The scope of VBR instances for which the VPC-connected instance authorization is being revoked. Valid values:
	//
	// - **ALL**: Revokes the VPC-connected instance authorization for all VBR instances in the specified region. In this case, the **VbrInstanceIds*	- parameter can be left empty.
	//
	// - **Specify**: Revokes the VPC-connected instance authorization for the specified VBR instances. In this case, the **VbrInstanceIds*	- parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// The instance ID of the VPC-connected instance for which you want to revoke the authorization.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1brjuegjc88v3u9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the VPC-connected instance for which you want to revoke the authorization.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of VBR instance IDs for which the VPC-connected instance authorization is being revoked.
	VbrInstanceIdsShrink *string `json:"VbrInstanceIds,omitempty" xml:"VbrInstanceIds,omitempty"`
	// The Alibaba Cloud account ID of the VBR instance for which the authorization is being revoked.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1210123456123456
	VbrOwnerUid *string `json:"VbrOwnerUid,omitempty" xml:"VbrOwnerUid,omitempty"`
	// The region ID of the VBR instance for which the VPC-connected instance authorization is being revoked.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VbrRegionNo *string `json:"VbrRegionNo,omitempty" xml:"VbrRegionNo,omitempty"`
}

func (s RevokeInstanceFromVbrShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromVbrShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromVbrShrinkRequest) GetGrantType() *string {
	return s.GrantType
}

func (s *RevokeInstanceFromVbrShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeInstanceFromVbrShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeInstanceFromVbrShrinkRequest) GetVbrInstanceIdsShrink() *string {
	return s.VbrInstanceIdsShrink
}

func (s *RevokeInstanceFromVbrShrinkRequest) GetVbrOwnerUid() *string {
	return s.VbrOwnerUid
}

func (s *RevokeInstanceFromVbrShrinkRequest) GetVbrRegionNo() *string {
	return s.VbrRegionNo
}

func (s *RevokeInstanceFromVbrShrinkRequest) SetGrantType(v string) *RevokeInstanceFromVbrShrinkRequest {
	s.GrantType = &v
	return s
}

func (s *RevokeInstanceFromVbrShrinkRequest) SetInstanceId(v string) *RevokeInstanceFromVbrShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeInstanceFromVbrShrinkRequest) SetRegionId(v string) *RevokeInstanceFromVbrShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeInstanceFromVbrShrinkRequest) SetVbrInstanceIdsShrink(v string) *RevokeInstanceFromVbrShrinkRequest {
	s.VbrInstanceIdsShrink = &v
	return s
}

func (s *RevokeInstanceFromVbrShrinkRequest) SetVbrOwnerUid(v string) *RevokeInstanceFromVbrShrinkRequest {
	s.VbrOwnerUid = &v
	return s
}

func (s *RevokeInstanceFromVbrShrinkRequest) SetVbrRegionNo(v string) *RevokeInstanceFromVbrShrinkRequest {
	s.VbrRegionNo = &v
	return s
}

func (s *RevokeInstanceFromVbrShrinkRequest) Validate() error {
	return dara.Validate(s)
}

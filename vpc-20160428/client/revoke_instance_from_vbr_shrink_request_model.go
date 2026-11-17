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
	// The scope of VBR instances from which the authorization of the VPC-connected instance is revoked. Valid values:
	//
	// - **ALL**: revokes the authorization between the VPC-connected instance and all VBR instances in the specified region. In this case, the **VbrInstanceIds*	- parameter can be left empty.
	//
	// - **Specify**: revokes the authorization between the VPC-connected instance and the specified VBR instances. In this case, the **VbrInstanceIds*	- parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// The VPC-connected instance ID for which you want to revoke the authorization.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1brjuegjc88v3u9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the VPC-connected instance for which you want to revoke the authorization. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of VBR instance IDs from which the authorization of the VPC-connected instance is revoked.
	VbrInstanceIdsShrink *string `json:"VbrInstanceIds,omitempty" xml:"VbrInstanceIds,omitempty"`
	// The ID of the Alibaba Cloud account that owns the VBR instance from which the authorization is revoked.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1210123456123456
	VbrOwnerUid *string `json:"VbrOwnerUid,omitempty" xml:"VbrOwnerUid,omitempty"`
	// The region ID of the VBR instance from which the authorization of the VPC-connected instance is revoked.
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

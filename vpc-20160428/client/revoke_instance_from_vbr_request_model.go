// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromVbrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrantType(v string) *RevokeInstanceFromVbrRequest
	GetGrantType() *string
	SetInstanceId(v string) *RevokeInstanceFromVbrRequest
	GetInstanceId() *string
	SetRegionId(v string) *RevokeInstanceFromVbrRequest
	GetRegionId() *string
	SetVbrInstanceIds(v []*string) *RevokeInstanceFromVbrRequest
	GetVbrInstanceIds() []*string
	SetVbrOwnerUid(v string) *RevokeInstanceFromVbrRequest
	GetVbrOwnerUid() *string
	SetVbrRegionNo(v string) *RevokeInstanceFromVbrRequest
	GetVbrRegionNo() *string
}

type RevokeInstanceFromVbrRequest struct {
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
	VbrInstanceIds []*string `json:"VbrInstanceIds,omitempty" xml:"VbrInstanceIds,omitempty" type:"Repeated"`
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

func (s RevokeInstanceFromVbrRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromVbrRequest) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromVbrRequest) GetGrantType() *string {
	return s.GrantType
}

func (s *RevokeInstanceFromVbrRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeInstanceFromVbrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeInstanceFromVbrRequest) GetVbrInstanceIds() []*string {
	return s.VbrInstanceIds
}

func (s *RevokeInstanceFromVbrRequest) GetVbrOwnerUid() *string {
	return s.VbrOwnerUid
}

func (s *RevokeInstanceFromVbrRequest) GetVbrRegionNo() *string {
	return s.VbrRegionNo
}

func (s *RevokeInstanceFromVbrRequest) SetGrantType(v string) *RevokeInstanceFromVbrRequest {
	s.GrantType = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) SetInstanceId(v string) *RevokeInstanceFromVbrRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) SetRegionId(v string) *RevokeInstanceFromVbrRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) SetVbrInstanceIds(v []*string) *RevokeInstanceFromVbrRequest {
	s.VbrInstanceIds = v
	return s
}

func (s *RevokeInstanceFromVbrRequest) SetVbrOwnerUid(v string) *RevokeInstanceFromVbrRequest {
	s.VbrOwnerUid = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) SetVbrRegionNo(v string) *RevokeInstanceFromVbrRequest {
	s.VbrRegionNo = &v
	return s
}

func (s *RevokeInstanceFromVbrRequest) Validate() error {
	return dara.Validate(s)
}

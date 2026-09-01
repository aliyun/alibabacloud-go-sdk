// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToVbrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrantType(v string) *GrantInstanceToVbrRequest
	GetGrantType() *string
	SetInstanceId(v string) *GrantInstanceToVbrRequest
	GetInstanceId() *string
	SetRegionId(v string) *GrantInstanceToVbrRequest
	GetRegionId() *string
	SetVbrInstanceIds(v []*string) *GrantInstanceToVbrRequest
	GetVbrInstanceIds() []*string
	SetVbrOwnerUid(v int64) *GrantInstanceToVbrRequest
	GetVbrOwnerUid() *int64
	SetVbrRegionNo(v string) *GrantInstanceToVbrRequest
	GetVbrRegionNo() *string
}

type GrantInstanceToVbrRequest struct {
	// The scope of the VBR instances to be authorized. Valid values:
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
	// The ID of the VPC-connected instance for which authorization is to be granted.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1lqhq93q8evjpky****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the VPC-connected instance for which authorization is to be granted.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of VBR instances to be authorized.
	//
	// if can be null:
	// true
	VbrInstanceIds []*string `json:"VbrInstanceIds,omitempty" xml:"VbrInstanceIds,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account that owns the VBR instance to be authorized. This account must be different from the caller\\"s account. You cannot specify the caller\\"s own account ID. This operation is used for cross-account authorization.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1210123456123456
	VbrOwnerUid *int64 `json:"VbrOwnerUid,omitempty" xml:"VbrOwnerUid,omitempty"`
	// The region ID of the VBR instance to be authorized.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VbrRegionNo *string `json:"VbrRegionNo,omitempty" xml:"VbrRegionNo,omitempty"`
}

func (s GrantInstanceToVbrRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToVbrRequest) GoString() string {
	return s.String()
}

func (s *GrantInstanceToVbrRequest) GetGrantType() *string {
	return s.GrantType
}

func (s *GrantInstanceToVbrRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GrantInstanceToVbrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantInstanceToVbrRequest) GetVbrInstanceIds() []*string {
	return s.VbrInstanceIds
}

func (s *GrantInstanceToVbrRequest) GetVbrOwnerUid() *int64 {
	return s.VbrOwnerUid
}

func (s *GrantInstanceToVbrRequest) GetVbrRegionNo() *string {
	return s.VbrRegionNo
}

func (s *GrantInstanceToVbrRequest) SetGrantType(v string) *GrantInstanceToVbrRequest {
	s.GrantType = &v
	return s
}

func (s *GrantInstanceToVbrRequest) SetInstanceId(v string) *GrantInstanceToVbrRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantInstanceToVbrRequest) SetRegionId(v string) *GrantInstanceToVbrRequest {
	s.RegionId = &v
	return s
}

func (s *GrantInstanceToVbrRequest) SetVbrInstanceIds(v []*string) *GrantInstanceToVbrRequest {
	s.VbrInstanceIds = v
	return s
}

func (s *GrantInstanceToVbrRequest) SetVbrOwnerUid(v int64) *GrantInstanceToVbrRequest {
	s.VbrOwnerUid = &v
	return s
}

func (s *GrantInstanceToVbrRequest) SetVbrRegionNo(v string) *GrantInstanceToVbrRequest {
	s.VbrRegionNo = &v
	return s
}

func (s *GrantInstanceToVbrRequest) Validate() error {
	return dara.Validate(s)
}

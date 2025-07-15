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
	// The VBRs that need to acquire permissions on the VPC. Valid values:
	//
	// 	- **All**: Permissions on the VPC are granted to all VBRs that belong to the specified region and Alibaba Cloud account. In this case, you can leave **VbrInstanceIds*	- empty.
	//
	// 	- **Specify**: Permissions on the VPC are granted to the specified VBRs. **VbrInstanceIds*	- must be assigned a value.
	//
	// This parameter is required.
	//
	// example:
	//
	// All
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1lqhq93q8evjpky****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the VPC is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about the VBRs.
	//
	// if can be null:
	// true
	VbrInstanceIds []*string `json:"VbrInstanceIds,omitempty" xml:"VbrInstanceIds,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account to which the VBR belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1210123456123456
	VbrOwnerUid *int64 `json:"VbrOwnerUid,omitempty" xml:"VbrOwnerUid,omitempty"`
	// The ID of the region where the VBR is deployed.
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

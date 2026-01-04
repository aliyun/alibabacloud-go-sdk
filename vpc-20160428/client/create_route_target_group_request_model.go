// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteTargetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRouteTargetGroupRequest
	GetClientToken() *string
	SetConfigMode(v string) *CreateRouteTargetGroupRequest
	GetConfigMode() *string
	SetRegionId(v string) *CreateRouteTargetGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateRouteTargetGroupRequest
	GetResourceGroupId() *string
	SetRouteTargetGroupDescription(v string) *CreateRouteTargetGroupRequest
	GetRouteTargetGroupDescription() *string
	SetRouteTargetGroupName(v string) *CreateRouteTargetGroupRequest
	GetRouteTargetGroupName() *string
	SetRouteTargetMemberList(v []*CreateRouteTargetGroupRequestRouteTargetMemberList) *CreateRouteTargetGroupRequest
	GetRouteTargetMemberList() []*CreateRouteTargetGroupRequestRouteTargetMemberList
	SetTag(v []*CreateRouteTargetGroupRequestTag) *CreateRouteTargetGroupRequest
	GetTag() []*CreateRouteTargetGroupRequestTag
	SetVpcId(v string) *CreateRouteTargetGroupRequest
	GetVpcId() *string
}

type CreateRouteTargetGroupRequest struct {
	// Client token used to ensure the idempotence of the request. Generate a parameter value from your client to ensure that it is unique across different requests. ClientToken supports only ASCII characters. Note: If you do not specify this, the system automatically uses the RequestId of the API request as the ClientToken identifier. Each API request has a different RequestId.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe6****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration mode of the route target group. Supported modes:
	//
	// - **Active-Standby**: Active-Standby mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// Active-Standby
	ConfigMode *string `json:"ConfigMode,omitempty" xml:"ConfigMode,omitempty"`
	// The region ID to which the route target group belongs. You can obtain the region ID by calling the DescribeRegions interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [What is a Resource Group](https://help.aliyun.com/document_detail/2381067.html).
	//
	// example:
	//
	// rg-acfmxazffggds****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The description of the route target group.
	//
	// The description length must be between 1 and 256 characters, and cannot start with http:// or https://.
	//
	// example:
	//
	// myRouteTargetGroupDescription
	RouteTargetGroupDescription *string `json:"RouteTargetGroupDescription,omitempty" xml:"RouteTargetGroupDescription,omitempty"`
	// The name of the route target group.
	//
	// The name length must be between 1 and 128 characters, and cannot start with http:// or https://.
	//
	// example:
	//
	// myRouteTargetGroupName
	RouteTargetGroupName *string `json:"RouteTargetGroupName,omitempty" xml:"RouteTargetGroupName,omitempty"`
	// The member list of the route target group.
	//
	// In Active-Standby mode, the following restrictions apply to the members of the route target group:
	//
	// 1. The number of route target group members must be 2.
	//
	// 2. The route target group members must belong to different availability zones.
	//
	// This parameter is required.
	RouteTargetMemberList []*CreateRouteTargetGroupRequestRouteTargetMemberList `json:"RouteTargetMemberList,omitempty" xml:"RouteTargetMemberList,omitempty" type:"Repeated"`
	// The tags of the resource.
	Tag []*CreateRouteTargetGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC to which the route target group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateRouteTargetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteTargetGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteTargetGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRouteTargetGroupRequest) GetConfigMode() *string {
	return s.ConfigMode
}

func (s *CreateRouteTargetGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRouteTargetGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRouteTargetGroupRequest) GetRouteTargetGroupDescription() *string {
	return s.RouteTargetGroupDescription
}

func (s *CreateRouteTargetGroupRequest) GetRouteTargetGroupName() *string {
	return s.RouteTargetGroupName
}

func (s *CreateRouteTargetGroupRequest) GetRouteTargetMemberList() []*CreateRouteTargetGroupRequestRouteTargetMemberList {
	return s.RouteTargetMemberList
}

func (s *CreateRouteTargetGroupRequest) GetTag() []*CreateRouteTargetGroupRequestTag {
	return s.Tag
}

func (s *CreateRouteTargetGroupRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateRouteTargetGroupRequest) SetClientToken(v string) *CreateRouteTargetGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRouteTargetGroupRequest) SetConfigMode(v string) *CreateRouteTargetGroupRequest {
	s.ConfigMode = &v
	return s
}

func (s *CreateRouteTargetGroupRequest) SetRegionId(v string) *CreateRouteTargetGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRouteTargetGroupRequest) SetResourceGroupId(v string) *CreateRouteTargetGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRouteTargetGroupRequest) SetRouteTargetGroupDescription(v string) *CreateRouteTargetGroupRequest {
	s.RouteTargetGroupDescription = &v
	return s
}

func (s *CreateRouteTargetGroupRequest) SetRouteTargetGroupName(v string) *CreateRouteTargetGroupRequest {
	s.RouteTargetGroupName = &v
	return s
}

func (s *CreateRouteTargetGroupRequest) SetRouteTargetMemberList(v []*CreateRouteTargetGroupRequestRouteTargetMemberList) *CreateRouteTargetGroupRequest {
	s.RouteTargetMemberList = v
	return s
}

func (s *CreateRouteTargetGroupRequest) SetTag(v []*CreateRouteTargetGroupRequestTag) *CreateRouteTargetGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateRouteTargetGroupRequest) SetVpcId(v string) *CreateRouteTargetGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateRouteTargetGroupRequest) Validate() error {
	if s.RouteTargetMemberList != nil {
		for _, item := range s.RouteTargetMemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRouteTargetGroupRequestRouteTargetMemberList struct {
	// The instance ID of the route target group member.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-xxxx
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The type of the route target group member.
	//
	// Currently supported types:
	//
	// - **GatewayLoadBalancerEndpoint**
	//
	// In Active-Standby mode, all members of the route target group must have the same type.
	//
	// This parameter is required.
	//
	// example:
	//
	// GatewayLoadBalancerEndpoint
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	// The weight value of the route target group member. Values:
	//
	// - **100**: Indicates the member is a primary instance.
	//
	// - **0**: Indicates the member is a standby instance.
	//
	// The weight value can only be set during creation and cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateRouteTargetGroupRequestRouteTargetMemberList) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteTargetGroupRequestRouteTargetMemberList) GoString() string {
	return s.String()
}

func (s *CreateRouteTargetGroupRequestRouteTargetMemberList) GetMemberId() *string {
	return s.MemberId
}

func (s *CreateRouteTargetGroupRequestRouteTargetMemberList) GetMemberType() *string {
	return s.MemberType
}

func (s *CreateRouteTargetGroupRequestRouteTargetMemberList) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateRouteTargetGroupRequestRouteTargetMemberList) SetMemberId(v string) *CreateRouteTargetGroupRequestRouteTargetMemberList {
	s.MemberId = &v
	return s
}

func (s *CreateRouteTargetGroupRequestRouteTargetMemberList) SetMemberType(v string) *CreateRouteTargetGroupRequestRouteTargetMemberList {
	s.MemberType = &v
	return s
}

func (s *CreateRouteTargetGroupRequestRouteTargetMemberList) SetWeight(v int32) *CreateRouteTargetGroupRequestRouteTargetMemberList {
	s.Weight = &v
	return s
}

func (s *CreateRouteTargetGroupRequestRouteTargetMemberList) Validate() error {
	return dara.Validate(s)
}

type CreateRouteTargetGroupRequestTag struct {
	// The tag key of the resource. Up to 20 tag keys are supported. If you need to pass this value, you cannot input an empty string.
	//
	// A tag key can have up to 128 characters and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. Up to 20 tag values are supported. If you need to pass this value, you can input an empty string.
	//
	// A tag value can have up to 128 characters and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRouteTargetGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteTargetGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRouteTargetGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateRouteTargetGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateRouteTargetGroupRequestTag) SetKey(v string) *CreateRouteTargetGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRouteTargetGroupRequestTag) SetValue(v string) *CreateRouteTargetGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateRouteTargetGroupRequestTag) Validate() error {
	return dara.Validate(s)
}

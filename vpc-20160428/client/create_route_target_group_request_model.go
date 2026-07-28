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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe6****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration mode of the route target group. Valid values:
	//
	// - **Active-Standby**: active/standby mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// Active-Standby
	ConfigMode *string `json:"ConfigMode,omitempty" xml:"ConfigMode,omitempty"`
	// The region ID of the route target group. You can call the DescribeRegions operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [What is a resource group?](https://help.aliyun.com/document_detail/2381067.html).
	//
	// example:
	//
	// rg-acfmxazffggds****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The description of the route target group.
	//
	// The description must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// myRouteTargetGroupDescription
	RouteTargetGroupDescription *string `json:"RouteTargetGroupDescription,omitempty" xml:"RouteTargetGroupDescription,omitempty"`
	// The name of the route target group.
	//
	// The name must be 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// myRouteTargetGroupName
	RouteTargetGroupName *string `json:"RouteTargetGroupName,omitempty" xml:"RouteTargetGroupName,omitempty"`
	// The member list of the route target group.
	//
	// In active/standby mode, the following limits apply to route target group members:
	//
	// 1. The number of route target group members must be 2.
	//
	// 2. The route target group members must be in different zones.
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
	// The member type of the route target group.
	//
	// Supported type:
	//
	// - **GatewayLoadBalancerEndpoint**
	//
	// In active/standby mode, all members of the route target group must be of the same type.
	//
	// This parameter is required.
	//
	// example:
	//
	// GatewayLoadBalancerEndpoint
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	// The weight of the route target group member. Valid values:
	//
	// - **100**: The member is the active instance.
	//
	// - **0**: The member is the standby instance.
	//
	// The weight can only be set during creation and cannot be modified.
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
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
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

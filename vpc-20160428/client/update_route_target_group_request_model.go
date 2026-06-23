// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteTargetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRouteTargetGroupRequest
	GetClientToken() *string
	SetRegionId(v string) *UpdateRouteTargetGroupRequest
	GetRegionId() *string
	SetRouteTargetGroupDescription(v string) *UpdateRouteTargetGroupRequest
	GetRouteTargetGroupDescription() *string
	SetRouteTargetGroupId(v string) *UpdateRouteTargetGroupRequest
	GetRouteTargetGroupId() *string
	SetRouteTargetGroupName(v string) *UpdateRouteTargetGroupRequest
	GetRouteTargetGroupName() *string
	SetRouteTargetMemberList(v []*UpdateRouteTargetGroupRequestRouteTargetMemberList) *UpdateRouteTargetGroupRequest
	GetRouteTargetMemberList() []*UpdateRouteTargetGroupRequestRouteTargetMemberList
}

type UpdateRouteTargetGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the RequestId value as the ClientToken value. The RequestId value may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the route target group instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the route target group.
	//
	// The description must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// myRouteTargetGroupDescription
	RouteTargetGroupDescription *string `json:"RouteTargetGroupDescription,omitempty" xml:"RouteTargetGroupDescription,omitempty"`
	// The routing target group instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtg-xxx
	RouteTargetGroupId *string `json:"RouteTargetGroupId,omitempty" xml:"RouteTargetGroupId,omitempty"`
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
	// 2. The route target group members must belong to different zones.
	RouteTargetMemberList []*UpdateRouteTargetGroupRequestRouteTargetMemberList `json:"RouteTargetMemberList,omitempty" xml:"RouteTargetMemberList,omitempty" type:"Repeated"`
}

func (s UpdateRouteTargetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteTargetGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateRouteTargetGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRouteTargetGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateRouteTargetGroupRequest) GetRouteTargetGroupDescription() *string {
	return s.RouteTargetGroupDescription
}

func (s *UpdateRouteTargetGroupRequest) GetRouteTargetGroupId() *string {
	return s.RouteTargetGroupId
}

func (s *UpdateRouteTargetGroupRequest) GetRouteTargetGroupName() *string {
	return s.RouteTargetGroupName
}

func (s *UpdateRouteTargetGroupRequest) GetRouteTargetMemberList() []*UpdateRouteTargetGroupRequestRouteTargetMemberList {
	return s.RouteTargetMemberList
}

func (s *UpdateRouteTargetGroupRequest) SetClientToken(v string) *UpdateRouteTargetGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRouteTargetGroupRequest) SetRegionId(v string) *UpdateRouteTargetGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateRouteTargetGroupRequest) SetRouteTargetGroupDescription(v string) *UpdateRouteTargetGroupRequest {
	s.RouteTargetGroupDescription = &v
	return s
}

func (s *UpdateRouteTargetGroupRequest) SetRouteTargetGroupId(v string) *UpdateRouteTargetGroupRequest {
	s.RouteTargetGroupId = &v
	return s
}

func (s *UpdateRouteTargetGroupRequest) SetRouteTargetGroupName(v string) *UpdateRouteTargetGroupRequest {
	s.RouteTargetGroupName = &v
	return s
}

func (s *UpdateRouteTargetGroupRequest) SetRouteTargetMemberList(v []*UpdateRouteTargetGroupRequestRouteTargetMemberList) *UpdateRouteTargetGroupRequest {
	s.RouteTargetMemberList = v
	return s
}

func (s *UpdateRouteTargetGroupRequest) Validate() error {
	if s.RouteTargetMemberList != nil {
		for _, item := range s.RouteTargetMemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRouteTargetGroupRequestRouteTargetMemberList struct {
	// The instance ID of the route target group member.
	//
	// example:
	//
	// ep-xxxx
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The member type of the route target group.
	//
	// Currently supported type:
	//
	// - **GatewayLoadBalancerEndpoint**
	//
	// In active/standby mode, all members of the route target group must be of the same type.
	//
	// example:
	//
	// GatewayLoadBalancerEndpoint
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	// The weight of the route target group member. Valid values:
	//
	// - 100: The member is the active instance.
	//
	// - 0: The member is the standby instance.
	//
	// The weight can only be set during creation and cannot be modified.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateRouteTargetGroupRequestRouteTargetMemberList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteTargetGroupRequestRouteTargetMemberList) GoString() string {
	return s.String()
}

func (s *UpdateRouteTargetGroupRequestRouteTargetMemberList) GetMemberId() *string {
	return s.MemberId
}

func (s *UpdateRouteTargetGroupRequestRouteTargetMemberList) GetMemberType() *string {
	return s.MemberType
}

func (s *UpdateRouteTargetGroupRequestRouteTargetMemberList) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateRouteTargetGroupRequestRouteTargetMemberList) SetMemberId(v string) *UpdateRouteTargetGroupRequestRouteTargetMemberList {
	s.MemberId = &v
	return s
}

func (s *UpdateRouteTargetGroupRequestRouteTargetMemberList) SetMemberType(v string) *UpdateRouteTargetGroupRequestRouteTargetMemberList {
	s.MemberType = &v
	return s
}

func (s *UpdateRouteTargetGroupRequestRouteTargetMemberList) SetWeight(v int32) *UpdateRouteTargetGroupRequestRouteTargetMemberList {
	s.Weight = &v
	return s
}

func (s *UpdateRouteTargetGroupRequestRouteTargetMemberList) Validate() error {
	return dara.Validate(s)
}

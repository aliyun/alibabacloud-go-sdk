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
	// Client Token, used to ensure the idempotence of requests. Generate a unique value for this parameter from your client for each request. ClientToken supports only ASCII characters. Note that if you do not specify this, the system will automatically use the RequestId of the API request as the ClientToken identifier. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the region to which the route target group instance belongs. You can obtain the region ID by calling the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Description of the route target group.
	//
	// The description length should be between 1 to 256 characters and must not start with http:// or https://.
	//
	// example:
	//
	// myRouteTargetGroupDescription
	RouteTargetGroupDescription *string `json:"RouteTargetGroupDescription,omitempty" xml:"RouteTargetGroupDescription,omitempty"`
	// The ID of the route target group instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtg-xxx
	RouteTargetGroupId *string `json:"RouteTargetGroupId,omitempty" xml:"RouteTargetGroupId,omitempty"`
	// The name of the route target group.
	//
	// The name length should be between 1 and 128 characters, and cannot start with http:// or https://.
	//
	// example:
	//
	// myRouteTargetGroupName
	RouteTargetGroupName *string `json:"RouteTargetGroupName,omitempty" xml:"RouteTargetGroupName,omitempty"`
	// List of members in the route target group.
	//
	// Under the primary-standby mode, there are the following restrictions on the members of the route target group:
	//
	// 1. The number of members in the route target group must be 2.
	//
	// 2. The members of the route target group must belong to different availability zones.
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
	// ID of the route target group member instance.
	//
	// example:
	//
	// ep-xxxx
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The member type of the route target group.
	//
	// Currently supported types: - **GatewayLoadBalancerEndpoint*	-
	//
	// In active-standby mode, all members of the route target group must be of the same type.
	//
	// example:
	//
	// GatewayLoadBalancerEndpoint
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	// The weight value of the route target group member. Values:
	//
	// - 100: indicates the member is the primary instance.
	//
	// - 0: indicates the member is the backup instance.
	//
	// The weight value can only be set during creation and cannot be modified.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRouteTargetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigMode(v string) *GetRouteTargetGroupResponseBody
	GetConfigMode() *string
	SetCreateTime(v string) *GetRouteTargetGroupResponseBody
	GetCreateTime() *string
	SetRegionId(v string) *GetRouteTargetGroupResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetRouteTargetGroupResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetRouteTargetGroupResponseBody
	GetResourceGroupId() *string
	SetRouteTargetGroupDescription(v string) *GetRouteTargetGroupResponseBody
	GetRouteTargetGroupDescription() *string
	SetRouteTargetGroupId(v string) *GetRouteTargetGroupResponseBody
	GetRouteTargetGroupId() *string
	SetRouteTargetGroupName(v string) *GetRouteTargetGroupResponseBody
	GetRouteTargetGroupName() *string
	SetRouteTargetMemberList(v []*GetRouteTargetGroupResponseBodyRouteTargetMemberList) *GetRouteTargetGroupResponseBody
	GetRouteTargetMemberList() []*GetRouteTargetGroupResponseBodyRouteTargetMemberList
	SetStatus(v string) *GetRouteTargetGroupResponseBody
	GetStatus() *string
	SetTags(v []*GetRouteTargetGroupResponseBodyTags) *GetRouteTargetGroupResponseBody
	GetTags() []*GetRouteTargetGroupResponseBodyTags
	SetVpcId(v string) *GetRouteTargetGroupResponseBody
	GetVpcId() *string
}

type GetRouteTargetGroupResponseBody struct {
	// The configuration mode of the route target group. Valid values:
	//
	// - **Active-Standby**: active/standby mode.
	//
	// example:
	//
	// Active-Standby
	ConfigMode *string `json:"ConfigMode,omitempty" xml:"ConfigMode,omitempty"`
	// The time when the route target group was created.
	//
	// example:
	//
	// 2025-12-30T06:40:50Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The region ID of the VPC to which the route target group belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the region ID.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AE05898-06E5-4782-xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the route target group belongs.
	//
	// example:
	//
	// rg-acfmxazdjdhd****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The description of the route target group.
	//
	// example:
	//
	// myRouteTargetGroupDescription
	RouteTargetGroupDescription *string `json:"RouteTargetGroupDescription,omitempty" xml:"RouteTargetGroupDescription,omitempty"`
	// The instance ID of the routing target group.
	//
	// example:
	//
	// rtg-xxxx
	RouteTargetGroupId *string `json:"RouteTargetGroupId,omitempty" xml:"RouteTargetGroupId,omitempty"`
	// The name of the route target group.
	//
	// example:
	//
	// myRouteTargetGroupName
	RouteTargetGroupName *string `json:"RouteTargetGroupName,omitempty" xml:"RouteTargetGroupName,omitempty"`
	// The member list of the route target group.
	RouteTargetMemberList []*GetRouteTargetGroupResponseBodyRouteTargetMemberList `json:"RouteTargetMemberList,omitempty" xml:"RouteTargetMemberList,omitempty" type:"Repeated"`
	// The status of the routing target group. Valid values:
	//
	// - **Recovering**: The active/standby switchback is in progress.
	//
	// - **Switched**: The active/standby switchover is complete.
	//
	// - **Available**: The routing target group is available.
	//
	// - **Abnormal**: The standby instance has instance failures.
	//
	// - **Pending**: The routing target group is being created.
	//
	// - **Switching**: The active/standby switchover is in progress.
	//
	// - **Deleting**: The routing target group is being deleted.
	//
	// - **Unavailable**: Both the primary and secondary instances have instance failures.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the route target group.
	Tags []*GetRouteTargetGroupResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC to which the route target group belongs.
	//
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetRouteTargetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRouteTargetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetRouteTargetGroupResponseBody) GetConfigMode() *string {
	return s.ConfigMode
}

func (s *GetRouteTargetGroupResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRouteTargetGroupResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRouteTargetGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRouteTargetGroupResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetRouteTargetGroupResponseBody) GetRouteTargetGroupDescription() *string {
	return s.RouteTargetGroupDescription
}

func (s *GetRouteTargetGroupResponseBody) GetRouteTargetGroupId() *string {
	return s.RouteTargetGroupId
}

func (s *GetRouteTargetGroupResponseBody) GetRouteTargetGroupName() *string {
	return s.RouteTargetGroupName
}

func (s *GetRouteTargetGroupResponseBody) GetRouteTargetMemberList() []*GetRouteTargetGroupResponseBodyRouteTargetMemberList {
	return s.RouteTargetMemberList
}

func (s *GetRouteTargetGroupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRouteTargetGroupResponseBody) GetTags() []*GetRouteTargetGroupResponseBodyTags {
	return s.Tags
}

func (s *GetRouteTargetGroupResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetRouteTargetGroupResponseBody) SetConfigMode(v string) *GetRouteTargetGroupResponseBody {
	s.ConfigMode = &v
	return s
}

func (s *GetRouteTargetGroupResponseBody) SetCreateTime(v string) *GetRouteTargetGroupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRouteTargetGroupResponseBody) SetRegionId(v string) *GetRouteTargetGroupResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetRouteTargetGroupResponseBody) SetRequestId(v string) *GetRouteTargetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRouteTargetGroupResponseBody) SetResourceGroupId(v string) *GetRouteTargetGroupResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetRouteTargetGroupResponseBody) SetRouteTargetGroupDescription(v string) *GetRouteTargetGroupResponseBody {
	s.RouteTargetGroupDescription = &v
	return s
}

func (s *GetRouteTargetGroupResponseBody) SetRouteTargetGroupId(v string) *GetRouteTargetGroupResponseBody {
	s.RouteTargetGroupId = &v
	return s
}

func (s *GetRouteTargetGroupResponseBody) SetRouteTargetGroupName(v string) *GetRouteTargetGroupResponseBody {
	s.RouteTargetGroupName = &v
	return s
}

func (s *GetRouteTargetGroupResponseBody) SetRouteTargetMemberList(v []*GetRouteTargetGroupResponseBodyRouteTargetMemberList) *GetRouteTargetGroupResponseBody {
	s.RouteTargetMemberList = v
	return s
}

func (s *GetRouteTargetGroupResponseBody) SetStatus(v string) *GetRouteTargetGroupResponseBody {
	s.Status = &v
	return s
}

func (s *GetRouteTargetGroupResponseBody) SetTags(v []*GetRouteTargetGroupResponseBodyTags) *GetRouteTargetGroupResponseBody {
	s.Tags = v
	return s
}

func (s *GetRouteTargetGroupResponseBody) SetVpcId(v string) *GetRouteTargetGroupResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetRouteTargetGroupResponseBody) Validate() error {
	if s.RouteTargetMemberList != nil {
		for _, item := range s.RouteTargetMemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRouteTargetGroupResponseBodyRouteTargetMemberList struct {
	// The enable status of the route target group member. Valid values:
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// Only route target group members in the disabled state can be replaced with other instances. Route target group members in the enabled state cannot be modified.
	//
	// example:
	//
	// Enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health check status of the route target group member. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **Abnormal**: Abnormal.
	//
	// example:
	//
	// Normal
	HealthCheckStatus *string `json:"HealthCheckStatus,omitempty" xml:"HealthCheckStatus,omitempty"`
	// The instance ID of the routing target group member.
	//
	// example:
	//
	// ep-xxxx
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The member type of the route target group.
	//
	// Currently supported types:
	//
	// - **GatewayLoadBalancerEndpoint**
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
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s GetRouteTargetGroupResponseBodyRouteTargetMemberList) String() string {
	return dara.Prettify(s)
}

func (s GetRouteTargetGroupResponseBodyRouteTargetMemberList) GoString() string {
	return s.String()
}

func (s *GetRouteTargetGroupResponseBodyRouteTargetMemberList) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *GetRouteTargetGroupResponseBodyRouteTargetMemberList) GetHealthCheckStatus() *string {
	return s.HealthCheckStatus
}

func (s *GetRouteTargetGroupResponseBodyRouteTargetMemberList) GetMemberId() *string {
	return s.MemberId
}

func (s *GetRouteTargetGroupResponseBodyRouteTargetMemberList) GetMemberType() *string {
	return s.MemberType
}

func (s *GetRouteTargetGroupResponseBodyRouteTargetMemberList) GetWeight() *int32 {
	return s.Weight
}

func (s *GetRouteTargetGroupResponseBodyRouteTargetMemberList) SetEnableStatus(v string) *GetRouteTargetGroupResponseBodyRouteTargetMemberList {
	s.EnableStatus = &v
	return s
}

func (s *GetRouteTargetGroupResponseBodyRouteTargetMemberList) SetHealthCheckStatus(v string) *GetRouteTargetGroupResponseBodyRouteTargetMemberList {
	s.HealthCheckStatus = &v
	return s
}

func (s *GetRouteTargetGroupResponseBodyRouteTargetMemberList) SetMemberId(v string) *GetRouteTargetGroupResponseBodyRouteTargetMemberList {
	s.MemberId = &v
	return s
}

func (s *GetRouteTargetGroupResponseBodyRouteTargetMemberList) SetMemberType(v string) *GetRouteTargetGroupResponseBodyRouteTargetMemberList {
	s.MemberType = &v
	return s
}

func (s *GetRouteTargetGroupResponseBodyRouteTargetMemberList) SetWeight(v int32) *GetRouteTargetGroupResponseBodyRouteTargetMemberList {
	s.Weight = &v
	return s
}

func (s *GetRouteTargetGroupResponseBodyRouteTargetMemberList) Validate() error {
	return dara.Validate(s)
}

type GetRouteTargetGroupResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRouteTargetGroupResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetRouteTargetGroupResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetRouteTargetGroupResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetRouteTargetGroupResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetRouteTargetGroupResponseBodyTags) SetKey(v string) *GetRouteTargetGroupResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetRouteTargetGroupResponseBodyTags) SetValue(v string) *GetRouteTargetGroupResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetRouteTargetGroupResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

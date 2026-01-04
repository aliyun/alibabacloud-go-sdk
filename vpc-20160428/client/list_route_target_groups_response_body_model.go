// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteTargetGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRouteTargetGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRouteTargetGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRouteTargetGroupsResponseBody
	GetRequestId() *string
	SetRouteTargetGroups(v []*ListRouteTargetGroupsResponseBodyRouteTargetGroups) *ListRouteTargetGroupsResponseBody
	GetRouteTargetGroups() []*ListRouteTargetGroupsResponseBodyRouteTargetGroups
	SetTotalCount(v int32) *ListRouteTargetGroupsResponseBody
	GetTotalCount() *int32
}

type ListRouteTargetGroupsResponseBody struct {
	// The page size.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token for the next query. Value: If NextToken is empty, it indicates there is no next query. If NextToken has a return value, it indicates the token for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// ID of the request
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of route target groups.
	RouteTargetGroups []*ListRouteTargetGroupsResponseBodyRouteTargetGroups `json:"RouteTargetGroups,omitempty" xml:"RouteTargetGroups,omitempty" type:"Repeated"`
	// Number of items in the list.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRouteTargetGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRouteTargetGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRouteTargetGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRouteTargetGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRouteTargetGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRouteTargetGroupsResponseBody) GetRouteTargetGroups() []*ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	return s.RouteTargetGroups
}

func (s *ListRouteTargetGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRouteTargetGroupsResponseBody) SetMaxResults(v int32) *ListRouteTargetGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBody) SetNextToken(v string) *ListRouteTargetGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBody) SetRequestId(v string) *ListRouteTargetGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBody) SetRouteTargetGroups(v []*ListRouteTargetGroupsResponseBodyRouteTargetGroups) *ListRouteTargetGroupsResponseBody {
	s.RouteTargetGroups = v
	return s
}

func (s *ListRouteTargetGroupsResponseBody) SetTotalCount(v int32) *ListRouteTargetGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBody) Validate() error {
	if s.RouteTargetGroups != nil {
		for _, item := range s.RouteTargetGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRouteTargetGroupsResponseBodyRouteTargetGroups struct {
	// The configuration mode of the route target group. Supported modes are as follows:
	//
	// - **Active-Standby**: Active-standby mode.
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
	// The region ID of the VPC to which the route target group belongs.
	//
	// You can obtain the region ID by calling the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) interface.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the route target group belongs.
	//
	// example:
	//
	// rg-acfm3swh6ta56ri
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Description of the route target group.
	//
	// example:
	//
	// myRouteTargetGroupDescription
	RouteTargetGroupDescription *string `json:"RouteTargetGroupDescription,omitempty" xml:"RouteTargetGroupDescription,omitempty"`
	// The ID of the route target group instance.
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
	// The list of route target group members.
	RouteTargetMemberList []*ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList `json:"RouteTargetMemberList,omitempty" xml:"RouteTargetMemberList,omitempty" type:"Repeated"`
	// Status of the route target group. Values:
	//
	// - **Recovering**: Active-Standby rollback in progress
	//
	// - **Switched**: Active-Standby switched
	//
	// - **Available**: Available
	//
	// - **Abnormal**: Standby instance abnormal
	//
	// - **Pending**: Creating
	//
	// - **Switching**: Active-Standby switching in progress
	//
	// - **Deleting**: Deleting
	//
	// - **Unavailable**: Both primary and standby instances are abnormal
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag values. A maximum of 20 tag values are supported. If you need to pass this value, you can input an empty string.
	//
	// A maximum of 128 characters are supported. The value cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	Tags []*ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC to which the route target group belongs.
	//
	// example:
	//
	// vpc-uf60y8uzhsvbhmuh3l654
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListRouteTargetGroupsResponseBodyRouteTargetGroups) String() string {
	return dara.Prettify(s)
}

func (s ListRouteTargetGroupsResponseBodyRouteTargetGroups) GoString() string {
	return s.String()
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) GetConfigMode() *string {
	return s.ConfigMode
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) GetRouteTargetGroupDescription() *string {
	return s.RouteTargetGroupDescription
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) GetRouteTargetGroupId() *string {
	return s.RouteTargetGroupId
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) GetRouteTargetGroupName() *string {
	return s.RouteTargetGroupName
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) GetRouteTargetMemberList() []*ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList {
	return s.RouteTargetMemberList
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) GetStatus() *string {
	return s.Status
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) GetTags() []*ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags {
	return s.Tags
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) GetVpcId() *string {
	return s.VpcId
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) SetConfigMode(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	s.ConfigMode = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) SetCreateTime(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	s.CreateTime = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) SetRegionId(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	s.RegionId = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) SetResourceGroupId(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) SetRouteTargetGroupDescription(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	s.RouteTargetGroupDescription = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) SetRouteTargetGroupId(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	s.RouteTargetGroupId = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) SetRouteTargetGroupName(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	s.RouteTargetGroupName = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) SetRouteTargetMemberList(v []*ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) *ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	s.RouteTargetMemberList = v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) SetStatus(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	s.Status = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) SetTags(v []*ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags) *ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	s.Tags = v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) SetVpcId(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroups {
	s.VpcId = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroups) Validate() error {
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

type ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList struct {
	// The enable status of the route target group member. Values:
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// Only disabled route target group members can be modified to other instances. Enabled route target group members cannot be modified.
	//
	// example:
	//
	// Enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health check status of the route target group member. Values:
	//
	// - **Normal**: Normal
	//
	// - **Abnormal**: Abnormal
	//
	// example:
	//
	// Normal
	HealthCheckStatus *string `json:"HealthCheckStatus,omitempty" xml:"HealthCheckStatus,omitempty"`
	// The ID of the route target group member instance.
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
	// example:
	//
	// GatewayLoadBalancerEndpoint
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	// The weight value of the route target group member. Values:
	//
	// - **100**: Indicates that the member is the primary instance.
	//
	// - **0**: Indicates that the member is the backup instance.
	//
	// The weight value can only be set during creation and cannot be modified.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) String() string {
	return dara.Prettify(s)
}

func (s ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) GoString() string {
	return s.String()
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) GetHealthCheckStatus() *string {
	return s.HealthCheckStatus
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) GetMemberId() *string {
	return s.MemberId
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) GetMemberType() *string {
	return s.MemberType
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) GetWeight() *int32 {
	return s.Weight
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) SetEnableStatus(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList {
	s.EnableStatus = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) SetHealthCheckStatus(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList {
	s.HealthCheckStatus = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) SetMemberId(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList {
	s.MemberId = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) SetMemberType(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList {
	s.MemberType = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) SetWeight(v int32) *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList {
	s.Weight = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsRouteTargetMemberList) Validate() error {
	return dara.Validate(s)
}

type ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags struct {
	// The key of the resource tag.
	//
	// example:
	//
	// image/upload/cbbec42e0be33abb27babefcbe0397f0
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the resource tag. Up to 20 tag values are supported. If you need to pass this value, you can input an empty string. A maximum of 128 characters is allowed. The value cannot start with `aliyun` or `acs:`, and it must not contain `http://` or `https://`.
	//
	// example:
	//
	// 8
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags) GoString() string {
	return s.String()
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags) GetKey() *string {
	return s.Key
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags) GetValue() *string {
	return s.Value
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags) SetKey(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags {
	s.Key = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags) SetValue(v string) *ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags {
	s.Value = &v
	return s
}

func (s *ListRouteTargetGroupsResponseBodyRouteTargetGroupsTags) Validate() error {
	return dara.Validate(s)
}

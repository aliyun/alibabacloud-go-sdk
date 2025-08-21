// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSecurityGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSecurityGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSecurityGroupsResponseBody
	GetRequestId() *string
	SetSecurityGroups(v *DescribeSecurityGroupsResponseBodySecurityGroups) *DescribeSecurityGroupsResponseBody
	GetSecurityGroups() *DescribeSecurityGroupsResponseBodySecurityGroups
	SetTotalCount(v int32) *DescribeSecurityGroupsResponseBody
	GetTotalCount() *int32
}

type DescribeSecurityGroupsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about security groups.
	SecurityGroups *DescribeSecurityGroupsResponseBodySecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Struct"`
	// The total number of returned pages.
	//
	// example:
	//
	// 49
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSecurityGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityGroupsResponseBody) GetSecurityGroups() *DescribeSecurityGroupsResponseBodySecurityGroups {
	return s.SecurityGroups
}

func (s *DescribeSecurityGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSecurityGroupsResponseBody) SetPageNumber(v int32) *DescribeSecurityGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetPageSize(v int32) *DescribeSecurityGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetRequestId(v string) *DescribeSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetSecurityGroups(v *DescribeSecurityGroupsResponseBodySecurityGroups) *DescribeSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetTotalCount(v int32) *DescribeSecurityGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityGroupsResponseBodySecurityGroups struct {
	SecurityGroup []*DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroups) GetSecurityGroup() []*DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	return s.SecurityGroup
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroups) SetSecurityGroup(v []*DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) *DescribeSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroup = v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup struct {
	// The creation time. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-11-01T06:08:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the security group.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of associated instances.
	//
	// example:
	//
	// 5
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The IDs of the instances that are associated with the security group.
	InstanceIds *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// The IDs of the ENIs that are associated with the security group.
	NetworkInterfaceIds *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupNetworkInterfaceIds `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Struct"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp67acfmxazb4ph***
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the security group.
	//
	// example:
	//
	// DocTest
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetInstanceIds() *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupInstanceIds {
	return s.InstanceIds
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetNetworkInterfaceIds() *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupNetworkInterfaceIds {
	return s.NetworkInterfaceIds
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetCreationTime(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.CreationTime = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetDescription(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.Description = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetInstanceCount(v int32) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.InstanceCount = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetInstanceIds(v *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupInstanceIds) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.InstanceIds = v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetNetworkInterfaceIds(v *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupNetworkInterfaceIds) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.NetworkInterfaceIds = v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetSecurityGroupId(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetSecurityGroupName(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupInstanceIds) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupInstanceIds) SetInstanceId(v []*string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupInstanceIds {
	s.InstanceId = v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupNetworkInterfaceIds struct {
	NetworkInterfaceId []*string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupNetworkInterfaceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupNetworkInterfaceIds) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupNetworkInterfaceIds) GetNetworkInterfaceId() []*string {
	return s.NetworkInterfaceId
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupNetworkInterfaceIds) SetNetworkInterfaceId(v []*string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupNetworkInterfaceIds {
	s.NetworkInterfaceId = v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupNetworkInterfaceIds) Validate() error {
	return dara.Validate(s)
}

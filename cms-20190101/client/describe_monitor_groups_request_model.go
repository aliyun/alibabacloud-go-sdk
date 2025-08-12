// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicTagRuleId(v string) *DescribeMonitorGroupsRequest
	GetDynamicTagRuleId() *string
	SetGroupFounderTagKey(v string) *DescribeMonitorGroupsRequest
	GetGroupFounderTagKey() *string
	SetGroupFounderTagValue(v string) *DescribeMonitorGroupsRequest
	GetGroupFounderTagValue() *string
	SetGroupId(v string) *DescribeMonitorGroupsRequest
	GetGroupId() *string
	SetGroupName(v string) *DescribeMonitorGroupsRequest
	GetGroupName() *string
	SetIncludeTemplateHistory(v bool) *DescribeMonitorGroupsRequest
	GetIncludeTemplateHistory() *bool
	SetInstanceId(v string) *DescribeMonitorGroupsRequest
	GetInstanceId() *string
	SetKeyword(v string) *DescribeMonitorGroupsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *DescribeMonitorGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMonitorGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMonitorGroupsRequest
	GetRegionId() *string
	SetSelectContactGroups(v bool) *DescribeMonitorGroupsRequest
	GetSelectContactGroups() *bool
	SetTag(v []*DescribeMonitorGroupsRequestTag) *DescribeMonitorGroupsRequest
	GetTag() []*DescribeMonitorGroupsRequestTag
	SetType(v string) *DescribeMonitorGroupsRequest
	GetType() *string
	SetTypes(v string) *DescribeMonitorGroupsRequest
	GetTypes() *string
}

type DescribeMonitorGroupsRequest struct {
	// The ID of the tag rule.
	//
	// example:
	//
	// 6b882d9a-5117-42e2-9d0c-4749a0c6****
	DynamicTagRuleId *string `json:"DynamicTagRuleId,omitempty" xml:"DynamicTagRuleId,omitempty"`
	// The tag key that is created for the application group by using the tag rule.
	//
	// example:
	//
	// GroupKey1
	GroupFounderTagKey *string `json:"GroupFounderTagKey,omitempty" xml:"GroupFounderTagKey,omitempty"`
	// The tag value that is created for the application group by using the tag rule.
	//
	// example:
	//
	// GroupValue1
	GroupFounderTagValue *string `json:"GroupFounderTagValue,omitempty" xml:"GroupFounderTagValue,omitempty"`
	// The ID of the application group. Separate multiple application group IDs with commas (,).
	//
	// example:
	//
	// 92****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// testGroup124
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Specifies whether to include the historical alert templates that are applied to the application group in the response. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IncludeTemplateHistory *bool `json:"IncludeTemplateHistory,omitempty" xml:"IncludeTemplateHistory,omitempty"`
	// The instance ID. This parameter is used to query the application group to which the specified instance belongs.
	//
	// example:
	//
	// i-abcdefgh12****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The keyword that is used for the search.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Pages start from page 1. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to include the alert contact groups in the response. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SelectContactGroups *bool `json:"SelectContactGroups,omitempty" xml:"SelectContactGroups,omitempty"`
	// The tags of the application group.
	Tag []*DescribeMonitorGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the application group. Valid values:
	//
	// 	- custom: a self-managed application group
	//
	// 	- ehpc_cluster: an application group that is synchronized from an E-HPC cluster
	//
	// 	- kubernetes: an application group that is synchronized from an ACK cluster
	//
	// example:
	//
	// custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The type of the application group. Valid values:
	//
	// 	- custom: a self-managed application group
	//
	// 	- ehpc_cluster: an application group that is synchronized from an Elastic High Performance Computing (E-HPC) cluster
	//
	// 	- kubernetes: an application group that is synchronized from a Container Service for Kubernetes (ACK) cluster
	//
	// 	- tag: an application group that is automatically created by using tags
	//
	// 	- resMgr: an application group that is created by using resource groups
	//
	// 	- ess: an application group that is synchronized from Auto Scaling (ESS)
	//
	// example:
	//
	// custom
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s DescribeMonitorGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsRequest) GetDynamicTagRuleId() *string {
	return s.DynamicTagRuleId
}

func (s *DescribeMonitorGroupsRequest) GetGroupFounderTagKey() *string {
	return s.GroupFounderTagKey
}

func (s *DescribeMonitorGroupsRequest) GetGroupFounderTagValue() *string {
	return s.GroupFounderTagValue
}

func (s *DescribeMonitorGroupsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeMonitorGroupsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeMonitorGroupsRequest) GetIncludeTemplateHistory() *bool {
	return s.IncludeTemplateHistory
}

func (s *DescribeMonitorGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitorGroupsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeMonitorGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMonitorGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMonitorGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitorGroupsRequest) GetSelectContactGroups() *bool {
	return s.SelectContactGroups
}

func (s *DescribeMonitorGroupsRequest) GetTag() []*DescribeMonitorGroupsRequestTag {
	return s.Tag
}

func (s *DescribeMonitorGroupsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeMonitorGroupsRequest) GetTypes() *string {
	return s.Types
}

func (s *DescribeMonitorGroupsRequest) SetDynamicTagRuleId(v string) *DescribeMonitorGroupsRequest {
	s.DynamicTagRuleId = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetGroupFounderTagKey(v string) *DescribeMonitorGroupsRequest {
	s.GroupFounderTagKey = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetGroupFounderTagValue(v string) *DescribeMonitorGroupsRequest {
	s.GroupFounderTagValue = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetGroupId(v string) *DescribeMonitorGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetGroupName(v string) *DescribeMonitorGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetIncludeTemplateHistory(v bool) *DescribeMonitorGroupsRequest {
	s.IncludeTemplateHistory = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetInstanceId(v string) *DescribeMonitorGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetKeyword(v string) *DescribeMonitorGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetPageNumber(v int32) *DescribeMonitorGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetPageSize(v int32) *DescribeMonitorGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetRegionId(v string) *DescribeMonitorGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetSelectContactGroups(v bool) *DescribeMonitorGroupsRequest {
	s.SelectContactGroups = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetTag(v []*DescribeMonitorGroupsRequestTag) *DescribeMonitorGroupsRequest {
	s.Tag = v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetType(v string) *DescribeMonitorGroupsRequest {
	s.Type = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetTypes(v string) *DescribeMonitorGroupsRequest {
	s.Types = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupsRequestTag struct {
	// The tag key of the application group. Valid values of N: 1 to 5.
	//
	// example:
	//
	// tagKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the application group. Valid values of N: 1 to 5.
	//
	// example:
	//
	// tagValue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMonitorGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeMonitorGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeMonitorGroupsRequestTag) SetKey(v string) *DescribeMonitorGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeMonitorGroupsRequestTag) SetValue(v string) *DescribeMonitorGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeMonitorGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}

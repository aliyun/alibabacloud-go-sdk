// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleBlackListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeMetricRuleBlackListRequest
	GetCategory() *string
	SetIds(v []*string) *DescribeMetricRuleBlackListRequest
	GetIds() []*string
	SetInstanceIds(v []*string) *DescribeMetricRuleBlackListRequest
	GetInstanceIds() []*string
	SetIsEnable(v bool) *DescribeMetricRuleBlackListRequest
	GetIsEnable() *bool
	SetName(v string) *DescribeMetricRuleBlackListRequest
	GetName() *string
	SetNamespace(v string) *DescribeMetricRuleBlackListRequest
	GetNamespace() *string
	SetOrder(v int32) *DescribeMetricRuleBlackListRequest
	GetOrder() *int32
	SetPageNumber(v int32) *DescribeMetricRuleBlackListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMetricRuleBlackListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMetricRuleBlackListRequest
	GetRegionId() *string
	SetScopeType(v string) *DescribeMetricRuleBlackListRequest
	GetScopeType() *string
}

type DescribeMetricRuleBlackListRequest struct {
	// The category of the Alibaba Cloud service. For example, Redis has different editions, such as `kvstore_standard` (Standard Edition), `kvstore_sharding` (Cluster Edition), and `kvstore_splitrw` (Read/write Splitting Edition).
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The IDs of the blacklist policies.
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The IDs of instances in the blacklist policy.
	//
	// The value of N can be an integer from 0 to 10.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The status of the blacklist policy. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// example:
	//
	// true
	IsEnable *bool `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	// The name of the blacklist policy.
	//
	// Fuzzy queries are supported.
	//
	// example:
	//
	// Blacklist-01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the Alibaba Cloud service.
	//
	// For more information, see [Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The order in which to sort the results by time. Valid values:
	//
	// - DESC (default): descending order.
	//
	// - ASC: ascending order.
	//
	// example:
	//
	// DESC
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 100
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scope of the blacklist policy. Valid values:
	//
	// - USER: The blacklist policy takes effect only for the current Alibaba Cloud account.
	//
	// - GROUP: The blacklist policy takes effect for the specified application groups.
	//
	// example:
	//
	// USER
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s DescribeMetricRuleBlackListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleBlackListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleBlackListRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeMetricRuleBlackListRequest) GetIds() []*string {
	return s.Ids
}

func (s *DescribeMetricRuleBlackListRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeMetricRuleBlackListRequest) GetIsEnable() *bool {
	return s.IsEnable
}

func (s *DescribeMetricRuleBlackListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeMetricRuleBlackListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricRuleBlackListRequest) GetOrder() *int32 {
	return s.Order
}

func (s *DescribeMetricRuleBlackListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMetricRuleBlackListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMetricRuleBlackListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricRuleBlackListRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *DescribeMetricRuleBlackListRequest) SetCategory(v string) *DescribeMetricRuleBlackListRequest {
	s.Category = &v
	return s
}

func (s *DescribeMetricRuleBlackListRequest) SetIds(v []*string) *DescribeMetricRuleBlackListRequest {
	s.Ids = v
	return s
}

func (s *DescribeMetricRuleBlackListRequest) SetInstanceIds(v []*string) *DescribeMetricRuleBlackListRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeMetricRuleBlackListRequest) SetIsEnable(v bool) *DescribeMetricRuleBlackListRequest {
	s.IsEnable = &v
	return s
}

func (s *DescribeMetricRuleBlackListRequest) SetName(v string) *DescribeMetricRuleBlackListRequest {
	s.Name = &v
	return s
}

func (s *DescribeMetricRuleBlackListRequest) SetNamespace(v string) *DescribeMetricRuleBlackListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleBlackListRequest) SetOrder(v int32) *DescribeMetricRuleBlackListRequest {
	s.Order = &v
	return s
}

func (s *DescribeMetricRuleBlackListRequest) SetPageNumber(v int32) *DescribeMetricRuleBlackListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMetricRuleBlackListRequest) SetPageSize(v int32) *DescribeMetricRuleBlackListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetricRuleBlackListRequest) SetRegionId(v string) *DescribeMetricRuleBlackListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricRuleBlackListRequest) SetScopeType(v string) *DescribeMetricRuleBlackListRequest {
	s.ScopeType = &v
	return s
}

func (s *DescribeMetricRuleBlackListRequest) Validate() error {
	return dara.Validate(s)
}

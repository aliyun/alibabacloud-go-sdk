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
	// The ID of the blacklist policy.
	//
	// example:
	//
	// ecs
	Category *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	Ids      []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The IDs of the instances in the blacklist policy.
	//
	// Valid values of N: 0 to 10.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The status of the blacklist policy. Valid values:
	//
	// 	- true: The blacklist policy is enabled.
	//
	// 	- false: The blacklist policy is disabled.
	//
	// example:
	//
	// true
	IsEnable *bool `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	// The name of the blacklist policy.
	//
	// This parameter supports fuzzy match.
	//
	// example:
	//
	// Blacklist-01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The timestamp when the blacklist policy expired.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// DESC
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The categories of the Alibaba Cloud service. For example, ApsaraDB for Redis includes the following categories: ApsaraDB for Redis (standard architecture), ApsaraDB for Redis (cluster architecture), and ApsaraDB for Redis (read/write splitting architecture). In this case, the valid values of this parameter for ApsaraDB for Redis include `kvstore_standard`, `kvstore_sharding`, and `kvstore_splitrw`.
	//
	// example:
	//
	// 100
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The effective scope of the blacklist policy. Valid values:
	//
	// 	- USER: The blacklist policy takes effect only within the current Alibaba Cloud account.
	//
	// 	- GROUP: The blacklist policy takes effect only within the specified application group.
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

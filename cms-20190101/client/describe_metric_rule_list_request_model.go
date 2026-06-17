// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertState(v string) *DescribeMetricRuleListRequest
	GetAlertState() *string
	SetDimensions(v string) *DescribeMetricRuleListRequest
	GetDimensions() *string
	SetEnableState(v bool) *DescribeMetricRuleListRequest
	GetEnableState() *bool
	SetGroupId(v string) *DescribeMetricRuleListRequest
	GetGroupId() *string
	SetMetricName(v string) *DescribeMetricRuleListRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeMetricRuleListRequest
	GetNamespace() *string
	SetPage(v int32) *DescribeMetricRuleListRequest
	GetPage() *int32
	SetPageSize(v int32) *DescribeMetricRuleListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMetricRuleListRequest
	GetRegionId() *string
	SetRuleIds(v string) *DescribeMetricRuleListRequest
	GetRuleIds() *string
	SetRuleName(v string) *DescribeMetricRuleListRequest
	GetRuleName() *string
}

type DescribeMetricRuleListRequest struct {
	// The state of the alert rule. Valid values:
	//
	// - OK: Normal.
	//
	// - ALARM: Alert.
	//
	// - INSUFFICIENT_DATA: No data.
	//
	// example:
	//
	// OK
	AlertState *string `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	// The monitoring dimensions of the specified resource.
	//
	// Format: a collection of `key:value` pairs. Example: `{"userId":"120886317861****"}` and `{"instanceId":"i-2ze2d6j5uhg20x47****"}`.
	//
	// example:
	//
	// {"instanceId":"i-2ze2d6j5uhg20x47****"}
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The enabled state of the alert rule. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// example:
	//
	// true
	EnableState *bool `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	// The ID of the application group.
	//
	// For information about how to obtain the application group ID, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// example:
	//
	// 7301****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The metric name.
	//
	// For information about how to obtain metric names, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// Note: For composite metrics, you can retrieve results only by the first metric.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the Alibaba Cloud service.
	//
	// For information about how to obtain the namespace of an Alibaba Cloud service, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number.
	//
	// Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the alert rules. Separate multiple IDs with commas (,). You can query up to 20 alert rules at a time.
	//
	// example:
	//
	// applyTemplate344cfd42-0f32-4fd6-805a-88d7908a****
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	// The name of the alert rule.
	//
	// This parameter supports fuzzy match.
	//
	// example:
	//
	// Rule_01
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeMetricRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListRequest) GetAlertState() *string {
	return s.AlertState
}

func (s *DescribeMetricRuleListRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeMetricRuleListRequest) GetEnableState() *bool {
	return s.EnableState
}

func (s *DescribeMetricRuleListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeMetricRuleListRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricRuleListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricRuleListRequest) GetPage() *int32 {
	return s.Page
}

func (s *DescribeMetricRuleListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMetricRuleListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricRuleListRequest) GetRuleIds() *string {
	return s.RuleIds
}

func (s *DescribeMetricRuleListRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeMetricRuleListRequest) SetAlertState(v string) *DescribeMetricRuleListRequest {
	s.AlertState = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetDimensions(v string) *DescribeMetricRuleListRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetEnableState(v bool) *DescribeMetricRuleListRequest {
	s.EnableState = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetGroupId(v string) *DescribeMetricRuleListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetMetricName(v string) *DescribeMetricRuleListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetNamespace(v string) *DescribeMetricRuleListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetPage(v int32) *DescribeMetricRuleListRequest {
	s.Page = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetPageSize(v int32) *DescribeMetricRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetRegionId(v string) *DescribeMetricRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetRuleIds(v string) *DescribeMetricRuleListRequest {
	s.RuleIds = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetRuleName(v string) *DescribeMetricRuleListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeMetricRuleListRequest) Validate() error {
	return dara.Validate(s)
}

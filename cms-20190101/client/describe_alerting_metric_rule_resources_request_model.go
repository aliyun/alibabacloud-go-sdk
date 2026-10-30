// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertingMetricRuleResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertBeforeTime(v string) *DescribeAlertingMetricRuleResourcesRequest
	GetAlertBeforeTime() *string
	SetDimensions(v string) *DescribeAlertingMetricRuleResourcesRequest
	GetDimensions() *string
	SetGroupId(v string) *DescribeAlertingMetricRuleResourcesRequest
	GetGroupId() *string
	SetNamespace(v string) *DescribeAlertingMetricRuleResourcesRequest
	GetNamespace() *string
	SetPage(v int32) *DescribeAlertingMetricRuleResourcesRequest
	GetPage() *int32
	SetPageSize(v int32) *DescribeAlertingMetricRuleResourcesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAlertingMetricRuleResourcesRequest
	GetRegionId() *string
	SetRuleId(v string) *DescribeAlertingMetricRuleResourcesRequest
	GetRuleId() *string
}

type DescribeAlertingMetricRuleResourcesRequest struct {
	// Queries alerts that occurred before this point in time. The value is a millisecond-level timestamp.
	//
	// example:
	//
	// 1698827400000
	AlertBeforeTime *string `json:"AlertBeforeTime,omitempty" xml:"AlertBeforeTime,omitempty"`
	// The dimension map used to query monitoring data of a specified resource.
	//
	// example:
	//
	// {\\"userId\\":\\"120886317861****\\",\\"region\\":\\"cn-huhehaote\\",\\"queue\\":\\"test-0128\\"}
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The ID of the application group. For more information about how to query application group IDs, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// example:
	//
	// 7671****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The data namespace of the Alibaba Cloud service.
	//
	// For more information about how to obtain the data namespace of an Alibaba Cloud service, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// acs_mns_new
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert rule. For more information about how to query alert rule IDs, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// example:
	//
	// putNewAlarm_user_7e78d765-0e3e-4671-ba6d-7ce39108****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeAlertingMetricRuleResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertingMetricRuleResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertingMetricRuleResourcesRequest) GetAlertBeforeTime() *string {
	return s.AlertBeforeTime
}

func (s *DescribeAlertingMetricRuleResourcesRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeAlertingMetricRuleResourcesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeAlertingMetricRuleResourcesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeAlertingMetricRuleResourcesRequest) GetPage() *int32 {
	return s.Page
}

func (s *DescribeAlertingMetricRuleResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertingMetricRuleResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertingMetricRuleResourcesRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeAlertingMetricRuleResourcesRequest) SetAlertBeforeTime(v string) *DescribeAlertingMetricRuleResourcesRequest {
	s.AlertBeforeTime = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesRequest) SetDimensions(v string) *DescribeAlertingMetricRuleResourcesRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesRequest) SetGroupId(v string) *DescribeAlertingMetricRuleResourcesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesRequest) SetNamespace(v string) *DescribeAlertingMetricRuleResourcesRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesRequest) SetPage(v int32) *DescribeAlertingMetricRuleResourcesRequest {
	s.Page = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesRequest) SetPageSize(v int32) *DescribeAlertingMetricRuleResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesRequest) SetRegionId(v string) *DescribeAlertingMetricRuleResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesRequest) SetRuleId(v string) *DescribeAlertingMetricRuleResourcesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertingMetricRuleResourcesRequest) Validate() error {
	return dara.Validate(s)
}

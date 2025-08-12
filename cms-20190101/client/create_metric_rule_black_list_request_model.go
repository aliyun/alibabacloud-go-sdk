// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetricRuleBlackListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateMetricRuleBlackListRequest
	GetCategory() *string
	SetEffectiveTime(v string) *CreateMetricRuleBlackListRequest
	GetEffectiveTime() *string
	SetEnableEndTime(v string) *CreateMetricRuleBlackListRequest
	GetEnableEndTime() *string
	SetEnableStartTime(v string) *CreateMetricRuleBlackListRequest
	GetEnableStartTime() *string
	SetInstances(v []*string) *CreateMetricRuleBlackListRequest
	GetInstances() []*string
	SetMetrics(v []*CreateMetricRuleBlackListRequestMetrics) *CreateMetricRuleBlackListRequest
	GetMetrics() []*CreateMetricRuleBlackListRequestMetrics
	SetName(v string) *CreateMetricRuleBlackListRequest
	GetName() *string
	SetNamespace(v string) *CreateMetricRuleBlackListRequest
	GetNamespace() *string
	SetRegionId(v string) *CreateMetricRuleBlackListRequest
	GetRegionId() *string
	SetScopeType(v string) *CreateMetricRuleBlackListRequest
	GetScopeType() *string
	SetScopeValue(v string) *CreateMetricRuleBlackListRequest
	GetScopeValue() *string
}

type CreateMetricRuleBlackListRequest struct {
	// The category of the cloud service. For example, ApsaraDB for Redis includes the following categories: ApsaraDB for Redis (standard architecture), ApsaraDB for Redis (cluster architecture), and ApsaraDB for Redis (read/write splitting architecture). In this case, the valid values of this parameter for ApsaraDB for Redis include `kvstore_standard`, `kvstore_sharding`, and `kvstore_splitrw`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time range within which the blacklist policy is effective.
	//
	// 	- If you do not configure this parameter, the blacklist policy is permanently effective.
	//
	// 	- If you configure this parameter, the blacklist policy is effective only within the specified time range. Examples:
	//
	//     	- `03:00-04:59`: The blacklist policy is effective from 03:00 to 05:00 local time. 05:00 local time is excluded.
	//
	//     	- `03:00-04:59 UTC+0700`: The blacklist policy is effective from 03:00 to 05:00 (UTC+7). 05:00 (UTC+7) is excluded.
	//
	// example:
	//
	// 03:00-04:59
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The timestamp when the blacklist policy expires.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1640608200000
	EnableEndTime *string `json:"EnableEndTime,omitempty" xml:"EnableEndTime,omitempty"`
	// The timestamp when the blacklist policy starts to take effect.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1640237400000
	EnableStartTime *string `json:"EnableStartTime,omitempty" xml:"EnableStartTime,omitempty"`
	// The IDs of the instances that belong to the specified cloud service.
	//
	// This parameter is required.
	Instances []*string `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The metrics of the instance.
	//
	// 	- If you do not configure this parameter, the blacklist policy applies to all metrics of the specified cloud service.
	//
	// 	- If you configure this parameter, the blacklist policy applies only to the current metric.
	Metrics []*CreateMetricRuleBlackListRequestMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The name of the blacklist policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// Blacklist-01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the cloud service.
	//
	// For more information about the namespaces of different cloud services, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The effective scope of the blacklist policy. Valid values:
	//
	// 	- USER (default): The blacklist policy takes effect only for the current Alibaba Cloud account.
	//
	// 	- GROUP: The blacklist policy takes effect only for the specified application group. For information about how to query the IDs of application groups, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// example:
	//
	// USER
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The ID of the application group. The value of this parameter is a JSON array.
	//
	// > This parameter must be specified when `ScopeType` is set to `GROUP`.
	//
	// example:
	//
	// ["67****","78****"]
	ScopeValue *string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty"`
}

func (s CreateMetricRuleBlackListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleBlackListRequest) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleBlackListRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateMetricRuleBlackListRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *CreateMetricRuleBlackListRequest) GetEnableEndTime() *string {
	return s.EnableEndTime
}

func (s *CreateMetricRuleBlackListRequest) GetEnableStartTime() *string {
	return s.EnableStartTime
}

func (s *CreateMetricRuleBlackListRequest) GetInstances() []*string {
	return s.Instances
}

func (s *CreateMetricRuleBlackListRequest) GetMetrics() []*CreateMetricRuleBlackListRequestMetrics {
	return s.Metrics
}

func (s *CreateMetricRuleBlackListRequest) GetName() *string {
	return s.Name
}

func (s *CreateMetricRuleBlackListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateMetricRuleBlackListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMetricRuleBlackListRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *CreateMetricRuleBlackListRequest) GetScopeValue() *string {
	return s.ScopeValue
}

func (s *CreateMetricRuleBlackListRequest) SetCategory(v string) *CreateMetricRuleBlackListRequest {
	s.Category = &v
	return s
}

func (s *CreateMetricRuleBlackListRequest) SetEffectiveTime(v string) *CreateMetricRuleBlackListRequest {
	s.EffectiveTime = &v
	return s
}

func (s *CreateMetricRuleBlackListRequest) SetEnableEndTime(v string) *CreateMetricRuleBlackListRequest {
	s.EnableEndTime = &v
	return s
}

func (s *CreateMetricRuleBlackListRequest) SetEnableStartTime(v string) *CreateMetricRuleBlackListRequest {
	s.EnableStartTime = &v
	return s
}

func (s *CreateMetricRuleBlackListRequest) SetInstances(v []*string) *CreateMetricRuleBlackListRequest {
	s.Instances = v
	return s
}

func (s *CreateMetricRuleBlackListRequest) SetMetrics(v []*CreateMetricRuleBlackListRequestMetrics) *CreateMetricRuleBlackListRequest {
	s.Metrics = v
	return s
}

func (s *CreateMetricRuleBlackListRequest) SetName(v string) *CreateMetricRuleBlackListRequest {
	s.Name = &v
	return s
}

func (s *CreateMetricRuleBlackListRequest) SetNamespace(v string) *CreateMetricRuleBlackListRequest {
	s.Namespace = &v
	return s
}

func (s *CreateMetricRuleBlackListRequest) SetRegionId(v string) *CreateMetricRuleBlackListRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMetricRuleBlackListRequest) SetScopeType(v string) *CreateMetricRuleBlackListRequest {
	s.ScopeType = &v
	return s
}

func (s *CreateMetricRuleBlackListRequest) SetScopeValue(v string) *CreateMetricRuleBlackListRequest {
	s.ScopeValue = &v
	return s
}

func (s *CreateMetricRuleBlackListRequest) Validate() error {
	return dara.Validate(s)
}

type CreateMetricRuleBlackListRequestMetrics struct {
	// The metric name.
	//
	// Valid values of N: 1 to 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// disk_utilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The extended dimension of the instance. For example, `{"device":"C:"}` specifies that the blacklist policy is applied to all C disks of the specified Elastic Compute Service (ECS) instance.
	//
	// Valid values of N: 1 to 10.
	//
	// example:
	//
	// {"device":"C:"}
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s CreateMetricRuleBlackListRequestMetrics) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleBlackListRequestMetrics) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleBlackListRequestMetrics) GetMetricName() *string {
	return s.MetricName
}

func (s *CreateMetricRuleBlackListRequestMetrics) GetResource() *string {
	return s.Resource
}

func (s *CreateMetricRuleBlackListRequestMetrics) SetMetricName(v string) *CreateMetricRuleBlackListRequestMetrics {
	s.MetricName = &v
	return s
}

func (s *CreateMetricRuleBlackListRequestMetrics) SetResource(v string) *CreateMetricRuleBlackListRequestMetrics {
	s.Resource = &v
	return s
}

func (s *CreateMetricRuleBlackListRequestMetrics) Validate() error {
	return dara.Validate(s)
}

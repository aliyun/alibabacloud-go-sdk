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
	// The category of the Alibaba Cloud service. For example, Redis includes `kvstore_standard` (Standard Edition), `kvstore_sharding` (Cluster Edition), and `kvstore_splitrw` (Read/Write Splitting Edition).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The effective period of the alert blacklist policy.
	//
	// - If you do not specify this parameter, the alert blacklist policy is permanently valid.
	//
	// - If you specify this parameter, the blacklist policy is valid only within the specified time range. This parameter supports the following two formats:
	//
	//     - `03:00-04:59`: The alert blacklist policy is valid from 03:00 to 04:59 (excluding 05:00) local time.
	//
	//     - `03:00-04:59 UTC+0700`: The alert blacklist policy is valid from 03:00 to 04:59 (excluding 05:00) UTC+07:00.
	//
	// example:
	//
	// 03:00-04:59
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The end timestamp when the alert blacklist policy takes effect.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1640608200000
	EnableEndTime *string `json:"EnableEndTime,omitempty" xml:"EnableEndTime,omitempty"`
	// The start timestamp when the alert blacklist policy takes effect.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1640237400000
	EnableStartTime *string `json:"EnableStartTime,omitempty" xml:"EnableStartTime,omitempty"`
	// The list of instances of the specified Alibaba Cloud service in the alert blacklist policy.
	//
	// This parameter is required.
	Instances []*string `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The metrics of the instances.
	//
	// - If you do not specify this parameter, the alert blacklist policy applies to all metrics of the specified Alibaba Cloud service.
	//
	// - If you specify this parameter, the blacklist policy applies only to the specified metrics.
	Metrics []*CreateMetricRuleBlackListRequestMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The name of the alert blacklist policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// Blacklist-01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data namespace of the Alibaba Cloud service.
	//
	// For more information about the namespaces of Alibaba Cloud services, see [Cloud service metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The effective scope of the alert blacklist policy. Valid values:
	//
	// - USER: The alert blacklist policy takes effect only within the current Alibaba Cloud account. This is the default value.
	//
	// - GROUP: The alert blacklist policy takes effect within the specified application group. You must specify an application group. For more information about how to obtain the ID of an application group, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// example:
	//
	// USER
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The list of application group IDs. The value is a JSON array.
	//
	// > This parameter is required only when `ScopeType` is set to `GROUP`.
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
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMetricRuleBlackListRequestMetrics struct {
	// The name of the metric.
	//
	// Valid values of N: 1 to 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// disk_utilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The extended dimension information of the instance. For example, `{"device":"C:"}` indicates that the blacklist policy applies to all C drives of the ECS instance.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMetricRuleBlackListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ModifyMetricRuleBlackListRequest
	GetCategory() *string
	SetEffectiveTime(v string) *ModifyMetricRuleBlackListRequest
	GetEffectiveTime() *string
	SetEnableEndTime(v string) *ModifyMetricRuleBlackListRequest
	GetEnableEndTime() *string
	SetEnableStartTime(v string) *ModifyMetricRuleBlackListRequest
	GetEnableStartTime() *string
	SetId(v string) *ModifyMetricRuleBlackListRequest
	GetId() *string
	SetInstances(v []*string) *ModifyMetricRuleBlackListRequest
	GetInstances() []*string
	SetMetrics(v []*ModifyMetricRuleBlackListRequestMetrics) *ModifyMetricRuleBlackListRequest
	GetMetrics() []*ModifyMetricRuleBlackListRequestMetrics
	SetName(v string) *ModifyMetricRuleBlackListRequest
	GetName() *string
	SetNamespace(v string) *ModifyMetricRuleBlackListRequest
	GetNamespace() *string
	SetRegionId(v string) *ModifyMetricRuleBlackListRequest
	GetRegionId() *string
	SetScopeType(v string) *ModifyMetricRuleBlackListRequest
	GetScopeType() *string
	SetScopeValue(v string) *ModifyMetricRuleBlackListRequest
	GetScopeValue() *string
}

type ModifyMetricRuleBlackListRequest struct {
	// The category of the cloud service. For example, ApsaraDB for Redis supports the standard architecture, the cluster architecture, and the read/write splitting architecture. In this case, the valid values of this parameter for ApsaraDB for Redis include `kvstore_standard`, `kvstore_sharding`, and `kvstore_splitrw`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time range within which the blacklist policy is effective. Take note of the following information:
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
	// The ID of the blacklist policy.
	//
	// For information about how to obtain the ID of a blacklist policy, see [DescribeMetricRuleBlackList](https://help.aliyun.com/document_detail/457257.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 93514c96-ceb8-47d8-8ee3-93b6d98b****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IDs of the instances that belong to the specified cloud service.
	//
	// This parameter is required.
	Instances []*string `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The metrics of the instance.
	//
	// 	- If you do not configure this parameter, the blacklist policy applies to all metrics of the specified cloud service.
	//
	// 	- If you configure this parameter, the blacklist policy applies only to the current metric.
	Metrics []*ModifyMetricRuleBlackListRequestMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The name of the blacklist policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// Blacklist-02
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the cloud service.
	//
	// For more information about the namespaces of cloud services, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
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
	// 	- USER: The blacklist policy takes effect only within the current Alibaba Cloud account.
	//
	// 	- GROUP (default): The blacklist policy takes effect only within the specified application group. For information about how to obtain the ID of an application group, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// example:
	//
	// USER
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The IDs of the application groups.
	//
	// >  This parameter is required only when `ScopeType` is set to `GROUP`.
	//
	// example:
	//
	// ["67****","78****"]
	ScopeValue *string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty"`
}

func (s ModifyMetricRuleBlackListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleBlackListRequest) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleBlackListRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyMetricRuleBlackListRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyMetricRuleBlackListRequest) GetEnableEndTime() *string {
	return s.EnableEndTime
}

func (s *ModifyMetricRuleBlackListRequest) GetEnableStartTime() *string {
	return s.EnableStartTime
}

func (s *ModifyMetricRuleBlackListRequest) GetId() *string {
	return s.Id
}

func (s *ModifyMetricRuleBlackListRequest) GetInstances() []*string {
	return s.Instances
}

func (s *ModifyMetricRuleBlackListRequest) GetMetrics() []*ModifyMetricRuleBlackListRequestMetrics {
	return s.Metrics
}

func (s *ModifyMetricRuleBlackListRequest) GetName() *string {
	return s.Name
}

func (s *ModifyMetricRuleBlackListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyMetricRuleBlackListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMetricRuleBlackListRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *ModifyMetricRuleBlackListRequest) GetScopeValue() *string {
	return s.ScopeValue
}

func (s *ModifyMetricRuleBlackListRequest) SetCategory(v string) *ModifyMetricRuleBlackListRequest {
	s.Category = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) SetEffectiveTime(v string) *ModifyMetricRuleBlackListRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) SetEnableEndTime(v string) *ModifyMetricRuleBlackListRequest {
	s.EnableEndTime = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) SetEnableStartTime(v string) *ModifyMetricRuleBlackListRequest {
	s.EnableStartTime = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) SetId(v string) *ModifyMetricRuleBlackListRequest {
	s.Id = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) SetInstances(v []*string) *ModifyMetricRuleBlackListRequest {
	s.Instances = v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) SetMetrics(v []*ModifyMetricRuleBlackListRequestMetrics) *ModifyMetricRuleBlackListRequest {
	s.Metrics = v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) SetName(v string) *ModifyMetricRuleBlackListRequest {
	s.Name = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) SetNamespace(v string) *ModifyMetricRuleBlackListRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) SetRegionId(v string) *ModifyMetricRuleBlackListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) SetScopeType(v string) *ModifyMetricRuleBlackListRequest {
	s.ScopeType = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) SetScopeValue(v string) *ModifyMetricRuleBlackListRequest {
	s.ScopeValue = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequest) Validate() error {
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

type ModifyMetricRuleBlackListRequestMetrics struct {
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
	// The extended dimension of the instance. For example, `{"device":"C:"}` specifies that the blacklist policy is applied to all C disks of the specified Elastic Compute Service (ECS) instance.
	//
	// Valid values of N: 1 to 10.
	//
	// example:
	//
	// {"device":"C:"}
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s ModifyMetricRuleBlackListRequestMetrics) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleBlackListRequestMetrics) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleBlackListRequestMetrics) GetMetricName() *string {
	return s.MetricName
}

func (s *ModifyMetricRuleBlackListRequestMetrics) GetResource() *string {
	return s.Resource
}

func (s *ModifyMetricRuleBlackListRequestMetrics) SetMetricName(v string) *ModifyMetricRuleBlackListRequestMetrics {
	s.MetricName = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequestMetrics) SetResource(v string) *ModifyMetricRuleBlackListRequestMetrics {
	s.Resource = &v
	return s
}

func (s *ModifyMetricRuleBlackListRequestMetrics) Validate() error {
	return dara.Validate(s)
}

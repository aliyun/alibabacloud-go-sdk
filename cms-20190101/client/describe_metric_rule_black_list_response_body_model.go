// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleBlackListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMetricRuleBlackListResponseBody
	GetCode() *string
	SetDescribeMetricRuleBlackList(v []*DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) *DescribeMetricRuleBlackListResponseBody
	GetDescribeMetricRuleBlackList() []*DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList
	SetMessage(v string) *DescribeMetricRuleBlackListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMetricRuleBlackListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMetricRuleBlackListResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeMetricRuleBlackListResponseBody
	GetTotal() *int32
}

type DescribeMetricRuleBlackListResponseBody struct {
	// The categories of the Alibaba Cloud service. For example, ApsaraDB for Redis includes the following categories: ApsaraDB for Redis (standard architecture), ApsaraDB for Redis (cluster architecture), and ApsaraDB for Redis (read/write splitting architecture). In this case, the valid values of this parameter for ApsaraDB for Redis include `kvstore_standard`, `kvstore_sharding`, and `kvstore_splitrw`.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried blacklist policies.
	DescribeMetricRuleBlackList []*DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList `json:"DescribeMetricRuleBlackList,omitempty" xml:"DescribeMetricRuleBlackList,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The namespace of the cloud service.
	//
	// example:
	//
	// D63E76CB-29AA-5B9F-88CE-400A6F28D428
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The namespace of the cloud service.
	//
	// For more information about the namespaces of different cloud services, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The timestamp when the blacklist policy was created.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMetricRuleBlackListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleBlackListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMetricRuleBlackListResponseBody) GetDescribeMetricRuleBlackList() []*DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	return s.DescribeMetricRuleBlackList
}

func (s *DescribeMetricRuleBlackListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricRuleBlackListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricRuleBlackListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMetricRuleBlackListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeMetricRuleBlackListResponseBody) SetCode(v string) *DescribeMetricRuleBlackListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBody) SetDescribeMetricRuleBlackList(v []*DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) *DescribeMetricRuleBlackListResponseBody {
	s.DescribeMetricRuleBlackList = v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBody) SetMessage(v string) *DescribeMetricRuleBlackListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBody) SetRequestId(v string) *DescribeMetricRuleBlackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBody) SetSuccess(v bool) *DescribeMetricRuleBlackListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBody) SetTotal(v int32) *DescribeMetricRuleBlackListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBody) Validate() error {
	if s.DescribeMetricRuleBlackList != nil {
		for _, item := range s.DescribeMetricRuleBlackList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList struct {
	// The category of the cloud service. For example, ApsaraDB for Redis includes the following categories: ApsaraDB for Redis (standard architecture), ApsaraDB for Redis (cluster architecture), and ApsaraDB for Redis (read/write splitting architecture). In this case, the valid values of this parameter for ApsaraDB for Redis include `kvstore_standard`, `kvstore_sharding`, and `kvstore_splitrw`.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The timestamp when the blacklist policy was created.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1665714561000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time range within which the blacklist policy is effective.
	//
	// example:
	//
	// 00:00-23:59
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The timestamp when the blacklist policy started to take effect.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1640608200000
	EnableEndTime *int64 `json:"EnableEndTime,omitempty" xml:"EnableEndTime,omitempty"`
	// The timestamp when the blacklist policy expired.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1640237400000
	EnableStartTime *int64 `json:"EnableStartTime,omitempty" xml:"EnableStartTime,omitempty"`
	// The ID of the blacklist policy.
	//
	// example:
	//
	// 93514c96-ceb8-47d8-8ee3-93b6d98b****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IDs of the instances that belong to the specified cloud service.
	Instances []*string `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
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
	// The metrics of the instance.
	Metrics []*DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The name of the blacklist policy.
	//
	// example:
	//
	// Blacklist-01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the cloud service.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
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
	// The IDs of the application groups.
	ScopeValue []*string `json:"ScopeValue,omitempty" xml:"ScopeValue,omitempty" type:"Repeated"`
	// The timestamp when the blacklist policy was modified.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1665718373000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetCategory() *string {
	return s.Category
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetEnableEndTime() *int64 {
	return s.EnableEndTime
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetEnableStartTime() *int64 {
	return s.EnableStartTime
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetId() *string {
	return s.Id
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetInstances() []*string {
	return s.Instances
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetIsEnable() *bool {
	return s.IsEnable
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetMetrics() []*DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics {
	return s.Metrics
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetName() *string {
	return s.Name
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetScopeType() *string {
	return s.ScopeType
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetScopeValue() []*string {
	return s.ScopeValue
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetCategory(v string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.Category = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetCreateTime(v string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.CreateTime = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetEffectiveTime(v string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.EffectiveTime = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetEnableEndTime(v int64) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.EnableEndTime = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetEnableStartTime(v int64) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.EnableStartTime = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetId(v string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.Id = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetInstances(v []*string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.Instances = v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetIsEnable(v bool) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.IsEnable = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetMetrics(v []*DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.Metrics = v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetName(v string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.Name = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetNamespace(v string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetScopeType(v string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.ScopeType = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetScopeValue(v []*string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.ScopeValue = v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) SetUpdateTime(v string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList {
	s.UpdateTime = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackList) Validate() error {
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

type DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics struct {
	// The metric name.
	//
	// example:
	//
	// disk_utilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The extended dimension of the instance. For example, `{"device":"C:"}` specifies that the blacklist policy is applied to all C disks of the specified Elastic Compute Service (ECS) instance.
	//
	// example:
	//
	// [{"device":"C:"}]
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics) GetResource() *string {
	return s.Resource
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics) SetMetricName(v string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics) SetResource(v string) *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics {
	s.Resource = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponseBodyDescribeMetricRuleBlackListMetrics) Validate() error {
	return dara.Validate(s)
}

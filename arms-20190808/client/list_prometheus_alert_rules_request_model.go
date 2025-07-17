// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListPrometheusAlertRulesRequest
	GetClusterId() *string
	SetMatchExpressions(v string) *ListPrometheusAlertRulesRequest
	GetMatchExpressions() *string
	SetName(v string) *ListPrometheusAlertRulesRequest
	GetName() *string
	SetRegionId(v string) *ListPrometheusAlertRulesRequest
	GetRegionId() *string
	SetStatus(v int32) *ListPrometheusAlertRulesRequest
	GetStatus() *int32
	SetTags(v []*ListPrometheusAlertRulesRequestTags) *ListPrometheusAlertRulesRequest
	GetTags() []*ListPrometheusAlertRulesRequestTags
	SetType(v string) *ListPrometheusAlertRulesRequest
	GetType() *string
}

type ListPrometheusAlertRulesRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// c0bad479465464e1d8c1e641b0afb****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The tag match conditions that are described in a JSON string. For more information about this parameter, see the **Additional description of the MatchExpressions parameter*	- section.
	//
	// example:
	//
	// [{"key":"severity","value":"critical","operator":"re"}]
	MatchExpressions *string `json:"MatchExpressions,omitempty" xml:"MatchExpressions,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// Prometheus_Alert
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether the alert rule is enabled. Valid values:
	//
	// - 1: enables the alert rule.
	//
	// - 0: disables the alert rule.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListPrometheusAlertRulesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the alert rule.
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrometheusAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPrometheusAlertRulesRequest) GetMatchExpressions() *string {
	return s.MatchExpressions
}

func (s *ListPrometheusAlertRulesRequest) GetName() *string {
	return s.Name
}

func (s *ListPrometheusAlertRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusAlertRulesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListPrometheusAlertRulesRequest) GetTags() []*ListPrometheusAlertRulesRequestTags {
	return s.Tags
}

func (s *ListPrometheusAlertRulesRequest) GetType() *string {
	return s.Type
}

func (s *ListPrometheusAlertRulesRequest) SetClusterId(v string) *ListPrometheusAlertRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetMatchExpressions(v string) *ListPrometheusAlertRulesRequest {
	s.MatchExpressions = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetName(v string) *ListPrometheusAlertRulesRequest {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetRegionId(v string) *ListPrometheusAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetStatus(v int32) *ListPrometheusAlertRulesRequest {
	s.Status = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetTags(v []*ListPrometheusAlertRulesRequestTags) *ListPrometheusAlertRulesRequest {
	s.Tags = v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetType(v string) *ListPrometheusAlertRulesRequest {
	s.Type = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusAlertRulesRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// zhangsan
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertRulesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesRequestTags) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListPrometheusAlertRulesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusAlertRulesRequestTags) SetKey(v string) *ListPrometheusAlertRulesRequestTags {
	s.Key = &v
	return s
}

func (s *ListPrometheusAlertRulesRequestTags) SetValue(v string) *ListPrometheusAlertRulesRequestTags {
	s.Value = &v
	return s
}

func (s *ListPrometheusAlertRulesRequestTags) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugResourceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DebugResourceRuleRequest
	GetInstanceId() *string
	SetMetricInfo(v map[string]interface{}) *DebugResourceRuleRequest
	GetMetricInfo() map[string]interface{}
	SetRegionId(v string) *DebugResourceRuleRequest
	GetRegionId() *string
}

type DebugResourceRuleRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MetricInfo map[string]interface{} `json:"MetricInfo,omitempty" xml:"MetricInfo,omitempty"`
	RegionId   *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DebugResourceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DebugResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *DebugResourceRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DebugResourceRuleRequest) GetMetricInfo() map[string]interface{} {
	return s.MetricInfo
}

func (s *DebugResourceRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DebugResourceRuleRequest) SetInstanceId(v string) *DebugResourceRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DebugResourceRuleRequest) SetMetricInfo(v map[string]interface{}) *DebugResourceRuleRequest {
	s.MetricInfo = v
	return s
}

func (s *DebugResourceRuleRequest) SetRegionId(v string) *DebugResourceRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DebugResourceRuleRequest) Validate() error {
	return dara.Validate(s)
}

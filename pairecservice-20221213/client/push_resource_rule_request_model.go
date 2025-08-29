// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResourceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PushResourceRuleRequest
	GetInstanceId() *string
	SetMetricInfo(v map[string]interface{}) *PushResourceRuleRequest
	GetMetricInfo() map[string]interface{}
}

type PushResourceRuleRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MetricInfo map[string]interface{} `json:"MetricInfo,omitempty" xml:"MetricInfo,omitempty"`
}

func (s PushResourceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s PushResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *PushResourceRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PushResourceRuleRequest) GetMetricInfo() map[string]interface{} {
	return s.MetricInfo
}

func (s *PushResourceRuleRequest) SetInstanceId(v string) *PushResourceRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *PushResourceRuleRequest) SetMetricInfo(v map[string]interface{}) *PushResourceRuleRequest {
	s.MetricInfo = v
	return s
}

func (s *PushResourceRuleRequest) Validate() error {
	return dara.Validate(s)
}

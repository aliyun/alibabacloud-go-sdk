// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugResourceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentValues(v map[string]interface{}) *DebugResourceRuleResponseBody
	GetCurrentValues() map[string]interface{}
	SetOutputValues(v map[string]interface{}) *DebugResourceRuleResponseBody
	GetOutputValues() map[string]interface{}
	SetRequestId(v string) *DebugResourceRuleResponseBody
	GetRequestId() *string
}

type DebugResourceRuleResponseBody struct {
	CurrentValues map[string]interface{} `json:"CurrentValues,omitempty" xml:"CurrentValues,omitempty"`
	OutputValues  map[string]interface{} `json:"OutputValues,omitempty" xml:"OutputValues,omitempty"`
	RequestId     *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DebugResourceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DebugResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DebugResourceRuleResponseBody) GetCurrentValues() map[string]interface{} {
	return s.CurrentValues
}

func (s *DebugResourceRuleResponseBody) GetOutputValues() map[string]interface{} {
	return s.OutputValues
}

func (s *DebugResourceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DebugResourceRuleResponseBody) SetCurrentValues(v map[string]interface{}) *DebugResourceRuleResponseBody {
	s.CurrentValues = v
	return s
}

func (s *DebugResourceRuleResponseBody) SetOutputValues(v map[string]interface{}) *DebugResourceRuleResponseBody {
	s.OutputValues = v
	return s
}

func (s *DebugResourceRuleResponseBody) SetRequestId(v string) *DebugResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugResourceRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

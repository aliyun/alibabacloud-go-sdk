// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvDropMetricsRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateEnvDropMetricsRuleResponseBody
	GetCode() *int32
	SetData(v string) *UpdateEnvDropMetricsRuleResponseBody
	GetData() *string
	SetMessage(v string) *UpdateEnvDropMetricsRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEnvDropMetricsRuleResponseBody
	GetRequestId() *string
}

type UpdateEnvDropMetricsRuleResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6A9AEA84-7186-4D8D-B498-4585C6A2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEnvDropMetricsRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvDropMetricsRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvDropMetricsRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateEnvDropMetricsRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateEnvDropMetricsRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEnvDropMetricsRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnvDropMetricsRuleResponseBody) SetCode(v int32) *UpdateEnvDropMetricsRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnvDropMetricsRuleResponseBody) SetData(v string) *UpdateEnvDropMetricsRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateEnvDropMetricsRuleResponseBody) SetMessage(v string) *UpdateEnvDropMetricsRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEnvDropMetricsRuleResponseBody) SetRequestId(v string) *UpdateEnvDropMetricsRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnvDropMetricsRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

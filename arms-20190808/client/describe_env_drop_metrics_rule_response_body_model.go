// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvDropMetricsRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEnvDropMetricsRuleResponseBody
	GetCode() *int32
	SetData(v *DescribeEnvDropMetricsRuleResponseBodyData) *DescribeEnvDropMetricsRuleResponseBody
	GetData() *DescribeEnvDropMetricsRuleResponseBodyData
	SetMessage(v string) *DescribeEnvDropMetricsRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEnvDropMetricsRuleResponseBody
	GetRequestId() *string
}

type DescribeEnvDropMetricsRuleResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *DescribeEnvDropMetricsRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// F7781D4A-2818-41E7-B7BB-79D809E9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnvDropMetricsRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvDropMetricsRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnvDropMetricsRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEnvDropMetricsRuleResponseBody) GetData() *DescribeEnvDropMetricsRuleResponseBodyData {
	return s.Data
}

func (s *DescribeEnvDropMetricsRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEnvDropMetricsRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnvDropMetricsRuleResponseBody) SetCode(v int32) *DescribeEnvDropMetricsRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnvDropMetricsRuleResponseBody) SetData(v *DescribeEnvDropMetricsRuleResponseBodyData) *DescribeEnvDropMetricsRuleResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnvDropMetricsRuleResponseBody) SetMessage(v string) *DescribeEnvDropMetricsRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEnvDropMetricsRuleResponseBody) SetRequestId(v string) *DescribeEnvDropMetricsRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnvDropMetricsRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnvDropMetricsRuleResponseBodyData struct {
	// The list of discarded metrics. Separate multiple metrics with line feeds.
	//
	// example:
	//
	// kube_pod_created
	DropMetrics *string `json:"DropMetrics,omitempty" xml:"DropMetrics,omitempty"`
	// The name of the discarded metric rule.
	//
	// example:
	//
	// ruleName1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeEnvDropMetricsRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvDropMetricsRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnvDropMetricsRuleResponseBodyData) GetDropMetrics() *string {
	return s.DropMetrics
}

func (s *DescribeEnvDropMetricsRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeEnvDropMetricsRuleResponseBodyData) SetDropMetrics(v string) *DescribeEnvDropMetricsRuleResponseBodyData {
	s.DropMetrics = &v
	return s
}

func (s *DescribeEnvDropMetricsRuleResponseBodyData) SetRuleName(v string) *DescribeEnvDropMetricsRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *DescribeEnvDropMetricsRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}

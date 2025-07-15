// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectTrafficQosRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQosId(v string) *CreateExpressConnectTrafficQosRuleResponseBody
	GetQosId() *string
	SetQueueId(v string) *CreateExpressConnectTrafficQosRuleResponseBody
	GetQueueId() *string
	SetRequestId(v string) *CreateExpressConnectTrafficQosRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *CreateExpressConnectTrafficQosRuleResponseBody
	GetRuleId() *string
}

type CreateExpressConnectTrafficQosRuleResponseBody struct {
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-2giu0a6vd5x0mv4700
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the QoS queue.
	//
	// example:
	//
	// qos-queue-9nyx2u7n71s2rcy4n5
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 606998F0-B94D-48FE-8316-ACA81BB230DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the QoS rule.
	//
	// example:
	//
	// qos-rule-iugg0l9x27f2nocouj
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateExpressConnectTrafficQosRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectTrafficQosRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectTrafficQosRuleResponseBody) GetQosId() *string {
	return s.QosId
}

func (s *CreateExpressConnectTrafficQosRuleResponseBody) GetQueueId() *string {
	return s.QueueId
}

func (s *CreateExpressConnectTrafficQosRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExpressConnectTrafficQosRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateExpressConnectTrafficQosRuleResponseBody) SetQosId(v string) *CreateExpressConnectTrafficQosRuleResponseBody {
	s.QosId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleResponseBody) SetQueueId(v string) *CreateExpressConnectTrafficQosRuleResponseBody {
	s.QueueId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleResponseBody) SetRequestId(v string) *CreateExpressConnectTrafficQosRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleResponseBody) SetRuleId(v string) *CreateExpressConnectTrafficQosRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

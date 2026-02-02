// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQosRuleId(v string) *CreateQosRuleResponseBody
	GetQosRuleId() *string
	SetRequestId(v string) *CreateQosRuleResponseBody
	GetRequestId() *string
}

type CreateQosRuleResponseBody struct {
	// example:
	//
	// qos-5605u0gelk200****
	QosRuleId *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	// example:
	//
	// 51592A88-0F2C-55E6-AD2C-2AD9C10D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQosRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQosRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQosRuleResponseBody) GetQosRuleId() *string {
	return s.QosRuleId
}

func (s *CreateQosRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQosRuleResponseBody) SetQosRuleId(v string) *CreateQosRuleResponseBody {
	s.QosRuleId = &v
	return s
}

func (s *CreateQosRuleResponseBody) SetRequestId(v string) *CreateQosRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQosRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyQosRuleResponseBody
	GetRequestId() *string
}

type ModifyQosRuleResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyQosRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyQosRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyQosRuleResponseBody) SetRequestId(v string) *ModifyQosRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyQosRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

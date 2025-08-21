// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSchedulerRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchSchedulerRuleResponseBody
	GetRequestId() *string
}

type SwitchSchedulerRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7E3C301F-84BB-50E4-9DB9-2937B2429C1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchSchedulerRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchSchedulerRuleResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchSchedulerRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchSchedulerRuleResponseBody) SetRequestId(v string) *SwitchSchedulerRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchSchedulerRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

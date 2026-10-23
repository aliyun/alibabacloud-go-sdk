// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostEmonTryAlarmRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *PostEmonTryAlarmRuleRequest
	GetBody() *string
}

type PostEmonTryAlarmRuleRequest struct {
	// The information about the test alert rule, notification method, and message recipients.
	//
	// example:
	//
	// {"alarmRuleName":"test_rule","AlarmChannel":{"phone":true,"sms":true,"dingWebHook":true},"receivers":[{"id":19,"contactGroup":false}]}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostEmonTryAlarmRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s PostEmonTryAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *PostEmonTryAlarmRuleRequest) GetBody() *string {
	return s.Body
}

func (s *PostEmonTryAlarmRuleRequest) SetBody(v string) *PostEmonTryAlarmRuleRequest {
	s.Body = &v
	return s
}

func (s *PostEmonTryAlarmRuleRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostEmonTryAlarmRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PostEmonTryAlarmRuleResponseBody
	GetCode() *string
	SetMessage(v string) *PostEmonTryAlarmRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *PostEmonTryAlarmRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PostEmonTryAlarmRuleResponseBody
	GetSuccess() *bool
}

type PostEmonTryAlarmRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3EC5731F-0944-4E4F-9DD5-1F976B3FCC3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostEmonTryAlarmRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostEmonTryAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PostEmonTryAlarmRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *PostEmonTryAlarmRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PostEmonTryAlarmRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PostEmonTryAlarmRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PostEmonTryAlarmRuleResponseBody) SetCode(v string) *PostEmonTryAlarmRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PostEmonTryAlarmRuleResponseBody) SetMessage(v string) *PostEmonTryAlarmRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PostEmonTryAlarmRuleResponseBody) SetRequestId(v string) *PostEmonTryAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostEmonTryAlarmRuleResponseBody) SetSuccess(v bool) *PostEmonTryAlarmRuleResponseBody {
	s.Success = &v
	return s
}

func (s *PostEmonTryAlarmRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

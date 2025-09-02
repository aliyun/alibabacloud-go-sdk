// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDIAlarmRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDIAlarmRuleId(v int64) *CreateDIAlarmRuleResponseBody
	GetDIAlarmRuleId() *int64
	SetRequestId(v string) *CreateDIAlarmRuleResponseBody
	GetRequestId() *string
}

type CreateDIAlarmRuleResponseBody struct {
	// The alert rule ID.
	//
	// example:
	//
	// 34988
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C636A747-7E4E-594D-94CD-2B4F8A9A9A63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDIAlarmRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDIAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleResponseBody) GetDIAlarmRuleId() *int64 {
	return s.DIAlarmRuleId
}

func (s *CreateDIAlarmRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDIAlarmRuleResponseBody) SetDIAlarmRuleId(v int64) *CreateDIAlarmRuleResponseBody {
	s.DIAlarmRuleId = &v
	return s
}

func (s *CreateDIAlarmRuleResponseBody) SetRequestId(v string) *CreateDIAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDIAlarmRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

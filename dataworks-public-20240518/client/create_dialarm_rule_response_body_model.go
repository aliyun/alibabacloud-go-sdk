// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDIAlarmRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDIAlarmRuleId(v string) *CreateDIAlarmRuleResponseBody
	GetDIAlarmRuleId() *string
	SetId(v int64) *CreateDIAlarmRuleResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateDIAlarmRuleResponseBody
	GetRequestId() *string
}

type CreateDIAlarmRuleResponseBody struct {
	// Deprecated
	//
	// **[Deprecated]*	- Use the Id parameter instead.
	//
	// example:
	//
	// 1
	DIAlarmRuleId *string `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// C636A747-7E4E-594D-94CD-2B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDIAlarmRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDIAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleResponseBody) GetDIAlarmRuleId() *string {
	return s.DIAlarmRuleId
}

func (s *CreateDIAlarmRuleResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateDIAlarmRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDIAlarmRuleResponseBody) SetDIAlarmRuleId(v string) *CreateDIAlarmRuleResponseBody {
	s.DIAlarmRuleId = &v
	return s
}

func (s *CreateDIAlarmRuleResponseBody) SetId(v int64) *CreateDIAlarmRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDIAlarmRuleResponseBody) SetRequestId(v string) *CreateDIAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDIAlarmRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

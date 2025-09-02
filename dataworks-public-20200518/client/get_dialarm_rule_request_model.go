// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDIAlarmRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIAlarmRuleId(v int64) *GetDIAlarmRuleRequest
	GetDIAlarmRuleId() *int64
}

type GetDIAlarmRuleRequest struct {
	// The alert rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34994
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
}

func (s GetDIAlarmRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDIAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *GetDIAlarmRuleRequest) GetDIAlarmRuleId() *int64 {
	return s.DIAlarmRuleId
}

func (s *GetDIAlarmRuleRequest) SetDIAlarmRuleId(v int64) *GetDIAlarmRuleRequest {
	s.DIAlarmRuleId = &v
	return s
}

func (s *GetDIAlarmRuleRequest) Validate() error {
	return dara.Validate(s)
}

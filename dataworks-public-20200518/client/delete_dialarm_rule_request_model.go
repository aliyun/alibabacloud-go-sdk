// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDIAlarmRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIAlarmRuleId(v int64) *DeleteDIAlarmRuleRequest
	GetDIAlarmRuleId() *int64
}

type DeleteDIAlarmRuleRequest struct {
	// The alert rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34971
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
}

func (s DeleteDIAlarmRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDIAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDIAlarmRuleRequest) GetDIAlarmRuleId() *int64 {
	return s.DIAlarmRuleId
}

func (s *DeleteDIAlarmRuleRequest) SetDIAlarmRuleId(v int64) *DeleteDIAlarmRuleRequest {
	s.DIAlarmRuleId = &v
	return s
}

func (s *DeleteDIAlarmRuleRequest) Validate() error {
	return dara.Validate(s)
}

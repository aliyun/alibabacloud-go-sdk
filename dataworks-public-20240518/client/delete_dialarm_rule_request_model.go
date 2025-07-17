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
	SetDIJobId(v int64) *DeleteDIAlarmRuleRequest
	GetDIJobId() *int64
	SetId(v int64) *DeleteDIAlarmRuleRequest
	GetId() *int64
}

type DeleteDIAlarmRuleRequest struct {
	// Deprecated
	//
	// This parameter is deprecated. Use the Id parameter instead.
	//
	// example:
	//
	// 2
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 1
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *DeleteDIAlarmRuleRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *DeleteDIAlarmRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDIAlarmRuleRequest) SetDIAlarmRuleId(v int64) *DeleteDIAlarmRuleRequest {
	s.DIAlarmRuleId = &v
	return s
}

func (s *DeleteDIAlarmRuleRequest) SetDIJobId(v int64) *DeleteDIAlarmRuleRequest {
	s.DIJobId = &v
	return s
}

func (s *DeleteDIAlarmRuleRequest) SetId(v int64) *DeleteDIAlarmRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteDIAlarmRuleRequest) Validate() error {
	return dara.Validate(s)
}

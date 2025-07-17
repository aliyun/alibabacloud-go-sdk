// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAlarmSeverityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmId(v int64) *ChangeAlarmSeverityRequest
	GetAlarmId() *int64
	SetHandlerId(v int64) *ChangeAlarmSeverityRequest
	GetHandlerId() *int64
	SetRegionId(v string) *ChangeAlarmSeverityRequest
	GetRegionId() *string
	SetSeverity(v string) *ChangeAlarmSeverityRequest
	GetSeverity() *string
}

type ChangeAlarmSeverityRequest struct {
	// The ID of the alert.
	//
	// For more information about how to obtain the ID of an alert, see [ListAlertEvents](https://help.aliyun.com/document_detail/2612346.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 155
	AlarmId *int64 `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	// The ID of the handler.
	//
	// example:
	//
	// 2046076
	HandlerId *int64 `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The severity level of the alert. Valid values: P1, P2, P3, and P4. P4 indicates the lowest severity, whereas P1 indicates the highest severity.
	//
	// This parameter is required.
	//
	// example:
	//
	// P1
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s ChangeAlarmSeverityRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeAlarmSeverityRequest) GoString() string {
	return s.String()
}

func (s *ChangeAlarmSeverityRequest) GetAlarmId() *int64 {
	return s.AlarmId
}

func (s *ChangeAlarmSeverityRequest) GetHandlerId() *int64 {
	return s.HandlerId
}

func (s *ChangeAlarmSeverityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChangeAlarmSeverityRequest) GetSeverity() *string {
	return s.Severity
}

func (s *ChangeAlarmSeverityRequest) SetAlarmId(v int64) *ChangeAlarmSeverityRequest {
	s.AlarmId = &v
	return s
}

func (s *ChangeAlarmSeverityRequest) SetHandlerId(v int64) *ChangeAlarmSeverityRequest {
	s.HandlerId = &v
	return s
}

func (s *ChangeAlarmSeverityRequest) SetRegionId(v string) *ChangeAlarmSeverityRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeAlarmSeverityRequest) SetSeverity(v string) *ChangeAlarmSeverityRequest {
	s.Severity = &v
	return s
}

func (s *ChangeAlarmSeverityRequest) Validate() error {
	return dara.Validate(s)
}

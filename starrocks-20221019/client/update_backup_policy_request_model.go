// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireDays(v int32) *UpdateBackupPolicyRequest
	GetExpireDays() *int32
	SetHour(v int32) *UpdateBackupPolicyRequest
	GetHour() *int32
	SetInstanceId(v string) *UpdateBackupPolicyRequest
	GetInstanceId() *string
	SetMinute(v int32) *UpdateBackupPolicyRequest
	GetMinute() *int32
	SetPolicyId(v string) *UpdateBackupPolicyRequest
	GetPolicyId() *string
	SetRecurrenceValues(v []*int32) *UpdateBackupPolicyRequest
	GetRecurrenceValues() []*int32
	SetTimeoutSeconds(v int32) *UpdateBackupPolicyRequest
	GetTimeoutSeconds() *int32
}

type UpdateBackupPolicyRequest struct {
	// example:
	//
	// 7
	ExpireDays *int32 `json:"ExpireDays,omitempty" xml:"ExpireDays,omitempty"`
	// example:
	//
	// 10
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// c-0104730e9de40215
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 15
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// example:
	//
	// bk-9812023
	PolicyId         *string  `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RecurrenceValues []*int32 `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// example:
	//
	// 3600
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s UpdateBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyRequest) GetExpireDays() *int32 {
	return s.ExpireDays
}

func (s *UpdateBackupPolicyRequest) GetHour() *int32 {
	return s.Hour
}

func (s *UpdateBackupPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateBackupPolicyRequest) GetMinute() *int32 {
	return s.Minute
}

func (s *UpdateBackupPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateBackupPolicyRequest) GetRecurrenceValues() []*int32 {
	return s.RecurrenceValues
}

func (s *UpdateBackupPolicyRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *UpdateBackupPolicyRequest) SetExpireDays(v int32) *UpdateBackupPolicyRequest {
	s.ExpireDays = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetHour(v int32) *UpdateBackupPolicyRequest {
	s.Hour = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetInstanceId(v string) *UpdateBackupPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetMinute(v int32) *UpdateBackupPolicyRequest {
	s.Minute = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetPolicyId(v string) *UpdateBackupPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetRecurrenceValues(v []*int32) *UpdateBackupPolicyRequest {
	s.RecurrenceValues = v
	return s
}

func (s *UpdateBackupPolicyRequest) SetTimeoutSeconds(v int32) *UpdateBackupPolicyRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *UpdateBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}

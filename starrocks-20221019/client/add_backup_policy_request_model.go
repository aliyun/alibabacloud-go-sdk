// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireDays(v int32) *AddBackupPolicyRequest
	GetExpireDays() *int32
	SetHour(v int32) *AddBackupPolicyRequest
	GetHour() *int32
	SetInstanceId(v string) *AddBackupPolicyRequest
	GetInstanceId() *string
	SetMinute(v int32) *AddBackupPolicyRequest
	GetMinute() *int32
	SetRecurrenceType(v string) *AddBackupPolicyRequest
	GetRecurrenceType() *string
	SetRecurrenceValues(v []*int32) *AddBackupPolicyRequest
	GetRecurrenceValues() []*int32
	SetTimeoutSeconds(v int32) *AddBackupPolicyRequest
	GetTimeoutSeconds() *int32
}

type AddBackupPolicyRequest struct {
	// example:
	//
	// 7
	ExpireDays *int32 `json:"ExpireDays,omitempty" xml:"ExpireDays,omitempty"`
	// example:
	//
	// 2
	Hour *int32 `json:"Hour,omitempty" xml:"Hour,omitempty"`
	// example:
	//
	// c-0104730e9de40215
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 30
	Minute *int32 `json:"Minute,omitempty" xml:"Minute,omitempty"`
	// example:
	//
	// MONTHLY
	RecurrenceType   *string  `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValues []*int32 `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// example:
	//
	// 3600
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s AddBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *AddBackupPolicyRequest) GetExpireDays() *int32 {
	return s.ExpireDays
}

func (s *AddBackupPolicyRequest) GetHour() *int32 {
	return s.Hour
}

func (s *AddBackupPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddBackupPolicyRequest) GetMinute() *int32 {
	return s.Minute
}

func (s *AddBackupPolicyRequest) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *AddBackupPolicyRequest) GetRecurrenceValues() []*int32 {
	return s.RecurrenceValues
}

func (s *AddBackupPolicyRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *AddBackupPolicyRequest) SetExpireDays(v int32) *AddBackupPolicyRequest {
	s.ExpireDays = &v
	return s
}

func (s *AddBackupPolicyRequest) SetHour(v int32) *AddBackupPolicyRequest {
	s.Hour = &v
	return s
}

func (s *AddBackupPolicyRequest) SetInstanceId(v string) *AddBackupPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *AddBackupPolicyRequest) SetMinute(v int32) *AddBackupPolicyRequest {
	s.Minute = &v
	return s
}

func (s *AddBackupPolicyRequest) SetRecurrenceType(v string) *AddBackupPolicyRequest {
	s.RecurrenceType = &v
	return s
}

func (s *AddBackupPolicyRequest) SetRecurrenceValues(v []*int32) *AddBackupPolicyRequest {
	s.RecurrenceValues = v
	return s
}

func (s *AddBackupPolicyRequest) SetTimeoutSeconds(v int32) *AddBackupPolicyRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *AddBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}

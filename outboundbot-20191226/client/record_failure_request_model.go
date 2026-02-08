// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordFailureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActualTime(v int64) *RecordFailureRequest
	GetActualTime() *int64
	SetCallId(v string) *RecordFailureRequest
	GetCallId() *string
	SetCalledNumber(v string) *RecordFailureRequest
	GetCalledNumber() *string
	SetCallingNumber(v string) *RecordFailureRequest
	GetCallingNumber() *string
	SetDispositionCode(v string) *RecordFailureRequest
	GetDispositionCode() *string
	SetExceptionCodes(v string) *RecordFailureRequest
	GetExceptionCodes() *string
	SetInstanceId(v string) *RecordFailureRequest
	GetInstanceId() *string
	SetTaskId(v string) *RecordFailureRequest
	GetTaskId() *string
}

type RecordFailureRequest struct {
	// Call Start Time
	//
	// This parameter is required.
	//
	// example:
	//
	// 1579055782000
	ActualTime *int64 `json:"ActualTime,omitempty" xml:"ActualTime,omitempty"`
	// Call ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 4f21446e-324e-46f2-bf62-7f341fb004ea
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Called number
	//
	// This parameter is required.
	//
	// example:
	//
	// 135815****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// Calling number
	//
	// This parameter is required.
	//
	// example:
	//
	// 10086
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// Reason for failure:
	//
	// - nonexistent number (NotExists)
	//
	// - Busy (Busy)
	//
	// - Not answered (NotAnswered)
	//
	// This parameter is required.
	//
	// example:
	//
	// NotExists
	DispositionCode *string `json:"DispositionCode,omitempty" xml:"DispositionCode,omitempty"`
	// Error encoding when the outbound call fails:
	//
	// - nonexistent number (NotExists)
	//
	// - Busy (Busy)
	//
	// - Not answered (NotAnswered)
	//
	// example:
	//
	// NotExists
	ExceptionCodes *string `json:"ExceptionCodes,omitempty" xml:"ExceptionCodes,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 9ab43460-c0b9-40e2-8447-48d82c97fc67
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job ID
	//
	// This parameter is required.
	//
	// example:
	//
	// d2295c0e-3bc3-48a5-9f56-b185db2be909
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecordFailureRequest) String() string {
	return dara.Prettify(s)
}

func (s RecordFailureRequest) GoString() string {
	return s.String()
}

func (s *RecordFailureRequest) GetActualTime() *int64 {
	return s.ActualTime
}

func (s *RecordFailureRequest) GetCallId() *string {
	return s.CallId
}

func (s *RecordFailureRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *RecordFailureRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *RecordFailureRequest) GetDispositionCode() *string {
	return s.DispositionCode
}

func (s *RecordFailureRequest) GetExceptionCodes() *string {
	return s.ExceptionCodes
}

func (s *RecordFailureRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RecordFailureRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RecordFailureRequest) SetActualTime(v int64) *RecordFailureRequest {
	s.ActualTime = &v
	return s
}

func (s *RecordFailureRequest) SetCallId(v string) *RecordFailureRequest {
	s.CallId = &v
	return s
}

func (s *RecordFailureRequest) SetCalledNumber(v string) *RecordFailureRequest {
	s.CalledNumber = &v
	return s
}

func (s *RecordFailureRequest) SetCallingNumber(v string) *RecordFailureRequest {
	s.CallingNumber = &v
	return s
}

func (s *RecordFailureRequest) SetDispositionCode(v string) *RecordFailureRequest {
	s.DispositionCode = &v
	return s
}

func (s *RecordFailureRequest) SetExceptionCodes(v string) *RecordFailureRequest {
	s.ExceptionCodes = &v
	return s
}

func (s *RecordFailureRequest) SetInstanceId(v string) *RecordFailureRequest {
	s.InstanceId = &v
	return s
}

func (s *RecordFailureRequest) SetTaskId(v string) *RecordFailureRequest {
	s.TaskId = &v
	return s
}

func (s *RecordFailureRequest) Validate() error {
	return dara.Validate(s)
}

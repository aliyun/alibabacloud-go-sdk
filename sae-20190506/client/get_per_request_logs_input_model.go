// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPerRequestLogsInput interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetPerRequestLogsInput
	GetEndTime() *int64
	SetErrorType(v string) *GetPerRequestLogsInput
	GetErrorType() *string
	SetForwardLine(v int64) *GetPerRequestLogsInput
	GetForwardLine() *int64
	SetInstanceID(v string) *GetPerRequestLogsInput
	GetInstanceID() *string
	SetIsColdStart(v bool) *GetPerRequestLogsInput
	GetIsColdStart() *bool
	SetRequestID(v string) *GetPerRequestLogsInput
	GetRequestID() *string
	SetStartTime(v int64) *GetPerRequestLogsInput
	GetStartTime() *int64
	SetTimestamp(v string) *GetPerRequestLogsInput
	GetTimestamp() *string
}

type GetPerRequestLogsInput struct {
	// This parameter is required.
	EndTime     *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ErrorType   *string `json:"errorType,omitempty" xml:"errorType,omitempty"`
	ForwardLine *int64  `json:"forwardLine,omitempty" xml:"forwardLine,omitempty"`
	InstanceID  *string `json:"instanceID,omitempty" xml:"instanceID,omitempty"`
	IsColdStart *bool   `json:"isColdStart,omitempty" xml:"isColdStart,omitempty"`
	// This parameter is required.
	RequestID *string `json:"requestID,omitempty" xml:"requestID,omitempty"`
	// This parameter is required.
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetPerRequestLogsInput) String() string {
	return dara.Prettify(s)
}

func (s GetPerRequestLogsInput) GoString() string {
	return s.String()
}

func (s *GetPerRequestLogsInput) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetPerRequestLogsInput) GetErrorType() *string {
	return s.ErrorType
}

func (s *GetPerRequestLogsInput) GetForwardLine() *int64 {
	return s.ForwardLine
}

func (s *GetPerRequestLogsInput) GetInstanceID() *string {
	return s.InstanceID
}

func (s *GetPerRequestLogsInput) GetIsColdStart() *bool {
	return s.IsColdStart
}

func (s *GetPerRequestLogsInput) GetRequestID() *string {
	return s.RequestID
}

func (s *GetPerRequestLogsInput) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetPerRequestLogsInput) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetPerRequestLogsInput) SetEndTime(v int64) *GetPerRequestLogsInput {
	s.EndTime = &v
	return s
}

func (s *GetPerRequestLogsInput) SetErrorType(v string) *GetPerRequestLogsInput {
	s.ErrorType = &v
	return s
}

func (s *GetPerRequestLogsInput) SetForwardLine(v int64) *GetPerRequestLogsInput {
	s.ForwardLine = &v
	return s
}

func (s *GetPerRequestLogsInput) SetInstanceID(v string) *GetPerRequestLogsInput {
	s.InstanceID = &v
	return s
}

func (s *GetPerRequestLogsInput) SetIsColdStart(v bool) *GetPerRequestLogsInput {
	s.IsColdStart = &v
	return s
}

func (s *GetPerRequestLogsInput) SetRequestID(v string) *GetPerRequestLogsInput {
	s.RequestID = &v
	return s
}

func (s *GetPerRequestLogsInput) SetStartTime(v int64) *GetPerRequestLogsInput {
	s.StartTime = &v
	return s
}

func (s *GetPerRequestLogsInput) SetTimestamp(v string) *GetPerRequestLogsInput {
	s.Timestamp = &v
	return s
}

func (s *GetPerRequestLogsInput) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationLogsInput interface {
	dara.Model
	String() string
	GoString() string
	SetBackwardLine(v int64) *GetApplicationLogsInput
	GetBackwardLine() *int64
	SetEndTime(v int64) *GetApplicationLogsInput
	GetEndTime() *int64
	SetForwardLine(v int64) *GetApplicationLogsInput
	GetForwardLine() *int64
	SetInstanceID(v string) *GetApplicationLogsInput
	GetInstanceID() *string
	SetIsTail(v bool) *GetApplicationLogsInput
	GetIsTail() *bool
	SetMatch(v string) *GetApplicationLogsInput
	GetMatch() *string
	SetMessage(v string) *GetApplicationLogsInput
	GetMessage() *string
	SetOffset(v int64) *GetApplicationLogsInput
	GetOffset() *int64
	SetPackID(v string) *GetApplicationLogsInput
	GetPackID() *string
	SetPackMeta(v string) *GetApplicationLogsInput
	GetPackMeta() *string
	SetStartTime(v int64) *GetApplicationLogsInput
	GetStartTime() *int64
	SetTimestamp(v string) *GetApplicationLogsInput
	GetTimestamp() *string
	SetVersionID(v string) *GetApplicationLogsInput
	GetVersionID() *string
}

type GetApplicationLogsInput struct {
	BackwardLine *int64 `json:"backwardLine,omitempty" xml:"backwardLine,omitempty"`
	// This parameter is required.
	EndTime     *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ForwardLine *int64  `json:"forwardLine,omitempty" xml:"forwardLine,omitempty"`
	InstanceID  *string `json:"instanceID,omitempty" xml:"instanceID,omitempty"`
	IsTail      *bool   `json:"isTail,omitempty" xml:"isTail,omitempty"`
	Match       *string `json:"match,omitempty" xml:"match,omitempty"`
	Message     *string `json:"message,omitempty" xml:"message,omitempty"`
	Offset      *int64  `json:"offset,omitempty" xml:"offset,omitempty"`
	PackID      *string `json:"packID,omitempty" xml:"packID,omitempty"`
	PackMeta    *string `json:"packMeta,omitempty" xml:"packMeta,omitempty"`
	// This parameter is required.
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	VersionID *string `json:"versionID,omitempty" xml:"versionID,omitempty"`
}

func (s GetApplicationLogsInput) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationLogsInput) GoString() string {
	return s.String()
}

func (s *GetApplicationLogsInput) GetBackwardLine() *int64 {
	return s.BackwardLine
}

func (s *GetApplicationLogsInput) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetApplicationLogsInput) GetForwardLine() *int64 {
	return s.ForwardLine
}

func (s *GetApplicationLogsInput) GetInstanceID() *string {
	return s.InstanceID
}

func (s *GetApplicationLogsInput) GetIsTail() *bool {
	return s.IsTail
}

func (s *GetApplicationLogsInput) GetMatch() *string {
	return s.Match
}

func (s *GetApplicationLogsInput) GetMessage() *string {
	return s.Message
}

func (s *GetApplicationLogsInput) GetOffset() *int64 {
	return s.Offset
}

func (s *GetApplicationLogsInput) GetPackID() *string {
	return s.PackID
}

func (s *GetApplicationLogsInput) GetPackMeta() *string {
	return s.PackMeta
}

func (s *GetApplicationLogsInput) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetApplicationLogsInput) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetApplicationLogsInput) GetVersionID() *string {
	return s.VersionID
}

func (s *GetApplicationLogsInput) SetBackwardLine(v int64) *GetApplicationLogsInput {
	s.BackwardLine = &v
	return s
}

func (s *GetApplicationLogsInput) SetEndTime(v int64) *GetApplicationLogsInput {
	s.EndTime = &v
	return s
}

func (s *GetApplicationLogsInput) SetForwardLine(v int64) *GetApplicationLogsInput {
	s.ForwardLine = &v
	return s
}

func (s *GetApplicationLogsInput) SetInstanceID(v string) *GetApplicationLogsInput {
	s.InstanceID = &v
	return s
}

func (s *GetApplicationLogsInput) SetIsTail(v bool) *GetApplicationLogsInput {
	s.IsTail = &v
	return s
}

func (s *GetApplicationLogsInput) SetMatch(v string) *GetApplicationLogsInput {
	s.Match = &v
	return s
}

func (s *GetApplicationLogsInput) SetMessage(v string) *GetApplicationLogsInput {
	s.Message = &v
	return s
}

func (s *GetApplicationLogsInput) SetOffset(v int64) *GetApplicationLogsInput {
	s.Offset = &v
	return s
}

func (s *GetApplicationLogsInput) SetPackID(v string) *GetApplicationLogsInput {
	s.PackID = &v
	return s
}

func (s *GetApplicationLogsInput) SetPackMeta(v string) *GetApplicationLogsInput {
	s.PackMeta = &v
	return s
}

func (s *GetApplicationLogsInput) SetStartTime(v int64) *GetApplicationLogsInput {
	s.StartTime = &v
	return s
}

func (s *GetApplicationLogsInput) SetTimestamp(v string) *GetApplicationLogsInput {
	s.Timestamp = &v
	return s
}

func (s *GetApplicationLogsInput) SetVersionID(v string) *GetApplicationLogsInput {
	s.VersionID = &v
	return s
}

func (s *GetApplicationLogsInput) Validate() error {
	return dara.Validate(s)
}

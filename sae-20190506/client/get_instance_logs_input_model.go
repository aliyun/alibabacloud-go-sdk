// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLogsInput interface {
	dara.Model
	String() string
	GoString() string
	SetBackwardLine(v int64) *GetInstanceLogsInput
	GetBackwardLine() *int64
	SetEndTime(v int64) *GetInstanceLogsInput
	GetEndTime() *int64
	SetForwardLine(v int64) *GetInstanceLogsInput
	GetForwardLine() *int64
	SetIsTail(v bool) *GetInstanceLogsInput
	GetIsTail() *bool
	SetMatch(v string) *GetInstanceLogsInput
	GetMatch() *string
	SetMessage(v string) *GetInstanceLogsInput
	GetMessage() *string
	SetOffset(v int64) *GetInstanceLogsInput
	GetOffset() *int64
	SetPackID(v string) *GetInstanceLogsInput
	GetPackID() *string
	SetPackMeta(v string) *GetInstanceLogsInput
	GetPackMeta() *string
	SetStartTime(v int64) *GetInstanceLogsInput
	GetStartTime() *int64
	SetTimestamp(v string) *GetInstanceLogsInput
	GetTimestamp() *string
	SetVersionID(v string) *GetInstanceLogsInput
	GetVersionID() *string
}

type GetInstanceLogsInput struct {
	BackwardLine *int64 `json:"backwardLine,omitempty" xml:"backwardLine,omitempty"`
	// This parameter is required.
	EndTime     *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ForwardLine *int64  `json:"forwardLine,omitempty" xml:"forwardLine,omitempty"`
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

func (s GetInstanceLogsInput) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLogsInput) GoString() string {
	return s.String()
}

func (s *GetInstanceLogsInput) GetBackwardLine() *int64 {
	return s.BackwardLine
}

func (s *GetInstanceLogsInput) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetInstanceLogsInput) GetForwardLine() *int64 {
	return s.ForwardLine
}

func (s *GetInstanceLogsInput) GetIsTail() *bool {
	return s.IsTail
}

func (s *GetInstanceLogsInput) GetMatch() *string {
	return s.Match
}

func (s *GetInstanceLogsInput) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceLogsInput) GetOffset() *int64 {
	return s.Offset
}

func (s *GetInstanceLogsInput) GetPackID() *string {
	return s.PackID
}

func (s *GetInstanceLogsInput) GetPackMeta() *string {
	return s.PackMeta
}

func (s *GetInstanceLogsInput) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetInstanceLogsInput) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetInstanceLogsInput) GetVersionID() *string {
	return s.VersionID
}

func (s *GetInstanceLogsInput) SetBackwardLine(v int64) *GetInstanceLogsInput {
	s.BackwardLine = &v
	return s
}

func (s *GetInstanceLogsInput) SetEndTime(v int64) *GetInstanceLogsInput {
	s.EndTime = &v
	return s
}

func (s *GetInstanceLogsInput) SetForwardLine(v int64) *GetInstanceLogsInput {
	s.ForwardLine = &v
	return s
}

func (s *GetInstanceLogsInput) SetIsTail(v bool) *GetInstanceLogsInput {
	s.IsTail = &v
	return s
}

func (s *GetInstanceLogsInput) SetMatch(v string) *GetInstanceLogsInput {
	s.Match = &v
	return s
}

func (s *GetInstanceLogsInput) SetMessage(v string) *GetInstanceLogsInput {
	s.Message = &v
	return s
}

func (s *GetInstanceLogsInput) SetOffset(v int64) *GetInstanceLogsInput {
	s.Offset = &v
	return s
}

func (s *GetInstanceLogsInput) SetPackID(v string) *GetInstanceLogsInput {
	s.PackID = &v
	return s
}

func (s *GetInstanceLogsInput) SetPackMeta(v string) *GetInstanceLogsInput {
	s.PackMeta = &v
	return s
}

func (s *GetInstanceLogsInput) SetStartTime(v int64) *GetInstanceLogsInput {
	s.StartTime = &v
	return s
}

func (s *GetInstanceLogsInput) SetTimestamp(v string) *GetInstanceLogsInput {
	s.Timestamp = &v
	return s
}

func (s *GetInstanceLogsInput) SetVersionID(v string) *GetInstanceLogsInput {
	s.VersionID = &v
	return s
}

func (s *GetInstanceLogsInput) Validate() error {
	return dara.Validate(s)
}

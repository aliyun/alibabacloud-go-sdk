// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAgentConversationLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentGzipB64(v string) *UploadAgentConversationLogRequest
	GetContentGzipB64() *string
	SetFileName(v string) *UploadAgentConversationLogRequest
	GetFileName() *string
	SetTimeRangeEnd(v string) *UploadAgentConversationLogRequest
	GetTimeRangeEnd() *string
	SetTimeRangeStart(v string) *UploadAgentConversationLogRequest
	GetTimeRangeStart() *string
}

type UploadAgentConversationLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// H4sIAAAAAAAAEzMyMDLTNbDUNTQKMTSwMgAhbQMLKwMDBU8/N3+F4sq8kozUksxkheSczNS8EoWc/HSF4pLEohIuIwydhgR1pualcAEAN/JXaXIAAAA=
	ContentGzipB64 *string `json:"contentGzipB64,omitempty" xml:"contentGzipB64,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// client.log
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 2026-09-12T11:00:00+08:00
	TimeRangeEnd *string `json:"timeRangeEnd,omitempty" xml:"timeRangeEnd,omitempty"`
	// example:
	//
	// 2026-09-12T10:00:00+08:00
	TimeRangeStart *string `json:"timeRangeStart,omitempty" xml:"timeRangeStart,omitempty"`
}

func (s UploadAgentConversationLogRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentConversationLogRequest) GoString() string {
	return s.String()
}

func (s *UploadAgentConversationLogRequest) GetContentGzipB64() *string {
	return s.ContentGzipB64
}

func (s *UploadAgentConversationLogRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadAgentConversationLogRequest) GetTimeRangeEnd() *string {
	return s.TimeRangeEnd
}

func (s *UploadAgentConversationLogRequest) GetTimeRangeStart() *string {
	return s.TimeRangeStart
}

func (s *UploadAgentConversationLogRequest) SetContentGzipB64(v string) *UploadAgentConversationLogRequest {
	s.ContentGzipB64 = &v
	return s
}

func (s *UploadAgentConversationLogRequest) SetFileName(v string) *UploadAgentConversationLogRequest {
	s.FileName = &v
	return s
}

func (s *UploadAgentConversationLogRequest) SetTimeRangeEnd(v string) *UploadAgentConversationLogRequest {
	s.TimeRangeEnd = &v
	return s
}

func (s *UploadAgentConversationLogRequest) SetTimeRangeStart(v string) *UploadAgentConversationLogRequest {
	s.TimeRangeStart = &v
	return s
}

func (s *UploadAgentConversationLogRequest) Validate() error {
	return dara.Validate(s)
}

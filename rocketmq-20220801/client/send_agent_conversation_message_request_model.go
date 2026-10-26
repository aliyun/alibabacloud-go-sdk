// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAgentConversationMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachments(v []*SendAgentConversationMessageRequestAttachments) *SendAgentConversationMessageRequest
	GetAttachments() []*SendAgentConversationMessageRequestAttachments
	SetClientContext(v string) *SendAgentConversationMessageRequest
	GetClientContext() *string
	SetImages(v []*SendAgentConversationMessageRequestImages) *SendAgentConversationMessageRequest
	GetImages() []*SendAgentConversationMessageRequestImages
	SetMessage(v string) *SendAgentConversationMessageRequest
	GetMessage() *string
	SetMessageType(v string) *SendAgentConversationMessageRequest
	GetMessageType() *string
	SetTicketRequestId(v string) *SendAgentConversationMessageRequest
	GetTicketRequestId() *string
}

type SendAgentConversationMessageRequest struct {
	Attachments []*SendAgentConversationMessageRequestAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// example:
	//
	// 告警时间为 2026-08-19 19:30:00
	ClientContext *string                                      `json:"clientContext,omitempty" xml:"clientContext,omitempty"`
	Images        []*SendAgentConversationMessageRequestImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 实例 rmq-cn-example 的消息为什么持续堆积？
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// ticket_request
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// example:
	//
	// ticket-request-example-001
	TicketRequestId *string `json:"ticketRequestId,omitempty" xml:"ticketRequestId,omitempty"`
}

func (s SendAgentConversationMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendAgentConversationMessageRequest) GoString() string {
	return s.String()
}

func (s *SendAgentConversationMessageRequest) GetAttachments() []*SendAgentConversationMessageRequestAttachments {
	return s.Attachments
}

func (s *SendAgentConversationMessageRequest) GetClientContext() *string {
	return s.ClientContext
}

func (s *SendAgentConversationMessageRequest) GetImages() []*SendAgentConversationMessageRequestImages {
	return s.Images
}

func (s *SendAgentConversationMessageRequest) GetMessage() *string {
	return s.Message
}

func (s *SendAgentConversationMessageRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *SendAgentConversationMessageRequest) GetTicketRequestId() *string {
	return s.TicketRequestId
}

func (s *SendAgentConversationMessageRequest) SetAttachments(v []*SendAgentConversationMessageRequestAttachments) *SendAgentConversationMessageRequest {
	s.Attachments = v
	return s
}

func (s *SendAgentConversationMessageRequest) SetClientContext(v string) *SendAgentConversationMessageRequest {
	s.ClientContext = &v
	return s
}

func (s *SendAgentConversationMessageRequest) SetImages(v []*SendAgentConversationMessageRequestImages) *SendAgentConversationMessageRequest {
	s.Images = v
	return s
}

func (s *SendAgentConversationMessageRequest) SetMessage(v string) *SendAgentConversationMessageRequest {
	s.Message = &v
	return s
}

func (s *SendAgentConversationMessageRequest) SetMessageType(v string) *SendAgentConversationMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendAgentConversationMessageRequest) SetTicketRequestId(v string) *SendAgentConversationMessageRequest {
	s.TicketRequestId = &v
	return s
}

func (s *SendAgentConversationMessageRequest) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendAgentConversationMessageRequestAttachments struct {
	// This parameter is required.
	//
	// example:
	//
	// client.log
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 2
	KeptLines *int64 `json:"keptLines,omitempty" xml:"keptLines,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// client_log
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// client-logs/0198c8a91f4d7f01a3c21e99/client.log
	SandboxPath *string `json:"sandboxPath,omitempty" xml:"sandboxPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 114
	SizeBytes *int64 `json:"sizeBytes,omitempty" xml:"sizeBytes,omitempty"`
	// example:
	//
	// 2026-09-12T11:00:00+08:00
	TimeRangeEnd *string `json:"timeRangeEnd,omitempty" xml:"timeRangeEnd,omitempty"`
	// example:
	//
	// 2026-09-12T10:00:00+08:00
	TimeRangeStart *string `json:"timeRangeStart,omitempty" xml:"timeRangeStart,omitempty"`
	// example:
	//
	// 2
	TotalLines *int64 `json:"totalLines,omitempty" xml:"totalLines,omitempty"`
}

func (s SendAgentConversationMessageRequestAttachments) String() string {
	return dara.Prettify(s)
}

func (s SendAgentConversationMessageRequestAttachments) GoString() string {
	return s.String()
}

func (s *SendAgentConversationMessageRequestAttachments) GetFileName() *string {
	return s.FileName
}

func (s *SendAgentConversationMessageRequestAttachments) GetKeptLines() *int64 {
	return s.KeptLines
}

func (s *SendAgentConversationMessageRequestAttachments) GetKind() *string {
	return s.Kind
}

func (s *SendAgentConversationMessageRequestAttachments) GetSandboxPath() *string {
	return s.SandboxPath
}

func (s *SendAgentConversationMessageRequestAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *SendAgentConversationMessageRequestAttachments) GetTimeRangeEnd() *string {
	return s.TimeRangeEnd
}

func (s *SendAgentConversationMessageRequestAttachments) GetTimeRangeStart() *string {
	return s.TimeRangeStart
}

func (s *SendAgentConversationMessageRequestAttachments) GetTotalLines() *int64 {
	return s.TotalLines
}

func (s *SendAgentConversationMessageRequestAttachments) SetFileName(v string) *SendAgentConversationMessageRequestAttachments {
	s.FileName = &v
	return s
}

func (s *SendAgentConversationMessageRequestAttachments) SetKeptLines(v int64) *SendAgentConversationMessageRequestAttachments {
	s.KeptLines = &v
	return s
}

func (s *SendAgentConversationMessageRequestAttachments) SetKind(v string) *SendAgentConversationMessageRequestAttachments {
	s.Kind = &v
	return s
}

func (s *SendAgentConversationMessageRequestAttachments) SetSandboxPath(v string) *SendAgentConversationMessageRequestAttachments {
	s.SandboxPath = &v
	return s
}

func (s *SendAgentConversationMessageRequestAttachments) SetSizeBytes(v int64) *SendAgentConversationMessageRequestAttachments {
	s.SizeBytes = &v
	return s
}

func (s *SendAgentConversationMessageRequestAttachments) SetTimeRangeEnd(v string) *SendAgentConversationMessageRequestAttachments {
	s.TimeRangeEnd = &v
	return s
}

func (s *SendAgentConversationMessageRequestAttachments) SetTimeRangeStart(v string) *SendAgentConversationMessageRequestAttachments {
	s.TimeRangeStart = &v
	return s
}

func (s *SendAgentConversationMessageRequestAttachments) SetTotalLines(v int64) *SendAgentConversationMessageRequestAttachments {
	s.TotalLines = &v
	return s
}

func (s *SendAgentConversationMessageRequestAttachments) Validate() error {
	return dara.Validate(s)
}

type SendAgentConversationMessageRequestImages struct {
	// This parameter is required.
	//
	// example:
	//
	// iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAC0lEQVR4nGNgAAIAAAUAAXpeqz8AAAAASUVORK5CYII=
	DataBase64 *string `json:"dataBase64,omitempty" xml:"dataBase64,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// image/png
	MediaType *string `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
}

func (s SendAgentConversationMessageRequestImages) String() string {
	return dara.Prettify(s)
}

func (s SendAgentConversationMessageRequestImages) GoString() string {
	return s.String()
}

func (s *SendAgentConversationMessageRequestImages) GetDataBase64() *string {
	return s.DataBase64
}

func (s *SendAgentConversationMessageRequestImages) GetMediaType() *string {
	return s.MediaType
}

func (s *SendAgentConversationMessageRequestImages) SetDataBase64(v string) *SendAgentConversationMessageRequestImages {
	s.DataBase64 = &v
	return s
}

func (s *SendAgentConversationMessageRequestImages) SetMediaType(v string) *SendAgentConversationMessageRequestImages {
	s.MediaType = &v
	return s
}

func (s *SendAgentConversationMessageRequestImages) Validate() error {
	return dara.Validate(s)
}

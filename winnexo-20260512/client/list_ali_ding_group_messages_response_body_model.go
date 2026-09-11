// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliDingGroupMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAliDingGroupMessagesResponseBody
	GetCode() *string
	SetHasMore(v bool) *ListAliDingGroupMessagesResponseBody
	GetHasMore() *bool
	SetItems(v []*ListAliDingGroupMessagesResponseBodyItems) *ListAliDingGroupMessagesResponseBody
	GetItems() []*ListAliDingGroupMessagesResponseBodyItems
	SetMessage(v string) *ListAliDingGroupMessagesResponseBody
	GetMessage() *string
	SetNextTime(v string) *ListAliDingGroupMessagesResponseBody
	GetNextTime() *string
	SetRequestId(v string) *ListAliDingGroupMessagesResponseBody
	GetRequestId() *string
}

type ListAliDingGroupMessagesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether more pages are available.
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The file information.
	Items []*ListAliDingGroupMessagesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The time when the next plan is scheduled.
	//
	// example:
	//
	// 2026-09-08T09:01:00+08:00
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-id
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAliDingGroupMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAliDingGroupMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliDingGroupMessagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAliDingGroupMessagesResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListAliDingGroupMessagesResponseBody) GetItems() []*ListAliDingGroupMessagesResponseBodyItems {
	return s.Items
}

func (s *ListAliDingGroupMessagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAliDingGroupMessagesResponseBody) GetNextTime() *string {
	return s.NextTime
}

func (s *ListAliDingGroupMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAliDingGroupMessagesResponseBody) SetCode(v string) *ListAliDingGroupMessagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBody) SetHasMore(v bool) *ListAliDingGroupMessagesResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBody) SetItems(v []*ListAliDingGroupMessagesResponseBodyItems) *ListAliDingGroupMessagesResponseBody {
	s.Items = v
	return s
}

func (s *ListAliDingGroupMessagesResponseBody) SetMessage(v string) *ListAliDingGroupMessagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBody) SetNextTime(v string) *ListAliDingGroupMessagesResponseBody {
	s.NextTime = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBody) SetRequestId(v string) *ListAliDingGroupMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAliDingGroupMessagesResponseBodyItems struct {
	// The comment attachments.
	Attachments []*ListAliDingGroupMessagesResponseBodyItemsAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// The returned content.
	//
	// example:
	//
	// See the attachment
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2026-09-08 09:01:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The message ID.
	//
	// example:
	//
	// msg-example
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The message type. Valid values:
	//
	// - **MARKDOWN**: Markdown message.
	//
	// - **ACTIONCARD**: card message.
	//
	// > Markdown messages do not support message buttons.
	//
	// example:
	//
	// FILE
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The DingTalk ID of the business-side customer service representative.
	//
	// example:
	//
	// user-example
	SenderId *string `json:"senderId,omitempty" xml:"senderId,omitempty"`
	// The name of the message sender.
	//
	// example:
	//
	// John
	SenderName *string `json:"senderName,omitempty" xml:"senderName,omitempty"`
}

func (s ListAliDingGroupMessagesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAliDingGroupMessagesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAliDingGroupMessagesResponseBodyItems) GetAttachments() []*ListAliDingGroupMessagesResponseBodyItemsAttachments {
	return s.Attachments
}

func (s *ListAliDingGroupMessagesResponseBodyItems) GetContent() *string {
	return s.Content
}

func (s *ListAliDingGroupMessagesResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAliDingGroupMessagesResponseBodyItems) GetMessageId() *string {
	return s.MessageId
}

func (s *ListAliDingGroupMessagesResponseBodyItems) GetMessageType() *string {
	return s.MessageType
}

func (s *ListAliDingGroupMessagesResponseBodyItems) GetSenderId() *string {
	return s.SenderId
}

func (s *ListAliDingGroupMessagesResponseBodyItems) GetSenderName() *string {
	return s.SenderName
}

func (s *ListAliDingGroupMessagesResponseBodyItems) SetAttachments(v []*ListAliDingGroupMessagesResponseBodyItemsAttachments) *ListAliDingGroupMessagesResponseBodyItems {
	s.Attachments = v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItems) SetContent(v string) *ListAliDingGroupMessagesResponseBodyItems {
	s.Content = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItems) SetCreateTime(v string) *ListAliDingGroupMessagesResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItems) SetMessageId(v string) *ListAliDingGroupMessagesResponseBodyItems {
	s.MessageId = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItems) SetMessageType(v string) *ListAliDingGroupMessagesResponseBodyItems {
	s.MessageType = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItems) SetSenderId(v string) *ListAliDingGroupMessagesResponseBodyItems {
	s.SenderId = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItems) SetSenderName(v string) *ListAliDingGroupMessagesResponseBodyItems {
	s.SenderName = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItems) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAliDingGroupMessagesResponseBodyItemsAttachments struct {
	// The attachment ID.
	//
	// example:
	//
	// attachment-example
	AttachmentId *string `json:"attachmentId,omitempty" xml:"attachmentId,omitempty"`
	// The attachment type.
	//
	// example:
	//
	// FILE
	AttachmentType *string `json:"attachmentType,omitempty" xml:"attachmentType,omitempty"`
	// The execution duration of the asynchronous task.
	//
	// example:
	//
	// 1000
	DurationMs *int64 `json:"durationMs,omitempty" xml:"durationMs,omitempty"`
	// The new file name. This parameter is optional. If you do not specify this parameter or set it to an empty string, the original file name is retained.
	//
	// example:
	//
	// Plan.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The file size, in **bytes**.
	//
	// example:
	//
	// 102400
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The thumbnail height, in pixels.
	//
	// example:
	//
	// 1080
	Height *int64 `json:"height,omitempty" xml:"height,omitempty"`
	// The media type. The file name extension is in uppercase, such as XLS, DOC, DOCX, PDF, or XLSX.
	//
	// example:
	//
	// application/pdf
	MimeType *string `json:"mimeType,omitempty" xml:"mimeType,omitempty"`
	// The image width, in pixels.
	//
	// example:
	//
	// 1920
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s ListAliDingGroupMessagesResponseBodyItemsAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListAliDingGroupMessagesResponseBodyItemsAttachments) GoString() string {
	return s.String()
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) GetAttachmentType() *string {
	return s.AttachmentType
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) GetDurationMs() *int64 {
	return s.DurationMs
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) GetFileName() *string {
	return s.FileName
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) GetFileSize() *int64 {
	return s.FileSize
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) GetHeight() *int64 {
	return s.Height
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) GetMimeType() *string {
	return s.MimeType
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) GetWidth() *int64 {
	return s.Width
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) SetAttachmentId(v string) *ListAliDingGroupMessagesResponseBodyItemsAttachments {
	s.AttachmentId = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) SetAttachmentType(v string) *ListAliDingGroupMessagesResponseBodyItemsAttachments {
	s.AttachmentType = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) SetDurationMs(v int64) *ListAliDingGroupMessagesResponseBodyItemsAttachments {
	s.DurationMs = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) SetFileName(v string) *ListAliDingGroupMessagesResponseBodyItemsAttachments {
	s.FileName = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) SetFileSize(v int64) *ListAliDingGroupMessagesResponseBodyItemsAttachments {
	s.FileSize = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) SetHeight(v int64) *ListAliDingGroupMessagesResponseBodyItemsAttachments {
	s.Height = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) SetMimeType(v string) *ListAliDingGroupMessagesResponseBodyItemsAttachments {
	s.MimeType = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) SetWidth(v int64) *ListAliDingGroupMessagesResponseBodyItemsAttachments {
	s.Width = &v
	return s
}

func (s *ListAliDingGroupMessagesResponseBodyItemsAttachments) Validate() error {
	return dara.Validate(s)
}

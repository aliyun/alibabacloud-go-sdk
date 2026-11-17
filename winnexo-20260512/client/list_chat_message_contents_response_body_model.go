// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatMessageContentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListChatMessageContentsResponseBody
	GetCode() *string
	SetHasMore(v bool) *ListChatMessageContentsResponseBody
	GetHasMore() *bool
	SetItems(v []*ListChatMessageContentsResponseBodyItems) *ListChatMessageContentsResponseBody
	GetItems() []*ListChatMessageContentsResponseBodyItems
	SetMessage(v string) *ListChatMessageContentsResponseBody
	GetMessage() *string
	SetPage(v int32) *ListChatMessageContentsResponseBody
	GetPage() *int32
	SetPageSize(v int32) *ListChatMessageContentsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListChatMessageContentsResponseBody
	GetRequestId() *string
	SetSourceId(v string) *ListChatMessageContentsResponseBody
	GetSourceId() *string
	SetTotal(v int64) *ListChatMessageContentsResponseBody
	GetTotal() *int64
}

type ListChatMessageContentsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Specifies whether there is a next page.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The list of MCP cards.
	Items []*ListChatMessageContentsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The description of the status code.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The maximum number of data records returned in this request.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// trace-chat-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the original project.
	//
	// example:
	//
	// source-chat-1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListChatMessageContentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatMessageContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatMessageContentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatMessageContentsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListChatMessageContentsResponseBody) GetItems() []*ListChatMessageContentsResponseBodyItems {
	return s.Items
}

func (s *ListChatMessageContentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatMessageContentsResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *ListChatMessageContentsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChatMessageContentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatMessageContentsResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *ListChatMessageContentsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListChatMessageContentsResponseBody) SetCode(v string) *ListChatMessageContentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatMessageContentsResponseBody) SetHasMore(v bool) *ListChatMessageContentsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListChatMessageContentsResponseBody) SetItems(v []*ListChatMessageContentsResponseBodyItems) *ListChatMessageContentsResponseBody {
	s.Items = v
	return s
}

func (s *ListChatMessageContentsResponseBody) SetMessage(v string) *ListChatMessageContentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatMessageContentsResponseBody) SetPage(v int32) *ListChatMessageContentsResponseBody {
	s.Page = &v
	return s
}

func (s *ListChatMessageContentsResponseBody) SetPageSize(v int32) *ListChatMessageContentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChatMessageContentsResponseBody) SetRequestId(v string) *ListChatMessageContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatMessageContentsResponseBody) SetSourceId(v string) *ListChatMessageContentsResponseBody {
	s.SourceId = &v
	return s
}

func (s *ListChatMessageContentsResponseBody) SetTotal(v int64) *ListChatMessageContentsResponseBody {
	s.Total = &v
	return s
}

func (s *ListChatMessageContentsResponseBody) Validate() error {
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

type ListChatMessageContentsResponseBodyItems struct {
	// The content type. Valid values:
	//
	// - 0: Text.
	//
	// - 1: Image material.
	//
	// - 2: Handwritten text.
	//
	// example:
	//
	// []
	Attachments []*ListChatMessageContentsResponseBodyItemsAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// The ID of the message.
	//
	// example:
	//
	// msg-1
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The message type. Valid values:
	//
	// - **MARKDOWN**: Markdown message.
	//
	// - **ACTIONCARD**: Card message.
	//
	// > Message buttons cannot be used in Markdown messages.
	//
	// example:
	//
	// POST
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The ID of the service record.
	//
	// example:
	//
	// 123
	RecordId *int64 `json:"recordId,omitempty" xml:"recordId,omitempty"`
	// The DingTalk ID of the customer service representative.
	//
	// example:
	//
	// sender-1
	SenderId *string `json:"senderId,omitempty" xml:"senderId,omitempty"`
	// The name of the message sender.
	//
	// example:
	//
	// Zhang San
	SenderName *string `json:"senderName,omitempty" xml:"senderName,omitempty"`
	// The time when the message was sent.
	//
	// example:
	//
	// 2026-11-09T10:00:00+08:00
	SentAt *string `json:"sentAt,omitempty" xml:"sentAt,omitempty"`
	// The batch text to be translated.
	//
	// example:
	//
	// Please check the attachment
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ListChatMessageContentsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListChatMessageContentsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListChatMessageContentsResponseBodyItems) GetAttachments() []*ListChatMessageContentsResponseBodyItemsAttachments {
	return s.Attachments
}

func (s *ListChatMessageContentsResponseBodyItems) GetMessageId() *string {
	return s.MessageId
}

func (s *ListChatMessageContentsResponseBodyItems) GetMessageType() *string {
	return s.MessageType
}

func (s *ListChatMessageContentsResponseBodyItems) GetRecordId() *int64 {
	return s.RecordId
}

func (s *ListChatMessageContentsResponseBodyItems) GetSenderId() *string {
	return s.SenderId
}

func (s *ListChatMessageContentsResponseBodyItems) GetSenderName() *string {
	return s.SenderName
}

func (s *ListChatMessageContentsResponseBodyItems) GetSentAt() *string {
	return s.SentAt
}

func (s *ListChatMessageContentsResponseBodyItems) GetText() *string {
	return s.Text
}

func (s *ListChatMessageContentsResponseBodyItems) SetAttachments(v []*ListChatMessageContentsResponseBodyItemsAttachments) *ListChatMessageContentsResponseBodyItems {
	s.Attachments = v
	return s
}

func (s *ListChatMessageContentsResponseBodyItems) SetMessageId(v string) *ListChatMessageContentsResponseBodyItems {
	s.MessageId = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItems) SetMessageType(v string) *ListChatMessageContentsResponseBodyItems {
	s.MessageType = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItems) SetRecordId(v int64) *ListChatMessageContentsResponseBodyItems {
	s.RecordId = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItems) SetSenderId(v string) *ListChatMessageContentsResponseBodyItems {
	s.SenderId = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItems) SetSenderName(v string) *ListChatMessageContentsResponseBodyItems {
	s.SenderName = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItems) SetSentAt(v string) *ListChatMessageContentsResponseBodyItems {
	s.SentAt = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItems) SetText(v string) *ListChatMessageContentsResponseBodyItems {
	s.Text = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItems) Validate() error {
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

type ListChatMessageContentsResponseBodyItemsAttachments struct {
	// The ID of the attachment.
	//
	// example:
	//
	// file-1
	AttachmentId *string `json:"attachmentId,omitempty" xml:"attachmentId,omitempty"`
	// The total number of characters in the content.
	//
	// example:
	//
	// 4
	CharCount *int64 `json:"charCount,omitempty" xml:"charCount,omitempty"`
	// The number of chunks.
	//
	// example:
	//
	// 1
	ChunkCount *int32 `json:"chunkCount,omitempty" xml:"chunkCount,omitempty"`
	// The content of the user message.
	//
	// example:
	//
	// Attachment content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The reason why the content is empty.
	//
	// example:
	//
	// None
	EmptyReason *string `json:"emptyReason,omitempty" xml:"emptyReason,omitempty"`
	// The failure code of the batch. Valid values:
	//
	// - ARREARAGE: The account has an overdue payment.
	//
	// - NO_TARGET: No valid inspection instance is available.
	//
	// - TIMEOUT: The batch execution timed out.
	//
	// - SYSTEM_ERROR: A system error occurred.
	//
	// - ALL_INSTANCE_FAILED: All instance inspections failed.
	//
	// This parameter is empty if the batch does not fail.
	//
	// example:
	//
	// ARREARAGE
	FailureCode *string `json:"failureCode,omitempty" xml:"failureCode,omitempty"`
	// The failure summary of the asynchronous task. This parameter is not returned if no failure information is available. You can use the taskId and requestId to troubleshoot the issue.
	//
	// example:
	//
	// Scheduling failed
	FailureReason *string `json:"failureReason,omitempty" xml:"failureReason,omitempty"`
	// The new file name. This parameter is optional. If you do not specify this parameter or specify an empty string, the original file name is retained.
	//
	// example:
	//
	// Weekly_Report.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Specifies whether the standard package instance is found.
	//
	// example:
	//
	// true
	Found *bool `json:"found,omitempty" xml:"found,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// msg-1
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The type of the attachment.
	//
	// example:
	//
	// doc
	Modality *string `json:"modality,omitempty" xml:"modality,omitempty"`
	// The status of the data source. Valid values:
	//
	// - **1**: Online.
	//
	// - **0**: Offline.
	//
	// example:
	//
	// PARSED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Specifies whether the scan result is truncated.
	//
	// example:
	//
	// false
	Truncated *bool `json:"truncated,omitempty" xml:"truncated,omitempty"`
}

func (s ListChatMessageContentsResponseBodyItemsAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListChatMessageContentsResponseBodyItemsAttachments) GoString() string {
	return s.String()
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetCharCount() *int64 {
	return s.CharCount
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetChunkCount() *int32 {
	return s.ChunkCount
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetContent() *string {
	return s.Content
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetEmptyReason() *string {
	return s.EmptyReason
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetFailureCode() *string {
	return s.FailureCode
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetFailureReason() *string {
	return s.FailureReason
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetFileName() *string {
	return s.FileName
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetFound() *bool {
	return s.Found
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetMessageId() *string {
	return s.MessageId
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetModality() *string {
	return s.Modality
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetStatus() *string {
	return s.Status
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) GetTruncated() *bool {
	return s.Truncated
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetAttachmentId(v string) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.AttachmentId = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetCharCount(v int64) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.CharCount = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetChunkCount(v int32) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.ChunkCount = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetContent(v string) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.Content = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetEmptyReason(v string) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.EmptyReason = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetFailureCode(v string) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.FailureCode = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetFailureReason(v string) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.FailureReason = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetFileName(v string) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.FileName = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetFound(v bool) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.Found = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetMessageId(v string) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.MessageId = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetModality(v string) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.Modality = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetStatus(v string) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.Status = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) SetTruncated(v bool) *ListChatMessageContentsResponseBodyItemsAttachments {
	s.Truncated = &v
	return s
}

func (s *ListChatMessageContentsResponseBodyItemsAttachments) Validate() error {
	return dara.Validate(s)
}

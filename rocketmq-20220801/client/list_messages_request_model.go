// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListMessagesRequest
	GetEndTime() *string
	SetLiteTopicName(v string) *ListMessagesRequest
	GetLiteTopicName() *string
	SetMessageId(v string) *ListMessagesRequest
	GetMessageId() *string
	SetMessageKey(v string) *ListMessagesRequest
	GetMessageKey() *string
	SetPageNumber(v int32) *ListMessagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMessagesRequest
	GetPageSize() *int32
	SetScrollId(v string) *ListMessagesRequest
	GetScrollId() *string
	SetStartTime(v string) *ListMessagesRequest
	GetStartTime() *string
}

type ListMessagesRequest struct {
	// The end of the time range to query.
	//
	// example:
	//
	// 2024-09-09 09:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// abc
	LiteTopicName *string `json:"liteTopicName,omitempty" xml:"liteTopicName,omitempty"`
	// Message Id.
	//
	// example:
	//
	// 7F00000100207A4F0F294A938F7807AE
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// Message key.
	//
	// example:
	//
	// XSCBillResult
	MessageKey *string `json:"messageKey,omitempty" xml:"messageKey,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The scroll ID of the request.
	//
	// You do not need to configure this parameter for the first page. This parameter is included in the pagination request based on the result returned for the first page.
	//
	// example:
	//
	// B13D0B07-F24B-4790-88D8-D47A38063D00
	ScrollId *string `json:"scrollId,omitempty" xml:"scrollId,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2024-09-09 08:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListMessagesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListMessagesRequest) GetLiteTopicName() *string {
	return s.LiteTopicName
}

func (s *ListMessagesRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *ListMessagesRequest) GetMessageKey() *string {
	return s.MessageKey
}

func (s *ListMessagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMessagesRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *ListMessagesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMessagesRequest) SetEndTime(v string) *ListMessagesRequest {
	s.EndTime = &v
	return s
}

func (s *ListMessagesRequest) SetLiteTopicName(v string) *ListMessagesRequest {
	s.LiteTopicName = &v
	return s
}

func (s *ListMessagesRequest) SetMessageId(v string) *ListMessagesRequest {
	s.MessageId = &v
	return s
}

func (s *ListMessagesRequest) SetMessageKey(v string) *ListMessagesRequest {
	s.MessageKey = &v
	return s
}

func (s *ListMessagesRequest) SetPageNumber(v int32) *ListMessagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMessagesRequest) SetPageSize(v int32) *ListMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessagesRequest) SetScrollId(v string) *ListMessagesRequest {
	s.ScrollId = &v
	return s
}

func (s *ListMessagesRequest) SetStartTime(v string) *ListMessagesRequest {
	s.StartTime = &v
	return s
}

func (s *ListMessagesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTracesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTracesRequest
	GetEndTime() *string
	SetMessageId(v string) *ListTracesRequest
	GetMessageId() *string
	SetMessageKey(v string) *ListTracesRequest
	GetMessageKey() *string
	SetPageNumber(v int32) *ListTracesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTracesRequest
	GetPageSize() *int32
	SetQueryType(v string) *ListTracesRequest
	GetQueryType() *string
	SetStartTime(v string) *ListTracesRequest
	GetStartTime() *string
}

type ListTracesRequest struct {
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-19 10:10:09
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The message ID.
	//
	// This parameter is required if you set queryType to MESSAGE_ID.
	//
	// example:
	//
	// 0100163E0EC1F1965C04C7906700000000
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The message key.
	//
	// This parameter is required if you set queryType to MESSAGE_ID.
	//
	// example:
	//
	// order_ceating
	MessageKey *string `json:"messageKey,omitempty" xml:"messageKey,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The query type.
	//
	// Valid values:
	//
	// 	- MESSAGE_ID
	//
	// 	- MESSAGE_KEY
	//
	// 	- TOPIC
	//
	// This parameter is required.
	//
	// example:
	//
	// MESSAGE_ID
	QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-10 10:42:11
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListTracesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTracesRequest) GoString() string {
	return s.String()
}

func (s *ListTracesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTracesRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *ListTracesRequest) GetMessageKey() *string {
	return s.MessageKey
}

func (s *ListTracesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTracesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTracesRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *ListTracesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTracesRequest) SetEndTime(v string) *ListTracesRequest {
	s.EndTime = &v
	return s
}

func (s *ListTracesRequest) SetMessageId(v string) *ListTracesRequest {
	s.MessageId = &v
	return s
}

func (s *ListTracesRequest) SetMessageKey(v string) *ListTracesRequest {
	s.MessageKey = &v
	return s
}

func (s *ListTracesRequest) SetPageNumber(v int32) *ListTracesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTracesRequest) SetPageSize(v int32) *ListTracesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTracesRequest) SetQueryType(v string) *ListTracesRequest {
	s.QueryType = &v
	return s
}

func (s *ListTracesRequest) SetStartTime(v string) *ListTracesRequest {
	s.StartTime = &v
	return s
}

func (s *ListTracesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMessagesRequest
	GetMaxResults() *int32
	SetMessageLevel(v string) *ListMessagesRequest
	GetMessageLevel() *string
	SetMessageTimeEarlierThan(v int64) *ListMessagesRequest
	GetMessageTimeEarlierThan() *int64
	SetMessageTimeLaterThan(v int64) *ListMessagesRequest
	GetMessageTimeLaterThan() *int64
	SetMessageType(v string) *ListMessagesRequest
	GetMessageType() *string
	SetNextToken(v string) *ListMessagesRequest
	GetNextToken() *string
}

type ListMessagesRequest struct {
	// The maximum number of records to return in this request.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The message level.
	//
	// example:
	//
	// WARNING
	MessageLevel *string `json:"MessageLevel,omitempty" xml:"MessageLevel,omitempty"`
	// Filters messages with a time earlier than the specified value.
	//
	// example:
	//
	// 1740019610
	MessageTimeEarlierThan *int64 `json:"MessageTimeEarlierThan,omitempty" xml:"MessageTimeEarlierThan,omitempty"`
	// Filters messages with a time later than the specified value.
	//
	// example:
	//
	// 1740019609
	MessageTimeLaterThan *int64 `json:"MessageTimeLaterThan,omitempty" xml:"MessageTimeLaterThan,omitempty"`
	// The message type.
	//
	// example:
	//
	// SUB_PROTECTION_POLICY_MODIFIED
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	// The pagination token. If there is a next page, this field has a return value. This parameter indicates that there is a next page as long as data is returned. You can use the returned NextToken as a request parameter to obtain the next page of data until Null is returned, which indicates that all data has been retrieved.
	//
	// example:
	//
	// cae***********99
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListMessagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMessagesRequest) GetMessageLevel() *string {
	return s.MessageLevel
}

func (s *ListMessagesRequest) GetMessageTimeEarlierThan() *int64 {
	return s.MessageTimeEarlierThan
}

func (s *ListMessagesRequest) GetMessageTimeLaterThan() *int64 {
	return s.MessageTimeLaterThan
}

func (s *ListMessagesRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *ListMessagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMessagesRequest) SetMaxResults(v int32) *ListMessagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMessagesRequest) SetMessageLevel(v string) *ListMessagesRequest {
	s.MessageLevel = &v
	return s
}

func (s *ListMessagesRequest) SetMessageTimeEarlierThan(v int64) *ListMessagesRequest {
	s.MessageTimeEarlierThan = &v
	return s
}

func (s *ListMessagesRequest) SetMessageTimeLaterThan(v int64) *ListMessagesRequest {
	s.MessageTimeLaterThan = &v
	return s
}

func (s *ListMessagesRequest) SetMessageType(v string) *ListMessagesRequest {
	s.MessageType = &v
	return s
}

func (s *ListMessagesRequest) SetNextToken(v string) *ListMessagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMessagesRequest) Validate() error {
	return dara.Validate(s)
}

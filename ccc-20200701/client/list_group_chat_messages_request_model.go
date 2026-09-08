// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupChatMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListGroupChatMessagesRequest
	GetInstanceId() *string
	SetJobId(v string) *ListGroupChatMessagesRequest
	GetJobId() *string
	SetNextPageToken(v string) *ListGroupChatMessagesRequest
	GetNextPageToken() *string
	SetPageSize(v int32) *ListGroupChatMessagesRequest
	GetPageSize() *int32
	SetSortOrder(v string) *ListGroupChatMessagesRequest
	GetSortOrder() *string
}

type ListGroupChatMessagesRequest struct {
	// Cloud Contact Center instance.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Call ID.
	//
	// example:
	//
	// chat-65382141036853491
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Token for the next page. This is a 32-character UUID. Leave this parameter empty when requesting the first page. For subsequent pages, use the NextPageToken value from the previous response.
	//
	// example:
	//
	// b2ad450b116e4f8396e58108acf5c020
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// Page size. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Sorting order. Default is descending.
	//
	// Valid values:
	//
	// ASC: ascending.
	//
	// DESC: descending.
	//
	// example:
	//
	// DESC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListGroupChatMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupChatMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListGroupChatMessagesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupChatMessagesRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListGroupChatMessagesRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListGroupChatMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGroupChatMessagesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListGroupChatMessagesRequest) SetInstanceId(v string) *ListGroupChatMessagesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGroupChatMessagesRequest) SetJobId(v string) *ListGroupChatMessagesRequest {
	s.JobId = &v
	return s
}

func (s *ListGroupChatMessagesRequest) SetNextPageToken(v string) *ListGroupChatMessagesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListGroupChatMessagesRequest) SetPageSize(v int32) *ListGroupChatMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupChatMessagesRequest) SetSortOrder(v string) *ListGroupChatMessagesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListGroupChatMessagesRequest) Validate() error {
	return dara.Validate(s)
}

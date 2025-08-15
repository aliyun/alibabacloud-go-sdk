// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListTopicsShrinkRequest
	GetFilter() *string
	SetMessageTypesShrink(v string) *ListTopicsShrinkRequest
	GetMessageTypesShrink() *string
	SetPageNumber(v int32) *ListTopicsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTopicsShrinkRequest
	GetPageSize() *int32
}

type ListTopicsShrinkRequest struct {
	// The filter condition for the query. If not provided, all topics under the instance will be queried.
	//
	// example:
	//
	// topic_test
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The message type of the topic.
	MessageTypesShrink *string `json:"messageTypes,omitempty" xml:"messageTypes,omitempty"`
	// Page number, indicating which page of results to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size, the maximum number of results to display per page.
	//
	// example:
	//
	// 3
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListTopicsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTopicsShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListTopicsShrinkRequest) GetMessageTypesShrink() *string {
	return s.MessageTypesShrink
}

func (s *ListTopicsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTopicsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTopicsShrinkRequest) SetFilter(v string) *ListTopicsShrinkRequest {
	s.Filter = &v
	return s
}

func (s *ListTopicsShrinkRequest) SetMessageTypesShrink(v string) *ListTopicsShrinkRequest {
	s.MessageTypesShrink = &v
	return s
}

func (s *ListTopicsShrinkRequest) SetPageNumber(v int32) *ListTopicsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTopicsShrinkRequest) SetPageSize(v int32) *ListTopicsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTopicsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

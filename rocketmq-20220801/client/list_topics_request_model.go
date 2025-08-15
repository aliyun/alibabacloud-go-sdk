// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListTopicsRequest
	GetFilter() *string
	SetMessageTypes(v []*string) *ListTopicsRequest
	GetMessageTypes() []*string
	SetPageNumber(v int32) *ListTopicsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTopicsRequest
	GetPageSize() *int32
}

type ListTopicsRequest struct {
	// The filter condition for the query. If not provided, all topics under the instance will be queried.
	//
	// example:
	//
	// topic_test
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The message type of the topic.
	MessageTypes []*string `json:"messageTypes,omitempty" xml:"messageTypes,omitempty" type:"Repeated"`
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

func (s ListTopicsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsRequest) GoString() string {
	return s.String()
}

func (s *ListTopicsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListTopicsRequest) GetMessageTypes() []*string {
	return s.MessageTypes
}

func (s *ListTopicsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTopicsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTopicsRequest) SetFilter(v string) *ListTopicsRequest {
	s.Filter = &v
	return s
}

func (s *ListTopicsRequest) SetMessageTypes(v []*string) *ListTopicsRequest {
	s.MessageTypes = v
	return s
}

func (s *ListTopicsRequest) SetPageNumber(v int32) *ListTopicsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTopicsRequest) SetPageSize(v int32) *ListTopicsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTopicsRequest) Validate() error {
	return dara.Validate(s)
}

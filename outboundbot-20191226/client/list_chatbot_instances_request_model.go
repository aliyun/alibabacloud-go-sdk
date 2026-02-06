// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatbotInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListChatbotInstancesRequest
	GetAgentKey() *string
	SetPageNumber(v int32) *ListChatbotInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListChatbotInstancesRequest
	GetPageSize() *int32
}

type ListChatbotInstancesRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListChatbotInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatbotInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListChatbotInstancesRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListChatbotInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListChatbotInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChatbotInstancesRequest) SetAgentKey(v string) *ListChatbotInstancesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetPageNumber(v int32) *ListChatbotInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListChatbotInstancesRequest) SetPageSize(v int32) *ListChatbotInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatbotInstancesRequest) Validate() error {
	return dara.Validate(s)
}

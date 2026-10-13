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
	// The access key of the chatbot workspace.
	//
	// > You can obtain the value by calling the ListChatbotAgents operation. If you leave this parameter empty, chatbots in all workspaces are queried.
	//
	// example:
	//
	// 9137ab9c27044921860030adf8590ec4_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// > The value of this parameter must be less than or equal to 50.
	//
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

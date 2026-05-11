// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConversationDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *ListConversationDetailsRequest
	GetConversationId() *string
	SetInstanceId(v string) *ListConversationDetailsRequest
	GetInstanceId() *string
}

type ListConversationDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a2c26e67-5984-4935-984e-bcee52971993
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 82b2eaae-ce5c-45f8-8231-f15b6b27e55c
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListConversationDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConversationDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListConversationDetailsRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListConversationDetailsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConversationDetailsRequest) SetConversationId(v string) *ListConversationDetailsRequest {
	s.ConversationId = &v
	return s
}

func (s *ListConversationDetailsRequest) SetInstanceId(v string) *ListConversationDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConversationDetailsRequest) Validate() error {
	return dara.Validate(s)
}

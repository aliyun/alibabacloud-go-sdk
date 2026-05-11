// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *DescribeConversationRequest
	GetConversationId() *string
	SetInstanceId(v string) *DescribeConversationRequest
	GetInstanceId() *string
}

type DescribeConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 15608cce-36be-43d5-9361-178cbe64127b
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5daac920-d6c1-429f-a95f-2a798f5255b5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConversationRequest) GoString() string {
	return s.String()
}

func (s *DescribeConversationRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *DescribeConversationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeConversationRequest) SetConversationId(v string) *DescribeConversationRequest {
	s.ConversationId = &v
	return s
}

func (s *DescribeConversationRequest) SetInstanceId(v string) *DescribeConversationRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeConversationRequest) Validate() error {
	return dara.Validate(s)
}

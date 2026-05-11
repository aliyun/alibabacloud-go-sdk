// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConversationContextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *DescribeConversationContextRequest
	GetConversationId() *string
	SetInstanceId(v string) *DescribeConversationContextRequest
	GetInstanceId() *string
}

type DescribeConversationContextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 026ca0f4-483b-4252-ae1d-1f15f056f8b9
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeConversationContextRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConversationContextRequest) GoString() string {
	return s.String()
}

func (s *DescribeConversationContextRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *DescribeConversationContextRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeConversationContextRequest) SetConversationId(v string) *DescribeConversationContextRequest {
	s.ConversationId = &v
	return s
}

func (s *DescribeConversationContextRequest) SetInstanceId(v string) *DescribeConversationContextRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeConversationContextRequest) Validate() error {
	return dara.Validate(s)
}

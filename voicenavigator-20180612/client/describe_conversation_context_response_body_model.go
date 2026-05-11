// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConversationContextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConversationContext(v string) *DescribeConversationContextResponseBody
	GetConversationContext() *string
	SetRequestId(v string) *DescribeConversationContextResponseBody
	GetRequestId() *string
}

type DescribeConversationContextResponseBody struct {
	// example:
	//
	// {         "CallingNumber": "135815***",         "AdditionalContext": "",         "ConversationId": "361c8a53-0e29-42f3-8aa7-c7752d010399"     }
	ConversationContext *string `json:"ConversationContext,omitempty" xml:"ConversationContext,omitempty"`
	// example:
	//
	// b19af5ce5314ac08108d1b33fe20e15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConversationContextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConversationContextResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConversationContextResponseBody) GetConversationContext() *string {
	return s.ConversationContext
}

func (s *DescribeConversationContextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConversationContextResponseBody) SetConversationContext(v string) *DescribeConversationContextResponseBody {
	s.ConversationContext = &v
	return s
}

func (s *DescribeConversationContextResponseBody) SetRequestId(v string) *DescribeConversationContextResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConversationContextResponseBody) Validate() error {
	return dara.Validate(s)
}

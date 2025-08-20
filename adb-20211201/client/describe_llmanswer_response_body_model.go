// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLLMAnswerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeLLMAnswerResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeLLMAnswerResponseBody
	GetRequestId() *string
	SetSessionId(v string) *DescribeLLMAnswerResponseBody
	GetSessionId() *string
}

type DescribeLLMAnswerResponseBody struct {
	// The answer by the intelligent assistant to the question.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 456
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribeLLMAnswerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLLMAnswerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLLMAnswerResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeLLMAnswerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLLMAnswerResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeLLMAnswerResponseBody) SetContent(v string) *DescribeLLMAnswerResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeLLMAnswerResponseBody) SetRequestId(v string) *DescribeLLMAnswerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLLMAnswerResponseBody) SetSessionId(v string) *DescribeLLMAnswerResponseBody {
	s.SessionId = &v
	return s
}

func (s *DescribeLLMAnswerResponseBody) Validate() error {
	return dara.Validate(s)
}

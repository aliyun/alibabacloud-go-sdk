// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextualAnswerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v *Answer) *ContextualAnswerResponseBody
	GetAnswer() *Answer
	SetCode(v string) *ContextualAnswerResponseBody
	GetCode() *string
	SetMessage(v string) *ContextualAnswerResponseBody
	GetMessage() *string
	SetRequestId(v string) *ContextualAnswerResponseBody
	GetRequestId() *string
}

type ContextualAnswerResponseBody struct {
	Answer  *Answer `json:"Answer,omitempty" xml:"Answer,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 22F081FB-90D7-525A-BFE4-D28DC906A28F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContextualAnswerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContextualAnswerResponseBody) GoString() string {
	return s.String()
}

func (s *ContextualAnswerResponseBody) GetAnswer() *Answer {
	return s.Answer
}

func (s *ContextualAnswerResponseBody) GetCode() *string {
	return s.Code
}

func (s *ContextualAnswerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ContextualAnswerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContextualAnswerResponseBody) SetAnswer(v *Answer) *ContextualAnswerResponseBody {
	s.Answer = v
	return s
}

func (s *ContextualAnswerResponseBody) SetCode(v string) *ContextualAnswerResponseBody {
	s.Code = &v
	return s
}

func (s *ContextualAnswerResponseBody) SetMessage(v string) *ContextualAnswerResponseBody {
	s.Message = &v
	return s
}

func (s *ContextualAnswerResponseBody) SetRequestId(v string) *ContextualAnswerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContextualAnswerResponseBody) Validate() error {
	return dara.Validate(s)
}

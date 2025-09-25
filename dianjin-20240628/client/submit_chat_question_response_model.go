// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitChatQuestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitChatQuestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitChatQuestionResponse
	GetStatusCode() *int32
	SetBody(v *SubmitChatQuestionResponseBody) *SubmitChatQuestionResponse
	GetBody() *SubmitChatQuestionResponseBody
}

type SubmitChatQuestionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitChatQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitChatQuestionResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitChatQuestionResponse) GoString() string {
	return s.String()
}

func (s *SubmitChatQuestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitChatQuestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitChatQuestionResponse) GetBody() *SubmitChatQuestionResponseBody {
	return s.Body
}

func (s *SubmitChatQuestionResponse) SetHeaders(v map[string]*string) *SubmitChatQuestionResponse {
	s.Headers = v
	return s
}

func (s *SubmitChatQuestionResponse) SetStatusCode(v int32) *SubmitChatQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitChatQuestionResponse) SetBody(v *SubmitChatQuestionResponseBody) *SubmitChatQuestionResponse {
	s.Body = v
	return s
}

func (s *SubmitChatQuestionResponse) Validate() error {
	return dara.Validate(s)
}

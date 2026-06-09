// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerSSEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnswerSSEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnswerSSEResponse
	GetStatusCode() *int32
	SetId(v string) *AnswerSSEResponse
	GetId() *string
	SetEvent(v string) *AnswerSSEResponse
	GetEvent() *string
	SetBody(v *AnswerSSEResponseBody) *AnswerSSEResponse
	GetBody() *AnswerSSEResponseBody
}

type AnswerSSEResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                `json:"event,omitempty" xml:"event,omitempty"`
	Body       *AnswerSSEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnswerSSEResponse) String() string {
	return dara.Prettify(s)
}

func (s AnswerSSEResponse) GoString() string {
	return s.String()
}

func (s *AnswerSSEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnswerSSEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnswerSSEResponse) GetId() *string {
	return s.Id
}

func (s *AnswerSSEResponse) GetEvent() *string {
	return s.Event
}

func (s *AnswerSSEResponse) GetBody() *AnswerSSEResponseBody {
	return s.Body
}

func (s *AnswerSSEResponse) SetHeaders(v map[string]*string) *AnswerSSEResponse {
	s.Headers = v
	return s
}

func (s *AnswerSSEResponse) SetStatusCode(v int32) *AnswerSSEResponse {
	s.StatusCode = &v
	return s
}

func (s *AnswerSSEResponse) SetId(v string) *AnswerSSEResponse {
	s.Id = &v
	return s
}

func (s *AnswerSSEResponse) SetEvent(v string) *AnswerSSEResponse {
	s.Event = &v
	return s
}

func (s *AnswerSSEResponse) SetBody(v *AnswerSSEResponseBody) *AnswerSSEResponse {
	s.Body = v
	return s
}

func (s *AnswerSSEResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextualAnswerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContextualAnswerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContextualAnswerResponse
	GetStatusCode() *int32
	SetId(v string) *ContextualAnswerResponse
	GetId() *string
	SetEvent(v string) *ContextualAnswerResponse
	GetEvent() *string
	SetBody(v *ContextualAnswerResponseBody) *ContextualAnswerResponse
	GetBody() *ContextualAnswerResponseBody
}

type ContextualAnswerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                       `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                       `json:"event,omitempty" xml:"event,omitempty"`
	Body       *ContextualAnswerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContextualAnswerResponse) String() string {
	return dara.Prettify(s)
}

func (s ContextualAnswerResponse) GoString() string {
	return s.String()
}

func (s *ContextualAnswerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContextualAnswerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContextualAnswerResponse) GetId() *string {
	return s.Id
}

func (s *ContextualAnswerResponse) GetEvent() *string {
	return s.Event
}

func (s *ContextualAnswerResponse) GetBody() *ContextualAnswerResponseBody {
	return s.Body
}

func (s *ContextualAnswerResponse) SetHeaders(v map[string]*string) *ContextualAnswerResponse {
	s.Headers = v
	return s
}

func (s *ContextualAnswerResponse) SetStatusCode(v int32) *ContextualAnswerResponse {
	s.StatusCode = &v
	return s
}

func (s *ContextualAnswerResponse) SetId(v string) *ContextualAnswerResponse {
	s.Id = &v
	return s
}

func (s *ContextualAnswerResponse) SetEvent(v string) *ContextualAnswerResponse {
	s.Event = &v
	return s
}

func (s *ContextualAnswerResponse) SetBody(v *ContextualAnswerResponseBody) *ContextualAnswerResponse {
	s.Body = v
	return s
}

func (s *ContextualAnswerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

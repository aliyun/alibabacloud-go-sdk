// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChatResponse
	GetStatusCode() *int32
	SetId(v string) *CreateChatResponse
	GetId() *string
	SetEvent(v string) *CreateChatResponse
	GetEvent() *string
	SetBody(v *CreateChatResponseBody) *CreateChatResponse
	GetBody() *CreateChatResponseBody
}

type CreateChatResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                 `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                 `json:"event,omitempty" xml:"event,omitempty"`
	Body       *CreateChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChatResponse) GoString() string {
	return s.String()
}

func (s *CreateChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChatResponse) GetId() *string {
	return s.Id
}

func (s *CreateChatResponse) GetEvent() *string {
	return s.Event
}

func (s *CreateChatResponse) GetBody() *CreateChatResponseBody {
	return s.Body
}

func (s *CreateChatResponse) SetHeaders(v map[string]*string) *CreateChatResponse {
	s.Headers = v
	return s
}

func (s *CreateChatResponse) SetStatusCode(v int32) *CreateChatResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatResponse) SetId(v string) *CreateChatResponse {
	s.Id = &v
	return s
}

func (s *CreateChatResponse) SetEvent(v string) *CreateChatResponse {
	s.Event = &v
	return s
}

func (s *CreateChatResponse) SetBody(v *CreateChatResponseBody) *CreateChatResponse {
	s.Body = v
	return s
}

func (s *CreateChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

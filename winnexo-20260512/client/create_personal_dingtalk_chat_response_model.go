// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalDingtalkChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalDingtalkChatResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalDingtalkChatResponseBody) *CreatePersonalDingtalkChatResponse
	GetBody() *CreatePersonalDingtalkChatResponseBody
}

type CreatePersonalDingtalkChatResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalDingtalkChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalDingtalkChatResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkChatResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalDingtalkChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalDingtalkChatResponse) GetBody() *CreatePersonalDingtalkChatResponseBody {
	return s.Body
}

func (s *CreatePersonalDingtalkChatResponse) SetHeaders(v map[string]*string) *CreatePersonalDingtalkChatResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalDingtalkChatResponse) SetStatusCode(v int32) *CreatePersonalDingtalkChatResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalDingtalkChatResponse) SetBody(v *CreatePersonalDingtalkChatResponseBody) *CreatePersonalDingtalkChatResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalDingtalkChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

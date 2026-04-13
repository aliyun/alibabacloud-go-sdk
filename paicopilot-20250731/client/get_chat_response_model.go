// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatResponse
	GetStatusCode() *int32
	SetBody(v *GetChatResponseBody) *GetChatResponse
	GetBody() *GetChatResponseBody
}

type GetChatResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatResponse) GoString() string {
	return s.String()
}

func (s *GetChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatResponse) GetBody() *GetChatResponseBody {
	return s.Body
}

func (s *GetChatResponse) SetHeaders(v map[string]*string) *GetChatResponse {
	s.Headers = v
	return s
}

func (s *GetChatResponse) SetStatusCode(v int32) *GetChatResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatResponse) SetBody(v *GetChatResponseBody) *GetChatResponse {
	s.Body = v
	return s
}

func (s *GetChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

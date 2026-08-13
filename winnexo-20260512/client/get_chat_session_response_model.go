// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatSessionResponse
	GetStatusCode() *int32
	SetBody(v *GetChatSessionResponseBody) *GetChatSessionResponse
	GetBody() *GetChatSessionResponseBody
}

type GetChatSessionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatSessionResponse) GoString() string {
	return s.String()
}

func (s *GetChatSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatSessionResponse) GetBody() *GetChatSessionResponseBody {
	return s.Body
}

func (s *GetChatSessionResponse) SetHeaders(v map[string]*string) *GetChatSessionResponse {
	s.Headers = v
	return s
}

func (s *GetChatSessionResponse) SetStatusCode(v int32) *GetChatSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatSessionResponse) SetBody(v *GetChatSessionResponseBody) *GetChatSessionResponse {
	s.Body = v
	return s
}

func (s *GetChatSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChatSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChatSessionResponse
	GetStatusCode() *int32
	SetBody(v *CreateChatSessionResponseBody) *CreateChatSessionResponse
	GetBody() *CreateChatSessionResponseBody
}

type CreateChatSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChatSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChatSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateChatSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChatSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChatSessionResponse) GetBody() *CreateChatSessionResponseBody {
	return s.Body
}

func (s *CreateChatSessionResponse) SetHeaders(v map[string]*string) *CreateChatSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateChatSessionResponse) SetStatusCode(v int32) *CreateChatSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChatSessionResponse) SetBody(v *CreateChatSessionResponseBody) *CreateChatSessionResponse {
	s.Body = v
	return s
}

func (s *CreateChatSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

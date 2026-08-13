// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateChatSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateChatSessionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateChatSessionResponseBody) *UpdateChatSessionResponse
	GetBody() *UpdateChatSessionResponseBody
}

type UpdateChatSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChatSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChatSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatSessionResponse) GoString() string {
	return s.String()
}

func (s *UpdateChatSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateChatSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateChatSessionResponse) GetBody() *UpdateChatSessionResponseBody {
	return s.Body
}

func (s *UpdateChatSessionResponse) SetHeaders(v map[string]*string) *UpdateChatSessionResponse {
	s.Headers = v
	return s
}

func (s *UpdateChatSessionResponse) SetStatusCode(v int32) *UpdateChatSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChatSessionResponse) SetBody(v *UpdateChatSessionResponseBody) *UpdateChatSessionResponse {
	s.Body = v
	return s
}

func (s *UpdateChatSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

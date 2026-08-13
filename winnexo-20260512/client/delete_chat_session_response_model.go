// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChatSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChatSessionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChatSessionResponseBody) *DeleteChatSessionResponse
	GetBody() *DeleteChatSessionResponseBody
}

type DeleteChatSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChatSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChatSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatSessionResponse) GoString() string {
	return s.String()
}

func (s *DeleteChatSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChatSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChatSessionResponse) GetBody() *DeleteChatSessionResponseBody {
	return s.Body
}

func (s *DeleteChatSessionResponse) SetHeaders(v map[string]*string) *DeleteChatSessionResponse {
	s.Headers = v
	return s
}

func (s *DeleteChatSessionResponse) SetStatusCode(v int32) *DeleteChatSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChatSessionResponse) SetBody(v *DeleteChatSessionResponseBody) *DeleteChatSessionResponse {
	s.Body = v
	return s
}

func (s *DeleteChatSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

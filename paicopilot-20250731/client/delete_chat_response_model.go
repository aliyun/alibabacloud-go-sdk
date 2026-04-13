// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChatResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChatResponseBody) *DeleteChatResponse
	GetBody() *DeleteChatResponseBody
}

type DeleteChatResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChatResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatResponse) GoString() string {
	return s.String()
}

func (s *DeleteChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChatResponse) GetBody() *DeleteChatResponseBody {
	return s.Body
}

func (s *DeleteChatResponse) SetHeaders(v map[string]*string) *DeleteChatResponse {
	s.Headers = v
	return s
}

func (s *DeleteChatResponse) SetStatusCode(v int32) *DeleteChatResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChatResponse) SetBody(v *DeleteChatResponseBody) *DeleteChatResponse {
	s.Body = v
	return s
}

func (s *DeleteChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

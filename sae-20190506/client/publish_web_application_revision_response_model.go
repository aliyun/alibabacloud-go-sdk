// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishWebApplicationRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishWebApplicationRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishWebApplicationRevisionResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationRevisionBody) *PublishWebApplicationRevisionResponse
	GetBody() *WebApplicationRevisionBody
}

type PublishWebApplicationRevisionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationRevisionBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishWebApplicationRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishWebApplicationRevisionResponse) GoString() string {
	return s.String()
}

func (s *PublishWebApplicationRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishWebApplicationRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishWebApplicationRevisionResponse) GetBody() *WebApplicationRevisionBody {
	return s.Body
}

func (s *PublishWebApplicationRevisionResponse) SetHeaders(v map[string]*string) *PublishWebApplicationRevisionResponse {
	s.Headers = v
	return s
}

func (s *PublishWebApplicationRevisionResponse) SetStatusCode(v int32) *PublishWebApplicationRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishWebApplicationRevisionResponse) SetBody(v *WebApplicationRevisionBody) *PublishWebApplicationRevisionResponse {
	s.Body = v
	return s
}

func (s *PublishWebApplicationRevisionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

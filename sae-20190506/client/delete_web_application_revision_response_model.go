// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebApplicationRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebApplicationRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebApplicationRevisionResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationRevisionBody) *DeleteWebApplicationRevisionResponse
	GetBody() *WebApplicationRevisionBody
}

type DeleteWebApplicationRevisionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationRevisionBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebApplicationRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebApplicationRevisionResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebApplicationRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebApplicationRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebApplicationRevisionResponse) GetBody() *WebApplicationRevisionBody {
	return s.Body
}

func (s *DeleteWebApplicationRevisionResponse) SetHeaders(v map[string]*string) *DeleteWebApplicationRevisionResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebApplicationRevisionResponse) SetStatusCode(v int32) *DeleteWebApplicationRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebApplicationRevisionResponse) SetBody(v *WebApplicationRevisionBody) *DeleteWebApplicationRevisionResponse {
	s.Body = v
	return s
}

func (s *DeleteWebApplicationRevisionResponse) Validate() error {
	return dara.Validate(s)
}

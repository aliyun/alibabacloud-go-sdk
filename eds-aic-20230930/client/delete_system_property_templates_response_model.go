// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSystemPropertyTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSystemPropertyTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSystemPropertyTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSystemPropertyTemplatesResponseBody) *DeleteSystemPropertyTemplatesResponse
	GetBody() *DeleteSystemPropertyTemplatesResponseBody
}

type DeleteSystemPropertyTemplatesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSystemPropertyTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSystemPropertyTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSystemPropertyTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteSystemPropertyTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSystemPropertyTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSystemPropertyTemplatesResponse) GetBody() *DeleteSystemPropertyTemplatesResponseBody {
	return s.Body
}

func (s *DeleteSystemPropertyTemplatesResponse) SetHeaders(v map[string]*string) *DeleteSystemPropertyTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteSystemPropertyTemplatesResponse) SetStatusCode(v int32) *DeleteSystemPropertyTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSystemPropertyTemplatesResponse) SetBody(v *DeleteSystemPropertyTemplatesResponseBody) *DeleteSystemPropertyTemplatesResponse {
	s.Body = v
	return s
}

func (s *DeleteSystemPropertyTemplatesResponse) Validate() error {
	return dara.Validate(s)
}

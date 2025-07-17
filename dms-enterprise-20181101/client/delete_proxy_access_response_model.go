// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProxyAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProxyAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProxyAccessResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProxyAccessResponseBody) *DeleteProxyAccessResponse
	GetBody() *DeleteProxyAccessResponseBody
}

type DeleteProxyAccessResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProxyAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProxyAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProxyAccessResponse) GoString() string {
	return s.String()
}

func (s *DeleteProxyAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProxyAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProxyAccessResponse) GetBody() *DeleteProxyAccessResponseBody {
	return s.Body
}

func (s *DeleteProxyAccessResponse) SetHeaders(v map[string]*string) *DeleteProxyAccessResponse {
	s.Headers = v
	return s
}

func (s *DeleteProxyAccessResponse) SetStatusCode(v int32) *DeleteProxyAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProxyAccessResponse) SetBody(v *DeleteProxyAccessResponseBody) *DeleteProxyAccessResponse {
	s.Body = v
	return s
}

func (s *DeleteProxyAccessResponse) Validate() error {
	return dara.Validate(s)
}

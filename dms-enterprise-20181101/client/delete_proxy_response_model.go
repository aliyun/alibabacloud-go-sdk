// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProxyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProxyResponseBody) *DeleteProxyResponse
	GetBody() *DeleteProxyResponseBody
}

type DeleteProxyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProxyResponse) GoString() string {
	return s.String()
}

func (s *DeleteProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProxyResponse) GetBody() *DeleteProxyResponseBody {
	return s.Body
}

func (s *DeleteProxyResponse) SetHeaders(v map[string]*string) *DeleteProxyResponse {
	s.Headers = v
	return s
}

func (s *DeleteProxyResponse) SetStatusCode(v int32) *DeleteProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProxyResponse) SetBody(v *DeleteProxyResponseBody) *DeleteProxyResponse {
	s.Body = v
	return s
}

func (s *DeleteProxyResponse) Validate() error {
	return dara.Validate(s)
}

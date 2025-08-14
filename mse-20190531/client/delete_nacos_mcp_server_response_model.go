// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNacosMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNacosMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNacosMcpServerResponseBody) *DeleteNacosMcpServerResponse
	GetBody() *DeleteNacosMcpServerResponseBody
}

type DeleteNacosMcpServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNacosMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNacosMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosMcpServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteNacosMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNacosMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNacosMcpServerResponse) GetBody() *DeleteNacosMcpServerResponseBody {
	return s.Body
}

func (s *DeleteNacosMcpServerResponse) SetHeaders(v map[string]*string) *DeleteNacosMcpServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteNacosMcpServerResponse) SetStatusCode(v int32) *DeleteNacosMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNacosMcpServerResponse) SetBody(v *DeleteNacosMcpServerResponseBody) *DeleteNacosMcpServerResponse {
	s.Body = v
	return s
}

func (s *DeleteNacosMcpServerResponse) Validate() error {
	return dara.Validate(s)
}

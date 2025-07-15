// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDhcpOptionsSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDhcpOptionsSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDhcpOptionsSetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDhcpOptionsSetResponseBody) *DeleteDhcpOptionsSetResponse
	GetBody() *DeleteDhcpOptionsSetResponseBody
}

type DeleteDhcpOptionsSetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDhcpOptionsSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDhcpOptionsSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDhcpOptionsSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDhcpOptionsSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDhcpOptionsSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDhcpOptionsSetResponse) GetBody() *DeleteDhcpOptionsSetResponseBody {
	return s.Body
}

func (s *DeleteDhcpOptionsSetResponse) SetHeaders(v map[string]*string) *DeleteDhcpOptionsSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDhcpOptionsSetResponse) SetStatusCode(v int32) *DeleteDhcpOptionsSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDhcpOptionsSetResponse) SetBody(v *DeleteDhcpOptionsSetResponseBody) *DeleteDhcpOptionsSetResponse {
	s.Body = v
	return s
}

func (s *DeleteDhcpOptionsSetResponse) Validate() error {
	return dara.Validate(s)
}

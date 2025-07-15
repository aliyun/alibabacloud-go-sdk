// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatIpCidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNatIpCidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNatIpCidrResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNatIpCidrResponseBody) *DeleteNatIpCidrResponse
	GetBody() *DeleteNatIpCidrResponseBody
}

type DeleteNatIpCidrResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNatIpCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNatIpCidrResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatIpCidrResponse) GoString() string {
	return s.String()
}

func (s *DeleteNatIpCidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNatIpCidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNatIpCidrResponse) GetBody() *DeleteNatIpCidrResponseBody {
	return s.Body
}

func (s *DeleteNatIpCidrResponse) SetHeaders(v map[string]*string) *DeleteNatIpCidrResponse {
	s.Headers = v
	return s
}

func (s *DeleteNatIpCidrResponse) SetStatusCode(v int32) *DeleteNatIpCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNatIpCidrResponse) SetBody(v *DeleteNatIpCidrResponseBody) *DeleteNatIpCidrResponse {
	s.Body = v
	return s
}

func (s *DeleteNatIpCidrResponse) Validate() error {
	return dara.Validate(s)
}

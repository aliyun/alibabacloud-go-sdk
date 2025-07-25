// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamPoolCidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpamPoolCidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpamPoolCidrResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpamPoolCidrResponseBody) *DeleteIpamPoolCidrResponse
	GetBody() *DeleteIpamPoolCidrResponseBody
}

type DeleteIpamPoolCidrResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpamPoolCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpamPoolCidrResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamPoolCidrResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolCidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpamPoolCidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpamPoolCidrResponse) GetBody() *DeleteIpamPoolCidrResponseBody {
	return s.Body
}

func (s *DeleteIpamPoolCidrResponse) SetHeaders(v map[string]*string) *DeleteIpamPoolCidrResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpamPoolCidrResponse) SetStatusCode(v int32) *DeleteIpamPoolCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpamPoolCidrResponse) SetBody(v *DeleteIpamPoolCidrResponseBody) *DeleteIpamPoolCidrResponse {
	s.Body = v
	return s
}

func (s *DeleteIpamPoolCidrResponse) Validate() error {
	return dara.Validate(s)
}

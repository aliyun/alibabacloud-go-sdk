// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicIpAddressPoolCidrBlocksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublicIpAddressPoolCidrBlocksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublicIpAddressPoolCidrBlocksResponse
	GetStatusCode() *int32
	SetBody(v *ListPublicIpAddressPoolCidrBlocksResponseBody) *ListPublicIpAddressPoolCidrBlocksResponse
	GetBody() *ListPublicIpAddressPoolCidrBlocksResponseBody
}

type ListPublicIpAddressPoolCidrBlocksResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicIpAddressPoolCidrBlocksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicIpAddressPoolCidrBlocksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublicIpAddressPoolCidrBlocksResponse) GoString() string {
	return s.String()
}

func (s *ListPublicIpAddressPoolCidrBlocksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublicIpAddressPoolCidrBlocksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublicIpAddressPoolCidrBlocksResponse) GetBody() *ListPublicIpAddressPoolCidrBlocksResponseBody {
	return s.Body
}

func (s *ListPublicIpAddressPoolCidrBlocksResponse) SetHeaders(v map[string]*string) *ListPublicIpAddressPoolCidrBlocksResponse {
	s.Headers = v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponse) SetStatusCode(v int32) *ListPublicIpAddressPoolCidrBlocksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponse) SetBody(v *ListPublicIpAddressPoolCidrBlocksResponseBody) *ListPublicIpAddressPoolCidrBlocksResponse {
	s.Body = v
	return s
}

func (s *ListPublicIpAddressPoolCidrBlocksResponse) Validate() error {
	return dara.Validate(s)
}

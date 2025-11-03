// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicIpAddressPoolCidrBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePublicIpAddressPoolCidrBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePublicIpAddressPoolCidrBlockResponse
	GetStatusCode() *int32
	SetBody(v *DeletePublicIpAddressPoolCidrBlockResponseBody) *DeletePublicIpAddressPoolCidrBlockResponse
	GetBody() *DeletePublicIpAddressPoolCidrBlockResponseBody
}

type DeletePublicIpAddressPoolCidrBlockResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePublicIpAddressPoolCidrBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePublicIpAddressPoolCidrBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicIpAddressPoolCidrBlockResponse) GoString() string {
	return s.String()
}

func (s *DeletePublicIpAddressPoolCidrBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePublicIpAddressPoolCidrBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePublicIpAddressPoolCidrBlockResponse) GetBody() *DeletePublicIpAddressPoolCidrBlockResponseBody {
	return s.Body
}

func (s *DeletePublicIpAddressPoolCidrBlockResponse) SetHeaders(v map[string]*string) *DeletePublicIpAddressPoolCidrBlockResponse {
	s.Headers = v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockResponse) SetStatusCode(v int32) *DeletePublicIpAddressPoolCidrBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockResponse) SetBody(v *DeletePublicIpAddressPoolCidrBlockResponseBody) *DeletePublicIpAddressPoolCidrBlockResponse {
	s.Body = v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

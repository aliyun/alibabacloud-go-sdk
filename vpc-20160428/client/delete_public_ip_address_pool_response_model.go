// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicIpAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePublicIpAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePublicIpAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *DeletePublicIpAddressPoolResponseBody) *DeletePublicIpAddressPoolResponse
	GetBody() *DeletePublicIpAddressPoolResponseBody
}

type DeletePublicIpAddressPoolResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePublicIpAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePublicIpAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicIpAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *DeletePublicIpAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePublicIpAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePublicIpAddressPoolResponse) GetBody() *DeletePublicIpAddressPoolResponseBody {
	return s.Body
}

func (s *DeletePublicIpAddressPoolResponse) SetHeaders(v map[string]*string) *DeletePublicIpAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *DeletePublicIpAddressPoolResponse) SetStatusCode(v int32) *DeletePublicIpAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePublicIpAddressPoolResponse) SetBody(v *DeletePublicIpAddressPoolResponseBody) *DeletePublicIpAddressPoolResponse {
	s.Body = v
	return s
}

func (s *DeletePublicIpAddressPoolResponse) Validate() error {
	return dara.Validate(s)
}

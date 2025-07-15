// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublicIpAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePublicIpAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePublicIpAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *CreatePublicIpAddressPoolResponseBody) *CreatePublicIpAddressPoolResponse
	GetBody() *CreatePublicIpAddressPoolResponseBody
}

type CreatePublicIpAddressPoolResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePublicIpAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePublicIpAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePublicIpAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *CreatePublicIpAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePublicIpAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePublicIpAddressPoolResponse) GetBody() *CreatePublicIpAddressPoolResponseBody {
	return s.Body
}

func (s *CreatePublicIpAddressPoolResponse) SetHeaders(v map[string]*string) *CreatePublicIpAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *CreatePublicIpAddressPoolResponse) SetStatusCode(v int32) *CreatePublicIpAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePublicIpAddressPoolResponse) SetBody(v *CreatePublicIpAddressPoolResponseBody) *CreatePublicIpAddressPoolResponse {
	s.Body = v
	return s
}

func (s *CreatePublicIpAddressPoolResponse) Validate() error {
	return dara.Validate(s)
}

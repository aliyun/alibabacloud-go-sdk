// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPublicIpAddressPoolServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenPublicIpAddressPoolServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenPublicIpAddressPoolServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenPublicIpAddressPoolServiceResponseBody) *OpenPublicIpAddressPoolServiceResponse
	GetBody() *OpenPublicIpAddressPoolServiceResponseBody
}

type OpenPublicIpAddressPoolServiceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenPublicIpAddressPoolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenPublicIpAddressPoolServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenPublicIpAddressPoolServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenPublicIpAddressPoolServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenPublicIpAddressPoolServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenPublicIpAddressPoolServiceResponse) GetBody() *OpenPublicIpAddressPoolServiceResponseBody {
	return s.Body
}

func (s *OpenPublicIpAddressPoolServiceResponse) SetHeaders(v map[string]*string) *OpenPublicIpAddressPoolServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenPublicIpAddressPoolServiceResponse) SetStatusCode(v int32) *OpenPublicIpAddressPoolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenPublicIpAddressPoolServiceResponse) SetBody(v *OpenPublicIpAddressPoolServiceResponseBody) *OpenPublicIpAddressPoolServiceResponse {
	s.Body = v
	return s
}

func (s *OpenPublicIpAddressPoolServiceResponse) Validate() error {
	return dara.Validate(s)
}

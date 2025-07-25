// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenVpcIpamServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenVpcIpamServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenVpcIpamServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenVpcIpamServiceResponseBody) *OpenVpcIpamServiceResponse
	GetBody() *OpenVpcIpamServiceResponseBody
}

type OpenVpcIpamServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenVpcIpamServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenVpcIpamServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenVpcIpamServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenVpcIpamServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenVpcIpamServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenVpcIpamServiceResponse) GetBody() *OpenVpcIpamServiceResponseBody {
	return s.Body
}

func (s *OpenVpcIpamServiceResponse) SetHeaders(v map[string]*string) *OpenVpcIpamServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenVpcIpamServiceResponse) SetStatusCode(v int32) *OpenVpcIpamServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenVpcIpamServiceResponse) SetBody(v *OpenVpcIpamServiceResponseBody) *OpenVpcIpamServiceResponse {
	s.Body = v
	return s
}

func (s *OpenVpcIpamServiceResponse) Validate() error {
	return dara.Validate(s)
}

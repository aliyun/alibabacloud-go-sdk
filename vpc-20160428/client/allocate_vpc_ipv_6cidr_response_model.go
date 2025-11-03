// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateVpcIpv6CidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateVpcIpv6CidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateVpcIpv6CidrResponse
	GetStatusCode() *int32
	SetBody(v *AllocateVpcIpv6CidrResponseBody) *AllocateVpcIpv6CidrResponse
	GetBody() *AllocateVpcIpv6CidrResponseBody
}

type AllocateVpcIpv6CidrResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateVpcIpv6CidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateVpcIpv6CidrResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateVpcIpv6CidrResponse) GoString() string {
	return s.String()
}

func (s *AllocateVpcIpv6CidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateVpcIpv6CidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateVpcIpv6CidrResponse) GetBody() *AllocateVpcIpv6CidrResponseBody {
	return s.Body
}

func (s *AllocateVpcIpv6CidrResponse) SetHeaders(v map[string]*string) *AllocateVpcIpv6CidrResponse {
	s.Headers = v
	return s
}

func (s *AllocateVpcIpv6CidrResponse) SetStatusCode(v int32) *AllocateVpcIpv6CidrResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateVpcIpv6CidrResponse) SetBody(v *AllocateVpcIpv6CidrResponseBody) *AllocateVpcIpv6CidrResponse {
	s.Body = v
	return s
}

func (s *AllocateVpcIpv6CidrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

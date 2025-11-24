// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIstioGatewayDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIstioGatewayDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIstioGatewayDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIstioGatewayDomainsResponseBody) *DescribeIstioGatewayDomainsResponse
	GetBody() *DescribeIstioGatewayDomainsResponseBody
}

type DescribeIstioGatewayDomainsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIstioGatewayDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIstioGatewayDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIstioGatewayDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIstioGatewayDomainsResponse) GetBody() *DescribeIstioGatewayDomainsResponseBody {
	return s.Body
}

func (s *DescribeIstioGatewayDomainsResponse) SetHeaders(v map[string]*string) *DescribeIstioGatewayDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIstioGatewayDomainsResponse) SetStatusCode(v int32) *DescribeIstioGatewayDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIstioGatewayDomainsResponse) SetBody(v *DescribeIstioGatewayDomainsResponseBody) *DescribeIstioGatewayDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeIstioGatewayDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

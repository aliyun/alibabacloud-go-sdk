// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDNSSLBSubDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDNSSLBSubDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDNSSLBSubDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDNSSLBSubDomainsResponseBody) *DescribeDNSSLBSubDomainsResponse
	GetBody() *DescribeDNSSLBSubDomainsResponseBody
}

type DescribeDNSSLBSubDomainsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDNSSLBSubDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDNSSLBSubDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDNSSLBSubDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDNSSLBSubDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDNSSLBSubDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDNSSLBSubDomainsResponse) GetBody() *DescribeDNSSLBSubDomainsResponseBody {
	return s.Body
}

func (s *DescribeDNSSLBSubDomainsResponse) SetHeaders(v map[string]*string) *DescribeDNSSLBSubDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponse) SetStatusCode(v int32) *DescribeDNSSLBSubDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponse) SetBody(v *DescribeDNSSLBSubDomainsResponseBody) *DescribeDNSSLBSubDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

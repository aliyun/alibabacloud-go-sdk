// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainUsedPortsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainUsedPortsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainUsedPortsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainUsedPortsResponseBody) *DescribeDomainUsedPortsResponse
	GetBody() *DescribeDomainUsedPortsResponseBody
}

type DescribeDomainUsedPortsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainUsedPortsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainUsedPortsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUsedPortsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsedPortsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainUsedPortsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainUsedPortsResponse) GetBody() *DescribeDomainUsedPortsResponseBody {
	return s.Body
}

func (s *DescribeDomainUsedPortsResponse) SetHeaders(v map[string]*string) *DescribeDomainUsedPortsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainUsedPortsResponse) SetStatusCode(v int32) *DescribeDomainUsedPortsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainUsedPortsResponse) SetBody(v *DescribeDomainUsedPortsResponseBody) *DescribeDomainUsedPortsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainUsedPortsResponse) Validate() error {
	return dara.Validate(s)
}

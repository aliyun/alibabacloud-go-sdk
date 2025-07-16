// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAverageResponseTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainAverageResponseTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainAverageResponseTimeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainAverageResponseTimeResponseBody) *DescribeDomainAverageResponseTimeResponse
	GetBody() *DescribeDomainAverageResponseTimeResponseBody
}

type DescribeDomainAverageResponseTimeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainAverageResponseTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainAverageResponseTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAverageResponseTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAverageResponseTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainAverageResponseTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainAverageResponseTimeResponse) GetBody() *DescribeDomainAverageResponseTimeResponseBody {
	return s.Body
}

func (s *DescribeDomainAverageResponseTimeResponse) SetHeaders(v map[string]*string) *DescribeDomainAverageResponseTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponse) SetStatusCode(v int32) *DescribeDomainAverageResponseTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponse) SetBody(v *DescribeDomainAverageResponseTimeResponseBody) *DescribeDomainAverageResponseTimeResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponse) Validate() error {
	return dara.Validate(s)
}

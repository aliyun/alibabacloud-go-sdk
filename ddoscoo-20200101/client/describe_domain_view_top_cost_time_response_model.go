// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewTopCostTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainViewTopCostTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainViewTopCostTimeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainViewTopCostTimeResponseBody) *DescribeDomainViewTopCostTimeResponse
	GetBody() *DescribeDomainViewTopCostTimeResponseBody
}

type DescribeDomainViewTopCostTimeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainViewTopCostTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainViewTopCostTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewTopCostTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopCostTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainViewTopCostTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainViewTopCostTimeResponse) GetBody() *DescribeDomainViewTopCostTimeResponseBody {
	return s.Body
}

func (s *DescribeDomainViewTopCostTimeResponse) SetHeaders(v map[string]*string) *DescribeDomainViewTopCostTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponse) SetStatusCode(v int32) *DescribeDomainViewTopCostTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponse) SetBody(v *DescribeDomainViewTopCostTimeResponseBody) *DescribeDomainViewTopCostTimeResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponse) Validate() error {
	return dara.Validate(s)
}

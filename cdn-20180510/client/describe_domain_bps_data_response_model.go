// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainBpsDataResponseBody) *DescribeDomainBpsDataResponse
	GetBody() *DescribeDomainBpsDataResponseBody
}

type DescribeDomainBpsDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainBpsDataResponse) GetBody() *DescribeDomainBpsDataResponseBody {
	return s.Body
}

func (s *DescribeDomainBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainBpsDataResponse) SetStatusCode(v int32) *DescribeDomainBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainBpsDataResponse) SetBody(v *DescribeDomainBpsDataResponseBody) *DescribeDomainBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainBpsDataResponse) Validate() error {
	return dara.Validate(s)
}

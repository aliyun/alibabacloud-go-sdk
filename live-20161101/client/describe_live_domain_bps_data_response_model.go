// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainBpsDataResponseBody) *DescribeLiveDomainBpsDataResponse
	GetBody() *DescribeLiveDomainBpsDataResponseBody
}

type DescribeLiveDomainBpsDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainBpsDataResponse) GetBody() *DescribeLiveDomainBpsDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainBpsDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainBpsDataResponse) SetStatusCode(v int32) *DescribeLiveDomainBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainBpsDataResponse) SetBody(v *DescribeLiveDomainBpsDataResponseBody) *DescribeLiveDomainBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainBpsDataResponse) Validate() error {
	return dara.Validate(s)
}

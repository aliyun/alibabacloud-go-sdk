// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSrcBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSrcBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSrcBpsDataResponseBody) *DescribeDomainSrcBpsDataResponse
	GetBody() *DescribeDomainSrcBpsDataResponseBody
}

type DescribeDomainSrcBpsDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSrcBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSrcBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSrcBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSrcBpsDataResponse) GetBody() *DescribeDomainSrcBpsDataResponseBody {
	return s.Body
}

func (s *DescribeDomainSrcBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcBpsDataResponse) SetStatusCode(v int32) *DescribeDomainSrcBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponse) SetBody(v *DescribeDomainSrcBpsDataResponseBody) *DescribeDomainSrcBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSrcBpsDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

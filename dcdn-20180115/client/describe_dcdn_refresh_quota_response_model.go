// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRefreshQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnRefreshQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnRefreshQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnRefreshQuotaResponseBody) *DescribeDcdnRefreshQuotaResponse
	GetBody() *DescribeDcdnRefreshQuotaResponseBody
}

type DescribeDcdnRefreshQuotaResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnRefreshQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnRefreshQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnRefreshQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnRefreshQuotaResponse) GetBody() *DescribeDcdnRefreshQuotaResponseBody {
	return s.Body
}

func (s *DescribeDcdnRefreshQuotaResponse) SetHeaders(v map[string]*string) *DescribeDcdnRefreshQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponse) SetStatusCode(v int32) *DescribeDcdnRefreshQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponse) SetBody(v *DescribeDcdnRefreshQuotaResponseBody) *DescribeDcdnRefreshQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnRefreshQuotaResponse) Validate() error {
	return dara.Validate(s)
}

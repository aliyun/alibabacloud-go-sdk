// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserQuotaResponseBody) *DescribeDcdnUserQuotaResponse
	GetBody() *DescribeDcdnUserQuotaResponseBody
}

type DescribeDcdnUserQuotaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserQuotaResponse) GetBody() *DescribeDcdnUserQuotaResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserQuotaResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserQuotaResponse) SetStatusCode(v int32) *DescribeDcdnUserQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserQuotaResponse) SetBody(v *DescribeDcdnUserQuotaResponseBody) *DescribeDcdnUserQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserQuotaResponse) Validate() error {
	return dara.Validate(s)
}

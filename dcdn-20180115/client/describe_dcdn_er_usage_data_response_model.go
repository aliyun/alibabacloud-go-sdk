// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnErUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnErUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnErUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnErUsageDataResponseBody) *DescribeDcdnErUsageDataResponse
	GetBody() *DescribeDcdnErUsageDataResponseBody
}

type DescribeDcdnErUsageDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnErUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnErUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnErUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnErUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnErUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnErUsageDataResponse) GetBody() *DescribeDcdnErUsageDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnErUsageDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnErUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnErUsageDataResponse) SetStatusCode(v int32) *DescribeDcdnErUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnErUsageDataResponse) SetBody(v *DescribeDcdnErUsageDataResponseBody) *DescribeDcdnErUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnErUsageDataResponse) Validate() error {
	return dara.Validate(s)
}

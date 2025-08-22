// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnBlockedRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnBlockedRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnBlockedRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnBlockedRegionsResponseBody) *DescribeDcdnBlockedRegionsResponse
	GetBody() *DescribeDcdnBlockedRegionsResponseBody
}

type DescribeDcdnBlockedRegionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnBlockedRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnBlockedRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBlockedRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBlockedRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnBlockedRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnBlockedRegionsResponse) GetBody() *DescribeDcdnBlockedRegionsResponseBody {
	return s.Body
}

func (s *DescribeDcdnBlockedRegionsResponse) SetHeaders(v map[string]*string) *DescribeDcdnBlockedRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponse) SetStatusCode(v int32) *DescribeDcdnBlockedRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponse) SetBody(v *DescribeDcdnBlockedRegionsResponseBody) *DescribeDcdnBlockedRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponse) Validate() error {
	return dara.Validate(s)
}

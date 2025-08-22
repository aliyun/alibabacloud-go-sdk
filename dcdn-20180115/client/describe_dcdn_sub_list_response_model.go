// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSubListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnSubListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnSubListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnSubListResponseBody) *DescribeDcdnSubListResponse
	GetBody() *DescribeDcdnSubListResponseBody
}

type DescribeDcdnSubListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnSubListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnSubListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSubListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSubListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnSubListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnSubListResponse) GetBody() *DescribeDcdnSubListResponseBody {
	return s.Body
}

func (s *DescribeDcdnSubListResponse) SetHeaders(v map[string]*string) *DescribeDcdnSubListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSubListResponse) SetStatusCode(v int32) *DescribeDcdnSubListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnSubListResponse) SetBody(v *DescribeDcdnSubListResponseBody) *DescribeDcdnSubListResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnSubListResponse) Validate() error {
	return dara.Validate(s)
}

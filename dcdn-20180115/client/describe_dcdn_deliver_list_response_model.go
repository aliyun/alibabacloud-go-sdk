// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDeliverListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDeliverListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDeliverListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDeliverListResponseBody) *DescribeDcdnDeliverListResponse
	GetBody() *DescribeDcdnDeliverListResponseBody
}

type DescribeDcdnDeliverListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDeliverListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDeliverListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDeliverListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeliverListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDeliverListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDeliverListResponse) GetBody() *DescribeDcdnDeliverListResponseBody {
	return s.Body
}

func (s *DescribeDcdnDeliverListResponse) SetHeaders(v map[string]*string) *DescribeDcdnDeliverListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDeliverListResponse) SetStatusCode(v int32) *DescribeDcdnDeliverListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDeliverListResponse) SetBody(v *DescribeDcdnDeliverListResponseBody) *DescribeDcdnDeliverListResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDeliverListResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsShardingDbsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsShardingDbsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsShardingDbsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsShardingDbsResponseBody) *DescribeDrdsShardingDbsResponse
	GetBody() *DescribeDrdsShardingDbsResponseBody
}

type DescribeDrdsShardingDbsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsShardingDbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsShardingDbsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsShardingDbsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsShardingDbsResponse) GetBody() *DescribeDrdsShardingDbsResponseBody {
	return s.Body
}

func (s *DescribeDrdsShardingDbsResponse) SetHeaders(v map[string]*string) *DescribeDrdsShardingDbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsShardingDbsResponse) SetStatusCode(v int32) *DescribeDrdsShardingDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponse) SetBody(v *DescribeDrdsShardingDbsResponseBody) *DescribeDrdsShardingDbsResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsShardingDbsResponse) Validate() error {
	return dara.Validate(s)
}

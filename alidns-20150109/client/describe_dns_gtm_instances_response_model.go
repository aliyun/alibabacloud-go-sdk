// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmInstancesResponseBody) *DescribeDnsGtmInstancesResponse
	GetBody() *DescribeDnsGtmInstancesResponseBody
}

type DescribeDnsGtmInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmInstancesResponse) GetBody() *DescribeDnsGtmInstancesResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmInstancesResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstancesResponse) SetStatusCode(v int32) *DescribeDnsGtmInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponse) SetBody(v *DescribeDnsGtmInstancesResponseBody) *DescribeDnsGtmInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmInstancesResponse) Validate() error {
	return dara.Validate(s)
}

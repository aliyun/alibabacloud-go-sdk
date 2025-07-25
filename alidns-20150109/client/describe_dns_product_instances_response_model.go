// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsProductInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsProductInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsProductInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsProductInstancesResponseBody) *DescribeDnsProductInstancesResponse
	GetBody() *DescribeDnsProductInstancesResponseBody
}

type DescribeDnsProductInstancesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsProductInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsProductInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsProductInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsProductInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsProductInstancesResponse) GetBody() *DescribeDnsProductInstancesResponseBody {
	return s.Body
}

func (s *DescribeDnsProductInstancesResponse) SetHeaders(v map[string]*string) *DescribeDnsProductInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsProductInstancesResponse) SetStatusCode(v int32) *DescribeDnsProductInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsProductInstancesResponse) SetBody(v *DescribeDnsProductInstancesResponseBody) *DescribeDnsProductInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsProductInstancesResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAccessStrategiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmAccessStrategiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmAccessStrategiesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmAccessStrategiesResponseBody) *DescribeGtmAccessStrategiesResponse
	GetBody() *DescribeGtmAccessStrategiesResponseBody
}

type DescribeGtmAccessStrategiesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmAccessStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmAccessStrategiesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmAccessStrategiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmAccessStrategiesResponse) GetBody() *DescribeGtmAccessStrategiesResponseBody {
	return s.Body
}

func (s *DescribeGtmAccessStrategiesResponse) SetHeaders(v map[string]*string) *DescribeGtmAccessStrategiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmAccessStrategiesResponse) SetStatusCode(v int32) *DescribeGtmAccessStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmAccessStrategiesResponse) SetBody(v *DescribeGtmAccessStrategiesResponseBody) *DescribeGtmAccessStrategiesResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmAccessStrategiesResponse) Validate() error {
	return dara.Validate(s)
}

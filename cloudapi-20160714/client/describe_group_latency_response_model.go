// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupLatencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupLatencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupLatencyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupLatencyResponseBody) *DescribeGroupLatencyResponse
	GetBody() *DescribeGroupLatencyResponseBody
}

type DescribeGroupLatencyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupLatencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupLatencyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupLatencyResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupLatencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupLatencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupLatencyResponse) GetBody() *DescribeGroupLatencyResponseBody {
	return s.Body
}

func (s *DescribeGroupLatencyResponse) SetHeaders(v map[string]*string) *DescribeGroupLatencyResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupLatencyResponse) SetStatusCode(v int32) *DescribeGroupLatencyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupLatencyResponse) SetBody(v *DescribeGroupLatencyResponseBody) *DescribeGroupLatencyResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupLatencyResponse) Validate() error {
	return dara.Validate(s)
}

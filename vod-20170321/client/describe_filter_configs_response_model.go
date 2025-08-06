// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilterConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFilterConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFilterConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFilterConfigsResponseBody) *DescribeFilterConfigsResponse
	GetBody() *DescribeFilterConfigsResponseBody
}

type DescribeFilterConfigsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFilterConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFilterConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilterConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFilterConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFilterConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFilterConfigsResponse) GetBody() *DescribeFilterConfigsResponseBody {
	return s.Body
}

func (s *DescribeFilterConfigsResponse) SetHeaders(v map[string]*string) *DescribeFilterConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFilterConfigsResponse) SetStatusCode(v int32) *DescribeFilterConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFilterConfigsResponse) SetBody(v *DescribeFilterConfigsResponseBody) *DescribeFilterConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeFilterConfigsResponse) Validate() error {
	return dara.Validate(s)
}

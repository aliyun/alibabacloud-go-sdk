// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficControlsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTrafficControlsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTrafficControlsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTrafficControlsResponseBody) *DescribeTrafficControlsResponse
	GetBody() *DescribeTrafficControlsResponseBody
}

type DescribeTrafficControlsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrafficControlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrafficControlsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficControlsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTrafficControlsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTrafficControlsResponse) GetBody() *DescribeTrafficControlsResponseBody {
	return s.Body
}

func (s *DescribeTrafficControlsResponse) SetHeaders(v map[string]*string) *DescribeTrafficControlsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrafficControlsResponse) SetStatusCode(v int32) *DescribeTrafficControlsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrafficControlsResponse) SetBody(v *DescribeTrafficControlsResponseBody) *DescribeTrafficControlsResponse {
	s.Body = v
	return s
}

func (s *DescribeTrafficControlsResponse) Validate() error {
	return dara.Validate(s)
}

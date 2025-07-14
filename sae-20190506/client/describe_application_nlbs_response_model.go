// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationNlbsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationNlbsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationNlbsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationNlbsResponseBody) *DescribeApplicationNlbsResponse
	GetBody() *DescribeApplicationNlbsResponseBody
}

type DescribeApplicationNlbsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationNlbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationNlbsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationNlbsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationNlbsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationNlbsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationNlbsResponse) GetBody() *DescribeApplicationNlbsResponseBody {
	return s.Body
}

func (s *DescribeApplicationNlbsResponse) SetHeaders(v map[string]*string) *DescribeApplicationNlbsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationNlbsResponse) SetStatusCode(v int32) *DescribeApplicationNlbsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationNlbsResponse) SetBody(v *DescribeApplicationNlbsResponseBody) *DescribeApplicationNlbsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationNlbsResponse) Validate() error {
	return dara.Validate(s)
}

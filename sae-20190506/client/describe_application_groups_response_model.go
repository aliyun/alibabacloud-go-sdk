// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationGroupsResponseBody) *DescribeApplicationGroupsResponse
	GetBody() *DescribeApplicationGroupsResponseBody
}

type DescribeApplicationGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationGroupsResponse) GetBody() *DescribeApplicationGroupsResponseBody {
	return s.Body
}

func (s *DescribeApplicationGroupsResponse) SetHeaders(v map[string]*string) *DescribeApplicationGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationGroupsResponse) SetStatusCode(v int32) *DescribeApplicationGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationGroupsResponse) SetBody(v *DescribeApplicationGroupsResponseBody) *DescribeApplicationGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationGroupsResponse) Validate() error {
	return dara.Validate(s)
}

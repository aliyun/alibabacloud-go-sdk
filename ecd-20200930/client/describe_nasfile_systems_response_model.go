// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNASFileSystemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNASFileSystemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNASFileSystemsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNASFileSystemsResponseBody) *DescribeNASFileSystemsResponse
	GetBody() *DescribeNASFileSystemsResponseBody
}

type DescribeNASFileSystemsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNASFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNASFileSystemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNASFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNASFileSystemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNASFileSystemsResponse) GetBody() *DescribeNASFileSystemsResponseBody {
	return s.Body
}

func (s *DescribeNASFileSystemsResponse) SetHeaders(v map[string]*string) *DescribeNASFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNASFileSystemsResponse) SetStatusCode(v int32) *DescribeNASFileSystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNASFileSystemsResponse) SetBody(v *DescribeNASFileSystemsResponseBody) *DescribeNASFileSystemsResponse {
	s.Body = v
	return s
}

func (s *DescribeNASFileSystemsResponse) Validate() error {
	return dara.Validate(s)
}

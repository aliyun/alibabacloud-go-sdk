// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDirectoriesResponseBody) *DescribeDirectoriesResponse
	GetBody() *DescribeDirectoriesResponseBody
}

type DescribeDirectoriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDirectoriesResponse) GetBody() *DescribeDirectoriesResponseBody {
	return s.Body
}

func (s *DescribeDirectoriesResponse) SetHeaders(v map[string]*string) *DescribeDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirectoriesResponse) SetStatusCode(v int32) *DescribeDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDirectoriesResponse) SetBody(v *DescribeDirectoriesResponseBody) *DescribeDirectoriesResponse {
	s.Body = v
	return s
}

func (s *DescribeDirectoriesResponse) Validate() error {
	return dara.Validate(s)
}

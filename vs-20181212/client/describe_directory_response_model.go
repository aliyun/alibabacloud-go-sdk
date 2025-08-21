// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDirectoryResponseBody) *DescribeDirectoryResponse
	GetBody() *DescribeDirectoryResponseBody
}

type DescribeDirectoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDirectoryResponse) GetBody() *DescribeDirectoryResponseBody {
	return s.Body
}

func (s *DescribeDirectoryResponse) SetHeaders(v map[string]*string) *DescribeDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirectoryResponse) SetStatusCode(v int32) *DescribeDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDirectoryResponse) SetBody(v *DescribeDirectoryResponseBody) *DescribeDirectoryResponse {
	s.Body = v
	return s
}

func (s *DescribeDirectoryResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarkPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMarkPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMarkPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMarkPageResponseBody) *DescribeMarkPageResponse
	GetBody() *DescribeMarkPageResponseBody
}

type DescribeMarkPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMarkPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMarkPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarkPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeMarkPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMarkPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMarkPageResponse) GetBody() *DescribeMarkPageResponseBody {
	return s.Body
}

func (s *DescribeMarkPageResponse) SetHeaders(v map[string]*string) *DescribeMarkPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeMarkPageResponse) SetStatusCode(v int32) *DescribeMarkPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMarkPageResponse) SetBody(v *DescribeMarkPageResponseBody) *DescribeMarkPageResponse {
	s.Body = v
	return s
}

func (s *DescribeMarkPageResponse) Validate() error {
	return dara.Validate(s)
}

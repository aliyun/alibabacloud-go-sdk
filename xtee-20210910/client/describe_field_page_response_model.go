// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFieldPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFieldPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFieldPageResponseBody) *DescribeFieldPageResponse
	GetBody() *DescribeFieldPageResponseBody
}

type DescribeFieldPageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFieldPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFieldPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeFieldPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFieldPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFieldPageResponse) GetBody() *DescribeFieldPageResponseBody {
	return s.Body
}

func (s *DescribeFieldPageResponse) SetHeaders(v map[string]*string) *DescribeFieldPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeFieldPageResponse) SetStatusCode(v int32) *DescribeFieldPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFieldPageResponse) SetBody(v *DescribeFieldPageResponseBody) *DescribeFieldPageResponse {
	s.Body = v
	return s
}

func (s *DescribeFieldPageResponse) Validate() error {
	return dara.Validate(s)
}

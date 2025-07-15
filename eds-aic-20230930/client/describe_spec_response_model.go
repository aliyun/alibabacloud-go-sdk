// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSpecResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSpecResponseBody) *DescribeSpecResponse
	GetBody() *DescribeSpecResponseBody
}

type DescribeSpecResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSpecResponse) GetBody() *DescribeSpecResponseBody {
	return s.Body
}

func (s *DescribeSpecResponse) SetHeaders(v map[string]*string) *DescribeSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeSpecResponse) SetStatusCode(v int32) *DescribeSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSpecResponse) SetBody(v *DescribeSpecResponseBody) *DescribeSpecResponse {
	s.Body = v
	return s
}

func (s *DescribeSpecResponse) Validate() error {
	return dara.Validate(s)
}

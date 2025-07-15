// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppSecurityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppSecurityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppSecurityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppSecurityResponseBody) *DescribeAppSecurityResponse
	GetBody() *DescribeAppSecurityResponseBody
}

type DescribeAppSecurityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppSecurityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppSecurityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppSecurityResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppSecurityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppSecurityResponse) GetBody() *DescribeAppSecurityResponseBody {
	return s.Body
}

func (s *DescribeAppSecurityResponse) SetHeaders(v map[string]*string) *DescribeAppSecurityResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppSecurityResponse) SetStatusCode(v int32) *DescribeAppSecurityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppSecurityResponse) SetBody(v *DescribeAppSecurityResponseBody) *DescribeAppSecurityResponse {
	s.Body = v
	return s
}

func (s *DescribeAppSecurityResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAreaBlockConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebAreaBlockConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebAreaBlockConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebAreaBlockConfigsResponseBody) *DescribeWebAreaBlockConfigsResponse
	GetBody() *DescribeWebAreaBlockConfigsResponseBody
}

type DescribeWebAreaBlockConfigsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebAreaBlockConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebAreaBlockConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAreaBlockConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebAreaBlockConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebAreaBlockConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebAreaBlockConfigsResponse) GetBody() *DescribeWebAreaBlockConfigsResponseBody {
	return s.Body
}

func (s *DescribeWebAreaBlockConfigsResponse) SetHeaders(v map[string]*string) *DescribeWebAreaBlockConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponse) SetStatusCode(v int32) *DescribeWebAreaBlockConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponse) SetBody(v *DescribeWebAreaBlockConfigsResponseBody) *DescribeWebAreaBlockConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeWebAreaBlockConfigsResponse) Validate() error {
	return dara.Validate(s)
}

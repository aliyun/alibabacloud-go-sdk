// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebApplicationResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationBody) *DescribeWebApplicationResponse
	GetBody() *WebApplicationBody
}

type DescribeWebApplicationResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebApplicationResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebApplicationResponse) GetBody() *WebApplicationBody {
	return s.Body
}

func (s *DescribeWebApplicationResponse) SetHeaders(v map[string]*string) *DescribeWebApplicationResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebApplicationResponse) SetStatusCode(v int32) *DescribeWebApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebApplicationResponse) SetBody(v *WebApplicationBody) *DescribeWebApplicationResponse {
	s.Body = v
	return s
}

func (s *DescribeWebApplicationResponse) Validate() error {
	return dara.Validate(s)
}

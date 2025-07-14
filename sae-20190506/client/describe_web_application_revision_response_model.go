// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebApplicationRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebApplicationRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebApplicationRevisionResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationRevisionBody) *DescribeWebApplicationRevisionResponse
	GetBody() *WebApplicationRevisionBody
}

type DescribeWebApplicationRevisionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationRevisionBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebApplicationRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebApplicationRevisionResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebApplicationRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebApplicationRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebApplicationRevisionResponse) GetBody() *WebApplicationRevisionBody {
	return s.Body
}

func (s *DescribeWebApplicationRevisionResponse) SetHeaders(v map[string]*string) *DescribeWebApplicationRevisionResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebApplicationRevisionResponse) SetStatusCode(v int32) *DescribeWebApplicationRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebApplicationRevisionResponse) SetBody(v *WebApplicationRevisionBody) *DescribeWebApplicationRevisionResponse {
	s.Body = v
	return s
}

func (s *DescribeWebApplicationRevisionResponse) Validate() error {
	return dara.Validate(s)
}

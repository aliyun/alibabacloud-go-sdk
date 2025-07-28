// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemPropertyTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSystemPropertyTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSystemPropertyTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSystemPropertyTemplatesResponseBody) *DescribeSystemPropertyTemplatesResponse
	GetBody() *DescribeSystemPropertyTemplatesResponseBody
}

type DescribeSystemPropertyTemplatesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSystemPropertyTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSystemPropertyTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemPropertyTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemPropertyTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSystemPropertyTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSystemPropertyTemplatesResponse) GetBody() *DescribeSystemPropertyTemplatesResponseBody {
	return s.Body
}

func (s *DescribeSystemPropertyTemplatesResponse) SetHeaders(v map[string]*string) *DescribeSystemPropertyTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponse) SetStatusCode(v int32) *DescribeSystemPropertyTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponse) SetBody(v *DescribeSystemPropertyTemplatesResponseBody) *DescribeSystemPropertyTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeSystemPropertyTemplatesResponse) Validate() error {
	return dara.Validate(s)
}

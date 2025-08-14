// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTemplateCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTemplateCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTemplateCountResponseBody) *DescribeTemplateCountResponse
	GetBody() *DescribeTemplateCountResponseBody
}

type DescribeTemplateCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplateCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplateCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTemplateCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTemplateCountResponse) GetBody() *DescribeTemplateCountResponseBody {
	return s.Body
}

func (s *DescribeTemplateCountResponse) SetHeaders(v map[string]*string) *DescribeTemplateCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateCountResponse) SetStatusCode(v int32) *DescribeTemplateCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateCountResponse) SetBody(v *DescribeTemplateCountResponseBody) *DescribeTemplateCountResponse {
	s.Body = v
	return s
}

func (s *DescribeTemplateCountResponse) Validate() error {
	return dara.Validate(s)
}

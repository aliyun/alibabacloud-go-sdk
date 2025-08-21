// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIndexTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIndexTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIndexTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIndexTemplateResponseBody) *DescribeIndexTemplateResponse
	GetBody() *DescribeIndexTemplateResponseBody
}

type DescribeIndexTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIndexTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIndexTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIndexTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeIndexTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIndexTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIndexTemplateResponse) GetBody() *DescribeIndexTemplateResponseBody {
	return s.Body
}

func (s *DescribeIndexTemplateResponse) SetHeaders(v map[string]*string) *DescribeIndexTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeIndexTemplateResponse) SetStatusCode(v int32) *DescribeIndexTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIndexTemplateResponse) SetBody(v *DescribeIndexTemplateResponseBody) *DescribeIndexTemplateResponse {
	s.Body = v
	return s
}

func (s *DescribeIndexTemplateResponse) Validate() error {
	return dara.Validate(s)
}

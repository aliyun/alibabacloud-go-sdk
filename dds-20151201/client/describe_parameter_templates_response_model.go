// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParameterTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParameterTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParameterTemplatesResponseBody) *DescribeParameterTemplatesResponse
	GetBody() *DescribeParameterTemplatesResponseBody
}

type DescribeParameterTemplatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParameterTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParameterTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParameterTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParameterTemplatesResponse) GetBody() *DescribeParameterTemplatesResponseBody {
	return s.Body
}

func (s *DescribeParameterTemplatesResponse) SetHeaders(v map[string]*string) *DescribeParameterTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterTemplatesResponse) SetStatusCode(v int32) *DescribeParameterTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParameterTemplatesResponse) SetBody(v *DescribeParameterTemplatesResponseBody) *DescribeParameterTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeParameterTemplatesResponse) Validate() error {
	return dara.Validate(s)
}

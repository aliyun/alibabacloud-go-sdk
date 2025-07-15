// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePluginSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePluginSchemasResponse
	GetStatusCode() *int32
	SetBody(v *DescribePluginSchemasResponseBody) *DescribePluginSchemasResponse
	GetBody() *DescribePluginSchemasResponseBody
}

type DescribePluginSchemasResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePluginSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePluginSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginSchemasResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePluginSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePluginSchemasResponse) GetBody() *DescribePluginSchemasResponseBody {
	return s.Body
}

func (s *DescribePluginSchemasResponse) SetHeaders(v map[string]*string) *DescribePluginSchemasResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginSchemasResponse) SetStatusCode(v int32) *DescribePluginSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePluginSchemasResponse) SetBody(v *DescribePluginSchemasResponseBody) *DescribePluginSchemasResponse {
	s.Body = v
	return s
}

func (s *DescribePluginSchemasResponse) Validate() error {
	return dara.Validate(s)
}

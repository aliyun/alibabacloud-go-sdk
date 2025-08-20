// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdbMySqlSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdbMySqlSchemasResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdbMySqlSchemasResponseBody) *DescribeAdbMySqlSchemasResponse
	GetBody() *DescribeAdbMySqlSchemasResponseBody
}

type DescribeAdbMySqlSchemasResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdbMySqlSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdbMySqlSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlSchemasResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdbMySqlSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdbMySqlSchemasResponse) GetBody() *DescribeAdbMySqlSchemasResponseBody {
	return s.Body
}

func (s *DescribeAdbMySqlSchemasResponse) SetHeaders(v map[string]*string) *DescribeAdbMySqlSchemasResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdbMySqlSchemasResponse) SetStatusCode(v int32) *DescribeAdbMySqlSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdbMySqlSchemasResponse) SetBody(v *DescribeAdbMySqlSchemasResponseBody) *DescribeAdbMySqlSchemasResponse {
	s.Body = v
	return s
}

func (s *DescribeAdbMySqlSchemasResponse) Validate() error {
	return dara.Validate(s)
}

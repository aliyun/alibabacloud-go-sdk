// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdbMySqlTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdbMySqlTablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdbMySqlTablesResponseBody) *DescribeAdbMySqlTablesResponse
	GetBody() *DescribeAdbMySqlTablesResponseBody
}

type DescribeAdbMySqlTablesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdbMySqlTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdbMySqlTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdbMySqlTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdbMySqlTablesResponse) GetBody() *DescribeAdbMySqlTablesResponseBody {
	return s.Body
}

func (s *DescribeAdbMySqlTablesResponse) SetHeaders(v map[string]*string) *DescribeAdbMySqlTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdbMySqlTablesResponse) SetStatusCode(v int32) *DescribeAdbMySqlTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdbMySqlTablesResponse) SetBody(v *DescribeAdbMySqlTablesResponseBody) *DescribeAdbMySqlTablesResponse {
	s.Body = v
	return s
}

func (s *DescribeAdbMySqlTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

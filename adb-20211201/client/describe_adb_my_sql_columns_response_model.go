// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlColumnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdbMySqlColumnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdbMySqlColumnsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdbMySqlColumnsResponseBody) *DescribeAdbMySqlColumnsResponse
	GetBody() *DescribeAdbMySqlColumnsResponseBody
}

type DescribeAdbMySqlColumnsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdbMySqlColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdbMySqlColumnsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlColumnsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlColumnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdbMySqlColumnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdbMySqlColumnsResponse) GetBody() *DescribeAdbMySqlColumnsResponseBody {
	return s.Body
}

func (s *DescribeAdbMySqlColumnsResponse) SetHeaders(v map[string]*string) *DescribeAdbMySqlColumnsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdbMySqlColumnsResponse) SetStatusCode(v int32) *DescribeAdbMySqlColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdbMySqlColumnsResponse) SetBody(v *DescribeAdbMySqlColumnsResponseBody) *DescribeAdbMySqlColumnsResponse {
	s.Body = v
	return s
}

func (s *DescribeAdbMySqlColumnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

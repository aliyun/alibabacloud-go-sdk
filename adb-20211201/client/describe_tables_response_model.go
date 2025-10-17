// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTablesResponseBody) *DescribeTablesResponse
	GetBody() *DescribeTablesResponseBody
}

type DescribeTablesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTablesResponse) GetBody() *DescribeTablesResponseBody {
	return s.Body
}

func (s *DescribeTablesResponse) SetHeaders(v map[string]*string) *DescribeTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTablesResponse) SetStatusCode(v int32) *DescribeTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTablesResponse) SetBody(v *DescribeTablesResponseBody) *DescribeTablesResponse {
	s.Body = v
	return s
}

func (s *DescribeTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

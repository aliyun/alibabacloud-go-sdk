// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInclinedTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInclinedTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInclinedTablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInclinedTablesResponseBody) *DescribeInclinedTablesResponse
	GetBody() *DescribeInclinedTablesResponseBody
}

type DescribeInclinedTablesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInclinedTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInclinedTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInclinedTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInclinedTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInclinedTablesResponse) GetBody() *DescribeInclinedTablesResponseBody {
	return s.Body
}

func (s *DescribeInclinedTablesResponse) SetHeaders(v map[string]*string) *DescribeInclinedTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInclinedTablesResponse) SetStatusCode(v int32) *DescribeInclinedTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInclinedTablesResponse) SetBody(v *DescribeInclinedTablesResponseBody) *DescribeInclinedTablesResponse {
	s.Body = v
	return s
}

func (s *DescribeInclinedTablesResponse) Validate() error {
	return dara.Validate(s)
}

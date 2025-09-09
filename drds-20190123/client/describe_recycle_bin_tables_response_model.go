// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecycleBinTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecycleBinTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecycleBinTablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecycleBinTablesResponseBody) *DescribeRecycleBinTablesResponse
	GetBody() *DescribeRecycleBinTablesResponseBody
}

type DescribeRecycleBinTablesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecycleBinTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecycleBinTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecycleBinTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecycleBinTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecycleBinTablesResponse) GetBody() *DescribeRecycleBinTablesResponseBody {
	return s.Body
}

func (s *DescribeRecycleBinTablesResponse) SetHeaders(v map[string]*string) *DescribeRecycleBinTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecycleBinTablesResponse) SetStatusCode(v int32) *DescribeRecycleBinTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecycleBinTablesResponse) SetBody(v *DescribeRecycleBinTablesResponseBody) *DescribeRecycleBinTablesResponse {
	s.Body = v
	return s
}

func (s *DescribeRecycleBinTablesResponse) Validate() error {
	return dara.Validate(s)
}

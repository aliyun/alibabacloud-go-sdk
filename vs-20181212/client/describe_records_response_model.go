// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecordsResponseBody) *DescribeRecordsResponse
	GetBody() *DescribeRecordsResponseBody
}

type DescribeRecordsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecordsResponse) GetBody() *DescribeRecordsResponseBody {
	return s.Body
}

func (s *DescribeRecordsResponse) SetHeaders(v map[string]*string) *DescribeRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordsResponse) SetStatusCode(v int32) *DescribeRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordsResponse) SetBody(v *DescribeRecordsResponseBody) *DescribeRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeRecordsResponse) Validate() error {
	return dara.Validate(s)
}

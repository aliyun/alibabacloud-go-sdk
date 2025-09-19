// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomBlockRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomBlockRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomBlockRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomBlockRecordsResponseBody) *DescribeCustomBlockRecordsResponse
	GetBody() *DescribeCustomBlockRecordsResponseBody
}

type DescribeCustomBlockRecordsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomBlockRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomBlockRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBlockRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomBlockRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomBlockRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomBlockRecordsResponse) GetBody() *DescribeCustomBlockRecordsResponseBody {
	return s.Body
}

func (s *DescribeCustomBlockRecordsResponse) SetHeaders(v map[string]*string) *DescribeCustomBlockRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomBlockRecordsResponse) SetStatusCode(v int32) *DescribeCustomBlockRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomBlockRecordsResponse) SetBody(v *DescribeCustomBlockRecordsResponseBody) *DescribeCustomBlockRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomBlockRecordsResponse) Validate() error {
	return dara.Validate(s)
}

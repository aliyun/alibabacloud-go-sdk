// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeErrorLogRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeErrorLogRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeErrorLogRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeErrorLogRecordsResponseBody) *DescribeErrorLogRecordsResponse
	GetBody() *DescribeErrorLogRecordsResponseBody
}

type DescribeErrorLogRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeErrorLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeErrorLogRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeErrorLogRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeErrorLogRecordsResponse) GetBody() *DescribeErrorLogRecordsResponseBody {
	return s.Body
}

func (s *DescribeErrorLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeErrorLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeErrorLogRecordsResponse) SetStatusCode(v int32) *DescribeErrorLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeErrorLogRecordsResponse) SetBody(v *DescribeErrorLogRecordsResponseBody) *DescribeErrorLogRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeErrorLogRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

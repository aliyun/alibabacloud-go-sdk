// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHALogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHALogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHALogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHALogsResponseBody) *DescribeHALogsResponse
	GetBody() *DescribeHALogsResponseBody
}

type DescribeHALogsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHALogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHALogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHALogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHALogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHALogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHALogsResponse) GetBody() *DescribeHALogsResponseBody {
	return s.Body
}

func (s *DescribeHALogsResponse) SetHeaders(v map[string]*string) *DescribeHALogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHALogsResponse) SetStatusCode(v int32) *DescribeHALogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHALogsResponse) SetBody(v *DescribeHALogsResponseBody) *DescribeHALogsResponse {
	s.Body = v
	return s
}

func (s *DescribeHALogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

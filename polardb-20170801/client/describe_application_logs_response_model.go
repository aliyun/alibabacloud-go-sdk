// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationLogsResponseBody) *DescribeApplicationLogsResponse
	GetBody() *DescribeApplicationLogsResponseBody
}

type DescribeApplicationLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationLogsResponse) GetBody() *DescribeApplicationLogsResponseBody {
	return s.Body
}

func (s *DescribeApplicationLogsResponse) SetHeaders(v map[string]*string) *DescribeApplicationLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationLogsResponse) SetStatusCode(v int32) *DescribeApplicationLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationLogsResponse) SetBody(v *DescribeApplicationLogsResponseBody) *DescribeApplicationLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

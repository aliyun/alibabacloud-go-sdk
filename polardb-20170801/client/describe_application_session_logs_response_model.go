// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSessionLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationSessionLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationSessionLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationSessionLogsResponseBody) *DescribeApplicationSessionLogsResponse
	GetBody() *DescribeApplicationSessionLogsResponseBody
}

type DescribeApplicationSessionLogsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationSessionLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationSessionLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSessionLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSessionLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationSessionLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationSessionLogsResponse) GetBody() *DescribeApplicationSessionLogsResponseBody {
	return s.Body
}

func (s *DescribeApplicationSessionLogsResponse) SetHeaders(v map[string]*string) *DescribeApplicationSessionLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationSessionLogsResponse) SetStatusCode(v int32) *DescribeApplicationSessionLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponse) SetBody(v *DescribeApplicationSessionLogsResponseBody) *DescribeApplicationSessionLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationSessionLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChangeLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChangeLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChangeLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChangeLogsResponseBody) *DescribeChangeLogsResponse
	GetBody() *DescribeChangeLogsResponseBody
}

type DescribeChangeLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChangeLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChangeLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChangeLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChangeLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChangeLogsResponse) GetBody() *DescribeChangeLogsResponseBody {
	return s.Body
}

func (s *DescribeChangeLogsResponse) SetHeaders(v map[string]*string) *DescribeChangeLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeChangeLogsResponse) SetStatusCode(v int32) *DescribeChangeLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChangeLogsResponse) SetBody(v *DescribeChangeLogsResponseBody) *DescribeChangeLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeChangeLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

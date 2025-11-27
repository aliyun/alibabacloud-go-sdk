// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeErrorLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeErrorLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeErrorLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeErrorLogsResponseBody) *DescribeErrorLogsResponse
	GetBody() *DescribeErrorLogsResponseBody
}

type DescribeErrorLogsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeErrorLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeErrorLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeErrorLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeErrorLogsResponse) GetBody() *DescribeErrorLogsResponseBody {
	return s.Body
}

func (s *DescribeErrorLogsResponse) SetHeaders(v map[string]*string) *DescribeErrorLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeErrorLogsResponse) SetStatusCode(v int32) *DescribeErrorLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeErrorLogsResponse) SetBody(v *DescribeErrorLogsResponseBody) *DescribeErrorLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeErrorLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeductLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeductLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeductLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeductLogsResponseBody) *DescribeDeductLogsResponse
	GetBody() *DescribeDeductLogsResponseBody
}

type DescribeDeductLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeductLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeductLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeductLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeductLogsResponse) GetBody() *DescribeDeductLogsResponseBody {
	return s.Body
}

func (s *DescribeDeductLogsResponse) SetHeaders(v map[string]*string) *DescribeDeductLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeductLogsResponse) SetStatusCode(v int32) *DescribeDeductLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeductLogsResponse) SetBody(v *DescribeDeductLogsResponseBody) *DescribeDeductLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeDeductLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

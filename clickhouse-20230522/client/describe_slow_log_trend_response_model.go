// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlowLogTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlowLogTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlowLogTrendResponseBody) *DescribeSlowLogTrendResponse
	GetBody() *DescribeSlowLogTrendResponseBody
}

type DescribeSlowLogTrendResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowLogTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlowLogTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlowLogTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlowLogTrendResponse) GetBody() *DescribeSlowLogTrendResponseBody {
	return s.Body
}

func (s *DescribeSlowLogTrendResponse) SetHeaders(v map[string]*string) *DescribeSlowLogTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogTrendResponse) SetStatusCode(v int32) *DescribeSlowLogTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowLogTrendResponse) SetBody(v *DescribeSlowLogTrendResponseBody) *DescribeSlowLogTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeSlowLogTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

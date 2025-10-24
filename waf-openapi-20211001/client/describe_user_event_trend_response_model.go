// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEventTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserEventTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserEventTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserEventTrendResponseBody) *DescribeUserEventTrendResponse
	GetBody() *DescribeUserEventTrendResponseBody
}

type DescribeUserEventTrendResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserEventTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserEventTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEventTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserEventTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserEventTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserEventTrendResponse) GetBody() *DescribeUserEventTrendResponseBody {
	return s.Body
}

func (s *DescribeUserEventTrendResponse) SetHeaders(v map[string]*string) *DescribeUserEventTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserEventTrendResponse) SetStatusCode(v int32) *DescribeUserEventTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserEventTrendResponse) SetBody(v *DescribeUserEventTrendResponseBody) *DescribeUserEventTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeUserEventTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

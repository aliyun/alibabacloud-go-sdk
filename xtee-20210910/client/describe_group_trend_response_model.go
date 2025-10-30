// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupTrendResponseBody) *DescribeGroupTrendResponse
	GetBody() *DescribeGroupTrendResponseBody
}

type DescribeGroupTrendResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupTrendResponse) GetBody() *DescribeGroupTrendResponseBody {
	return s.Body
}

func (s *DescribeGroupTrendResponse) SetHeaders(v map[string]*string) *DescribeGroupTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupTrendResponse) SetStatusCode(v int32) *DescribeGroupTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupTrendResponse) SetBody(v *DescribeGroupTrendResponseBody) *DescribeGroupTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnprotectedPortTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUnprotectedPortTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUnprotectedPortTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUnprotectedPortTrendResponseBody) *DescribeUnprotectedPortTrendResponse
	GetBody() *DescribeUnprotectedPortTrendResponseBody
}

type DescribeUnprotectedPortTrendResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUnprotectedPortTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUnprotectedPortTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnprotectedPortTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeUnprotectedPortTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUnprotectedPortTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUnprotectedPortTrendResponse) GetBody() *DescribeUnprotectedPortTrendResponseBody {
	return s.Body
}

func (s *DescribeUnprotectedPortTrendResponse) SetHeaders(v map[string]*string) *DescribeUnprotectedPortTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeUnprotectedPortTrendResponse) SetStatusCode(v int32) *DescribeUnprotectedPortTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUnprotectedPortTrendResponse) SetBody(v *DescribeUnprotectedPortTrendResponseBody) *DescribeUnprotectedPortTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeUnprotectedPortTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

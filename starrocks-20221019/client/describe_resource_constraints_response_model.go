// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceConstraintsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceConstraintsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceConstraintsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceConstraintsResponseBody) *DescribeResourceConstraintsResponse
	GetBody() *DescribeResourceConstraintsResponseBody
}

type DescribeResourceConstraintsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceConstraintsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceConstraintsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceConstraintsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceConstraintsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceConstraintsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceConstraintsResponse) GetBody() *DescribeResourceConstraintsResponseBody {
	return s.Body
}

func (s *DescribeResourceConstraintsResponse) SetHeaders(v map[string]*string) *DescribeResourceConstraintsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceConstraintsResponse) SetStatusCode(v int32) *DescribeResourceConstraintsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceConstraintsResponse) SetBody(v *DescribeResourceConstraintsResponseBody) *DescribeResourceConstraintsResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceConstraintsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

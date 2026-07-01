// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBProjectsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBProjectsResponseBody) *DescribeAgenticDBProjectsResponse
	GetBody() *DescribeAgenticDBProjectsResponseBody
}

type DescribeAgenticDBProjectsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBProjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBProjectsResponse) GetBody() *DescribeAgenticDBProjectsResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBProjectsResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBProjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBProjectsResponse) SetStatusCode(v int32) *DescribeAgenticDBProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBProjectsResponse) SetBody(v *DescribeAgenticDBProjectsResponseBody) *DescribeAgenticDBProjectsResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

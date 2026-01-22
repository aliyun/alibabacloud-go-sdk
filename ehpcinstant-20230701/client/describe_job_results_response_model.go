// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobResultsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobResultsResponseBody) *DescribeJobResultsResponse
	GetBody() *DescribeJobResultsResponseBody
}

type DescribeJobResultsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobResultsResponse) GetBody() *DescribeJobResultsResponseBody {
	return s.Body
}

func (s *DescribeJobResultsResponse) SetHeaders(v map[string]*string) *DescribeJobResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobResultsResponse) SetStatusCode(v int32) *DescribeJobResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobResultsResponse) SetBody(v *DescribeJobResultsResponseBody) *DescribeJobResultsResponse {
	s.Body = v
	return s
}

func (s *DescribeJobResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

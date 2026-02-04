// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowlogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFlowlogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFlowlogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFlowlogsResponseBody) *DescribeFlowlogsResponse
	GetBody() *DescribeFlowlogsResponseBody
}

type DescribeFlowlogsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowlogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowlogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowlogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowlogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFlowlogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFlowlogsResponse) GetBody() *DescribeFlowlogsResponseBody {
	return s.Body
}

func (s *DescribeFlowlogsResponse) SetHeaders(v map[string]*string) *DescribeFlowlogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowlogsResponse) SetStatusCode(v int32) *DescribeFlowlogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowlogsResponse) SetBody(v *DescribeFlowlogsResponseBody) *DescribeFlowlogsResponse {
	s.Body = v
	return s
}

func (s *DescribeFlowlogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

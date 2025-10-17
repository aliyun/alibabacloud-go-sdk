// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBNodesParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBNodesParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBNodesParametersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBNodesParametersResponseBody) *DescribeDBNodesParametersResponse
	GetBody() *DescribeDBNodesParametersResponseBody
}

type DescribeDBNodesParametersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBNodesParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBNodesParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodesParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBNodesParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBNodesParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBNodesParametersResponse) GetBody() *DescribeDBNodesParametersResponseBody {
	return s.Body
}

func (s *DescribeDBNodesParametersResponse) SetHeaders(v map[string]*string) *DescribeDBNodesParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBNodesParametersResponse) SetStatusCode(v int32) *DescribeDBNodesParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBNodesParametersResponse) SetBody(v *DescribeDBNodesParametersResponseBody) *DescribeDBNodesParametersResponse {
	s.Body = v
	return s
}

func (s *DescribeDBNodesParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

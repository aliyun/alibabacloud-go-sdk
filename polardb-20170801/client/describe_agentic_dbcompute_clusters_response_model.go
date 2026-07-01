// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBComputeClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBComputeClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBComputeClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBComputeClustersResponseBody) *DescribeAgenticDBComputeClustersResponse
	GetBody() *DescribeAgenticDBComputeClustersResponseBody
}

type DescribeAgenticDBComputeClustersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBComputeClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBComputeClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBComputeClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBComputeClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBComputeClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBComputeClustersResponse) GetBody() *DescribeAgenticDBComputeClustersResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBComputeClustersResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBComputeClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponse) SetStatusCode(v int32) *DescribeAgenticDBComputeClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponse) SetBody(v *DescribeAgenticDBComputeClustersResponseBody) *DescribeAgenticDBComputeClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

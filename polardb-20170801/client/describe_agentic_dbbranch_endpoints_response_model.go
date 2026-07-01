// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBBranchEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBBranchEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBBranchEndpointsResponseBody) *DescribeAgenticDBBranchEndpointsResponse
	GetBody() *DescribeAgenticDBBranchEndpointsResponseBody
}

type DescribeAgenticDBBranchEndpointsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBBranchEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBBranchEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBBranchEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBBranchEndpointsResponse) GetBody() *DescribeAgenticDBBranchEndpointsResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBBranchEndpointsResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBBranchEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponse) SetStatusCode(v int32) *DescribeAgenticDBBranchEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponse) SetBody(v *DescribeAgenticDBBranchEndpointsResponseBody) *DescribeAgenticDBBranchEndpointsResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBBranchEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

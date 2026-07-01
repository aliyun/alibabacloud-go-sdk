// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBBranchesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBBranchesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBBranchesResponseBody) *DescribeAgenticDBBranchesResponse
	GetBody() *DescribeAgenticDBBranchesResponseBody
}

type DescribeAgenticDBBranchesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBBranchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBBranchesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBBranchesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBBranchesResponse) GetBody() *DescribeAgenticDBBranchesResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBBranchesResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBBranchesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBBranchesResponse) SetStatusCode(v int32) *DescribeAgenticDBBranchesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponse) SetBody(v *DescribeAgenticDBBranchesResponseBody) *DescribeAgenticDBBranchesResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBBranchesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchLineageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBBranchLineageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBBranchLineageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBBranchLineageResponseBody) *DescribeAgenticDBBranchLineageResponse
	GetBody() *DescribeAgenticDBBranchLineageResponseBody
}

type DescribeAgenticDBBranchLineageResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBBranchLineageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBBranchLineageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchLineageResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchLineageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBBranchLineageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBBranchLineageResponse) GetBody() *DescribeAgenticDBBranchLineageResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBBranchLineageResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBBranchLineageResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponse) SetStatusCode(v int32) *DescribeAgenticDBBranchLineageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponse) SetBody(v *DescribeAgenticDBBranchLineageResponseBody) *DescribeAgenticDBBranchLineageResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

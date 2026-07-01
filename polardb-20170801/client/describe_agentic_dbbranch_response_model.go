// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBBranchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBBranchResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBBranchResponseBody) *DescribeAgenticDBBranchResponse
	GetBody() *DescribeAgenticDBBranchResponseBody
}

type DescribeAgenticDBBranchResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBBranchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBBranchResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBBranchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBBranchResponse) GetBody() *DescribeAgenticDBBranchResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBBranchResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBBranchResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBBranchResponse) SetStatusCode(v int32) *DescribeAgenticDBBranchResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBBranchResponse) SetBody(v *DescribeAgenticDBBranchResponseBody) *DescribeAgenticDBBranchResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBBranchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

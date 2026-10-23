// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchRestoreWindowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBBranchRestoreWindowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBBranchRestoreWindowResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBBranchRestoreWindowResponseBody) *DescribeAgenticDBBranchRestoreWindowResponse
	GetBody() *DescribeAgenticDBBranchRestoreWindowResponseBody
}

type DescribeAgenticDBBranchRestoreWindowResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBBranchRestoreWindowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBBranchRestoreWindowResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchRestoreWindowResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchRestoreWindowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBBranchRestoreWindowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBBranchRestoreWindowResponse) GetBody() *DescribeAgenticDBBranchRestoreWindowResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBBranchRestoreWindowResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBBranchRestoreWindowResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowResponse) SetStatusCode(v int32) *DescribeAgenticDBBranchRestoreWindowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowResponse) SetBody(v *DescribeAgenticDBBranchRestoreWindowResponseBody) *DescribeAgenticDBBranchRestoreWindowResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBBranchRestoreWindowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

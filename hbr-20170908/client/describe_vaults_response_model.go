// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVaultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVaultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVaultsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVaultsResponseBody) *DescribeVaultsResponse
	GetBody() *DescribeVaultsResponseBody
}

type DescribeVaultsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVaultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVaultsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVaultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVaultsResponse) GetBody() *DescribeVaultsResponseBody {
	return s.Body
}

func (s *DescribeVaultsResponse) SetHeaders(v map[string]*string) *DescribeVaultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVaultsResponse) SetStatusCode(v int32) *DescribeVaultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVaultsResponse) SetBody(v *DescribeVaultsResponseBody) *DescribeVaultsResponse {
	s.Body = v
	return s
}

func (s *DescribeVaultsResponse) Validate() error {
	return dara.Validate(s)
}

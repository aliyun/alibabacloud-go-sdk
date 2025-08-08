// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProjectNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProjectNodesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProjectNodesResponseBody) *DescribeProjectNodesResponse
	GetBody() *DescribeProjectNodesResponseBody
}

type DescribeProjectNodesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProjectNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProjectNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProjectNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProjectNodesResponse) GetBody() *DescribeProjectNodesResponseBody {
	return s.Body
}

func (s *DescribeProjectNodesResponse) SetHeaders(v map[string]*string) *DescribeProjectNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectNodesResponse) SetStatusCode(v int32) *DescribeProjectNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectNodesResponse) SetBody(v *DescribeProjectNodesResponseBody) *DescribeProjectNodesResponse {
	s.Body = v
	return s
}

func (s *DescribeProjectNodesResponse) Validate() error {
	return dara.Validate(s)
}

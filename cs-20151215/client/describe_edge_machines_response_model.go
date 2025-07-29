// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMachinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEdgeMachinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEdgeMachinesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEdgeMachinesResponseBody) *DescribeEdgeMachinesResponse
	GetBody() *DescribeEdgeMachinesResponseBody
}

type DescribeEdgeMachinesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEdgeMachinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEdgeMachinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEdgeMachinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEdgeMachinesResponse) GetBody() *DescribeEdgeMachinesResponseBody {
	return s.Body
}

func (s *DescribeEdgeMachinesResponse) SetHeaders(v map[string]*string) *DescribeEdgeMachinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdgeMachinesResponse) SetStatusCode(v int32) *DescribeEdgeMachinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEdgeMachinesResponse) SetBody(v *DescribeEdgeMachinesResponseBody) *DescribeEdgeMachinesResponse {
	s.Body = v
	return s
}

func (s *DescribeEdgeMachinesResponse) Validate() error {
	return dara.Validate(s)
}

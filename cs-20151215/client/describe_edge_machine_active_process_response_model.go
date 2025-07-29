// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMachineActiveProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEdgeMachineActiveProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEdgeMachineActiveProcessResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEdgeMachineActiveProcessResponseBody) *DescribeEdgeMachineActiveProcessResponse
	GetBody() *DescribeEdgeMachineActiveProcessResponseBody
}

type DescribeEdgeMachineActiveProcessResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEdgeMachineActiveProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEdgeMachineActiveProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachineActiveProcessResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachineActiveProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEdgeMachineActiveProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEdgeMachineActiveProcessResponse) GetBody() *DescribeEdgeMachineActiveProcessResponseBody {
	return s.Body
}

func (s *DescribeEdgeMachineActiveProcessResponse) SetHeaders(v map[string]*string) *DescribeEdgeMachineActiveProcessResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdgeMachineActiveProcessResponse) SetStatusCode(v int32) *DescribeEdgeMachineActiveProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEdgeMachineActiveProcessResponse) SetBody(v *DescribeEdgeMachineActiveProcessResponseBody) *DescribeEdgeMachineActiveProcessResponse {
	s.Body = v
	return s
}

func (s *DescribeEdgeMachineActiveProcessResponse) Validate() error {
	return dara.Validate(s)
}

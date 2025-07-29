// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMachineModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEdgeMachineModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEdgeMachineModelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEdgeMachineModelsResponseBody) *DescribeEdgeMachineModelsResponse
	GetBody() *DescribeEdgeMachineModelsResponseBody
}

type DescribeEdgeMachineModelsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEdgeMachineModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEdgeMachineModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachineModelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachineModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEdgeMachineModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEdgeMachineModelsResponse) GetBody() *DescribeEdgeMachineModelsResponseBody {
	return s.Body
}

func (s *DescribeEdgeMachineModelsResponse) SetHeaders(v map[string]*string) *DescribeEdgeMachineModelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdgeMachineModelsResponse) SetStatusCode(v int32) *DescribeEdgeMachineModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEdgeMachineModelsResponse) SetBody(v *DescribeEdgeMachineModelsResponseBody) *DescribeEdgeMachineModelsResponse {
	s.Body = v
	return s
}

func (s *DescribeEdgeMachineModelsResponse) Validate() error {
	return dara.Validate(s)
}

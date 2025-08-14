// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIAgentInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIAgentInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIAgentInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIAgentInstanceResponseBody) *DescribeAIAgentInstanceResponse
	GetBody() *DescribeAIAgentInstanceResponseBody
}

type DescribeAIAgentInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIAgentInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIAgentInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIAgentInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIAgentInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIAgentInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIAgentInstanceResponse) GetBody() *DescribeAIAgentInstanceResponseBody {
	return s.Body
}

func (s *DescribeAIAgentInstanceResponse) SetHeaders(v map[string]*string) *DescribeAIAgentInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIAgentInstanceResponse) SetStatusCode(v int32) *DescribeAIAgentInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIAgentInstanceResponse) SetBody(v *DescribeAIAgentInstanceResponseBody) *DescribeAIAgentInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeAIAgentInstanceResponse) Validate() error {
	return dara.Validate(s)
}

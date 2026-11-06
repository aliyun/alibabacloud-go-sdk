// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIAgentOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIAgentOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIAgentOrderResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIAgentOrderResponseBody) *DescribeAIAgentOrderResponse
	GetBody() *DescribeAIAgentOrderResponseBody
}

type DescribeAIAgentOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIAgentOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIAgentOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIAgentOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIAgentOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIAgentOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIAgentOrderResponse) GetBody() *DescribeAIAgentOrderResponseBody {
	return s.Body
}

func (s *DescribeAIAgentOrderResponse) SetHeaders(v map[string]*string) *DescribeAIAgentOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIAgentOrderResponse) SetStatusCode(v int32) *DescribeAIAgentOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIAgentOrderResponse) SetBody(v *DescribeAIAgentOrderResponseBody) *DescribeAIAgentOrderResponse {
	s.Body = v
	return s
}

func (s *DescribeAIAgentOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

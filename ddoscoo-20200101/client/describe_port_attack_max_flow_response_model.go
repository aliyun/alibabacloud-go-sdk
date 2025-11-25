// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortAttackMaxFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortAttackMaxFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortAttackMaxFlowResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortAttackMaxFlowResponseBody) *DescribePortAttackMaxFlowResponse
	GetBody() *DescribePortAttackMaxFlowResponseBody
}

type DescribePortAttackMaxFlowResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortAttackMaxFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortAttackMaxFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortAttackMaxFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribePortAttackMaxFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortAttackMaxFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortAttackMaxFlowResponse) GetBody() *DescribePortAttackMaxFlowResponseBody {
	return s.Body
}

func (s *DescribePortAttackMaxFlowResponse) SetHeaders(v map[string]*string) *DescribePortAttackMaxFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribePortAttackMaxFlowResponse) SetStatusCode(v int32) *DescribePortAttackMaxFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortAttackMaxFlowResponse) SetBody(v *DescribePortAttackMaxFlowResponseBody) *DescribePortAttackMaxFlowResponse {
	s.Body = v
	return s
}

func (s *DescribePortAttackMaxFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

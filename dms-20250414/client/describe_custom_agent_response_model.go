// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomAgentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomAgentResponseBody) *DescribeCustomAgentResponse
	GetBody() *DescribeCustomAgentResponseBody
}

type DescribeCustomAgentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomAgentResponse) GetBody() *DescribeCustomAgentResponseBody {
	return s.Body
}

func (s *DescribeCustomAgentResponse) SetHeaders(v map[string]*string) *DescribeCustomAgentResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomAgentResponse) SetStatusCode(v int32) *DescribeCustomAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomAgentResponse) SetBody(v *DescribeCustomAgentResponseBody) *DescribeCustomAgentResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

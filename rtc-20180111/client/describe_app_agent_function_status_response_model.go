// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppAgentFunctionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppAgentFunctionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppAgentFunctionStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppAgentFunctionStatusResponseBody) *DescribeAppAgentFunctionStatusResponse
	GetBody() *DescribeAppAgentFunctionStatusResponseBody
}

type DescribeAppAgentFunctionStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppAgentFunctionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppAgentFunctionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentFunctionStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentFunctionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppAgentFunctionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppAgentFunctionStatusResponse) GetBody() *DescribeAppAgentFunctionStatusResponseBody {
	return s.Body
}

func (s *DescribeAppAgentFunctionStatusResponse) SetHeaders(v map[string]*string) *DescribeAppAgentFunctionStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppAgentFunctionStatusResponse) SetStatusCode(v int32) *DescribeAppAgentFunctionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppAgentFunctionStatusResponse) SetBody(v *DescribeAppAgentFunctionStatusResponseBody) *DescribeAppAgentFunctionStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAppAgentFunctionStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExternalAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExternalAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExternalAgentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExternalAgentResponseBody) *DescribeExternalAgentResponse
	GetBody() *DescribeExternalAgentResponseBody
}

type DescribeExternalAgentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExternalAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExternalAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExternalAgentResponse) GoString() string {
	return s.String()
}

func (s *DescribeExternalAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExternalAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExternalAgentResponse) GetBody() *DescribeExternalAgentResponseBody {
	return s.Body
}

func (s *DescribeExternalAgentResponse) SetHeaders(v map[string]*string) *DescribeExternalAgentResponse {
	s.Headers = v
	return s
}

func (s *DescribeExternalAgentResponse) SetStatusCode(v int32) *DescribeExternalAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExternalAgentResponse) SetBody(v *DescribeExternalAgentResponseBody) *DescribeExternalAgentResponse {
	s.Body = v
	return s
}

func (s *DescribeExternalAgentResponse) Validate() error {
	return dara.Validate(s)
}

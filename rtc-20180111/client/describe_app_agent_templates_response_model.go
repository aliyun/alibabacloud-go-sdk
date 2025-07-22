// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppAgentTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppAgentTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppAgentTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppAgentTemplatesResponseBody) *DescribeAppAgentTemplatesResponse
	GetBody() *DescribeAppAgentTemplatesResponseBody
}

type DescribeAppAgentTemplatesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppAgentTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppAgentTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppAgentTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppAgentTemplatesResponse) GetBody() *DescribeAppAgentTemplatesResponseBody {
	return s.Body
}

func (s *DescribeAppAgentTemplatesResponse) SetHeaders(v map[string]*string) *DescribeAppAgentTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppAgentTemplatesResponse) SetStatusCode(v int32) *DescribeAppAgentTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponse) SetBody(v *DescribeAppAgentTemplatesResponseBody) *DescribeAppAgentTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeAppAgentTemplatesResponse) Validate() error {
	return dara.Validate(s)
}

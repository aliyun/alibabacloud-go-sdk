// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgentTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgentTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgentTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgentTaskResponseBody) *DescribeAgentTaskResponse
	GetBody() *DescribeAgentTaskResponseBody
}

type DescribeAgentTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgentTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgentTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgentTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgentTaskResponse) GetBody() *DescribeAgentTaskResponseBody {
	return s.Body
}

func (s *DescribeAgentTaskResponse) SetHeaders(v map[string]*string) *DescribeAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgentTaskResponse) SetStatusCode(v int32) *DescribeAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgentTaskResponse) SetBody(v *DescribeAgentTaskResponseBody) *DescribeAgentTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeAgentTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

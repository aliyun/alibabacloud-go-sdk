// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyWorkflowsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComfyWorkflowsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComfyWorkflowsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComfyWorkflowsResponseBody) *DescribeComfyWorkflowsResponse
	GetBody() *DescribeComfyWorkflowsResponseBody
}

type DescribeComfyWorkflowsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComfyWorkflowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComfyWorkflowsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyWorkflowsResponse) GoString() string {
	return s.String()
}

func (s *DescribeComfyWorkflowsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComfyWorkflowsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComfyWorkflowsResponse) GetBody() *DescribeComfyWorkflowsResponseBody {
	return s.Body
}

func (s *DescribeComfyWorkflowsResponse) SetHeaders(v map[string]*string) *DescribeComfyWorkflowsResponse {
	s.Headers = v
	return s
}

func (s *DescribeComfyWorkflowsResponse) SetStatusCode(v int32) *DescribeComfyWorkflowsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComfyWorkflowsResponse) SetBody(v *DescribeComfyWorkflowsResponseBody) *DescribeComfyWorkflowsResponse {
	s.Body = v
	return s
}

func (s *DescribeComfyWorkflowsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

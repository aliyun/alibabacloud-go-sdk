// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWorkflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWorkflowResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWorkflowResponseBody) *DescribeWorkflowResponse
	GetBody() *DescribeWorkflowResponseBody
}

type DescribeWorkflowResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWorkflowResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkflowResponse) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWorkflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWorkflowResponse) GetBody() *DescribeWorkflowResponseBody {
	return s.Body
}

func (s *DescribeWorkflowResponse) SetHeaders(v map[string]*string) *DescribeWorkflowResponse {
	s.Headers = v
	return s
}

func (s *DescribeWorkflowResponse) SetStatusCode(v int32) *DescribeWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWorkflowResponse) SetBody(v *DescribeWorkflowResponseBody) *DescribeWorkflowResponse {
	s.Body = v
	return s
}

func (s *DescribeWorkflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkflowInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWorkflowInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWorkflowInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWorkflowInstanceResponseBody) *DescribeWorkflowInstanceResponse
	GetBody() *DescribeWorkflowInstanceResponseBody
}

type DescribeWorkflowInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWorkflowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWorkflowInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkflowInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeWorkflowInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWorkflowInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWorkflowInstanceResponse) GetBody() *DescribeWorkflowInstanceResponseBody {
	return s.Body
}

func (s *DescribeWorkflowInstanceResponse) SetHeaders(v map[string]*string) *DescribeWorkflowInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeWorkflowInstanceResponse) SetStatusCode(v int32) *DescribeWorkflowInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWorkflowInstanceResponse) SetBody(v *DescribeWorkflowInstanceResponseBody) *DescribeWorkflowInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeWorkflowInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

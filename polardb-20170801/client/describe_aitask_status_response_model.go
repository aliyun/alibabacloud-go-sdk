// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAITaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAITaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAITaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAITaskStatusResponseBody) *DescribeAITaskStatusResponse
	GetBody() *DescribeAITaskStatusResponseBody
}

type DescribeAITaskStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAITaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAITaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAITaskStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAITaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAITaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAITaskStatusResponse) GetBody() *DescribeAITaskStatusResponseBody {
	return s.Body
}

func (s *DescribeAITaskStatusResponse) SetHeaders(v map[string]*string) *DescribeAITaskStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAITaskStatusResponse) SetStatusCode(v int32) *DescribeAITaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAITaskStatusResponse) SetBody(v *DescribeAITaskStatusResponseBody) *DescribeAITaskStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAITaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

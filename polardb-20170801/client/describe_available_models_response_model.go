// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableModelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableModelsResponseBody) *DescribeAvailableModelsResponse
	GetBody() *DescribeAvailableModelsResponseBody
}

type DescribeAvailableModelsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableModelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableModelsResponse) GetBody() *DescribeAvailableModelsResponseBody {
	return s.Body
}

func (s *DescribeAvailableModelsResponse) SetHeaders(v map[string]*string) *DescribeAvailableModelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableModelsResponse) SetStatusCode(v int32) *DescribeAvailableModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableModelsResponse) SetBody(v *DescribeAvailableModelsResponseBody) *DescribeAvailableModelsResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

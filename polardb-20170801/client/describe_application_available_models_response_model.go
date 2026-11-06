// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAvailableModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationAvailableModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationAvailableModelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationAvailableModelsResponseBody) *DescribeApplicationAvailableModelsResponse
	GetBody() *DescribeApplicationAvailableModelsResponseBody
}

type DescribeApplicationAvailableModelsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationAvailableModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationAvailableModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationAvailableModelsResponse) GetBody() *DescribeApplicationAvailableModelsResponseBody {
	return s.Body
}

func (s *DescribeApplicationAvailableModelsResponse) SetHeaders(v map[string]*string) *DescribeApplicationAvailableModelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponse) SetStatusCode(v int32) *DescribeApplicationAvailableModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponse) SetBody(v *DescribeApplicationAvailableModelsResponseBody) *DescribeApplicationAvailableModelsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

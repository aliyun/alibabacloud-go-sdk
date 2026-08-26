// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKnowledgeBaseAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKnowledgeBaseAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKnowledgeBaseAttributeResponseBody) *DescribeKnowledgeBaseAttributeResponse
	GetBody() *DescribeKnowledgeBaseAttributeResponseBody
}

type DescribeKnowledgeBaseAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKnowledgeBaseAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKnowledgeBaseAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKnowledgeBaseAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKnowledgeBaseAttributeResponse) GetBody() *DescribeKnowledgeBaseAttributeResponseBody {
	return s.Body
}

func (s *DescribeKnowledgeBaseAttributeResponse) SetHeaders(v map[string]*string) *DescribeKnowledgeBaseAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponse) SetStatusCode(v int32) *DescribeKnowledgeBaseAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponse) SetBody(v *DescribeKnowledgeBaseAttributeResponseBody) *DescribeKnowledgeBaseAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeKnowledgeBaseAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

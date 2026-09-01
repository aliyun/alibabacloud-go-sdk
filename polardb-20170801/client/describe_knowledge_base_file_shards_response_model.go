// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseFileShardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKnowledgeBaseFileShardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKnowledgeBaseFileShardsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKnowledgeBaseFileShardsResponseBody) *DescribeKnowledgeBaseFileShardsResponse
	GetBody() *DescribeKnowledgeBaseFileShardsResponseBody
}

type DescribeKnowledgeBaseFileShardsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKnowledgeBaseFileShardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKnowledgeBaseFileShardsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFileShardsResponse) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFileShardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKnowledgeBaseFileShardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKnowledgeBaseFileShardsResponse) GetBody() *DescribeKnowledgeBaseFileShardsResponseBody {
	return s.Body
}

func (s *DescribeKnowledgeBaseFileShardsResponse) SetHeaders(v map[string]*string) *DescribeKnowledgeBaseFileShardsResponse {
	s.Headers = v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponse) SetStatusCode(v int32) *DescribeKnowledgeBaseFileShardsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponse) SetBody(v *DescribeKnowledgeBaseFileShardsResponseBody) *DescribeKnowledgeBaseFileShardsResponse {
	s.Body = v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseAnswerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKnowledgeBaseAnswerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKnowledgeBaseAnswerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKnowledgeBaseAnswerResponseBody) *DescribeKnowledgeBaseAnswerResponse
	GetBody() *DescribeKnowledgeBaseAnswerResponseBody
}

type DescribeKnowledgeBaseAnswerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKnowledgeBaseAnswerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKnowledgeBaseAnswerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseAnswerResponse) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseAnswerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKnowledgeBaseAnswerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKnowledgeBaseAnswerResponse) GetBody() *DescribeKnowledgeBaseAnswerResponseBody {
	return s.Body
}

func (s *DescribeKnowledgeBaseAnswerResponse) SetHeaders(v map[string]*string) *DescribeKnowledgeBaseAnswerResponse {
	s.Headers = v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponse) SetStatusCode(v int32) *DescribeKnowledgeBaseAnswerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponse) SetBody(v *DescribeKnowledgeBaseAnswerResponseBody) *DescribeKnowledgeBaseAnswerResponse {
	s.Body = v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

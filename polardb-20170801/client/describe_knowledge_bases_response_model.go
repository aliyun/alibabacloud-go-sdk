// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKnowledgeBasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKnowledgeBasesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKnowledgeBasesResponseBody) *DescribeKnowledgeBasesResponse
	GetBody() *DescribeKnowledgeBasesResponseBody
}

type DescribeKnowledgeBasesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKnowledgeBasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKnowledgeBasesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKnowledgeBasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKnowledgeBasesResponse) GetBody() *DescribeKnowledgeBasesResponseBody {
	return s.Body
}

func (s *DescribeKnowledgeBasesResponse) SetHeaders(v map[string]*string) *DescribeKnowledgeBasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeKnowledgeBasesResponse) SetStatusCode(v int32) *DescribeKnowledgeBasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKnowledgeBasesResponse) SetBody(v *DescribeKnowledgeBasesResponseBody) *DescribeKnowledgeBasesResponse {
	s.Body = v
	return s
}

func (s *DescribeKnowledgeBasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

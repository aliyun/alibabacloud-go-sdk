// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKnowledgeBaseFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKnowledgeBaseFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKnowledgeBaseFilesResponseBody) *DescribeKnowledgeBaseFilesResponse
	GetBody() *DescribeKnowledgeBaseFilesResponseBody
}

type DescribeKnowledgeBaseFilesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKnowledgeBaseFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKnowledgeBaseFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKnowledgeBaseFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKnowledgeBaseFilesResponse) GetBody() *DescribeKnowledgeBaseFilesResponseBody {
	return s.Body
}

func (s *DescribeKnowledgeBaseFilesResponse) SetHeaders(v map[string]*string) *DescribeKnowledgeBaseFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponse) SetStatusCode(v int32) *DescribeKnowledgeBaseFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponse) SetBody(v *DescribeKnowledgeBaseFilesResponseBody) *DescribeKnowledgeBaseFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

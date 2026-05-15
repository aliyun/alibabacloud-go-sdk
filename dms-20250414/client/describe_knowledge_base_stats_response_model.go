// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKnowledgeBaseStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKnowledgeBaseStatsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKnowledgeBaseStatsResponseBody) *DescribeKnowledgeBaseStatsResponse
	GetBody() *DescribeKnowledgeBaseStatsResponseBody
}

type DescribeKnowledgeBaseStatsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKnowledgeBaseStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKnowledgeBaseStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseStatsResponse) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKnowledgeBaseStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKnowledgeBaseStatsResponse) GetBody() *DescribeKnowledgeBaseStatsResponseBody {
	return s.Body
}

func (s *DescribeKnowledgeBaseStatsResponse) SetHeaders(v map[string]*string) *DescribeKnowledgeBaseStatsResponse {
	s.Headers = v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponse) SetStatusCode(v int32) *DescribeKnowledgeBaseStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponse) SetBody(v *DescribeKnowledgeBaseStatsResponseBody) *DescribeKnowledgeBaseStatsResponse {
	s.Body = v
	return s
}

func (s *DescribeKnowledgeBaseStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeBasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeBasesResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeBasesResult) *ListKnowledgeBasesResponse
	GetBody() *ListKnowledgeBasesResult
}

type ListKnowledgeBasesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeBasesResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeBasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeBasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeBasesResponse) GetBody() *ListKnowledgeBasesResult {
	return s.Body
}

func (s *ListKnowledgeBasesResponse) SetHeaders(v map[string]*string) *ListKnowledgeBasesResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeBasesResponse) SetStatusCode(v int32) *ListKnowledgeBasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeBasesResponse) SetBody(v *ListKnowledgeBasesResult) *ListKnowledgeBasesResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeBasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

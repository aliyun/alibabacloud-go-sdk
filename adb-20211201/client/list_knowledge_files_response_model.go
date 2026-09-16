// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeFilesResponseBody) *ListKnowledgeFilesResponse
	GetBody() *ListKnowledgeFilesResponseBody
}

type ListKnowledgeFilesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeFilesResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeFilesResponse) GetBody() *ListKnowledgeFilesResponseBody {
	return s.Body
}

func (s *ListKnowledgeFilesResponse) SetHeaders(v map[string]*string) *ListKnowledgeFilesResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeFilesResponse) SetStatusCode(v int32) *ListKnowledgeFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeFilesResponse) SetBody(v *ListKnowledgeFilesResponseBody) *ListKnowledgeFilesResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

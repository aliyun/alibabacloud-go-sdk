// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseBuildTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeBaseBuildTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeBaseBuildTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeBaseBuildTasksResponseBody) *ListKnowledgeBaseBuildTasksResponse
	GetBody() *ListKnowledgeBaseBuildTasksResponseBody
}

type ListKnowledgeBaseBuildTasksResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeBaseBuildTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeBaseBuildTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTasksResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeBaseBuildTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeBaseBuildTasksResponse) GetBody() *ListKnowledgeBaseBuildTasksResponseBody {
	return s.Body
}

func (s *ListKnowledgeBaseBuildTasksResponse) SetHeaders(v map[string]*string) *ListKnowledgeBaseBuildTasksResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponse) SetStatusCode(v int32) *ListKnowledgeBaseBuildTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponse) SetBody(v *ListKnowledgeBaseBuildTasksResponseBody) *ListKnowledgeBaseBuildTasksResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseBuildTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeBaseBuildTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeBaseBuildTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeBaseBuildTaskInstancesResponseBody) *ListKnowledgeBaseBuildTaskInstancesResponse
	GetBody() *ListKnowledgeBaseBuildTaskInstancesResponseBody
}

type ListKnowledgeBaseBuildTaskInstancesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeBaseBuildTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeBaseBuildTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponse) GetBody() *ListKnowledgeBaseBuildTaskInstancesResponseBody {
	return s.Body
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponse) SetHeaders(v map[string]*string) *ListKnowledgeBaseBuildTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponse) SetStatusCode(v int32) *ListKnowledgeBaseBuildTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponse) SetBody(v *ListKnowledgeBaseBuildTaskInstancesResponseBody) *ListKnowledgeBaseBuildTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

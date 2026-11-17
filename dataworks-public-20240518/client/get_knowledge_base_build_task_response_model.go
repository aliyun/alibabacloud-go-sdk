// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseBuildTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeBaseBuildTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeBaseBuildTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeBaseBuildTaskResponseBody) *GetKnowledgeBaseBuildTaskResponse
	GetBody() *GetKnowledgeBaseBuildTaskResponseBody
}

type GetKnowledgeBaseBuildTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeBaseBuildTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeBaseBuildTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeBaseBuildTaskResponse) GetBody() *GetKnowledgeBaseBuildTaskResponseBody {
	return s.Body
}

func (s *GetKnowledgeBaseBuildTaskResponse) SetHeaders(v map[string]*string) *GetKnowledgeBaseBuildTaskResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponse) SetStatusCode(v int32) *GetKnowledgeBaseBuildTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponse) SetBody(v *GetKnowledgeBaseBuildTaskResponseBody) *GetKnowledgeBaseBuildTaskResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

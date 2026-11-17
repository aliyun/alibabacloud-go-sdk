// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseBuildTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeBaseBuildTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeBaseBuildTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeBaseBuildTaskResponseBody) *CreateKnowledgeBaseBuildTaskResponse
	GetBody() *CreateKnowledgeBaseBuildTaskResponseBody
}

type CreateKnowledgeBaseBuildTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeBaseBuildTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeBaseBuildTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeBaseBuildTaskResponse) GetBody() *CreateKnowledgeBaseBuildTaskResponseBody {
	return s.Body
}

func (s *CreateKnowledgeBaseBuildTaskResponse) SetHeaders(v map[string]*string) *CreateKnowledgeBaseBuildTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponse) SetStatusCode(v int32) *CreateKnowledgeBaseBuildTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponse) SetBody(v *CreateKnowledgeBaseBuildTaskResponseBody) *CreateKnowledgeBaseBuildTaskResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

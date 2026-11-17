// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseBuildTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeBaseBuildTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeBaseBuildTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKnowledgeBaseBuildTaskResponseBody) *UpdateKnowledgeBaseBuildTaskResponse
	GetBody() *UpdateKnowledgeBaseBuildTaskResponseBody
}

type UpdateKnowledgeBaseBuildTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeBaseBuildTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeBaseBuildTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeBaseBuildTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeBaseBuildTaskResponse) GetBody() *UpdateKnowledgeBaseBuildTaskResponseBody {
	return s.Body
}

func (s *UpdateKnowledgeBaseBuildTaskResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeBaseBuildTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponse) SetStatusCode(v int32) *UpdateKnowledgeBaseBuildTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponse) SetBody(v *UpdateKnowledgeBaseBuildTaskResponseBody) *UpdateKnowledgeBaseBuildTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

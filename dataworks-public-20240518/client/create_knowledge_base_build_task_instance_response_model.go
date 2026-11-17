// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseBuildTaskInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeBaseBuildTaskInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeBaseBuildTaskInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeBaseBuildTaskInstanceResponseBody) *CreateKnowledgeBaseBuildTaskInstanceResponse
	GetBody() *CreateKnowledgeBaseBuildTaskInstanceResponseBody
}

type CreateKnowledgeBaseBuildTaskInstanceResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeBaseBuildTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponse) GetBody() *CreateKnowledgeBaseBuildTaskInstanceResponseBody {
	return s.Body
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponse) SetHeaders(v map[string]*string) *CreateKnowledgeBaseBuildTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponse) SetStatusCode(v int32) *CreateKnowledgeBaseBuildTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponse) SetBody(v *CreateKnowledgeBaseBuildTaskInstanceResponseBody) *CreateKnowledgeBaseBuildTaskInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

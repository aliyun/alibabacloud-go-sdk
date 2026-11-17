// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseBuildTaskInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeBaseBuildTaskInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeBaseBuildTaskInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeBaseBuildTaskInstanceResponseBody) *GetKnowledgeBaseBuildTaskInstanceResponse
	GetBody() *GetKnowledgeBaseBuildTaskInstanceResponseBody
}

type GetKnowledgeBaseBuildTaskInstanceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeBaseBuildTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponse) GetBody() *GetKnowledgeBaseBuildTaskInstanceResponseBody {
	return s.Body
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponse) SetHeaders(v map[string]*string) *GetKnowledgeBaseBuildTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponse) SetStatusCode(v int32) *GetKnowledgeBaseBuildTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponse) SetBody(v *GetKnowledgeBaseBuildTaskInstanceResponseBody) *GetKnowledgeBaseBuildTaskInstanceResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

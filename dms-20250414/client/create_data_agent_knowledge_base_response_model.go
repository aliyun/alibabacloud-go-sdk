// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataAgentKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataAgentKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataAgentKnowledgeBaseResponseBody) *CreateDataAgentKnowledgeBaseResponse
	GetBody() *CreateDataAgentKnowledgeBaseResponseBody
}

type CreateDataAgentKnowledgeBaseResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataAgentKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataAgentKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *CreateDataAgentKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataAgentKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataAgentKnowledgeBaseResponse) GetBody() *CreateDataAgentKnowledgeBaseResponseBody {
	return s.Body
}

func (s *CreateDataAgentKnowledgeBaseResponse) SetHeaders(v map[string]*string) *CreateDataAgentKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *CreateDataAgentKnowledgeBaseResponse) SetStatusCode(v int32) *CreateDataAgentKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataAgentKnowledgeBaseResponse) SetBody(v *CreateDataAgentKnowledgeBaseResponseBody) *CreateDataAgentKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *CreateDataAgentKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

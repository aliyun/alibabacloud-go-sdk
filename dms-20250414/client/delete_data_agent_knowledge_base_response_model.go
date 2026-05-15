// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataAgentKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataAgentKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataAgentKnowledgeBaseResponseBody) *DeleteDataAgentKnowledgeBaseResponse
	GetBody() *DeleteDataAgentKnowledgeBaseResponseBody
}

type DeleteDataAgentKnowledgeBaseResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataAgentKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataAgentKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataAgentKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataAgentKnowledgeBaseResponse) GetBody() *DeleteDataAgentKnowledgeBaseResponseBody {
	return s.Body
}

func (s *DeleteDataAgentKnowledgeBaseResponse) SetHeaders(v map[string]*string) *DeleteDataAgentKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseResponse) SetStatusCode(v int32) *DeleteDataAgentKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseResponse) SetBody(v *DeleteDataAgentKnowledgeBaseResponseBody) *DeleteDataAgentKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *DeleteDataAgentKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

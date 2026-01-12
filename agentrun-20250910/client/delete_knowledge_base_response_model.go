// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *KnowledgeBaseResult) *DeleteKnowledgeBaseResponse
	GetBody() *KnowledgeBaseResult
}

type DeleteKnowledgeBaseResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KnowledgeBaseResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKnowledgeBaseResponse) GetBody() *KnowledgeBaseResult {
	return s.Body
}

func (s *DeleteKnowledgeBaseResponse) SetHeaders(v map[string]*string) *DeleteKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteKnowledgeBaseResponse) SetStatusCode(v int32) *DeleteKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKnowledgeBaseResponse) SetBody(v *KnowledgeBaseResult) *DeleteKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *DeleteKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

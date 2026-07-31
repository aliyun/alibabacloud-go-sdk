// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeRecallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeRecallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeRecallResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeRecallResponseBody) *GetKnowledgeRecallResponse
	GetBody() *GetKnowledgeRecallResponseBody
}

type GetKnowledgeRecallResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeRecallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeRecallResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeRecallResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeRecallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeRecallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeRecallResponse) GetBody() *GetKnowledgeRecallResponseBody {
	return s.Body
}

func (s *GetKnowledgeRecallResponse) SetHeaders(v map[string]*string) *GetKnowledgeRecallResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeRecallResponse) SetStatusCode(v int32) *GetKnowledgeRecallResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeRecallResponse) SetBody(v *GetKnowledgeRecallResponseBody) *GetKnowledgeRecallResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeRecallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

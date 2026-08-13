// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSourceContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeBaseSourceContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeBaseSourceContentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKnowledgeBaseSourceContentResponseBody) *UpdateKnowledgeBaseSourceContentResponse
	GetBody() *UpdateKnowledgeBaseSourceContentResponseBody
}

type UpdateKnowledgeBaseSourceContentResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeBaseSourceContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeBaseSourceContentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSourceContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSourceContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeBaseSourceContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeBaseSourceContentResponse) GetBody() *UpdateKnowledgeBaseSourceContentResponseBody {
	return s.Body
}

func (s *UpdateKnowledgeBaseSourceContentResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeBaseSourceContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentResponse) SetStatusCode(v int32) *UpdateKnowledgeBaseSourceContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentResponse) SetBody(v *UpdateKnowledgeBaseSourceContentResponseBody) *UpdateKnowledgeBaseSourceContentResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

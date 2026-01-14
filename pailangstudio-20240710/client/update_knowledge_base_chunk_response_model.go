// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseChunkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeBaseChunkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeBaseChunkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKnowledgeBaseChunkResponseBody) *UpdateKnowledgeBaseChunkResponse
	GetBody() *UpdateKnowledgeBaseChunkResponseBody
}

type UpdateKnowledgeBaseChunkResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeBaseChunkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeBaseChunkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseChunkResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseChunkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeBaseChunkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeBaseChunkResponse) GetBody() *UpdateKnowledgeBaseChunkResponseBody {
	return s.Body
}

func (s *UpdateKnowledgeBaseChunkResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeBaseChunkResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeBaseChunkResponse) SetStatusCode(v int32) *UpdateKnowledgeBaseChunkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeBaseChunkResponse) SetBody(v *UpdateKnowledgeBaseChunkResponseBody) *UpdateKnowledgeBaseChunkResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeBaseChunkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

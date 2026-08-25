// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiModalKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMultiModalKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMultiModalKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMultiModalKnowledgeBaseResponseBody) *DeleteMultiModalKnowledgeBaseResponse
	GetBody() *DeleteMultiModalKnowledgeBaseResponseBody
}

type DeleteMultiModalKnowledgeBaseResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultiModalKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultiModalKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiModalKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultiModalKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMultiModalKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMultiModalKnowledgeBaseResponse) GetBody() *DeleteMultiModalKnowledgeBaseResponseBody {
	return s.Body
}

func (s *DeleteMultiModalKnowledgeBaseResponse) SetHeaders(v map[string]*string) *DeleteMultiModalKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultiModalKnowledgeBaseResponse) SetStatusCode(v int32) *DeleteMultiModalKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultiModalKnowledgeBaseResponse) SetBody(v *DeleteMultiModalKnowledgeBaseResponseBody) *DeleteMultiModalKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *DeleteMultiModalKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

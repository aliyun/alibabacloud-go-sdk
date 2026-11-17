// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeKnowledgeBaseRetrievalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeKnowledgeBaseRetrievalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeKnowledgeBaseRetrievalResponse
	GetStatusCode() *int32
	SetBody(v *RevokeKnowledgeBaseRetrievalResponseBody) *RevokeKnowledgeBaseRetrievalResponse
	GetBody() *RevokeKnowledgeBaseRetrievalResponseBody
}

type RevokeKnowledgeBaseRetrievalResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeKnowledgeBaseRetrievalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeKnowledgeBaseRetrievalResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeKnowledgeBaseRetrievalResponse) GoString() string {
	return s.String()
}

func (s *RevokeKnowledgeBaseRetrievalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeKnowledgeBaseRetrievalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeKnowledgeBaseRetrievalResponse) GetBody() *RevokeKnowledgeBaseRetrievalResponseBody {
	return s.Body
}

func (s *RevokeKnowledgeBaseRetrievalResponse) SetHeaders(v map[string]*string) *RevokeKnowledgeBaseRetrievalResponse {
	s.Headers = v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalResponse) SetStatusCode(v int32) *RevokeKnowledgeBaseRetrievalResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalResponse) SetBody(v *RevokeKnowledgeBaseRetrievalResponseBody) *RevokeKnowledgeBaseRetrievalResponse {
	s.Body = v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

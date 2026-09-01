// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKnowledgeBaseFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKnowledgeBaseFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKnowledgeBaseFileResponseBody) *DeleteKnowledgeBaseFileResponse
	GetBody() *DeleteKnowledgeBaseFileResponseBody
}

type DeleteKnowledgeBaseFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKnowledgeBaseFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKnowledgeBaseFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKnowledgeBaseFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKnowledgeBaseFileResponse) GetBody() *DeleteKnowledgeBaseFileResponseBody {
	return s.Body
}

func (s *DeleteKnowledgeBaseFileResponse) SetHeaders(v map[string]*string) *DeleteKnowledgeBaseFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteKnowledgeBaseFileResponse) SetStatusCode(v int32) *DeleteKnowledgeBaseFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKnowledgeBaseFileResponse) SetBody(v *DeleteKnowledgeBaseFileResponseBody) *DeleteKnowledgeBaseFileResponse {
	s.Body = v
	return s
}

func (s *DeleteKnowledgeBaseFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

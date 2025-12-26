// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKnowledgeBaseJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKnowledgeBaseJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKnowledgeBaseJobResponseBody) *DeleteKnowledgeBaseJobResponse
	GetBody() *DeleteKnowledgeBaseJobResponseBody
}

type DeleteKnowledgeBaseJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKnowledgeBaseJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKnowledgeBaseJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKnowledgeBaseJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKnowledgeBaseJobResponse) GetBody() *DeleteKnowledgeBaseJobResponseBody {
	return s.Body
}

func (s *DeleteKnowledgeBaseJobResponse) SetHeaders(v map[string]*string) *DeleteKnowledgeBaseJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteKnowledgeBaseJobResponse) SetStatusCode(v int32) *DeleteKnowledgeBaseJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKnowledgeBaseJobResponse) SetBody(v *DeleteKnowledgeBaseJobResponseBody) *DeleteKnowledgeBaseJobResponse {
	s.Body = v
	return s
}

func (s *DeleteKnowledgeBaseJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKnowledgeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKnowledgeSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKnowledgeSourceResponseBody) *DeleteKnowledgeSourceResponse
	GetBody() *DeleteKnowledgeSourceResponseBody
}

type DeleteKnowledgeSourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKnowledgeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKnowledgeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKnowledgeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKnowledgeSourceResponse) GetBody() *DeleteKnowledgeSourceResponseBody {
	return s.Body
}

func (s *DeleteKnowledgeSourceResponse) SetHeaders(v map[string]*string) *DeleteKnowledgeSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteKnowledgeSourceResponse) SetStatusCode(v int32) *DeleteKnowledgeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKnowledgeSourceResponse) SetBody(v *DeleteKnowledgeSourceResponseBody) *DeleteKnowledgeSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteKnowledgeSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

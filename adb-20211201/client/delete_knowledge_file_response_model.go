// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKnowledgeFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKnowledgeFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKnowledgeFileResponseBody) *DeleteKnowledgeFileResponse
	GetBody() *DeleteKnowledgeFileResponseBody
}

type DeleteKnowledgeFileResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKnowledgeFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKnowledgeFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKnowledgeFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKnowledgeFileResponse) GetBody() *DeleteKnowledgeFileResponseBody {
	return s.Body
}

func (s *DeleteKnowledgeFileResponse) SetHeaders(v map[string]*string) *DeleteKnowledgeFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteKnowledgeFileResponse) SetStatusCode(v int32) *DeleteKnowledgeFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKnowledgeFileResponse) SetBody(v *DeleteKnowledgeFileResponseBody) *DeleteKnowledgeFileResponse {
	s.Body = v
	return s
}

func (s *DeleteKnowledgeFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

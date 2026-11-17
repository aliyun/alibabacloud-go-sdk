// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseBuildTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKnowledgeBaseBuildTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKnowledgeBaseBuildTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKnowledgeBaseBuildTaskResponseBody) *DeleteKnowledgeBaseBuildTaskResponse
	GetBody() *DeleteKnowledgeBaseBuildTaskResponseBody
}

type DeleteKnowledgeBaseBuildTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKnowledgeBaseBuildTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKnowledgeBaseBuildTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseBuildTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseBuildTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKnowledgeBaseBuildTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKnowledgeBaseBuildTaskResponse) GetBody() *DeleteKnowledgeBaseBuildTaskResponseBody {
	return s.Body
}

func (s *DeleteKnowledgeBaseBuildTaskResponse) SetHeaders(v map[string]*string) *DeleteKnowledgeBaseBuildTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskResponse) SetStatusCode(v int32) *DeleteKnowledgeBaseBuildTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskResponse) SetBody(v *DeleteKnowledgeBaseBuildTaskResponseBody) *DeleteKnowledgeBaseBuildTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSourceTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeBaseSourceTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeBaseSourceTagsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKnowledgeBaseSourceTagsResponseBody) *UpdateKnowledgeBaseSourceTagsResponse
	GetBody() *UpdateKnowledgeBaseSourceTagsResponseBody
}

type UpdateKnowledgeBaseSourceTagsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeBaseSourceTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeBaseSourceTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSourceTagsResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSourceTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeBaseSourceTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeBaseSourceTagsResponse) GetBody() *UpdateKnowledgeBaseSourceTagsResponseBody {
	return s.Body
}

func (s *UpdateKnowledgeBaseSourceTagsResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeBaseSourceTagsResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsResponse) SetStatusCode(v int32) *UpdateKnowledgeBaseSourceTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsResponse) SetBody(v *UpdateKnowledgeBaseSourceTagsResponseBody) *UpdateKnowledgeBaseSourceTagsResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

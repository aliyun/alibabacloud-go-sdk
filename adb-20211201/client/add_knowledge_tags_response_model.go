// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddKnowledgeTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddKnowledgeTagsResponse
	GetStatusCode() *int32
	SetBody(v *AddKnowledgeTagsResponseBody) *AddKnowledgeTagsResponse
	GetBody() *AddKnowledgeTagsResponseBody
}

type AddKnowledgeTagsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddKnowledgeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddKnowledgeTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeTagsResponse) GoString() string {
	return s.String()
}

func (s *AddKnowledgeTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddKnowledgeTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddKnowledgeTagsResponse) GetBody() *AddKnowledgeTagsResponseBody {
	return s.Body
}

func (s *AddKnowledgeTagsResponse) SetHeaders(v map[string]*string) *AddKnowledgeTagsResponse {
	s.Headers = v
	return s
}

func (s *AddKnowledgeTagsResponse) SetStatusCode(v int32) *AddKnowledgeTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddKnowledgeTagsResponse) SetBody(v *AddKnowledgeTagsResponseBody) *AddKnowledgeTagsResponse {
	s.Body = v
	return s
}

func (s *AddKnowledgeTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

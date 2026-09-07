// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveKnowledgeTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveKnowledgeTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveKnowledgeTagsResponse
	GetStatusCode() *int32
	SetBody(v *RemoveKnowledgeTagsResponseBody) *RemoveKnowledgeTagsResponse
	GetBody() *RemoveKnowledgeTagsResponseBody
}

type RemoveKnowledgeTagsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveKnowledgeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveKnowledgeTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeTagsResponse) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveKnowledgeTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveKnowledgeTagsResponse) GetBody() *RemoveKnowledgeTagsResponseBody {
	return s.Body
}

func (s *RemoveKnowledgeTagsResponse) SetHeaders(v map[string]*string) *RemoveKnowledgeTagsResponse {
	s.Headers = v
	return s
}

func (s *RemoveKnowledgeTagsResponse) SetStatusCode(v int32) *RemoveKnowledgeTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveKnowledgeTagsResponse) SetBody(v *RemoveKnowledgeTagsResponseBody) *RemoveKnowledgeTagsResponse {
	s.Body = v
	return s
}

func (s *RemoveKnowledgeTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

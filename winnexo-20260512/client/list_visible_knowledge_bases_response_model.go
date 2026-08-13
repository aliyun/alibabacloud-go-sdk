// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVisibleKnowledgeBasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVisibleKnowledgeBasesResponse
	GetStatusCode() *int32
	SetBody(v *ListVisibleKnowledgeBasesResponseBody) *ListVisibleKnowledgeBasesResponse
	GetBody() *ListVisibleKnowledgeBasesResponseBody
}

type ListVisibleKnowledgeBasesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVisibleKnowledgeBasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVisibleKnowledgeBasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBasesResponse) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVisibleKnowledgeBasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVisibleKnowledgeBasesResponse) GetBody() *ListVisibleKnowledgeBasesResponseBody {
	return s.Body
}

func (s *ListVisibleKnowledgeBasesResponse) SetHeaders(v map[string]*string) *ListVisibleKnowledgeBasesResponse {
	s.Headers = v
	return s
}

func (s *ListVisibleKnowledgeBasesResponse) SetStatusCode(v int32) *ListVisibleKnowledgeBasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVisibleKnowledgeBasesResponse) SetBody(v *ListVisibleKnowledgeBasesResponseBody) *ListVisibleKnowledgeBasesResponse {
	s.Body = v
	return s
}

func (s *ListVisibleKnowledgeBasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

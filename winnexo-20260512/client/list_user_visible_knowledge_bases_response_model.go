// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVisibleKnowledgeBasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserVisibleKnowledgeBasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserVisibleKnowledgeBasesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserVisibleKnowledgeBasesResponseBody) *ListUserVisibleKnowledgeBasesResponse
	GetBody() *ListUserVisibleKnowledgeBasesResponseBody
}

type ListUserVisibleKnowledgeBasesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserVisibleKnowledgeBasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserVisibleKnowledgeBasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBasesResponse) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserVisibleKnowledgeBasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserVisibleKnowledgeBasesResponse) GetBody() *ListUserVisibleKnowledgeBasesResponseBody {
	return s.Body
}

func (s *ListUserVisibleKnowledgeBasesResponse) SetHeaders(v map[string]*string) *ListUserVisibleKnowledgeBasesResponse {
	s.Headers = v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponse) SetStatusCode(v int32) *ListUserVisibleKnowledgeBasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponse) SetBody(v *ListUserVisibleKnowledgeBasesResponseBody) *ListUserVisibleKnowledgeBasesResponse {
	s.Body = v
	return s
}

func (s *ListUserVisibleKnowledgeBasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

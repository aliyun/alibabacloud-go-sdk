// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaKnowledgeBasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLumaKnowledgeBasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLumaKnowledgeBasesResponse
	GetStatusCode() *int32
	SetBody(v *ListLumaKnowledgeBasesResponseBody) *ListLumaKnowledgeBasesResponse
	GetBody() *ListLumaKnowledgeBasesResponseBody
}

type ListLumaKnowledgeBasesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLumaKnowledgeBasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLumaKnowledgeBasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLumaKnowledgeBasesResponse) GoString() string {
	return s.String()
}

func (s *ListLumaKnowledgeBasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLumaKnowledgeBasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLumaKnowledgeBasesResponse) GetBody() *ListLumaKnowledgeBasesResponseBody {
	return s.Body
}

func (s *ListLumaKnowledgeBasesResponse) SetHeaders(v map[string]*string) *ListLumaKnowledgeBasesResponse {
	s.Headers = v
	return s
}

func (s *ListLumaKnowledgeBasesResponse) SetStatusCode(v int32) *ListLumaKnowledgeBasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLumaKnowledgeBasesResponse) SetBody(v *ListLumaKnowledgeBasesResponseBody) *ListLumaKnowledgeBasesResponse {
	s.Body = v
	return s
}

func (s *ListLumaKnowledgeBasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

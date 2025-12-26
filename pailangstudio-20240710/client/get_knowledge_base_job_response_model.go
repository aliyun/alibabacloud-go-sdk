// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeBaseJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeBaseJobResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeBaseJobResponseBody) *GetKnowledgeBaseJobResponse
	GetBody() *GetKnowledgeBaseJobResponseBody
}

type GetKnowledgeBaseJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeBaseJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeBaseJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseJobResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeBaseJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeBaseJobResponse) GetBody() *GetKnowledgeBaseJobResponseBody {
	return s.Body
}

func (s *GetKnowledgeBaseJobResponse) SetHeaders(v map[string]*string) *GetKnowledgeBaseJobResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeBaseJobResponse) SetStatusCode(v int32) *GetKnowledgeBaseJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeBaseJobResponse) SetBody(v *GetKnowledgeBaseJobResponseBody) *GetKnowledgeBaseJobResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeBaseJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

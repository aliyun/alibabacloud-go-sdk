// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeBaseJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeBaseJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeBaseJobResponseBody) *CreateKnowledgeBaseJobResponse
	GetBody() *CreateKnowledgeBaseJobResponseBody
}

type CreateKnowledgeBaseJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeBaseJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBaseJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseJobResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeBaseJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeBaseJobResponse) GetBody() *CreateKnowledgeBaseJobResponseBody {
	return s.Body
}

func (s *CreateKnowledgeBaseJobResponse) SetHeaders(v map[string]*string) *CreateKnowledgeBaseJobResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeBaseJobResponse) SetStatusCode(v int32) *CreateKnowledgeBaseJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeBaseJobResponse) SetBody(v *CreateKnowledgeBaseJobResponseBody) *CreateKnowledgeBaseJobResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBaseJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeBaseJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeBaseJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKnowledgeBaseJobResponseBody) *UpdateKnowledgeBaseJobResponse
	GetBody() *UpdateKnowledgeBaseJobResponseBody
}

type UpdateKnowledgeBaseJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeBaseJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeBaseJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeBaseJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeBaseJobResponse) GetBody() *UpdateKnowledgeBaseJobResponseBody {
	return s.Body
}

func (s *UpdateKnowledgeBaseJobResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeBaseJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeBaseJobResponse) SetStatusCode(v int32) *UpdateKnowledgeBaseJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeBaseJobResponse) SetBody(v *UpdateKnowledgeBaseJobResponseBody) *UpdateKnowledgeBaseJobResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeBaseJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

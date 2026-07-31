// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddKnowledgeFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddKnowledgeFileResponse
	GetStatusCode() *int32
	SetBody(v *AddKnowledgeFileResponseBody) *AddKnowledgeFileResponse
	GetBody() *AddKnowledgeFileResponseBody
}

type AddKnowledgeFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddKnowledgeFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddKnowledgeFileResponse) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeFileResponse) GoString() string {
	return s.String()
}

func (s *AddKnowledgeFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddKnowledgeFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddKnowledgeFileResponse) GetBody() *AddKnowledgeFileResponseBody {
	return s.Body
}

func (s *AddKnowledgeFileResponse) SetHeaders(v map[string]*string) *AddKnowledgeFileResponse {
	s.Headers = v
	return s
}

func (s *AddKnowledgeFileResponse) SetStatusCode(v int32) *AddKnowledgeFileResponse {
	s.StatusCode = &v
	return s
}

func (s *AddKnowledgeFileResponse) SetBody(v *AddKnowledgeFileResponseBody) *AddKnowledgeFileResponse {
	s.Body = v
	return s
}

func (s *AddKnowledgeFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

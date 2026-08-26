// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeSpaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeSpaceResponseBody) *CreateKnowledgeSpaceResponse
	GetBody() *CreateKnowledgeSpaceResponseBody
}

type CreateKnowledgeSpaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSpaceResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeSpaceResponse) GetBody() *CreateKnowledgeSpaceResponseBody {
	return s.Body
}

func (s *CreateKnowledgeSpaceResponse) SetHeaders(v map[string]*string) *CreateKnowledgeSpaceResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeSpaceResponse) SetStatusCode(v int32) *CreateKnowledgeSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeSpaceResponse) SetBody(v *CreateKnowledgeSpaceResponseBody) *CreateKnowledgeSpaceResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

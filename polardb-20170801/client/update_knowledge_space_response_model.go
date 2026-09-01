// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeSpaceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKnowledgeSpaceResponseBody) *UpdateKnowledgeSpaceResponse
	GetBody() *UpdateKnowledgeSpaceResponseBody
}

type UpdateKnowledgeSpaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSpaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeSpaceResponse) GetBody() *UpdateKnowledgeSpaceResponseBody {
	return s.Body
}

func (s *UpdateKnowledgeSpaceResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeSpaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeSpaceResponse) SetStatusCode(v int32) *UpdateKnowledgeSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeSpaceResponse) SetBody(v *UpdateKnowledgeSpaceResponseBody) *UpdateKnowledgeSpaceResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

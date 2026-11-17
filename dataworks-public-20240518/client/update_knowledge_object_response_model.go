// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeObjectResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKnowledgeObjectResponseBody) *UpdateKnowledgeObjectResponse
	GetBody() *UpdateKnowledgeObjectResponseBody
}

type UpdateKnowledgeObjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeObjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeObjectResponse) GetBody() *UpdateKnowledgeObjectResponseBody {
	return s.Body
}

func (s *UpdateKnowledgeObjectResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeObjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeObjectResponse) SetStatusCode(v int32) *UpdateKnowledgeObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeObjectResponse) SetBody(v *UpdateKnowledgeObjectResponseBody) *UpdateKnowledgeObjectResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeObjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeObjectResponseBody) *CreateKnowledgeObjectResponse
	GetBody() *CreateKnowledgeObjectResponseBody
}

type CreateKnowledgeObjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeObjectResponse) GetBody() *CreateKnowledgeObjectResponseBody {
	return s.Body
}

func (s *CreateKnowledgeObjectResponse) SetHeaders(v map[string]*string) *CreateKnowledgeObjectResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeObjectResponse) SetStatusCode(v int32) *CreateKnowledgeObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeObjectResponse) SetBody(v *CreateKnowledgeObjectResponseBody) *CreateKnowledgeObjectResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

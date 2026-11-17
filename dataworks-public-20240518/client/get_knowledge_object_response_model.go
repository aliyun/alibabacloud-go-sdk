// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeObjectResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeObjectResponseBody) *GetKnowledgeObjectResponse
	GetBody() *GetKnowledgeObjectResponseBody
}

type GetKnowledgeObjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeObjectResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeObjectResponse) GetBody() *GetKnowledgeObjectResponseBody {
	return s.Body
}

func (s *GetKnowledgeObjectResponse) SetHeaders(v map[string]*string) *GetKnowledgeObjectResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeObjectResponse) SetStatusCode(v int32) *GetKnowledgeObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeObjectResponse) SetBody(v *GetKnowledgeObjectResponseBody) *GetKnowledgeObjectResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

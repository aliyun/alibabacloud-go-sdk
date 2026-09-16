// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLumaKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLumaKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *GetLumaKnowledgeBaseResponseBody) *GetLumaKnowledgeBaseResponse
	GetBody() *GetLumaKnowledgeBaseResponseBody
}

type GetLumaKnowledgeBaseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLumaKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLumaKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLumaKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *GetLumaKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLumaKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLumaKnowledgeBaseResponse) GetBody() *GetLumaKnowledgeBaseResponseBody {
	return s.Body
}

func (s *GetLumaKnowledgeBaseResponse) SetHeaders(v map[string]*string) *GetLumaKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *GetLumaKnowledgeBaseResponse) SetStatusCode(v int32) *GetLumaKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLumaKnowledgeBaseResponse) SetBody(v *GetLumaKnowledgeBaseResponseBody) *GetLumaKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *GetLumaKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

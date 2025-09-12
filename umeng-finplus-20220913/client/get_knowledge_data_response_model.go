// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeDataResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgeDataResponseBody) *GetKnowledgeDataResponse
	GetBody() *GetKnowledgeDataResponseBody
}

type GetKnowledgeDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeDataResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeDataResponse) GetBody() *GetKnowledgeDataResponseBody {
	return s.Body
}

func (s *GetKnowledgeDataResponse) SetHeaders(v map[string]*string) *GetKnowledgeDataResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeDataResponse) SetStatusCode(v int32) *GetKnowledgeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeDataResponse) SetBody(v *GetKnowledgeDataResponseBody) *GetKnowledgeDataResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeDataResponse) Validate() error {
	return dara.Validate(s)
}

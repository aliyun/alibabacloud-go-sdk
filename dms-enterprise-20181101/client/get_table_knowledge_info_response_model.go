// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableKnowledgeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableKnowledgeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableKnowledgeInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetTableKnowledgeInfoResponseBody) *GetTableKnowledgeInfoResponse
	GetBody() *GetTableKnowledgeInfoResponseBody
}

type GetTableKnowledgeInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableKnowledgeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableKnowledgeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableKnowledgeInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTableKnowledgeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableKnowledgeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableKnowledgeInfoResponse) GetBody() *GetTableKnowledgeInfoResponseBody {
	return s.Body
}

func (s *GetTableKnowledgeInfoResponse) SetHeaders(v map[string]*string) *GetTableKnowledgeInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTableKnowledgeInfoResponse) SetStatusCode(v int32) *GetTableKnowledgeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableKnowledgeInfoResponse) SetBody(v *GetTableKnowledgeInfoResponseBody) *GetTableKnowledgeInfoResponse {
	s.Body = v
	return s
}

func (s *GetTableKnowledgeInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

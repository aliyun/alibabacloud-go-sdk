// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgePointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgePointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgePointResponse
	GetStatusCode() *int32
	SetBody(v *GetKnowledgePointResponseBody) *GetKnowledgePointResponse
	GetBody() *GetKnowledgePointResponseBody
}

type GetKnowledgePointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKnowledgePointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgePointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgePointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgePointResponse) GetBody() *GetKnowledgePointResponseBody {
	return s.Body
}

func (s *GetKnowledgePointResponse) SetHeaders(v map[string]*string) *GetKnowledgePointResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgePointResponse) SetStatusCode(v int32) *GetKnowledgePointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgePointResponse) SetBody(v *GetKnowledgePointResponseBody) *GetKnowledgePointResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgePointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

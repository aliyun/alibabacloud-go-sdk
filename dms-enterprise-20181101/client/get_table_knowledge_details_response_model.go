// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableKnowledgeDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableKnowledgeDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableKnowledgeDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetTableKnowledgeDetailsResponseBody) *GetTableKnowledgeDetailsResponse
	GetBody() *GetTableKnowledgeDetailsResponseBody
}

type GetTableKnowledgeDetailsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableKnowledgeDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableKnowledgeDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableKnowledgeDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetTableKnowledgeDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableKnowledgeDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableKnowledgeDetailsResponse) GetBody() *GetTableKnowledgeDetailsResponseBody {
	return s.Body
}

func (s *GetTableKnowledgeDetailsResponse) SetHeaders(v map[string]*string) *GetTableKnowledgeDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetTableKnowledgeDetailsResponse) SetStatusCode(v int32) *GetTableKnowledgeDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableKnowledgeDetailsResponse) SetBody(v *GetTableKnowledgeDetailsResponseBody) *GetTableKnowledgeDetailsResponse {
	s.Body = v
	return s
}

func (s *GetTableKnowledgeDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

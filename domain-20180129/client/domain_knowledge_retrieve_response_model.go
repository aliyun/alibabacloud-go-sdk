// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomainKnowledgeRetrieveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DomainKnowledgeRetrieveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DomainKnowledgeRetrieveResponse
	GetStatusCode() *int32
	SetBody(v *DomainKnowledgeRetrieveResponseBody) *DomainKnowledgeRetrieveResponse
	GetBody() *DomainKnowledgeRetrieveResponseBody
}

type DomainKnowledgeRetrieveResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DomainKnowledgeRetrieveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DomainKnowledgeRetrieveResponse) String() string {
	return dara.Prettify(s)
}

func (s DomainKnowledgeRetrieveResponse) GoString() string {
	return s.String()
}

func (s *DomainKnowledgeRetrieveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DomainKnowledgeRetrieveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DomainKnowledgeRetrieveResponse) GetBody() *DomainKnowledgeRetrieveResponseBody {
	return s.Body
}

func (s *DomainKnowledgeRetrieveResponse) SetHeaders(v map[string]*string) *DomainKnowledgeRetrieveResponse {
	s.Headers = v
	return s
}

func (s *DomainKnowledgeRetrieveResponse) SetStatusCode(v int32) *DomainKnowledgeRetrieveResponse {
	s.StatusCode = &v
	return s
}

func (s *DomainKnowledgeRetrieveResponse) SetBody(v *DomainKnowledgeRetrieveResponseBody) *DomainKnowledgeRetrieveResponse {
	s.Body = v
	return s
}

func (s *DomainKnowledgeRetrieveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInventoryKnowledgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchInventoryKnowledgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchInventoryKnowledgeResponse
	GetStatusCode() *int32
	SetBody(v *SearchInventoryKnowledgeResponseBody) *SearchInventoryKnowledgeResponse
	GetBody() *SearchInventoryKnowledgeResponseBody
}

type SearchInventoryKnowledgeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchInventoryKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchInventoryKnowledgeResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *SearchInventoryKnowledgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchInventoryKnowledgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchInventoryKnowledgeResponse) GetBody() *SearchInventoryKnowledgeResponseBody {
	return s.Body
}

func (s *SearchInventoryKnowledgeResponse) SetHeaders(v map[string]*string) *SearchInventoryKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *SearchInventoryKnowledgeResponse) SetStatusCode(v int32) *SearchInventoryKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchInventoryKnowledgeResponse) SetBody(v *SearchInventoryKnowledgeResponseBody) *SearchInventoryKnowledgeResponse {
	s.Body = v
	return s
}

func (s *SearchInventoryKnowledgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

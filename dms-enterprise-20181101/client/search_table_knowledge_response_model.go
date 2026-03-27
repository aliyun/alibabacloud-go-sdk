// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTableKnowledgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchTableKnowledgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchTableKnowledgeResponse
	GetStatusCode() *int32
	SetBody(v *SearchTableKnowledgeResponseBody) *SearchTableKnowledgeResponse
	GetBody() *SearchTableKnowledgeResponseBody
}

type SearchTableKnowledgeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTableKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTableKnowledgeResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchTableKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *SearchTableKnowledgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchTableKnowledgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchTableKnowledgeResponse) GetBody() *SearchTableKnowledgeResponseBody {
	return s.Body
}

func (s *SearchTableKnowledgeResponse) SetHeaders(v map[string]*string) *SearchTableKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *SearchTableKnowledgeResponse) SetStatusCode(v int32) *SearchTableKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTableKnowledgeResponse) SetBody(v *SearchTableKnowledgeResponseBody) *SearchTableKnowledgeResponse {
	s.Body = v
	return s
}

func (s *SearchTableKnowledgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

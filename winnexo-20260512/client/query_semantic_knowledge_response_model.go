// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySemanticKnowledgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySemanticKnowledgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySemanticKnowledgeResponse
	GetStatusCode() *int32
	SetBody(v *QuerySemanticKnowledgeResponseBody) *QuerySemanticKnowledgeResponse
	GetBody() *QuerySemanticKnowledgeResponseBody
}

type QuerySemanticKnowledgeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySemanticKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySemanticKnowledgeResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySemanticKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *QuerySemanticKnowledgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySemanticKnowledgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySemanticKnowledgeResponse) GetBody() *QuerySemanticKnowledgeResponseBody {
	return s.Body
}

func (s *QuerySemanticKnowledgeResponse) SetHeaders(v map[string]*string) *QuerySemanticKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *QuerySemanticKnowledgeResponse) SetStatusCode(v int32) *QuerySemanticKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySemanticKnowledgeResponse) SetBody(v *QuerySemanticKnowledgeResponseBody) *QuerySemanticKnowledgeResponse {
	s.Body = v
	return s
}

func (s *QuerySemanticKnowledgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

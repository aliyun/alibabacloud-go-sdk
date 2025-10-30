// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryKnowledgeBasesContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryKnowledgeBasesContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryKnowledgeBasesContentResponse
	GetStatusCode() *int32
	SetBody(v *QueryKnowledgeBasesContentResponseBody) *QueryKnowledgeBasesContentResponse
	GetBody() *QueryKnowledgeBasesContentResponseBody
}

type QueryKnowledgeBasesContentResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryKnowledgeBasesContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryKnowledgeBasesContentResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentResponse) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryKnowledgeBasesContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryKnowledgeBasesContentResponse) GetBody() *QueryKnowledgeBasesContentResponseBody {
	return s.Body
}

func (s *QueryKnowledgeBasesContentResponse) SetHeaders(v map[string]*string) *QueryKnowledgeBasesContentResponse {
	s.Headers = v
	return s
}

func (s *QueryKnowledgeBasesContentResponse) SetStatusCode(v int32) *QueryKnowledgeBasesContentResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponse) SetBody(v *QueryKnowledgeBasesContentResponseBody) *QueryKnowledgeBasesContentResponse {
	s.Body = v
	return s
}

func (s *QueryKnowledgeBasesContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

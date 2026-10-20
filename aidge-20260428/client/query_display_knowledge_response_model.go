// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDisplayKnowledgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDisplayKnowledgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDisplayKnowledgeResponse
	GetStatusCode() *int32
	SetBody(v *QueryDisplayKnowledgeResponseBody) *QueryDisplayKnowledgeResponse
	GetBody() *QueryDisplayKnowledgeResponseBody
}

type QueryDisplayKnowledgeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDisplayKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDisplayKnowledgeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDisplayKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *QueryDisplayKnowledgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDisplayKnowledgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDisplayKnowledgeResponse) GetBody() *QueryDisplayKnowledgeResponseBody {
	return s.Body
}

func (s *QueryDisplayKnowledgeResponse) SetHeaders(v map[string]*string) *QueryDisplayKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *QueryDisplayKnowledgeResponse) SetStatusCode(v int32) *QueryDisplayKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDisplayKnowledgeResponse) SetBody(v *QueryDisplayKnowledgeResponseBody) *QueryDisplayKnowledgeResponse {
	s.Body = v
	return s
}

func (s *QueryDisplayKnowledgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

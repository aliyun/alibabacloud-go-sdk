// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDisplayKnowledgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDisplayKnowledgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDisplayKnowledgeResponse
	GetStatusCode() *int32
	SetBody(v *AddDisplayKnowledgeResponseBody) *AddDisplayKnowledgeResponse
	GetBody() *AddDisplayKnowledgeResponseBody
}

type AddDisplayKnowledgeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDisplayKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDisplayKnowledgeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDisplayKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *AddDisplayKnowledgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDisplayKnowledgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDisplayKnowledgeResponse) GetBody() *AddDisplayKnowledgeResponseBody {
	return s.Body
}

func (s *AddDisplayKnowledgeResponse) SetHeaders(v map[string]*string) *AddDisplayKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *AddDisplayKnowledgeResponse) SetStatusCode(v int32) *AddDisplayKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDisplayKnowledgeResponse) SetBody(v *AddDisplayKnowledgeResponseBody) *AddDisplayKnowledgeResponse {
	s.Body = v
	return s
}

func (s *AddDisplayKnowledgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

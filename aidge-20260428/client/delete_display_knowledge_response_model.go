// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDisplayKnowledgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDisplayKnowledgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDisplayKnowledgeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDisplayKnowledgeResponseBody) *DeleteDisplayKnowledgeResponse
	GetBody() *DeleteDisplayKnowledgeResponseBody
}

type DeleteDisplayKnowledgeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDisplayKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDisplayKnowledgeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDisplayKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *DeleteDisplayKnowledgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDisplayKnowledgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDisplayKnowledgeResponse) GetBody() *DeleteDisplayKnowledgeResponseBody {
	return s.Body
}

func (s *DeleteDisplayKnowledgeResponse) SetHeaders(v map[string]*string) *DeleteDisplayKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *DeleteDisplayKnowledgeResponse) SetStatusCode(v int32) *DeleteDisplayKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDisplayKnowledgeResponse) SetBody(v *DeleteDisplayKnowledgeResponseBody) *DeleteDisplayKnowledgeResponse {
	s.Body = v
	return s
}

func (s *DeleteDisplayKnowledgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

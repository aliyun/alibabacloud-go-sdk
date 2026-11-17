// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKnowledgeObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKnowledgeObjectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKnowledgeObjectResponseBody) *DeleteKnowledgeObjectResponse
	GetBody() *DeleteKnowledgeObjectResponseBody
}

type DeleteKnowledgeObjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKnowledgeObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKnowledgeObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeObjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKnowledgeObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKnowledgeObjectResponse) GetBody() *DeleteKnowledgeObjectResponseBody {
	return s.Body
}

func (s *DeleteKnowledgeObjectResponse) SetHeaders(v map[string]*string) *DeleteKnowledgeObjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteKnowledgeObjectResponse) SetStatusCode(v int32) *DeleteKnowledgeObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKnowledgeObjectResponse) SetBody(v *DeleteKnowledgeObjectResponseBody) *DeleteKnowledgeObjectResponse {
	s.Body = v
	return s
}

func (s *DeleteKnowledgeObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

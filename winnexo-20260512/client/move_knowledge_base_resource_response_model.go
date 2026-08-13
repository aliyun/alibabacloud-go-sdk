// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveKnowledgeBaseResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveKnowledgeBaseResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveKnowledgeBaseResourceResponse
	GetStatusCode() *int32
	SetBody(v *MoveKnowledgeBaseResourceResponseBody) *MoveKnowledgeBaseResourceResponse
	GetBody() *MoveKnowledgeBaseResourceResponseBody
}

type MoveKnowledgeBaseResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveKnowledgeBaseResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveKnowledgeBaseResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveKnowledgeBaseResourceResponse) GoString() string {
	return s.String()
}

func (s *MoveKnowledgeBaseResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveKnowledgeBaseResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveKnowledgeBaseResourceResponse) GetBody() *MoveKnowledgeBaseResourceResponseBody {
	return s.Body
}

func (s *MoveKnowledgeBaseResourceResponse) SetHeaders(v map[string]*string) *MoveKnowledgeBaseResourceResponse {
	s.Headers = v
	return s
}

func (s *MoveKnowledgeBaseResourceResponse) SetStatusCode(v int32) *MoveKnowledgeBaseResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveKnowledgeBaseResourceResponse) SetBody(v *MoveKnowledgeBaseResourceResponseBody) *MoveKnowledgeBaseResourceResponse {
	s.Body = v
	return s
}

func (s *MoveKnowledgeBaseResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

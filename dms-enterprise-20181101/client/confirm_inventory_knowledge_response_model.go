// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmInventoryKnowledgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmInventoryKnowledgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmInventoryKnowledgeResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmInventoryKnowledgeResponseBody) *ConfirmInventoryKnowledgeResponse
	GetBody() *ConfirmInventoryKnowledgeResponseBody
}

type ConfirmInventoryKnowledgeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmInventoryKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmInventoryKnowledgeResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmInventoryKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *ConfirmInventoryKnowledgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmInventoryKnowledgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmInventoryKnowledgeResponse) GetBody() *ConfirmInventoryKnowledgeResponseBody {
	return s.Body
}

func (s *ConfirmInventoryKnowledgeResponse) SetHeaders(v map[string]*string) *ConfirmInventoryKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *ConfirmInventoryKnowledgeResponse) SetStatusCode(v int32) *ConfirmInventoryKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmInventoryKnowledgeResponse) SetBody(v *ConfirmInventoryKnowledgeResponseBody) *ConfirmInventoryKnowledgeResponse {
	s.Body = v
	return s
}

func (s *ConfirmInventoryKnowledgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

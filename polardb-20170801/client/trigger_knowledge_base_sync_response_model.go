// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerKnowledgeBaseSyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerKnowledgeBaseSyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerKnowledgeBaseSyncResponse
	GetStatusCode() *int32
	SetBody(v *TriggerKnowledgeBaseSyncResponseBody) *TriggerKnowledgeBaseSyncResponse
	GetBody() *TriggerKnowledgeBaseSyncResponseBody
}

type TriggerKnowledgeBaseSyncResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerKnowledgeBaseSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerKnowledgeBaseSyncResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerKnowledgeBaseSyncResponse) GoString() string {
	return s.String()
}

func (s *TriggerKnowledgeBaseSyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerKnowledgeBaseSyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerKnowledgeBaseSyncResponse) GetBody() *TriggerKnowledgeBaseSyncResponseBody {
	return s.Body
}

func (s *TriggerKnowledgeBaseSyncResponse) SetHeaders(v map[string]*string) *TriggerKnowledgeBaseSyncResponse {
	s.Headers = v
	return s
}

func (s *TriggerKnowledgeBaseSyncResponse) SetStatusCode(v int32) *TriggerKnowledgeBaseSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerKnowledgeBaseSyncResponse) SetBody(v *TriggerKnowledgeBaseSyncResponseBody) *TriggerKnowledgeBaseSyncResponse {
	s.Body = v
	return s
}

func (s *TriggerKnowledgeBaseSyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

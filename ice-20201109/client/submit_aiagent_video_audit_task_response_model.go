// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIAgentVideoAuditTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIAgentVideoAuditTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIAgentVideoAuditTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIAgentVideoAuditTaskResponseBody) *SubmitAIAgentVideoAuditTaskResponse
	GetBody() *SubmitAIAgentVideoAuditTaskResponseBody
}

type SubmitAIAgentVideoAuditTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIAgentVideoAuditTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIAgentVideoAuditTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIAgentVideoAuditTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIAgentVideoAuditTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIAgentVideoAuditTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIAgentVideoAuditTaskResponse) GetBody() *SubmitAIAgentVideoAuditTaskResponseBody {
	return s.Body
}

func (s *SubmitAIAgentVideoAuditTaskResponse) SetHeaders(v map[string]*string) *SubmitAIAgentVideoAuditTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskResponse) SetStatusCode(v int32) *SubmitAIAgentVideoAuditTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskResponse) SetBody(v *SubmitAIAgentVideoAuditTaskResponseBody) *SubmitAIAgentVideoAuditTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

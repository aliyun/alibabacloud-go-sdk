// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisplayAIAuditSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisplayAIAuditSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisplayAIAuditSwitchResponse
	GetStatusCode() *int32
	SetBody(v *DisplayAIAuditSwitchResponseBody) *DisplayAIAuditSwitchResponse
	GetBody() *DisplayAIAuditSwitchResponseBody
}

type DisplayAIAuditSwitchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisplayAIAuditSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisplayAIAuditSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s DisplayAIAuditSwitchResponse) GoString() string {
	return s.String()
}

func (s *DisplayAIAuditSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisplayAIAuditSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisplayAIAuditSwitchResponse) GetBody() *DisplayAIAuditSwitchResponseBody {
	return s.Body
}

func (s *DisplayAIAuditSwitchResponse) SetHeaders(v map[string]*string) *DisplayAIAuditSwitchResponse {
	s.Headers = v
	return s
}

func (s *DisplayAIAuditSwitchResponse) SetStatusCode(v int32) *DisplayAIAuditSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *DisplayAIAuditSwitchResponse) SetBody(v *DisplayAIAuditSwitchResponseBody) *DisplayAIAuditSwitchResponse {
	s.Body = v
	return s
}

func (s *DisplayAIAuditSwitchResponse) Validate() error {
	return dara.Validate(s)
}

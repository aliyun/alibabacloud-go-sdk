// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConversationalAutomationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConversationalAutomationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConversationalAutomationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConversationalAutomationResponseBody) *UpdateConversationalAutomationResponse
	GetBody() *UpdateConversationalAutomationResponseBody
}

type UpdateConversationalAutomationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConversationalAutomationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConversationalAutomationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConversationalAutomationResponse) GoString() string {
	return s.String()
}

func (s *UpdateConversationalAutomationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConversationalAutomationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConversationalAutomationResponse) GetBody() *UpdateConversationalAutomationResponseBody {
	return s.Body
}

func (s *UpdateConversationalAutomationResponse) SetHeaders(v map[string]*string) *UpdateConversationalAutomationResponse {
	s.Headers = v
	return s
}

func (s *UpdateConversationalAutomationResponse) SetStatusCode(v int32) *UpdateConversationalAutomationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConversationalAutomationResponse) SetBody(v *UpdateConversationalAutomationResponseBody) *UpdateConversationalAutomationResponse {
	s.Body = v
	return s
}

func (s *UpdateConversationalAutomationResponse) Validate() error {
	return dara.Validate(s)
}

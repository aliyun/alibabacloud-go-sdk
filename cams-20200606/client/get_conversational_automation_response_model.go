// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversationalAutomationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConversationalAutomationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConversationalAutomationResponse
	GetStatusCode() *int32
	SetBody(v *GetConversationalAutomationResponseBody) *GetConversationalAutomationResponse
	GetBody() *GetConversationalAutomationResponseBody
}

type GetConversationalAutomationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConversationalAutomationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConversationalAutomationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConversationalAutomationResponse) GoString() string {
	return s.String()
}

func (s *GetConversationalAutomationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConversationalAutomationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConversationalAutomationResponse) GetBody() *GetConversationalAutomationResponseBody {
	return s.Body
}

func (s *GetConversationalAutomationResponse) SetHeaders(v map[string]*string) *GetConversationalAutomationResponse {
	s.Headers = v
	return s
}

func (s *GetConversationalAutomationResponse) SetStatusCode(v int32) *GetConversationalAutomationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConversationalAutomationResponse) SetBody(v *GetConversationalAutomationResponseBody) *GetConversationalAutomationResponse {
	s.Body = v
	return s
}

func (s *GetConversationalAutomationResponse) Validate() error {
	return dara.Validate(s)
}

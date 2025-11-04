// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAgentOutboundCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAIAgentOutboundCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAIAgentOutboundCallResponse
	GetStatusCode() *int32
	SetBody(v *StartAIAgentOutboundCallResponseBody) *StartAIAgentOutboundCallResponse
	GetBody() *StartAIAgentOutboundCallResponseBody
}

type StartAIAgentOutboundCallResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAIAgentOutboundCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAIAgentOutboundCallResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAIAgentOutboundCallResponse) GoString() string {
	return s.String()
}

func (s *StartAIAgentOutboundCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAIAgentOutboundCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAIAgentOutboundCallResponse) GetBody() *StartAIAgentOutboundCallResponseBody {
	return s.Body
}

func (s *StartAIAgentOutboundCallResponse) SetHeaders(v map[string]*string) *StartAIAgentOutboundCallResponse {
	s.Headers = v
	return s
}

func (s *StartAIAgentOutboundCallResponse) SetStatusCode(v int32) *StartAIAgentOutboundCallResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAIAgentOutboundCallResponse) SetBody(v *StartAIAgentOutboundCallResponseBody) *StartAIAgentOutboundCallResponse {
	s.Body = v
	return s
}

func (s *StartAIAgentOutboundCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

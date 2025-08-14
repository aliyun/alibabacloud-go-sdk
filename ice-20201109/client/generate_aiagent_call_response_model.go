// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAIAgentCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateAIAgentCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateAIAgentCallResponse
	GetStatusCode() *int32
	SetBody(v *GenerateAIAgentCallResponseBody) *GenerateAIAgentCallResponse
	GetBody() *GenerateAIAgentCallResponseBody
}

type GenerateAIAgentCallResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAIAgentCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAIAgentCallResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateAIAgentCallResponse) GoString() string {
	return s.String()
}

func (s *GenerateAIAgentCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateAIAgentCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateAIAgentCallResponse) GetBody() *GenerateAIAgentCallResponseBody {
	return s.Body
}

func (s *GenerateAIAgentCallResponse) SetHeaders(v map[string]*string) *GenerateAIAgentCallResponse {
	s.Headers = v
	return s
}

func (s *GenerateAIAgentCallResponse) SetStatusCode(v int32) *GenerateAIAgentCallResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAIAgentCallResponse) SetBody(v *GenerateAIAgentCallResponseBody) *GenerateAIAgentCallResponse {
	s.Body = v
	return s
}

func (s *GenerateAIAgentCallResponse) Validate() error {
	return dara.Validate(s)
}

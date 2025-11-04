// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAIAgentVoiceprintResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAIAgentVoiceprintResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAIAgentVoiceprintResponse
	GetStatusCode() *int32
	SetBody(v *SetAIAgentVoiceprintResponseBody) *SetAIAgentVoiceprintResponse
	GetBody() *SetAIAgentVoiceprintResponseBody
}

type SetAIAgentVoiceprintResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAIAgentVoiceprintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAIAgentVoiceprintResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAIAgentVoiceprintResponse) GoString() string {
	return s.String()
}

func (s *SetAIAgentVoiceprintResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAIAgentVoiceprintResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAIAgentVoiceprintResponse) GetBody() *SetAIAgentVoiceprintResponseBody {
	return s.Body
}

func (s *SetAIAgentVoiceprintResponse) SetHeaders(v map[string]*string) *SetAIAgentVoiceprintResponse {
	s.Headers = v
	return s
}

func (s *SetAIAgentVoiceprintResponse) SetStatusCode(v int32) *SetAIAgentVoiceprintResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAIAgentVoiceprintResponse) SetBody(v *SetAIAgentVoiceprintResponseBody) *SetAIAgentVoiceprintResponse {
	s.Body = v
	return s
}

func (s *SetAIAgentVoiceprintResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

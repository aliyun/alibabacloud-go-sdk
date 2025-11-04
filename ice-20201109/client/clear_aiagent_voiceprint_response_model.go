// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAIAgentVoiceprintResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearAIAgentVoiceprintResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearAIAgentVoiceprintResponse
	GetStatusCode() *int32
	SetBody(v *ClearAIAgentVoiceprintResponseBody) *ClearAIAgentVoiceprintResponse
	GetBody() *ClearAIAgentVoiceprintResponseBody
}

type ClearAIAgentVoiceprintResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearAIAgentVoiceprintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearAIAgentVoiceprintResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearAIAgentVoiceprintResponse) GoString() string {
	return s.String()
}

func (s *ClearAIAgentVoiceprintResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearAIAgentVoiceprintResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearAIAgentVoiceprintResponse) GetBody() *ClearAIAgentVoiceprintResponseBody {
	return s.Body
}

func (s *ClearAIAgentVoiceprintResponse) SetHeaders(v map[string]*string) *ClearAIAgentVoiceprintResponse {
	s.Headers = v
	return s
}

func (s *ClearAIAgentVoiceprintResponse) SetStatusCode(v int32) *ClearAIAgentVoiceprintResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearAIAgentVoiceprintResponse) SetBody(v *ClearAIAgentVoiceprintResponseBody) *ClearAIAgentVoiceprintResponse {
	s.Body = v
	return s
}

func (s *ClearAIAgentVoiceprintResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentVoiceprintsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIAgentVoiceprintsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIAgentVoiceprintsResponse
	GetStatusCode() *int32
	SetBody(v *ListAIAgentVoiceprintsResponseBody) *ListAIAgentVoiceprintsResponse
	GetBody() *ListAIAgentVoiceprintsResponseBody
}

type ListAIAgentVoiceprintsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIAgentVoiceprintsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIAgentVoiceprintsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentVoiceprintsResponse) GoString() string {
	return s.String()
}

func (s *ListAIAgentVoiceprintsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIAgentVoiceprintsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIAgentVoiceprintsResponse) GetBody() *ListAIAgentVoiceprintsResponseBody {
	return s.Body
}

func (s *ListAIAgentVoiceprintsResponse) SetHeaders(v map[string]*string) *ListAIAgentVoiceprintsResponse {
	s.Headers = v
	return s
}

func (s *ListAIAgentVoiceprintsResponse) SetStatusCode(v int32) *ListAIAgentVoiceprintsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIAgentVoiceprintsResponse) SetBody(v *ListAIAgentVoiceprintsResponseBody) *ListAIAgentVoiceprintsResponse {
	s.Body = v
	return s
}

func (s *ListAIAgentVoiceprintsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

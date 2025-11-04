// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIAgentDialogueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAIAgentDialogueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAIAgentDialogueResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAIAgentDialogueResponseBody) *DeleteAIAgentDialogueResponse
	GetBody() *DeleteAIAgentDialogueResponseBody
}

type DeleteAIAgentDialogueResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAIAgentDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAIAgentDialogueResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIAgentDialogueResponse) GoString() string {
	return s.String()
}

func (s *DeleteAIAgentDialogueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAIAgentDialogueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAIAgentDialogueResponse) GetBody() *DeleteAIAgentDialogueResponseBody {
	return s.Body
}

func (s *DeleteAIAgentDialogueResponse) SetHeaders(v map[string]*string) *DeleteAIAgentDialogueResponse {
	s.Headers = v
	return s
}

func (s *DeleteAIAgentDialogueResponse) SetStatusCode(v int32) *DeleteAIAgentDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAIAgentDialogueResponse) SetBody(v *DeleteAIAgentDialogueResponseBody) *DeleteAIAgentDialogueResponse {
	s.Body = v
	return s
}

func (s *DeleteAIAgentDialogueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

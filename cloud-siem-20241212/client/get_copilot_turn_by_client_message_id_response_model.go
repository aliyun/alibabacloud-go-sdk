// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotTurnByClientMessageIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCopilotTurnByClientMessageIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCopilotTurnByClientMessageIdResponse
	GetStatusCode() *int32
	SetBody(v *GetCopilotTurnByClientMessageIdResponseBody) *GetCopilotTurnByClientMessageIdResponse
	GetBody() *GetCopilotTurnByClientMessageIdResponseBody
}

type GetCopilotTurnByClientMessageIdResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCopilotTurnByClientMessageIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCopilotTurnByClientMessageIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotTurnByClientMessageIdResponse) GoString() string {
	return s.String()
}

func (s *GetCopilotTurnByClientMessageIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCopilotTurnByClientMessageIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCopilotTurnByClientMessageIdResponse) GetBody() *GetCopilotTurnByClientMessageIdResponseBody {
	return s.Body
}

func (s *GetCopilotTurnByClientMessageIdResponse) SetHeaders(v map[string]*string) *GetCopilotTurnByClientMessageIdResponse {
	s.Headers = v
	return s
}

func (s *GetCopilotTurnByClientMessageIdResponse) SetStatusCode(v int32) *GetCopilotTurnByClientMessageIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCopilotTurnByClientMessageIdResponse) SetBody(v *GetCopilotTurnByClientMessageIdResponseBody) *GetCopilotTurnByClientMessageIdResponse {
	s.Body = v
	return s
}

func (s *GetCopilotTurnByClientMessageIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

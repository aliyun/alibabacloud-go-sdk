// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotTurnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCopilotTurnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCopilotTurnResponse
	GetStatusCode() *int32
	SetBody(v *GetCopilotTurnResponseBody) *GetCopilotTurnResponse
	GetBody() *GetCopilotTurnResponseBody
}

type GetCopilotTurnResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCopilotTurnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCopilotTurnResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotTurnResponse) GoString() string {
	return s.String()
}

func (s *GetCopilotTurnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCopilotTurnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCopilotTurnResponse) GetBody() *GetCopilotTurnResponseBody {
	return s.Body
}

func (s *GetCopilotTurnResponse) SetHeaders(v map[string]*string) *GetCopilotTurnResponse {
	s.Headers = v
	return s
}

func (s *GetCopilotTurnResponse) SetStatusCode(v int32) *GetCopilotTurnResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCopilotTurnResponse) SetBody(v *GetCopilotTurnResponseBody) *GetCopilotTurnResponse {
	s.Body = v
	return s
}

func (s *GetCopilotTurnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

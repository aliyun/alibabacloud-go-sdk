// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCopilotTurnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCopilotTurnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCopilotTurnResponse
	GetStatusCode() *int32
	SetBody(v *CancelCopilotTurnResponseBody) *CancelCopilotTurnResponse
	GetBody() *CancelCopilotTurnResponseBody
}

type CancelCopilotTurnResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCopilotTurnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCopilotTurnResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCopilotTurnResponse) GoString() string {
	return s.String()
}

func (s *CancelCopilotTurnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCopilotTurnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCopilotTurnResponse) GetBody() *CancelCopilotTurnResponseBody {
	return s.Body
}

func (s *CancelCopilotTurnResponse) SetHeaders(v map[string]*string) *CancelCopilotTurnResponse {
	s.Headers = v
	return s
}

func (s *CancelCopilotTurnResponse) SetStatusCode(v int32) *CancelCopilotTurnResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCopilotTurnResponse) SetBody(v *CancelCopilotTurnResponseBody) *CancelCopilotTurnResponse {
	s.Body = v
	return s
}

func (s *CancelCopilotTurnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

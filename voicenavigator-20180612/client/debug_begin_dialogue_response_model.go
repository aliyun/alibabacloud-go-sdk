// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugBeginDialogueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DebugBeginDialogueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DebugBeginDialogueResponse
	GetStatusCode() *int32
	SetBody(v *DebugBeginDialogueResponseBody) *DebugBeginDialogueResponse
	GetBody() *DebugBeginDialogueResponseBody
}

type DebugBeginDialogueResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugBeginDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugBeginDialogueResponse) String() string {
	return dara.Prettify(s)
}

func (s DebugBeginDialogueResponse) GoString() string {
	return s.String()
}

func (s *DebugBeginDialogueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DebugBeginDialogueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DebugBeginDialogueResponse) GetBody() *DebugBeginDialogueResponseBody {
	return s.Body
}

func (s *DebugBeginDialogueResponse) SetHeaders(v map[string]*string) *DebugBeginDialogueResponse {
	s.Headers = v
	return s
}

func (s *DebugBeginDialogueResponse) SetStatusCode(v int32) *DebugBeginDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugBeginDialogueResponse) SetBody(v *DebugBeginDialogueResponseBody) *DebugBeginDialogueResponse {
	s.Body = v
	return s
}

func (s *DebugBeginDialogueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

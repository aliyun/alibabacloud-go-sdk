// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugDialogueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DebugDialogueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DebugDialogueResponse
	GetStatusCode() *int32
	SetBody(v *DebugDialogueResponseBody) *DebugDialogueResponse
	GetBody() *DebugDialogueResponseBody
}

type DebugDialogueResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugDialogueResponse) String() string {
	return dara.Prettify(s)
}

func (s DebugDialogueResponse) GoString() string {
	return s.String()
}

func (s *DebugDialogueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DebugDialogueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DebugDialogueResponse) GetBody() *DebugDialogueResponseBody {
	return s.Body
}

func (s *DebugDialogueResponse) SetHeaders(v map[string]*string) *DebugDialogueResponse {
	s.Headers = v
	return s
}

func (s *DebugDialogueResponse) SetStatusCode(v int32) *DebugDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugDialogueResponse) SetBody(v *DebugDialogueResponseBody) *DebugDialogueResponse {
	s.Body = v
	return s
}

func (s *DebugDialogueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

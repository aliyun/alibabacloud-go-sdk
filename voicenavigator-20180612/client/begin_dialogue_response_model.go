// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBeginDialogueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BeginDialogueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BeginDialogueResponse
	GetStatusCode() *int32
	SetBody(v *BeginDialogueResponseBody) *BeginDialogueResponse
	GetBody() *BeginDialogueResponseBody
}

type BeginDialogueResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BeginDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BeginDialogueResponse) String() string {
	return dara.Prettify(s)
}

func (s BeginDialogueResponse) GoString() string {
	return s.String()
}

func (s *BeginDialogueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BeginDialogueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BeginDialogueResponse) GetBody() *BeginDialogueResponseBody {
	return s.Body
}

func (s *BeginDialogueResponse) SetHeaders(v map[string]*string) *BeginDialogueResponse {
	s.Headers = v
	return s
}

func (s *BeginDialogueResponse) SetStatusCode(v int32) *BeginDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *BeginDialogueResponse) SetBody(v *BeginDialogueResponseBody) *BeginDialogueResponse {
	s.Body = v
	return s
}

func (s *BeginDialogueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDialogueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DialogueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DialogueResponse
	GetStatusCode() *int32
	SetBody(v *DialogueResponseBody) *DialogueResponse
	GetBody() *DialogueResponseBody
}

type DialogueResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DialogueResponse) String() string {
	return dara.Prettify(s)
}

func (s DialogueResponse) GoString() string {
	return s.String()
}

func (s *DialogueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DialogueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DialogueResponse) GetBody() *DialogueResponseBody {
	return s.Body
}

func (s *DialogueResponse) SetHeaders(v map[string]*string) *DialogueResponse {
	s.Headers = v
	return s
}

func (s *DialogueResponse) SetStatusCode(v int32) *DialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *DialogueResponse) SetBody(v *DialogueResponseBody) *DialogueResponse {
	s.Body = v
	return s
}

func (s *DialogueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

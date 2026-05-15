// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartChatWorkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartChatWorkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartChatWorkResponse
	GetStatusCode() *int32
	SetBody(v *StartChatWorkResponseBody) *StartChatWorkResponse
	GetBody() *StartChatWorkResponseBody
}

type StartChatWorkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartChatWorkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartChatWorkResponse) String() string {
	return dara.Prettify(s)
}

func (s StartChatWorkResponse) GoString() string {
	return s.String()
}

func (s *StartChatWorkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartChatWorkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartChatWorkResponse) GetBody() *StartChatWorkResponseBody {
	return s.Body
}

func (s *StartChatWorkResponse) SetHeaders(v map[string]*string) *StartChatWorkResponse {
	s.Headers = v
	return s
}

func (s *StartChatWorkResponse) SetStatusCode(v int32) *StartChatWorkResponse {
	s.StatusCode = &v
	return s
}

func (s *StartChatWorkResponse) SetBody(v *StartChatWorkResponseBody) *StartChatWorkResponse {
	s.Body = v
	return s
}

func (s *StartChatWorkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

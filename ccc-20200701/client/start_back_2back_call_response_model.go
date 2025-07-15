// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBack2BackCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartBack2BackCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartBack2BackCallResponse
	GetStatusCode() *int32
	SetBody(v *StartBack2BackCallResponseBody) *StartBack2BackCallResponse
	GetBody() *StartBack2BackCallResponseBody
}

type StartBack2BackCallResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartBack2BackCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartBack2BackCallResponse) String() string {
	return dara.Prettify(s)
}

func (s StartBack2BackCallResponse) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartBack2BackCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartBack2BackCallResponse) GetBody() *StartBack2BackCallResponseBody {
	return s.Body
}

func (s *StartBack2BackCallResponse) SetHeaders(v map[string]*string) *StartBack2BackCallResponse {
	s.Headers = v
	return s
}

func (s *StartBack2BackCallResponse) SetStatusCode(v int32) *StartBack2BackCallResponse {
	s.StatusCode = &v
	return s
}

func (s *StartBack2BackCallResponse) SetBody(v *StartBack2BackCallResponseBody) *StartBack2BackCallResponse {
	s.Body = v
	return s
}

func (s *StartBack2BackCallResponse) Validate() error {
	return dara.Validate(s)
}

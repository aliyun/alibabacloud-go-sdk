// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerminalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTerminalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTerminalResponse
	GetStatusCode() *int32
	SetBody(v *ListTerminalResponseBody) *ListTerminalResponse
	GetBody() *ListTerminalResponseBody
}

type ListTerminalResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTerminalResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalResponse) GoString() string {
	return s.String()
}

func (s *ListTerminalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTerminalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTerminalResponse) GetBody() *ListTerminalResponseBody {
	return s.Body
}

func (s *ListTerminalResponse) SetHeaders(v map[string]*string) *ListTerminalResponse {
	s.Headers = v
	return s
}

func (s *ListTerminalResponse) SetStatusCode(v int32) *ListTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTerminalResponse) SetBody(v *ListTerminalResponseBody) *ListTerminalResponse {
	s.Body = v
	return s
}

func (s *ListTerminalResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceWebTerminalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckInstanceWebTerminalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckInstanceWebTerminalResponse
	GetStatusCode() *int32
	SetBody(v *CheckInstanceWebTerminalResponseBody) *CheckInstanceWebTerminalResponse
	GetBody() *CheckInstanceWebTerminalResponseBody
}

type CheckInstanceWebTerminalResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckInstanceWebTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckInstanceWebTerminalResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceWebTerminalResponse) GoString() string {
	return s.String()
}

func (s *CheckInstanceWebTerminalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckInstanceWebTerminalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckInstanceWebTerminalResponse) GetBody() *CheckInstanceWebTerminalResponseBody {
	return s.Body
}

func (s *CheckInstanceWebTerminalResponse) SetHeaders(v map[string]*string) *CheckInstanceWebTerminalResponse {
	s.Headers = v
	return s
}

func (s *CheckInstanceWebTerminalResponse) SetStatusCode(v int32) *CheckInstanceWebTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckInstanceWebTerminalResponse) SetBody(v *CheckInstanceWebTerminalResponseBody) *CheckInstanceWebTerminalResponse {
	s.Body = v
	return s
}

func (s *CheckInstanceWebTerminalResponse) Validate() error {
	return dara.Validate(s)
}

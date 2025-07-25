// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceWebTerminalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceWebTerminalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceWebTerminalResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceWebTerminalResponseBody) *CreateInstanceWebTerminalResponse
	GetBody() *CreateInstanceWebTerminalResponseBody
}

type CreateInstanceWebTerminalResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceWebTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceWebTerminalResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceWebTerminalResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceWebTerminalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceWebTerminalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceWebTerminalResponse) GetBody() *CreateInstanceWebTerminalResponseBody {
	return s.Body
}

func (s *CreateInstanceWebTerminalResponse) SetHeaders(v map[string]*string) *CreateInstanceWebTerminalResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceWebTerminalResponse) SetStatusCode(v int32) *CreateInstanceWebTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceWebTerminalResponse) SetBody(v *CreateInstanceWebTerminalResponseBody) *CreateInstanceWebTerminalResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceWebTerminalResponse) Validate() error {
	return dara.Validate(s)
}

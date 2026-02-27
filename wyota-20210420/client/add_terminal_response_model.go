// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTerminalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTerminalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTerminalResponse
	GetStatusCode() *int32
	SetBody(v *AddTerminalResponseBody) *AddTerminalResponse
	GetBody() *AddTerminalResponseBody
}

type AddTerminalResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTerminalResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTerminalResponse) GoString() string {
	return s.String()
}

func (s *AddTerminalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTerminalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTerminalResponse) GetBody() *AddTerminalResponseBody {
	return s.Body
}

func (s *AddTerminalResponse) SetHeaders(v map[string]*string) *AddTerminalResponse {
	s.Headers = v
	return s
}

func (s *AddTerminalResponse) SetStatusCode(v int32) *AddTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTerminalResponse) SetBody(v *AddTerminalResponseBody) *AddTerminalResponse {
	s.Body = v
	return s
}

func (s *AddTerminalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

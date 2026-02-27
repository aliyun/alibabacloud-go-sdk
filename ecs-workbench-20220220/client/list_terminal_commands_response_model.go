// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerminalCommandsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTerminalCommandsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTerminalCommandsResponse
	GetStatusCode() *int32
	SetBody(v *ListTerminalCommandsResponseBody) *ListTerminalCommandsResponse
	GetBody() *ListTerminalCommandsResponseBody
}

type ListTerminalCommandsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTerminalCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTerminalCommandsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalCommandsResponse) GoString() string {
	return s.String()
}

func (s *ListTerminalCommandsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTerminalCommandsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTerminalCommandsResponse) GetBody() *ListTerminalCommandsResponseBody {
	return s.Body
}

func (s *ListTerminalCommandsResponse) SetHeaders(v map[string]*string) *ListTerminalCommandsResponse {
	s.Headers = v
	return s
}

func (s *ListTerminalCommandsResponse) SetStatusCode(v int32) *ListTerminalCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTerminalCommandsResponse) SetBody(v *ListTerminalCommandsResponseBody) *ListTerminalCommandsResponse {
	s.Body = v
	return s
}

func (s *ListTerminalCommandsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

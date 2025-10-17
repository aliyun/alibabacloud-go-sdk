// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebTerminalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWebTerminalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWebTerminalResponse
	GetStatusCode() *int32
	SetBody(v *GetWebTerminalResponseBody) *GetWebTerminalResponse
	GetBody() *GetWebTerminalResponseBody
}

type GetWebTerminalResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWebTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWebTerminalResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWebTerminalResponse) GoString() string {
	return s.String()
}

func (s *GetWebTerminalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWebTerminalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWebTerminalResponse) GetBody() *GetWebTerminalResponseBody {
	return s.Body
}

func (s *GetWebTerminalResponse) SetHeaders(v map[string]*string) *GetWebTerminalResponse {
	s.Headers = v
	return s
}

func (s *GetWebTerminalResponse) SetStatusCode(v int32) *GetWebTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWebTerminalResponse) SetBody(v *GetWebTerminalResponseBody) *GetWebTerminalResponse {
	s.Body = v
	return s
}

func (s *GetWebTerminalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

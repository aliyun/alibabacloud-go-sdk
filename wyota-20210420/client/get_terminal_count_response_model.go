// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTerminalCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTerminalCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTerminalCountResponse
	GetStatusCode() *int32
	SetBody(v *GetTerminalCountResponseBody) *GetTerminalCountResponse
	GetBody() *GetTerminalCountResponseBody
}

type GetTerminalCountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTerminalCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTerminalCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTerminalCountResponse) GoString() string {
	return s.String()
}

func (s *GetTerminalCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTerminalCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTerminalCountResponse) GetBody() *GetTerminalCountResponseBody {
	return s.Body
}

func (s *GetTerminalCountResponse) SetHeaders(v map[string]*string) *GetTerminalCountResponse {
	s.Headers = v
	return s
}

func (s *GetTerminalCountResponse) SetStatusCode(v int32) *GetTerminalCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTerminalCountResponse) SetBody(v *GetTerminalCountResponseBody) *GetTerminalCountResponse {
	s.Body = v
	return s
}

func (s *GetTerminalCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

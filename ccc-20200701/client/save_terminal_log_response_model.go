// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTerminalLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTerminalLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTerminalLogResponse
	GetStatusCode() *int32
	SetBody(v *SaveTerminalLogResponseBody) *SaveTerminalLogResponse
	GetBody() *SaveTerminalLogResponseBody
}

type SaveTerminalLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTerminalLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTerminalLogResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTerminalLogResponse) GoString() string {
	return s.String()
}

func (s *SaveTerminalLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTerminalLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTerminalLogResponse) GetBody() *SaveTerminalLogResponseBody {
	return s.Body
}

func (s *SaveTerminalLogResponse) SetHeaders(v map[string]*string) *SaveTerminalLogResponse {
	s.Headers = v
	return s
}

func (s *SaveTerminalLogResponse) SetStatusCode(v int32) *SaveTerminalLogResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTerminalLogResponse) SetBody(v *SaveTerminalLogResponseBody) *SaveTerminalLogResponse {
	s.Body = v
	return s
}

func (s *SaveTerminalLogResponse) Validate() error {
	return dara.Validate(s)
}

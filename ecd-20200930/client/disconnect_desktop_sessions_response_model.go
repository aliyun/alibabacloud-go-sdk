// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisconnectDesktopSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisconnectDesktopSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisconnectDesktopSessionsResponse
	GetStatusCode() *int32
	SetBody(v *DisconnectDesktopSessionsResponseBody) *DisconnectDesktopSessionsResponse
	GetBody() *DisconnectDesktopSessionsResponseBody
}

type DisconnectDesktopSessionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisconnectDesktopSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisconnectDesktopSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DisconnectDesktopSessionsResponse) GoString() string {
	return s.String()
}

func (s *DisconnectDesktopSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisconnectDesktopSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisconnectDesktopSessionsResponse) GetBody() *DisconnectDesktopSessionsResponseBody {
	return s.Body
}

func (s *DisconnectDesktopSessionsResponse) SetHeaders(v map[string]*string) *DisconnectDesktopSessionsResponse {
	s.Headers = v
	return s
}

func (s *DisconnectDesktopSessionsResponse) SetStatusCode(v int32) *DisconnectDesktopSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DisconnectDesktopSessionsResponse) SetBody(v *DisconnectDesktopSessionsResponseBody) *DisconnectDesktopSessionsResponse {
	s.Body = v
	return s
}

func (s *DisconnectDesktopSessionsResponse) Validate() error {
	return dara.Validate(s)
}

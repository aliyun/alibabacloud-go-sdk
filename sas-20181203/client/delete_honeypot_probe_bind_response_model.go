// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotProbeBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHoneypotProbeBindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHoneypotProbeBindResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHoneypotProbeBindResponseBody) *DeleteHoneypotProbeBindResponse
	GetBody() *DeleteHoneypotProbeBindResponseBody
}

type DeleteHoneypotProbeBindResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHoneypotProbeBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHoneypotProbeBindResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotProbeBindResponse) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotProbeBindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHoneypotProbeBindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHoneypotProbeBindResponse) GetBody() *DeleteHoneypotProbeBindResponseBody {
	return s.Body
}

func (s *DeleteHoneypotProbeBindResponse) SetHeaders(v map[string]*string) *DeleteHoneypotProbeBindResponse {
	s.Headers = v
	return s
}

func (s *DeleteHoneypotProbeBindResponse) SetStatusCode(v int32) *DeleteHoneypotProbeBindResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHoneypotProbeBindResponse) SetBody(v *DeleteHoneypotProbeBindResponseBody) *DeleteHoneypotProbeBindResponse {
	s.Body = v
	return s
}

func (s *DeleteHoneypotProbeBindResponse) Validate() error {
	return dara.Validate(s)
}

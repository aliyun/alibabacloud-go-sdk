// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotProbeBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHoneypotProbeBindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHoneypotProbeBindResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHoneypotProbeBindResponseBody) *UpdateHoneypotProbeBindResponse
	GetBody() *UpdateHoneypotProbeBindResponseBody
}

type UpdateHoneypotProbeBindResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHoneypotProbeBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHoneypotProbeBindResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotProbeBindResponse) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotProbeBindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHoneypotProbeBindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHoneypotProbeBindResponse) GetBody() *UpdateHoneypotProbeBindResponseBody {
	return s.Body
}

func (s *UpdateHoneypotProbeBindResponse) SetHeaders(v map[string]*string) *UpdateHoneypotProbeBindResponse {
	s.Headers = v
	return s
}

func (s *UpdateHoneypotProbeBindResponse) SetStatusCode(v int32) *UpdateHoneypotProbeBindResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHoneypotProbeBindResponse) SetBody(v *UpdateHoneypotProbeBindResponseBody) *UpdateHoneypotProbeBindResponse {
	s.Body = v
	return s
}

func (s *UpdateHoneypotProbeBindResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

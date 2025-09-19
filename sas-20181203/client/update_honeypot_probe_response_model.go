// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotProbeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHoneypotProbeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHoneypotProbeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHoneypotProbeResponseBody) *UpdateHoneypotProbeResponse
	GetBody() *UpdateHoneypotProbeResponseBody
}

type UpdateHoneypotProbeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHoneypotProbeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHoneypotProbeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotProbeResponse) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotProbeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHoneypotProbeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHoneypotProbeResponse) GetBody() *UpdateHoneypotProbeResponseBody {
	return s.Body
}

func (s *UpdateHoneypotProbeResponse) SetHeaders(v map[string]*string) *UpdateHoneypotProbeResponse {
	s.Headers = v
	return s
}

func (s *UpdateHoneypotProbeResponse) SetStatusCode(v int32) *UpdateHoneypotProbeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHoneypotProbeResponse) SetBody(v *UpdateHoneypotProbeResponseBody) *UpdateHoneypotProbeResponse {
	s.Body = v
	return s
}

func (s *UpdateHoneypotProbeResponse) Validate() error {
	return dara.Validate(s)
}

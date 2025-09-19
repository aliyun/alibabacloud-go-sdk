// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotProbeBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHoneypotProbeBindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHoneypotProbeBindResponse
	GetStatusCode() *int32
	SetBody(v *CreateHoneypotProbeBindResponseBody) *CreateHoneypotProbeBindResponse
	GetBody() *CreateHoneypotProbeBindResponseBody
}

type CreateHoneypotProbeBindResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHoneypotProbeBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHoneypotProbeBindResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotProbeBindResponse) GoString() string {
	return s.String()
}

func (s *CreateHoneypotProbeBindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHoneypotProbeBindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHoneypotProbeBindResponse) GetBody() *CreateHoneypotProbeBindResponseBody {
	return s.Body
}

func (s *CreateHoneypotProbeBindResponse) SetHeaders(v map[string]*string) *CreateHoneypotProbeBindResponse {
	s.Headers = v
	return s
}

func (s *CreateHoneypotProbeBindResponse) SetStatusCode(v int32) *CreateHoneypotProbeBindResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHoneypotProbeBindResponse) SetBody(v *CreateHoneypotProbeBindResponseBody) *CreateHoneypotProbeBindResponse {
	s.Body = v
	return s
}

func (s *CreateHoneypotProbeBindResponse) Validate() error {
	return dara.Validate(s)
}

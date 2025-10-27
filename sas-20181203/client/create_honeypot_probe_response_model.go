// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotProbeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHoneypotProbeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHoneypotProbeResponse
	GetStatusCode() *int32
	SetBody(v *CreateHoneypotProbeResponseBody) *CreateHoneypotProbeResponse
	GetBody() *CreateHoneypotProbeResponseBody
}

type CreateHoneypotProbeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHoneypotProbeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHoneypotProbeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotProbeResponse) GoString() string {
	return s.String()
}

func (s *CreateHoneypotProbeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHoneypotProbeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHoneypotProbeResponse) GetBody() *CreateHoneypotProbeResponseBody {
	return s.Body
}

func (s *CreateHoneypotProbeResponse) SetHeaders(v map[string]*string) *CreateHoneypotProbeResponse {
	s.Headers = v
	return s
}

func (s *CreateHoneypotProbeResponse) SetStatusCode(v int32) *CreateHoneypotProbeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHoneypotProbeResponse) SetBody(v *CreateHoneypotProbeResponseBody) *CreateHoneypotProbeResponse {
	s.Body = v
	return s
}

func (s *CreateHoneypotProbeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

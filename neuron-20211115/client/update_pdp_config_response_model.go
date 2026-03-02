// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePdpConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePdpConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePdpConfigResponse
	GetStatusCode() *int32
	SetBody(v *PdpConfig) *UpdatePdpConfigResponse
	GetBody() *PdpConfig
}

type UpdatePdpConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpConfig         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePdpConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePdpConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdatePdpConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePdpConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePdpConfigResponse) GetBody() *PdpConfig {
	return s.Body
}

func (s *UpdatePdpConfigResponse) SetHeaders(v map[string]*string) *UpdatePdpConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdatePdpConfigResponse) SetStatusCode(v int32) *UpdatePdpConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePdpConfigResponse) SetBody(v *PdpConfig) *UpdatePdpConfigResponse {
	s.Body = v
	return s
}

func (s *UpdatePdpConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

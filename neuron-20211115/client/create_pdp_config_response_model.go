// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdpConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePdpConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePdpConfigResponse
	GetStatusCode() *int32
	SetBody(v *PdpConfig) *CreatePdpConfigResponse
	GetBody() *PdpConfig
}

type CreatePdpConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpConfig         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdpConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePdpConfigResponse) GoString() string {
	return s.String()
}

func (s *CreatePdpConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePdpConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePdpConfigResponse) GetBody() *PdpConfig {
	return s.Body
}

func (s *CreatePdpConfigResponse) SetHeaders(v map[string]*string) *CreatePdpConfigResponse {
	s.Headers = v
	return s
}

func (s *CreatePdpConfigResponse) SetStatusCode(v int32) *CreatePdpConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePdpConfigResponse) SetBody(v *PdpConfig) *CreatePdpConfigResponse {
	s.Body = v
	return s
}

func (s *CreatePdpConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

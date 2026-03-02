// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPdpConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPdpConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPdpConfigResponse
	GetStatusCode() *int32
	SetBody(v *PdpConfig) *GetPdpConfigResponse
	GetBody() *PdpConfig
}

type GetPdpConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpConfig         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPdpConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPdpConfigResponse) GoString() string {
	return s.String()
}

func (s *GetPdpConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPdpConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPdpConfigResponse) GetBody() *PdpConfig {
	return s.Body
}

func (s *GetPdpConfigResponse) SetHeaders(v map[string]*string) *GetPdpConfigResponse {
	s.Headers = v
	return s
}

func (s *GetPdpConfigResponse) SetStatusCode(v int32) *GetPdpConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPdpConfigResponse) SetBody(v *PdpConfig) *GetPdpConfigResponse {
	s.Body = v
	return s
}

func (s *GetPdpConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

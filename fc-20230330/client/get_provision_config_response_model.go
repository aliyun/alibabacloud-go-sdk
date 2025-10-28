// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProvisionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProvisionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProvisionConfigResponse
	GetStatusCode() *int32
	SetBody(v *ProvisionConfig) *GetProvisionConfigResponse
	GetBody() *ProvisionConfig
}

type GetProvisionConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProvisionConfig   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProvisionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionConfigResponse) GoString() string {
	return s.String()
}

func (s *GetProvisionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProvisionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProvisionConfigResponse) GetBody() *ProvisionConfig {
	return s.Body
}

func (s *GetProvisionConfigResponse) SetHeaders(v map[string]*string) *GetProvisionConfigResponse {
	s.Headers = v
	return s
}

func (s *GetProvisionConfigResponse) SetStatusCode(v int32) *GetProvisionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProvisionConfigResponse) SetBody(v *ProvisionConfig) *GetProvisionConfigResponse {
	s.Body = v
	return s
}

func (s *GetProvisionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

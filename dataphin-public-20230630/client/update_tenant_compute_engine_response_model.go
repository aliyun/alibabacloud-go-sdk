// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantComputeEngineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTenantComputeEngineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTenantComputeEngineResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTenantComputeEngineResponseBody) *UpdateTenantComputeEngineResponse
	GetBody() *UpdateTenantComputeEngineResponseBody
}

type UpdateTenantComputeEngineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTenantComputeEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTenantComputeEngineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantComputeEngineResponse) GoString() string {
	return s.String()
}

func (s *UpdateTenantComputeEngineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTenantComputeEngineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTenantComputeEngineResponse) GetBody() *UpdateTenantComputeEngineResponseBody {
	return s.Body
}

func (s *UpdateTenantComputeEngineResponse) SetHeaders(v map[string]*string) *UpdateTenantComputeEngineResponse {
	s.Headers = v
	return s
}

func (s *UpdateTenantComputeEngineResponse) SetStatusCode(v int32) *UpdateTenantComputeEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTenantComputeEngineResponse) SetBody(v *UpdateTenantComputeEngineResponseBody) *UpdateTenantComputeEngineResponse {
	s.Body = v
	return s
}

func (s *UpdateTenantComputeEngineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

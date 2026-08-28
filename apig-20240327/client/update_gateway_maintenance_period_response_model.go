// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayMaintenancePeriodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayMaintenancePeriodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayMaintenancePeriodResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayMaintenancePeriodResponseBody) *UpdateGatewayMaintenancePeriodResponse
	GetBody() *UpdateGatewayMaintenancePeriodResponseBody
}

type UpdateGatewayMaintenancePeriodResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayMaintenancePeriodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayMaintenancePeriodResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayMaintenancePeriodResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayMaintenancePeriodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayMaintenancePeriodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayMaintenancePeriodResponse) GetBody() *UpdateGatewayMaintenancePeriodResponseBody {
	return s.Body
}

func (s *UpdateGatewayMaintenancePeriodResponse) SetHeaders(v map[string]*string) *UpdateGatewayMaintenancePeriodResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayMaintenancePeriodResponse) SetStatusCode(v int32) *UpdateGatewayMaintenancePeriodResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayMaintenancePeriodResponse) SetBody(v *UpdateGatewayMaintenancePeriodResponseBody) *UpdateGatewayMaintenancePeriodResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayMaintenancePeriodResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

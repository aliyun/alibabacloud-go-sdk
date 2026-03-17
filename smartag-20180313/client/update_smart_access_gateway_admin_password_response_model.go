// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayAdminPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayAdminPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAccessGatewayAdminPasswordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAccessGatewayAdminPasswordResponseBody) *UpdateSmartAccessGatewayAdminPasswordResponse
	GetBody() *UpdateSmartAccessGatewayAdminPasswordResponseBody
}

type UpdateSmartAccessGatewayAdminPasswordResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAccessGatewayAdminPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAccessGatewayAdminPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayAdminPasswordResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponse) GetBody() *UpdateSmartAccessGatewayAdminPasswordResponseBody {
	return s.Body
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponse) SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayAdminPasswordResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponse) SetStatusCode(v int32) *UpdateSmartAccessGatewayAdminPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponse) SetBody(v *UpdateSmartAccessGatewayAdminPasswordResponseBody) *UpdateSmartAccessGatewayAdminPasswordResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAccessGatewayAdminPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

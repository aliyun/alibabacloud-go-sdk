// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateASMGatewayImportedServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateASMGatewayImportedServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateASMGatewayImportedServicesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateASMGatewayImportedServicesResponseBody) *UpdateASMGatewayImportedServicesResponse
	GetBody() *UpdateASMGatewayImportedServicesResponseBody
}

type UpdateASMGatewayImportedServicesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateASMGatewayImportedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateASMGatewayImportedServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateASMGatewayImportedServicesResponse) GoString() string {
	return s.String()
}

func (s *UpdateASMGatewayImportedServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateASMGatewayImportedServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateASMGatewayImportedServicesResponse) GetBody() *UpdateASMGatewayImportedServicesResponseBody {
	return s.Body
}

func (s *UpdateASMGatewayImportedServicesResponse) SetHeaders(v map[string]*string) *UpdateASMGatewayImportedServicesResponse {
	s.Headers = v
	return s
}

func (s *UpdateASMGatewayImportedServicesResponse) SetStatusCode(v int32) *UpdateASMGatewayImportedServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateASMGatewayImportedServicesResponse) SetBody(v *UpdateASMGatewayImportedServicesResponseBody) *UpdateASMGatewayImportedServicesResponse {
	s.Body = v
	return s
}

func (s *UpdateASMGatewayImportedServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

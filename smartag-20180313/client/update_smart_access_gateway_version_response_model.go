// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAccessGatewayVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAccessGatewayVersionResponseBody) *UpdateSmartAccessGatewayVersionResponse
	GetBody() *UpdateSmartAccessGatewayVersionResponseBody
}

type UpdateSmartAccessGatewayVersionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAccessGatewayVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAccessGatewayVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAccessGatewayVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAccessGatewayVersionResponse) GetBody() *UpdateSmartAccessGatewayVersionResponseBody {
	return s.Body
}

func (s *UpdateSmartAccessGatewayVersionResponse) SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAccessGatewayVersionResponse) SetStatusCode(v int32) *UpdateSmartAccessGatewayVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAccessGatewayVersionResponse) SetBody(v *UpdateSmartAccessGatewayVersionResponseBody) *UpdateSmartAccessGatewayVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAccessGatewayVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

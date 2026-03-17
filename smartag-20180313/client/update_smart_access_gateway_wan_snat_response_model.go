// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayWanSnatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayWanSnatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAccessGatewayWanSnatResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAccessGatewayWanSnatResponseBody) *UpdateSmartAccessGatewayWanSnatResponse
	GetBody() *UpdateSmartAccessGatewayWanSnatResponseBody
}

type UpdateSmartAccessGatewayWanSnatResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAccessGatewayWanSnatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAccessGatewayWanSnatResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayWanSnatResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayWanSnatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAccessGatewayWanSnatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAccessGatewayWanSnatResponse) GetBody() *UpdateSmartAccessGatewayWanSnatResponseBody {
	return s.Body
}

func (s *UpdateSmartAccessGatewayWanSnatResponse) SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayWanSnatResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatResponse) SetStatusCode(v int32) *UpdateSmartAccessGatewayWanSnatResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatResponse) SetBody(v *UpdateSmartAccessGatewayWanSnatResponseBody) *UpdateSmartAccessGatewayWanSnatResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

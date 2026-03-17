// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayDnsForwardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayDnsForwardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAccessGatewayDnsForwardResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAccessGatewayDnsForwardResponseBody) *UpdateSmartAccessGatewayDnsForwardResponse
	GetBody() *UpdateSmartAccessGatewayDnsForwardResponseBody
}

type UpdateSmartAccessGatewayDnsForwardResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAccessGatewayDnsForwardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAccessGatewayDnsForwardResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayDnsForwardResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayDnsForwardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAccessGatewayDnsForwardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAccessGatewayDnsForwardResponse) GetBody() *UpdateSmartAccessGatewayDnsForwardResponseBody {
	return s.Body
}

func (s *UpdateSmartAccessGatewayDnsForwardResponse) SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayDnsForwardResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAccessGatewayDnsForwardResponse) SetStatusCode(v int32) *UpdateSmartAccessGatewayDnsForwardResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsForwardResponse) SetBody(v *UpdateSmartAccessGatewayDnsForwardResponseBody) *UpdateSmartAccessGatewayDnsForwardResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAccessGatewayDnsForwardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

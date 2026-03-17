// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayDnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayDnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAccessGatewayDnsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAccessGatewayDnsResponseBody) *UpdateSmartAccessGatewayDnsResponse
	GetBody() *UpdateSmartAccessGatewayDnsResponseBody
}

type UpdateSmartAccessGatewayDnsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAccessGatewayDnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAccessGatewayDnsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayDnsResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayDnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAccessGatewayDnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAccessGatewayDnsResponse) GetBody() *UpdateSmartAccessGatewayDnsResponseBody {
	return s.Body
}

func (s *UpdateSmartAccessGatewayDnsResponse) SetHeaders(v map[string]*string) *UpdateSmartAccessGatewayDnsResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAccessGatewayDnsResponse) SetStatusCode(v int32) *UpdateSmartAccessGatewayDnsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsResponse) SetBody(v *UpdateSmartAccessGatewayDnsResponseBody) *UpdateSmartAccessGatewayDnsResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAccessGatewayDnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

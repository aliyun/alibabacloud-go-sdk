// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmartAccessGatewayDnsForwardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSmartAccessGatewayDnsForwardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSmartAccessGatewayDnsForwardResponse
	GetStatusCode() *int32
	SetBody(v *AddSmartAccessGatewayDnsForwardResponseBody) *AddSmartAccessGatewayDnsForwardResponse
	GetBody() *AddSmartAccessGatewayDnsForwardResponseBody
}

type AddSmartAccessGatewayDnsForwardResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSmartAccessGatewayDnsForwardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSmartAccessGatewayDnsForwardResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSmartAccessGatewayDnsForwardResponse) GoString() string {
	return s.String()
}

func (s *AddSmartAccessGatewayDnsForwardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSmartAccessGatewayDnsForwardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSmartAccessGatewayDnsForwardResponse) GetBody() *AddSmartAccessGatewayDnsForwardResponseBody {
	return s.Body
}

func (s *AddSmartAccessGatewayDnsForwardResponse) SetHeaders(v map[string]*string) *AddSmartAccessGatewayDnsForwardResponse {
	s.Headers = v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponse) SetStatusCode(v int32) *AddSmartAccessGatewayDnsForwardResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponse) SetBody(v *AddSmartAccessGatewayDnsForwardResponseBody) *AddSmartAccessGatewayDnsForwardResponse {
	s.Body = v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

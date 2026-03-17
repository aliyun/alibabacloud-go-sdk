// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartAccessGatewayDnsForwardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSmartAccessGatewayDnsForwardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSmartAccessGatewayDnsForwardResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSmartAccessGatewayDnsForwardResponseBody) *DeleteSmartAccessGatewayDnsForwardResponse
	GetBody() *DeleteSmartAccessGatewayDnsForwardResponseBody
}

type DeleteSmartAccessGatewayDnsForwardResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSmartAccessGatewayDnsForwardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSmartAccessGatewayDnsForwardResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartAccessGatewayDnsForwardResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmartAccessGatewayDnsForwardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSmartAccessGatewayDnsForwardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSmartAccessGatewayDnsForwardResponse) GetBody() *DeleteSmartAccessGatewayDnsForwardResponseBody {
	return s.Body
}

func (s *DeleteSmartAccessGatewayDnsForwardResponse) SetHeaders(v map[string]*string) *DeleteSmartAccessGatewayDnsForwardResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardResponse) SetStatusCode(v int32) *DeleteSmartAccessGatewayDnsForwardResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardResponse) SetBody(v *DeleteSmartAccessGatewayDnsForwardResponseBody) *DeleteSmartAccessGatewayDnsForwardResponse {
	s.Body = v
	return s
}

func (s *DeleteSmartAccessGatewayDnsForwardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

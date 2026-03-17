// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayDnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ViewSmartAccessGatewayDnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ViewSmartAccessGatewayDnsResponse
	GetStatusCode() *int32
	SetBody(v *ViewSmartAccessGatewayDnsResponseBody) *ViewSmartAccessGatewayDnsResponse
	GetBody() *ViewSmartAccessGatewayDnsResponseBody
}

type ViewSmartAccessGatewayDnsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ViewSmartAccessGatewayDnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ViewSmartAccessGatewayDnsResponse) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayDnsResponse) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayDnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ViewSmartAccessGatewayDnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ViewSmartAccessGatewayDnsResponse) GetBody() *ViewSmartAccessGatewayDnsResponseBody {
	return s.Body
}

func (s *ViewSmartAccessGatewayDnsResponse) SetHeaders(v map[string]*string) *ViewSmartAccessGatewayDnsResponse {
	s.Headers = v
	return s
}

func (s *ViewSmartAccessGatewayDnsResponse) SetStatusCode(v int32) *ViewSmartAccessGatewayDnsResponse {
	s.StatusCode = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsResponse) SetBody(v *ViewSmartAccessGatewayDnsResponseBody) *ViewSmartAccessGatewayDnsResponse {
	s.Body = v
	return s
}

func (s *ViewSmartAccessGatewayDnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

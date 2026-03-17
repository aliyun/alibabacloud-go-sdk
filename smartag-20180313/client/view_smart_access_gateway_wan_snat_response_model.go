// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayWanSnatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ViewSmartAccessGatewayWanSnatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ViewSmartAccessGatewayWanSnatResponse
	GetStatusCode() *int32
	SetBody(v *ViewSmartAccessGatewayWanSnatResponseBody) *ViewSmartAccessGatewayWanSnatResponse
	GetBody() *ViewSmartAccessGatewayWanSnatResponseBody
}

type ViewSmartAccessGatewayWanSnatResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ViewSmartAccessGatewayWanSnatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ViewSmartAccessGatewayWanSnatResponse) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayWanSnatResponse) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayWanSnatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ViewSmartAccessGatewayWanSnatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ViewSmartAccessGatewayWanSnatResponse) GetBody() *ViewSmartAccessGatewayWanSnatResponseBody {
	return s.Body
}

func (s *ViewSmartAccessGatewayWanSnatResponse) SetHeaders(v map[string]*string) *ViewSmartAccessGatewayWanSnatResponse {
	s.Headers = v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatResponse) SetStatusCode(v int32) *ViewSmartAccessGatewayWanSnatResponse {
	s.StatusCode = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatResponse) SetBody(v *ViewSmartAccessGatewayWanSnatResponseBody) *ViewSmartAccessGatewayWanSnatResponse {
	s.Body = v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

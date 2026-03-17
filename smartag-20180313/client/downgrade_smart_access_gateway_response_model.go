// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradeSmartAccessGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DowngradeSmartAccessGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DowngradeSmartAccessGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DowngradeSmartAccessGatewayResponseBody) *DowngradeSmartAccessGatewayResponse
	GetBody() *DowngradeSmartAccessGatewayResponseBody
}

type DowngradeSmartAccessGatewayResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DowngradeSmartAccessGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DowngradeSmartAccessGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DowngradeSmartAccessGatewayResponse) GoString() string {
	return s.String()
}

func (s *DowngradeSmartAccessGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DowngradeSmartAccessGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DowngradeSmartAccessGatewayResponse) GetBody() *DowngradeSmartAccessGatewayResponseBody {
	return s.Body
}

func (s *DowngradeSmartAccessGatewayResponse) SetHeaders(v map[string]*string) *DowngradeSmartAccessGatewayResponse {
	s.Headers = v
	return s
}

func (s *DowngradeSmartAccessGatewayResponse) SetStatusCode(v int32) *DowngradeSmartAccessGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DowngradeSmartAccessGatewayResponse) SetBody(v *DowngradeSmartAccessGatewayResponseBody) *DowngradeSmartAccessGatewayResponse {
	s.Body = v
	return s
}

func (s *DowngradeSmartAccessGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

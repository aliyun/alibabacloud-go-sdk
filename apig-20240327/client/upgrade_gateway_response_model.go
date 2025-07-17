// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeGatewayResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeGatewayResponseBody) *UpgradeGatewayResponse
	GetBody() *UpgradeGatewayResponseBody
}

type UpgradeGatewayResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpgradeGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeGatewayResponse) GetBody() *UpgradeGatewayResponseBody {
	return s.Body
}

func (s *UpgradeGatewayResponse) SetHeaders(v map[string]*string) *UpgradeGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpgradeGatewayResponse) SetStatusCode(v int32) *UpgradeGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeGatewayResponse) SetBody(v *UpgradeGatewayResponseBody) *UpgradeGatewayResponse {
	s.Body = v
	return s
}

func (s *UpgradeGatewayResponse) Validate() error {
	return dara.Validate(s)
}

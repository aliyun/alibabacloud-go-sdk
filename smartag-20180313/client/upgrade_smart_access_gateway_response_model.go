// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeSmartAccessGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeSmartAccessGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeSmartAccessGatewayResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeSmartAccessGatewayResponseBody) *UpgradeSmartAccessGatewayResponse
	GetBody() *UpgradeSmartAccessGatewayResponseBody
}

type UpgradeSmartAccessGatewayResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeSmartAccessGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeSmartAccessGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeSmartAccessGatewayResponse) GoString() string {
	return s.String()
}

func (s *UpgradeSmartAccessGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeSmartAccessGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeSmartAccessGatewayResponse) GetBody() *UpgradeSmartAccessGatewayResponseBody {
	return s.Body
}

func (s *UpgradeSmartAccessGatewayResponse) SetHeaders(v map[string]*string) *UpgradeSmartAccessGatewayResponse {
	s.Headers = v
	return s
}

func (s *UpgradeSmartAccessGatewayResponse) SetStatusCode(v int32) *UpgradeSmartAccessGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeSmartAccessGatewayResponse) SetBody(v *UpgradeSmartAccessGatewayResponseBody) *UpgradeSmartAccessGatewayResponse {
	s.Body = v
	return s
}

func (s *UpgradeSmartAccessGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

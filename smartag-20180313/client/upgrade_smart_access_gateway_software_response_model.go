// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeSmartAccessGatewaySoftwareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeSmartAccessGatewaySoftwareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeSmartAccessGatewaySoftwareResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeSmartAccessGatewaySoftwareResponseBody) *UpgradeSmartAccessGatewaySoftwareResponse
	GetBody() *UpgradeSmartAccessGatewaySoftwareResponseBody
}

type UpgradeSmartAccessGatewaySoftwareResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeSmartAccessGatewaySoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeSmartAccessGatewaySoftwareResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeSmartAccessGatewaySoftwareResponse) GoString() string {
	return s.String()
}

func (s *UpgradeSmartAccessGatewaySoftwareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeSmartAccessGatewaySoftwareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeSmartAccessGatewaySoftwareResponse) GetBody() *UpgradeSmartAccessGatewaySoftwareResponseBody {
	return s.Body
}

func (s *UpgradeSmartAccessGatewaySoftwareResponse) SetHeaders(v map[string]*string) *UpgradeSmartAccessGatewaySoftwareResponse {
	s.Headers = v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareResponse) SetStatusCode(v int32) *UpgradeSmartAccessGatewaySoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareResponse) SetBody(v *UpgradeSmartAccessGatewaySoftwareResponseBody) *UpgradeSmartAccessGatewaySoftwareResponse {
	s.Body = v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

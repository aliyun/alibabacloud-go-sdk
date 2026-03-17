// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradeSmartAccessGatewaySoftwareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DowngradeSmartAccessGatewaySoftwareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DowngradeSmartAccessGatewaySoftwareResponse
	GetStatusCode() *int32
	SetBody(v *DowngradeSmartAccessGatewaySoftwareResponseBody) *DowngradeSmartAccessGatewaySoftwareResponse
	GetBody() *DowngradeSmartAccessGatewaySoftwareResponseBody
}

type DowngradeSmartAccessGatewaySoftwareResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DowngradeSmartAccessGatewaySoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DowngradeSmartAccessGatewaySoftwareResponse) String() string {
	return dara.Prettify(s)
}

func (s DowngradeSmartAccessGatewaySoftwareResponse) GoString() string {
	return s.String()
}

func (s *DowngradeSmartAccessGatewaySoftwareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DowngradeSmartAccessGatewaySoftwareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DowngradeSmartAccessGatewaySoftwareResponse) GetBody() *DowngradeSmartAccessGatewaySoftwareResponseBody {
	return s.Body
}

func (s *DowngradeSmartAccessGatewaySoftwareResponse) SetHeaders(v map[string]*string) *DowngradeSmartAccessGatewaySoftwareResponse {
	s.Headers = v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareResponse) SetStatusCode(v int32) *DowngradeSmartAccessGatewaySoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareResponse) SetBody(v *DowngradeSmartAccessGatewaySoftwareResponseBody) *DowngradeSmartAccessGatewaySoftwareResponse {
	s.Body = v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempUpgradeCenBandwidthPackageSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TempUpgradeCenBandwidthPackageSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TempUpgradeCenBandwidthPackageSpecResponse
	GetStatusCode() *int32
	SetBody(v *TempUpgradeCenBandwidthPackageSpecResponseBody) *TempUpgradeCenBandwidthPackageSpecResponse
	GetBody() *TempUpgradeCenBandwidthPackageSpecResponseBody
}

type TempUpgradeCenBandwidthPackageSpecResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TempUpgradeCenBandwidthPackageSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TempUpgradeCenBandwidthPackageSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s TempUpgradeCenBandwidthPackageSpecResponse) GoString() string {
	return s.String()
}

func (s *TempUpgradeCenBandwidthPackageSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TempUpgradeCenBandwidthPackageSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TempUpgradeCenBandwidthPackageSpecResponse) GetBody() *TempUpgradeCenBandwidthPackageSpecResponseBody {
	return s.Body
}

func (s *TempUpgradeCenBandwidthPackageSpecResponse) SetHeaders(v map[string]*string) *TempUpgradeCenBandwidthPackageSpecResponse {
	s.Headers = v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecResponse) SetStatusCode(v int32) *TempUpgradeCenBandwidthPackageSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecResponse) SetBody(v *TempUpgradeCenBandwidthPackageSpecResponseBody) *TempUpgradeCenBandwidthPackageSpecResponse {
	s.Body = v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecResponse) Validate() error {
	return dara.Validate(s)
}

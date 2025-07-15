// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCommonBandwidthPackageIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCommonBandwidthPackageIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCommonBandwidthPackageIpResponse
	GetStatusCode() *int32
	SetBody(v *AddCommonBandwidthPackageIpResponseBody) *AddCommonBandwidthPackageIpResponse
	GetBody() *AddCommonBandwidthPackageIpResponseBody
}

type AddCommonBandwidthPackageIpResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCommonBandwidthPackageIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCommonBandwidthPackageIpResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCommonBandwidthPackageIpResponse) GoString() string {
	return s.String()
}

func (s *AddCommonBandwidthPackageIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCommonBandwidthPackageIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCommonBandwidthPackageIpResponse) GetBody() *AddCommonBandwidthPackageIpResponseBody {
	return s.Body
}

func (s *AddCommonBandwidthPackageIpResponse) SetHeaders(v map[string]*string) *AddCommonBandwidthPackageIpResponse {
	s.Headers = v
	return s
}

func (s *AddCommonBandwidthPackageIpResponse) SetStatusCode(v int32) *AddCommonBandwidthPackageIpResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCommonBandwidthPackageIpResponse) SetBody(v *AddCommonBandwidthPackageIpResponseBody) *AddCommonBandwidthPackageIpResponse {
	s.Body = v
	return s
}

func (s *AddCommonBandwidthPackageIpResponse) Validate() error {
	return dara.Validate(s)
}

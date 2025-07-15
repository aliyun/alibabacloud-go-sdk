// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCommonBandwidthPackageIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCommonBandwidthPackageIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCommonBandwidthPackageIpsResponse
	GetStatusCode() *int32
	SetBody(v *AddCommonBandwidthPackageIpsResponseBody) *AddCommonBandwidthPackageIpsResponse
	GetBody() *AddCommonBandwidthPackageIpsResponseBody
}

type AddCommonBandwidthPackageIpsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCommonBandwidthPackageIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCommonBandwidthPackageIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCommonBandwidthPackageIpsResponse) GoString() string {
	return s.String()
}

func (s *AddCommonBandwidthPackageIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCommonBandwidthPackageIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCommonBandwidthPackageIpsResponse) GetBody() *AddCommonBandwidthPackageIpsResponseBody {
	return s.Body
}

func (s *AddCommonBandwidthPackageIpsResponse) SetHeaders(v map[string]*string) *AddCommonBandwidthPackageIpsResponse {
	s.Headers = v
	return s
}

func (s *AddCommonBandwidthPackageIpsResponse) SetStatusCode(v int32) *AddCommonBandwidthPackageIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCommonBandwidthPackageIpsResponse) SetBody(v *AddCommonBandwidthPackageIpsResponseBody) *AddCommonBandwidthPackageIpsResponse {
	s.Body = v
	return s
}

func (s *AddCommonBandwidthPackageIpsResponse) Validate() error {
	return dara.Validate(s)
}

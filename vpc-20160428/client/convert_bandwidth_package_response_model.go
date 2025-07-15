// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *ConvertBandwidthPackageResponseBody) *ConvertBandwidthPackageResponse
	GetBody() *ConvertBandwidthPackageResponseBody
}

type ConvertBandwidthPackageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *ConvertBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertBandwidthPackageResponse) GetBody() *ConvertBandwidthPackageResponseBody {
	return s.Body
}

func (s *ConvertBandwidthPackageResponse) SetHeaders(v map[string]*string) *ConvertBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *ConvertBandwidthPackageResponse) SetStatusCode(v int32) *ConvertBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertBandwidthPackageResponse) SetBody(v *ConvertBandwidthPackageResponseBody) *ConvertBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *ConvertBandwidthPackageResponse) Validate() error {
	return dara.Validate(s)
}

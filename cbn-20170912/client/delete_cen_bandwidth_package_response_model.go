// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCenBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCenBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCenBandwidthPackageResponseBody) *DeleteCenBandwidthPackageResponse
	GetBody() *DeleteCenBandwidthPackageResponseBody
}

type DeleteCenBandwidthPackageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCenBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCenBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCenBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCenBandwidthPackageResponse) GetBody() *DeleteCenBandwidthPackageResponseBody {
	return s.Body
}

func (s *DeleteCenBandwidthPackageResponse) SetHeaders(v map[string]*string) *DeleteCenBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenBandwidthPackageResponse) SetStatusCode(v int32) *DeleteCenBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCenBandwidthPackageResponse) SetBody(v *DeleteCenBandwidthPackageResponseBody) *DeleteCenBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *DeleteCenBandwidthPackageResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCenBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCenBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *CreateCenBandwidthPackageResponseBody) *CreateCenBandwidthPackageResponse
	GetBody() *CreateCenBandwidthPackageResponseBody
}

type CreateCenBandwidthPackageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCenBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCenBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCenBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateCenBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCenBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCenBandwidthPackageResponse) GetBody() *CreateCenBandwidthPackageResponseBody {
	return s.Body
}

func (s *CreateCenBandwidthPackageResponse) SetHeaders(v map[string]*string) *CreateCenBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateCenBandwidthPackageResponse) SetStatusCode(v int32) *CreateCenBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCenBandwidthPackageResponse) SetBody(v *CreateCenBandwidthPackageResponseBody) *CreateCenBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *CreateCenBandwidthPackageResponse) Validate() error {
	return dara.Validate(s)
}

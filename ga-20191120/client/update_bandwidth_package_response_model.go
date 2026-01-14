// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBandwidthPackageResponseBody) *UpdateBandwidthPackageResponse
	GetBody() *UpdateBandwidthPackageResponseBody
}

type UpdateBandwidthPackageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBandwidthPackageResponse) GetBody() *UpdateBandwidthPackageResponseBody {
	return s.Body
}

func (s *UpdateBandwidthPackageResponse) SetHeaders(v map[string]*string) *UpdateBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *UpdateBandwidthPackageResponse) SetStatusCode(v int32) *UpdateBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBandwidthPackageResponse) SetBody(v *UpdateBandwidthPackageResponseBody) *UpdateBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *UpdateBandwidthPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

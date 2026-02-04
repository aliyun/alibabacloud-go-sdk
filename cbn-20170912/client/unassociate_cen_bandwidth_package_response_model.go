// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateCenBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassociateCenBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassociateCenBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *UnassociateCenBandwidthPackageResponseBody) *UnassociateCenBandwidthPackageResponse
	GetBody() *UnassociateCenBandwidthPackageResponseBody
}

type UnassociateCenBandwidthPackageResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociateCenBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassociateCenBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassociateCenBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *UnassociateCenBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassociateCenBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassociateCenBandwidthPackageResponse) GetBody() *UnassociateCenBandwidthPackageResponseBody {
	return s.Body
}

func (s *UnassociateCenBandwidthPackageResponse) SetHeaders(v map[string]*string) *UnassociateCenBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *UnassociateCenBandwidthPackageResponse) SetStatusCode(v int32) *UnassociateCenBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociateCenBandwidthPackageResponse) SetBody(v *UnassociateCenBandwidthPackageResponseBody) *UnassociateCenBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *UnassociateCenBandwidthPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

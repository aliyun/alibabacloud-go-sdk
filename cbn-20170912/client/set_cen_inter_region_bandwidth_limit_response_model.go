// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCenInterRegionBandwidthLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCenInterRegionBandwidthLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCenInterRegionBandwidthLimitResponse
	GetStatusCode() *int32
	SetBody(v *SetCenInterRegionBandwidthLimitResponseBody) *SetCenInterRegionBandwidthLimitResponse
	GetBody() *SetCenInterRegionBandwidthLimitResponseBody
}

type SetCenInterRegionBandwidthLimitResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCenInterRegionBandwidthLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCenInterRegionBandwidthLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCenInterRegionBandwidthLimitResponse) GoString() string {
	return s.String()
}

func (s *SetCenInterRegionBandwidthLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCenInterRegionBandwidthLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCenInterRegionBandwidthLimitResponse) GetBody() *SetCenInterRegionBandwidthLimitResponseBody {
	return s.Body
}

func (s *SetCenInterRegionBandwidthLimitResponse) SetHeaders(v map[string]*string) *SetCenInterRegionBandwidthLimitResponse {
	s.Headers = v
	return s
}

func (s *SetCenInterRegionBandwidthLimitResponse) SetStatusCode(v int32) *SetCenInterRegionBandwidthLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitResponse) SetBody(v *SetCenInterRegionBandwidthLimitResponseBody) *SetCenInterRegionBandwidthLimitResponse {
	s.Body = v
	return s
}

func (s *SetCenInterRegionBandwidthLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

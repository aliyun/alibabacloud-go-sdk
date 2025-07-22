// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeExpressConnectRouterInterRegionBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse
	GetBody() *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody
}

type SynchronizeExpressConnectRouterInterRegionBandwidthResponse struct {
	Headers    map[string]*string                                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponse) GoString() string {
	return s.String()
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) GetBody() *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	return s.Body
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) SetHeaders(v map[string]*string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse {
	s.Headers = v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) SetStatusCode(v int32) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) SetBody(v *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse {
	s.Body = v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) Validate() error {
	return dara.Validate(s)
}

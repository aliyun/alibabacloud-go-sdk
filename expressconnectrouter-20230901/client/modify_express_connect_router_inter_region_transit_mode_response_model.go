// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterInterRegionTransitModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExpressConnectRouterInterRegionTransitModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExpressConnectRouterInterRegionTransitModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) *ModifyExpressConnectRouterInterRegionTransitModeResponse
	GetBody() *ModifyExpressConnectRouterInterRegionTransitModeResponseBody
}

type ModifyExpressConnectRouterInterRegionTransitModeResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectRouterInterRegionTransitModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectRouterInterRegionTransitModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) GetBody() *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	return s.Body
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectRouterInterRegionTransitModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) SetStatusCode(v int32) *ModifyExpressConnectRouterInterRegionTransitModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) SetBody(v *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) *ModifyExpressConnectRouterInterRegionTransitModeResponse {
	s.Body = v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) Validate() error {
	return dara.Validate(s)
}

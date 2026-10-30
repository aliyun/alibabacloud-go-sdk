// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteMapRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRouteMapRuleResponseBody
	GetRequestId() *string
	SetRouteMapRuleId(v string) *CreateRouteMapRuleResponseBody
	GetRouteMapRuleId() *string
}

type CreateRouteMapRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EF44D7B7-6248-5A60-8A4C-A92388605A57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the route map rule.
	//
	// example:
	//
	// rmr-3dm13f66wtctcnydup
	RouteMapRuleId *string `json:"RouteMapRuleId,omitempty" xml:"RouteMapRuleId,omitempty"`
}

func (s CreateRouteMapRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteMapRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRouteMapRuleResponseBody) GetRouteMapRuleId() *string {
	return s.RouteMapRuleId
}

func (s *CreateRouteMapRuleResponseBody) SetRequestId(v string) *CreateRouteMapRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouteMapRuleResponseBody) SetRouteMapRuleId(v string) *CreateRouteMapRuleResponseBody {
	s.RouteMapRuleId = &v
	return s
}

func (s *CreateRouteMapRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

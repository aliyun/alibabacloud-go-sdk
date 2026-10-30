// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteMapRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRouteMapRuleResponseBody
	GetRequestId() *string
}

type DeleteRouteMapRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EF44D7B7-6248-5A60-8A4C-A92388605A57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRouteMapRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteMapRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteMapRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteMapRuleResponseBody) SetRequestId(v string) *DeleteRouteMapRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteMapRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

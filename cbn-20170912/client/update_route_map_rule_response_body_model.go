// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteMapRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRouteMapRuleResponseBody
	GetRequestId() *string
}

type UpdateRouteMapRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EF44D7B7-6248-5A60-8A4C-A92388605A57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRouteMapRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteMapRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRouteMapRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRouteMapRuleResponseBody) SetRequestId(v string) *UpdateRouteMapRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRouteMapRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouteDistributionStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRouteDistributionStrategyResponseBody
	GetRequestId() *string
}

type ModifyRouteDistributionStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 24C58BD0-1679-4942-9D42-00B635DAAADB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRouteDistributionStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouteDistributionStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRouteDistributionStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRouteDistributionStrategyResponseBody) SetRequestId(v string) *ModifyRouteDistributionStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRouteDistributionStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}

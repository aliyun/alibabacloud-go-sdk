// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoGroupingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAutoGroupingRuleResponseBody
	GetRequestId() *string
}

type UpdateAutoGroupingRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAutoGroupingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoGroupingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoGroupingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAutoGroupingRuleResponseBody) SetRequestId(v string) *UpdateAutoGroupingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoGroupingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCostRuleResponseBody
	GetRequestId() *string
}

type DeleteCostRuleResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7E2FE3BB-C677-5FF9-9FC5-9CF364BD6BE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCostRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCostRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCostRuleResponseBody) SetRequestId(v string) *DeleteCostRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCostRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

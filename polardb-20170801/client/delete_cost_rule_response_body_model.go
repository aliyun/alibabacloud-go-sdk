// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedModelApiIds(v []*string) *DeleteCostRuleResponseBody
	GetAffectedModelApiIds() []*string
	SetRequestId(v string) *DeleteCostRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCostRuleResponseBody
	GetSuccess() *bool
}

type DeleteCostRuleResponseBody struct {
	// The list of model API IDs affected by the deletion.
	//
	// example:
	//
	// ["ma-2ze3x6i1d7g5rto28"]
	AffectedModelApiIds []*string `json:"AffectedModelApiIds,omitempty" xml:"AffectedModelApiIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7E2FE3BB-C677-5FF9-9FC5-9CF364BD6BE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the deletion was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCostRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCostRuleResponseBody) GetAffectedModelApiIds() []*string {
	return s.AffectedModelApiIds
}

func (s *DeleteCostRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCostRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCostRuleResponseBody) SetAffectedModelApiIds(v []*string) *DeleteCostRuleResponseBody {
	s.AffectedModelApiIds = v
	return s
}

func (s *DeleteCostRuleResponseBody) SetRequestId(v string) *DeleteCostRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCostRuleResponseBody) SetSuccess(v bool) *DeleteCostRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCostRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

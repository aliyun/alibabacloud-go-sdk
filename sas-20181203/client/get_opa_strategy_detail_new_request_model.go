// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaStrategyDetailNewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyId(v int64) *GetOpaStrategyDetailNewRequest
	GetStrategyId() *int64
}

type GetOpaStrategyDetailNewRequest struct {
	// The rule ID.
	//
	// >  You can call the [ListOpaClusterStrategyNew](https://help.aliyun.com/document_detail/2623574.html) operation to query the rule ID.
	//
	// example:
	//
	// 1349
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s GetOpaStrategyDetailNewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyDetailNewRequest) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyDetailNewRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *GetOpaStrategyDetailNewRequest) SetStrategyId(v int64) *GetOpaStrategyDetailNewRequest {
	s.StrategyId = &v
	return s
}

func (s *GetOpaStrategyDetailNewRequest) Validate() error {
	return dara.Validate(s)
}

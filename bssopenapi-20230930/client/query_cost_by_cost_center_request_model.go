// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostByCostCenterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingMonth(v int32) *QueryCostByCostCenterRequest
	GetBillingMonth() *int32
	SetDisplayZeroAmountBills(v bool) *QueryCostByCostCenterRequest
	GetDisplayZeroAmountBills() *bool
	SetGroupByCostCenterLevel(v bool) *QueryCostByCostCenterRequest
	GetGroupByCostCenterLevel() *bool
	SetMetrics(v string) *QueryCostByCostCenterRequest
	GetMetrics() *string
	SetOwnerAccountId(v int64) *QueryCostByCostCenterRequest
	GetOwnerAccountId() *int64
}

type QueryCostByCostCenterRequest struct {
	// The billing cycle month in the format of YYYYMM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 202506
	BillingMonth *int32 `json:"BillingMonth,omitempty" xml:"BillingMonth,omitempty"`
	// Specifies whether to display data rows with a payable amount of 0.
	//
	// example:
	//
	// false
	DisplayZeroAmountBills *bool `json:"DisplayZeroAmountBills,omitempty" xml:"DisplayZeroAmountBills,omitempty"`
	// Specifies whether to display results grouped by financial unit level.
	//
	// example:
	//
	// false
	GroupByCostCenterLevel *bool `json:"GroupByCostCenterLevel,omitempty" xml:"GroupByCostCenterLevel,omitempty"`
	// The cost type.
	//
	// This parameter is required.
	//
	// example:
	//
	// REQUIRE_AMOUNT
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The account ID of the resource ownership.
	//
	// example:
	//
	// 1374729705039203
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s QueryCostByCostCenterRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCostByCostCenterRequest) GoString() string {
	return s.String()
}

func (s *QueryCostByCostCenterRequest) GetBillingMonth() *int32 {
	return s.BillingMonth
}

func (s *QueryCostByCostCenterRequest) GetDisplayZeroAmountBills() *bool {
	return s.DisplayZeroAmountBills
}

func (s *QueryCostByCostCenterRequest) GetGroupByCostCenterLevel() *bool {
	return s.GroupByCostCenterLevel
}

func (s *QueryCostByCostCenterRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *QueryCostByCostCenterRequest) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostByCostCenterRequest) SetBillingMonth(v int32) *QueryCostByCostCenterRequest {
	s.BillingMonth = &v
	return s
}

func (s *QueryCostByCostCenterRequest) SetDisplayZeroAmountBills(v bool) *QueryCostByCostCenterRequest {
	s.DisplayZeroAmountBills = &v
	return s
}

func (s *QueryCostByCostCenterRequest) SetGroupByCostCenterLevel(v bool) *QueryCostByCostCenterRequest {
	s.GroupByCostCenterLevel = &v
	return s
}

func (s *QueryCostByCostCenterRequest) SetMetrics(v string) *QueryCostByCostCenterRequest {
	s.Metrics = &v
	return s
}

func (s *QueryCostByCostCenterRequest) SetOwnerAccountId(v int64) *QueryCostByCostCenterRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostByCostCenterRequest) Validate() error {
	return dara.Validate(s)
}

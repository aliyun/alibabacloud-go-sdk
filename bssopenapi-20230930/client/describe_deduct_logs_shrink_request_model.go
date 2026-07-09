// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeductLogsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillInstanceId(v string) *DescribeDeductLogsShrinkRequest
	GetBillInstanceId() *string
	SetBillingCommodityCode(v string) *DescribeDeductLogsShrinkRequest
	GetBillingCommodityCode() *string
	SetBillingEndTime(v int64) *DescribeDeductLogsShrinkRequest
	GetBillingEndTime() *int64
	SetBillingStartTime(v int64) *DescribeDeductLogsShrinkRequest
	GetBillingStartTime() *int64
	SetCommodityCode(v string) *DescribeDeductLogsShrinkRequest
	GetCommodityCode() *string
	SetEcIdAccountIdsShrink(v string) *DescribeDeductLogsShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetGroup(v string) *DescribeDeductLogsShrinkRequest
	GetGroup() *string
	SetInstanceId(v string) *DescribeDeductLogsShrinkRequest
	GetInstanceId() *string
	SetNbid(v string) *DescribeDeductLogsShrinkRequest
	GetNbid() *string
	SetPageNum(v int32) *DescribeDeductLogsShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeDeductLogsShrinkRequest
	GetPageSize() *int32
	SetRelationAccountIdsShrink(v string) *DescribeDeductLogsShrinkRequest
	GetRelationAccountIdsShrink() *string
}

type DescribeDeductLogsShrinkRequest struct {
	// The instance ID for billing deduction.
	//
	// example:
	//
	// ecs-******
	BillInstanceId *string `json:"BillInstanceId,omitempty" xml:"BillInstanceId,omitempty"`
	// The commodity code of the deducted item.
	//
	// example:
	//
	// snapshot
	BillingCommodityCode *string `json:"BillingCommodityCode,omitempty" xml:"BillingCommodityCode,omitempty"`
	// The billing end time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1679068799999
	BillingEndTime *int64 `json:"BillingEndTime,omitempty" xml:"BillingEndTime,omitempty"`
	// The billing start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1647446400000
	BillingStartTime *int64 `json:"BillingStartTime,omitempty" xml:"BillingStartTime,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// ossbag
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The enterprise and account list. If this parameter is empty, the current account is queried.
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// The resource dimension for the query.
	//
	// example:
	//
	// oss_rc
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The instance name.
	//
	// example:
	//
	// OSSBAG-cn******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The primary marketplace ID. If this parameter is empty, the marketplace ID of the current user is used by default.
	//
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of deduction accounts.
	RelationAccountIdsShrink *string `json:"RelationAccountIds,omitempty" xml:"RelationAccountIds,omitempty"`
}

func (s DescribeDeductLogsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsShrinkRequest) GetBillInstanceId() *string {
	return s.BillInstanceId
}

func (s *DescribeDeductLogsShrinkRequest) GetBillingCommodityCode() *string {
	return s.BillingCommodityCode
}

func (s *DescribeDeductLogsShrinkRequest) GetBillingEndTime() *int64 {
	return s.BillingEndTime
}

func (s *DescribeDeductLogsShrinkRequest) GetBillingStartTime() *int64 {
	return s.BillingStartTime
}

func (s *DescribeDeductLogsShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDeductLogsShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *DescribeDeductLogsShrinkRequest) GetGroup() *string {
	return s.Group
}

func (s *DescribeDeductLogsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDeductLogsShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeDeductLogsShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeDeductLogsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeductLogsShrinkRequest) GetRelationAccountIdsShrink() *string {
	return s.RelationAccountIdsShrink
}

func (s *DescribeDeductLogsShrinkRequest) SetBillInstanceId(v string) *DescribeDeductLogsShrinkRequest {
	s.BillInstanceId = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) SetBillingCommodityCode(v string) *DescribeDeductLogsShrinkRequest {
	s.BillingCommodityCode = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) SetBillingEndTime(v int64) *DescribeDeductLogsShrinkRequest {
	s.BillingEndTime = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) SetBillingStartTime(v int64) *DescribeDeductLogsShrinkRequest {
	s.BillingStartTime = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) SetCommodityCode(v string) *DescribeDeductLogsShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) SetEcIdAccountIdsShrink(v string) *DescribeDeductLogsShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) SetGroup(v string) *DescribeDeductLogsShrinkRequest {
	s.Group = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) SetInstanceId(v string) *DescribeDeductLogsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) SetNbid(v string) *DescribeDeductLogsShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) SetPageNum(v int32) *DescribeDeductLogsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) SetPageSize(v int32) *DescribeDeductLogsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) SetRelationAccountIdsShrink(v string) *DescribeDeductLogsShrinkRequest {
	s.RelationAccountIdsShrink = &v
	return s
}

func (s *DescribeDeductLogsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeductLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillInstanceId(v string) *DescribeDeductLogsRequest
	GetBillInstanceId() *string
	SetBillingCommodityCode(v string) *DescribeDeductLogsRequest
	GetBillingCommodityCode() *string
	SetBillingEndTime(v int64) *DescribeDeductLogsRequest
	GetBillingEndTime() *int64
	SetBillingStartTime(v int64) *DescribeDeductLogsRequest
	GetBillingStartTime() *int64
	SetCommodityCode(v string) *DescribeDeductLogsRequest
	GetCommodityCode() *string
	SetEcIdAccountIds(v []*DescribeDeductLogsRequestEcIdAccountIds) *DescribeDeductLogsRequest
	GetEcIdAccountIds() []*DescribeDeductLogsRequestEcIdAccountIds
	SetGroup(v string) *DescribeDeductLogsRequest
	GetGroup() *string
	SetInstanceId(v string) *DescribeDeductLogsRequest
	GetInstanceId() *string
	SetNbid(v string) *DescribeDeductLogsRequest
	GetNbid() *string
	SetPageNum(v int32) *DescribeDeductLogsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeDeductLogsRequest
	GetPageSize() *int32
	SetRelationAccountIds(v []*int64) *DescribeDeductLogsRequest
	GetRelationAccountIds() []*int64
}

type DescribeDeductLogsRequest struct {
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
	EcIdAccountIds []*DescribeDeductLogsRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
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
	RelationAccountIds []*int64 `json:"RelationAccountIds,omitempty" xml:"RelationAccountIds,omitempty" type:"Repeated"`
}

func (s DescribeDeductLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsRequest) GetBillInstanceId() *string {
	return s.BillInstanceId
}

func (s *DescribeDeductLogsRequest) GetBillingCommodityCode() *string {
	return s.BillingCommodityCode
}

func (s *DescribeDeductLogsRequest) GetBillingEndTime() *int64 {
	return s.BillingEndTime
}

func (s *DescribeDeductLogsRequest) GetBillingStartTime() *int64 {
	return s.BillingStartTime
}

func (s *DescribeDeductLogsRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDeductLogsRequest) GetEcIdAccountIds() []*DescribeDeductLogsRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *DescribeDeductLogsRequest) GetGroup() *string {
	return s.Group
}

func (s *DescribeDeductLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDeductLogsRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeDeductLogsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeDeductLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeductLogsRequest) GetRelationAccountIds() []*int64 {
	return s.RelationAccountIds
}

func (s *DescribeDeductLogsRequest) SetBillInstanceId(v string) *DescribeDeductLogsRequest {
	s.BillInstanceId = &v
	return s
}

func (s *DescribeDeductLogsRequest) SetBillingCommodityCode(v string) *DescribeDeductLogsRequest {
	s.BillingCommodityCode = &v
	return s
}

func (s *DescribeDeductLogsRequest) SetBillingEndTime(v int64) *DescribeDeductLogsRequest {
	s.BillingEndTime = &v
	return s
}

func (s *DescribeDeductLogsRequest) SetBillingStartTime(v int64) *DescribeDeductLogsRequest {
	s.BillingStartTime = &v
	return s
}

func (s *DescribeDeductLogsRequest) SetCommodityCode(v string) *DescribeDeductLogsRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDeductLogsRequest) SetEcIdAccountIds(v []*DescribeDeductLogsRequestEcIdAccountIds) *DescribeDeductLogsRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *DescribeDeductLogsRequest) SetGroup(v string) *DescribeDeductLogsRequest {
	s.Group = &v
	return s
}

func (s *DescribeDeductLogsRequest) SetInstanceId(v string) *DescribeDeductLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDeductLogsRequest) SetNbid(v string) *DescribeDeductLogsRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeDeductLogsRequest) SetPageNum(v int32) *DescribeDeductLogsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeDeductLogsRequest) SetPageSize(v int32) *DescribeDeductLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeductLogsRequest) SetRelationAccountIds(v []*int64) *DescribeDeductLogsRequest {
	s.RelationAccountIds = v
	return s
}

func (s *DescribeDeductLogsRequest) Validate() error {
	if s.EcIdAccountIds != nil {
		for _, item := range s.EcIdAccountIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeductLogsRequestEcIdAccountIds struct {
	// The list of accounts to access. If this parameter is empty, all accounts under the current entity ID are selected.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The enterprise entity ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1004064243473974
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DescribeDeductLogsRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *DescribeDeductLogsRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *DescribeDeductLogsRequestEcIdAccountIds) SetAccountIds(v []*int64) *DescribeDeductLogsRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *DescribeDeductLogsRequestEcIdAccountIds) SetEcId(v string) *DescribeDeductLogsRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *DescribeDeductLogsRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}

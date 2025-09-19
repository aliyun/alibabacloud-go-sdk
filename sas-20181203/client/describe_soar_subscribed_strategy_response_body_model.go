// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarSubscribedStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSoarSubscribedStrategyResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSoarSubscribedStrategyResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSoarSubscribedStrategyResponseBody
	GetRequestId() *string
	SetSoarStrategies(v []*DescribeSoarSubscribedStrategyResponseBodySoarStrategies) *DescribeSoarSubscribedStrategyResponseBody
	GetSoarStrategies() []*DescribeSoarSubscribedStrategyResponseBodySoarStrategies
	SetTotalCount(v int32) *DescribeSoarSubscribedStrategyResponseBody
	GetTotalCount() *int32
}

type DescribeSoarSubscribedStrategyResponseBody struct {
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The policies.
	SoarStrategies []*DescribeSoarSubscribedStrategyResponseBodySoarStrategies `json:"SoarStrategies,omitempty" xml:"SoarStrategies,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 101
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSoarSubscribedStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarSubscribedStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarSubscribedStrategyResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSoarSubscribedStrategyResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSoarSubscribedStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSoarSubscribedStrategyResponseBody) GetSoarStrategies() []*DescribeSoarSubscribedStrategyResponseBodySoarStrategies {
	return s.SoarStrategies
}

func (s *DescribeSoarSubscribedStrategyResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSoarSubscribedStrategyResponseBody) SetPageNumber(v int32) *DescribeSoarSubscribedStrategyResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBody) SetPageSize(v int32) *DescribeSoarSubscribedStrategyResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBody) SetRequestId(v string) *DescribeSoarSubscribedStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBody) SetSoarStrategies(v []*DescribeSoarSubscribedStrategyResponseBodySoarStrategies) *DescribeSoarSubscribedStrategyResponseBody {
	s.SoarStrategies = v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBody) SetTotalCount(v int32) *DescribeSoarSubscribedStrategyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSoarSubscribedStrategyResponseBodySoarStrategies struct {
	// The Alibaba Cloud account ID of the creator.
	//
	// example:
	//
	// 1276085*****4392
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// strategy_description_01
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The total number of times that the policy is executed.
	//
	// example:
	//
	// 10
	ExecuteNum *string `json:"ExecuteNum,omitempty" xml:"ExecuteNum,omitempty"`
	// The timestamp when the policy was created. Unit: milliseconds.
	//
	// example:
	//
	// 1716344106000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the policy was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1652672104000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 300063
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// strategy_name01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution mode. Valid values:
	//
	// 	- runmode_TRIGGER_BY_USER: manually executed
	//
	// example:
	//
	// runmode_TRIGGER_BY_USER
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- type_vulfix: vulnerability operations
	//
	// example:
	//
	// type_vulfix
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSoarSubscribedStrategyResponseBodySoarStrategies) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarSubscribedStrategyResponseBodySoarStrategies) GoString() string {
	return s.String()
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) GetCreator() *string {
	return s.Creator
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) GetDescription() *string {
	return s.Description
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) GetExecuteNum() *string {
	return s.ExecuteNum
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) GetId() *int64 {
	return s.Id
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) GetName() *string {
	return s.Name
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) GetRunMode() *string {
	return s.RunMode
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) GetType() *string {
	return s.Type
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) SetCreator(v string) *DescribeSoarSubscribedStrategyResponseBodySoarStrategies {
	s.Creator = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) SetDescription(v string) *DescribeSoarSubscribedStrategyResponseBodySoarStrategies {
	s.Description = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) SetExecuteNum(v string) *DescribeSoarSubscribedStrategyResponseBodySoarStrategies {
	s.ExecuteNum = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) SetGmtCreate(v int64) *DescribeSoarSubscribedStrategyResponseBodySoarStrategies {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) SetGmtModified(v int64) *DescribeSoarSubscribedStrategyResponseBodySoarStrategies {
	s.GmtModified = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) SetId(v int64) *DescribeSoarSubscribedStrategyResponseBodySoarStrategies {
	s.Id = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) SetName(v string) *DescribeSoarSubscribedStrategyResponseBodySoarStrategies {
	s.Name = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) SetRunMode(v string) *DescribeSoarSubscribedStrategyResponseBodySoarStrategies {
	s.RunMode = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) SetType(v string) *DescribeSoarSubscribedStrategyResponseBodySoarStrategies {
	s.Type = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyResponseBodySoarStrategies) Validate() error {
	return dara.Validate(s)
}

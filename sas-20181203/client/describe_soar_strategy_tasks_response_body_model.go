// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSoarStrategyTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSoarStrategyTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSoarStrategyTasksResponseBody
	GetRequestId() *string
	SetSoarStrategyTasks(v []*DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) *DescribeSoarStrategyTasksResponseBody
	GetSoarStrategyTasks() []*DescribeSoarStrategyTasksResponseBodySoarStrategyTasks
	SetTotalCount(v int32) *DescribeSoarStrategyTasksResponseBody
	GetTotalCount() *int32
}

type DescribeSoarStrategyTasksResponseBody struct {
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
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The policy tasks.
	SoarStrategyTasks []*DescribeSoarStrategyTasksResponseBodySoarStrategyTasks `json:"SoarStrategyTasks,omitempty" xml:"SoarStrategyTasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSoarStrategyTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSoarStrategyTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSoarStrategyTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSoarStrategyTasksResponseBody) GetSoarStrategyTasks() []*DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	return s.SoarStrategyTasks
}

func (s *DescribeSoarStrategyTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSoarStrategyTasksResponseBody) SetPageNumber(v int32) *DescribeSoarStrategyTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBody) SetPageSize(v int32) *DescribeSoarStrategyTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBody) SetRequestId(v string) *DescribeSoarStrategyTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBody) SetSoarStrategyTasks(v []*DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) *DescribeSoarStrategyTasksResponseBody {
	s.SoarStrategyTasks = v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBody) SetTotalCount(v int32) *DescribeSoarStrategyTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBody) Validate() error {
	if s.SoarStrategyTasks != nil {
		for _, item := range s.SoarStrategyTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSoarStrategyTasksResponseBodySoarStrategyTasks struct {
	// The number of execution failures.
	//
	// example:
	//
	// 20
	FailedNum *int32 `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	// The timestamp when the policy task was created. Unit: milliseconds.
	//
	// example:
	//
	// 1708481235000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the policy task was complete. Unit: milliseconds.
	//
	// example:
	//
	// 1586739841000
	GmtFinish *int64 `json:"GmtFinish,omitempty" xml:"GmtFinish,omitempty"`
	// The timestamp when the policy task was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1669869436000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the policy task.
	//
	// example:
	//
	// 5374
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy task.
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
	// The status of the policy task. Valid values:
	//
	// 	- \\-1: waiting
	//
	// 	- 0: starting
	//
	// 	- 1: running
	//
	// 	- 2: finished
	//
	// 	- 3: schedule
	//
	// 	- 4: pause
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 8000
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The number of successful executions.
	//
	// example:
	//
	// 20
	SuccessNum *int32 `json:"SuccessNum,omitempty" xml:"SuccessNum,omitempty"`
	// The total number of executions.
	//
	// example:
	//
	// 20
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GetFailedNum() *int32 {
	return s.FailedNum
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GetGmtFinish() *int64 {
	return s.GmtFinish
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GetId() *int64 {
	return s.Id
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GetName() *string {
	return s.Name
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GetRunMode() *string {
	return s.RunMode
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GetSuccessNum() *int32 {
	return s.SuccessNum
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) SetFailedNum(v int32) *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	s.FailedNum = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) SetGmtCreate(v int64) *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) SetGmtFinish(v int64) *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	s.GmtFinish = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) SetGmtModified(v int64) *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	s.GmtModified = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) SetId(v int64) *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	s.Id = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) SetName(v string) *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	s.Name = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) SetRunMode(v string) *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	s.RunMode = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) SetStatus(v string) *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	s.Status = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) SetStrategyId(v int64) *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	s.StrategyId = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) SetSuccessNum(v int32) *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	s.SuccessNum = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) SetTotalNum(v int32) *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks {
	s.TotalNum = &v
	return s
}

func (s *DescribeSoarStrategyTasksResponseBodySoarStrategyTasks) Validate() error {
	return dara.Validate(s)
}

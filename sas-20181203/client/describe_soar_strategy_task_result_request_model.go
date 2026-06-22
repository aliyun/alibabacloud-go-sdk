// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *DescribeSoarStrategyTaskResultRequest
	GetCondition() *string
	SetCurrentPage(v int32) *DescribeSoarStrategyTaskResultRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSoarStrategyTaskResultRequest
	GetPageSize() *int32
	SetStrategyTaskId(v int64) *DescribeSoarStrategyTaskResultRequest
	GetStrategyTaskId() *int64
}

type DescribeSoarStrategyTaskResultRequest struct {
	// The condition parameter for task scheduling.
	//
	// example:
	//
	// {"status":1}
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The page number of the current page for a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries per page for a paged query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the policy task.
	//
	// >Call the [DescribeSoarStrategyTasks](~~DescribeSoarStrategyTasks~~) operation to obtain this parameter.
	//
	// example:
	//
	// 100
	StrategyTaskId *int64 `json:"StrategyTaskId,omitempty" xml:"StrategyTaskId,omitempty"`
}

func (s DescribeSoarStrategyTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTaskResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTaskResultRequest) GetCondition() *string {
	return s.Condition
}

func (s *DescribeSoarStrategyTaskResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSoarStrategyTaskResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSoarStrategyTaskResultRequest) GetStrategyTaskId() *int64 {
	return s.StrategyTaskId
}

func (s *DescribeSoarStrategyTaskResultRequest) SetCondition(v string) *DescribeSoarStrategyTaskResultRequest {
	s.Condition = &v
	return s
}

func (s *DescribeSoarStrategyTaskResultRequest) SetCurrentPage(v int32) *DescribeSoarStrategyTaskResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSoarStrategyTaskResultRequest) SetPageSize(v int32) *DescribeSoarStrategyTaskResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarStrategyTaskResultRequest) SetStrategyTaskId(v int64) *DescribeSoarStrategyTaskResultRequest {
	s.StrategyTaskId = &v
	return s
}

func (s *DescribeSoarStrategyTaskResultRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsUnfreezeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *UpdateOssCheckResultsUnfreezeRequest
	GetEndDate() *string
	SetFreezeItems(v string) *UpdateOssCheckResultsUnfreezeRequest
	GetFreezeItems() *string
	SetRegionId(v string) *UpdateOssCheckResultsUnfreezeRequest
	GetRegionId() *string
	SetStartDate(v string) *UpdateOssCheckResultsUnfreezeRequest
	GetStartDate() *string
	SetTaskId(v string) *UpdateOssCheckResultsUnfreezeRequest
	GetTaskId() *string
}

type UpdateOssCheckResultsUnfreezeRequest struct {
	// The end time. Format: YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2023-10-21 16:08:38
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The frozen items.
	//
	// example:
	//
	// []
	FreezeItems *string `json:"FreezeItems,omitempty" xml:"FreezeItems,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Format: YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2023-08-21 16:08:38
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The task ID.
	//
	// example:
	//
	// P_B6YRVD
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateOssCheckResultsUnfreezeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsUnfreezeRequest) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsUnfreezeRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *UpdateOssCheckResultsUnfreezeRequest) GetFreezeItems() *string {
	return s.FreezeItems
}

func (s *UpdateOssCheckResultsUnfreezeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateOssCheckResultsUnfreezeRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *UpdateOssCheckResultsUnfreezeRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateOssCheckResultsUnfreezeRequest) SetEndDate(v string) *UpdateOssCheckResultsUnfreezeRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeRequest) SetFreezeItems(v string) *UpdateOssCheckResultsUnfreezeRequest {
	s.FreezeItems = &v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeRequest) SetRegionId(v string) *UpdateOssCheckResultsUnfreezeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeRequest) SetStartDate(v string) *UpdateOssCheckResultsUnfreezeRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeRequest) SetTaskId(v string) *UpdateOssCheckResultsUnfreezeRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateOssCheckResultsUnfreezeRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssCheckResultsFreezeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *UpdateOssCheckResultsFreezeRequest
	GetEndDate() *string
	SetFreezeItems(v string) *UpdateOssCheckResultsFreezeRequest
	GetFreezeItems() *string
	SetFreezeRestorePath(v string) *UpdateOssCheckResultsFreezeRequest
	GetFreezeRestorePath() *string
	SetFreezeType(v string) *UpdateOssCheckResultsFreezeRequest
	GetFreezeType() *string
	SetRegionId(v string) *UpdateOssCheckResultsFreezeRequest
	GetRegionId() *string
	SetStartDate(v string) *UpdateOssCheckResultsFreezeRequest
	GetStartDate() *string
	SetTaskId(v string) *UpdateOssCheckResultsFreezeRequest
	GetTaskId() *string
}

type UpdateOssCheckResultsFreezeRequest struct {
	// example:
	//
	// 2023-10-21 16:08:38
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// []
	FreezeItems       *string `json:"FreezeItems,omitempty" xml:"FreezeItems,omitempty"`
	FreezeRestorePath *string `json:"FreezeRestorePath,omitempty" xml:"FreezeRestorePath,omitempty"`
	// example:
	//
	// ACL
	FreezeType *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2023-08-21 16:08:38
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// P_15BU42
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateOssCheckResultsFreezeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssCheckResultsFreezeRequest) GoString() string {
	return s.String()
}

func (s *UpdateOssCheckResultsFreezeRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *UpdateOssCheckResultsFreezeRequest) GetFreezeItems() *string {
	return s.FreezeItems
}

func (s *UpdateOssCheckResultsFreezeRequest) GetFreezeRestorePath() *string {
	return s.FreezeRestorePath
}

func (s *UpdateOssCheckResultsFreezeRequest) GetFreezeType() *string {
	return s.FreezeType
}

func (s *UpdateOssCheckResultsFreezeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateOssCheckResultsFreezeRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *UpdateOssCheckResultsFreezeRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateOssCheckResultsFreezeRequest) SetEndDate(v string) *UpdateOssCheckResultsFreezeRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeRequest) SetFreezeItems(v string) *UpdateOssCheckResultsFreezeRequest {
	s.FreezeItems = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeRequest) SetFreezeRestorePath(v string) *UpdateOssCheckResultsFreezeRequest {
	s.FreezeRestorePath = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeRequest) SetFreezeType(v string) *UpdateOssCheckResultsFreezeRequest {
	s.FreezeType = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeRequest) SetRegionId(v string) *UpdateOssCheckResultsFreezeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeRequest) SetStartDate(v string) *UpdateOssCheckResultsFreezeRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeRequest) SetTaskId(v string) *UpdateOssCheckResultsFreezeRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateOssCheckResultsFreezeRequest) Validate() error {
	return dara.Validate(s)
}

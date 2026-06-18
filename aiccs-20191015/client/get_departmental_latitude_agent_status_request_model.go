// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDepartmentalLatitudeAgentStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *GetDepartmentalLatitudeAgentStatusRequest
	GetCurrentPage() *int64
	SetDepIds(v []*int64) *GetDepartmentalLatitudeAgentStatusRequest
	GetDepIds() []*int64
	SetEndDate(v int64) *GetDepartmentalLatitudeAgentStatusRequest
	GetEndDate() *int64
	SetExistDepartmentGrouping(v bool) *GetDepartmentalLatitudeAgentStatusRequest
	GetExistDepartmentGrouping() *bool
	SetInstanceId(v string) *GetDepartmentalLatitudeAgentStatusRequest
	GetInstanceId() *string
	SetPageSize(v int64) *GetDepartmentalLatitudeAgentStatusRequest
	GetPageSize() *int64
	SetStartDate(v int64) *GetDepartmentalLatitudeAgentStatusRequest
	GetStartDate() *int64
}

type GetDepartmentalLatitudeAgentStatusRequest struct {
	// Current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of department IDs.
	DepIds []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	// End UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1617761765000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Indicates whether to query by department grouping. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ExistDepartmentGrouping *bool `json:"ExistDepartmentGrouping,omitempty" xml:"ExistDepartmentGrouping,omitempty"`
	// AICCS instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page size. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start date UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1615083365000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetDepartmentalLatitudeAgentStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDepartmentalLatitudeAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) GetDepIds() []*int64 {
	return s.DepIds
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) GetExistDepartmentGrouping() *bool {
	return s.ExistDepartmentGrouping
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetCurrentPage(v int64) *GetDepartmentalLatitudeAgentStatusRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetDepIds(v []*int64) *GetDepartmentalLatitudeAgentStatusRequest {
	s.DepIds = v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetEndDate(v int64) *GetDepartmentalLatitudeAgentStatusRequest {
	s.EndDate = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetExistDepartmentGrouping(v bool) *GetDepartmentalLatitudeAgentStatusRequest {
	s.ExistDepartmentGrouping = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetInstanceId(v string) *GetDepartmentalLatitudeAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetPageSize(v int64) *GetDepartmentalLatitudeAgentStatusRequest {
	s.PageSize = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) SetStartDate(v int64) *GetDepartmentalLatitudeAgentStatusRequest {
	s.StartDate = &v
	return s
}

func (s *GetDepartmentalLatitudeAgentStatusRequest) Validate() error {
	return dara.Validate(s)
}

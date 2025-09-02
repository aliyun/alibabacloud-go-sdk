// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineTypes(v string) *ListBaselinesRequest
	GetBaselineTypes() *string
	SetEnable(v bool) *ListBaselinesRequest
	GetEnable() *bool
	SetOwner(v string) *ListBaselinesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListBaselinesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBaselinesRequest
	GetPageSize() *int32
	SetPriority(v string) *ListBaselinesRequest
	GetPriority() *string
	SetProjectId(v int64) *ListBaselinesRequest
	GetProjectId() *int64
	SetSearchText(v string) *ListBaselinesRequest
	GetSearchText() *string
}

type ListBaselinesRequest struct {
	// The type of the baseline. Valid values: DAILY and HOURLY. You can specify multiple types. Separate multiple types with commas (,).
	//
	// example:
	//
	// DAILY
	BaselineTypes *string `json:"BaselineTypes,omitempty" xml:"BaselineTypes,omitempty"`
	// Specifies whether to enable the baseline. Valid values: true and false.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The owner.
	//
	// example:
	//
	// 3726346****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Pages start from page 1. Default value: 1. Maximum value: 30.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The priority of the baseline. Valid values: {1,3,5,7,8}.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The keyword in the baseline name, which is used to search for the baseline.
	//
	// example:
	//
	// baselineName
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
}

func (s ListBaselinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBaselinesRequest) GoString() string {
	return s.String()
}

func (s *ListBaselinesRequest) GetBaselineTypes() *string {
	return s.BaselineTypes
}

func (s *ListBaselinesRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ListBaselinesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListBaselinesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBaselinesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBaselinesRequest) GetPriority() *string {
	return s.Priority
}

func (s *ListBaselinesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBaselinesRequest) GetSearchText() *string {
	return s.SearchText
}

func (s *ListBaselinesRequest) SetBaselineTypes(v string) *ListBaselinesRequest {
	s.BaselineTypes = &v
	return s
}

func (s *ListBaselinesRequest) SetEnable(v bool) *ListBaselinesRequest {
	s.Enable = &v
	return s
}

func (s *ListBaselinesRequest) SetOwner(v string) *ListBaselinesRequest {
	s.Owner = &v
	return s
}

func (s *ListBaselinesRequest) SetPageNumber(v int32) *ListBaselinesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBaselinesRequest) SetPageSize(v int32) *ListBaselinesRequest {
	s.PageSize = &v
	return s
}

func (s *ListBaselinesRequest) SetPriority(v string) *ListBaselinesRequest {
	s.Priority = &v
	return s
}

func (s *ListBaselinesRequest) SetProjectId(v int64) *ListBaselinesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListBaselinesRequest) SetSearchText(v string) *ListBaselinesRequest {
	s.SearchText = &v
	return s
}

func (s *ListBaselinesRequest) Validate() error {
	return dara.Validate(s)
}

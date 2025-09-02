// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselineConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineTypes(v string) *ListBaselineConfigsRequest
	GetBaselineTypes() *string
	SetOwner(v string) *ListBaselineConfigsRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListBaselineConfigsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBaselineConfigsRequest
	GetPageSize() *int32
	SetPriority(v string) *ListBaselineConfigsRequest
	GetPriority() *string
	SetProjectId(v int64) *ListBaselineConfigsRequest
	GetProjectId() *int64
	SetSearchText(v string) *ListBaselineConfigsRequest
	GetSearchText() *string
	SetUseflag(v bool) *ListBaselineConfigsRequest
	GetUseflag() *bool
}

type ListBaselineConfigsRequest struct {
	// The type of the baseline. Valid values: DAILY and HOURLY. Separate multiple baseline types with commas (,).
	//
	// example:
	//
	// DAILY,HOURLY
	BaselineTypes *string `json:"BaselineTypes,omitempty" xml:"BaselineTypes,omitempty"`
	// The ID of the Alibaba Cloud account used by the baseline owner.
	//
	// example:
	//
	// 95279527****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Valid values: 1 to 30. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The priority of the baseline. Valid values: {1,3,5,7,8}. Separate multiple priorities with commas (,).
	//
	// example:
	//
	// 1,3,5,7,8
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The workspace ID. You can call the ListProjects operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The keyword in the baseline name, which is used to search for the baseline.
	//
	// example:
	//
	// Baseline name search keywords
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	// Specifies whether to enable the baseline. Valid values: true and false.
	//
	// example:
	//
	// true
	Useflag *bool `json:"Useflag,omitempty" xml:"Useflag,omitempty"`
}

func (s ListBaselineConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBaselineConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListBaselineConfigsRequest) GetBaselineTypes() *string {
	return s.BaselineTypes
}

func (s *ListBaselineConfigsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListBaselineConfigsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBaselineConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBaselineConfigsRequest) GetPriority() *string {
	return s.Priority
}

func (s *ListBaselineConfigsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBaselineConfigsRequest) GetSearchText() *string {
	return s.SearchText
}

func (s *ListBaselineConfigsRequest) GetUseflag() *bool {
	return s.Useflag
}

func (s *ListBaselineConfigsRequest) SetBaselineTypes(v string) *ListBaselineConfigsRequest {
	s.BaselineTypes = &v
	return s
}

func (s *ListBaselineConfigsRequest) SetOwner(v string) *ListBaselineConfigsRequest {
	s.Owner = &v
	return s
}

func (s *ListBaselineConfigsRequest) SetPageNumber(v int32) *ListBaselineConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBaselineConfigsRequest) SetPageSize(v int32) *ListBaselineConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListBaselineConfigsRequest) SetPriority(v string) *ListBaselineConfigsRequest {
	s.Priority = &v
	return s
}

func (s *ListBaselineConfigsRequest) SetProjectId(v int64) *ListBaselineConfigsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListBaselineConfigsRequest) SetSearchText(v string) *ListBaselineConfigsRequest {
	s.SearchText = &v
	return s
}

func (s *ListBaselineConfigsRequest) SetUseflag(v bool) *ListBaselineConfigsRequest {
	s.Useflag = &v
	return s
}

func (s *ListBaselineConfigsRequest) Validate() error {
	return dara.Validate(s)
}

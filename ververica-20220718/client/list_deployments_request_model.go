// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *ListDeploymentsRequest
	GetCreator() *string
	SetExecutionMode(v string) *ListDeploymentsRequest
	GetExecutionMode() *string
	SetLabelKey(v string) *ListDeploymentsRequest
	GetLabelKey() *string
	SetLabelValueArray(v string) *ListDeploymentsRequest
	GetLabelValueArray() *string
	SetModifier(v string) *ListDeploymentsRequest
	GetModifier() *string
	SetName(v string) *ListDeploymentsRequest
	GetName() *string
	SetPageIndex(v int32) *ListDeploymentsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDeploymentsRequest
	GetPageSize() *int32
	SetSortName(v string) *ListDeploymentsRequest
	GetSortName() *string
	SetStatus(v string) *ListDeploymentsRequest
	GetStatus() *string
}

type ListDeploymentsRequest struct {
	// Creator UID.
	//
	// example:
	//
	// 183899668*******
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// Deployment execution mode.
	//
	// example:
	//
	// STREAMING
	ExecutionMode *string `json:"executionMode,omitempty" xml:"executionMode,omitempty"`
	// Label name.
	//
	// example:
	//
	// key
	LabelKey *string `json:"labelKey,omitempty" xml:"labelKey,omitempty"`
	// Tag values. Multiple values are separated by semicolons.
	//
	// example:
	//
	// value1,value2
	LabelValueArray *string `json:"labelValueArray,omitempty" xml:"labelValueArray,omitempty"`
	// Modifier UID.
	//
	// example:
	//
	// 183899668*******
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// Deployment name.
	//
	// example:
	//
	// vvp_ds_0522
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Pagination parameter. Page index indicating the requested page number. Minimum value is 1. Default value is 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// Pagination parameter. Number of elements on the requested page. Maximum value is 100, minimum value is 1, and default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Sorting method. Supports returning data in descending order by creation time or updated time.
	//
	// example:
	//
	// gmt_create
	SortName *string `json:"sortName,omitempty" xml:"sortName,omitempty"`
	// The status of the latest job for the deployment.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListDeploymentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentsRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListDeploymentsRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *ListDeploymentsRequest) GetLabelKey() *string {
	return s.LabelKey
}

func (s *ListDeploymentsRequest) GetLabelValueArray() *string {
	return s.LabelValueArray
}

func (s *ListDeploymentsRequest) GetModifier() *string {
	return s.Modifier
}

func (s *ListDeploymentsRequest) GetName() *string {
	return s.Name
}

func (s *ListDeploymentsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDeploymentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentsRequest) GetSortName() *string {
	return s.SortName
}

func (s *ListDeploymentsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDeploymentsRequest) SetCreator(v string) *ListDeploymentsRequest {
	s.Creator = &v
	return s
}

func (s *ListDeploymentsRequest) SetExecutionMode(v string) *ListDeploymentsRequest {
	s.ExecutionMode = &v
	return s
}

func (s *ListDeploymentsRequest) SetLabelKey(v string) *ListDeploymentsRequest {
	s.LabelKey = &v
	return s
}

func (s *ListDeploymentsRequest) SetLabelValueArray(v string) *ListDeploymentsRequest {
	s.LabelValueArray = &v
	return s
}

func (s *ListDeploymentsRequest) SetModifier(v string) *ListDeploymentsRequest {
	s.Modifier = &v
	return s
}

func (s *ListDeploymentsRequest) SetName(v string) *ListDeploymentsRequest {
	s.Name = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageIndex(v int32) *ListDeploymentsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageSize(v int32) *ListDeploymentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsRequest) SetSortName(v string) *ListDeploymentsRequest {
	s.SortName = &v
	return s
}

func (s *ListDeploymentsRequest) SetStatus(v string) *ListDeploymentsRequest {
	s.Status = &v
	return s
}

func (s *ListDeploymentsRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v string) *ListExperimentShrinkRequest
	GetLabels() *string
	SetMaxResults(v int64) *ListExperimentShrinkRequest
	GetMaxResults() *int64
	SetName(v string) *ListExperimentShrinkRequest
	GetName() *string
	SetOptionsShrink(v string) *ListExperimentShrinkRequest
	GetOptionsShrink() *string
	SetOrder(v string) *ListExperimentShrinkRequest
	GetOrder() *string
	SetOrderBy(v string) *ListExperimentShrinkRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *ListExperimentShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExperimentShrinkRequest
	GetPageSize() *int32
	SetPageToken(v int64) *ListExperimentShrinkRequest
	GetPageToken() *int64
	SetSortBy(v string) *ListExperimentShrinkRequest
	GetSortBy() *string
	SetVerbose(v bool) *ListExperimentShrinkRequest
	GetVerbose() *bool
	SetWorkspaceId(v string) *ListExperimentShrinkRequest
	GetWorkspaceId() *string
}

type ListExperimentShrinkRequest struct {
	// The tag filter conditions. Multiple conditions are separated by commas (,). The format of a single condition filter is `key=value`.
	//
	// example:
	//
	// is_evaluation:true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum number of entries in the request. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The experiment name.
	//
	// example:
	//
	// exp-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The optional parameters.
	OptionsShrink *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The order of specific fields of results in a paged query (ascending or descending).
	//
	// 	- ASC: ascending order
	//
	// 	- DESC: descending order. This is the default value.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The strings used for sorting. The following fields can be used for sorting: GmtCreateTime, Name, GmtModifiedTime, and ExperimentId. The sorting order can be ASC (default) and DESC.
	//
	// example:
	//
	// GmtCreateTime DESC,Name ASC
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. The value starts from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pagination token, which starts from 0. Default value: 0.
	//
	// example:
	//
	// 0
	PageToken *int64 `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	// The field used for sorting. The GmtCreateTime field is used.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Specifies whether to obtain the LatestRun value that is related to the experiment.
	//
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// The ID of the workspace to which the experiment belongs. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// >  If you do not specify a workspace ID, the system returns the experiments in the default workspace.
	//
	// example:
	//
	// 151739
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListExperimentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentShrinkRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListExperimentShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListExperimentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListExperimentShrinkRequest) GetOptionsShrink() *string {
	return s.OptionsShrink
}

func (s *ListExperimentShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListExperimentShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListExperimentShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExperimentShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExperimentShrinkRequest) GetPageToken() *int64 {
	return s.PageToken
}

func (s *ListExperimentShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListExperimentShrinkRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListExperimentShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListExperimentShrinkRequest) SetLabels(v string) *ListExperimentShrinkRequest {
	s.Labels = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetMaxResults(v int64) *ListExperimentShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetName(v string) *ListExperimentShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetOptionsShrink(v string) *ListExperimentShrinkRequest {
	s.OptionsShrink = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetOrder(v string) *ListExperimentShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetOrderBy(v string) *ListExperimentShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetPageNumber(v int32) *ListExperimentShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetPageSize(v int32) *ListExperimentShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetPageToken(v int64) *ListExperimentShrinkRequest {
	s.PageToken = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetSortBy(v string) *ListExperimentShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetVerbose(v bool) *ListExperimentShrinkRequest {
	s.Verbose = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetWorkspaceId(v string) *ListExperimentShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListExperimentShrinkRequest) Validate() error {
	return dara.Validate(s)
}

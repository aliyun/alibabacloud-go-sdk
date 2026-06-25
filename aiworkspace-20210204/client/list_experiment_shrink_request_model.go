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
	// The filter conditions for labels. Separate multiple conditions with commas (,). A single filter condition must be in the `Key=Value` format.
	//
	// example:
	//
	// is_evaluation:true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum number of results to return. The default is 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the experiment.
	//
	// example:
	//
	// exp-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Optional parameters.
	OptionsShrink *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The order in which to sort the results of a paged query. Valid values:
	//
	// - ASC: ascending order.
	//
	// - DESC (default): descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// A list of sorting methods as strings. You can sort by the following fields: GmtCreateTime, Name, GmtModifiedTime, or ExperimentId. The sorting methods are DESC and ASC. The default is ASC.
	//
	// example:
	//
	// GmtCreateTime DESC,Name ASC
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from 1.
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
	// The paging token. It starts from 0. The default is 0.
	//
	// example:
	//
	// 0
	PageToken *int64 `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	// The field to use for sorting in a paged query. Currently, only the GmtCreateTime field is supported for sorting.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Specifies whether to retrieve the LatestRun information related to the experiment.
	//
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// The ID of the workspace where the experiment resides. For more information about how to obtain a workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// > If you do not specify a workspace ID, the system returns the list of experiments in the default workspace.
	//
	// example:
	//
	// 1517**
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

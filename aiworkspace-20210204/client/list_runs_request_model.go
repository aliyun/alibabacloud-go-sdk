// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExperimentId(v string) *ListRunsRequest
	GetExperimentId() *string
	SetGmtCreateTime(v string) *ListRunsRequest
	GetGmtCreateTime() *string
	SetLabels(v string) *ListRunsRequest
	GetLabels() *string
	SetMaxResults(v int64) *ListRunsRequest
	GetMaxResults() *int64
	SetName(v string) *ListRunsRequest
	GetName() *string
	SetOrder(v string) *ListRunsRequest
	GetOrder() *string
	SetOrderBy(v string) *ListRunsRequest
	GetOrderBy() *string
	SetPageNumber(v int64) *ListRunsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRunsRequest
	GetPageSize() *int64
	SetPageToken(v int64) *ListRunsRequest
	GetPageToken() *int64
	SetSortBy(v string) *ListRunsRequest
	GetSortBy() *string
	SetSourceId(v string) *ListRunsRequest
	GetSourceId() *string
	SetSourceType(v string) *ListRunsRequest
	GetSourceType() *string
	SetVerbose(v bool) *ListRunsRequest
	GetVerbose() *bool
	SetWorkspaceId(v string) *ListRunsRequest
	GetWorkspaceId() *string
}

type ListRunsRequest struct {
	// The ID of the experiment to which the run belongs.
	//
	// example:
	//
	// exp-1zpfthdx******
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The labels of the run for an exact match. The following formats are supported:
	//
	// - Single-label query: "is_evaluation:true"
	//
	// - Multi-label query: "is_evaluation:true,LLM_evaluation:true". This method is not recommended for common scenarios because it may degrade performance. Use commas (,) to separate multiple labels. The system matches all specified key-value pairs.
	//
	// example:
	//
	// is_evaluation:true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum number of results to return. The default value is 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the run.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order for the paged query. Use this parameter with SortBy.
	//
	// - ASC: ascending order.
	//
	// - DESC (default): descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The fields to sort by and the sort order. You can sort by GmtCreateTime and Name. Valid sort orders are DESC and ASC. The default is ASC. To sort by multiple fields, separate them with a comma (,).
	//
	// example:
	//
	// GmtCreateTime DESC,Name ASC
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. The value must be greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records to display on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The paging token. The value starts from 0. The default value is 0.
	//
	// example:
	//
	// 0
	PageToken *int64 `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	// The field to use for sorting. Valid values:
	//
	// - Name: the name of the run.
	//
	// - GmtCreateTime (default): the time when the run was created.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The ID of the PAI workload associated with the run.
	//
	// example:
	//
	// job-rbvg5wzlj****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The type of the PAI workload associated with the run.
	//
	// example:
	//
	// TrainingService
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Specifies whether to display details, including Metrics, Params, and Labels. Valid values:
	//
	// - true: displays details.
	//
	// - false (default): does not display details.
	//
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// The ID of the workspace where the experiment resides. For more information about how to obtain a workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// > If you do not specify a workspace ID, the system returns the list of runs in the default workspace.
	//
	// example:
	//
	// 228**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRunsRequest) GoString() string {
	return s.String()
}

func (s *ListRunsRequest) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *ListRunsRequest) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListRunsRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListRunsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListRunsRequest) GetName() *string {
	return s.Name
}

func (s *ListRunsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListRunsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListRunsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRunsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRunsRequest) GetPageToken() *int64 {
	return s.PageToken
}

func (s *ListRunsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListRunsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ListRunsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListRunsRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListRunsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListRunsRequest) SetExperimentId(v string) *ListRunsRequest {
	s.ExperimentId = &v
	return s
}

func (s *ListRunsRequest) SetGmtCreateTime(v string) *ListRunsRequest {
	s.GmtCreateTime = &v
	return s
}

func (s *ListRunsRequest) SetLabels(v string) *ListRunsRequest {
	s.Labels = &v
	return s
}

func (s *ListRunsRequest) SetMaxResults(v int64) *ListRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRunsRequest) SetName(v string) *ListRunsRequest {
	s.Name = &v
	return s
}

func (s *ListRunsRequest) SetOrder(v string) *ListRunsRequest {
	s.Order = &v
	return s
}

func (s *ListRunsRequest) SetOrderBy(v string) *ListRunsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListRunsRequest) SetPageNumber(v int64) *ListRunsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRunsRequest) SetPageSize(v int64) *ListRunsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRunsRequest) SetPageToken(v int64) *ListRunsRequest {
	s.PageToken = &v
	return s
}

func (s *ListRunsRequest) SetSortBy(v string) *ListRunsRequest {
	s.SortBy = &v
	return s
}

func (s *ListRunsRequest) SetSourceId(v string) *ListRunsRequest {
	s.SourceId = &v
	return s
}

func (s *ListRunsRequest) SetSourceType(v string) *ListRunsRequest {
	s.SourceType = &v
	return s
}

func (s *ListRunsRequest) SetVerbose(v bool) *ListRunsRequest {
	s.Verbose = &v
	return s
}

func (s *ListRunsRequest) SetWorkspaceId(v string) *ListRunsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListRunsRequest) Validate() error {
	return dara.Validate(s)
}

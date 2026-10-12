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
	// The time when the instance was created. The time follows the ISO 8601 standard in UTC. Format: YYYY-MM-DDTHH:mm:ss.sssZ.
	//
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The labels of the task for exact match. Supported formats:
	//
	// - Single-label query: "is_evaluation:true"
	//
	// - Multi-label scenario (not recommended except in special cases due to potential performance issues): "is_evaluation:true,LLM_evaluation:true" separated by commas (,), indicating that multiple label key-value pairs must all match.
	//
	// example:
	//
	// is_evaluation:true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum number of results to return. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The run name.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sorting order for the specified field in a paging query. Used together with SortBy. Valid values:
	//
	// - ASC: Ascending order.
	//
	// - DESC (default): Descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// A comma-separated list of sorting strings. Supported fields: GmtCreateTime and Name. Supported sorting orders: DESC and ASC. Default: ASC. Separate multiple fields with commas (,).
	//
	// example:
	//
	// GmtCreateTime DESC,Name ASC
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Valid values: greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The pagination token. Starts from 0. Default value: 0.
	//
	// example:
	//
	// 0
	PageToken *int64 `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	// The field used for sorting. Valid values:
	//
	// - Name: The run name.
	//
	// - GmtCreateTime (default): The run creation time.
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
	// Specifies whether to display detailed information, including Metrics, Params, and Labels. Valid values:
	//
	// - true: Detailed information is displayed.
	//
	// - false (default): Detailed information is not displayed.
	//
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// The ID of the workspace where the experiment resides. For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// > If you do not specify a workspace ID, the system returns the run list of the default workspace.
	//
	// >
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

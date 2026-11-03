// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v string) *ListExperimentRequest
	GetLabels() *string
	SetMaxResults(v int64) *ListExperimentRequest
	GetMaxResults() *int64
	SetName(v string) *ListExperimentRequest
	GetName() *string
	SetOptions(v *ListExperimentRequestOptions) *ListExperimentRequest
	GetOptions() *ListExperimentRequestOptions
	SetOrder(v string) *ListExperimentRequest
	GetOrder() *string
	SetOrderBy(v string) *ListExperimentRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *ListExperimentRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExperimentRequest
	GetPageSize() *int32
	SetPageToken(v int64) *ListExperimentRequest
	GetPageToken() *int64
	SetSortBy(v string) *ListExperimentRequest
	GetSortBy() *string
	SetVerbose(v bool) *ListExperimentRequest
	GetVerbose() *bool
	SetWorkspaceId(v string) *ListExperimentRequest
	GetWorkspaceId() *string
}

type ListExperimentRequest struct {
	// The label filter conditions. Separate multiple conditions with commas (,).
	//
	// The format of a single filter condition is `Key=Value`.
	//
	// example:
	//
	// is_evaluation:true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum number of results to return for the request. Default value: 10.
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
	// The optional parameters.
	Options *ListExperimentRequestOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// The sort order of the specified field in a paged query. Valid values:
	//
	// - ASC: ascending order.
	//
	// - DESC: descending order.
	//
	// Default value: DESC.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The list of sort order strings. Supported fields for sorting include GmtCreateTime, Name, GmtModifiedTime, and ExperimentId. The sort order can be DESC or ASC. Default value: ASC.
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
	// The pagination token. Pages start from 0. Default value: 0.
	//
	// example:
	//
	// 0
	PageToken *int64 `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	// The field by which to sort the results of a paged query. Currently, only the GmtCreateTime field is supported for sorting.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Specifies whether to retrieve the LatestRun information associated with the experiment.
	//
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// The ID of the workspace to which the experiment belongs. For more information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// > If you do not specify the workspace ID, the system returns the list of experiments in the default workspace.
	//
	// >
	//
	// example:
	//
	// 1517**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListExperimentRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListExperimentRequest) GetName() *string {
	return s.Name
}

func (s *ListExperimentRequest) GetOptions() *ListExperimentRequestOptions {
	return s.Options
}

func (s *ListExperimentRequest) GetOrder() *string {
	return s.Order
}

func (s *ListExperimentRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListExperimentRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExperimentRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExperimentRequest) GetPageToken() *int64 {
	return s.PageToken
}

func (s *ListExperimentRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListExperimentRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListExperimentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListExperimentRequest) SetLabels(v string) *ListExperimentRequest {
	s.Labels = &v
	return s
}

func (s *ListExperimentRequest) SetMaxResults(v int64) *ListExperimentRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExperimentRequest) SetName(v string) *ListExperimentRequest {
	s.Name = &v
	return s
}

func (s *ListExperimentRequest) SetOptions(v *ListExperimentRequestOptions) *ListExperimentRequest {
	s.Options = v
	return s
}

func (s *ListExperimentRequest) SetOrder(v string) *ListExperimentRequest {
	s.Order = &v
	return s
}

func (s *ListExperimentRequest) SetOrderBy(v string) *ListExperimentRequest {
	s.OrderBy = &v
	return s
}

func (s *ListExperimentRequest) SetPageNumber(v int32) *ListExperimentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExperimentRequest) SetPageSize(v int32) *ListExperimentRequest {
	s.PageSize = &v
	return s
}

func (s *ListExperimentRequest) SetPageToken(v int64) *ListExperimentRequest {
	s.PageToken = &v
	return s
}

func (s *ListExperimentRequest) SetSortBy(v string) *ListExperimentRequest {
	s.SortBy = &v
	return s
}

func (s *ListExperimentRequest) SetVerbose(v bool) *ListExperimentRequest {
	s.Verbose = &v
	return s
}

func (s *ListExperimentRequest) SetWorkspaceId(v string) *ListExperimentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListExperimentRequest) Validate() error {
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExperimentRequestOptions struct {
	// Specifies whether the name is an exact match. Valid values: true and false.
	//
	// example:
	//
	// true
	MatchNameExactly *string `json:"match_name_exactly,omitempty" xml:"match_name_exactly,omitempty"`
}

func (s ListExperimentRequestOptions) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentRequestOptions) GoString() string {
	return s.String()
}

func (s *ListExperimentRequestOptions) GetMatchNameExactly() *string {
	return s.MatchNameExactly
}

func (s *ListExperimentRequestOptions) SetMatchNameExactly(v string) *ListExperimentRequestOptions {
	s.MatchNameExactly = &v
	return s
}

func (s *ListExperimentRequestOptions) Validate() error {
	return dara.Validate(s)
}

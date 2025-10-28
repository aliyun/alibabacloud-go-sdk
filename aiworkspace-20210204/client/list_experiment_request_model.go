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
	Options *ListExperimentRequestOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
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
	// Specifies whether to exactly match the experiment by name. Valid values: true and false.
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

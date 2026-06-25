// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetVersion(v string) *ListDatasetJobsRequest
	GetDatasetVersion() *string
	SetJobAction(v string) *ListDatasetJobsRequest
	GetJobAction() *string
	SetOrder(v string) *ListDatasetJobsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListDatasetJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetJobsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListDatasetJobsRequest
	GetSortBy() *string
	SetStatus(v string) *ListDatasetJobsRequest
	GetStatus() *string
	SetWithLogs(v bool) *ListDatasetJobsRequest
	GetWithLogs() *bool
	SetWorkspaceId(v string) *ListDatasetJobsRequest
	GetWorkspaceId() *string
}

type ListDatasetJobsRequest struct {
	// The name of the dataset version.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The job action.
	//
	// example:
	//
	// SemanticIndex
	JobAction *string `json:"JobAction,omitempty" xml:"JobAction,omitempty"`
	// The order in which to sort the results. This parameter is used with `SortBy`. Default: DESC.
	//
	// - ASC: ascending order.
	//
	// - DESC: descending order.
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages are 1-indexed. Default: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field by which to sort the results. By default, the results are sorted by `CreateTime` in descending order.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The job status.
	//
	// example:
	//
	// Running
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WithLogs *bool   `json:"WithLogs,omitempty" xml:"WithLogs,omitempty"`
	// The ID of the workspace. To obtain this ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasetJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetJobsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetJobsRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *ListDatasetJobsRequest) GetJobAction() *string {
	return s.JobAction
}

func (s *ListDatasetJobsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDatasetJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDatasetJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDatasetJobsRequest) GetWithLogs() *bool {
	return s.WithLogs
}

func (s *ListDatasetJobsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDatasetJobsRequest) SetDatasetVersion(v string) *ListDatasetJobsRequest {
	s.DatasetVersion = &v
	return s
}

func (s *ListDatasetJobsRequest) SetJobAction(v string) *ListDatasetJobsRequest {
	s.JobAction = &v
	return s
}

func (s *ListDatasetJobsRequest) SetOrder(v string) *ListDatasetJobsRequest {
	s.Order = &v
	return s
}

func (s *ListDatasetJobsRequest) SetPageNumber(v int32) *ListDatasetJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetJobsRequest) SetPageSize(v int32) *ListDatasetJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetJobsRequest) SetSortBy(v string) *ListDatasetJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasetJobsRequest) SetStatus(v string) *ListDatasetJobsRequest {
	s.Status = &v
	return s
}

func (s *ListDatasetJobsRequest) SetWithLogs(v bool) *ListDatasetJobsRequest {
	s.WithLogs = &v
	return s
}

func (s *ListDatasetJobsRequest) SetWorkspaceId(v string) *ListDatasetJobsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasetJobsRequest) Validate() error {
	return dara.Validate(s)
}

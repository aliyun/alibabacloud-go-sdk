// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetFileMetasStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregateBy(v string) *GetDatasetFileMetasStatisticsRequest
	GetAggregateBy() *string
	SetDatasetVersion(v string) *GetDatasetFileMetasStatisticsRequest
	GetDatasetVersion() *string
	SetMaxResults(v int32) *GetDatasetFileMetasStatisticsRequest
	GetMaxResults() *int32
	SetWorkspaceId(v string) *GetDatasetFileMetasStatisticsRequest
	GetWorkspaceId() *string
}

type GetDatasetFileMetasStatisticsRequest struct {
	// Aggregates statistics based on the specified metadata field. The value is not case-sensitive. If not specified, the total number of dataset file metadata will be returned, instead of aggregation lists. Valid values:
	//
	// 	- filedir: the directory path of the file
	//
	// 	- file_type: the file type
	//
	// 	- tags.user: user-defined tag
	//
	// 	- tags.user-delete-ai-tags: algorithm tags deleted by the user
	//
	// 	- tags.ai: algorithm tags (aggregated by all tagging tasks)
	//
	// 	- tags.all: algorithm tags and user-defined tags (excluding alogorithm tags deleted by the user)
	//
	// example:
	//
	// filedir
	AggregateBy *string `json:"AggregateBy,omitempty" xml:"AggregateBy,omitempty"`
	// The dataset version.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The maximum number of results to be returned from a single query when the NextToken parameter is used in the query. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 145883
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetFileMetasStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetFileMetasStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetFileMetasStatisticsRequest) GetAggregateBy() *string {
	return s.AggregateBy
}

func (s *GetDatasetFileMetasStatisticsRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *GetDatasetFileMetasStatisticsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetDatasetFileMetasStatisticsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDatasetFileMetasStatisticsRequest) SetAggregateBy(v string) *GetDatasetFileMetasStatisticsRequest {
	s.AggregateBy = &v
	return s
}

func (s *GetDatasetFileMetasStatisticsRequest) SetDatasetVersion(v string) *GetDatasetFileMetasStatisticsRequest {
	s.DatasetVersion = &v
	return s
}

func (s *GetDatasetFileMetasStatisticsRequest) SetMaxResults(v int32) *GetDatasetFileMetasStatisticsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetDatasetFileMetasStatisticsRequest) SetWorkspaceId(v string) *GetDatasetFileMetasStatisticsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDatasetFileMetasStatisticsRequest) Validate() error {
	return dara.Validate(s)
}

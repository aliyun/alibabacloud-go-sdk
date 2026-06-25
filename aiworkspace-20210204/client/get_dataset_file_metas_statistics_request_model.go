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
	// The metadata field used for statistical aggregation. The value is not case-sensitive. If you do not specify this parameter, the total number of file metadata entries in the dataset is returned, and the aggregation list is not returned.
	//
	// Valid values:
	//
	// - filedir: The directory path of the file.
	//
	// - filetype: The file type.
	//
	// - tags.user: Custom user tags.
	//
	// - tags.user-delete-ai-tags: Algorithm tags deleted by the user.
	//
	// - tags.ai: Algorithm tags that are aggregated from all labeling tasks.
	//
	// - tags.all: A combination of algorithm tags and custom user tags, excluding any algorithm tags deleted by the user.
	//
	// example:
	//
	// filedir
	AggregateBy *string `json:"AggregateBy,omitempty" xml:"AggregateBy,omitempty"`
	// The name of the dataset version.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The maximum number of results to return for each query that uses the NextToken parameter. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The workspace ID. For more information about how to obtain a workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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

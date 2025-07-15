// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *ListDatasetsRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *ListDatasetsRequest
	GetDatasetName() *string
	SetDatasetType(v string) *ListDatasetsRequest
	GetDatasetType() *string
	SetEndTime(v string) *ListDatasetsRequest
	GetEndTime() *string
	SetIncludeConfig(v bool) *ListDatasetsRequest
	GetIncludeConfig() *bool
	SetPageNumber(v int32) *ListDatasetsRequest
	GetPageNumber() *int32
	SetPageSize(v string) *ListDatasetsRequest
	GetPageSize() *string
	SetSearchDatasetEnable(v int32) *ListDatasetsRequest
	GetSearchDatasetEnable() *int32
	SetStartTime(v string) *ListDatasetsRequest
	GetStartTime() *string
	SetWorkspaceId(v string) *ListDatasetsRequest
	GetWorkspaceId() *string
}

type ListDatasetsRequest struct {
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// businessDataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// CustomSemanticSearch
	DatasetType *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	// example:
	//
	// 创建时间-结束
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IncludeConfig *bool   `json:"IncludeConfig,omitempty" xml:"IncludeConfig,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3
	SearchDatasetEnable *int32 `json:"SearchDatasetEnable,omitempty" xml:"SearchDatasetEnable,omitempty"`
	// example:
	//
	// 创建时间-开始
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *ListDatasetsRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListDatasetsRequest) GetDatasetType() *string {
	return s.DatasetType
}

func (s *ListDatasetsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDatasetsRequest) GetIncludeConfig() *bool {
	return s.IncludeConfig
}

func (s *ListDatasetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDatasetsRequest) GetSearchDatasetEnable() *int32 {
	return s.SearchDatasetEnable
}

func (s *ListDatasetsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDatasetsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDatasetsRequest) SetDatasetId(v int64) *ListDatasetsRequest {
	s.DatasetId = &v
	return s
}

func (s *ListDatasetsRequest) SetDatasetName(v string) *ListDatasetsRequest {
	s.DatasetName = &v
	return s
}

func (s *ListDatasetsRequest) SetDatasetType(v string) *ListDatasetsRequest {
	s.DatasetType = &v
	return s
}

func (s *ListDatasetsRequest) SetEndTime(v string) *ListDatasetsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDatasetsRequest) SetIncludeConfig(v bool) *ListDatasetsRequest {
	s.IncludeConfig = &v
	return s
}

func (s *ListDatasetsRequest) SetPageNumber(v int32) *ListDatasetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetsRequest) SetPageSize(v string) *ListDatasetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetsRequest) SetSearchDatasetEnable(v int32) *ListDatasetsRequest {
	s.SearchDatasetEnable = &v
	return s
}

func (s *ListDatasetsRequest) SetStartTime(v string) *ListDatasetsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDatasetsRequest) SetWorkspaceId(v string) *ListDatasetsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasetsRequest) Validate() error {
	return dara.Validate(s)
}

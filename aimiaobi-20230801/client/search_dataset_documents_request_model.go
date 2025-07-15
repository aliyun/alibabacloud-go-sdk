// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDatasetDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v int64) *SearchDatasetDocumentsRequest
	GetDatasetId() *int64
	SetDatasetName(v string) *SearchDatasetDocumentsRequest
	GetDatasetName() *string
	SetExtend1(v string) *SearchDatasetDocumentsRequest
	GetExtend1() *string
	SetIncludeContent(v bool) *SearchDatasetDocumentsRequest
	GetIncludeContent() *bool
	SetPageSize(v string) *SearchDatasetDocumentsRequest
	GetPageSize() *string
	SetQuery(v string) *SearchDatasetDocumentsRequest
	GetQuery() *string
	SetWorkspaceId(v string) *SearchDatasetDocumentsRequest
	GetWorkspaceId() *string
}

type SearchDatasetDocumentsRequest struct {
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// 业务参数
	Extend1 *string `json:"Extend1,omitempty" xml:"Extend1,omitempty"`
	// example:
	//
	// false
	IncludeContent *bool `json:"IncludeContent,omitempty" xml:"IncludeContent,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 搜索内容
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SearchDatasetDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchDatasetDocumentsRequest) GoString() string {
	return s.String()
}

func (s *SearchDatasetDocumentsRequest) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *SearchDatasetDocumentsRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *SearchDatasetDocumentsRequest) GetExtend1() *string {
	return s.Extend1
}

func (s *SearchDatasetDocumentsRequest) GetIncludeContent() *bool {
	return s.IncludeContent
}

func (s *SearchDatasetDocumentsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *SearchDatasetDocumentsRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchDatasetDocumentsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SearchDatasetDocumentsRequest) SetDatasetId(v int64) *SearchDatasetDocumentsRequest {
	s.DatasetId = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetDatasetName(v string) *SearchDatasetDocumentsRequest {
	s.DatasetName = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetExtend1(v string) *SearchDatasetDocumentsRequest {
	s.Extend1 = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetIncludeContent(v bool) *SearchDatasetDocumentsRequest {
	s.IncludeContent = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetPageSize(v string) *SearchDatasetDocumentsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetQuery(v string) *SearchDatasetDocumentsRequest {
	s.Query = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) SetWorkspaceId(v string) *SearchDatasetDocumentsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SearchDatasetDocumentsRequest) Validate() error {
	return dara.Validate(s)
}

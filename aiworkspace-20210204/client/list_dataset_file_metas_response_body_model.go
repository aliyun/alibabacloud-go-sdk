// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetFileMetasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetFileMetas(v []*DatasetFileMeta) *ListDatasetFileMetasResponseBody
	GetDatasetFileMetas() []*DatasetFileMeta
	SetDatasetId(v string) *ListDatasetFileMetasResponseBody
	GetDatasetId() *string
	SetDatasetVersion(v string) *ListDatasetFileMetasResponseBody
	GetDatasetVersion() *string
	SetMaxResults(v int32) *ListDatasetFileMetasResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDatasetFileMetasResponseBody
	GetNextToken() *string
	SetPageSize(v int32) *ListDatasetFileMetasResponseBody
	GetPageSize() *int32
	SetTotalCount(v int32) *ListDatasetFileMetasResponseBody
	GetTotalCount() *int32
	SetWorkspaceId(v string) *ListDatasetFileMetasResponseBody
	GetWorkspaceId() *string
}

type ListDatasetFileMetasResponseBody struct {
	// The metadata records of the dataset files.
	DatasetFileMetas []*DatasetFileMeta `json:"DatasetFileMetas,omitempty" xml:"DatasetFileMetas,omitempty" type:"Repeated"`
	// The dataset ID.
	//
	// example:
	//
	// d-rbvg5*****jhc9ks92
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The dataset version.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. If the number of results exceeds the maximum number of entries allowed per page, a pagination token is returned. This token can be used as an input parameter to obtain the next page of results. If all results are obtained, no token is returned.
	//
	// example:
	//
	// 90******-f5c5-4cd4-927e-1f45e1cb8b62_1729644433000
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Deprecated
	//
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 123
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 105173
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasetFileMetasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetFileMetasResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetFileMetasResponseBody) GetDatasetFileMetas() []*DatasetFileMeta {
	return s.DatasetFileMetas
}

func (s *ListDatasetFileMetasResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *ListDatasetFileMetasResponseBody) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *ListDatasetFileMetasResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatasetFileMetasResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasetFileMetasResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetFileMetasResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDatasetFileMetasResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDatasetFileMetasResponseBody) SetDatasetFileMetas(v []*DatasetFileMeta) *ListDatasetFileMetasResponseBody {
	s.DatasetFileMetas = v
	return s
}

func (s *ListDatasetFileMetasResponseBody) SetDatasetId(v string) *ListDatasetFileMetasResponseBody {
	s.DatasetId = &v
	return s
}

func (s *ListDatasetFileMetasResponseBody) SetDatasetVersion(v string) *ListDatasetFileMetasResponseBody {
	s.DatasetVersion = &v
	return s
}

func (s *ListDatasetFileMetasResponseBody) SetMaxResults(v int32) *ListDatasetFileMetasResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDatasetFileMetasResponseBody) SetNextToken(v string) *ListDatasetFileMetasResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDatasetFileMetasResponseBody) SetPageSize(v int32) *ListDatasetFileMetasResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDatasetFileMetasResponseBody) SetTotalCount(v int32) *ListDatasetFileMetasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatasetFileMetasResponseBody) SetWorkspaceId(v string) *ListDatasetFileMetasResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasetFileMetasResponseBody) Validate() error {
	if s.DatasetFileMetas != nil {
		for _, item := range s.DatasetFileMetas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkLogAnalyzeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListSparkLogAnalyzeTasksRequest
	GetDBClusterId() *string
	SetPageNumber(v int64) *ListSparkLogAnalyzeTasksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSparkLogAnalyzeTasksRequest
	GetPageSize() *int64
}

type ListSparkLogAnalyzeTasksRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-9scxs****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSparkLogAnalyzeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSparkLogAnalyzeTasksRequest) GoString() string {
	return s.String()
}

func (s *ListSparkLogAnalyzeTasksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListSparkLogAnalyzeTasksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSparkLogAnalyzeTasksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSparkLogAnalyzeTasksRequest) SetDBClusterId(v string) *ListSparkLogAnalyzeTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksRequest) SetPageNumber(v int64) *ListSparkLogAnalyzeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksRequest) SetPageSize(v int64) *ListSparkLogAnalyzeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListSparkLogAnalyzeTasksRequest) Validate() error {
	return dara.Validate(s)
}

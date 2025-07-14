// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiDatasourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWord(v string) *ListApiDatasourceRequest
	GetKeyWord() *string
	SetPageNum(v int32) *ListApiDatasourceRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListApiDatasourceRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListApiDatasourceRequest
	GetWorkspaceId() *string
}

type ListApiDatasourceRequest struct {
	// The keyword of the API data source name.
	//
	// example:
	//
	// test
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// Current page number for API data source list:
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page in a paged query.
	//
	// 	- Default value: 10.
	//
	// 	- Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 89713491-cb4f-4579-b889-e82c35f1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListApiDatasourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiDatasourceRequest) GoString() string {
	return s.String()
}

func (s *ListApiDatasourceRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ListApiDatasourceRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListApiDatasourceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiDatasourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListApiDatasourceRequest) SetKeyWord(v string) *ListApiDatasourceRequest {
	s.KeyWord = &v
	return s
}

func (s *ListApiDatasourceRequest) SetPageNum(v int32) *ListApiDatasourceRequest {
	s.PageNum = &v
	return s
}

func (s *ListApiDatasourceRequest) SetPageSize(v int32) *ListApiDatasourceRequest {
	s.PageSize = &v
	return s
}

func (s *ListApiDatasourceRequest) SetWorkspaceId(v string) *ListApiDatasourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListApiDatasourceRequest) Validate() error {
	return dara.Validate(s)
}

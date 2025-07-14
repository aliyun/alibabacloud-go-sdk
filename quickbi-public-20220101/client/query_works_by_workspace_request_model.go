// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksByWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *QueryWorksByWorkspaceRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryWorksByWorkspaceRequest
	GetPageSize() *int32
	SetStatus(v int32) *QueryWorksByWorkspaceRequest
	GetStatus() *int32
	SetThirdPartAuthFlag(v int32) *QueryWorksByWorkspaceRequest
	GetThirdPartAuthFlag() *int32
	SetWorksType(v string) *QueryWorksByWorkspaceRequest
	GetWorksType() *string
	SetWorkspaceId(v string) *QueryWorksByWorkspaceRequest
	GetWorkspaceId() *string
}

type QueryWorksByWorkspaceRequest struct {
	// The page number of the returned page.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the work. Valid values:
	//
	// 	- 0: unpublished
	//
	// 	- 1: published
	//
	// 	- 2: modified but not published
	//
	// 	- 3: unpublished
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Third-party embedding status. Valid values:
	//
	// 	- 0: The embed service is not enabled.
	//
	// 	- 1: Embed is enabled.
	//
	// example:
	//
	// 0
	ThirdPartAuthFlag *int32 `json:"ThirdPartAuthFlag,omitempty" xml:"ThirdPartAuthFlag,omitempty"`
	// The type of the work. Valid values:
	//
	// 	- DATAPRODUCT: BI portal
	//
	// 	- PAGE: Dashboard
	//
	// 	- FULLPAGE: full-screen dashboards
	//
	// 	- REPORT: workbook
	//
	// example:
	//
	// PAGE
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 89713491-cb4f-4579-b889-e82c35f1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryWorksByWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksByWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *QueryWorksByWorkspaceRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryWorksByWorkspaceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryWorksByWorkspaceRequest) GetStatus() *int32 {
	return s.Status
}

func (s *QueryWorksByWorkspaceRequest) GetThirdPartAuthFlag() *int32 {
	return s.ThirdPartAuthFlag
}

func (s *QueryWorksByWorkspaceRequest) GetWorksType() *string {
	return s.WorksType
}

func (s *QueryWorksByWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryWorksByWorkspaceRequest) SetPageNum(v int32) *QueryWorksByWorkspaceRequest {
	s.PageNum = &v
	return s
}

func (s *QueryWorksByWorkspaceRequest) SetPageSize(v int32) *QueryWorksByWorkspaceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryWorksByWorkspaceRequest) SetStatus(v int32) *QueryWorksByWorkspaceRequest {
	s.Status = &v
	return s
}

func (s *QueryWorksByWorkspaceRequest) SetThirdPartAuthFlag(v int32) *QueryWorksByWorkspaceRequest {
	s.ThirdPartAuthFlag = &v
	return s
}

func (s *QueryWorksByWorkspaceRequest) SetWorksType(v string) *QueryWorksByWorkspaceRequest {
	s.WorksType = &v
	return s
}

func (s *QueryWorksByWorkspaceRequest) SetWorkspaceId(v string) *QueryWorksByWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryWorksByWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}

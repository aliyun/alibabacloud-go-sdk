// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSharedReportsResponseBody
	GetRequestId() *string
	SetResult(v *ListSharedReportsResponseBodyResult) *ListSharedReportsResponseBody
	GetResult() *ListSharedReportsResponseBodyResult
	SetSuccess(v bool) *ListSharedReportsResponseBody
	GetSuccess() *bool
}

type ListSharedReportsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The query results are returned.
	Result *ListSharedReportsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSharedReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSharedReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSharedReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSharedReportsResponseBody) GetResult() *ListSharedReportsResponseBodyResult {
	return s.Result
}

func (s *ListSharedReportsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSharedReportsResponseBody) SetRequestId(v string) *ListSharedReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSharedReportsResponseBody) SetResult(v *ListSharedReportsResponseBodyResult) *ListSharedReportsResponseBody {
	s.Result = v
	return s
}

func (s *ListSharedReportsResponseBody) SetSuccess(v bool) *ListSharedReportsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSharedReportsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSharedReportsResponseBodyResult struct {
	// The list of queried works.
	Data []*ListSharedReportsResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page set when the interface is requested.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of rows in the table.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListSharedReportsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSharedReportsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSharedReportsResponseBodyResult) GetData() []*ListSharedReportsResponseBodyResultData {
	return s.Data
}

func (s *ListSharedReportsResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListSharedReportsResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSharedReportsResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListSharedReportsResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListSharedReportsResponseBodyResult) SetData(v []*ListSharedReportsResponseBodyResultData) *ListSharedReportsResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListSharedReportsResponseBodyResult) SetPageNum(v int32) *ListSharedReportsResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListSharedReportsResponseBodyResult) SetPageSize(v int32) *ListSharedReportsResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListSharedReportsResponseBodyResult) SetTotalNum(v int32) *ListSharedReportsResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListSharedReportsResponseBodyResult) SetTotalPages(v int32) *ListSharedReportsResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *ListSharedReportsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListSharedReportsResponseBodyResultData struct {
	// Queries whether the user has collected the work.
	//
	// example:
	//
	// true
	Favorite *bool `json:"Favorite,omitempty" xml:"Favorite,omitempty"`
	// The timestamp when the work was created.
	//
	// example:
	//
	// 1640088615000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the work was modified.
	//
	// example:
	//
	// 1644373980000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The query user has the editing rights of the work.
	//
	// example:
	//
	// true
	HasEditAuth *bool `json:"HasEditAuth,omitempty" xml:"HasEditAuth,omitempty"`
	// The query user has the permission to view the work.
	//
	// example:
	//
	// true
	HasViewAuth *bool `json:"HasViewAuth,omitempty" xml:"HasViewAuth,omitempty"`
	// The name of the work.
	//
	// example:
	//
	// Test report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account name of the work owner.
	//
	// example:
	//
	// test account
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The UserID of the work owner.
	//
	// example:
	//
	// 1365*****238860
	OwnerNum *string `json:"OwnerNum,omitempty" xml:"OwnerNum,omitempty"`
	// The publication status of the work. Valid values:
	//
	// 	- 0: unpublished
	//
	// 	- 1: published
	//
	// 	- 2: modified and saved but not published.
	//
	// 	- 3: unpublished
	//
	// example:
	//
	// 1
	PublishStatus *int32 `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// The ID of the work.
	//
	// example:
	//
	// 977c7698-****-****-****-44b7304d20fc
	TreeId *string `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
	// The type of the work. Valid values:
	//
	// 	- DATAPRODUCT: BI portal
	//
	// 	- PAGE: Dashboard
	//
	// 	- REPORT: workbook
	//
	// example:
	//
	// PAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the workspace to which the work belongs.
	//
	// example:
	//
	// gfidm145-****-****-9426-8f93be23****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the work belongs.
	//
	// example:
	//
	// Test Workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListSharedReportsResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s ListSharedReportsResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListSharedReportsResponseBodyResultData) GetFavorite() *bool {
	return s.Favorite
}

func (s *ListSharedReportsResponseBodyResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListSharedReportsResponseBodyResultData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListSharedReportsResponseBodyResultData) GetHasEditAuth() *bool {
	return s.HasEditAuth
}

func (s *ListSharedReportsResponseBodyResultData) GetHasViewAuth() *bool {
	return s.HasViewAuth
}

func (s *ListSharedReportsResponseBodyResultData) GetName() *string {
	return s.Name
}

func (s *ListSharedReportsResponseBodyResultData) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListSharedReportsResponseBodyResultData) GetOwnerNum() *string {
	return s.OwnerNum
}

func (s *ListSharedReportsResponseBodyResultData) GetPublishStatus() *int32 {
	return s.PublishStatus
}

func (s *ListSharedReportsResponseBodyResultData) GetTreeId() *string {
	return s.TreeId
}

func (s *ListSharedReportsResponseBodyResultData) GetType() *string {
	return s.Type
}

func (s *ListSharedReportsResponseBodyResultData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListSharedReportsResponseBodyResultData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListSharedReportsResponseBodyResultData) SetFavorite(v bool) *ListSharedReportsResponseBodyResultData {
	s.Favorite = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetGmtCreate(v string) *ListSharedReportsResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetGmtModified(v string) *ListSharedReportsResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetHasEditAuth(v bool) *ListSharedReportsResponseBodyResultData {
	s.HasEditAuth = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetHasViewAuth(v bool) *ListSharedReportsResponseBodyResultData {
	s.HasViewAuth = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetName(v string) *ListSharedReportsResponseBodyResultData {
	s.Name = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetOwnerName(v string) *ListSharedReportsResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetOwnerNum(v string) *ListSharedReportsResponseBodyResultData {
	s.OwnerNum = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetPublishStatus(v int32) *ListSharedReportsResponseBodyResultData {
	s.PublishStatus = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetTreeId(v string) *ListSharedReportsResponseBodyResultData {
	s.TreeId = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetType(v string) *ListSharedReportsResponseBodyResultData {
	s.Type = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetWorkspaceId(v string) *ListSharedReportsResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) SetWorkspaceName(v string) *ListSharedReportsResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

func (s *ListSharedReportsResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}

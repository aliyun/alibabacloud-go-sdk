// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentViewReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRecentViewReportsResponseBody
	GetRequestId() *string
	SetResult(v *ListRecentViewReportsResponseBodyResult) *ListRecentViewReportsResponseBody
	GetResult() *ListRecentViewReportsResponseBodyResult
	SetSuccess(v bool) *ListRecentViewReportsResponseBody
	GetSuccess() *bool
}

type ListRecentViewReportsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The query results are returned.
	Result *ListRecentViewReportsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s ListRecentViewReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecentViewReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentViewReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecentViewReportsResponseBody) GetResult() *ListRecentViewReportsResponseBodyResult {
	return s.Result
}

func (s *ListRecentViewReportsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRecentViewReportsResponseBody) SetRequestId(v string) *ListRecentViewReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecentViewReportsResponseBody) SetResult(v *ListRecentViewReportsResponseBodyResult) *ListRecentViewReportsResponseBody {
	s.Result = v
	return s
}

func (s *ListRecentViewReportsResponseBody) SetSuccess(v bool) *ListRecentViewReportsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRecentViewReportsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRecentViewReportsResponseBodyResult struct {
	// Attention
	//
	// example:
	//
	// test
	Attention *string `json:"Attention,omitempty" xml:"Attention,omitempty"`
	// The list of queried works.
	Data []*ListRecentViewReportsResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// 	- Default value: 10.
	//
	// 	- Maximum of 100 articles
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

func (s ListRecentViewReportsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListRecentViewReportsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRecentViewReportsResponseBodyResult) GetAttention() *string {
	return s.Attention
}

func (s *ListRecentViewReportsResponseBodyResult) GetData() []*ListRecentViewReportsResponseBodyResultData {
	return s.Data
}

func (s *ListRecentViewReportsResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListRecentViewReportsResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecentViewReportsResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListRecentViewReportsResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListRecentViewReportsResponseBodyResult) SetAttention(v string) *ListRecentViewReportsResponseBodyResult {
	s.Attention = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResult) SetData(v []*ListRecentViewReportsResponseBodyResultData) *ListRecentViewReportsResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListRecentViewReportsResponseBodyResult) SetPageNum(v int32) *ListRecentViewReportsResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResult) SetPageSize(v int32) *ListRecentViewReportsResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResult) SetTotalNum(v int32) *ListRecentViewReportsResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResult) SetTotalPages(v int32) *ListRecentViewReportsResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListRecentViewReportsResponseBodyResultData struct {
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
	// 1496651577000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the work was modified.
	//
	// example:
	//
	// 1640595729000
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
	// The timestamp when the work was last accessed.
	//
	// example:
	//
	// 1642067498000
	LatestViewTime *string `json:"LatestViewTime,omitempty" xml:"LatestViewTime,omitempty"`
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
	// test
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
	// The number of times the work was accessed.
	//
	// example:
	//
	// 7
	ViewCount *int64 `json:"ViewCount,omitempty" xml:"ViewCount,omitempty"`
	// The ID of the workspace to which the work belongs.
	//
	// example:
	//
	// 523793cb-****-****-****-aa71c65ffa39
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the work belongs.
	//
	// example:
	//
	// Test Workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListRecentViewReportsResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s ListRecentViewReportsResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListRecentViewReportsResponseBodyResultData) GetFavorite() *bool {
	return s.Favorite
}

func (s *ListRecentViewReportsResponseBodyResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListRecentViewReportsResponseBodyResultData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListRecentViewReportsResponseBodyResultData) GetHasEditAuth() *bool {
	return s.HasEditAuth
}

func (s *ListRecentViewReportsResponseBodyResultData) GetHasViewAuth() *bool {
	return s.HasViewAuth
}

func (s *ListRecentViewReportsResponseBodyResultData) GetLatestViewTime() *string {
	return s.LatestViewTime
}

func (s *ListRecentViewReportsResponseBodyResultData) GetName() *string {
	return s.Name
}

func (s *ListRecentViewReportsResponseBodyResultData) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListRecentViewReportsResponseBodyResultData) GetOwnerNum() *string {
	return s.OwnerNum
}

func (s *ListRecentViewReportsResponseBodyResultData) GetPublishStatus() *int32 {
	return s.PublishStatus
}

func (s *ListRecentViewReportsResponseBodyResultData) GetTreeId() *string {
	return s.TreeId
}

func (s *ListRecentViewReportsResponseBodyResultData) GetType() *string {
	return s.Type
}

func (s *ListRecentViewReportsResponseBodyResultData) GetViewCount() *int64 {
	return s.ViewCount
}

func (s *ListRecentViewReportsResponseBodyResultData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListRecentViewReportsResponseBodyResultData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListRecentViewReportsResponseBodyResultData) SetFavorite(v bool) *ListRecentViewReportsResponseBodyResultData {
	s.Favorite = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetGmtCreate(v string) *ListRecentViewReportsResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetGmtModified(v string) *ListRecentViewReportsResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetHasEditAuth(v bool) *ListRecentViewReportsResponseBodyResultData {
	s.HasEditAuth = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetHasViewAuth(v bool) *ListRecentViewReportsResponseBodyResultData {
	s.HasViewAuth = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetLatestViewTime(v string) *ListRecentViewReportsResponseBodyResultData {
	s.LatestViewTime = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetName(v string) *ListRecentViewReportsResponseBodyResultData {
	s.Name = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetOwnerName(v string) *ListRecentViewReportsResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetOwnerNum(v string) *ListRecentViewReportsResponseBodyResultData {
	s.OwnerNum = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetPublishStatus(v int32) *ListRecentViewReportsResponseBodyResultData {
	s.PublishStatus = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetTreeId(v string) *ListRecentViewReportsResponseBodyResultData {
	s.TreeId = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetType(v string) *ListRecentViewReportsResponseBodyResultData {
	s.Type = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetViewCount(v int64) *ListRecentViewReportsResponseBodyResultData {
	s.ViewCount = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetWorkspaceId(v string) *ListRecentViewReportsResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) SetWorkspaceName(v string) *ListRecentViewReportsResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

func (s *ListRecentViewReportsResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}

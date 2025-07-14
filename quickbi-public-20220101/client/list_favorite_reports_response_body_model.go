// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFavoriteReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListFavoriteReportsResponseBody
	GetRequestId() *string
	SetResult(v *ListFavoriteReportsResponseBodyResult) *ListFavoriteReportsResponseBody
	GetResult() *ListFavoriteReportsResponseBodyResult
	SetSuccess(v bool) *ListFavoriteReportsResponseBody
	GetSuccess() *bool
}

type ListFavoriteReportsResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the query result.
	Result *ListFavoriteReportsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFavoriteReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFavoriteReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFavoriteReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFavoriteReportsResponseBody) GetResult() *ListFavoriteReportsResponseBodyResult {
	return s.Result
}

func (s *ListFavoriteReportsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListFavoriteReportsResponseBody) SetRequestId(v string) *ListFavoriteReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFavoriteReportsResponseBody) SetResult(v *ListFavoriteReportsResponseBodyResult) *ListFavoriteReportsResponseBody {
	s.Result = v
	return s
}

func (s *ListFavoriteReportsResponseBody) SetSuccess(v bool) *ListFavoriteReportsResponseBody {
	s.Success = &v
	return s
}

func (s *ListFavoriteReportsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFavoriteReportsResponseBodyResult struct {
	// List of works queried.
	Data []*ListFavoriteReportsResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of rows per page set when requesting the interface.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of rows.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListFavoriteReportsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListFavoriteReportsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFavoriteReportsResponseBodyResult) GetData() []*ListFavoriteReportsResponseBodyResultData {
	return s.Data
}

func (s *ListFavoriteReportsResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListFavoriteReportsResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFavoriteReportsResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListFavoriteReportsResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListFavoriteReportsResponseBodyResult) SetData(v []*ListFavoriteReportsResponseBodyResultData) *ListFavoriteReportsResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListFavoriteReportsResponseBodyResult) SetPageNum(v int32) *ListFavoriteReportsResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResult) SetPageSize(v int32) *ListFavoriteReportsResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResult) SetTotalNum(v int32) *ListFavoriteReportsResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResult) SetTotalPages(v int32) *ListFavoriteReportsResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListFavoriteReportsResponseBodyResultData struct {
	// Indicates whether the user has favorited the work.
	//
	// example:
	//
	// true
	Favorite *bool `json:"Favorite,omitempty" xml:"Favorite,omitempty"`
	// The timestamp when the work was favorited.
	//
	// example:
	//
	// 1640088615000
	FavoriteDate *string `json:"FavoriteDate,omitempty" xml:"FavoriteDate,omitempty"`
	// Timestamp of the work creation.
	//
	// example:
	//
	// 1640088615000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Timestamp of the work modification.
	//
	// example:
	//
	// 1640595729000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the user has edit permission for the work.
	//
	// example:
	//
	// true
	HasEditAuth *bool `json:"HasEditAuth,omitempty" xml:"HasEditAuth,omitempty"`
	// Check if the user has the permission to view the work.
	//
	// example:
	//
	// true
	HasViewAuth *bool `json:"HasViewAuth,omitempty" xml:"HasViewAuth,omitempty"`
	// Name of the work.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Alibaba Cloud account name of the work owner.
	//
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// UserID of the work owner.
	//
	// example:
	//
	// 1365*****238860
	OwnerNum *string `json:"OwnerNum,omitempty" xml:"OwnerNum,omitempty"`
	// Publication status of the work. Value range:
	//
	// - 0: Not published
	//
	// - 1: Published
	//
	// - 2: Saved with modifications, not published
	//
	// - 3: Offline
	//
	// example:
	//
	// 1
	PublishStatus *int32 `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// Work ID.
	//
	// example:
	//
	// 977c7698-****-****-****-44b7304d20fc
	TreeId *string `json:"TreeId,omitempty" xml:"TreeId,omitempty"`
	// Type of the work. Value range:
	//
	// - DATAPRODUCT: Data Portal
	//
	// - PAGE: Dashboard
	//
	// - REPORT: Spreadsheet
	//
	// example:
	//
	// PAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// test
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListFavoriteReportsResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s ListFavoriteReportsResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListFavoriteReportsResponseBodyResultData) GetFavorite() *bool {
	return s.Favorite
}

func (s *ListFavoriteReportsResponseBodyResultData) GetFavoriteDate() *string {
	return s.FavoriteDate
}

func (s *ListFavoriteReportsResponseBodyResultData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListFavoriteReportsResponseBodyResultData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListFavoriteReportsResponseBodyResultData) GetHasEditAuth() *bool {
	return s.HasEditAuth
}

func (s *ListFavoriteReportsResponseBodyResultData) GetHasViewAuth() *bool {
	return s.HasViewAuth
}

func (s *ListFavoriteReportsResponseBodyResultData) GetName() *string {
	return s.Name
}

func (s *ListFavoriteReportsResponseBodyResultData) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListFavoriteReportsResponseBodyResultData) GetOwnerNum() *string {
	return s.OwnerNum
}

func (s *ListFavoriteReportsResponseBodyResultData) GetPublishStatus() *int32 {
	return s.PublishStatus
}

func (s *ListFavoriteReportsResponseBodyResultData) GetTreeId() *string {
	return s.TreeId
}

func (s *ListFavoriteReportsResponseBodyResultData) GetType() *string {
	return s.Type
}

func (s *ListFavoriteReportsResponseBodyResultData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListFavoriteReportsResponseBodyResultData) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListFavoriteReportsResponseBodyResultData) SetFavorite(v bool) *ListFavoriteReportsResponseBodyResultData {
	s.Favorite = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetFavoriteDate(v string) *ListFavoriteReportsResponseBodyResultData {
	s.FavoriteDate = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetGmtCreate(v string) *ListFavoriteReportsResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetGmtModified(v string) *ListFavoriteReportsResponseBodyResultData {
	s.GmtModified = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetHasEditAuth(v bool) *ListFavoriteReportsResponseBodyResultData {
	s.HasEditAuth = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetHasViewAuth(v bool) *ListFavoriteReportsResponseBodyResultData {
	s.HasViewAuth = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetName(v string) *ListFavoriteReportsResponseBodyResultData {
	s.Name = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetOwnerName(v string) *ListFavoriteReportsResponseBodyResultData {
	s.OwnerName = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetOwnerNum(v string) *ListFavoriteReportsResponseBodyResultData {
	s.OwnerNum = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetPublishStatus(v int32) *ListFavoriteReportsResponseBodyResultData {
	s.PublishStatus = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetTreeId(v string) *ListFavoriteReportsResponseBodyResultData {
	s.TreeId = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetType(v string) *ListFavoriteReportsResponseBodyResultData {
	s.Type = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetWorkspaceId(v string) *ListFavoriteReportsResponseBodyResultData {
	s.WorkspaceId = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) SetWorkspaceName(v string) *ListFavoriteReportsResponseBodyResultData {
	s.WorkspaceName = &v
	return s
}

func (s *ListFavoriteReportsResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}

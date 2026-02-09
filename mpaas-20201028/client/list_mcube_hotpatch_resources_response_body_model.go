// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeHotpatchResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListHotpatchResourceResult(v *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) *ListMcubeHotpatchResourcesResponseBody
	GetListHotpatchResourceResult() *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult
	SetRequestId(v string) *ListMcubeHotpatchResourcesResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeHotpatchResourcesResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeHotpatchResourcesResponseBody
	GetResultMessage() *string
}

type ListMcubeHotpatchResourcesResponseBody struct {
	ListHotpatchResourceResult *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult `json:"ListHotpatchResourceResult,omitempty" xml:"ListHotpatchResourceResult,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcubeHotpatchResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeHotpatchResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeHotpatchResourcesResponseBody) GetListHotpatchResourceResult() *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult {
	return s.ListHotpatchResourceResult
}

func (s *ListMcubeHotpatchResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeHotpatchResourcesResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcubeHotpatchResourcesResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcubeHotpatchResourcesResponseBody) SetListHotpatchResourceResult(v *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) *ListMcubeHotpatchResourcesResponseBody {
	s.ListHotpatchResourceResult = v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBody) SetRequestId(v string) *ListMcubeHotpatchResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBody) SetResultCode(v string) *ListMcubeHotpatchResourcesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBody) SetResultMessage(v string) *ListMcubeHotpatchResourcesResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBody) Validate() error {
	if s.ListHotpatchResourceResult != nil {
		if err := s.ListHotpatchResourceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult struct {
	// example:
	//
	// 3
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// False
	HasMore              *bool                                                                                   `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	HotpatchResourceInfo []*ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo `json:"HotpatchResourceInfo,omitempty" xml:"HotpatchResourceInfo,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1C23E812-217E-5065-B778-D34586E2105E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 71
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) GoString() string {
	return s.String()
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) GetHotpatchResourceInfo() []*ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	return s.HotpatchResourceInfo
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) SetCurrentPage(v int32) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult {
	s.CurrentPage = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) SetErrorCode(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult {
	s.ErrorCode = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) SetHasMore(v bool) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult {
	s.HasMore = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) SetHotpatchResourceInfo(v []*ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult {
	s.HotpatchResourceInfo = v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) SetPageSize(v int32) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult {
	s.PageSize = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) SetRequestId(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult {
	s.RequestId = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) SetResultMsg(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) SetSuccess(v bool) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult {
	s.Success = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) SetTotalCount(v int64) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult {
	s.TotalCount = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResult) Validate() error {
	if s.HotpatchResourceInfo != nil {
		for _, item := range s.HotpatchResourceInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo struct {
	// example:
	//
	// ALIPUBE5C3F6D091419-default
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// example:
	//
	// xxx
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// https://xxxxx.jar
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// 528
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// 1745892911000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1574261514000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 81c90a2cafdc6dfc54201e70845b5708
	HotpatchVersion *string `json:"HotpatchVersion,omitempty" xml:"HotpatchVersion,omitempty"`
	// example:
	//
	// 1358
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// E268154063D1256B4E60FE82B48E0811
	Md5  *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// example:
	//
	// modifier
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// 0
	PackageId *int64 `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// example:
	//
	// iOS,Android
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// ALIPUBE5C3F6D091419_Android-default
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1.0.0
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// example:
	//
	// 0
	PublishPeriod *int64 `json:"PublishPeriod,omitempty" xml:"PublishPeriod,omitempty"`
	// example:
	//
	// 81c90a2cafdc6dfc54201e70845b5708
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// example:
	//
	// mpaas.jar
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
}

func (s ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GoString() string {
	return s.String()
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetCreator() *string {
	return s.Creator
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetHotpatchVersion() *string {
	return s.HotpatchVersion
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetMd5() *string {
	return s.Md5
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetMemo() *string {
	return s.Memo
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetModifier() *string {
	return s.Modifier
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetPackageId() *int64 {
	return s.PackageId
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetPlatform() *string {
	return s.Platform
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetProductId() *string {
	return s.ProductId
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetPublishPeriod() *int64 {
	return s.PublishPeriod
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) GetSourceName() *string {
	return s.SourceName
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetAppCode(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.AppCode = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetCreator(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.Creator = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetDownloadUrl(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.DownloadUrl = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetFileSize(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.FileSize = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetGmtCreate(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.GmtCreate = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetGmtModified(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.GmtModified = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetHotpatchVersion(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.HotpatchVersion = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetId(v int64) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.Id = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetMd5(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.Md5 = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetMemo(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.Memo = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetModifier(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.Modifier = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetPackageId(v int64) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.PackageId = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetPlatform(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.Platform = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetProductId(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.ProductId = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetProductVersion(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.ProductVersion = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetPublishPeriod(v int64) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.PublishPeriod = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetReleaseVersion(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.ReleaseVersion = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) SetSourceName(v string) *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo {
	s.SourceName = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponseBodyListHotpatchResourceResultHotpatchResourceInfo) Validate() error {
	return dara.Validate(s)
}

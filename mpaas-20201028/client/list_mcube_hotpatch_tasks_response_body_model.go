// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeHotpatchTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListHotpatchTasksResult(v *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) *ListMcubeHotpatchTasksResponseBody
	GetListHotpatchTasksResult() *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult
	SetRequestId(v string) *ListMcubeHotpatchTasksResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeHotpatchTasksResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeHotpatchTasksResponseBody
	GetResultMessage() *string
}

type ListMcubeHotpatchTasksResponseBody struct {
	ListHotpatchTasksResult *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult `json:"ListHotpatchTasksResult,omitempty" xml:"ListHotpatchTasksResult,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcubeHotpatchTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeHotpatchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeHotpatchTasksResponseBody) GetListHotpatchTasksResult() *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult {
	return s.ListHotpatchTasksResult
}

func (s *ListMcubeHotpatchTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeHotpatchTasksResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcubeHotpatchTasksResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcubeHotpatchTasksResponseBody) SetListHotpatchTasksResult(v *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) *ListMcubeHotpatchTasksResponseBody {
	s.ListHotpatchTasksResult = v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBody) SetRequestId(v string) *ListMcubeHotpatchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBody) SetResultCode(v string) *ListMcubeHotpatchTasksResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBody) SetResultMessage(v string) *ListMcubeHotpatchTasksResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBody) Validate() error {
	if s.ListHotpatchTasksResult != nil {
		if err := s.ListHotpatchTasksResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult struct {
	// example:
	//
	// OK
	ErrorCode        *string                                                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HotpatchTaskInfo []*ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo `json:"HotpatchTaskInfo,omitempty" xml:"HotpatchTaskInfo,omitempty" type:"Repeated"`
	// example:
	//
	// 6BD4C876-47B4-56CF-84C5-57389EE1EDFE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) GoString() string {
	return s.String()
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) GetHotpatchTaskInfo() []*ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	return s.HotpatchTaskInfo
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) SetErrorCode(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult {
	s.ErrorCode = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) SetHotpatchTaskInfo(v []*ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult {
	s.HotpatchTaskInfo = v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) SetRequestId(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult {
	s.RequestId = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) SetResultMsg(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) SetSuccess(v bool) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult {
	s.Success = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResult) Validate() error {
	if s.HotpatchTaskInfo != nil {
		for _, item := range s.HotpatchTaskInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo struct {
	// example:
	//
	// ALIPUB40DB571101207-default
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// example:
	//
	// ***
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 1751594649000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-10-29 18:01:32
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 2024-10-29 18:01:32
	GmtModifiedStr *string `json:"GmtModifiedStr,omitempty" xml:"GmtModifiedStr,omitempty"`
	GreyConfigInfo *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	// example:
	//
	// 2020-03-18 20:12:26
	GreyEndtime     *string `json:"GreyEndtime,omitempty" xml:"GreyEndtime,omitempty"`
	GreyEndtimeData *string `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	// example:
	//
	// 10
	GreyNum *int64 `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	// example:
	//
	// 1486
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// example:
	//
	// xxxx
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// 1664552
	PackageId *int64 `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// example:
	//
	// iOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// ALIPUB40DB571101207_ANDROID-default
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1.0.0
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// example:
	//
	// 1
	PublishMode *int64 `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	// example:
	//
	// 3
	PublishType *int64 `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// example:
	//
	// 81c90a2cafdc6dfc54201e70845b5708
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// example:
	//
	// 1786
	ResIds *string `json:"ResIds,omitempty" xml:"ResIds,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int64 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 825827
	WhitelistIds *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
}

func (s ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GoString() string {
	return s.String()
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetCreator() *string {
	return s.Creator
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetGmtModifiedStr() *string {
	return s.GmtModifiedStr
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetGreyEndtime() *string {
	return s.GreyEndtime
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetGreyNum() *int64 {
	return s.GreyNum
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetMemo() *string {
	return s.Memo
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetModifier() *string {
	return s.Modifier
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetPackageId() *int64 {
	return s.PackageId
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetPlatform() *string {
	return s.Platform
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetProductId() *string {
	return s.ProductId
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetPublishMode() *int64 {
	return s.PublishMode
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetPublishType() *int64 {
	return s.PublishType
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetResIds() *string {
	return s.ResIds
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetTaskStatus() *int64 {
	return s.TaskStatus
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetAppCode(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.AppCode = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetCreator(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.Creator = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetGmtCreate(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.GmtCreate = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetGmtModified(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.GmtModified = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetGmtModifiedStr(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.GmtModifiedStr = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetGreyConfigInfo(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.GreyConfigInfo = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetGreyEndtime(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.GreyEndtime = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetGreyEndtimeData(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.GreyEndtimeData = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetGreyNum(v int64) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.GreyNum = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetId(v int64) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.Id = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetMemo(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.Memo = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetModifier(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.Modifier = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetPackageId(v int64) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.PackageId = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetPlatform(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.Platform = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetProductId(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.ProductId = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetProductVersion(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.ProductVersion = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetPublishMode(v int64) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.PublishMode = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetPublishType(v int64) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.PublishType = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetReleaseVersion(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.ReleaseVersion = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetResIds(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.ResIds = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetTaskStatus(v int64) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.TaskStatus = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) SetWhitelistIds(v string) *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo {
	s.WhitelistIds = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponseBodyListHotpatchTasksResultHotpatchTaskInfo) Validate() error {
	return dara.Validate(s)
}

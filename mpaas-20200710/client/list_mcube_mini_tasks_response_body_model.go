// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeMiniTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListMiniTaskResult(v *ListMcubeMiniTasksResponseBodyListMiniTaskResult) *ListMcubeMiniTasksResponseBody
	GetListMiniTaskResult() *ListMcubeMiniTasksResponseBodyListMiniTaskResult
	SetRequestId(v string) *ListMcubeMiniTasksResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeMiniTasksResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeMiniTasksResponseBody
	GetResultMessage() *string
}

type ListMcubeMiniTasksResponseBody struct {
	ListMiniTaskResult *ListMcubeMiniTasksResponseBodyListMiniTaskResult `json:"ListMiniTaskResult,omitempty" xml:"ListMiniTaskResult,omitempty" type:"Struct"`
	RequestId          *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                                           `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                                           `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcubeMiniTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniTasksResponseBody) GetListMiniTaskResult() *ListMcubeMiniTasksResponseBodyListMiniTaskResult {
	return s.ListMiniTaskResult
}

func (s *ListMcubeMiniTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeMiniTasksResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcubeMiniTasksResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcubeMiniTasksResponseBody) SetListMiniTaskResult(v *ListMcubeMiniTasksResponseBodyListMiniTaskResult) *ListMcubeMiniTasksResponseBody {
	s.ListMiniTaskResult = v
	return s
}

func (s *ListMcubeMiniTasksResponseBody) SetRequestId(v string) *ListMcubeMiniTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBody) SetResultCode(v string) *ListMcubeMiniTasksResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBody) SetResultMessage(v string) *ListMcubeMiniTasksResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBody) Validate() error {
	if s.ListMiniTaskResult != nil {
		if err := s.ListMiniTaskResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcubeMiniTasksResponseBodyListMiniTaskResult struct {
	MiniTaskList []*ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList `json:"MiniTaskList,omitempty" xml:"MiniTaskList,omitempty" type:"Repeated"`
	ResultMsg    *string                                                         `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success      *bool                                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMcubeMiniTasksResponseBodyListMiniTaskResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniTasksResponseBodyListMiniTaskResult) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResult) GetMiniTaskList() []*ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	return s.MiniTaskList
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResult) SetMiniTaskList(v []*ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) *ListMcubeMiniTasksResponseBodyListMiniTaskResult {
	s.MiniTaskList = v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResult) SetResultMsg(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResult) SetSuccess(v bool) *ListMcubeMiniTasksResponseBodyListMiniTaskResult {
	s.Success = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResult) Validate() error {
	if s.MiniTaskList != nil {
		for _, item := range s.MiniTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList struct {
	AppCode         *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GreyConfigInfo  *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtime     *string `json:"GreyEndtime,omitempty" xml:"GreyEndtime,omitempty"`
	GreyEndtimeData *string `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	GreyNum         *int64  `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Memo            *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	PackageId       *int64  `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	Platform        *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProductVersion  *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	PublishMode     *int64  `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	PublishType     *int64  `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskStatus      *int64  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	WhitelistIds    *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
}

func (s ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetGreyEndtime() *string {
	return s.GreyEndtime
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetGreyNum() *int64 {
	return s.GreyNum
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetMemo() *string {
	return s.Memo
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetPackageId() *int64 {
	return s.PackageId
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetPlatform() *string {
	return s.Platform
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetPublishMode() *int64 {
	return s.PublishMode
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetPublishType() *int64 {
	return s.PublishType
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetStatus() *string {
	return s.Status
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetTaskStatus() *int64 {
	return s.TaskStatus
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetAppCode(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.AppCode = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetGmtCreate(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.GmtCreate = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetGmtModified(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.GmtModified = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetGreyConfigInfo(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.GreyConfigInfo = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetGreyEndtime(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.GreyEndtime = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetGreyEndtimeData(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.GreyEndtimeData = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetGreyNum(v int64) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.GreyNum = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetId(v int64) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.Id = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetMemo(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.Memo = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetPackageId(v int64) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.PackageId = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetPlatform(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.Platform = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetProductVersion(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.ProductVersion = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetPublishMode(v int64) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.PublishMode = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetPublishType(v int64) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.PublishType = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetStatus(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.Status = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetTaskStatus(v int64) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.TaskStatus = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) SetWhitelistIds(v string) *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList {
	s.WhitelistIds = &v
	return s
}

func (s *ListMcubeMiniTasksResponseBodyListMiniTaskResultMiniTaskList) Validate() error {
	return dara.Validate(s)
}

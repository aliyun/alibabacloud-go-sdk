// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeMiniTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryMiniTaskResult(v *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult) *QueryMcubeMiniTaskResponseBody
	GetQueryMiniTaskResult() *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult
	SetRequestId(v string) *QueryMcubeMiniTaskResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMcubeMiniTaskResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *QueryMcubeMiniTaskResponseBody
	GetResultMessage() *string
}

type QueryMcubeMiniTaskResponseBody struct {
	QueryMiniTaskResult *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult `json:"QueryMiniTaskResult,omitempty" xml:"QueryMiniTaskResult,omitempty" type:"Struct"`
	RequestId           *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode          *string                                            `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage       *string                                            `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMcubeMiniTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeMiniTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMcubeMiniTaskResponseBody) GetQueryMiniTaskResult() *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult {
	return s.QueryMiniTaskResult
}

func (s *QueryMcubeMiniTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMcubeMiniTaskResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMcubeMiniTaskResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMcubeMiniTaskResponseBody) SetQueryMiniTaskResult(v *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult) *QueryMcubeMiniTaskResponseBody {
	s.QueryMiniTaskResult = v
	return s
}

func (s *QueryMcubeMiniTaskResponseBody) SetRequestId(v string) *QueryMcubeMiniTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBody) SetResultCode(v string) *QueryMcubeMiniTaskResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBody) SetResultMessage(v string) *QueryMcubeMiniTaskResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult struct {
	MiniTaskInfo *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo `json:"MiniTaskInfo,omitempty" xml:"MiniTaskInfo,omitempty" type:"Struct"`
	ResultMsg    *string                                                        `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success      *bool                                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult) GoString() string {
	return s.String()
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult) GetMiniTaskInfo() *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	return s.MiniTaskInfo
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult) SetMiniTaskInfo(v *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult {
	s.MiniTaskInfo = v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult) SetResultMsg(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult {
	s.ResultMsg = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult) SetSuccess(v bool) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult {
	s.Success = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResult) Validate() error {
	return dara.Validate(s)
}

type QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo struct {
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

func (s QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GoString() string {
	return s.String()
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetAppCode() *string {
	return s.AppCode
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetGreyEndtime() *string {
	return s.GreyEndtime
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetGreyNum() *int64 {
	return s.GreyNum
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetId() *int64 {
	return s.Id
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetMemo() *string {
	return s.Memo
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetPackageId() *int64 {
	return s.PackageId
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetPlatform() *string {
	return s.Platform
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetPublishMode() *int64 {
	return s.PublishMode
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetPublishType() *int64 {
	return s.PublishType
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetTaskStatus() *int64 {
	return s.TaskStatus
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetAppCode(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.AppCode = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetGmtCreate(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.GmtCreate = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetGmtModified(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.GmtModified = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetGreyConfigInfo(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.GreyConfigInfo = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetGreyEndtime(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.GreyEndtime = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetGreyEndtimeData(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.GreyEndtimeData = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetGreyNum(v int64) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.GreyNum = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetId(v int64) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.Id = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetMemo(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.Memo = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetPackageId(v int64) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.PackageId = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetPlatform(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.Platform = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetProductVersion(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.ProductVersion = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetPublishMode(v int64) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.PublishMode = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetPublishType(v int64) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.PublishType = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetStatus(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.Status = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetTaskStatus(v int64) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.TaskStatus = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) SetWhitelistIds(v string) *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo {
	s.WhitelistIds = &v
	return s
}

func (s *QueryMcubeMiniTaskResponseBodyQueryMiniTaskResultMiniTaskInfo) Validate() error {
	return dara.Validate(s)
}

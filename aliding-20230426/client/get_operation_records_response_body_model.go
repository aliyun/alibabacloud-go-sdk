// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOperationRecordsResponseBody
	GetRequestId() *string
	SetResult(v []*GetOperationRecordsResponseBodyResult) *GetOperationRecordsResponseBody
	GetResult() []*GetOperationRecordsResponseBodyResult
	SetVendorRequestId(v string) *GetOperationRecordsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetOperationRecordsResponseBody
	GetVendorType() *string
}

type GetOperationRecordsResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*GetOperationRecordsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetOperationRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOperationRecordsResponseBody) GetResult() []*GetOperationRecordsResponseBodyResult {
	return s.Result
}

func (s *GetOperationRecordsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetOperationRecordsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetOperationRecordsResponseBody) SetRequestId(v string) *GetOperationRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOperationRecordsResponseBody) SetResult(v []*GetOperationRecordsResponseBodyResult) *GetOperationRecordsResponseBody {
	s.Result = v
	return s
}

func (s *GetOperationRecordsResponseBody) SetVendorRequestId(v string) *GetOperationRecordsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetOperationRecordsResponseBody) SetVendorType(v string) *GetOperationRecordsResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetOperationRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOperationRecordsResponseBodyResult struct {
	// example:
	//
	// return
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// 同意
	ActionExit *string `json:"ActionExit,omitempty" xml:"ActionExit,omitempty"`
	// example:
	//
	// 2021-02-01
	ActiveTimeGMT *string `json:"ActiveTimeGMT,omitempty" xml:"ActiveTimeGMT,omitempty"`
	// example:
	//
	// act-xxaanfaf
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 12345
	DataId *int64 `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// https://oss.com/Signature.pdf
	DigitalSign *string `json:"DigitalSign,omitempty" xml:"DigitalSign,omitempty"`
	// example:
	//
	// https://oss.com/a.pdf
	Files *string `json:"Files,omitempty" xml:"Files,omitempty"`
	// example:
	//
	// 2021-01-01
	OperateTimeGMT *string `json:"OperateTimeGMT,omitempty" xml:"OperateTimeGMT,omitempty"`
	// example:
	//
	// remove
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// 张三
	OperatorDisplayName *string `json:"OperatorDisplayName,omitempty" xml:"OperatorDisplayName,omitempty"`
	// example:
	//
	// 李四
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// example:
	//
	// 无冬
	OperatorNickName *string `json:"OperatorNickName,omitempty" xml:"OperatorNickName,omitempty"`
	// example:
	//
	// https://oss.com/a.jpeg
	OperatorPhotoUrl *string `json:"OperatorPhotoUrl,omitempty" xml:"OperatorPhotoUrl,omitempty"`
	// example:
	//
	// 良好
	OperatorStatus *string `json:"OperatorStatus,omitempty" xml:"OperatorStatus,omitempty"`
	// example:
	//
	// manager123
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
	// example:
	//
	// f30233fb-72e1-4af4-8cb8-c7e0ea9ee530
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// example:
	//
	// 确认同意
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 请购类型
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// example:
	//
	// 12
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 同步
	TaskExecuteType *string `json:"TaskExecuteType,omitempty" xml:"TaskExecuteType,omitempty"`
	// example:
	//
	// 2021-01-01
	TaskHoldTimeGMT *int64 `json:"TaskHoldTimeGMT,omitempty" xml:"TaskHoldTimeGMT,omitempty"`
	// example:
	//
	// task-123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// append task
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// i18n
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetOperationRecordsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsResponseBodyResult) GetAction() *string {
	return s.Action
}

func (s *GetOperationRecordsResponseBodyResult) GetActionExit() *string {
	return s.ActionExit
}

func (s *GetOperationRecordsResponseBodyResult) GetActiveTimeGMT() *string {
	return s.ActiveTimeGMT
}

func (s *GetOperationRecordsResponseBodyResult) GetActivityId() *string {
	return s.ActivityId
}

func (s *GetOperationRecordsResponseBodyResult) GetDataId() *int64 {
	return s.DataId
}

func (s *GetOperationRecordsResponseBodyResult) GetDigitalSign() *string {
	return s.DigitalSign
}

func (s *GetOperationRecordsResponseBodyResult) GetFiles() *string {
	return s.Files
}

func (s *GetOperationRecordsResponseBodyResult) GetOperateTimeGMT() *string {
	return s.OperateTimeGMT
}

func (s *GetOperationRecordsResponseBodyResult) GetOperateType() *string {
	return s.OperateType
}

func (s *GetOperationRecordsResponseBodyResult) GetOperatorDisplayName() *string {
	return s.OperatorDisplayName
}

func (s *GetOperationRecordsResponseBodyResult) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetOperationRecordsResponseBodyResult) GetOperatorNickName() *string {
	return s.OperatorNickName
}

func (s *GetOperationRecordsResponseBodyResult) GetOperatorPhotoUrl() *string {
	return s.OperatorPhotoUrl
}

func (s *GetOperationRecordsResponseBodyResult) GetOperatorStatus() *string {
	return s.OperatorStatus
}

func (s *GetOperationRecordsResponseBodyResult) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *GetOperationRecordsResponseBodyResult) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetOperationRecordsResponseBodyResult) GetRemark() *string {
	return s.Remark
}

func (s *GetOperationRecordsResponseBodyResult) GetShowName() *string {
	return s.ShowName
}

func (s *GetOperationRecordsResponseBodyResult) GetSize() *int32 {
	return s.Size
}

func (s *GetOperationRecordsResponseBodyResult) GetTaskExecuteType() *string {
	return s.TaskExecuteType
}

func (s *GetOperationRecordsResponseBodyResult) GetTaskHoldTimeGMT() *int64 {
	return s.TaskHoldTimeGMT
}

func (s *GetOperationRecordsResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetOperationRecordsResponseBodyResult) GetTaskType() *string {
	return s.TaskType
}

func (s *GetOperationRecordsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetOperationRecordsResponseBodyResult) SetAction(v string) *GetOperationRecordsResponseBodyResult {
	s.Action = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetActionExit(v string) *GetOperationRecordsResponseBodyResult {
	s.ActionExit = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetActiveTimeGMT(v string) *GetOperationRecordsResponseBodyResult {
	s.ActiveTimeGMT = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetActivityId(v string) *GetOperationRecordsResponseBodyResult {
	s.ActivityId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetDataId(v int64) *GetOperationRecordsResponseBodyResult {
	s.DataId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetDigitalSign(v string) *GetOperationRecordsResponseBodyResult {
	s.DigitalSign = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetFiles(v string) *GetOperationRecordsResponseBodyResult {
	s.Files = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperateTimeGMT(v string) *GetOperationRecordsResponseBodyResult {
	s.OperateTimeGMT = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperateType(v string) *GetOperationRecordsResponseBodyResult {
	s.OperateType = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorDisplayName(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorDisplayName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorName(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorNickName(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorNickName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorPhotoUrl(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorPhotoUrl = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorStatus(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorStatus = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorUserId(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorUserId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetProcessInstanceId(v string) *GetOperationRecordsResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetRemark(v string) *GetOperationRecordsResponseBodyResult {
	s.Remark = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetShowName(v string) *GetOperationRecordsResponseBodyResult {
	s.ShowName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetSize(v int32) *GetOperationRecordsResponseBodyResult {
	s.Size = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskExecuteType(v string) *GetOperationRecordsResponseBodyResult {
	s.TaskExecuteType = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskHoldTimeGMT(v int64) *GetOperationRecordsResponseBodyResult {
	s.TaskHoldTimeGMT = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskId(v string) *GetOperationRecordsResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskType(v string) *GetOperationRecordsResponseBodyResult {
	s.TaskType = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetType(v string) *GetOperationRecordsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

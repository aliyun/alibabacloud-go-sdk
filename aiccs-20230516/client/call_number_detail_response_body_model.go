// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallNumberDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CallNumberDetailResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int64) *CallNumberDetailResponseBody
	GetCode() *int64
	SetMessage(v string) *CallNumberDetailResponseBody
	GetMessage() *string
	SetModel(v *CallNumberDetailResponseBodyModel) *CallNumberDetailResponseBody
	GetModel() *CallNumberDetailResponseBodyModel
	SetRequestId(v string) *CallNumberDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CallNumberDetailResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *CallNumberDetailResponseBody
	GetTimestamp() *int64
}

type CallNumberDetailResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *CallNumberDetailResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 48
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s CallNumberDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CallNumberDetailResponseBody) GoString() string {
	return s.String()
}

func (s *CallNumberDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CallNumberDetailResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CallNumberDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CallNumberDetailResponseBody) GetModel() *CallNumberDetailResponseBodyModel {
	return s.Model
}

func (s *CallNumberDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CallNumberDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CallNumberDetailResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CallNumberDetailResponseBody) SetAccessDeniedDetail(v string) *CallNumberDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CallNumberDetailResponseBody) SetCode(v int64) *CallNumberDetailResponseBody {
	s.Code = &v
	return s
}

func (s *CallNumberDetailResponseBody) SetMessage(v string) *CallNumberDetailResponseBody {
	s.Message = &v
	return s
}

func (s *CallNumberDetailResponseBody) SetModel(v *CallNumberDetailResponseBodyModel) *CallNumberDetailResponseBody {
	s.Model = v
	return s
}

func (s *CallNumberDetailResponseBody) SetRequestId(v string) *CallNumberDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *CallNumberDetailResponseBody) SetSuccess(v bool) *CallNumberDetailResponseBody {
	s.Success = &v
	return s
}

func (s *CallNumberDetailResponseBody) SetTimestamp(v int64) *CallNumberDetailResponseBody {
	s.Timestamp = &v
	return s
}

func (s *CallNumberDetailResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CallNumberDetailResponseBodyModel struct {
	// ai计费时长
	//
	// example:
	//
	// 1
	AiBill *int64 `json:"AiBill,omitempty" xml:"AiBill,omitempty"`
	// 接通转接类型 1-不转人工；2-接通转人工；3-智能转人工
	//
	// example:
	//
	// 1
	AnswerTransferType *int64 `json:"AnswerTransferType,omitempty" xml:"AnswerTransferType,omitempty"`
	// 导入号码时返回的批次号
	//
	// example:
	//
	// 1
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 通话时长，单位为毫秒，实际计费需向上取整转换为秒
	//
	// example:
	//
	// 1
	Bill *int64 `json:"Bill,omitempty" xml:"Bill,omitempty"`
	// 转接计费时长
	//
	// example:
	//
	// 1
	BridgeBill *int64 `json:"BridgeBill,omitempty" xml:"BridgeBill,omitempty"`
	// 每次呼叫的唯一标识
	//
	// example:
	//
	// a
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 可选项 1-AI外呼，6-语音通知
	//
	// example:
	//
	// 1
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 客户详情url
	//
	// example:
	//
	// a
	ClientUrl *string `json:"ClientUrl,omitempty" xml:"ClientUrl,omitempty"`
	// 创建时间 格式 2026-01-01 00:00:00
	//
	// example:
	//
	// 2026-01-01 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 通话轮次
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 线路id
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// 号码编号
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 意向标签
	//
	// example:
	//
	// a
	IntentTag *string `json:"IntentTag,omitempty" xml:"IntentTag,omitempty"`
	// 回复关键词
	//
	// example:
	//
	// a
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// 外呼号码
	//
	// example:
	//
	// a
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 外呼号码MD5
	//
	// example:
	//
	// a
	NumberMd5 *string `json:"NumberMd5,omitempty" xml:"NumberMd5,omitempty"`
	// 自定义参数 json
	//
	// example:
	//
	// {"名字":"a"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// 个性标签
	//
	// example:
	//
	// a
	PersonalityTag *string `json:"PersonalityTag,omitempty" xml:"PersonalityTag,omitempty"`
	// 备注信息
	//
	// example:
	//
	// a
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 坐席id
	//
	// example:
	//
	// 1
	Sid *int64 `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// 外呼状态编码
	//
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 用户自定义标签
	//
	// example:
	//
	// a
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// 任务ID
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 外呼使用的模板ID
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 转人工状态 0-未触发,
	//
	// example:
	//
	// 0
	TransferStatus *int64 `json:"TransferStatus,omitempty" xml:"TransferStatus,omitempty"`
}

func (s CallNumberDetailResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s CallNumberDetailResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CallNumberDetailResponseBodyModel) GetAiBill() *int64 {
	return s.AiBill
}

func (s *CallNumberDetailResponseBodyModel) GetAnswerTransferType() *int64 {
	return s.AnswerTransferType
}

func (s *CallNumberDetailResponseBodyModel) GetBatchId() *string {
	return s.BatchId
}

func (s *CallNumberDetailResponseBodyModel) GetBill() *int64 {
	return s.Bill
}

func (s *CallNumberDetailResponseBodyModel) GetBridgeBill() *int64 {
	return s.BridgeBill
}

func (s *CallNumberDetailResponseBodyModel) GetCallId() *string {
	return s.CallId
}

func (s *CallNumberDetailResponseBodyModel) GetCallType() *int64 {
	return s.CallType
}

func (s *CallNumberDetailResponseBodyModel) GetClientUrl() *string {
	return s.ClientUrl
}

func (s *CallNumberDetailResponseBodyModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CallNumberDetailResponseBodyModel) GetDuration() *int64 {
	return s.Duration
}

func (s *CallNumberDetailResponseBodyModel) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *CallNumberDetailResponseBodyModel) GetId() *int64 {
	return s.Id
}

func (s *CallNumberDetailResponseBodyModel) GetIntentTag() *string {
	return s.IntentTag
}

func (s *CallNumberDetailResponseBodyModel) GetKeywords() *string {
	return s.Keywords
}

func (s *CallNumberDetailResponseBodyModel) GetNumber() *string {
	return s.Number
}

func (s *CallNumberDetailResponseBodyModel) GetNumberMd5() *string {
	return s.NumberMd5
}

func (s *CallNumberDetailResponseBodyModel) GetParams() *string {
	return s.Params
}

func (s *CallNumberDetailResponseBodyModel) GetPersonalityTag() *string {
	return s.PersonalityTag
}

func (s *CallNumberDetailResponseBodyModel) GetRemark() *string {
	return s.Remark
}

func (s *CallNumberDetailResponseBodyModel) GetSid() *int64 {
	return s.Sid
}

func (s *CallNumberDetailResponseBodyModel) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *CallNumberDetailResponseBodyModel) GetTag() *string {
	return s.Tag
}

func (s *CallNumberDetailResponseBodyModel) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CallNumberDetailResponseBodyModel) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CallNumberDetailResponseBodyModel) GetTransferStatus() *int64 {
	return s.TransferStatus
}

func (s *CallNumberDetailResponseBodyModel) SetAiBill(v int64) *CallNumberDetailResponseBodyModel {
	s.AiBill = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetAnswerTransferType(v int64) *CallNumberDetailResponseBodyModel {
	s.AnswerTransferType = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetBatchId(v string) *CallNumberDetailResponseBodyModel {
	s.BatchId = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetBill(v int64) *CallNumberDetailResponseBodyModel {
	s.Bill = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetBridgeBill(v int64) *CallNumberDetailResponseBodyModel {
	s.BridgeBill = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetCallId(v string) *CallNumberDetailResponseBodyModel {
	s.CallId = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetCallType(v int64) *CallNumberDetailResponseBodyModel {
	s.CallType = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetClientUrl(v string) *CallNumberDetailResponseBodyModel {
	s.ClientUrl = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetCreateTime(v string) *CallNumberDetailResponseBodyModel {
	s.CreateTime = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetDuration(v int64) *CallNumberDetailResponseBodyModel {
	s.Duration = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetGatewayId(v int64) *CallNumberDetailResponseBodyModel {
	s.GatewayId = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetId(v int64) *CallNumberDetailResponseBodyModel {
	s.Id = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetIntentTag(v string) *CallNumberDetailResponseBodyModel {
	s.IntentTag = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetKeywords(v string) *CallNumberDetailResponseBodyModel {
	s.Keywords = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetNumber(v string) *CallNumberDetailResponseBodyModel {
	s.Number = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetNumberMd5(v string) *CallNumberDetailResponseBodyModel {
	s.NumberMd5 = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetParams(v string) *CallNumberDetailResponseBodyModel {
	s.Params = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetPersonalityTag(v string) *CallNumberDetailResponseBodyModel {
	s.PersonalityTag = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetRemark(v string) *CallNumberDetailResponseBodyModel {
	s.Remark = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetSid(v int64) *CallNumberDetailResponseBodyModel {
	s.Sid = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetStatusCode(v int64) *CallNumberDetailResponseBodyModel {
	s.StatusCode = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetTag(v string) *CallNumberDetailResponseBodyModel {
	s.Tag = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetTaskId(v int64) *CallNumberDetailResponseBodyModel {
	s.TaskId = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetTemplateId(v int64) *CallNumberDetailResponseBodyModel {
	s.TemplateId = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) SetTransferStatus(v int64) *CallNumberDetailResponseBodyModel {
	s.TransferStatus = &v
	return s
}

func (s *CallNumberDetailResponseBodyModel) Validate() error {
	return dara.Validate(s)
}

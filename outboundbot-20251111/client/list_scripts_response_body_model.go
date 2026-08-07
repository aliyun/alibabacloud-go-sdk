// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScriptsResponseBody
	GetCode() *string
	SetData(v *ListScriptsResponseBodyData) *ListScriptsResponseBody
	GetData() *ListScriptsResponseBodyData
	SetHttpStatusCode(v int32) *ListScriptsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScriptsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListScriptsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListScriptsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListScriptsResponseBody
	GetSuccess() *bool
}

type ListScriptsResponseBody struct {
	// 返回码
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *ListScriptsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP状态码
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 错误信息
	//
	// example:
	//
	// Instance does not exist. Instance=ob-0987654321
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 错误信息中的变量值列表
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// 请求ID
	//
	// example:
	//
	// 019FDAC7-13C5-1B64-A853-999DF105B9EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否调用成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScriptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScriptsResponseBody) GetData() *ListScriptsResponseBodyData {
	return s.Data
}

func (s *ListScriptsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScriptsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScriptsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListScriptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScriptsResponseBody) SetCode(v string) *ListScriptsResponseBody {
	s.Code = &v
	return s
}

func (s *ListScriptsResponseBody) SetData(v *ListScriptsResponseBodyData) *ListScriptsResponseBody {
	s.Data = v
	return s
}

func (s *ListScriptsResponseBody) SetHttpStatusCode(v int32) *ListScriptsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScriptsResponseBody) SetMessage(v string) *ListScriptsResponseBody {
	s.Message = &v
	return s
}

func (s *ListScriptsResponseBody) SetParams(v []*string) *ListScriptsResponseBody {
	s.Params = v
	return s
}

func (s *ListScriptsResponseBody) SetRequestId(v string) *ListScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptsResponseBody) SetSuccess(v bool) *ListScriptsResponseBody {
	s.Success = &v
	return s
}

func (s *ListScriptsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScriptsResponseBodyData struct {
	// 页码，从1开始
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页记录数
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 数据列表
	Scripts []*ListScriptsResponseBodyDataScripts `json:"Scripts,omitempty" xml:"Scripts,omitempty" type:"Repeated"`
	// 符合条件的记录总数
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScriptsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptsResponseBodyData) GetScripts() []*ListScriptsResponseBodyDataScripts {
	return s.Scripts
}

func (s *ListScriptsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListScriptsResponseBodyData) SetPageNumber(v int32) *ListScriptsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListScriptsResponseBodyData) SetPageSize(v int32) *ListScriptsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListScriptsResponseBodyData) SetScripts(v []*ListScriptsResponseBodyDataScripts) *ListScriptsResponseBodyData {
	s.Scripts = v
	return s
}

func (s *ListScriptsResponseBodyData) SetTotalCount(v int32) *ListScriptsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListScriptsResponseBodyData) Validate() error {
	if s.Scripts != nil {
		for _, item := range s.Scripts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptsResponseBodyDataScripts struct {
	// 并发数
	//
	// example:
	//
	// 10
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// 创建时间，毫秒级时间戳
	//
	// example:
	//
	// 1735660800000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 描述
	//
	// example:
	//
	// 询问用户对本次服务是否满意
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 草稿版本ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b60
	DraftVersionId *string `json:"DraftVersionId,omitempty" xml:"DraftVersionId,omitempty"`
	// 名称
	//
	// example:
	//
	// 满意度调研
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// NLU接入类型
	//
	// example:
	//
	// MANAGED
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// NLU引擎类型
	//
	// example:
	//
	// BEEBOT
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// 电话号码
	//
	// example:
	//
	// 01057316547
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 已发布版本ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b41
	PublishedVersionId *string `json:"PublishedVersionId,omitempty" xml:"PublishedVersionId,omitempty"`
	// 场景ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// 状态
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间，毫秒级时间戳
	//
	// example:
	//
	// 1735660800000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListScriptsResponseBodyDataScripts) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponseBodyDataScripts) GoString() string {
	return s.String()
}

func (s *ListScriptsResponseBodyDataScripts) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *ListScriptsResponseBodyDataScripts) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListScriptsResponseBodyDataScripts) GetDescription() *string {
	return s.Description
}

func (s *ListScriptsResponseBodyDataScripts) GetDraftVersionId() *string {
	return s.DraftVersionId
}

func (s *ListScriptsResponseBodyDataScripts) GetName() *string {
	return s.Name
}

func (s *ListScriptsResponseBodyDataScripts) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *ListScriptsResponseBodyDataScripts) GetNluEngine() *string {
	return s.NluEngine
}

func (s *ListScriptsResponseBodyDataScripts) GetNumber() *string {
	return s.Number
}

func (s *ListScriptsResponseBodyDataScripts) GetPublishedVersionId() *string {
	return s.PublishedVersionId
}

func (s *ListScriptsResponseBodyDataScripts) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptsResponseBodyDataScripts) GetStatus() *string {
	return s.Status
}

func (s *ListScriptsResponseBodyDataScripts) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListScriptsResponseBodyDataScripts) SetConcurrency(v int32) *ListScriptsResponseBodyDataScripts {
	s.Concurrency = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetCreatedTime(v int64) *ListScriptsResponseBodyDataScripts {
	s.CreatedTime = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetDescription(v string) *ListScriptsResponseBodyDataScripts {
	s.Description = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetDraftVersionId(v string) *ListScriptsResponseBodyDataScripts {
	s.DraftVersionId = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetName(v string) *ListScriptsResponseBodyDataScripts {
	s.Name = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetNluAccessType(v string) *ListScriptsResponseBodyDataScripts {
	s.NluAccessType = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetNluEngine(v string) *ListScriptsResponseBodyDataScripts {
	s.NluEngine = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetNumber(v string) *ListScriptsResponseBodyDataScripts {
	s.Number = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetPublishedVersionId(v string) *ListScriptsResponseBodyDataScripts {
	s.PublishedVersionId = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetScriptId(v string) *ListScriptsResponseBodyDataScripts {
	s.ScriptId = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetStatus(v string) *ListScriptsResponseBodyDataScripts {
	s.Status = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) SetUpdatedTime(v int64) *ListScriptsResponseBodyDataScripts {
	s.UpdatedTime = &v
	return s
}

func (s *ListScriptsResponseBodyDataScripts) Validate() error {
	return dara.Validate(s)
}

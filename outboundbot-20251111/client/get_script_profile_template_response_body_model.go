// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptProfileTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetScriptProfileTemplateResponseBody
	GetCode() *string
	SetData(v *GetScriptProfileTemplateResponseBodyData) *GetScriptProfileTemplateResponseBody
	GetData() *GetScriptProfileTemplateResponseBodyData
	SetHttpStatusCode(v int32) *GetScriptProfileTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetScriptProfileTemplateResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetScriptProfileTemplateResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetScriptProfileTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScriptProfileTemplateResponseBody
	GetSuccess() *bool
}

type GetScriptProfileTemplateResponseBody struct {
	// 返回码
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data *GetScriptProfileTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Instance does not exist. Instance=392db13c-8901-4a25-b566-91d0d8114cec
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

func (s GetScriptProfileTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScriptProfileTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetScriptProfileTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScriptProfileTemplateResponseBody) GetData() *GetScriptProfileTemplateResponseBodyData {
	return s.Data
}

func (s *GetScriptProfileTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetScriptProfileTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScriptProfileTemplateResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetScriptProfileTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScriptProfileTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScriptProfileTemplateResponseBody) SetCode(v string) *GetScriptProfileTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetData(v *GetScriptProfileTemplateResponseBodyData) *GetScriptProfileTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetHttpStatusCode(v int32) *GetScriptProfileTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetMessage(v string) *GetScriptProfileTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetParams(v []*string) *GetScriptProfileTemplateResponseBody {
	s.Params = v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetRequestId(v string) *GetScriptProfileTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) SetSuccess(v bool) *GetScriptProfileTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScriptProfileTemplateResponseBodyData struct {
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
	// 作为调研专员，对服务总体满意度、服务亮点、改进建议、服务效率、员工态度、再次选择意愿进行依次询问，并采集信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 标签定义
	//
	// example:
	//
	// [{\\"name\\":\\"总体满意度\\",\\"description\\":\\"收集对服务的整体满意度的评价信息\\",\\"candidateValues\\":[\\"非常满意\\",\\"满意\\",\\"一般\\",\\"不满意\\",\\"非常不满意\\"]},{\\"name\\":\\"服务亮点\\",\\"description\\":\\"客户对于服务亮点或者满意的地方的描述\\",\\"candidateValues\\":[]},{\\"name\\":\\"改进建议\\",\\"description\\":\\"客户对于改进意见的描述\\",\\"candidateValues\\":[]},{\\"name\\":\\"服务效率\\",\\"description\\":\\"客户对于服务响应速度和服务完成的时效性的反馈\\",\\"candidateValues\\":[]},{\\"name\\":\\"员工态度\\",\\"description\\":\\"客户对于对于服务人员的专业度和态度的评价\\",\\"candidateValues\\":[]},{\\"name\\":\\"再次选择意愿\\",\\"description\\":\\"客户是否愿意再次选择\\",\\"candidateValues\\":[\\"是\\",\\"否\\"]}]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 名称
	//
	// example:
	//
	// 服务满意度调研
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// schema定义
	//
	// example:
	//
	// {\\"name\\":\\"李明\\",\\"gender\\":\\"男\\"}
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// 模板ID
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b59
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 更新时间，毫秒级时间戳
	//
	// example:
	//
	// 1735660800000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// 变量定义
	//
	// example:
	//
	// [{\\"name\\":\\"name\\",\\"description\\":\\"客户姓名\\"},{\\"name\\":\\"gender\\",\\"description\\":\\"客户性别\\"}]
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s GetScriptProfileTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScriptProfileTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScriptProfileTemplateResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetScriptProfileTemplateResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetScriptProfileTemplateResponseBodyData) GetLabels() *string {
	return s.Labels
}

func (s *GetScriptProfileTemplateResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetScriptProfileTemplateResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *GetScriptProfileTemplateResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetScriptProfileTemplateResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetScriptProfileTemplateResponseBodyData) GetVariables() *string {
	return s.Variables
}

func (s *GetScriptProfileTemplateResponseBodyData) SetCreatedTime(v int64) *GetScriptProfileTemplateResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetDescription(v string) *GetScriptProfileTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetLabels(v string) *GetScriptProfileTemplateResponseBodyData {
	s.Labels = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetName(v string) *GetScriptProfileTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetSchema(v string) *GetScriptProfileTemplateResponseBodyData {
	s.Schema = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetTemplateId(v string) *GetScriptProfileTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetUpdatedTime(v int64) *GetScriptProfileTemplateResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) SetVariables(v string) *GetScriptProfileTemplateResponseBodyData {
	s.Variables = &v
	return s
}

func (s *GetScriptProfileTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptProfileTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScriptProfileTemplatesResponseBody
	GetCode() *string
	SetData(v []*ListScriptProfileTemplatesResponseBodyData) *ListScriptProfileTemplatesResponseBody
	GetData() []*ListScriptProfileTemplatesResponseBodyData
	SetHttpStatusCode(v int32) *ListScriptProfileTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScriptProfileTemplatesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListScriptProfileTemplatesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListScriptProfileTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListScriptProfileTemplatesResponseBody
	GetSuccess() *bool
}

type ListScriptProfileTemplatesResponseBody struct {
	// 返回码
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据
	Data []*ListScriptProfileTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// Instance does not exist. Instance=ob-9876543210.
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

func (s ListScriptProfileTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptProfileTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptProfileTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScriptProfileTemplatesResponseBody) GetData() []*ListScriptProfileTemplatesResponseBodyData {
	return s.Data
}

func (s *ListScriptProfileTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScriptProfileTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScriptProfileTemplatesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListScriptProfileTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptProfileTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScriptProfileTemplatesResponseBody) SetCode(v string) *ListScriptProfileTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetData(v []*ListScriptProfileTemplatesResponseBodyData) *ListScriptProfileTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetHttpStatusCode(v int32) *ListScriptProfileTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetMessage(v string) *ListScriptProfileTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetParams(v []*string) *ListScriptProfileTemplatesResponseBody {
	s.Params = v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetRequestId(v string) *ListScriptProfileTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) SetSuccess(v bool) *ListScriptProfileTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptProfileTemplatesResponseBodyData struct {
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
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
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

func (s ListScriptProfileTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListScriptProfileTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetLabels() *string {
	return s.Labels
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListScriptProfileTemplatesResponseBodyData) GetVariables() *string {
	return s.Variables
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetCreatedTime(v int64) *ListScriptProfileTemplatesResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetDescription(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetLabels(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Labels = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetName(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetSchema(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Schema = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetTemplateId(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetUpdatedTime(v int64) *ListScriptProfileTemplatesResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) SetVariables(v string) *ListScriptProfileTemplatesResponseBodyData {
	s.Variables = &v
	return s
}

func (s *ListScriptProfileTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

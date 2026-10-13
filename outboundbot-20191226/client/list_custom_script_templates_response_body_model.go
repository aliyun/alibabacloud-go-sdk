// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomScriptTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCustomScriptTemplatesResponseBody
	GetCode() *string
	SetData(v *ListCustomScriptTemplatesResponseBodyData) *ListCustomScriptTemplatesResponseBody
	GetData() *ListCustomScriptTemplatesResponseBodyData
	SetHttpStatusCode(v int32) *ListCustomScriptTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListCustomScriptTemplatesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListCustomScriptTemplatesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListCustomScriptTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCustomScriptTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomScriptTemplatesResponseBody
	GetSuccess() *bool
}

type ListCustomScriptTemplatesResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCustomScriptTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 暂无使用
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 暂无使用
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCustomScriptTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomScriptTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomScriptTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCustomScriptTemplatesResponseBody) GetData() *ListCustomScriptTemplatesResponseBodyData {
	return s.Data
}

func (s *ListCustomScriptTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCustomScriptTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCustomScriptTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCustomScriptTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomScriptTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomScriptTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomScriptTemplatesResponseBody) SetCode(v string) *ListCustomScriptTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBody) SetData(v *ListCustomScriptTemplatesResponseBodyData) *ListCustomScriptTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomScriptTemplatesResponseBody) SetHttpStatusCode(v int32) *ListCustomScriptTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBody) SetMaxResults(v int32) *ListCustomScriptTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBody) SetMessage(v string) *ListCustomScriptTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBody) SetNextToken(v string) *ListCustomScriptTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBody) SetRequestId(v string) *ListCustomScriptTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBody) SetSuccess(v bool) *ListCustomScriptTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomScriptTemplatesResponseBodyData struct {
	List []*ListCustomScriptTemplatesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomScriptTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomScriptTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomScriptTemplatesResponseBodyData) GetList() []*ListCustomScriptTemplatesResponseBodyDataList {
	return s.List
}

func (s *ListCustomScriptTemplatesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomScriptTemplatesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomScriptTemplatesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCustomScriptTemplatesResponseBodyData) SetList(v []*ListCustomScriptTemplatesResponseBodyDataList) *ListCustomScriptTemplatesResponseBodyData {
	s.List = v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyData) SetPageNumber(v int32) *ListCustomScriptTemplatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyData) SetPageSize(v int32) *ListCustomScriptTemplatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyData) SetTotalCount(v int32) *ListCustomScriptTemplatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomScriptTemplatesResponseBodyDataList struct {
	// example:
	//
	// d31794e2a51f47d2901b4094d88311d7
	AgentProfileId *string `json:"AgentProfileId,omitempty" xml:"AgentProfileId,omitempty"`
	// example:
	//
	// 1741338619000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 这是一个大模型机器人
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// e49ad122-15a1-443a-a102-84a78a93be79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// [{\\"name\\":\\"是否送达\\",\\"description\\":\\"购买的家电是否已经送达\\",\\"valueList\\":\\"[\\\\\\"是\\\\\\",\\\\\\"否\\\\\\"]\\"},{\\"name\\":\\"预约上门时间\\",\\"description\\":\\"收集客户期望的上门安装时间\\",\\"valueList\\":\\"[]\\"}]
	LabelsJson *string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty"`
	// example:
	//
	// model_002
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 测试场景
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// \\n名称：安妮 \\n身份：游戏推广员\\n背景：通过电话向玩家宣传最新福利活动，确保每一位玩家都能及时掌握信息，享受游戏的乐趣。\\n技能：作为游戏客服，向用户推荐限时游戏福利活动，强调福利亮点并告知参与流程。\\n约束条件：你需要记住你是安妮，
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// {"name":"小x","gender":"男","age":18,"role":"游戏推广员","communicationStyle":["亲切"],"goals":"你好","background":"不是很好","openingPrompt":"你好，我是xxx"}
	PromptJson *string `json:"PromptJson,omitempty" xml:"PromptJson,omitempty"`
	// example:
	//
	// 43972417-0657-452a-81c2-c59d8245a9b2
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListCustomScriptTemplatesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListCustomScriptTemplatesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) GetAgentProfileId() *string {
	return s.AgentProfileId
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) GetLabelsJson() *string {
	return s.LabelsJson
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) GetModel() *string {
	return s.Model
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) GetPrompt() *string {
	return s.Prompt
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) GetPromptJson() *string {
	return s.PromptJson
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) SetAgentProfileId(v string) *ListCustomScriptTemplatesResponseBodyDataList {
	s.AgentProfileId = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) SetCreateTime(v string) *ListCustomScriptTemplatesResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) SetDescription(v string) *ListCustomScriptTemplatesResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) SetInstanceId(v string) *ListCustomScriptTemplatesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) SetLabelsJson(v string) *ListCustomScriptTemplatesResponseBodyDataList {
	s.LabelsJson = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) SetModel(v string) *ListCustomScriptTemplatesResponseBodyDataList {
	s.Model = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) SetName(v string) *ListCustomScriptTemplatesResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) SetPrompt(v string) *ListCustomScriptTemplatesResponseBodyDataList {
	s.Prompt = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) SetPromptJson(v string) *ListCustomScriptTemplatesResponseBodyDataList {
	s.PromptJson = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) SetScriptId(v string) *ListCustomScriptTemplatesResponseBodyDataList {
	s.ScriptId = &v
	return s
}

func (s *ListCustomScriptTemplatesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

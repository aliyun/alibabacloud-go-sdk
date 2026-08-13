// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArguments(v []*GetSkillResponseBodyArguments) *GetSkillResponseBody
	GetArguments() []*GetSkillResponseBodyArguments
	SetCode(v string) *GetSkillResponseBody
	GetCode() *string
	SetCreatedTime(v string) *GetSkillResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *GetSkillResponseBody
	GetDescription() *string
	SetDetailLogic(v string) *GetSkillResponseBody
	GetDetailLogic() *string
	SetDisplayName(v string) *GetSkillResponseBody
	GetDisplayName() *string
	SetExecuteMode(v string) *GetSkillResponseBody
	GetExecuteMode() *string
	SetGlobalAccess(v bool) *GetSkillResponseBody
	GetGlobalAccess() *bool
	SetHasDraftChanges(v bool) *GetSkillResponseBody
	GetHasDraftChanges() *bool
	SetInputConfig(v string) *GetSkillResponseBody
	GetInputConfig() *string
	SetInputConfigFormatted(v []map[string]interface{}) *GetSkillResponseBody
	GetInputConfigFormatted() []map[string]interface{}
	SetMessage(v string) *GetSkillResponseBody
	GetMessage() *string
	SetName(v string) *GetSkillResponseBody
	GetName() *string
	SetRequestId(v string) *GetSkillResponseBody
	GetRequestId() *string
	SetSkillCode(v string) *GetSkillResponseBody
	GetSkillCode() *string
	SetSkillFiles(v []map[string]interface{}) *GetSkillResponseBody
	GetSkillFiles() []map[string]interface{}
	SetSkillHubDefinitionId(v int64) *GetSkillResponseBody
	GetSkillHubDefinitionId() *int64
	SetSkillMdSummary(v string) *GetSkillResponseBody
	GetSkillMdSummary() *string
	SetSourceType(v string) *GetSkillResponseBody
	GetSourceType() *string
	SetStatus(v string) *GetSkillResponseBody
	GetStatus() *string
	SetTags(v []*string) *GetSkillResponseBody
	GetTags() []*string
	SetUpdatedTime(v string) *GetSkillResponseBody
	GetUpdatedTime() *string
	SetVersionCount(v int64) *GetSkillResponseBody
	GetVersionCount() *int64
	SetVersionNumber(v string) *GetSkillResponseBody
	GetVersionNumber() *string
}

type GetSkillResponseBody struct {
	Arguments []*GetSkillResponseBodyArguments `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 创建时间，ISO8601 格式
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// 技能描述（已 i18n 解析）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 技能详细逻辑
	//
	// example:
	//
	// string_value
	DetailLogic *string `json:"detailLogic,omitempty" xml:"detailLogic,omitempty"`
	// 展示名称
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 执行模式
	//
	// example:
	//
	// string_value
	ExecuteMode *string `json:"executeMode,omitempty" xml:"executeMode,omitempty"`
	// 是否全局可访问
	//
	// example:
	//
	// true
	GlobalAccess *bool `json:"globalAccess,omitempty" xml:"globalAccess,omitempty"`
	// 是否存在未发布的草稿修改
	//
	// example:
	//
	// true
	HasDraftChanges *bool `json:"hasDraftChanges,omitempty" xml:"hasDraftChanges,omitempty"`
	// 入参配置原文
	//
	// example:
	//
	// string_value
	InputConfig          *string                  `json:"inputConfig,omitempty" xml:"inputConfig,omitempty"`
	InputConfigFormatted []map[string]interface{} `json:"inputConfigFormatted,omitempty" xml:"inputConfigFormatted,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 技能编码（全局唯一）
	//
	// example:
	//
	// string_value
	SkillCode  *string                  `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	SkillFiles []map[string]interface{} `json:"skillFiles,omitempty" xml:"skillFiles,omitempty" type:"Repeated"`
	// 技能定义 ID
	//
	// example:
	//
	// 1
	SkillHubDefinitionId *int64 `json:"skillHubDefinitionId,omitempty" xml:"skillHubDefinitionId,omitempty"`
	// SKILL.md 简介（由 LLM 生成）
	//
	// example:
	//
	// string_value
	SkillMdSummary *string `json:"skillMdSummary,omitempty" xml:"skillMdSummary,omitempty"`
	// 来源类型: BUILTIN / CUSTOM
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 技能状态: ACTIVE / DRAFT
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// tags
	//
	// example:
	//
	// string_value
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// 修改时间，ISO8601 格式
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	UpdatedTime *string `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	// 版本总数
	//
	// example:
	//
	// 1
	VersionCount *int64 `json:"versionCount,omitempty" xml:"versionCount,omitempty"`
	// 版本号
	//
	// example:
	//
	// string_value
	VersionNumber *string `json:"versionNumber,omitempty" xml:"versionNumber,omitempty"`
}

func (s GetSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillResponseBody) GetArguments() []*GetSkillResponseBodyArguments {
	return s.Arguments
}

func (s *GetSkillResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSkillResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetSkillResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetSkillResponseBody) GetDetailLogic() *string {
	return s.DetailLogic
}

func (s *GetSkillResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetSkillResponseBody) GetExecuteMode() *string {
	return s.ExecuteMode
}

func (s *GetSkillResponseBody) GetGlobalAccess() *bool {
	return s.GlobalAccess
}

func (s *GetSkillResponseBody) GetHasDraftChanges() *bool {
	return s.HasDraftChanges
}

func (s *GetSkillResponseBody) GetInputConfig() *string {
	return s.InputConfig
}

func (s *GetSkillResponseBody) GetInputConfigFormatted() []map[string]interface{} {
	return s.InputConfigFormatted
}

func (s *GetSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSkillResponseBody) GetName() *string {
	return s.Name
}

func (s *GetSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillResponseBody) GetSkillCode() *string {
	return s.SkillCode
}

func (s *GetSkillResponseBody) GetSkillFiles() []map[string]interface{} {
	return s.SkillFiles
}

func (s *GetSkillResponseBody) GetSkillHubDefinitionId() *int64 {
	return s.SkillHubDefinitionId
}

func (s *GetSkillResponseBody) GetSkillMdSummary() *string {
	return s.SkillMdSummary
}

func (s *GetSkillResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetSkillResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSkillResponseBody) GetTags() []*string {
	return s.Tags
}

func (s *GetSkillResponseBody) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *GetSkillResponseBody) GetVersionCount() *int64 {
	return s.VersionCount
}

func (s *GetSkillResponseBody) GetVersionNumber() *string {
	return s.VersionNumber
}

func (s *GetSkillResponseBody) SetArguments(v []*GetSkillResponseBodyArguments) *GetSkillResponseBody {
	s.Arguments = v
	return s
}

func (s *GetSkillResponseBody) SetCode(v string) *GetSkillResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillResponseBody) SetCreatedTime(v string) *GetSkillResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetSkillResponseBody) SetDescription(v string) *GetSkillResponseBody {
	s.Description = &v
	return s
}

func (s *GetSkillResponseBody) SetDetailLogic(v string) *GetSkillResponseBody {
	s.DetailLogic = &v
	return s
}

func (s *GetSkillResponseBody) SetDisplayName(v string) *GetSkillResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetSkillResponseBody) SetExecuteMode(v string) *GetSkillResponseBody {
	s.ExecuteMode = &v
	return s
}

func (s *GetSkillResponseBody) SetGlobalAccess(v bool) *GetSkillResponseBody {
	s.GlobalAccess = &v
	return s
}

func (s *GetSkillResponseBody) SetHasDraftChanges(v bool) *GetSkillResponseBody {
	s.HasDraftChanges = &v
	return s
}

func (s *GetSkillResponseBody) SetInputConfig(v string) *GetSkillResponseBody {
	s.InputConfig = &v
	return s
}

func (s *GetSkillResponseBody) SetInputConfigFormatted(v []map[string]interface{}) *GetSkillResponseBody {
	s.InputConfigFormatted = v
	return s
}

func (s *GetSkillResponseBody) SetMessage(v string) *GetSkillResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillResponseBody) SetName(v string) *GetSkillResponseBody {
	s.Name = &v
	return s
}

func (s *GetSkillResponseBody) SetRequestId(v string) *GetSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillResponseBody) SetSkillCode(v string) *GetSkillResponseBody {
	s.SkillCode = &v
	return s
}

func (s *GetSkillResponseBody) SetSkillFiles(v []map[string]interface{}) *GetSkillResponseBody {
	s.SkillFiles = v
	return s
}

func (s *GetSkillResponseBody) SetSkillHubDefinitionId(v int64) *GetSkillResponseBody {
	s.SkillHubDefinitionId = &v
	return s
}

func (s *GetSkillResponseBody) SetSkillMdSummary(v string) *GetSkillResponseBody {
	s.SkillMdSummary = &v
	return s
}

func (s *GetSkillResponseBody) SetSourceType(v string) *GetSkillResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetSkillResponseBody) SetStatus(v string) *GetSkillResponseBody {
	s.Status = &v
	return s
}

func (s *GetSkillResponseBody) SetTags(v []*string) *GetSkillResponseBody {
	s.Tags = v
	return s
}

func (s *GetSkillResponseBody) SetUpdatedTime(v string) *GetSkillResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *GetSkillResponseBody) SetVersionCount(v int64) *GetSkillResponseBody {
	s.VersionCount = &v
	return s
}

func (s *GetSkillResponseBody) SetVersionNumber(v string) *GetSkillResponseBody {
	s.VersionNumber = &v
	return s
}

func (s *GetSkillResponseBody) Validate() error {
	if s.Arguments != nil {
		for _, item := range s.Arguments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSkillResponseBodyArguments struct {
	// 默认值
	//
	// example:
	//
	// string_value
	Default *string `json:"default,omitempty" xml:"default,omitempty"`
	// 参数说明
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// enum
	//
	// example:
	//
	// string_value
	Enum []*string `json:"enum,omitempty" xml:"enum,omitempty" type:"Repeated"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否必填
	//
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// 参数类型: string / number / boolean / array
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetSkillResponseBodyArguments) String() string {
	return dara.Prettify(s)
}

func (s GetSkillResponseBodyArguments) GoString() string {
	return s.String()
}

func (s *GetSkillResponseBodyArguments) GetDefault() *string {
	return s.Default
}

func (s *GetSkillResponseBodyArguments) GetDescription() *string {
	return s.Description
}

func (s *GetSkillResponseBodyArguments) GetEnum() []*string {
	return s.Enum
}

func (s *GetSkillResponseBodyArguments) GetName() *string {
	return s.Name
}

func (s *GetSkillResponseBodyArguments) GetRequired() *bool {
	return s.Required
}

func (s *GetSkillResponseBodyArguments) GetType() *string {
	return s.Type
}

func (s *GetSkillResponseBodyArguments) SetDefault(v string) *GetSkillResponseBodyArguments {
	s.Default = &v
	return s
}

func (s *GetSkillResponseBodyArguments) SetDescription(v string) *GetSkillResponseBodyArguments {
	s.Description = &v
	return s
}

func (s *GetSkillResponseBodyArguments) SetEnum(v []*string) *GetSkillResponseBodyArguments {
	s.Enum = v
	return s
}

func (s *GetSkillResponseBodyArguments) SetName(v string) *GetSkillResponseBodyArguments {
	s.Name = &v
	return s
}

func (s *GetSkillResponseBodyArguments) SetRequired(v bool) *GetSkillResponseBodyArguments {
	s.Required = &v
	return s
}

func (s *GetSkillResponseBodyArguments) SetType(v string) *GetSkillResponseBodyArguments {
	s.Type = &v
	return s
}

func (s *GetSkillResponseBodyArguments) Validate() error {
	return dara.Validate(s)
}

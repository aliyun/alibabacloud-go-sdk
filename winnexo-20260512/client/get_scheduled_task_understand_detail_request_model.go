// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskUnderstandDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetScheduledTaskUnderstandDetailRequest
	GetCollaborationGroupId() *string
	SetDigitalEmployeeName(v []*string) *GetScheduledTaskUnderstandDetailRequest
	GetDigitalEmployeeName() []*string
	SetSegments(v []*GetScheduledTaskUnderstandDetailRequestSegments) *GetScheduledTaskUnderstandDetailRequest
	GetSegments() []*GetScheduledTaskUnderstandDetailRequestSegments
	SetTenantId(v string) *GetScheduledTaskUnderstandDetailRequest
	GetTenantId() *string
	SetUserInput(v string) *GetScheduledTaskUnderstandDetailRequest
	GetUserInput() *string
}

type GetScheduledTaskUnderstandDetailRequest struct {
	// 所属协作群组 ID（如 cg_101）；群任务理解时传入（调用者需为有效群成员），候选技能额外并入群绑定技能
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 数字员工名称列表，用于过滤可用技能；必传（传空列表表示仅用租户 global 技能）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	DigitalEmployeeName []*string                                          `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty" type:"Repeated"`
	Segments            []*GetScheduledTaskUnderstandDetailRequestSegments `json:"segments,omitempty" xml:"segments,omitempty" type:"Repeated"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 自然语言任务描述
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	UserInput *string `json:"userInput,omitempty" xml:"userInput,omitempty"`
}

func (s GetScheduledTaskUnderstandDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskUnderstandDetailRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskUnderstandDetailRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetScheduledTaskUnderstandDetailRequest) GetDigitalEmployeeName() []*string {
	return s.DigitalEmployeeName
}

func (s *GetScheduledTaskUnderstandDetailRequest) GetSegments() []*GetScheduledTaskUnderstandDetailRequestSegments {
	return s.Segments
}

func (s *GetScheduledTaskUnderstandDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScheduledTaskUnderstandDetailRequest) GetUserInput() *string {
	return s.UserInput
}

func (s *GetScheduledTaskUnderstandDetailRequest) SetCollaborationGroupId(v string) *GetScheduledTaskUnderstandDetailRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequest) SetDigitalEmployeeName(v []*string) *GetScheduledTaskUnderstandDetailRequest {
	s.DigitalEmployeeName = v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequest) SetSegments(v []*GetScheduledTaskUnderstandDetailRequestSegments) *GetScheduledTaskUnderstandDetailRequest {
	s.Segments = v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequest) SetTenantId(v string) *GetScheduledTaskUnderstandDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequest) SetUserInput(v string) *GetScheduledTaskUnderstandDetailRequest {
	s.UserInput = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequest) Validate() error {
	if s.Segments != nil {
		for _, item := range s.Segments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScheduledTaskUnderstandDetailRequestSegments struct {
	// 文本内容，type=text 时必填
	//
	// example:
	//
	// 示例内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 功能开关，type=web_search 时可选
	//
	// example:
	//
	// true
	Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 对象 ID，type=mention 时有值
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 对象类型如 customer，type=mention 时有值
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 技能编码，type=skill 时有值
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// 元素类型：text|web_search|mention|skill
	//
	// This parameter is required.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetScheduledTaskUnderstandDetailRequestSegments) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskUnderstandDetailRequestSegments) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) GetContent() *string {
	return s.Content
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) GetEnabled() *string {
	return s.Enabled
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) GetName() *string {
	return s.Name
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) GetSkillCode() *string {
	return s.SkillCode
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) GetType() *string {
	return s.Type
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) SetContent(v string) *GetScheduledTaskUnderstandDetailRequestSegments {
	s.Content = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) SetEnabled(v string) *GetScheduledTaskUnderstandDetailRequestSegments {
	s.Enabled = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) SetName(v string) *GetScheduledTaskUnderstandDetailRequestSegments {
	s.Name = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) SetObjectId(v string) *GetScheduledTaskUnderstandDetailRequestSegments {
	s.ObjectId = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) SetObjectType(v string) *GetScheduledTaskUnderstandDetailRequestSegments {
	s.ObjectType = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) SetSkillCode(v string) *GetScheduledTaskUnderstandDetailRequestSegments {
	s.SkillCode = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) SetType(v string) *GetScheduledTaskUnderstandDetailRequestSegments {
	s.Type = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailRequestSegments) Validate() error {
	return dara.Validate(s)
}

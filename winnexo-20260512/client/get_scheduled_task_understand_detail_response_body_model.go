// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskUnderstandDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetScheduledTaskUnderstandDetailResponseBody
	GetCode() *string
	SetMessage(v string) *GetScheduledTaskUnderstandDetailResponseBody
	GetMessage() *string
	SetRelatedObjects(v []*GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) *GetScheduledTaskUnderstandDetailResponseBody
	GetRelatedObjects() []*GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects
	SetRelatedSemantics(v []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics) *GetScheduledTaskUnderstandDetailResponseBody
	GetRelatedSemantics() []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics
	SetRelatedSkills(v []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) *GetScheduledTaskUnderstandDetailResponseBody
	GetRelatedSkills() []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills
	SetRequestId(v string) *GetScheduledTaskUnderstandDetailResponseBody
	GetRequestId() *string
	SetTaskUnderstand(v string) *GetScheduledTaskUnderstandDetailResponseBody
	GetTaskUnderstand() *string
}

type GetScheduledTaskUnderstandDetailResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message          *string                                                         `json:"message,omitempty" xml:"message,omitempty"`
	RelatedObjects   []*GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects   `json:"relatedObjects,omitempty" xml:"relatedObjects,omitempty" type:"Repeated"`
	RelatedSemantics []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics `json:"relatedSemantics,omitempty" xml:"relatedSemantics,omitempty" type:"Repeated"`
	RelatedSkills    []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills    `json:"relatedSkills,omitempty" xml:"relatedSkills,omitempty" type:"Repeated"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 润色后的任务理解
	//
	// example:
	//
	// string_value
	TaskUnderstand *string `json:"taskUnderstand,omitempty" xml:"taskUnderstand,omitempty"`
}

func (s GetScheduledTaskUnderstandDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskUnderstandDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) GetRelatedObjects() []*GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects {
	return s.RelatedObjects
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) GetRelatedSemantics() []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics {
	return s.RelatedSemantics
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) GetRelatedSkills() []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills {
	return s.RelatedSkills
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) GetTaskUnderstand() *string {
	return s.TaskUnderstand
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) SetCode(v string) *GetScheduledTaskUnderstandDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) SetMessage(v string) *GetScheduledTaskUnderstandDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) SetRelatedObjects(v []*GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) *GetScheduledTaskUnderstandDetailResponseBody {
	s.RelatedObjects = v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) SetRelatedSemantics(v []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics) *GetScheduledTaskUnderstandDetailResponseBody {
	s.RelatedSemantics = v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) SetRelatedSkills(v []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) *GetScheduledTaskUnderstandDetailResponseBody {
	s.RelatedSkills = v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) SetRequestId(v string) *GetScheduledTaskUnderstandDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) SetTaskUnderstand(v string) *GetScheduledTaskUnderstandDetailResponseBody {
	s.TaskUnderstand = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBody) Validate() error {
	if s.RelatedObjects != nil {
		for _, item := range s.RelatedObjects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedSemantics != nil {
		for _, item := range s.RelatedSemantics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedSkills != nil {
		for _, item := range s.RelatedSkills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects struct {
	// 提及类型
	//
	// example:
	//
	// string_value
	MentionType *string `json:"mentionType,omitempty" xml:"mentionType,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 对象 ID
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 对象类型
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) GetMentionType() *string {
	return s.MentionType
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) GetName() *string {
	return s.Name
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) SetMentionType(v string) *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects {
	s.MentionType = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) SetName(v string) *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects {
	s.Name = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) SetObjectId(v string) *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects {
	s.ObjectId = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) SetObjectType(v string) *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects {
	s.ObjectType = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects) Validate() error {
	return dara.Validate(s)
}

type GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics struct {
	// 语义属性（JSON 字符串），用于语义检索时过滤
	//
	// example:
	//
	// {"level": "VIP"}
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 语义实体名，如客户/机会
	//
	// example:
	//
	// customer
	Entity *string `json:"entity,omitempty" xml:"entity,omitempty"`
}

func (s GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics) GetAttributes() *string {
	return s.Attributes
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics) GetEntity() *string {
	return s.Entity
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics) SetAttributes(v string) *GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics {
	s.Attributes = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics) SetEntity(v string) *GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics {
	s.Entity = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics) Validate() error {
	return dara.Validate(s)
}

type GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills struct {
	// 技能展示名称
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 技能代码
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// sourceIds
	//
	// example:
	//
	// string_value
	SourceIds []*string `json:"sourceIds,omitempty" xml:"sourceIds,omitempty" type:"Repeated"`
}

func (s GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) GetName() *string {
	return s.Name
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) GetSkillCode() *string {
	return s.SkillCode
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) SetDisplayName(v string) *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills {
	s.DisplayName = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) SetName(v string) *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills {
	s.Name = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) SetSkillCode(v string) *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills {
	s.SkillCode = &v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) SetSourceIds(v []*string) *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills {
	s.SourceIds = v
	return s
}

func (s *GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills) Validate() error {
	return dara.Validate(s)
}

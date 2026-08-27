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
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The related objects.
	RelatedObjects []*GetScheduledTaskUnderstandDetailResponseBodyRelatedObjects `json:"relatedObjects,omitempty" xml:"relatedObjects,omitempty" type:"Repeated"`
	// The related semantics.
	RelatedSemantics []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSemantics `json:"relatedSemantics,omitempty" xml:"relatedSemantics,omitempty" type:"Repeated"`
	// The related skills.
	RelatedSkills []*GetScheduledTaskUnderstandDetailResponseBodyRelatedSkills `json:"relatedSkills,omitempty" xml:"relatedSkills,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The task understanding description polished by the LLM.
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
	// The mention type, such as objects.
	//
	// example:
	//
	// string_value
	MentionType *string `json:"mentionType,omitempty" xml:"mentionType,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The object ID. Pass the project task ID.
	//
	// - For internal enterprise applications, this is the taskId obtained by calling the [Create a project task](https://open.dingtalk.com/document/orgapp-server/create-a-project-task) operation.
	//
	// - For third-party enterprise applications, this is the taskId obtained by calling the [Create a project task](https://open.dingtalk.com/document/isvapp-server/create-a-project-task) operation.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
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
	// The information type.
	//
	// example:
	//
	// {"level": "VIP"}
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// The semantic entity name, such as customer or opportunity.
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
	// The display name of the MCP service.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The skill code.
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

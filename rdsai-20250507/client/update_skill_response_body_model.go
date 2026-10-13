// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *UpdateSkillResponseBody
	GetAction() *string
	SetBundleHash(v string) *UpdateSkillResponseBody
	GetBundleHash() *string
	SetContent(v map[string]interface{}) *UpdateSkillResponseBody
	GetContent() map[string]interface{}
	SetDbtypes(v []*string) *UpdateSkillResponseBody
	GetDbtypes() []*string
	SetDescription(v string) *UpdateSkillResponseBody
	GetDescription() *string
	SetId(v string) *UpdateSkillResponseBody
	GetId() *string
	SetName(v string) *UpdateSkillResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateSkillResponseBody
	GetRequestId() *string
	SetSkillId(v string) *UpdateSkillResponseBody
	GetSkillId() *string
	SetSkillType(v string) *UpdateSkillResponseBody
	GetSkillType() *string
	SetUpdatedAt(v string) *UpdateSkillResponseBody
	GetUpdatedAt() *string
	SetVersionId(v string) *UpdateSkillResponseBody
	GetVersionId() *string
}

type UpdateSkillResponseBody struct {
	// The update action performed.
	//
	// example:
	//
	// ADD_VERSION
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The content digest of the Skill package.
	//
	// example:
	//
	// 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
	BundleHash *string `json:"BundleHash,omitempty" xml:"BundleHash,omitempty"`
	// The content.
	//
	// example:
	//
	// {"MySQL": "MySQL optimization guide...","PostgreSQL": "PostgreSQL optimization guide..."}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The list of database types.
	Dbtypes []*string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
	// The description of the Skill. The description can be up to 1000 characters in length.
	//
	// example:
	//
	// SQL Review Expert: Comprehensively reviews SQL for security, performance, and compliance, identifies risks, and provides optimization suggestions. Activated immediately when a user submits SQL or asks "SQL review", "SQL Review", "any risks?", or "how to optimize"
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the Skill.
	//
	// example:
	//
	// d1b7d639-f34e-44c7-8231-987da14d****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the Skill. The name can contain only lowercase letters, digits, and hyphens.
	//
	// example:
	//
	// sql-review
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the updated Skill.
	//
	// example:
	//
	// skill-example
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The Skill type.
	//
	// example:
	//
	// user
	SkillType *string `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
	// The time when the Skill was updated.
	//
	// example:
	//
	// 2026-02-04T21:14:45Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// The ID of the Skill version created in this update.
	//
	// example:
	//
	// version-example
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillResponseBody) GetAction() *string {
	return s.Action
}

func (s *UpdateSkillResponseBody) GetBundleHash() *string {
	return s.BundleHash
}

func (s *UpdateSkillResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *UpdateSkillResponseBody) GetDbtypes() []*string {
	return s.Dbtypes
}

func (s *UpdateSkillResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateSkillResponseBody) GetId() *string {
	return s.Id
}

func (s *UpdateSkillResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSkillResponseBody) GetSkillId() *string {
	return s.SkillId
}

func (s *UpdateSkillResponseBody) GetSkillType() *string {
	return s.SkillType
}

func (s *UpdateSkillResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateSkillResponseBody) GetVersionId() *string {
	return s.VersionId
}

func (s *UpdateSkillResponseBody) SetAction(v string) *UpdateSkillResponseBody {
	s.Action = &v
	return s
}

func (s *UpdateSkillResponseBody) SetBundleHash(v string) *UpdateSkillResponseBody {
	s.BundleHash = &v
	return s
}

func (s *UpdateSkillResponseBody) SetContent(v map[string]interface{}) *UpdateSkillResponseBody {
	s.Content = v
	return s
}

func (s *UpdateSkillResponseBody) SetDbtypes(v []*string) *UpdateSkillResponseBody {
	s.Dbtypes = v
	return s
}

func (s *UpdateSkillResponseBody) SetDescription(v string) *UpdateSkillResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateSkillResponseBody) SetId(v string) *UpdateSkillResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateSkillResponseBody) SetName(v string) *UpdateSkillResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateSkillResponseBody) SetRequestId(v string) *UpdateSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillResponseBody) SetSkillId(v string) *UpdateSkillResponseBody {
	s.SkillId = &v
	return s
}

func (s *UpdateSkillResponseBody) SetSkillType(v string) *UpdateSkillResponseBody {
	s.SkillType = &v
	return s
}

func (s *UpdateSkillResponseBody) SetUpdatedAt(v string) *UpdateSkillResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateSkillResponseBody) SetVersionId(v string) *UpdateSkillResponseBody {
	s.VersionId = &v
	return s
}

func (s *UpdateSkillResponseBody) Validate() error {
	return dara.Validate(s)
}

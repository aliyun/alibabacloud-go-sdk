// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *GetSkillResponseBody
	GetContent() map[string]interface{}
	SetCreatedAt(v string) *GetSkillResponseBody
	GetCreatedAt() *string
	SetDbtypes(v []*string) *GetSkillResponseBody
	GetDbtypes() []*string
	SetDescription(v string) *GetSkillResponseBody
	GetDescription() *string
	SetId(v string) *GetSkillResponseBody
	GetId() *string
	SetName(v string) *GetSkillResponseBody
	GetName() *string
	SetRequestId(v string) *GetSkillResponseBody
	GetRequestId() *string
	SetSkillType(v string) *GetSkillResponseBody
	GetSkillType() *string
	SetUpdatedAt(v string) *GetSkillResponseBody
	GetUpdatedAt() *string
}

type GetSkillResponseBody struct {
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2025-06-04T02:25:43Z
	CreatedAt   *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Dbtypes     []*string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// d1b7d639-f34e-44c7-8231-987da14d****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// sql-optimization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// user
	SkillType *string `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
	// example:
	//
	// 2026-02-04T21:14:45Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s GetSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *GetSkillResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetSkillResponseBody) GetDbtypes() []*string {
	return s.Dbtypes
}

func (s *GetSkillResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetSkillResponseBody) GetId() *string {
	return s.Id
}

func (s *GetSkillResponseBody) GetName() *string {
	return s.Name
}

func (s *GetSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillResponseBody) GetSkillType() *string {
	return s.SkillType
}

func (s *GetSkillResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetSkillResponseBody) SetContent(v map[string]interface{}) *GetSkillResponseBody {
	s.Content = v
	return s
}

func (s *GetSkillResponseBody) SetCreatedAt(v string) *GetSkillResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetSkillResponseBody) SetDbtypes(v []*string) *GetSkillResponseBody {
	s.Dbtypes = v
	return s
}

func (s *GetSkillResponseBody) SetDescription(v string) *GetSkillResponseBody {
	s.Description = &v
	return s
}

func (s *GetSkillResponseBody) SetId(v string) *GetSkillResponseBody {
	s.Id = &v
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

func (s *GetSkillResponseBody) SetSkillType(v string) *GetSkillResponseBody {
	s.SkillType = &v
	return s
}

func (s *GetSkillResponseBody) SetUpdatedAt(v string) *GetSkillResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetSkillResponseBody) Validate() error {
	return dara.Validate(s)
}

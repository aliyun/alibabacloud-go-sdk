// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetSkillResponseBody
	GetCreateTime() *string
	SetRequestId(v string) *GetSkillResponseBody
	GetRequestId() *string
	SetSkillDescription(v string) *GetSkillResponseBody
	GetSkillDescription() *string
	SetSkillId(v string) *GetSkillResponseBody
	GetSkillId() *string
	SetSkillLabels(v []*string) *GetSkillResponseBody
	GetSkillLabels() []*string
	SetSkillName(v string) *GetSkillResponseBody
	GetSkillName() *string
	SetSkillSpaceId(v string) *GetSkillResponseBody
	GetSkillSpaceId() *string
	SetUpdateTime(v string) *GetSkillResponseBody
	GetUpdateTime() *string
}

type GetSkillResponseBody struct {
	// The time the Skill was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The description of the Skill.
	//
	// example:
	//
	// 11111111
	SkillDescription *string `json:"SkillDescription,omitempty" xml:"SkillDescription,omitempty"`
	// The ID of the Skill.
	//
	// example:
	//
	// s-04zzrgosj6xd11yah
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The labels of the Skill.
	SkillLabels []*string `json:"SkillLabels,omitempty" xml:"SkillLabels,omitempty" type:"Repeated"`
	// The name of the Skill.
	//
	// example:
	//
	// skill-hello
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The ID of the SkillSpace to which the Skill belongs.
	//
	// example:
	//
	// ss-111111111
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
	// The time the Skill was last updated.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillResponseBody) GetSkillDescription() *string {
	return s.SkillDescription
}

func (s *GetSkillResponseBody) GetSkillId() *string {
	return s.SkillId
}

func (s *GetSkillResponseBody) GetSkillLabels() []*string {
	return s.SkillLabels
}

func (s *GetSkillResponseBody) GetSkillName() *string {
	return s.SkillName
}

func (s *GetSkillResponseBody) GetSkillSpaceId() *string {
	return s.SkillSpaceId
}

func (s *GetSkillResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetSkillResponseBody) SetCreateTime(v string) *GetSkillResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSkillResponseBody) SetRequestId(v string) *GetSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillResponseBody) SetSkillDescription(v string) *GetSkillResponseBody {
	s.SkillDescription = &v
	return s
}

func (s *GetSkillResponseBody) SetSkillId(v string) *GetSkillResponseBody {
	s.SkillId = &v
	return s
}

func (s *GetSkillResponseBody) SetSkillLabels(v []*string) *GetSkillResponseBody {
	s.SkillLabels = v
	return s
}

func (s *GetSkillResponseBody) SetSkillName(v string) *GetSkillResponseBody {
	s.SkillName = &v
	return s
}

func (s *GetSkillResponseBody) SetSkillSpaceId(v string) *GetSkillResponseBody {
	s.SkillSpaceId = &v
	return s
}

func (s *GetSkillResponseBody) SetUpdateTime(v string) *GetSkillResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetSkillResponseBody) Validate() error {
	return dara.Validate(s)
}

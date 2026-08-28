// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillVersionDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSkillVersionDetailResponseBodyData) *GetSkillVersionDetailResponseBody
	GetData() *GetSkillVersionDetailResponseBodyData
	SetRequestId(v string) *GetSkillVersionDetailResponseBody
	GetRequestId() *string
}

type GetSkillVersionDetailResponseBody struct {
	// The response data.
	Data *GetSkillVersionDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSkillVersionDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillVersionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillVersionDetailResponseBody) GetData() *GetSkillVersionDetailResponseBodyData {
	return s.Data
}

func (s *GetSkillVersionDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillVersionDetailResponseBody) SetData(v *GetSkillVersionDetailResponseBodyData) *GetSkillVersionDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillVersionDetailResponseBody) SetRequestId(v string) *GetSkillVersionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillVersionDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillVersionDetailResponseBodyData struct {
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name.
	//
	// example:
	//
	// skill-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource mapping (key is the resource name).
	Resource map[string]*DataResourceValue `json:"resource,omitempty" xml:"resource,omitempty"`
	// The Skill card content (SKILL.md).
	//
	// example:
	//
	// # Sample Skill
	//
	// Used to demonstrate Skill document content
	SkillMd *string `json:"skillMd,omitempty" xml:"skillMd,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetSkillVersionDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillVersionDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillVersionDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetSkillVersionDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetSkillVersionDetailResponseBodyData) GetResource() map[string]*DataResourceValue {
	return s.Resource
}

func (s *GetSkillVersionDetailResponseBodyData) GetSkillMd() *string {
	return s.SkillMd
}

func (s *GetSkillVersionDetailResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetSkillVersionDetailResponseBodyData) SetDescription(v string) *GetSkillVersionDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetSkillVersionDetailResponseBodyData) SetName(v string) *GetSkillVersionDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSkillVersionDetailResponseBodyData) SetResource(v map[string]*DataResourceValue) *GetSkillVersionDetailResponseBodyData {
	s.Resource = v
	return s
}

func (s *GetSkillVersionDetailResponseBodyData) SetSkillMd(v string) *GetSkillVersionDetailResponseBodyData {
	s.SkillMd = &v
	return s
}

func (s *GetSkillVersionDetailResponseBodyData) SetWorkspaceId(v string) *GetSkillVersionDetailResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetSkillVersionDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

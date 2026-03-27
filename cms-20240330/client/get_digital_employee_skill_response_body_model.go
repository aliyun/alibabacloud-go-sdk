// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetDigitalEmployeeSkillResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetDigitalEmployeeSkillResponseBody
	GetDescription() *string
	SetDisplayName(v string) *GetDigitalEmployeeSkillResponseBody
	GetDisplayName() *string
	SetEnable(v bool) *GetDigitalEmployeeSkillResponseBody
	GetEnable() *bool
	SetFiles(v []*GetDigitalEmployeeSkillResponseBodyFiles) *GetDigitalEmployeeSkillResponseBody
	GetFiles() []*GetDigitalEmployeeSkillResponseBodyFiles
	SetRemark(v string) *GetDigitalEmployeeSkillResponseBody
	GetRemark() *string
	SetRequestId(v string) *GetDigitalEmployeeSkillResponseBody
	GetRequestId() *string
	SetSkillName(v string) *GetDigitalEmployeeSkillResponseBody
	GetSkillName() *string
	SetUpdateTime(v string) *GetDigitalEmployeeSkillResponseBody
	GetUpdateTime() *string
	SetVersion(v string) *GetDigitalEmployeeSkillResponseBody
	GetVersion() *string
}

type GetDigitalEmployeeSkillResponseBody struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-02-06T14:09:11Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// true
	Enable *bool                                       `json:"enable,omitempty" xml:"enable,omitempty"`
	Files  []*GetDigitalEmployeeSkillResponseBodyFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// example:
	//
	// remark
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// test
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-02-06T14:09:11Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// 1770386951147366810
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDigitalEmployeeSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeSkillResponseBody) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeSkillResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDigitalEmployeeSkillResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetDigitalEmployeeSkillResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetDigitalEmployeeSkillResponseBody) GetEnable() *bool {
	return s.Enable
}

func (s *GetDigitalEmployeeSkillResponseBody) GetFiles() []*GetDigitalEmployeeSkillResponseBodyFiles {
	return s.Files
}

func (s *GetDigitalEmployeeSkillResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *GetDigitalEmployeeSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDigitalEmployeeSkillResponseBody) GetSkillName() *string {
	return s.SkillName
}

func (s *GetDigitalEmployeeSkillResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetDigitalEmployeeSkillResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetDigitalEmployeeSkillResponseBody) SetCreateTime(v string) *GetDigitalEmployeeSkillResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBody) SetDescription(v string) *GetDigitalEmployeeSkillResponseBody {
	s.Description = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBody) SetDisplayName(v string) *GetDigitalEmployeeSkillResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBody) SetEnable(v bool) *GetDigitalEmployeeSkillResponseBody {
	s.Enable = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBody) SetFiles(v []*GetDigitalEmployeeSkillResponseBodyFiles) *GetDigitalEmployeeSkillResponseBody {
	s.Files = v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBody) SetRemark(v string) *GetDigitalEmployeeSkillResponseBody {
	s.Remark = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBody) SetRequestId(v string) *GetDigitalEmployeeSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBody) SetSkillName(v string) *GetDigitalEmployeeSkillResponseBody {
	s.SkillName = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBody) SetUpdateTime(v string) *GetDigitalEmployeeSkillResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBody) SetVersion(v string) *GetDigitalEmployeeSkillResponseBody {
	s.Version = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBody) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDigitalEmployeeSkillResponseBodyFiles struct {
	// example:
	//
	// ---
	//
	// name: skill
	//
	// description: description
	//
	// ---
	//
	// # skill
	//
	// skill test
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// SKILL.md
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetDigitalEmployeeSkillResponseBodyFiles) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeSkillResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeSkillResponseBodyFiles) GetContent() *string {
	return s.Content
}

func (s *GetDigitalEmployeeSkillResponseBodyFiles) GetName() *string {
	return s.Name
}

func (s *GetDigitalEmployeeSkillResponseBodyFiles) SetContent(v string) *GetDigitalEmployeeSkillResponseBodyFiles {
	s.Content = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBodyFiles) SetName(v string) *GetDigitalEmployeeSkillResponseBodyFiles {
	s.Name = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponseBodyFiles) Validate() error {
	return dara.Validate(s)
}

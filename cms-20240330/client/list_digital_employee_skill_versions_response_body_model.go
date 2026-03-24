// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDigitalEmployeeSkillVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDigitalEmployeeSkillVersionsResponseBodyData) *ListDigitalEmployeeSkillVersionsResponseBody
	GetData() []*ListDigitalEmployeeSkillVersionsResponseBodyData
	SetRequestId(v string) *ListDigitalEmployeeSkillVersionsResponseBody
	GetRequestId() *string
}

type ListDigitalEmployeeSkillVersionsResponseBody struct {
	// The historical versions.
	Data []*ListDigitalEmployeeSkillVersionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListDigitalEmployeeSkillVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeeSkillVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeeSkillVersionsResponseBody) GetData() []*ListDigitalEmployeeSkillVersionsResponseBodyData {
	return s.Data
}

func (s *ListDigitalEmployeeSkillVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDigitalEmployeeSkillVersionsResponseBody) SetData(v []*ListDigitalEmployeeSkillVersionsResponseBodyData) *ListDigitalEmployeeSkillVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponseBody) SetRequestId(v string) *ListDigitalEmployeeSkillVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDigitalEmployeeSkillVersionsResponseBodyData struct {
	// The time when the skill was created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-02-06T14:09:11Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name.
	//
	// example:
	//
	// test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Indicates whether the skill is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The remarks.
	//
	// example:
	//
	// remark
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The name of the skill.
	//
	// example:
	//
	// test
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The time when the skill was last updated.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-02-06T14:09:11Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1770386951147366810
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListDigitalEmployeeSkillVersionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeeSkillVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) GetSkillName() *string {
	return s.SkillName
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) SetCreateTime(v string) *ListDigitalEmployeeSkillVersionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) SetDescription(v string) *ListDigitalEmployeeSkillVersionsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) SetDisplayName(v string) *ListDigitalEmployeeSkillVersionsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) SetEnable(v bool) *ListDigitalEmployeeSkillVersionsResponseBodyData {
	s.Enable = &v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) SetRemark(v string) *ListDigitalEmployeeSkillVersionsResponseBodyData {
	s.Remark = &v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) SetSkillName(v string) *ListDigitalEmployeeSkillVersionsResponseBodyData {
	s.SkillName = &v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) SetUpdateTime(v string) *ListDigitalEmployeeSkillVersionsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) SetVersion(v string) *ListDigitalEmployeeSkillVersionsResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

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
	Data      *GetSkillVersionDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Description *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                       `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId *string                       `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Resource    map[string]*DataResourceValue `json:"Resource,omitempty" xml:"Resource,omitempty"`
	SkillMd     *string                       `json:"SkillMd,omitempty" xml:"SkillMd,omitempty"`
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

func (s *GetSkillVersionDetailResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetSkillVersionDetailResponseBodyData) GetResource() map[string]*DataResourceValue {
	return s.Resource
}

func (s *GetSkillVersionDetailResponseBodyData) GetSkillMd() *string {
	return s.SkillMd
}

func (s *GetSkillVersionDetailResponseBodyData) SetDescription(v string) *GetSkillVersionDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetSkillVersionDetailResponseBodyData) SetName(v string) *GetSkillVersionDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSkillVersionDetailResponseBodyData) SetNamespaceId(v string) *GetSkillVersionDetailResponseBodyData {
	s.NamespaceId = &v
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

func (s *GetSkillVersionDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

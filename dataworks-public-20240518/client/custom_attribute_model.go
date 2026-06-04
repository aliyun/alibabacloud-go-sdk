// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomAttribute interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CustomAttribute
	GetComment() *string
	SetCreateTime(v int64) *CustomAttribute
	GetCreateTime() *int64
	SetDisplayEnabled(v bool) *CustomAttribute
	GetDisplayEnabled() *bool
	SetDisplayName(v string) *CustomAttribute
	GetDisplayName() *string
	SetEntityTypes(v []*string) *CustomAttribute
	GetEntityTypes() []*string
	SetId(v string) *CustomAttribute
	GetId() *string
	SetModifyTime(v int64) *CustomAttribute
	GetModifyTime() *int64
	SetSearchFilterEnabled(v bool) *CustomAttribute
	GetSearchFilterEnabled() *bool
	SetType(v string) *CustomAttribute
	GetType() *string
	SetValueEnums(v []*string) *CustomAttribute
	GetValueEnums() []*string
}

type CustomAttribute struct {
	Comment    *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	DisplayEnabled *bool `json:"DisplayEnabled,omitempty" xml:"DisplayEnabled,omitempty"`
	// example:
	//
	// 业务负责人
	DisplayName *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EntityTypes []*string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty" type:"Repeated"`
	// example:
	//
	// custom-attribute:biz_owner
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// false
	SearchFilterEnabled *bool `json:"SearchFilterEnabled,omitempty" xml:"SearchFilterEnabled,omitempty"`
	// example:
	//
	// TEXT
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	ValueEnums []*string `json:"ValueEnums,omitempty" xml:"ValueEnums,omitempty" type:"Repeated"`
}

func (s CustomAttribute) String() string {
	return dara.Prettify(s)
}

func (s CustomAttribute) GoString() string {
	return s.String()
}

func (s *CustomAttribute) GetComment() *string {
	return s.Comment
}

func (s *CustomAttribute) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CustomAttribute) GetDisplayEnabled() *bool {
	return s.DisplayEnabled
}

func (s *CustomAttribute) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CustomAttribute) GetEntityTypes() []*string {
	return s.EntityTypes
}

func (s *CustomAttribute) GetId() *string {
	return s.Id
}

func (s *CustomAttribute) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *CustomAttribute) GetSearchFilterEnabled() *bool {
	return s.SearchFilterEnabled
}

func (s *CustomAttribute) GetType() *string {
	return s.Type
}

func (s *CustomAttribute) GetValueEnums() []*string {
	return s.ValueEnums
}

func (s *CustomAttribute) SetComment(v string) *CustomAttribute {
	s.Comment = &v
	return s
}

func (s *CustomAttribute) SetCreateTime(v int64) *CustomAttribute {
	s.CreateTime = &v
	return s
}

func (s *CustomAttribute) SetDisplayEnabled(v bool) *CustomAttribute {
	s.DisplayEnabled = &v
	return s
}

func (s *CustomAttribute) SetDisplayName(v string) *CustomAttribute {
	s.DisplayName = &v
	return s
}

func (s *CustomAttribute) SetEntityTypes(v []*string) *CustomAttribute {
	s.EntityTypes = v
	return s
}

func (s *CustomAttribute) SetId(v string) *CustomAttribute {
	s.Id = &v
	return s
}

func (s *CustomAttribute) SetModifyTime(v int64) *CustomAttribute {
	s.ModifyTime = &v
	return s
}

func (s *CustomAttribute) SetSearchFilterEnabled(v bool) *CustomAttribute {
	s.SearchFilterEnabled = &v
	return s
}

func (s *CustomAttribute) SetType(v string) *CustomAttribute {
	s.Type = &v
	return s
}

func (s *CustomAttribute) SetValueEnums(v []*string) *CustomAttribute {
	s.ValueEnums = v
	return s
}

func (s *CustomAttribute) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateFieldRequest
	GetLang() *string
	SetClassify(v string) *CreateFieldRequest
	GetClassify() *string
	SetDescription(v string) *CreateFieldRequest
	GetDescription() *string
	SetEnumData(v string) *CreateFieldRequest
	GetEnumData() *string
	SetName(v string) *CreateFieldRequest
	GetName() *string
	SetRegId(v string) *CreateFieldRequest
	GetRegId() *string
	SetSource(v string) *CreateFieldRequest
	GetSource() *string
	SetTitle(v string) *CreateFieldRequest
	GetTitle() *string
	SetType(v string) *CreateFieldRequest
	GetType() *string
}

type CreateFieldRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Field classification.
	//
	// example:
	//
	// REQUEST_PARAM
	Classify *string `json:"classify,omitempty" xml:"classify,omitempty"`
	// Description information.
	//
	// example:
	//
	// 字段描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Enum data.
	//
	// example:
	//
	// enum
	EnumData *string `json:"enumData,omitempty" xml:"enumData,omitempty"`
	// Field name.
	//
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Region code.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Business source.
	//
	// example:
	//
	// DEFINE
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// Title.
	//
	// example:
	//
	// 年龄
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Field type.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFieldRequest) GoString() string {
	return s.String()
}

func (s *CreateFieldRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateFieldRequest) GetClassify() *string {
	return s.Classify
}

func (s *CreateFieldRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFieldRequest) GetEnumData() *string {
	return s.EnumData
}

func (s *CreateFieldRequest) GetName() *string {
	return s.Name
}

func (s *CreateFieldRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateFieldRequest) GetSource() *string {
	return s.Source
}

func (s *CreateFieldRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateFieldRequest) GetType() *string {
	return s.Type
}

func (s *CreateFieldRequest) SetLang(v string) *CreateFieldRequest {
	s.Lang = &v
	return s
}

func (s *CreateFieldRequest) SetClassify(v string) *CreateFieldRequest {
	s.Classify = &v
	return s
}

func (s *CreateFieldRequest) SetDescription(v string) *CreateFieldRequest {
	s.Description = &v
	return s
}

func (s *CreateFieldRequest) SetEnumData(v string) *CreateFieldRequest {
	s.EnumData = &v
	return s
}

func (s *CreateFieldRequest) SetName(v string) *CreateFieldRequest {
	s.Name = &v
	return s
}

func (s *CreateFieldRequest) SetRegId(v string) *CreateFieldRequest {
	s.RegId = &v
	return s
}

func (s *CreateFieldRequest) SetSource(v string) *CreateFieldRequest {
	s.Source = &v
	return s
}

func (s *CreateFieldRequest) SetTitle(v string) *CreateFieldRequest {
	s.Title = &v
	return s
}

func (s *CreateFieldRequest) SetType(v string) *CreateFieldRequest {
	s.Type = &v
	return s
}

func (s *CreateFieldRequest) Validate() error {
	return dara.Validate(s)
}

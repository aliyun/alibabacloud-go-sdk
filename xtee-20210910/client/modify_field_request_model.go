// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyFieldRequest
	GetLang() *string
	SetClassify(v string) *ModifyFieldRequest
	GetClassify() *string
	SetDescription(v string) *ModifyFieldRequest
	GetDescription() *string
	SetEnumData(v string) *ModifyFieldRequest
	GetEnumData() *string
	SetId(v int64) *ModifyFieldRequest
	GetId() *int64
	SetName(v string) *ModifyFieldRequest
	GetName() *string
	SetRegId(v string) *ModifyFieldRequest
	GetRegId() *string
	SetTitle(v string) *ModifyFieldRequest
	GetTitle() *string
}

type ModifyFieldRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Field classification
	//
	// example:
	//
	// REQUEST_PARAM
	Classify *string `json:"classify,omitempty" xml:"classify,omitempty"`
	// Description information.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Enum type
	//
	// example:
	//
	// [{\\"name\\":\\"ENABLE\\",\\"value\\":\\"禁用\\"}]
	EnumData *string `json:"enumData,omitempty" xml:"enumData,omitempty"`
	// Variable ID
	//
	// example:
	//
	// 376773
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Variable name
	//
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Title.
	//
	// example:
	//
	// 年龄
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ModifyFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFieldRequest) GoString() string {
	return s.String()
}

func (s *ModifyFieldRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyFieldRequest) GetClassify() *string {
	return s.Classify
}

func (s *ModifyFieldRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFieldRequest) GetEnumData() *string {
	return s.EnumData
}

func (s *ModifyFieldRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyFieldRequest) GetName() *string {
	return s.Name
}

func (s *ModifyFieldRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModifyFieldRequest) GetTitle() *string {
	return s.Title
}

func (s *ModifyFieldRequest) SetLang(v string) *ModifyFieldRequest {
	s.Lang = &v
	return s
}

func (s *ModifyFieldRequest) SetClassify(v string) *ModifyFieldRequest {
	s.Classify = &v
	return s
}

func (s *ModifyFieldRequest) SetDescription(v string) *ModifyFieldRequest {
	s.Description = &v
	return s
}

func (s *ModifyFieldRequest) SetEnumData(v string) *ModifyFieldRequest {
	s.EnumData = &v
	return s
}

func (s *ModifyFieldRequest) SetId(v int64) *ModifyFieldRequest {
	s.Id = &v
	return s
}

func (s *ModifyFieldRequest) SetName(v string) *ModifyFieldRequest {
	s.Name = &v
	return s
}

func (s *ModifyFieldRequest) SetRegId(v string) *ModifyFieldRequest {
	s.RegId = &v
	return s
}

func (s *ModifyFieldRequest) SetTitle(v string) *ModifyFieldRequest {
	s.Title = &v
	return s
}

func (s *ModifyFieldRequest) Validate() error {
	return dara.Validate(s)
}

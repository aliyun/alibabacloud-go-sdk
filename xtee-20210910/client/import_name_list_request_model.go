// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNameListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ImportNameListRequest
	GetLang() *string
	SetCreateType(v string) *ImportNameListRequest
	GetCreateType() *string
	SetData(v string) *ImportNameListRequest
	GetData() *string
	SetDescription(v string) *ImportNameListRequest
	GetDescription() *string
	SetImportType(v string) *ImportNameListRequest
	GetImportType() *string
	SetMemo(v string) *ImportNameListRequest
	GetMemo() *string
	SetNameListType(v string) *ImportNameListRequest
	GetNameListType() *string
	SetRegId(v string) *ImportNameListRequest
	GetRegId() *string
	SetTitle(v string) *ImportNameListRequest
	GetTitle() *string
	SetVariableId(v int64) *ImportNameListRequest
	GetVariableId() *int64
}

type ImportNameListRequest struct {
	// Set the language type for request and response messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Import name list.
	//
	// example:
	//
	// aa\\nbb\\ncc
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// Description information.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Document import type:
	//
	//
	// INPUT: Text input
	//
	// CSV: CSV upload
	//
	// NONE: Do not upload for now
	//
	// This parameter is required.
	//
	// example:
	//
	// CSV
	ImportType *string `json:"importType,omitempty" xml:"importType,omitempty"`
	// name content memo
	//
	// example:
	//
	// 名单内容描述
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// nameListType
	//
	// example:
	//
	// accountId
	NameListType *string `json:"nameListType,omitempty" xml:"nameListType,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Title.
	//
	// This parameter is required.
	//
	// example:
	//
	// 变量title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Variable ID
	//
	// example:
	//
	// 393314
	VariableId *int64 `json:"variableId,omitempty" xml:"variableId,omitempty"`
}

func (s ImportNameListRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportNameListRequest) GoString() string {
	return s.String()
}

func (s *ImportNameListRequest) GetLang() *string {
	return s.Lang
}

func (s *ImportNameListRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *ImportNameListRequest) GetData() *string {
	return s.Data
}

func (s *ImportNameListRequest) GetDescription() *string {
	return s.Description
}

func (s *ImportNameListRequest) GetImportType() *string {
	return s.ImportType
}

func (s *ImportNameListRequest) GetMemo() *string {
	return s.Memo
}

func (s *ImportNameListRequest) GetNameListType() *string {
	return s.NameListType
}

func (s *ImportNameListRequest) GetRegId() *string {
	return s.RegId
}

func (s *ImportNameListRequest) GetTitle() *string {
	return s.Title
}

func (s *ImportNameListRequest) GetVariableId() *int64 {
	return s.VariableId
}

func (s *ImportNameListRequest) SetLang(v string) *ImportNameListRequest {
	s.Lang = &v
	return s
}

func (s *ImportNameListRequest) SetCreateType(v string) *ImportNameListRequest {
	s.CreateType = &v
	return s
}

func (s *ImportNameListRequest) SetData(v string) *ImportNameListRequest {
	s.Data = &v
	return s
}

func (s *ImportNameListRequest) SetDescription(v string) *ImportNameListRequest {
	s.Description = &v
	return s
}

func (s *ImportNameListRequest) SetImportType(v string) *ImportNameListRequest {
	s.ImportType = &v
	return s
}

func (s *ImportNameListRequest) SetMemo(v string) *ImportNameListRequest {
	s.Memo = &v
	return s
}

func (s *ImportNameListRequest) SetNameListType(v string) *ImportNameListRequest {
	s.NameListType = &v
	return s
}

func (s *ImportNameListRequest) SetRegId(v string) *ImportNameListRequest {
	s.RegId = &v
	return s
}

func (s *ImportNameListRequest) SetTitle(v string) *ImportNameListRequest {
	s.Title = &v
	return s
}

func (s *ImportNameListRequest) SetVariableId(v int64) *ImportNameListRequest {
	s.VariableId = &v
	return s
}

func (s *ImportNameListRequest) Validate() error {
	return dara.Validate(s)
}

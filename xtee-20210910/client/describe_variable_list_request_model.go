// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeVariableListRequest
	GetLang() *string
	SetCreateType(v string) *DescribeVariableListRequest
	GetCreateType() *string
	SetCurrentPage(v string) *DescribeVariableListRequest
	GetCurrentPage() *string
	SetPageSize(v string) *DescribeVariableListRequest
	GetPageSize() *string
	SetRefObjId(v string) *DescribeVariableListRequest
	GetRefObjId() *string
	SetRegId(v string) *DescribeVariableListRequest
	GetRegId() *string
	SetSourceType(v string) *DescribeVariableListRequest
	GetSourceType() *string
	SetType(v string) *DescribeVariableListRequest
	GetType() *string
	SetTypesStr(v string) *DescribeVariableListRequest
	GetTypesStr() *string
	SetValue(v string) *DescribeVariableListRequest
	GetValue() *string
}

type DescribeVariableListRequest struct {
	// Set the language type for requests and responses, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Creation type.
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Associated event ID.
	//
	// example:
	//
	// de_avypfd8253
	RefObjId *string `json:"refObjId,omitempty" xml:"refObjId,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Source type.
	//
	// example:
	//
	// SAF
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// Type.
	//
	// example:
	//
	// IDENTIFY_SERVICE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// JSON array string of types.
	//
	// example:
	//
	// [\\"coupon_abuse_detection\\"]
	TypesStr *string `json:"typesStr,omitempty" xml:"typesStr,omitempty"`
	// Value for fuzzy search.
	//
	// example:
	//
	// 注册
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeVariableListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVariableListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVariableListRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeVariableListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVariableListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVariableListRequest) GetRefObjId() *string {
	return s.RefObjId
}

func (s *DescribeVariableListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeVariableListRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeVariableListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeVariableListRequest) GetTypesStr() *string {
	return s.TypesStr
}

func (s *DescribeVariableListRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeVariableListRequest) SetLang(v string) *DescribeVariableListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVariableListRequest) SetCreateType(v string) *DescribeVariableListRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeVariableListRequest) SetCurrentPage(v string) *DescribeVariableListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVariableListRequest) SetPageSize(v string) *DescribeVariableListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVariableListRequest) SetRefObjId(v string) *DescribeVariableListRequest {
	s.RefObjId = &v
	return s
}

func (s *DescribeVariableListRequest) SetRegId(v string) *DescribeVariableListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeVariableListRequest) SetSourceType(v string) *DescribeVariableListRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeVariableListRequest) SetType(v string) *DescribeVariableListRequest {
	s.Type = &v
	return s
}

func (s *DescribeVariableListRequest) SetTypesStr(v string) *DescribeVariableListRequest {
	s.TypesStr = &v
	return s
}

func (s *DescribeVariableListRequest) SetValue(v string) *DescribeVariableListRequest {
	s.Value = &v
	return s
}

func (s *DescribeVariableListRequest) Validate() error {
	return dara.Validate(s)
}

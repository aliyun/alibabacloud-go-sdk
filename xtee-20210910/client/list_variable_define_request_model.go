// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariableDefineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListVariableDefineRequest
	GetLang() *string
	SetAllowBind(v string) *ListVariableDefineRequest
	GetAllowBind() *string
	SetChargingMode(v string) *ListVariableDefineRequest
	GetChargingMode() *string
	SetCurrentPage(v int32) *ListVariableDefineRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListVariableDefineRequest
	GetPageSize() *int32
	SetPaging(v string) *ListVariableDefineRequest
	GetPaging() *string
	SetQueryContent(v string) *ListVariableDefineRequest
	GetQueryContent() *string
	SetRegId(v string) *ListVariableDefineRequest
	GetRegId() *string
	SetRoleType(v string) *ListVariableDefineRequest
	GetRoleType() *string
	SetScenesStr(v string) *ListVariableDefineRequest
	GetScenesStr() *string
	SetSource(v string) *ListVariableDefineRequest
	GetSource() *string
	SetStatus(v string) *ListVariableDefineRequest
	GetStatus() *string
	SetTitle(v string) *ListVariableDefineRequest
	GetTitle() *string
	SetTypesStr(v string) *ListVariableDefineRequest
	GetTypesStr() *string
}

type ListVariableDefineRequest struct {
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
	// Whether binding is allowed, default is ENABLE
	//
	// example:
	//
	// ENABLE
	AllowBind *string `json:"allowBind,omitempty" xml:"allowBind,omitempty"`
	// Charging mode
	//
	// example:
	//
	// FREE
	ChargingMode *string `json:"chargingMode,omitempty" xml:"chargingMode,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Paging identifier
	//
	// example:
	//
	// false
	Paging *string `json:"paging,omitempty" xml:"paging,omitempty"`
	// Query content
	//
	// example:
	//
	// age
	QueryContent *string `json:"queryContent,omitempty" xml:"queryContent,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Authorization type
	//
	// example:
	//
	// 1
	RoleType *string `json:"roleType,omitempty" xml:"roleType,omitempty"`
	// Scenario
	//
	// example:
	//
	// [\\"EVENT_ACTION\\"]
	ScenesStr *string `json:"scenesStr,omitempty" xml:"scenesStr,omitempty"`
	// Source
	//
	// example:
	//
	// SAF
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Title.
	//
	// example:
	//
	// 变量的title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Type
	//
	// example:
	//
	// [\\"IDENTIFY_SERVICE\\"]
	TypesStr *string `json:"typesStr,omitempty" xml:"typesStr,omitempty"`
}

func (s ListVariableDefineRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVariableDefineRequest) GoString() string {
	return s.String()
}

func (s *ListVariableDefineRequest) GetLang() *string {
	return s.Lang
}

func (s *ListVariableDefineRequest) GetAllowBind() *string {
	return s.AllowBind
}

func (s *ListVariableDefineRequest) GetChargingMode() *string {
	return s.ChargingMode
}

func (s *ListVariableDefineRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListVariableDefineRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVariableDefineRequest) GetPaging() *string {
	return s.Paging
}

func (s *ListVariableDefineRequest) GetQueryContent() *string {
	return s.QueryContent
}

func (s *ListVariableDefineRequest) GetRegId() *string {
	return s.RegId
}

func (s *ListVariableDefineRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *ListVariableDefineRequest) GetScenesStr() *string {
	return s.ScenesStr
}

func (s *ListVariableDefineRequest) GetSource() *string {
	return s.Source
}

func (s *ListVariableDefineRequest) GetStatus() *string {
	return s.Status
}

func (s *ListVariableDefineRequest) GetTitle() *string {
	return s.Title
}

func (s *ListVariableDefineRequest) GetTypesStr() *string {
	return s.TypesStr
}

func (s *ListVariableDefineRequest) SetLang(v string) *ListVariableDefineRequest {
	s.Lang = &v
	return s
}

func (s *ListVariableDefineRequest) SetAllowBind(v string) *ListVariableDefineRequest {
	s.AllowBind = &v
	return s
}

func (s *ListVariableDefineRequest) SetChargingMode(v string) *ListVariableDefineRequest {
	s.ChargingMode = &v
	return s
}

func (s *ListVariableDefineRequest) SetCurrentPage(v int32) *ListVariableDefineRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListVariableDefineRequest) SetPageSize(v int32) *ListVariableDefineRequest {
	s.PageSize = &v
	return s
}

func (s *ListVariableDefineRequest) SetPaging(v string) *ListVariableDefineRequest {
	s.Paging = &v
	return s
}

func (s *ListVariableDefineRequest) SetQueryContent(v string) *ListVariableDefineRequest {
	s.QueryContent = &v
	return s
}

func (s *ListVariableDefineRequest) SetRegId(v string) *ListVariableDefineRequest {
	s.RegId = &v
	return s
}

func (s *ListVariableDefineRequest) SetRoleType(v string) *ListVariableDefineRequest {
	s.RoleType = &v
	return s
}

func (s *ListVariableDefineRequest) SetScenesStr(v string) *ListVariableDefineRequest {
	s.ScenesStr = &v
	return s
}

func (s *ListVariableDefineRequest) SetSource(v string) *ListVariableDefineRequest {
	s.Source = &v
	return s
}

func (s *ListVariableDefineRequest) SetStatus(v string) *ListVariableDefineRequest {
	s.Status = &v
	return s
}

func (s *ListVariableDefineRequest) SetTitle(v string) *ListVariableDefineRequest {
	s.Title = &v
	return s
}

func (s *ListVariableDefineRequest) SetTypesStr(v string) *ListVariableDefineRequest {
	s.TypesStr = &v
	return s
}

func (s *ListVariableDefineRequest) Validate() error {
	return dara.Validate(s)
}

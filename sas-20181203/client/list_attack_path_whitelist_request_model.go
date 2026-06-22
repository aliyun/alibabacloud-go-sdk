// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttackPathWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAttackPathWhitelistRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListAttackPathWhitelistRequest
	GetLang() *string
	SetPageSize(v int32) *ListAttackPathWhitelistRequest
	GetPageSize() *int32
	SetPathNameDesc(v string) *ListAttackPathWhitelistRequest
	GetPathNameDesc() *string
	SetPathType(v string) *ListAttackPathWhitelistRequest
	GetPathType() *string
	SetWhitelistName(v string) *ListAttackPathWhitelistRequest
	GetWhitelistName() *string
}

type ListAttackPathWhitelistRequest struct {
	// The page number when using paging. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page when using paging. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The path name description.
	//
	// > Call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query path name descriptions.
	//
	// example:
	//
	// ECS Instance Can Obtain Long-term Access Credential by Enabling Console Logon for RAM User
	PathNameDesc *string `json:"PathNameDesc,omitempty" xml:"PathNameDesc,omitempty"`
	// The path type.
	//
	// > Call [ListAvailableAttackPath](~~ListAvailableAttackPath~~) to query path types.
	//
	// example:
	//
	// role_escalation
	PathType *string `json:"PathType,omitempty" xml:"PathType,omitempty"`
	// The whitelist name.
	//
	// example:
	//
	// test
	WhitelistName *string `json:"WhitelistName,omitempty" xml:"WhitelistName,omitempty"`
}

func (s ListAttackPathWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAttackPathWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ListAttackPathWhitelistRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAttackPathWhitelistRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAttackPathWhitelistRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAttackPathWhitelistRequest) GetPathNameDesc() *string {
	return s.PathNameDesc
}

func (s *ListAttackPathWhitelistRequest) GetPathType() *string {
	return s.PathType
}

func (s *ListAttackPathWhitelistRequest) GetWhitelistName() *string {
	return s.WhitelistName
}

func (s *ListAttackPathWhitelistRequest) SetCurrentPage(v int32) *ListAttackPathWhitelistRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAttackPathWhitelistRequest) SetLang(v string) *ListAttackPathWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *ListAttackPathWhitelistRequest) SetPageSize(v int32) *ListAttackPathWhitelistRequest {
	s.PageSize = &v
	return s
}

func (s *ListAttackPathWhitelistRequest) SetPathNameDesc(v string) *ListAttackPathWhitelistRequest {
	s.PathNameDesc = &v
	return s
}

func (s *ListAttackPathWhitelistRequest) SetPathType(v string) *ListAttackPathWhitelistRequest {
	s.PathType = &v
	return s
}

func (s *ListAttackPathWhitelistRequest) SetWhitelistName(v string) *ListAttackPathWhitelistRequest {
	s.WhitelistName = &v
	return s
}

func (s *ListAttackPathWhitelistRequest) Validate() error {
	return dara.Validate(s)
}

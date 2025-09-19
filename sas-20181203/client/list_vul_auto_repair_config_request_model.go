// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulAutoRepairConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *ListVulAutoRepairConfigRequest
	GetAliasName() *string
	SetCurrentPage(v int32) *ListVulAutoRepairConfigRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListVulAutoRepairConfigRequest
	GetLang() *string
	SetPageSize(v int32) *ListVulAutoRepairConfigRequest
	GetPageSize() *int32
	SetType(v string) *ListVulAutoRepairConfigRequest
	GetType() *string
}

type ListVulAutoRepairConfigRequest struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// RHSA-2022:0274-Important: polkit pkexec Local Privilege Escalation Vulnerability(CVE-2021-4034)
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListVulAutoRepairConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVulAutoRepairConfigRequest) GoString() string {
	return s.String()
}

func (s *ListVulAutoRepairConfigRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *ListVulAutoRepairConfigRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListVulAutoRepairConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *ListVulAutoRepairConfigRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVulAutoRepairConfigRequest) GetType() *string {
	return s.Type
}

func (s *ListVulAutoRepairConfigRequest) SetAliasName(v string) *ListVulAutoRepairConfigRequest {
	s.AliasName = &v
	return s
}

func (s *ListVulAutoRepairConfigRequest) SetCurrentPage(v int32) *ListVulAutoRepairConfigRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListVulAutoRepairConfigRequest) SetLang(v string) *ListVulAutoRepairConfigRequest {
	s.Lang = &v
	return s
}

func (s *ListVulAutoRepairConfigRequest) SetPageSize(v int32) *ListVulAutoRepairConfigRequest {
	s.PageSize = &v
	return s
}

func (s *ListVulAutoRepairConfigRequest) SetType(v string) *ListVulAutoRepairConfigRequest {
	s.Type = &v
	return s
}

func (s *ListVulAutoRepairConfigRequest) Validate() error {
	return dara.Validate(s)
}

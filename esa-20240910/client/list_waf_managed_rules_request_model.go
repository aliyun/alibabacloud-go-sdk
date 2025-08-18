// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafManagedRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackType(v int32) *ListWafManagedRulesRequest
	GetAttackType() *int32
	SetId(v int64) *ListWafManagedRulesRequest
	GetId() *int64
	SetLanguage(v string) *ListWafManagedRulesRequest
	GetLanguage() *string
	SetPageNumber(v int32) *ListWafManagedRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWafManagedRulesRequest
	GetPageSize() *int32
	SetProtectionLevel(v int32) *ListWafManagedRulesRequest
	GetProtectionLevel() *int32
	SetQueryArgs(v *ListWafManagedRulesRequestQueryArgs) *ListWafManagedRulesRequest
	GetQueryArgs() *ListWafManagedRulesRequestQueryArgs
	SetSiteId(v int64) *ListWafManagedRulesRequest
	GetSiteId() *int64
}

type ListWafManagedRulesRequest struct {
	// Attack type of the vulnerability protection event. Values:
	//
	// - SQL injection
	//
	// - Cross-site scripting
	//
	// - Code execution
	//
	// - CRLF
	//
	// - Local file inclusion
	//
	// - Remote file inclusion
	//
	// - Webshell
	//
	// - Cross-site request forgery
	//
	// - Other
	//
	// - SEMA
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	AttackType *int32 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// ID of the WAF rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Language type, which will be used to return the response. Value range:
	//
	// - **en**: English.
	//
	// - **zh**: Chinese.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Query page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Query page size.
	//
	// example:
	//
	// 20
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProtectionLevel *int32 `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// Query conditions.
	QueryArgs *ListWafManagedRulesRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListWafManagedRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWafManagedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWafManagedRulesRequest) GetAttackType() *int32 {
	return s.AttackType
}

func (s *ListWafManagedRulesRequest) GetId() *int64 {
	return s.Id
}

func (s *ListWafManagedRulesRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListWafManagedRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWafManagedRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWafManagedRulesRequest) GetProtectionLevel() *int32 {
	return s.ProtectionLevel
}

func (s *ListWafManagedRulesRequest) GetQueryArgs() *ListWafManagedRulesRequestQueryArgs {
	return s.QueryArgs
}

func (s *ListWafManagedRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWafManagedRulesRequest) SetAttackType(v int32) *ListWafManagedRulesRequest {
	s.AttackType = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetId(v int64) *ListWafManagedRulesRequest {
	s.Id = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetLanguage(v string) *ListWafManagedRulesRequest {
	s.Language = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetPageNumber(v int32) *ListWafManagedRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetPageSize(v int32) *ListWafManagedRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetProtectionLevel(v int32) *ListWafManagedRulesRequest {
	s.ProtectionLevel = &v
	return s
}

func (s *ListWafManagedRulesRequest) SetQueryArgs(v *ListWafManagedRulesRequestQueryArgs) *ListWafManagedRulesRequest {
	s.QueryArgs = v
	return s
}

func (s *ListWafManagedRulesRequest) SetSiteId(v int64) *ListWafManagedRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafManagedRulesRequest) Validate() error {
	return dara.Validate(s)
}

type ListWafManagedRulesRequestQueryArgs struct {
	// Action.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// Fuzzy search for rule ID or rule name.
	//
	// example:
	//
	// example
	IdNameLike *string `json:"IdNameLike,omitempty" xml:"IdNameLike,omitempty"`
	// List of rule protection levels.
	ProtectionLevels []*int32 `json:"ProtectionLevels,omitempty" xml:"ProtectionLevels,omitempty" type:"Repeated"`
	// Status.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWafManagedRulesRequestQueryArgs) String() string {
	return dara.Prettify(s)
}

func (s ListWafManagedRulesRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListWafManagedRulesRequestQueryArgs) GetAction() *string {
	return s.Action
}

func (s *ListWafManagedRulesRequestQueryArgs) GetIdNameLike() *string {
	return s.IdNameLike
}

func (s *ListWafManagedRulesRequestQueryArgs) GetProtectionLevels() []*int32 {
	return s.ProtectionLevels
}

func (s *ListWafManagedRulesRequestQueryArgs) GetStatus() *string {
	return s.Status
}

func (s *ListWafManagedRulesRequestQueryArgs) SetAction(v string) *ListWafManagedRulesRequestQueryArgs {
	s.Action = &v
	return s
}

func (s *ListWafManagedRulesRequestQueryArgs) SetIdNameLike(v string) *ListWafManagedRulesRequestQueryArgs {
	s.IdNameLike = &v
	return s
}

func (s *ListWafManagedRulesRequestQueryArgs) SetProtectionLevels(v []*int32) *ListWafManagedRulesRequestQueryArgs {
	s.ProtectionLevels = v
	return s
}

func (s *ListWafManagedRulesRequestQueryArgs) SetStatus(v string) *ListWafManagedRulesRequestQueryArgs {
	s.Status = &v
	return s
}

func (s *ListWafManagedRulesRequestQueryArgs) Validate() error {
	return dara.Validate(s)
}

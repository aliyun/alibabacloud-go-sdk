// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafManagedRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackType(v int32) *ListWafManagedRulesShrinkRequest
	GetAttackType() *int32
	SetId(v int64) *ListWafManagedRulesShrinkRequest
	GetId() *int64
	SetInstanceId(v string) *ListWafManagedRulesShrinkRequest
	GetInstanceId() *string
	SetLanguage(v string) *ListWafManagedRulesShrinkRequest
	GetLanguage() *string
	SetManagedRulesetShrink(v string) *ListWafManagedRulesShrinkRequest
	GetManagedRulesetShrink() *string
	SetPageNumber(v int32) *ListWafManagedRulesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWafManagedRulesShrinkRequest
	GetPageSize() *int32
	SetProtectionLevel(v int32) *ListWafManagedRulesShrinkRequest
	GetProtectionLevel() *int32
	SetQueryArgsShrink(v string) *ListWafManagedRulesShrinkRequest
	GetQueryArgsShrink() *string
	SetSiteId(v int64) *ListWafManagedRulesShrinkRequest
	GetSiteId() *int64
}

type ListWafManagedRulesShrinkRequest struct {
	// The attack type to filter the results by. Valid values:
	//
	// - SQL injection
	//
	// - cross-site scripting
	//
	// - code execution
	//
	// - CRLF
	//
	// - local file inclusion
	//
	// - remote file inclusion
	//
	// - webshell
	//
	// - cross-site request forgery
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
	// The ID of the WAF rule.
	//
	// example:
	//
	// 10000001
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The response language. Valid values:
	//
	// - **en**: English.
	//
	// - **zh**: Chinese.
	//
	// example:
	//
	// zh
	Language             *string `json:"Language,omitempty" xml:"Language,omitempty"`
	ManagedRulesetShrink *string `json:"ManagedRuleset,omitempty" xml:"ManagedRuleset,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProtectionLevel *int32 `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// The query conditions.
	QueryArgsShrink *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
	// The ID of the site. Call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain this ID.
	//
	// example:
	//
	// 1
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListWafManagedRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWafManagedRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWafManagedRulesShrinkRequest) GetAttackType() *int32 {
	return s.AttackType
}

func (s *ListWafManagedRulesShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *ListWafManagedRulesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWafManagedRulesShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListWafManagedRulesShrinkRequest) GetManagedRulesetShrink() *string {
	return s.ManagedRulesetShrink
}

func (s *ListWafManagedRulesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWafManagedRulesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWafManagedRulesShrinkRequest) GetProtectionLevel() *int32 {
	return s.ProtectionLevel
}

func (s *ListWafManagedRulesShrinkRequest) GetQueryArgsShrink() *string {
	return s.QueryArgsShrink
}

func (s *ListWafManagedRulesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWafManagedRulesShrinkRequest) SetAttackType(v int32) *ListWafManagedRulesShrinkRequest {
	s.AttackType = &v
	return s
}

func (s *ListWafManagedRulesShrinkRequest) SetId(v int64) *ListWafManagedRulesShrinkRequest {
	s.Id = &v
	return s
}

func (s *ListWafManagedRulesShrinkRequest) SetInstanceId(v string) *ListWafManagedRulesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListWafManagedRulesShrinkRequest) SetLanguage(v string) *ListWafManagedRulesShrinkRequest {
	s.Language = &v
	return s
}

func (s *ListWafManagedRulesShrinkRequest) SetManagedRulesetShrink(v string) *ListWafManagedRulesShrinkRequest {
	s.ManagedRulesetShrink = &v
	return s
}

func (s *ListWafManagedRulesShrinkRequest) SetPageNumber(v int32) *ListWafManagedRulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWafManagedRulesShrinkRequest) SetPageSize(v int32) *ListWafManagedRulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListWafManagedRulesShrinkRequest) SetProtectionLevel(v int32) *ListWafManagedRulesShrinkRequest {
	s.ProtectionLevel = &v
	return s
}

func (s *ListWafManagedRulesShrinkRequest) SetQueryArgsShrink(v string) *ListWafManagedRulesShrinkRequest {
	s.QueryArgsShrink = &v
	return s
}

func (s *ListWafManagedRulesShrinkRequest) SetSiteId(v int64) *ListWafManagedRulesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafManagedRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

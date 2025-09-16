// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpIncomingRequestHeaderModificationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListHttpIncomingRequestHeaderModificationRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListHttpIncomingRequestHeaderModificationRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListHttpIncomingRequestHeaderModificationRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpIncomingRequestHeaderModificationRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListHttpIncomingRequestHeaderModificationRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListHttpIncomingRequestHeaderModificationRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListHttpIncomingRequestHeaderModificationRulesRequest
	GetSiteVersion() *int32
}

type ListHttpIncomingRequestHeaderModificationRulesRequest struct {
	// example:
	//
	// 424371770570752
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 608665779308176
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListHttpIncomingRequestHeaderModificationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpIncomingRequestHeaderModificationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) SetConfigId(v int64) *ListHttpIncomingRequestHeaderModificationRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) SetConfigType(v string) *ListHttpIncomingRequestHeaderModificationRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) SetPageNumber(v int32) *ListHttpIncomingRequestHeaderModificationRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) SetPageSize(v int32) *ListHttpIncomingRequestHeaderModificationRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) SetRuleName(v string) *ListHttpIncomingRequestHeaderModificationRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) SetSiteId(v int64) *ListHttpIncomingRequestHeaderModificationRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) SetSiteVersion(v int32) *ListHttpIncomingRequestHeaderModificationRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpIncomingRequestHeaderModificationRulesRequest) Validate() error {
	return dara.Validate(s)
}

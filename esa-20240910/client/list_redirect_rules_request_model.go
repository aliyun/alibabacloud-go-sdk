// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRedirectRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListRedirectRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListRedirectRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListRedirectRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRedirectRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListRedirectRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListRedirectRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListRedirectRulesRequest
	GetSiteVersion() *int32
}

type ListRedirectRulesRequest struct {
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListRedirectRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRedirectRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRedirectRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListRedirectRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListRedirectRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRedirectRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRedirectRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRedirectRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListRedirectRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListRedirectRulesRequest) SetConfigId(v int64) *ListRedirectRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListRedirectRulesRequest) SetConfigType(v string) *ListRedirectRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListRedirectRulesRequest) SetPageNumber(v int32) *ListRedirectRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRedirectRulesRequest) SetPageSize(v int32) *ListRedirectRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRedirectRulesRequest) SetRuleName(v string) *ListRedirectRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListRedirectRulesRequest) SetSiteId(v int64) *ListRedirectRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListRedirectRulesRequest) SetSiteVersion(v int32) *ListRedirectRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListRedirectRulesRequest) Validate() error {
	return dara.Validate(s)
}

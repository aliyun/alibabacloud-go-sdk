// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCacheRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListCacheRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListCacheRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListCacheRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCacheRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListCacheRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListCacheRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListCacheRulesRequest
	GetSiteVersion() *int32
}

type ListCacheRulesRequest struct {
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListCacheRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCacheRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCacheRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListCacheRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListCacheRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCacheRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCacheRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCacheRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListCacheRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListCacheRulesRequest) SetConfigId(v int64) *ListCacheRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListCacheRulesRequest) SetConfigType(v string) *ListCacheRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListCacheRulesRequest) SetPageNumber(v int32) *ListCacheRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCacheRulesRequest) SetPageSize(v int32) *ListCacheRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCacheRulesRequest) SetRuleName(v string) *ListCacheRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListCacheRulesRequest) SetSiteId(v int64) *ListCacheRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListCacheRulesRequest) SetSiteVersion(v int32) *ListCacheRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListCacheRulesRequest) Validate() error {
	return dara.Validate(s)
}

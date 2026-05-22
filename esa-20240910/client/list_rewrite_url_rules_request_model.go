// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRewriteUrlRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListRewriteUrlRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListRewriteUrlRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListRewriteUrlRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRewriteUrlRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListRewriteUrlRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListRewriteUrlRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListRewriteUrlRulesRequest
	GetSiteVersion() *int32
}

type ListRewriteUrlRulesRequest struct {
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListRewriteUrlRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRewriteUrlRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRewriteUrlRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListRewriteUrlRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListRewriteUrlRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRewriteUrlRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRewriteUrlRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRewriteUrlRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListRewriteUrlRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListRewriteUrlRulesRequest) SetConfigId(v int64) *ListRewriteUrlRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetConfigType(v string) *ListRewriteUrlRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetPageNumber(v int32) *ListRewriteUrlRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetPageSize(v int32) *ListRewriteUrlRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetRuleName(v string) *ListRewriteUrlRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetSiteId(v int64) *ListRewriteUrlRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) SetSiteVersion(v int32) *ListRewriteUrlRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListRewriteUrlRulesRequest) Validate() error {
	return dara.Validate(s)
}

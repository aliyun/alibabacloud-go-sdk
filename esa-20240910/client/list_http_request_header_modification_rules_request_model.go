// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpRequestHeaderModificationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListHttpRequestHeaderModificationRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListHttpRequestHeaderModificationRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListHttpRequestHeaderModificationRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpRequestHeaderModificationRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListHttpRequestHeaderModificationRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListHttpRequestHeaderModificationRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListHttpRequestHeaderModificationRulesRequest
	GetSiteVersion() *int32
}

type ListHttpRequestHeaderModificationRulesRequest struct {
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListHttpRequestHeaderModificationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpRequestHeaderModificationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListHttpRequestHeaderModificationRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetConfigId(v int64) *ListHttpRequestHeaderModificationRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetConfigType(v string) *ListHttpRequestHeaderModificationRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetPageNumber(v int32) *ListHttpRequestHeaderModificationRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetPageSize(v int32) *ListHttpRequestHeaderModificationRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetRuleName(v string) *ListHttpRequestHeaderModificationRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetSiteId(v int64) *ListHttpRequestHeaderModificationRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) SetSiteVersion(v int32) *ListHttpRequestHeaderModificationRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpRequestHeaderModificationRulesRequest) Validate() error {
	return dara.Validate(s)
}

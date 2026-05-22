// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpResponseHeaderModificationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListHttpResponseHeaderModificationRulesRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListHttpResponseHeaderModificationRulesRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListHttpResponseHeaderModificationRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpResponseHeaderModificationRulesRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListHttpResponseHeaderModificationRulesRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListHttpResponseHeaderModificationRulesRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListHttpResponseHeaderModificationRulesRequest
	GetSiteVersion() *int32
}

type ListHttpResponseHeaderModificationRulesRequest struct {
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListHttpResponseHeaderModificationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpResponseHeaderModificationRulesRequest) GoString() string {
	return s.String()
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListHttpResponseHeaderModificationRulesRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetConfigId(v int64) *ListHttpResponseHeaderModificationRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetConfigType(v string) *ListHttpResponseHeaderModificationRulesRequest {
	s.ConfigType = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetPageNumber(v int32) *ListHttpResponseHeaderModificationRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetPageSize(v int32) *ListHttpResponseHeaderModificationRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetRuleName(v string) *ListHttpResponseHeaderModificationRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetSiteId(v int64) *ListHttpResponseHeaderModificationRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) SetSiteVersion(v int32) *ListHttpResponseHeaderModificationRulesRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListHttpResponseHeaderModificationRulesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageTransformsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListImageTransformsRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListImageTransformsRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListImageTransformsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImageTransformsRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListImageTransformsRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListImageTransformsRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListImageTransformsRequest
	GetSiteVersion() *int32
}

type ListImageTransformsRequest struct {
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListImageTransformsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageTransformsRequest) GoString() string {
	return s.String()
}

func (s *ListImageTransformsRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListImageTransformsRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListImageTransformsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImageTransformsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageTransformsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListImageTransformsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListImageTransformsRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListImageTransformsRequest) SetConfigId(v int64) *ListImageTransformsRequest {
	s.ConfigId = &v
	return s
}

func (s *ListImageTransformsRequest) SetConfigType(v string) *ListImageTransformsRequest {
	s.ConfigType = &v
	return s
}

func (s *ListImageTransformsRequest) SetPageNumber(v int32) *ListImageTransformsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImageTransformsRequest) SetPageSize(v int32) *ListImageTransformsRequest {
	s.PageSize = &v
	return s
}

func (s *ListImageTransformsRequest) SetRuleName(v string) *ListImageTransformsRequest {
	s.RuleName = &v
	return s
}

func (s *ListImageTransformsRequest) SetSiteId(v int64) *ListImageTransformsRequest {
	s.SiteId = &v
	return s
}

func (s *ListImageTransformsRequest) SetSiteVersion(v int32) *ListImageTransformsRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListImageTransformsRequest) Validate() error {
	return dara.Validate(s)
}

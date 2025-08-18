// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVideoProcessingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListVideoProcessingsRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListVideoProcessingsRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListVideoProcessingsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVideoProcessingsRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListVideoProcessingsRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListVideoProcessingsRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *ListVideoProcessingsRequest
	GetSiteVersion() *int32
}

type ListVideoProcessingsRequest struct {
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// global
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
	// 12345
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListVideoProcessingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVideoProcessingsRequest) GoString() string {
	return s.String()
}

func (s *ListVideoProcessingsRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListVideoProcessingsRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListVideoProcessingsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVideoProcessingsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVideoProcessingsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListVideoProcessingsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListVideoProcessingsRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListVideoProcessingsRequest) SetConfigId(v int64) *ListVideoProcessingsRequest {
	s.ConfigId = &v
	return s
}

func (s *ListVideoProcessingsRequest) SetConfigType(v string) *ListVideoProcessingsRequest {
	s.ConfigType = &v
	return s
}

func (s *ListVideoProcessingsRequest) SetPageNumber(v int32) *ListVideoProcessingsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVideoProcessingsRequest) SetPageSize(v int32) *ListVideoProcessingsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVideoProcessingsRequest) SetRuleName(v string) *ListVideoProcessingsRequest {
	s.RuleName = &v
	return s
}

func (s *ListVideoProcessingsRequest) SetSiteId(v int64) *ListVideoProcessingsRequest {
	s.SiteId = &v
	return s
}

func (s *ListVideoProcessingsRequest) SetSiteVersion(v int32) *ListVideoProcessingsRequest {
	s.SiteVersion = &v
	return s
}

func (s *ListVideoProcessingsRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpsBasicConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListHttpsBasicConfigurationsRequest
	GetConfigId() *int64
	SetConfigType(v string) *ListHttpsBasicConfigurationsRequest
	GetConfigType() *string
	SetPageNumber(v int32) *ListHttpsBasicConfigurationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpsBasicConfigurationsRequest
	GetPageSize() *int32
	SetRuleName(v string) *ListHttpsBasicConfigurationsRequest
	GetRuleName() *string
	SetSiteId(v int64) *ListHttpsBasicConfigurationsRequest
	GetSiteId() *int64
}

type ListHttpsBasicConfigurationsRequest struct {
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListHttpsBasicConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHttpsBasicConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListHttpsBasicConfigurationsRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpsBasicConfigurationsRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpsBasicConfigurationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpsBasicConfigurationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpsBasicConfigurationsRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpsBasicConfigurationsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListHttpsBasicConfigurationsRequest) SetConfigId(v int64) *ListHttpsBasicConfigurationsRequest {
	s.ConfigId = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) SetConfigType(v string) *ListHttpsBasicConfigurationsRequest {
	s.ConfigType = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) SetPageNumber(v int32) *ListHttpsBasicConfigurationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) SetPageSize(v int32) *ListHttpsBasicConfigurationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) SetRuleName(v string) *ListHttpsBasicConfigurationsRequest {
	s.RuleName = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) SetSiteId(v int64) *ListHttpsBasicConfigurationsRequest {
	s.SiteId = &v
	return s
}

func (s *ListHttpsBasicConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}

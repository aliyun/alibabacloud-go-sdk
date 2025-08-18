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
	// Configuration ID. Can be obtained by calling the [ListImageTransforms](https://help.aliyun.com/document_detail/2869056.html) interface.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule configurations. Possible values:
	//
	// - global: Query global configuration;
	//
	// - rule: Query rule configuration;
	//
	// This parameter is optional. If not provided, it will not distinguish between global and rule configurations.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Page number. The default value is 1 if not provided.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of items per page. The maximum value is 500, and the default value is 500 if not provided.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Rule name. This parameter is not required when adding a global configuration.
	//
	// example:
	//
	// test1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Site ID. Can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Site version number. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. The default value is version 0.
	//
	// example:
	//
	// 1
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

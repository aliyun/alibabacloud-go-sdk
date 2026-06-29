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
	// The configuration ID. You can call the [ListVideoProcessings](~~ListVideoProcessings~~) operation to obtain the configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. You can use this parameter to query global or rule configurations. This parameter takes effect only when functionName is specified.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The page number for a paged query. The value must be greater than or equal to 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Valid values: 1 to 500.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule name. You can use this parameter to query the rule that matches the specified name. This parameter takes effect only when functionName is specified.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
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

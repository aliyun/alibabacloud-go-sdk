// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaseInsensitive(v string) *UpdateCacheTagRequest
	GetCaseInsensitive() *string
	SetSiteId(v int64) *UpdateCacheTagRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *UpdateCacheTagRequest
	GetSiteVersion() *int32
	SetTagName(v string) *UpdateCacheTagRequest
	GetTagName() *string
}

type UpdateCacheTagRequest struct {
	// Specifies whether to ignore case sensitivity. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	CaseInsensitive *string `json:"CaseInsensitive,omitempty" xml:"CaseInsensitive,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the website configurations. You can use this parameter to specify a version of your website to apply the feature settings. By default, version 0 is used.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// The name of the custom cache tag.
	//
	// example:
	//
	// example
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s UpdateCacheTagRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheTagRequest) GoString() string {
	return s.String()
}

func (s *UpdateCacheTagRequest) GetCaseInsensitive() *string {
	return s.CaseInsensitive
}

func (s *UpdateCacheTagRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateCacheTagRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *UpdateCacheTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *UpdateCacheTagRequest) SetCaseInsensitive(v string) *UpdateCacheTagRequest {
	s.CaseInsensitive = &v
	return s
}

func (s *UpdateCacheTagRequest) SetSiteId(v int64) *UpdateCacheTagRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateCacheTagRequest) SetSiteVersion(v int32) *UpdateCacheTagRequest {
	s.SiteVersion = &v
	return s
}

func (s *UpdateCacheTagRequest) SetTagName(v string) *UpdateCacheTagRequest {
	s.TagName = &v
	return s
}

func (s *UpdateCacheTagRequest) Validate() error {
	return dara.Validate(s)
}

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
	CaseInsensitive *string `json:"CaseInsensitive,omitempty" xml:"CaseInsensitive,omitempty"`
	// This parameter is required.
	SiteId      *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	TagName     *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
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

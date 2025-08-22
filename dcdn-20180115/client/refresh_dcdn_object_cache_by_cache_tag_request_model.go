// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDcdnObjectCacheByCacheTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheTag(v string) *RefreshDcdnObjectCacheByCacheTagRequest
	GetCacheTag() *string
	SetDomainName(v string) *RefreshDcdnObjectCacheByCacheTagRequest
	GetDomainName() *string
	SetForce(v bool) *RefreshDcdnObjectCacheByCacheTagRequest
	GetForce() *bool
}

type RefreshDcdnObjectCacheByCacheTagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tag1,tag2
	CacheTag *string `json:"CacheTag,omitempty" xml:"CacheTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s RefreshDcdnObjectCacheByCacheTagRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshDcdnObjectCacheByCacheTagRequest) GoString() string {
	return s.String()
}

func (s *RefreshDcdnObjectCacheByCacheTagRequest) GetCacheTag() *string {
	return s.CacheTag
}

func (s *RefreshDcdnObjectCacheByCacheTagRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *RefreshDcdnObjectCacheByCacheTagRequest) GetForce() *bool {
	return s.Force
}

func (s *RefreshDcdnObjectCacheByCacheTagRequest) SetCacheTag(v string) *RefreshDcdnObjectCacheByCacheTagRequest {
	s.CacheTag = &v
	return s
}

func (s *RefreshDcdnObjectCacheByCacheTagRequest) SetDomainName(v string) *RefreshDcdnObjectCacheByCacheTagRequest {
	s.DomainName = &v
	return s
}

func (s *RefreshDcdnObjectCacheByCacheTagRequest) SetForce(v bool) *RefreshDcdnObjectCacheByCacheTagRequest {
	s.Force = &v
	return s
}

func (s *RefreshDcdnObjectCacheByCacheTagRequest) Validate() error {
	return dara.Validate(s)
}

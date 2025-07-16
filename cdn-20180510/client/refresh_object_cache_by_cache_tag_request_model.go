// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshObjectCacheByCacheTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheTag(v string) *RefreshObjectCacheByCacheTagRequest
	GetCacheTag() *string
	SetDomainName(v string) *RefreshObjectCacheByCacheTagRequest
	GetDomainName() *string
	SetForce(v bool) *RefreshObjectCacheByCacheTagRequest
	GetForce() *bool
}

type RefreshObjectCacheByCacheTagRequest struct {
	// The tags of Cache. If multiple tags are returned, the tags are separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// tag1,tag2
	CacheTag *string `json:"CacheTag,omitempty" xml:"CacheTag,omitempty"`
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to purge all resources that you submit if the requested content is one of the resources that you submit to purge. Default value: false.
	//
	// 	- **true**: The nearest POP fetches all resources from the origin server, delivers them to the client, and updates the cache with the new version.
	//
	// 	- **false**: The nearest POP checks the Last-Modified parameter of the resource on the origin server. If the parameter value is the same as that of the cached resource, the POP serves the cached resource. If the parameter value is not the same as that of the cached resource, the POP fetches the latest version from the origin server, delivers it to the client, and updates the cache with the new version.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s RefreshObjectCacheByCacheTagRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshObjectCacheByCacheTagRequest) GoString() string {
	return s.String()
}

func (s *RefreshObjectCacheByCacheTagRequest) GetCacheTag() *string {
	return s.CacheTag
}

func (s *RefreshObjectCacheByCacheTagRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *RefreshObjectCacheByCacheTagRequest) GetForce() *bool {
	return s.Force
}

func (s *RefreshObjectCacheByCacheTagRequest) SetCacheTag(v string) *RefreshObjectCacheByCacheTagRequest {
	s.CacheTag = &v
	return s
}

func (s *RefreshObjectCacheByCacheTagRequest) SetDomainName(v string) *RefreshObjectCacheByCacheTagRequest {
	s.DomainName = &v
	return s
}

func (s *RefreshObjectCacheByCacheTagRequest) SetForce(v bool) *RefreshObjectCacheByCacheTagRequest {
	s.Force = &v
	return s
}

func (s *RefreshObjectCacheByCacheTagRequest) Validate() error {
	return dara.Validate(s)
}

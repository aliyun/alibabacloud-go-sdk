// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshErObjectCachesByCacheTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheTag(v string) *RefreshErObjectCachesByCacheTagRequest
	GetCacheTag() *string
	SetDomain(v string) *RefreshErObjectCachesByCacheTagRequest
	GetDomain() *string
	SetForce(v bool) *RefreshErObjectCachesByCacheTagRequest
	GetForce() *bool
	SetMergeDomainName(v string) *RefreshErObjectCachesByCacheTagRequest
	GetMergeDomainName() *string
}

type RefreshErObjectCachesByCacheTagRequest struct {
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
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// example:
	//
	// a.test.com,b.test.com
	MergeDomainName *string `json:"MergeDomainName,omitempty" xml:"MergeDomainName,omitempty"`
}

func (s RefreshErObjectCachesByCacheTagRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshErObjectCachesByCacheTagRequest) GoString() string {
	return s.String()
}

func (s *RefreshErObjectCachesByCacheTagRequest) GetCacheTag() *string {
	return s.CacheTag
}

func (s *RefreshErObjectCachesByCacheTagRequest) GetDomain() *string {
	return s.Domain
}

func (s *RefreshErObjectCachesByCacheTagRequest) GetForce() *bool {
	return s.Force
}

func (s *RefreshErObjectCachesByCacheTagRequest) GetMergeDomainName() *string {
	return s.MergeDomainName
}

func (s *RefreshErObjectCachesByCacheTagRequest) SetCacheTag(v string) *RefreshErObjectCachesByCacheTagRequest {
	s.CacheTag = &v
	return s
}

func (s *RefreshErObjectCachesByCacheTagRequest) SetDomain(v string) *RefreshErObjectCachesByCacheTagRequest {
	s.Domain = &v
	return s
}

func (s *RefreshErObjectCachesByCacheTagRequest) SetForce(v bool) *RefreshErObjectCachesByCacheTagRequest {
	s.Force = &v
	return s
}

func (s *RefreshErObjectCachesByCacheTagRequest) SetMergeDomainName(v string) *RefreshErObjectCachesByCacheTagRequest {
	s.MergeDomainName = &v
	return s
}

func (s *RefreshErObjectCachesByCacheTagRequest) Validate() error {
	return dara.Validate(s)
}

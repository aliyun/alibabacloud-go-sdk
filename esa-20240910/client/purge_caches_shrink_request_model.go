// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeCachesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentShrink(v string) *PurgeCachesShrinkRequest
	GetContentShrink() *string
	SetEdgeComputePurge(v bool) *PurgeCachesShrinkRequest
	GetEdgeComputePurge() *bool
	SetForce(v bool) *PurgeCachesShrinkRequest
	GetForce() *bool
	SetSiteId(v int64) *PurgeCachesShrinkRequest
	GetSiteId() *int64
	SetType(v string) *PurgeCachesShrinkRequest
	GetType() *string
}

type PurgeCachesShrinkRequest struct {
	// Content to be refreshed.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Used for refreshing cached resources in edge computing, such as allowing the refresh of content cached using the CacheAPI interface of an edge function.
	//
	// example:
	//
	// true
	EdgeComputePurge *bool `json:"EdgeComputePurge,omitempty" xml:"EdgeComputePurge,omitempty"`
	// Indicates whether to refresh all resources under the directory when the content from the origin and the source resource are inconsistent. The default is false.
	//
	// - **true**: Refreshes all resources under the specified directory.
	//
	// - **false**: Refreshes only the changed resources under the specified directory.
	//
	// >
	//
	// >  Applies to: Directory refresh, cachetag refresh, ignoreParams refresh, hostname refresh, and purge all cache of the site.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The type of refresh task. Possible values:
	//
	// - **file*	- (default): File refresh.
	//
	// - **cachekey**: Cachekey refresh.
	//
	// - **cachetag**: Cachetag refresh.
	//
	// - **directory**: Directory refresh.
	//
	// - **ignoreParams**: Ignore parameters refresh. Ignoring parameters means removing the ? and everything after it in the request URL. When performing an ignore parameters refresh, the user first submits the URL without parameters through the interface. The submitted URLs to be refreshed will then be matched against the cached resource URLs with the parameters removed. If the cached resource URL, after removing the parameters, matches the URL to be refreshed, the CDN node will refresh the cached resources.
	//
	// - **hostname**: Hostname refresh.
	//
	// - **purgeall**: Purge all cache under the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PurgeCachesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PurgeCachesShrinkRequest) GoString() string {
	return s.String()
}

func (s *PurgeCachesShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *PurgeCachesShrinkRequest) GetEdgeComputePurge() *bool {
	return s.EdgeComputePurge
}

func (s *PurgeCachesShrinkRequest) GetForce() *bool {
	return s.Force
}

func (s *PurgeCachesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *PurgeCachesShrinkRequest) GetType() *string {
	return s.Type
}

func (s *PurgeCachesShrinkRequest) SetContentShrink(v string) *PurgeCachesShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *PurgeCachesShrinkRequest) SetEdgeComputePurge(v bool) *PurgeCachesShrinkRequest {
	s.EdgeComputePurge = &v
	return s
}

func (s *PurgeCachesShrinkRequest) SetForce(v bool) *PurgeCachesShrinkRequest {
	s.Force = &v
	return s
}

func (s *PurgeCachesShrinkRequest) SetSiteId(v int64) *PurgeCachesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *PurgeCachesShrinkRequest) SetType(v string) *PurgeCachesShrinkRequest {
	s.Type = &v
	return s
}

func (s *PurgeCachesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

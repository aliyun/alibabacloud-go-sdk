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
	// The refresh content.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Specifies whether to refresh edge computing cached resources. For example, this allows you to refresh content cached by the Edge Routine CacheAPI API operation using the edge function.
	//
	// example:
	//
	// true
	EdgeComputePurge *bool `json:"EdgeComputePurge,omitempty" xml:"EdgeComputePurge,omitempty"`
	// Specifies whether to refresh all resources under the corresponding directory when the back-to-origin content is inconsistent with the origin server resources. Default value: false.
	//
	// - **true**: Refreshes all resources under the corresponding directory.
	//
	// - **false**: Refreshes only the changed resources under the corresponding directory.
	//
	// >
	//
	// >  This parameter takes effect for directory refresh, cache tag refresh, parameter-ignored refresh, hostname refresh, and full site refresh.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The type of the refresh node. Valid values:
	//
	// - **file*	- (default): file refresh.
	//
	// - **cachekey**: cache key refresh.
	//
	// - **cachetag**: cache label refresh.
	//
	// - **directory**: folder refresh.
	//
	// - **ignoreParams**: parameter-ignored refresh. This refers to removing the question mark (?) and all parameters after it from the request URL. When you commit a parameter-stripped URL through this API operation, the committed URL is matched against cached resource URLs after their parameters are also stripped. If a cached resource URL matches the committed URL after parameter stripping, the point of presence executes the refresh on the cached resource.
	//
	// - **hostname**: hostname refresh.
	//
	// - **purgeall**: refreshes all cached content under the site.
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

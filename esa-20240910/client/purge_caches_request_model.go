// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *PurgeCachesRequestContent) *PurgeCachesRequest
	GetContent() *PurgeCachesRequestContent
	SetEdgeComputePurge(v bool) *PurgeCachesRequest
	GetEdgeComputePurge() *bool
	SetForce(v bool) *PurgeCachesRequest
	GetForce() *bool
	SetSiteId(v int64) *PurgeCachesRequest
	GetSiteId() *int64
	SetType(v string) *PurgeCachesRequest
	GetType() *string
}

type PurgeCachesRequest struct {
	// The refresh content.
	Content *PurgeCachesRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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

func (s PurgeCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s PurgeCachesRequest) GoString() string {
	return s.String()
}

func (s *PurgeCachesRequest) GetContent() *PurgeCachesRequestContent {
	return s.Content
}

func (s *PurgeCachesRequest) GetEdgeComputePurge() *bool {
	return s.EdgeComputePurge
}

func (s *PurgeCachesRequest) GetForce() *bool {
	return s.Force
}

func (s *PurgeCachesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *PurgeCachesRequest) GetType() *string {
	return s.Type
}

func (s *PurgeCachesRequest) SetContent(v *PurgeCachesRequestContent) *PurgeCachesRequest {
	s.Content = v
	return s
}

func (s *PurgeCachesRequest) SetEdgeComputePurge(v bool) *PurgeCachesRequest {
	s.EdgeComputePurge = &v
	return s
}

func (s *PurgeCachesRequest) SetForce(v bool) *PurgeCachesRequest {
	s.Force = &v
	return s
}

func (s *PurgeCachesRequest) SetSiteId(v int64) *PurgeCachesRequest {
	s.SiteId = &v
	return s
}

func (s *PurgeCachesRequest) SetType(v string) *PurgeCachesRequest {
	s.Type = &v
	return s
}

func (s *PurgeCachesRequest) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PurgeCachesRequestContent struct {
	// The list of cache keys to refresh. This parameter is required when Type is set to cachekey.
	CacheKeys []*PurgeCachesRequestContentCacheKeys `json:"CacheKeys,omitempty" xml:"CacheKeys,omitempty" type:"Repeated"`
	// The list of cache tags to refresh. This parameter is required when Type is set to cachetag.
	CacheTags []*string `json:"CacheTags,omitempty" xml:"CacheTags,omitempty" type:"Repeated"`
	// The list of directories to refresh. This parameter is required when Type is set to directory.
	Directories []*string `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// The list of files to refresh. This parameter is required when Type is set to file.
	Files []interface{} `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The list of hostnames to refresh. This parameter is required when Type is set to hostname.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The list of files with parameters ignored. This parameter is required when Type is set to ignoreParams.
	IgnoreParams []*string `json:"IgnoreParams,omitempty" xml:"IgnoreParams,omitempty" type:"Repeated"`
	// Specifies whether to refresh the entire site. Default value: false. Set this parameter to true when Type is set to purgeall.
	//
	// example:
	//
	// true
	PurgeAll *bool `json:"PurgeAll,omitempty" xml:"PurgeAll,omitempty"`
}

func (s PurgeCachesRequestContent) String() string {
	return dara.Prettify(s)
}

func (s PurgeCachesRequestContent) GoString() string {
	return s.String()
}

func (s *PurgeCachesRequestContent) GetCacheKeys() []*PurgeCachesRequestContentCacheKeys {
	return s.CacheKeys
}

func (s *PurgeCachesRequestContent) GetCacheTags() []*string {
	return s.CacheTags
}

func (s *PurgeCachesRequestContent) GetDirectories() []*string {
	return s.Directories
}

func (s *PurgeCachesRequestContent) GetFiles() []interface{} {
	return s.Files
}

func (s *PurgeCachesRequestContent) GetHostnames() []*string {
	return s.Hostnames
}

func (s *PurgeCachesRequestContent) GetIgnoreParams() []*string {
	return s.IgnoreParams
}

func (s *PurgeCachesRequestContent) GetPurgeAll() *bool {
	return s.PurgeAll
}

func (s *PurgeCachesRequestContent) SetCacheKeys(v []*PurgeCachesRequestContentCacheKeys) *PurgeCachesRequestContent {
	s.CacheKeys = v
	return s
}

func (s *PurgeCachesRequestContent) SetCacheTags(v []*string) *PurgeCachesRequestContent {
	s.CacheTags = v
	return s
}

func (s *PurgeCachesRequestContent) SetDirectories(v []*string) *PurgeCachesRequestContent {
	s.Directories = v
	return s
}

func (s *PurgeCachesRequestContent) SetFiles(v []interface{}) *PurgeCachesRequestContent {
	s.Files = v
	return s
}

func (s *PurgeCachesRequestContent) SetHostnames(v []*string) *PurgeCachesRequestContent {
	s.Hostnames = v
	return s
}

func (s *PurgeCachesRequestContent) SetIgnoreParams(v []*string) *PurgeCachesRequestContent {
	s.IgnoreParams = v
	return s
}

func (s *PurgeCachesRequestContent) SetPurgeAll(v bool) *PurgeCachesRequestContent {
	s.PurgeAll = &v
	return s
}

func (s *PurgeCachesRequestContent) Validate() error {
	if s.CacheKeys != nil {
		for _, item := range s.CacheKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PurgeCachesRequestContentCacheKeys struct {
	// The header information corresponding to the cache key specified during the refresh. When the custom cache key feature is enabled, the cache key is generated based on the specified headers for the refresh.
	//
	// **UserGeo: country/region**
	//
	// - Country/region codes follow the ISO 3166-2 standard.
	//
	// **UserDeviceType: device type. Valid values:**
	//
	// - desktop
	//
	// - tablet
	//
	// - mobile
	//
	// **UserLanguage: language**
	//
	// - Language codes follow the ISO 639-1 or BCP 47 standard. For example, set this to zh to refresh content in Chinese.
	Headers map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The URL to refresh.
	//
	// example:
	//
	// http://a.com/1.jpg?b=1
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PurgeCachesRequestContentCacheKeys) String() string {
	return dara.Prettify(s)
}

func (s PurgeCachesRequestContentCacheKeys) GoString() string {
	return s.String()
}

func (s *PurgeCachesRequestContentCacheKeys) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurgeCachesRequestContentCacheKeys) GetUrl() *string {
	return s.Url
}

func (s *PurgeCachesRequestContentCacheKeys) SetHeaders(v map[string]*string) *PurgeCachesRequestContentCacheKeys {
	s.Headers = v
	return s
}

func (s *PurgeCachesRequestContentCacheKeys) SetUrl(v string) *PurgeCachesRequestContentCacheKeys {
	s.Url = &v
	return s
}

func (s *PurgeCachesRequestContentCacheKeys) Validate() error {
	return dara.Validate(s)
}

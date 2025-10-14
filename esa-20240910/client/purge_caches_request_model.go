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
	// Content to be refreshed.
	Content *PurgeCachesRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// List of cachekeys to be refreshed, required when the type is cachekey.
	CacheKeys []*PurgeCachesRequestContentCacheKeys `json:"CacheKeys,omitempty" xml:"CacheKeys,omitempty" type:"Repeated"`
	// List of cachetags to be refreshed, required when the type is cachetag.
	CacheTags []*string `json:"CacheTags,omitempty" xml:"CacheTags,omitempty" type:"Repeated"`
	// List of directories to be refreshed, required when the type is directory.
	Directories []*string `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// List of files to be refreshed, required when the type is file.
	Files []interface{} `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// List of hostnames to be refreshed, required when the type is hostname.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// List of files with ignored parameters, required when the type is ignoreParams.
	IgnoreParams []*string `json:"IgnoreParams,omitempty" xml:"IgnoreParams,omitempty" type:"Repeated"`
	// Flag for purging all content. Default is false, set to true when the type is purgeall.
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
	// When refreshing, specify the header information corresponding to the cache key. When the custom cache key feature switch is enabled, the cache key will be generated based on the specified header for refreshing.
	//
	// **UserGeo: Country/Region**
	//
	// - Country/region codes follow the ISO 3166-2 standard.
	//
	// **UserDeviceType: Device Type, currently there are three enum values**
	//
	// - desktop
	//
	//  - tablet
	//
	//  - mobile
	//
	// **UserLanguage: Language**
	//
	// - Language codes follow the ISO 639-1 or BCP47 standards. For example, input \\"zh\\" to refresh content in Chinese.
	Headers map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// URL address to be refreshed.
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

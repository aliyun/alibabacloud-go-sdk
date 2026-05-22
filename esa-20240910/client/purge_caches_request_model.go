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
	Content          *PurgeCachesRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	EdgeComputePurge *bool                      `json:"EdgeComputePurge,omitempty" xml:"EdgeComputePurge,omitempty"`
	Force            *bool                      `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
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
	CacheKeys    []*PurgeCachesRequestContentCacheKeys `json:"CacheKeys,omitempty" xml:"CacheKeys,omitempty" type:"Repeated"`
	CacheTags    []*string                             `json:"CacheTags,omitempty" xml:"CacheTags,omitempty" type:"Repeated"`
	Directories  []*string                             `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	Files        []interface{}                         `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	Hostnames    []*string                             `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	IgnoreParams []*string                             `json:"IgnoreParams,omitempty" xml:"IgnoreParams,omitempty" type:"Repeated"`
	PurgeAll     *bool                                 `json:"PurgeAll,omitempty" xml:"PurgeAll,omitempty"`
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
	Headers map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	Url     *string            `json:"Url,omitempty" xml:"Url,omitempty"`
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

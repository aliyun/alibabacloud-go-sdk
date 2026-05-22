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
	ContentShrink    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EdgeComputePurge *bool   `json:"EdgeComputePurge,omitempty" xml:"EdgeComputePurge,omitempty"`
	Force            *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
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

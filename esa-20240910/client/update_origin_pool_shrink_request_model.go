// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginPoolShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateOriginPoolShrinkRequest
	GetEnabled() *bool
	SetId(v int64) *UpdateOriginPoolShrinkRequest
	GetId() *int64
	SetOriginsShrink(v string) *UpdateOriginPoolShrinkRequest
	GetOriginsShrink() *string
	SetSiteId(v int64) *UpdateOriginPoolShrinkRequest
	GetSiteId() *int64
}

type UpdateOriginPoolShrinkRequest struct {
	// Whether the origin pool is enabled:
	//
	// - true: Enabled;
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the origin pool, which can be obtained by calling the [ListOriginPools](https://help.aliyun.com/document_detail/2863947.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1038520525196928
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Information about the origins added to the origin pool. Multiple origins are passed as an array.
	OriginsShrink *string `json:"Origins,omitempty" xml:"Origins,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 216558609793952
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateOriginPoolShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginPoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateOriginPoolShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateOriginPoolShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateOriginPoolShrinkRequest) GetOriginsShrink() *string {
	return s.OriginsShrink
}

func (s *UpdateOriginPoolShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateOriginPoolShrinkRequest) SetEnabled(v bool) *UpdateOriginPoolShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateOriginPoolShrinkRequest) SetId(v int64) *UpdateOriginPoolShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateOriginPoolShrinkRequest) SetOriginsShrink(v string) *UpdateOriginPoolShrinkRequest {
	s.OriginsShrink = &v
	return s
}

func (s *UpdateOriginPoolShrinkRequest) SetSiteId(v int64) *UpdateOriginPoolShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateOriginPoolShrinkRequest) Validate() error {
	return dara.Validate(s)
}

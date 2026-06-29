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
	// Specifies whether the origin address pool is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The origin address pool ID. You can call the [ListOriginPools](~~ListOriginPools~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1038520525196928
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The origin server information added to the origin address pool. Use an array to pass multiple origin servers.
	OriginsShrink *string `json:"Origins,omitempty" xml:"Origins,omitempty"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the ID.
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

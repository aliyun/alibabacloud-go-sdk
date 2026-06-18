// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginPoolShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *CreateOriginPoolShrinkRequest
	GetEnabled() *bool
	SetName(v string) *CreateOriginPoolShrinkRequest
	GetName() *string
	SetOriginsShrink(v string) *CreateOriginPoolShrinkRequest
	GetOriginsShrink() *string
	SetSiteId(v int64) *CreateOriginPoolShrinkRequest
	GetSiteId() *int64
}

type CreateOriginPoolShrinkRequest struct {
	// Specifies whether the origin pool is enabled.
	//
	// - `true`: enabled
	//
	// - `false`: disabled
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the origin pool. The name must be unique within a site.
	//
	// This parameter is required.
	//
	// example:
	//
	// pool1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of origins to add to the origin pool. Use an array to specify multiple origins.
	OriginsShrink *string `json:"Origins,omitempty" xml:"Origins,omitempty"`
	// The site ID. To obtain this ID, call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21655860979****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s CreateOriginPoolShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginPoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOriginPoolShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateOriginPoolShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateOriginPoolShrinkRequest) GetOriginsShrink() *string {
	return s.OriginsShrink
}

func (s *CreateOriginPoolShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateOriginPoolShrinkRequest) SetEnabled(v bool) *CreateOriginPoolShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *CreateOriginPoolShrinkRequest) SetName(v string) *CreateOriginPoolShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateOriginPoolShrinkRequest) SetOriginsShrink(v string) *CreateOriginPoolShrinkRequest {
	s.OriginsShrink = &v
	return s
}

func (s *CreateOriginPoolShrinkRequest) SetSiteId(v int64) *CreateOriginPoolShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateOriginPoolShrinkRequest) Validate() error {
	return dara.Validate(s)
}

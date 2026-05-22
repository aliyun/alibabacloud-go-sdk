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
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OriginsShrink *string `json:"Origins,omitempty" xml:"Origins,omitempty"`
	// This parameter is required.
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

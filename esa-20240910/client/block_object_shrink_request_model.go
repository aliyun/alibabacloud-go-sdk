// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlockObjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentShrink(v string) *BlockObjectShrinkRequest
	GetContentShrink() *string
	SetMaxage(v int32) *BlockObjectShrinkRequest
	GetMaxage() *int32
	SetSiteId(v int64) *BlockObjectShrinkRequest
	GetSiteId() *int64
	SetType(v string) *BlockObjectShrinkRequest
	GetType() *string
}

type BlockObjectShrinkRequest struct {
	// This parameter is required.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Maxage        *int32  `json:"Maxage,omitempty" xml:"Maxage,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BlockObjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BlockObjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *BlockObjectShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *BlockObjectShrinkRequest) GetMaxage() *int32 {
	return s.Maxage
}

func (s *BlockObjectShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *BlockObjectShrinkRequest) GetType() *string {
	return s.Type
}

func (s *BlockObjectShrinkRequest) SetContentShrink(v string) *BlockObjectShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *BlockObjectShrinkRequest) SetMaxage(v int32) *BlockObjectShrinkRequest {
	s.Maxage = &v
	return s
}

func (s *BlockObjectShrinkRequest) SetSiteId(v int64) *BlockObjectShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *BlockObjectShrinkRequest) SetType(v string) *BlockObjectShrinkRequest {
	s.Type = &v
	return s
}

func (s *BlockObjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}

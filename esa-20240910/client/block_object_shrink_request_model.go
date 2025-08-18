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
	SetExtension(v string) *BlockObjectShrinkRequest
	GetExtension() *string
	SetMaxage(v int32) *BlockObjectShrinkRequest
	GetMaxage() *int32
	SetSiteId(v int64) *BlockObjectShrinkRequest
	GetSiteId() *int64
	SetType(v string) *BlockObjectShrinkRequest
	GetType() *string
}

type BlockObjectShrinkRequest struct {
	// The content to block.
	//
	// This parameter is required.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The blocking period that you can extend. Set the value to 2year.
	//
	// example:
	//
	// 2year
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The period of time during which the URL is blocked. Unit: seconds. Specify this parameter if Type is set to block.
	//
	// example:
	//
	// 864000
	Maxage *int32 `json:"Maxage,omitempty" xml:"Maxage,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// BlockObject
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The type. Valid values:
	//
	// 	- **block**
	//
	// 	- **unblock**
	//
	// This parameter is required.
	//
	// example:
	//
	// block
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

func (s *BlockObjectShrinkRequest) GetExtension() *string {
	return s.Extension
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

func (s *BlockObjectShrinkRequest) SetExtension(v string) *BlockObjectShrinkRequest {
	s.Extension = &v
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

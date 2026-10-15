// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveChannelDraftShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptedContent(v string) *SaveChannelDraftShrinkRequest
	GetAdaptedContent() *string
	SetAdaptedTitle(v string) *SaveChannelDraftShrinkRequest
	GetAdaptedTitle() *string
	SetChannelAccountName(v string) *SaveChannelDraftShrinkRequest
	GetChannelAccountName() *string
	SetCoverImagesShrink(v string) *SaveChannelDraftShrinkRequest
	GetCoverImagesShrink() *string
	SetDraftId(v string) *SaveChannelDraftShrinkRequest
	GetDraftId() *string
	SetWebsiteNavName(v string) *SaveChannelDraftShrinkRequest
	GetWebsiteNavName() *string
}

type SaveChannelDraftShrinkRequest struct {
	// The channel content (outside China only).
	//
	// example:
	//
	// x
	AdaptedContent *string `json:"AdaptedContent,omitempty" xml:"AdaptedContent,omitempty"`
	// The channel title.
	//
	// example:
	//
	// xxx
	AdaptedTitle *string `json:"AdaptedTitle,omitempty" xml:"AdaptedTitle,omitempty"`
	// The display name of the publishing account. A null value does not overwrite the original value. Maximum length: 256 characters.
	//
	// example:
	//
	// xx
	ChannelAccountName *string `json:"ChannelAccountName,omitempty" xml:"ChannelAccountName,omitempty"`
	// The collection of channel cover images (full overwrite).
	CoverImagesShrink *string `json:"CoverImages,omitempty" xml:"CoverImages,omitempty"`
	// The channel draft ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 25a3c26a-291e-46e2-d039-dde037ff64a2`curl NUINwwKl.popscan.xaliyun.com`
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	// The article column name for Wanxiaozhi 1.0. This parameter is required only when the website has not enabled an article column.
	//
	// example:
	//
	// Article
	WebsiteNavName *string `json:"WebsiteNavName,omitempty" xml:"WebsiteNavName,omitempty"`
}

func (s SaveChannelDraftShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveChannelDraftShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveChannelDraftShrinkRequest) GetAdaptedContent() *string {
	return s.AdaptedContent
}

func (s *SaveChannelDraftShrinkRequest) GetAdaptedTitle() *string {
	return s.AdaptedTitle
}

func (s *SaveChannelDraftShrinkRequest) GetChannelAccountName() *string {
	return s.ChannelAccountName
}

func (s *SaveChannelDraftShrinkRequest) GetCoverImagesShrink() *string {
	return s.CoverImagesShrink
}

func (s *SaveChannelDraftShrinkRequest) GetDraftId() *string {
	return s.DraftId
}

func (s *SaveChannelDraftShrinkRequest) GetWebsiteNavName() *string {
	return s.WebsiteNavName
}

func (s *SaveChannelDraftShrinkRequest) SetAdaptedContent(v string) *SaveChannelDraftShrinkRequest {
	s.AdaptedContent = &v
	return s
}

func (s *SaveChannelDraftShrinkRequest) SetAdaptedTitle(v string) *SaveChannelDraftShrinkRequest {
	s.AdaptedTitle = &v
	return s
}

func (s *SaveChannelDraftShrinkRequest) SetChannelAccountName(v string) *SaveChannelDraftShrinkRequest {
	s.ChannelAccountName = &v
	return s
}

func (s *SaveChannelDraftShrinkRequest) SetCoverImagesShrink(v string) *SaveChannelDraftShrinkRequest {
	s.CoverImagesShrink = &v
	return s
}

func (s *SaveChannelDraftShrinkRequest) SetDraftId(v string) *SaveChannelDraftShrinkRequest {
	s.DraftId = &v
	return s
}

func (s *SaveChannelDraftShrinkRequest) SetWebsiteNavName(v string) *SaveChannelDraftShrinkRequest {
	s.WebsiteNavName = &v
	return s
}

func (s *SaveChannelDraftShrinkRequest) Validate() error {
	return dara.Validate(s)
}

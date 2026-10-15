// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveChannelDraftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptedContent(v string) *SaveChannelDraftRequest
	GetAdaptedContent() *string
	SetAdaptedTitle(v string) *SaveChannelDraftRequest
	GetAdaptedTitle() *string
	SetChannelAccountName(v string) *SaveChannelDraftRequest
	GetChannelAccountName() *string
	SetCoverImages(v []*SaveChannelDraftRequestCoverImages) *SaveChannelDraftRequest
	GetCoverImages() []*SaveChannelDraftRequestCoverImages
	SetDraftId(v string) *SaveChannelDraftRequest
	GetDraftId() *string
	SetWebsiteNavName(v string) *SaveChannelDraftRequest
	GetWebsiteNavName() *string
}

type SaveChannelDraftRequest struct {
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
	CoverImages []*SaveChannelDraftRequestCoverImages `json:"CoverImages,omitempty" xml:"CoverImages,omitempty" type:"Repeated"`
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

func (s SaveChannelDraftRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveChannelDraftRequest) GoString() string {
	return s.String()
}

func (s *SaveChannelDraftRequest) GetAdaptedContent() *string {
	return s.AdaptedContent
}

func (s *SaveChannelDraftRequest) GetAdaptedTitle() *string {
	return s.AdaptedTitle
}

func (s *SaveChannelDraftRequest) GetChannelAccountName() *string {
	return s.ChannelAccountName
}

func (s *SaveChannelDraftRequest) GetCoverImages() []*SaveChannelDraftRequestCoverImages {
	return s.CoverImages
}

func (s *SaveChannelDraftRequest) GetDraftId() *string {
	return s.DraftId
}

func (s *SaveChannelDraftRequest) GetWebsiteNavName() *string {
	return s.WebsiteNavName
}

func (s *SaveChannelDraftRequest) SetAdaptedContent(v string) *SaveChannelDraftRequest {
	s.AdaptedContent = &v
	return s
}

func (s *SaveChannelDraftRequest) SetAdaptedTitle(v string) *SaveChannelDraftRequest {
	s.AdaptedTitle = &v
	return s
}

func (s *SaveChannelDraftRequest) SetChannelAccountName(v string) *SaveChannelDraftRequest {
	s.ChannelAccountName = &v
	return s
}

func (s *SaveChannelDraftRequest) SetCoverImages(v []*SaveChannelDraftRequestCoverImages) *SaveChannelDraftRequest {
	s.CoverImages = v
	return s
}

func (s *SaveChannelDraftRequest) SetDraftId(v string) *SaveChannelDraftRequest {
	s.DraftId = &v
	return s
}

func (s *SaveChannelDraftRequest) SetWebsiteNavName(v string) *SaveChannelDraftRequest {
	s.WebsiteNavName = &v
	return s
}

func (s *SaveChannelDraftRequest) Validate() error {
	if s.CoverImages != nil {
		for _, item := range s.CoverImages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveChannelDraftRequestCoverImages struct {
	// The image URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://campus-retail-prod.oss-cn-hangzhou.aliyuncs.com/d1c5efed9b3a419bb54b75056d54dae6.jpg?Expires=4928263770&OSSAccessKeyId=LTAI5tLYirCpzCVXoArh1VYt&Signature=******************************
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The material center file ID (optional).
	//
	// example:
	//
	// xxx
	MaterialFileId *string `json:"MaterialFileId,omitempty" xml:"MaterialFileId,omitempty"`
	// The sort order.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ascending
	SortOrder *int32 `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s SaveChannelDraftRequestCoverImages) String() string {
	return dara.Prettify(s)
}

func (s SaveChannelDraftRequestCoverImages) GoString() string {
	return s.String()
}

func (s *SaveChannelDraftRequestCoverImages) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SaveChannelDraftRequestCoverImages) GetMaterialFileId() *string {
	return s.MaterialFileId
}

func (s *SaveChannelDraftRequestCoverImages) GetSortOrder() *int32 {
	return s.SortOrder
}

func (s *SaveChannelDraftRequestCoverImages) SetImageUrl(v string) *SaveChannelDraftRequestCoverImages {
	s.ImageUrl = &v
	return s
}

func (s *SaveChannelDraftRequestCoverImages) SetMaterialFileId(v string) *SaveChannelDraftRequestCoverImages {
	s.MaterialFileId = &v
	return s
}

func (s *SaveChannelDraftRequestCoverImages) SetSortOrder(v int32) *SaveChannelDraftRequestCoverImages {
	s.SortOrder = &v
	return s
}

func (s *SaveChannelDraftRequestCoverImages) Validate() error {
	return dara.Validate(s)
}

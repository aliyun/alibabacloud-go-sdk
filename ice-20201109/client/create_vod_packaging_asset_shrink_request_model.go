// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodPackagingAssetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetName(v string) *CreateVodPackagingAssetShrinkRequest
	GetAssetName() *string
	SetContentId(v string) *CreateVodPackagingAssetShrinkRequest
	GetContentId() *string
	SetDescription(v string) *CreateVodPackagingAssetShrinkRequest
	GetDescription() *string
	SetGroupName(v string) *CreateVodPackagingAssetShrinkRequest
	GetGroupName() *string
	SetInputShrink(v string) *CreateVodPackagingAssetShrinkRequest
	GetInputShrink() *string
}

type CreateVodPackagingAssetShrinkRequest struct {
	// The name of the asset. The name must be unique and can be up to 128 characters in length. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// hls_3s
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The content ID in the digital rights management (DRM) system. The maximum length is 256 characters. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// movie
	ContentId *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	// The asset description.
	//
	// example:
	//
	// HLS 3 second packaging
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the packaging group.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The asset input configurations.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
}

func (s CreateVodPackagingAssetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingAssetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingAssetShrinkRequest) GetAssetName() *string {
	return s.AssetName
}

func (s *CreateVodPackagingAssetShrinkRequest) GetContentId() *string {
	return s.ContentId
}

func (s *CreateVodPackagingAssetShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVodPackagingAssetShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateVodPackagingAssetShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *CreateVodPackagingAssetShrinkRequest) SetAssetName(v string) *CreateVodPackagingAssetShrinkRequest {
	s.AssetName = &v
	return s
}

func (s *CreateVodPackagingAssetShrinkRequest) SetContentId(v string) *CreateVodPackagingAssetShrinkRequest {
	s.ContentId = &v
	return s
}

func (s *CreateVodPackagingAssetShrinkRequest) SetDescription(v string) *CreateVodPackagingAssetShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateVodPackagingAssetShrinkRequest) SetGroupName(v string) *CreateVodPackagingAssetShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateVodPackagingAssetShrinkRequest) SetInputShrink(v string) *CreateVodPackagingAssetShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *CreateVodPackagingAssetShrinkRequest) Validate() error {
	return dara.Validate(s)
}

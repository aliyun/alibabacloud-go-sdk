// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateVodPackagingAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssets(v []*BatchCreateVodPackagingAssetRequestAssets) *BatchCreateVodPackagingAssetRequest
	GetAssets() []*BatchCreateVodPackagingAssetRequestAssets
	SetGroupName(v string) *BatchCreateVodPackagingAssetRequest
	GetGroupName() *string
}

type BatchCreateVodPackagingAssetRequest struct {
	// The assets that you want to ingest.
	Assets []*BatchCreateVodPackagingAssetRequestAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// The name of the packaging group.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s BatchCreateVodPackagingAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateVodPackagingAssetRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateVodPackagingAssetRequest) GetAssets() []*BatchCreateVodPackagingAssetRequestAssets {
	return s.Assets
}

func (s *BatchCreateVodPackagingAssetRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *BatchCreateVodPackagingAssetRequest) SetAssets(v []*BatchCreateVodPackagingAssetRequestAssets) *BatchCreateVodPackagingAssetRequest {
	s.Assets = v
	return s
}

func (s *BatchCreateVodPackagingAssetRequest) SetGroupName(v string) *BatchCreateVodPackagingAssetRequest {
	s.GroupName = &v
	return s
}

func (s *BatchCreateVodPackagingAssetRequest) Validate() error {
	return dara.Validate(s)
}

type BatchCreateVodPackagingAssetRequestAssets struct {
	// The name of the asset. The name must be unique and can be up to 128 characters in length. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// 30min_movie
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The content ID in the digital rights management (DRM) system. The maximum length is 256 characters. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// movie
	ContentId *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	// The asset input configurations.
	Input *BatchCreateVodPackagingAssetRequestAssetsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s BatchCreateVodPackagingAssetRequestAssets) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateVodPackagingAssetRequestAssets) GoString() string {
	return s.String()
}

func (s *BatchCreateVodPackagingAssetRequestAssets) GetAssetName() *string {
	return s.AssetName
}

func (s *BatchCreateVodPackagingAssetRequestAssets) GetContentId() *string {
	return s.ContentId
}

func (s *BatchCreateVodPackagingAssetRequestAssets) GetInput() *BatchCreateVodPackagingAssetRequestAssetsInput {
	return s.Input
}

func (s *BatchCreateVodPackagingAssetRequestAssets) SetAssetName(v string) *BatchCreateVodPackagingAssetRequestAssets {
	s.AssetName = &v
	return s
}

func (s *BatchCreateVodPackagingAssetRequestAssets) SetContentId(v string) *BatchCreateVodPackagingAssetRequestAssets {
	s.ContentId = &v
	return s
}

func (s *BatchCreateVodPackagingAssetRequestAssets) SetInput(v *BatchCreateVodPackagingAssetRequestAssetsInput) *BatchCreateVodPackagingAssetRequestAssets {
	s.Input = v
	return s
}

func (s *BatchCreateVodPackagingAssetRequestAssets) Validate() error {
	return dara.Validate(s)
}

type BatchCreateVodPackagingAssetRequestAssetsInput struct {
	// The URL of the media file. You can only specify a M3U8 file stored in Object Storage Service (OSS).
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The input type. Only OSS is supported.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BatchCreateVodPackagingAssetRequestAssetsInput) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateVodPackagingAssetRequestAssetsInput) GoString() string {
	return s.String()
}

func (s *BatchCreateVodPackagingAssetRequestAssetsInput) GetMedia() *string {
	return s.Media
}

func (s *BatchCreateVodPackagingAssetRequestAssetsInput) GetType() *string {
	return s.Type
}

func (s *BatchCreateVodPackagingAssetRequestAssetsInput) SetMedia(v string) *BatchCreateVodPackagingAssetRequestAssetsInput {
	s.Media = &v
	return s
}

func (s *BatchCreateVodPackagingAssetRequestAssetsInput) SetType(v string) *BatchCreateVodPackagingAssetRequestAssetsInput {
	s.Type = &v
	return s
}

func (s *BatchCreateVodPackagingAssetRequestAssetsInput) Validate() error {
	return dara.Validate(s)
}

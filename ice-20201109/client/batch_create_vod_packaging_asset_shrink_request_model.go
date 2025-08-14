// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateVodPackagingAssetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsShrink(v string) *BatchCreateVodPackagingAssetShrinkRequest
	GetAssetsShrink() *string
	SetGroupName(v string) *BatchCreateVodPackagingAssetShrinkRequest
	GetGroupName() *string
}

type BatchCreateVodPackagingAssetShrinkRequest struct {
	// The assets that you want to ingest.
	AssetsShrink *string `json:"Assets,omitempty" xml:"Assets,omitempty"`
	// The name of the packaging group.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s BatchCreateVodPackagingAssetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateVodPackagingAssetShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateVodPackagingAssetShrinkRequest) GetAssetsShrink() *string {
	return s.AssetsShrink
}

func (s *BatchCreateVodPackagingAssetShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *BatchCreateVodPackagingAssetShrinkRequest) SetAssetsShrink(v string) *BatchCreateVodPackagingAssetShrinkRequest {
	s.AssetsShrink = &v
	return s
}

func (s *BatchCreateVodPackagingAssetShrinkRequest) SetGroupName(v string) *BatchCreateVodPackagingAssetShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *BatchCreateVodPackagingAssetShrinkRequest) Validate() error {
	return dara.Validate(s)
}

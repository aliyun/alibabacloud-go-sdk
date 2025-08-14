// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodPackagingAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetName(v string) *DeleteVodPackagingAssetRequest
	GetAssetName() *string
}

type DeleteVodPackagingAssetRequest struct {
	// The name of the VOD packaging asset.
	//
	// example:
	//
	// 30min_movie
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
}

func (s DeleteVodPackagingAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodPackagingAssetRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodPackagingAssetRequest) GetAssetName() *string {
	return s.AssetName
}

func (s *DeleteVodPackagingAssetRequest) SetAssetName(v string) *DeleteVodPackagingAssetRequest {
	s.AssetName = &v
	return s
}

func (s *DeleteVodPackagingAssetRequest) Validate() error {
	return dara.Validate(s)
}

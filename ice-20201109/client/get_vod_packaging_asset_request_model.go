// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodPackagingAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetName(v string) *GetVodPackagingAssetRequest
	GetAssetName() *string
}

type GetVodPackagingAssetRequest struct {
	// The name of the VOD packaging asset.
	//
	// example:
	//
	// 30min_movie
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
}

func (s GetVodPackagingAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingAssetRequest) GoString() string {
	return s.String()
}

func (s *GetVodPackagingAssetRequest) GetAssetName() *string {
	return s.AssetName
}

func (s *GetVodPackagingAssetRequest) SetAssetName(v string) *GetVodPackagingAssetRequest {
	s.AssetName = &v
	return s
}

func (s *GetVodPackagingAssetRequest) Validate() error {
	return dara.Validate(s)
}

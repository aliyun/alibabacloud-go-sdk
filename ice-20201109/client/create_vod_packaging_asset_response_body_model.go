// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodPackagingAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsset(v *VodPackagingAsset) *CreateVodPackagingAssetResponseBody
	GetAsset() *VodPackagingAsset
	SetRequestId(v string) *CreateVodPackagingAssetResponseBody
	GetRequestId() *string
}

type CreateVodPackagingAssetResponseBody struct {
	// The information about the asset.
	Asset *VodPackagingAsset `json:"Asset,omitempty" xml:"Asset,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVodPackagingAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingAssetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingAssetResponseBody) GetAsset() *VodPackagingAsset {
	return s.Asset
}

func (s *CreateVodPackagingAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVodPackagingAssetResponseBody) SetAsset(v *VodPackagingAsset) *CreateVodPackagingAssetResponseBody {
	s.Asset = v
	return s
}

func (s *CreateVodPackagingAssetResponseBody) SetRequestId(v string) *CreateVodPackagingAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVodPackagingAssetResponseBody) Validate() error {
	if s.Asset != nil {
		if err := s.Asset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

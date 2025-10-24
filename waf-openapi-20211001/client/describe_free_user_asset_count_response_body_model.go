// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserAssetCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsset(v *DescribeFreeUserAssetCountResponseBodyAsset) *DescribeFreeUserAssetCountResponseBody
	GetAsset() *DescribeFreeUserAssetCountResponseBodyAsset
	SetRequestId(v string) *DescribeFreeUserAssetCountResponseBody
	GetRequestId() *string
}

type DescribeFreeUserAssetCountResponseBody struct {
	// The asset statistics provided by basic detection.
	Asset *DescribeFreeUserAssetCountResponseBodyAsset `json:"Asset,omitempty" xml:"Asset,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 30488BF0-FD58-52DD-B396-D014549F43A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFreeUserAssetCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserAssetCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserAssetCountResponseBody) GetAsset() *DescribeFreeUserAssetCountResponseBodyAsset {
	return s.Asset
}

func (s *DescribeFreeUserAssetCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFreeUserAssetCountResponseBody) SetAsset(v *DescribeFreeUserAssetCountResponseBodyAsset) *DescribeFreeUserAssetCountResponseBody {
	s.Asset = v
	return s
}

func (s *DescribeFreeUserAssetCountResponseBody) SetRequestId(v string) *DescribeFreeUserAssetCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFreeUserAssetCountResponseBody) Validate() error {
	if s.Asset != nil {
		if err := s.Asset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFreeUserAssetCountResponseBodyAsset struct {
	// The number of active APIs.
	//
	// example:
	//
	// 34
	AssetActive *int64 `json:"AssetActive,omitempty" xml:"AssetActive,omitempty"`
	// The total number of APIs.
	//
	// example:
	//
	// 15
	AssetCount *int64 `json:"AssetCount,omitempty" xml:"AssetCount,omitempty"`
	// The number of deactivated APIs.
	//
	// example:
	//
	// 13
	AssetOffline *int64 `json:"AssetOffline,omitempty" xml:"AssetOffline,omitempty"`
}

func (s DescribeFreeUserAssetCountResponseBodyAsset) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserAssetCountResponseBodyAsset) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserAssetCountResponseBodyAsset) GetAssetActive() *int64 {
	return s.AssetActive
}

func (s *DescribeFreeUserAssetCountResponseBodyAsset) GetAssetCount() *int64 {
	return s.AssetCount
}

func (s *DescribeFreeUserAssetCountResponseBodyAsset) GetAssetOffline() *int64 {
	return s.AssetOffline
}

func (s *DescribeFreeUserAssetCountResponseBodyAsset) SetAssetActive(v int64) *DescribeFreeUserAssetCountResponseBodyAsset {
	s.AssetActive = &v
	return s
}

func (s *DescribeFreeUserAssetCountResponseBodyAsset) SetAssetCount(v int64) *DescribeFreeUserAssetCountResponseBodyAsset {
	s.AssetCount = &v
	return s
}

func (s *DescribeFreeUserAssetCountResponseBodyAsset) SetAssetOffline(v int64) *DescribeFreeUserAssetCountResponseBodyAsset {
	s.AssetOffline = &v
	return s
}

func (s *DescribeFreeUserAssetCountResponseBodyAsset) Validate() error {
	return dara.Validate(s)
}

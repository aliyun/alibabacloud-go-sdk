// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComponentAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetId(v int64) *DeleteComponentAssetRequest
	GetAssetId() *int64
	SetLang(v string) *DeleteComponentAssetRequest
	GetLang() *string
}

type DeleteComponentAssetRequest struct {
	// The ID of the asset.
	//
	// >  You can call the [DescribeComponentAssets](~~DescribeComponentAssets~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12x
	AssetId *int64 `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DeleteComponentAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteComponentAssetRequest) GoString() string {
	return s.String()
}

func (s *DeleteComponentAssetRequest) GetAssetId() *int64 {
	return s.AssetId
}

func (s *DeleteComponentAssetRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteComponentAssetRequest) SetAssetId(v int64) *DeleteComponentAssetRequest {
	s.AssetId = &v
	return s
}

func (s *DeleteComponentAssetRequest) SetLang(v string) *DeleteComponentAssetRequest {
	s.Lang = &v
	return s
}

func (s *DeleteComponentAssetRequest) Validate() error {
	return dara.Validate(s)
}

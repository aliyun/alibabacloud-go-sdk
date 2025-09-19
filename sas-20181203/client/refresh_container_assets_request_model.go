// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshContainerAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetType(v string) *RefreshContainerAssetsRequest
	GetAssetType() *string
}

type RefreshContainerAssetsRequest struct {
	// The type of the container asset whose statistics you want to refresh. Valid values:
	//
	// 	- **IMAGE**
	//
	// 	- **CONTAINER**
	//
	// This parameter is required.
	//
	// example:
	//
	// IMAGE
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
}

func (s RefreshContainerAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshContainerAssetsRequest) GoString() string {
	return s.String()
}

func (s *RefreshContainerAssetsRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *RefreshContainerAssetsRequest) SetAssetType(v string) *RefreshContainerAssetsRequest {
	s.AssetType = &v
	return s
}

func (s *RefreshContainerAssetsRequest) Validate() error {
	return dara.Validate(s)
}

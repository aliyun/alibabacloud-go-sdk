// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssetWatchCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAssetId(v int64) *CreateAssetWatchCmd
	GetAssetId() *int64
	SetAssetType(v string) *CreateAssetWatchCmd
	GetAssetType() *string
	SetCompanyId(v int64) *CreateAssetWatchCmd
	GetCompanyId() *int64
	SetMarketId(v int64) *CreateAssetWatchCmd
	GetMarketId() *int64
}

type CreateAssetWatchCmd struct {
	AssetId   *int64  `json:"assetId,omitempty" xml:"assetId,omitempty"`
	AssetType *string `json:"assetType,omitempty" xml:"assetType,omitempty"`
	CompanyId *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	MarketId  *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
}

func (s CreateAssetWatchCmd) String() string {
	return dara.Prettify(s)
}

func (s CreateAssetWatchCmd) GoString() string {
	return s.String()
}

func (s *CreateAssetWatchCmd) GetAssetId() *int64 {
	return s.AssetId
}

func (s *CreateAssetWatchCmd) GetAssetType() *string {
	return s.AssetType
}

func (s *CreateAssetWatchCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *CreateAssetWatchCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *CreateAssetWatchCmd) SetAssetId(v int64) *CreateAssetWatchCmd {
	s.AssetId = &v
	return s
}

func (s *CreateAssetWatchCmd) SetAssetType(v string) *CreateAssetWatchCmd {
	s.AssetType = &v
	return s
}

func (s *CreateAssetWatchCmd) SetCompanyId(v int64) *CreateAssetWatchCmd {
	s.CompanyId = &v
	return s
}

func (s *CreateAssetWatchCmd) SetMarketId(v int64) *CreateAssetWatchCmd {
	s.MarketId = &v
	return s
}

func (s *CreateAssetWatchCmd) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAssetWatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetType(v string) *RemoveAssetWatchRequest
	GetAssetType() *string
	SetCompanyId(v int64) *RemoveAssetWatchRequest
	GetCompanyId() *int64
}

type RemoveAssetWatchRequest struct {
	// This parameter is required.
	AssetType *string `json:"assetType,omitempty" xml:"assetType,omitempty"`
	CompanyId *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
}

func (s RemoveAssetWatchRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveAssetWatchRequest) GoString() string {
	return s.String()
}

func (s *RemoveAssetWatchRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *RemoveAssetWatchRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *RemoveAssetWatchRequest) SetAssetType(v string) *RemoveAssetWatchRequest {
	s.AssetType = &v
	return s
}

func (s *RemoveAssetWatchRequest) SetCompanyId(v int64) *RemoveAssetWatchRequest {
	s.CompanyId = &v
	return s
}

func (s *RemoveAssetWatchRequest) Validate() error {
	return dara.Validate(s)
}

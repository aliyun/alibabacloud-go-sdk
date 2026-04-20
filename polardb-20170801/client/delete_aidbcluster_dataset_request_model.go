// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteAIDBClusterDatasetRequest
	GetDBClusterId() *string
	SetDatasetId(v string) *DeleteAIDBClusterDatasetRequest
	GetDatasetId() *string
	SetOwnerAccount(v string) *DeleteAIDBClusterDatasetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteAIDBClusterDatasetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteAIDBClusterDatasetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteAIDBClusterDatasetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAIDBClusterDatasetRequest
	GetResourceOwnerId() *int64
}

type DeleteAIDBClusterDatasetRequest struct {
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d-0npezigk8d3amuojqc
	DatasetId    *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteAIDBClusterDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAIDBClusterDatasetRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DeleteAIDBClusterDatasetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteAIDBClusterDatasetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAIDBClusterDatasetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAIDBClusterDatasetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAIDBClusterDatasetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAIDBClusterDatasetRequest) SetDBClusterId(v string) *DeleteAIDBClusterDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAIDBClusterDatasetRequest) SetDatasetId(v string) *DeleteAIDBClusterDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *DeleteAIDBClusterDatasetRequest) SetOwnerAccount(v string) *DeleteAIDBClusterDatasetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAIDBClusterDatasetRequest) SetOwnerId(v int64) *DeleteAIDBClusterDatasetRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAIDBClusterDatasetRequest) SetRegionId(v string) *DeleteAIDBClusterDatasetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAIDBClusterDatasetRequest) SetResourceOwnerAccount(v string) *DeleteAIDBClusterDatasetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAIDBClusterDatasetRequest) SetResourceOwnerId(v int64) *DeleteAIDBClusterDatasetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAIDBClusterDatasetRequest) Validate() error {
	return dara.Validate(s)
}

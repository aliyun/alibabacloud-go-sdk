// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteAIDBClusterTaskRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeleteAIDBClusterTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteAIDBClusterTaskRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteAIDBClusterTaskRequest
	GetRegionId() *string
	SetRelativeDBClusterId(v string) *DeleteAIDBClusterTaskRequest
	GetRelativeDBClusterId() *string
	SetResourceOwnerAccount(v string) *DeleteAIDBClusterTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAIDBClusterTaskRequest
	GetResourceOwnerId() *int64
}

type DeleteAIDBClusterTaskRequest struct {
	// example:
	//
	// pm-2ze9***
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// pc-2zejpr***
	RelativeDBClusterId  *string `json:"RelativeDBClusterId,omitempty" xml:"RelativeDBClusterId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteAIDBClusterTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterTaskRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAIDBClusterTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteAIDBClusterTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAIDBClusterTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAIDBClusterTaskRequest) GetRelativeDBClusterId() *string {
	return s.RelativeDBClusterId
}

func (s *DeleteAIDBClusterTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAIDBClusterTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAIDBClusterTaskRequest) SetDBClusterId(v string) *DeleteAIDBClusterTaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAIDBClusterTaskRequest) SetOwnerAccount(v string) *DeleteAIDBClusterTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAIDBClusterTaskRequest) SetOwnerId(v int64) *DeleteAIDBClusterTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAIDBClusterTaskRequest) SetRegionId(v string) *DeleteAIDBClusterTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAIDBClusterTaskRequest) SetRelativeDBClusterId(v string) *DeleteAIDBClusterTaskRequest {
	s.RelativeDBClusterId = &v
	return s
}

func (s *DeleteAIDBClusterTaskRequest) SetResourceOwnerAccount(v string) *DeleteAIDBClusterTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAIDBClusterTaskRequest) SetResourceOwnerId(v int64) *DeleteAIDBClusterTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAIDBClusterTaskRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseAITaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CloseAITaskRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CloseAITaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CloseAITaskRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CloseAITaskRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CloseAITaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloseAITaskRequest
	GetResourceOwnerId() *int64
}

type CloseAITaskRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloseAITaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseAITaskRequest) GoString() string {
	return s.String()
}

func (s *CloseAITaskRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CloseAITaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CloseAITaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloseAITaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloseAITaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloseAITaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloseAITaskRequest) SetDBClusterId(v string) *CloseAITaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *CloseAITaskRequest) SetOwnerAccount(v string) *CloseAITaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CloseAITaskRequest) SetOwnerId(v int64) *CloseAITaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CloseAITaskRequest) SetRegionId(v string) *CloseAITaskRequest {
	s.RegionId = &v
	return s
}

func (s *CloseAITaskRequest) SetResourceOwnerAccount(v string) *CloseAITaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloseAITaskRequest) SetResourceOwnerId(v int64) *CloseAITaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloseAITaskRequest) Validate() error {
	return dara.Validate(s)
}

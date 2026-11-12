// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAIDBClusterTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyAIDBClusterTaskRequest
	GetDBClusterId() *string
	SetModifyMode(v string) *ModifyAIDBClusterTaskRequest
	GetModifyMode() *string
	SetOwnerAccount(v string) *ModifyAIDBClusterTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAIDBClusterTaskRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyAIDBClusterTaskRequest
	GetRegionId() *string
	SetRelativeDBClusterId(v string) *ModifyAIDBClusterTaskRequest
	GetRelativeDBClusterId() *string
	SetResourceOwnerAccount(v string) *ModifyAIDBClusterTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAIDBClusterTaskRequest
	GetResourceOwnerId() *int64
}

type ModifyAIDBClusterTaskRequest struct {
	// The task ID. Note: This parameter refers to the training task ID, not the cluster ID. You can obtain a valid value by calling CreateAIDBClusterTask.
	//
	// example:
	//
	// pm-2ze9***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The operation to perform. Valid values:
	//
	// 	- **stop**: Stops the training. This value takes effect only for tasks that are in the training state.
	//
	// example:
	//
	// stop
	ModifyMode   *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the PolarDB database for the AI model service.
	//
	// example:
	//
	// pc-bp10ze***
	RelativeDBClusterId  *string `json:"RelativeDBClusterId,omitempty" xml:"RelativeDBClusterId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAIDBClusterTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAIDBClusterTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyAIDBClusterTaskRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAIDBClusterTaskRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyAIDBClusterTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAIDBClusterTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAIDBClusterTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAIDBClusterTaskRequest) GetRelativeDBClusterId() *string {
	return s.RelativeDBClusterId
}

func (s *ModifyAIDBClusterTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAIDBClusterTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAIDBClusterTaskRequest) SetDBClusterId(v string) *ModifyAIDBClusterTaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAIDBClusterTaskRequest) SetModifyMode(v string) *ModifyAIDBClusterTaskRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyAIDBClusterTaskRequest) SetOwnerAccount(v string) *ModifyAIDBClusterTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAIDBClusterTaskRequest) SetOwnerId(v int64) *ModifyAIDBClusterTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAIDBClusterTaskRequest) SetRegionId(v string) *ModifyAIDBClusterTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAIDBClusterTaskRequest) SetRelativeDBClusterId(v string) *ModifyAIDBClusterTaskRequest {
	s.RelativeDBClusterId = &v
	return s
}

func (s *ModifyAIDBClusterTaskRequest) SetResourceOwnerAccount(v string) *ModifyAIDBClusterTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAIDBClusterTaskRequest) SetResourceOwnerId(v int64) *ModifyAIDBClusterTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAIDBClusterTaskRequest) Validate() error {
	return dara.Validate(s)
}

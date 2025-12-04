// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMaintainTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterMaintainTimeRequest
	GetDBClusterId() *string
	SetMaintainTime(v string) *ModifyDBClusterMaintainTimeRequest
	GetMaintainTime() *string
	SetOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest
	GetResourceOwnerId() *int64
}

type ModifyDBClusterMaintainTimeRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The maintenance window of the cluster. Specify the time in the HH:mmZ-HH:mmZ format. The time must be in Coordinated Universal Time (UTC).
	//
	// For example, a value of 00:00Z-01:00Z indicates that routine maintenance can be performed on the cluster from 08:00 (UTC+8) to 09:00 (UTC+8).
	//
	// >  You can set the start time and end time of the maintenance window to the time on the hour, and the maintenance window is 1 hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00:00Z-01:00Z
	MaintainTime         *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterMaintainTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterMaintainTimeRequest) GetMaintainTime() *string {
	return s.MaintainTime
}

func (s *ModifyDBClusterMaintainTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterMaintainTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterMaintainTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterMaintainTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterMaintainTimeRequest) SetDBClusterId(v string) *ModifyDBClusterMaintainTimeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetMaintainTime(v string) *ModifyDBClusterMaintainTimeRequest {
	s.MaintainTime = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) Validate() error {
	return dara.Validate(s)
}

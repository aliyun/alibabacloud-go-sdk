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
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The maintenance window of the cluster. Specify the maintenance window in the `HH:mmZ-HH:mmZ` format. For example, the value `16:00Z-17:00Z` indicates that the cluster can be maintained from 00:00 to 01:00 (UTC+8).
	//
	//
	//
	// >  The maintenance window must start on the hour and last for an hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 02:00Z-03:00Z
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

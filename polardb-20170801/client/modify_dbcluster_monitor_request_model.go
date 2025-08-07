// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterMonitorRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyDBClusterMonitorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterMonitorRequest
	GetOwnerId() *int64
	SetPeriod(v string) *ModifyDBClusterMonitorRequest
	GetPeriod() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterMonitorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterMonitorRequest
	GetResourceOwnerId() *int64
}

type ModifyDBClusterMonitorRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The interval at which monitoring data is collected. Valid values: **5*	- and **60**. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMonitorRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMonitorRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterMonitorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterMonitorRequest) GetPeriod() *string {
	return s.Period
}

func (s *ModifyDBClusterMonitorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterMonitorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterMonitorRequest) SetDBClusterId(v string) *ModifyDBClusterMonitorRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetOwnerAccount(v string) *ModifyDBClusterMonitorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetOwnerId(v int64) *ModifyDBClusterMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetPeriod(v string) *ModifyDBClusterMonitorRequest {
	s.Period = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) Validate() error {
	return dara.Validate(s)
}

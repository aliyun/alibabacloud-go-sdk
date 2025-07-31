// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceMonitorRequest
	GetDBInstanceId() *string
	SetGranularity(v string) *ModifyDBInstanceMonitorRequest
	GetGranularity() *string
	SetOwnerAccount(v string) *ModifyDBInstanceMonitorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceMonitorRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceMonitorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceMonitorRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceMonitorRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The collection frequency of monitoring data. Valid values: **1*	- or **300**. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Granularity          *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMonitorRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMonitorRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceMonitorRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *ModifyDBInstanceMonitorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceMonitorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceMonitorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceMonitorRequest) SetDBInstanceId(v string) *ModifyDBInstanceMonitorRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetGranularity(v string) *ModifyDBInstanceMonitorRequest {
	s.Granularity = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetOwnerAccount(v string) *ModifyDBInstanceMonitorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetOwnerId(v int64) *ModifyDBInstanceMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) Validate() error {
	return dara.Validate(s)
}

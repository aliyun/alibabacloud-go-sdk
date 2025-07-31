// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMaintainTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceMaintainTimeRequest
	GetDBInstanceId() *string
	SetMaintainEndTime(v string) *ModifyDBInstanceMaintainTimeRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *ModifyDBInstanceMaintainTimeRequest
	GetMaintainStartTime() *string
	SetOwnerAccount(v string) *ModifyDBInstanceMaintainTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceMaintainTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceMaintainTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceMaintainTimeRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceMaintainTimeRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpac2345****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end time of the maintenance window. Specify the time in the ISO 8601 standard in the *HH:mm*Z format. The time must be in UTC.
	//
	// >  The end time must be later than the start time of the maintenance window.
	//
	// This parameter is required.
	//
	// example:
	//
	// 02:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. Specify the time in the ISO 8601 standard in the *HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01:00Z
	MaintainStartTime    *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetDBInstanceId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetMaintainEndTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetMaintainStartTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetOwnerAccount(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetOwnerId(v int64) *ModifyDBInstanceMaintainTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceMaintainTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) Validate() error {
	return dara.Validate(s)
}

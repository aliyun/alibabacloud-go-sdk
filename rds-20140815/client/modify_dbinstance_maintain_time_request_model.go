// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMaintainTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBInstanceMaintainTimeRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyDBInstanceMaintainTimeRequest
	GetDBInstanceId() *string
	SetMaintainTime(v string) *ModifyDBInstanceMaintainTimeRequest
	GetMaintainTime() *string
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The maintenance window of the instance. Specify the time in the *HH:mm*Z-*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22:00Z-02:00Z
	MaintainTime         *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
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

func (s *ModifyDBInstanceMaintainTimeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetMaintainTime() *string {
	return s.MaintainTime
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

func (s *ModifyDBInstanceMaintainTimeRequest) SetClientToken(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetDBInstanceId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetMaintainTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.MaintainTime = &v
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

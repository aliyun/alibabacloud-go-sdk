// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBInstanceMonitorRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyDBInstanceMonitorRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyDBInstanceMonitorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceMonitorRequest
	GetOwnerId() *int64
	SetPeriod(v string) *ModifyDBInstanceMonitorRequest
	GetPeriod() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceMonitorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceMonitorRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceMonitorRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The monitoring frequency that you want to use. Valid values:
	//
	// 	- **5**
	//
	// 	- **10**
	//
	// 	- **60**
	//
	// 	- **300**
	//
	// Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMonitorRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMonitorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBInstanceMonitorRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceMonitorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceMonitorRequest) GetPeriod() *string {
	return s.Period
}

func (s *ModifyDBInstanceMonitorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceMonitorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceMonitorRequest) SetClientToken(v string) *ModifyDBInstanceMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceMonitorRequest) SetDBInstanceId(v string) *ModifyDBInstanceMonitorRequest {
	s.DBInstanceId = &v
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

func (s *ModifyDBInstanceMonitorRequest) SetPeriod(v string) *ModifyDBInstanceMonitorRequest {
	s.Period = &v
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceHARequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *SwitchDBInstanceHARequest
	GetDBInstanceId() *string
	SetEffectiveTime(v string) *SwitchDBInstanceHARequest
	GetEffectiveTime() *string
	SetForce(v string) *SwitchDBInstanceHARequest
	GetForce() *string
	SetNodeId(v string) *SwitchDBInstanceHARequest
	GetNodeId() *string
	SetOwnerAccount(v string) *SwitchDBInstanceHARequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SwitchDBInstanceHARequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SwitchDBInstanceHARequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SwitchDBInstanceHARequest
	GetResourceOwnerId() *int64
}

type SwitchDBInstanceHARequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The time when the switching takes effect. Valid values:
	//
	// 	- **Immediate**: The switching immediately takes effect.
	//
	// 	- **MaintainTime**: The switching takes effect during the maintenance time.
	//
	// Default value: **Immediate**.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// Specifies whether to enable forcible switching. Valid values:
	//
	// 	- **Yes**
	//
	// 	- **No**
	//
	// Default value: **No**.
	//
	// example:
	//
	// No
	Force *string `json:"Force,omitempty" xml:"Force,omitempty"`
	// The secondary instance ID. You can call the DescribeDBInstanceHAConfig operation to query the secondary instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 349054
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SwitchDBInstanceHARequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceHARequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceHARequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *SwitchDBInstanceHARequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *SwitchDBInstanceHARequest) GetForce() *string {
	return s.Force
}

func (s *SwitchDBInstanceHARequest) GetNodeId() *string {
	return s.NodeId
}

func (s *SwitchDBInstanceHARequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SwitchDBInstanceHARequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SwitchDBInstanceHARequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SwitchDBInstanceHARequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchDBInstanceHARequest) SetDBInstanceId(v string) *SwitchDBInstanceHARequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetEffectiveTime(v string) *SwitchDBInstanceHARequest {
	s.EffectiveTime = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetForce(v string) *SwitchDBInstanceHARequest {
	s.Force = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetNodeId(v string) *SwitchDBInstanceHARequest {
	s.NodeId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetOwnerAccount(v string) *SwitchDBInstanceHARequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetOwnerId(v int64) *SwitchDBInstanceHARequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetResourceOwnerAccount(v string) *SwitchDBInstanceHARequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchDBInstanceHARequest) SetResourceOwnerId(v int64) *SwitchDBInstanceHARequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchDBInstanceHARequest) Validate() error {
	return dara.Validate(s)
}

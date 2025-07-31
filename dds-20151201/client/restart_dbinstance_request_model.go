// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *RestartDBInstanceRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *RestartDBInstanceRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *RestartDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartDBInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RestartDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartDBInstanceRequest
	GetResourceOwnerId() *int64
	SetSwitchMode(v string) *RestartDBInstanceRequest
	GetSwitchMode() *string
}

type RestartDBInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the shard or mongos node in the sharded cluster instance.
	//
	// > The sharded cluster instance is restarted if you do not specify this parameter.
	//
	// example:
	//
	// d-bpxxxxxxxx
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SwitchMode           *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s RestartDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestartDBInstanceRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *RestartDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartDBInstanceRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *RestartDBInstanceRequest) SetDBInstanceId(v string) *RestartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetNodeId(v string) *RestartDBInstanceRequest {
	s.NodeId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetOwnerAccount(v string) *RestartDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartDBInstanceRequest) SetOwnerId(v int64) *RestartDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetResourceOwnerAccount(v string) *RestartDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartDBInstanceRequest) SetResourceOwnerId(v int64) *RestartDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartDBInstanceRequest) SetSwitchMode(v string) *RestartDBInstanceRequest {
	s.SwitchMode = &v
	return s
}

func (s *RestartDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}

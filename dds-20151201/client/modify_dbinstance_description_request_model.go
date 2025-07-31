// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest
	GetDBInstanceDescription() *string
	SetDBInstanceId(v string) *ModifyDBInstanceDescriptionRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *ModifyDBInstanceDescriptionRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *ModifyDBInstanceDescriptionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceDescriptionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceDescriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceDescriptionRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceDescriptionRequest struct {
	// The name of the instance.
	//
	// >
	//
	// 	- The name cannot start with `http://` or `https://`.
	//
	// 	- It must start with a letter.
	//
	// 	- It must be 2 to 256 characters in length, and can contain letters, underscores (_), hyphens (-), and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdata
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// >  To modify the name of a shard or mongos node in a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp2234****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the shard or mongos node in the sharded cluster instance.
	//
	// >  This parameter is valid only if you set the **DBInstanceId*	- parameter to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bp89067****
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *ModifyDBInstanceDescriptionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceDescriptionRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ModifyDBInstanceDescriptionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceDescriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceDescriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceDescriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceId(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetNodeId(v string) *ModifyDBInstanceDescriptionRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetOwnerAccount(v string) *ModifyDBInstanceDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetOwnerId(v int64) *ModifyDBInstanceDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCharacterType(v string) *DescribeParametersRequest
	GetCharacterType() *string
	SetDBInstanceId(v string) *DescribeParametersRequest
	GetDBInstanceId() *string
	SetExtraParam(v string) *DescribeParametersRequest
	GetExtraParam() *string
	SetNodeId(v string) *DescribeParametersRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeParametersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParametersRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeParametersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParametersRequest
	GetResourceOwnerId() *int64
}

type DescribeParametersRequest struct {
	// The role of the instance. Valid values:
	//
	// 	- db: a shard node.
	//
	// 	- cs: a Configserver node.
	//
	// 	- mongos: a mongos node.
	//
	// example:
	//
	// mongos
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The instance ID.
	//
	// >  If you set this parameter to the ID of a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The parameter that is available in the future.
	//
	// example:
	//
	// terrform
	ExtraParam *string `json:"ExtraParam,omitempty" xml:"ExtraParam,omitempty"`
	// The ID of the mongos or shard node in the specified sharded cluster instance.
	//
	// >  This parameter is valid when the **DBInstanceId*	- parameter is set to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bpxxxxxxxx
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeParametersRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeParametersRequest) GetExtraParam() *string {
	return s.ExtraParam
}

func (s *DescribeParametersRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeParametersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParametersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParametersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParametersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParametersRequest) SetCharacterType(v string) *DescribeParametersRequest {
	s.CharacterType = &v
	return s
}

func (s *DescribeParametersRequest) SetDBInstanceId(v string) *DescribeParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeParametersRequest) SetExtraParam(v string) *DescribeParametersRequest {
	s.ExtraParam = &v
	return s
}

func (s *DescribeParametersRequest) SetNodeId(v string) *DescribeParametersRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeParametersRequest) SetOwnerAccount(v string) *DescribeParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParametersRequest) SetOwnerId(v int64) *DescribeParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParametersRequest) SetResourceOwnerAccount(v string) *DescribeParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParametersRequest) SetResourceOwnerId(v int64) *DescribeParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParametersRequest) Validate() error {
	return dara.Validate(s)
}

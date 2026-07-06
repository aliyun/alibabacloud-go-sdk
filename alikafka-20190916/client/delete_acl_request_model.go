// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclOperationType(v string) *DeleteAclRequest
	GetAclOperationType() *string
	SetAclOperationTypes(v string) *DeleteAclRequest
	GetAclOperationTypes() *string
	SetAclPermissionType(v string) *DeleteAclRequest
	GetAclPermissionType() *string
	SetAclResourceName(v string) *DeleteAclRequest
	GetAclResourceName() *string
	SetAclResourcePatternType(v string) *DeleteAclRequest
	GetAclResourcePatternType() *string
	SetAclResourceType(v string) *DeleteAclRequest
	GetAclResourceType() *string
	SetHost(v string) *DeleteAclRequest
	GetHost() *string
	SetInstanceId(v string) *DeleteAclRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteAclRequest
	GetRegionId() *string
	SetUsername(v string) *DeleteAclRequest
	GetUsername() *string
}

type DeleteAclRequest struct {
	// Operation type. Valid values:
	//
	// - **Write**: Write
	//
	// - **Read**: Read
	//
	// - **Describe**: Read TransactionalId
	//
	// - **IdempotentWrite**: Idempotent write to Cluster
	//
	// - **IDEMPOTENT_WRITE**: Idempotent write to Cluster, only available for Serverless instances.
	//
	// - **DESCRIBE_CONFIGS**: Query configuration, only available for Serverless instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// Write
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// Batch authorization operation types. Multiple operations are separated by commas (,).
	//
	// Valid values:
	//
	// - **Write**: Read
	//
	// - **Read**: Write
	//
	// - **Describe**: Read TransactionalId
	//
	// - **IdempotentWrite**: Idempotent write to Cluster
	//
	// - **IDEMPOTENT_WRITE**: Idempotent write to Cluster, only available for Serverless instances.
	//
	// - **DESCRIBE_CONFIGS**: Query configuration, only available for Serverless instances.
	//
	// > This parameter is only supported for Serverless instances.
	//
	// example:
	//
	// Write,Read
	AclOperationTypes *string `json:"AclOperationTypes,omitempty" xml:"AclOperationTypes,omitempty"`
	// Authorization method. Valid values:
	//
	// - DENY: Deny
	//
	// - ALLOW: Allow
	//
	// > This parameter is only supported for Serverless instances.
	//
	// example:
	//
	// DENY
	AclPermissionType *string `json:"AclPermissionType,omitempty" xml:"AclPermissionType,omitempty"`
	// Resource name.
	//
	// - Topic name or Group name.
	//
	// - Asterisk (\\*) represents all Topic or Group names.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// Matching pattern. Valid values:
	//
	// - **LITERAL**: Exact matching pattern
	//
	// - **PREFIXED**: Prefix matching pattern
	//
	// This parameter is required.
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// Resource type.
	//
	// - **Topic**: Message topic.
	//
	// - **Group**: Consumer group.
	//
	// - **Cluster**: Instance.
	//
	// - **TransactionalId**: Transaction ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Topic
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// Source IP.
	//
	// > - Only supports specific IP addresses or setting \\	- (all IPs), does not support IP segments.
	//
	// >
	//
	// > - This parameter is only supported for Serverless instances.
	//
	// example:
	//
	// *
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Username.
	//
	// This parameter is required.
	//
	// example:
	//
	// test12****
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DeleteAclRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclRequest) GoString() string {
	return s.String()
}

func (s *DeleteAclRequest) GetAclOperationType() *string {
	return s.AclOperationType
}

func (s *DeleteAclRequest) GetAclOperationTypes() *string {
	return s.AclOperationTypes
}

func (s *DeleteAclRequest) GetAclPermissionType() *string {
	return s.AclPermissionType
}

func (s *DeleteAclRequest) GetAclResourceName() *string {
	return s.AclResourceName
}

func (s *DeleteAclRequest) GetAclResourcePatternType() *string {
	return s.AclResourcePatternType
}

func (s *DeleteAclRequest) GetAclResourceType() *string {
	return s.AclResourceType
}

func (s *DeleteAclRequest) GetHost() *string {
	return s.Host
}

func (s *DeleteAclRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAclRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAclRequest) GetUsername() *string {
	return s.Username
}

func (s *DeleteAclRequest) SetAclOperationType(v string) *DeleteAclRequest {
	s.AclOperationType = &v
	return s
}

func (s *DeleteAclRequest) SetAclOperationTypes(v string) *DeleteAclRequest {
	s.AclOperationTypes = &v
	return s
}

func (s *DeleteAclRequest) SetAclPermissionType(v string) *DeleteAclRequest {
	s.AclPermissionType = &v
	return s
}

func (s *DeleteAclRequest) SetAclResourceName(v string) *DeleteAclRequest {
	s.AclResourceName = &v
	return s
}

func (s *DeleteAclRequest) SetAclResourcePatternType(v string) *DeleteAclRequest {
	s.AclResourcePatternType = &v
	return s
}

func (s *DeleteAclRequest) SetAclResourceType(v string) *DeleteAclRequest {
	s.AclResourceType = &v
	return s
}

func (s *DeleteAclRequest) SetHost(v string) *DeleteAclRequest {
	s.Host = &v
	return s
}

func (s *DeleteAclRequest) SetInstanceId(v string) *DeleteAclRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAclRequest) SetRegionId(v string) *DeleteAclRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAclRequest) SetUsername(v string) *DeleteAclRequest {
	s.Username = &v
	return s
}

func (s *DeleteAclRequest) Validate() error {
	return dara.Validate(s)
}

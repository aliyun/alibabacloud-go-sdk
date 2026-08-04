// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclOperationType(v string) *CreateAclRequest
	GetAclOperationType() *string
	SetAclOperationTypes(v string) *CreateAclRequest
	GetAclOperationTypes() *string
	SetAclPermissionType(v string) *CreateAclRequest
	GetAclPermissionType() *string
	SetAclResourceName(v string) *CreateAclRequest
	GetAclResourceName() *string
	SetAclResourcePatternType(v string) *CreateAclRequest
	GetAclResourcePatternType() *string
	SetAclResourceType(v string) *CreateAclRequest
	GetAclResourceType() *string
	SetHost(v string) *CreateAclRequest
	GetHost() *string
	SetInstanceId(v string) *CreateAclRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateAclRequest
	GetRegionId() *string
	SetUsername(v string) *CreateAclRequest
	GetUsername() *string
}

type CreateAclRequest struct {
	// The operation type. Valid values:
	//
	// - **Write**: write.
	//
	// - **Read**: read.
	//
	// - **Describe**: read TransactionalId.
	//
	// - **IdempotentWrite**: idempotent write to Cluster.
	//
	// - **IDEMPOTENT_WRITE**: idempotent write to Cluster. This value is available only for serverless instances.
	//
	// - **DESCRIBE_CONFIGS**: query configurations. This value is available only for serverless instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// Read
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// The operation types for batch authorization. Separate multiple operations with commas (,).
	//
	// Valid values:
	//
	// - **Write**: write.
	//
	// - **Read**: read.
	//
	// - **Describe**: read TransactionalId.
	//
	// - **IdempotentWrite**: idempotent write to Cluster.
	//
	// - **IDEMPOTENT_WRITE**: idempotent write to Cluster. This value is available only for serverless instances.
	//
	// - **DESCRIBE_CONFIGS**: query configurations. This value is available only for serverless instances.
	//
	//
	// > This parameter is available only for serverless instances.
	//
	// example:
	//
	// Write,Read
	AclOperationTypes *string `json:"AclOperationTypes,omitempty" xml:"AclOperationTypes,omitempty"`
	// The authorization method. Valid values:
	//
	// - **DENY**: Denied.
	//
	//
	//
	// - **ALLOW**: Allowed.
	//
	// > This parameter is available only for serverless instances.
	//
	// example:
	//
	// DENY
	AclPermissionType *string `json:"AclPermissionType,omitempty" xml:"AclPermissionType,omitempty"`
	// The resource name.
	//
	// - The name of the resource, which can be a topic name, group ID, cluster name, or transaction ID.
	//
	// - You can use an asterisk (\\*) to represent all resources of this type.
	//
	//
	// >
	//
	// > - The asterisk (\\*) returns results only after permissions are granted to all resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// X****
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The matching mode. Valid values:
	//
	// - **LITERAL**: exact match.
	//
	// - **PREFIXED**: prefix match.
	//
	// This parameter is required.
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The resource type. Valid values:
	//
	// - **Topic**: message topic.
	//
	// - **Group**: consumer group.
	//
	// - **Cluster**: instance.
	//
	// - **TransactionalId**: transaction ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Group
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The source IP address.
	//
	// >- Only specific IP addresses or an asterisk (\\*) to allow all IP addresses are supported. IP CIDR blocks are not supported.
	//
	// >- This parameter is available only for serverless instances.
	//
	// example:
	//
	// *
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The username.
	//
	// - You can use an asterisk (\\*) to represent all usernames.
	//
	//
	//
	// >
	//
	// > - The asterisk (\\*) returns results only after permissions are granted to all users.
	//
	// This parameter is required.
	//
	// example:
	//
	// test***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateAclRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAclRequest) GoString() string {
	return s.String()
}

func (s *CreateAclRequest) GetAclOperationType() *string {
	return s.AclOperationType
}

func (s *CreateAclRequest) GetAclOperationTypes() *string {
	return s.AclOperationTypes
}

func (s *CreateAclRequest) GetAclPermissionType() *string {
	return s.AclPermissionType
}

func (s *CreateAclRequest) GetAclResourceName() *string {
	return s.AclResourceName
}

func (s *CreateAclRequest) GetAclResourcePatternType() *string {
	return s.AclResourcePatternType
}

func (s *CreateAclRequest) GetAclResourceType() *string {
	return s.AclResourceType
}

func (s *CreateAclRequest) GetHost() *string {
	return s.Host
}

func (s *CreateAclRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAclRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAclRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateAclRequest) SetAclOperationType(v string) *CreateAclRequest {
	s.AclOperationType = &v
	return s
}

func (s *CreateAclRequest) SetAclOperationTypes(v string) *CreateAclRequest {
	s.AclOperationTypes = &v
	return s
}

func (s *CreateAclRequest) SetAclPermissionType(v string) *CreateAclRequest {
	s.AclPermissionType = &v
	return s
}

func (s *CreateAclRequest) SetAclResourceName(v string) *CreateAclRequest {
	s.AclResourceName = &v
	return s
}

func (s *CreateAclRequest) SetAclResourcePatternType(v string) *CreateAclRequest {
	s.AclResourcePatternType = &v
	return s
}

func (s *CreateAclRequest) SetAclResourceType(v string) *CreateAclRequest {
	s.AclResourceType = &v
	return s
}

func (s *CreateAclRequest) SetHost(v string) *CreateAclRequest {
	s.Host = &v
	return s
}

func (s *CreateAclRequest) SetInstanceId(v string) *CreateAclRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAclRequest) SetRegionId(v string) *CreateAclRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAclRequest) SetUsername(v string) *CreateAclRequest {
	s.Username = &v
	return s
}

func (s *CreateAclRequest) Validate() error {
	return dara.Validate(s)
}

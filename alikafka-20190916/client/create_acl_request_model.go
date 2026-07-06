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
	// Operation type. Valid values:
	//
	// - **Write**: write
	//
	// - **Read**: read
	//
	// - **Describe**: read TransactionalId
	//
	// - **IdempotentWrite**: idempotent write to Cluster
	//
	// - **IDEMPOTENT_WRITE**: idempotent write to Cluster, only available for Serverless instances.
	//
	// - **DESCRIBE_CONFIGS**: query configuration, only available for Serverless instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// Read
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// Batch authorization operation types. Multiple operations are separated by commas (,).
	//
	// Valid values:
	//
	// - **Write**: read
	//
	// - **Read**: write
	//
	// - **Describe**: read TransactionalId
	//
	// - **IdempotentWrite**: idempotent write to Cluster
	//
	// - **IDEMPOTENT_WRITE**: idempotent write to Cluster, only available for Serverless instances.
	//
	// - **DESCRIBE_CONFIGS**: query configuration, only available for Serverless instances.
	//
	// > This parameter is only supported for Serverless instances.
	//
	// example:
	//
	// Write,Read
	AclOperationTypes *string `json:"AclOperationTypes,omitempty" xml:"AclOperationTypes,omitempty"`
	// Authorization method. Valid values:
	//
	// - **DENY**: deny.
	//
	// - **ALLOW**: allow.
	//
	// > This parameter is only supported for Serverless instances.
	//
	// example:
	//
	// DENY
	AclPermissionType *string `json:"AclPermissionType,omitempty" xml:"AclPermissionType,omitempty"`
	// Resource name.
	//
	// - The name of the resource, which can be a topic name, Group ID, cluster name, or transaction ID.
	//
	// - You can use an asterisk (\\*) to represent all resources of this type.
	//
	// > 	- Only after authorization is granted to all resources can you query the authorized resources using an asterisk (\\*).
	//
	// This parameter is required.
	//
	// example:
	//
	// X****
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// Matching pattern. Valid values:
	//
	// - **LITERAL**: exact match
	//
	// - **PREFIXED**: prefix match
	//
	// This parameter is required.
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// Resource type. Valid values:
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
	// Source IP.
	//
	// > - Only specific IP addresses or \\	- (all IPs) are supported. IP address ranges are not supported.
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
	// - You can use an asterisk (\\*) to represent all usernames.
	//
	// > 	- Only after authorization is granted to all users can you query the authorized users using an asterisk (\\*).
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

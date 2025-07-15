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
	// The type of the operation allowed by the access control list (ACL). Valid values:
	//
	// 	- **Write**: data writes.
	//
	// 	- **Read**: data reads.
	//
	// 	- **Describe**: reads of transaction IDs.
	//
	// 	- **IdempotentWrite**: idempotent data writes to clusters.
	//
	// 	- **IDEMPOTENT_WRITE**: idempotent data writes to clusters. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// 	- **DESCRIBE_CONFIGS**: configuration query. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// Write
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// The types of operations allowed by the ACL. Separate multiple operations with commas (,).
	//
	// Valid values:
	//
	// 	- **Write**: data writes.
	//
	// 	- **Read**: data reads.
	//
	// 	- **Describe**: reads of transaction IDs.
	//
	// 	- **IdempotentWrite**: idempotent data writes to clusters.
	//
	// 	- **IDEMPOTENT_WRITE**: idempotent data writes to clusters. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// 	- **DESCRIBE_CONFIGS**: configuration query. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// >  This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// Write,Read
	AclOperationTypes *string `json:"AclOperationTypes,omitempty" xml:"AclOperationTypes,omitempty"`
	// The authorization method. Valid values:
	//
	// 	- Deny
	//
	// 	- ALLOW
	//
	// >  This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// DENY
	AclPermissionType *string `json:"AclPermissionType,omitempty" xml:"AclPermissionType,omitempty"`
	// The name of the resource.
	//
	// 	- The value can be the name of a topic or consumer group.
	//
	// 	- You can use an asterisk (\\*) to indicate the names of all topics or consumer groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The mode that is used to match resources. Valid values:
	//
	// 	- **LITERAL:*	- full match
	//
	// 	- **PREFIXED**: prefix match
	//
	// This parameter is required.
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **Topic**: topic
	//
	// 	- **Group**: consumer group
	//
	// 	- **Cluster**: cluster
	//
	// 	- **TransactionalId**: transactional ID
	//
	// This parameter is required.
	//
	// example:
	//
	// Topic
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The IP address of the source.
	//
	// >
	//
	// 	- You can specify only a specific IP address or use the asterisk (\\*) wildcard character to specify all IP addresses. CIDR blocks are not supported.
	//
	// 	- This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// *
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the user.
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

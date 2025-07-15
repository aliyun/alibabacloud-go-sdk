// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclOperationType(v string) *DescribeAclsRequest
	GetAclOperationType() *string
	SetAclPermissionType(v string) *DescribeAclsRequest
	GetAclPermissionType() *string
	SetAclResourceName(v string) *DescribeAclsRequest
	GetAclResourceName() *string
	SetAclResourcePatternType(v string) *DescribeAclsRequest
	GetAclResourcePatternType() *string
	SetAclResourceType(v string) *DescribeAclsRequest
	GetAclResourceType() *string
	SetHost(v string) *DescribeAclsRequest
	GetHost() *string
	SetInstanceId(v string) *DescribeAclsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeAclsRequest
	GetRegionId() *string
	SetUsername(v string) *DescribeAclsRequest
	GetUsername() *string
}

type DescribeAclsRequest struct {
	// The types of operations allowed by the ACL. Separate multiple operation types with commas (,).
	//
	// - Valid values:
	//
	// - Write
	//
	// - Read
	//
	// - Describe: reads of transactional IDs.
	//
	// - IdempotentWrite: idempotent data writes to clusters.
	//
	// - IDEMPOTENT_WRITE: idempotent data writes to clusters. This value is available only for ApsaraMQ for Kafka V3 instances.
	//
	// - DESCRIBE_CONFIGS: queries of configurations. This value is available only for ApsaraMQ for Kafka V3 instances.
	//
	// > This parameter is available only for ApsaraMQ for Kafka V3 serverless instances.
	//
	// example:
	//
	// Write
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// The authorization method. Valid values:
	//
	// - DENY
	//
	// - ALLOW
	//
	// > This parameter is available only for ApsaraMQ for Kafka V3 serverless instances.
	//
	// example:
	//
	// DENY
	AclPermissionType *string `json:"AclPermissionType,omitempty" xml:"AclPermissionType,omitempty"`
	// The resource name.
	//
	// 	- The value can be the name of a topic or consumer group.
	//
	// 	- You can use an asterisk (\\*) to specify the names of all topics or consumer groups.
	//
	// > You can query the resources on which permissions are granted only after you grant the user the required permissions on all resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The match mode. Valid values:
	//
	// 	- LITERAL: full-name match
	//
	// 	- PREFIXED: prefix match
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **Topic**
	//
	// 	- **Group**
	//
	// This parameter is required.
	//
	// example:
	//
	// Topic
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The source IP address.
	//
	// >-  You can specify only a specific IP address or use the asterisk (*) wildcard character to specify all IP addresses. CIDR blocks are not supported.
	//
	// > - This parameter is available only for ApsaraMQ for Kafka V3 serverless instances.
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
	// alikafka_pre-cn-v0h1cng****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The username.
	//
	// 	- You can use an asterisk (\\*) to specify all users.
	//
	// > You can use an asterisk (\\*) to query the authorized users only after you grant the required permissions to all users.
	//
	// This parameter is required.
	//
	// example:
	//
	// test12****
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeAclsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclsRequest) GetAclOperationType() *string {
	return s.AclOperationType
}

func (s *DescribeAclsRequest) GetAclPermissionType() *string {
	return s.AclPermissionType
}

func (s *DescribeAclsRequest) GetAclResourceName() *string {
	return s.AclResourceName
}

func (s *DescribeAclsRequest) GetAclResourcePatternType() *string {
	return s.AclResourcePatternType
}

func (s *DescribeAclsRequest) GetAclResourceType() *string {
	return s.AclResourceType
}

func (s *DescribeAclsRequest) GetHost() *string {
	return s.Host
}

func (s *DescribeAclsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAclsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAclsRequest) GetUsername() *string {
	return s.Username
}

func (s *DescribeAclsRequest) SetAclOperationType(v string) *DescribeAclsRequest {
	s.AclOperationType = &v
	return s
}

func (s *DescribeAclsRequest) SetAclPermissionType(v string) *DescribeAclsRequest {
	s.AclPermissionType = &v
	return s
}

func (s *DescribeAclsRequest) SetAclResourceName(v string) *DescribeAclsRequest {
	s.AclResourceName = &v
	return s
}

func (s *DescribeAclsRequest) SetAclResourcePatternType(v string) *DescribeAclsRequest {
	s.AclResourcePatternType = &v
	return s
}

func (s *DescribeAclsRequest) SetAclResourceType(v string) *DescribeAclsRequest {
	s.AclResourceType = &v
	return s
}

func (s *DescribeAclsRequest) SetHost(v string) *DescribeAclsRequest {
	s.Host = &v
	return s
}

func (s *DescribeAclsRequest) SetInstanceId(v string) *DescribeAclsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAclsRequest) SetRegionId(v string) *DescribeAclsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAclsRequest) SetUsername(v string) *DescribeAclsRequest {
	s.Username = &v
	return s
}

func (s *DescribeAclsRequest) Validate() error {
	return dara.Validate(s)
}

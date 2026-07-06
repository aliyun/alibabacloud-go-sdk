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
	// The operation type. Valid values:
	//
	// - **Write**
	//
	// - **Read**
	//
	// - **Describe**: reads a transactional ID.
	//
	// - **IdempotentWrite**: performs an idempotent write to a cluster. This value is not supported by Serverless instances. For Serverless instances, use IDEMPOTENT_WRITE.
	//
	// - **IDEMPOTENT_WRITE**: performs an idempotent write to a cluster. This value is available only for Serverless instances.
	//
	// - **DESCRIBE_CONFIGS**: queries configurations. This value is available only for Serverless instances.
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
	// > This parameter is available only for Serverless instances.
	//
	// example:
	//
	// DENY
	AclPermissionType *string `json:"AclPermissionType,omitempty" xml:"AclPermissionType,omitempty"`
	// The name of the resource.
	//
	// - The name can be a topic name or a group name.
	//
	// - You can use an asterisk (\\*) to represent all topic names or group names.
	//
	// > 	- You can use an asterisk (\\*) only after you grant permissions to all resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The match mode. Valid values:
	//
	// - LITERAL: an exact match
	//
	// - PREFIXED: a prefix match
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The type of the resource. Valid values:
	//
	// - **Topic**
	//
	// - **Group**
	//
	// - **Cluster**
	//
	// - **TransactionalId**
	//
	// This parameter is required.
	//
	// example:
	//
	// Topic
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The source IP address.
	//
	// > - You can set this parameter to a specific IP address or an asterisk (\\*). An asterisk (\\*) indicates all IP addresses. CIDR blocks are not supported.
	//
	// >
	//
	// > - This parameter is available only for Serverless instances.
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
	// - An asterisk (\\*) can be used to represent all users.
	//
	// > 	- A query with an asterisk (\\*) returns authorizations only if authorization has been granted to all users.
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

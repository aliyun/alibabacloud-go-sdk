// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAclsResponseBody
	GetCode() *int32
	SetKafkaAclList(v *DescribeAclsResponseBodyKafkaAclList) *DescribeAclsResponseBody
	GetKafkaAclList() *DescribeAclsResponseBodyKafkaAclList
	SetMessage(v string) *DescribeAclsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAclsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAclsResponseBody
	GetSuccess() *bool
}

type DescribeAclsResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The access control lists (ACLs).
	KafkaAclList *DescribeAclsResponseBodyKafkaAclList `json:"KafkaAclList,omitempty" xml:"KafkaAclList,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 46496E38-881E-4719-A2F3-F3DA6AE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAclsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAclsResponseBody) GetKafkaAclList() *DescribeAclsResponseBodyKafkaAclList {
	return s.KafkaAclList
}

func (s *DescribeAclsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAclsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAclsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAclsResponseBody) SetCode(v int32) *DescribeAclsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAclsResponseBody) SetKafkaAclList(v *DescribeAclsResponseBodyKafkaAclList) *DescribeAclsResponseBody {
	s.KafkaAclList = v
	return s
}

func (s *DescribeAclsResponseBody) SetMessage(v string) *DescribeAclsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAclsResponseBody) SetRequestId(v string) *DescribeAclsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclsResponseBody) SetSuccess(v bool) *DescribeAclsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAclsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAclsResponseBodyKafkaAclList struct {
	KafkaAclVO []*DescribeAclsResponseBodyKafkaAclListKafkaAclVO `json:"KafkaAclVO,omitempty" xml:"KafkaAclVO,omitempty" type:"Repeated"`
}

func (s DescribeAclsResponseBodyKafkaAclList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclsResponseBodyKafkaAclList) GoString() string {
	return s.String()
}

func (s *DescribeAclsResponseBodyKafkaAclList) GetKafkaAclVO() []*DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	return s.KafkaAclVO
}

func (s *DescribeAclsResponseBodyKafkaAclList) SetKafkaAclVO(v []*DescribeAclsResponseBodyKafkaAclListKafkaAclVO) *DescribeAclsResponseBodyKafkaAclList {
	s.KafkaAclVO = v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclList) Validate() error {
	return dara.Validate(s)
}

type DescribeAclsResponseBodyKafkaAclListKafkaAclVO struct {
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
	// 	- You can use the asterisk (\\*) wildcard character to specify the names of all topics or consumer groups.
	//
	// example:
	//
	// demo
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The matching mode. Valid values:
	//
	// 	- **LITERAL:*	- full-name match
	//
	// 	- **PREFIXED**: prefix match
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
	// example:
	//
	// Topic
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The host.
	//
	// example:
	//
	// ****
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The username.
	//
	// example:
	//
	// test12***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeAclsResponseBodyKafkaAclListKafkaAclVO) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclsResponseBodyKafkaAclListKafkaAclVO) GoString() string {
	return s.String()
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) GetAclOperationType() *string {
	return s.AclOperationType
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) GetAclPermissionType() *string {
	return s.AclPermissionType
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) GetAclResourceName() *string {
	return s.AclResourceName
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) GetAclResourcePatternType() *string {
	return s.AclResourcePatternType
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) GetAclResourceType() *string {
	return s.AclResourceType
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) GetHost() *string {
	return s.Host
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) GetUsername() *string {
	return s.Username
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetAclOperationType(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.AclOperationType = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetAclPermissionType(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.AclPermissionType = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetAclResourceName(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.AclResourceName = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetAclResourcePatternType(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.AclResourcePatternType = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetAclResourceType(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.AclResourceType = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetHost(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.Host = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetUsername(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.Username = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) Validate() error {
	return dara.Validate(s)
}

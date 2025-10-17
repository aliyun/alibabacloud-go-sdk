// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSubscriptionInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceEndpoint(v *ConfigureSubscriptionInstanceRequestSourceEndpoint) *ConfigureSubscriptionInstanceRequest
	GetSourceEndpoint() *ConfigureSubscriptionInstanceRequestSourceEndpoint
	SetSubscriptionDataType(v *ConfigureSubscriptionInstanceRequestSubscriptionDataType) *ConfigureSubscriptionInstanceRequest
	GetSubscriptionDataType() *ConfigureSubscriptionInstanceRequestSubscriptionDataType
	SetSubscriptionInstance(v *ConfigureSubscriptionInstanceRequestSubscriptionInstance) *ConfigureSubscriptionInstanceRequest
	GetSubscriptionInstance() *ConfigureSubscriptionInstanceRequestSubscriptionInstance
	SetAccountId(v string) *ConfigureSubscriptionInstanceRequest
	GetAccountId() *string
	SetOwnerId(v string) *ConfigureSubscriptionInstanceRequest
	GetOwnerId() *string
	SetRegionId(v string) *ConfigureSubscriptionInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ConfigureSubscriptionInstanceRequest
	GetResourceGroupId() *string
	SetSubscriptionInstanceId(v string) *ConfigureSubscriptionInstanceRequest
	GetSubscriptionInstanceId() *string
	SetSubscriptionInstanceName(v string) *ConfigureSubscriptionInstanceRequest
	GetSubscriptionInstanceName() *string
	SetSubscriptionInstanceNetworkType(v string) *ConfigureSubscriptionInstanceRequest
	GetSubscriptionInstanceNetworkType() *string
	SetSubscriptionObject(v string) *ConfigureSubscriptionInstanceRequest
	GetSubscriptionObject() *string
}

type ConfigureSubscriptionInstanceRequest struct {
	SourceEndpoint       *ConfigureSubscriptionInstanceRequestSourceEndpoint       `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	SubscriptionDataType *ConfigureSubscriptionInstanceRequestSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	SubscriptionInstance *ConfigureSubscriptionInstanceRequestSubscriptionInstance `json:"SubscriptionInstance,omitempty" xml:"SubscriptionInstance,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter is about to be discontinued.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源组ID。
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the change tracking instance. You can call the [DescribeSubscriptionInstances](https://help.aliyun.com/document_detail/49442.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtshp8n2ze4r5x****
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	// The name of the change tracking instance.
	//
	// > We recommend that you specify a descriptive name for easy identification. You do not need to use a unique name.
	//
	// example:
	//
	// MySQL Subscription
	SubscriptionInstanceName *string `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	// The network type of the change tracking instance. Set the value to **vpc**, which specifies the Virtual Private Cloud (VPC) network type.
	//
	// >
	//
	// 	- To use the new version of the change tracking feature, you must specify the SubscriptionInstanceNetworkType parameter. You must also specify the **SubscriptionInstance.VPCId*	- and **SubscriptionInstance.VSwitchID*	- parameters. If you do not specify the SubscriptionInstanceNetworkType parameter, the previous version of the change tracking feature is used.
	//
	// 	- The previous version of the change tracking feature supports self-managed MySQL databases, ApsaraDB RDS for MySQL instances, and PolarDB-X 1.0 instances. The new version of the change tracking feature supports self-managed MySQL databases, ApsaraDB RDS for MySQL instances, PolarDB for MySQL clusters, and Oracle databases.
	//
	// example:
	//
	// vpc
	SubscriptionInstanceNetworkType *string `json:"SubscriptionInstanceNetworkType,omitempty" xml:"SubscriptionInstanceNetworkType,omitempty"`
	// The objects for which you want to track data changes. The value is a JSON string and can contain regular expressions. For more information, see [SubscriptionObjects](https://help.aliyun.com/document_detail/141902.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// [{     "DBName": "dtstestdata" }]
	SubscriptionObject *string `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty"`
}

func (s ConfigureSubscriptionInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceRequest) GetSourceEndpoint() *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	return s.SourceEndpoint
}

func (s *ConfigureSubscriptionInstanceRequest) GetSubscriptionDataType() *ConfigureSubscriptionInstanceRequestSubscriptionDataType {
	return s.SubscriptionDataType
}

func (s *ConfigureSubscriptionInstanceRequest) GetSubscriptionInstance() *ConfigureSubscriptionInstanceRequestSubscriptionInstance {
	return s.SubscriptionInstance
}

func (s *ConfigureSubscriptionInstanceRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ConfigureSubscriptionInstanceRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ConfigureSubscriptionInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureSubscriptionInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigureSubscriptionInstanceRequest) GetSubscriptionInstanceId() *string {
	return s.SubscriptionInstanceId
}

func (s *ConfigureSubscriptionInstanceRequest) GetSubscriptionInstanceName() *string {
	return s.SubscriptionInstanceName
}

func (s *ConfigureSubscriptionInstanceRequest) GetSubscriptionInstanceNetworkType() *string {
	return s.SubscriptionInstanceNetworkType
}

func (s *ConfigureSubscriptionInstanceRequest) GetSubscriptionObject() *string {
	return s.SubscriptionObject
}

func (s *ConfigureSubscriptionInstanceRequest) SetSourceEndpoint(v *ConfigureSubscriptionInstanceRequestSourceEndpoint) *ConfigureSubscriptionInstanceRequest {
	s.SourceEndpoint = v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionDataType(v *ConfigureSubscriptionInstanceRequestSubscriptionDataType) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionDataType = v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionInstance(v *ConfigureSubscriptionInstanceRequestSubscriptionInstance) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionInstance = v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetAccountId(v string) *ConfigureSubscriptionInstanceRequest {
	s.AccountId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetOwnerId(v string) *ConfigureSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetRegionId(v string) *ConfigureSubscriptionInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetResourceGroupId(v string) *ConfigureSubscriptionInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionInstanceId(v string) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionInstanceName(v string) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionInstanceNetworkType(v string) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionInstanceNetworkType = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionObject(v string) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionObject = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) Validate() error {
	if s.SourceEndpoint != nil {
		if err := s.SourceEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.SubscriptionDataType != nil {
		if err := s.SubscriptionDataType.Validate(); err != nil {
			return err
		}
	}
	if s.SubscriptionInstance != nil {
		if err := s.SubscriptionInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigureSubscriptionInstanceRequestSourceEndpoint struct {
	// The name of the source database.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The endpoint of the source database.
	//
	// > This parameter is available and required only if the source database is a self-managed database.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the source instance.
	//
	// > This parameter is available and required only if the source instance is an ApsaraDB RDS for MySQL instance, a PolarDB-X 1.0 instance, or a PolarDB for MySQL cluster.
	//
	// example:
	//
	// rm-bp1zc3iyqe3qw****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the source instance. Valid values:
	//
	// 	- **RDS**: ApsaraDB RDS for MySQL instance
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster.
	//
	// 	- **LocalInstance**: self-managed database with a public IP address
	//
	// 	- **ECS**: self-managed database hosted on an Elastic Compute Service (ECS) instance
	//
	// 	- **Express**: self-managed database connected over Express Connect
	//
	// 	- **CEN**: self-managed database connected over Cloud Enterprise Network (CEN)
	//
	// 	- **dg**: self-managed database connected over Database Gateway
	//
	// > The engine of a self-managed database can be MySQL or Oracle. You must specify the engine type when you call the [CreateSubscriptionInstance](https://help.aliyun.com/document_detail/49436.html) operation.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The system ID (SID) of the Oracle database.
	//
	// > This parameter is available and required only if the source database is a self-managed Oracle database and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The ID of the Alibaba Cloud account to which the source database belongs.
	//
	// > This parameter is available and required only if you track data changes across different Alibaba Cloud accounts.
	//
	// example:
	//
	// 140692647406****
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The password of the account that is used to connect to the source database.
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The service port number of the source database.
	//
	// > This parameter is available and required only if the source database is a self-managed database.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The RAM role that is authorized to access the source database. This parameter is required if the source database does not belong to the Alibaba Cloud account that you use to configure the change tracking task. In this case, you must authorize the Alibaba Cloud account to access the source database by using a RAM role.
	//
	// > For more information about the permissions that are required for the RAM role and how to grant permissions to the RAM role, see [Configure RAM authorization for cross-account data migration and synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The username of the account that is used to connect to the source database.
	//
	// > The permissions that are required for the database account vary based on change tracking scenarios. For more information, see [Overview of change tracking scenarios](https://help.aliyun.com/document_detail/145715.html).
	//
	// example:
	//
	// dtstestaccount
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ConfigureSubscriptionInstanceRequestSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionInstanceRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) GetIP() *string {
	return s.IP
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) GetOwnerID() *string {
	return s.OwnerID
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) GetPassword() *string {
	return s.Password
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) GetRole() *string {
	return s.Role
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetDatabaseName(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetIP(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetInstanceID(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetInstanceType(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetOracleSID(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetOwnerID(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.OwnerID = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetPassword(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetPort(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetRole(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.Role = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetUserName(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type ConfigureSubscriptionInstanceRequestSubscriptionDataType struct {
	// Specifies whether to track DDL statements. Default value: true. Valid values:
	//
	// 	- **true**: tracks DDL statements.
	//
	// 	- **false**: does not track DDL statements.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DDL *bool `json:"DDL,omitempty" xml:"DDL,omitempty"`
	// Specifies whether to track DML statements. Default value: true. Valid values:
	//
	// 	- **true**: tracks DML statements.
	//
	// 	- **false**: does not track DML statements.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DML *bool `json:"DML,omitempty" xml:"DML,omitempty"`
}

func (s ConfigureSubscriptionInstanceRequestSubscriptionDataType) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionInstanceRequestSubscriptionDataType) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionDataType) GetDDL() *bool {
	return s.DDL
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionDataType) GetDML() *bool {
	return s.DML
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionDataType) SetDDL(v bool) *ConfigureSubscriptionInstanceRequestSubscriptionDataType {
	s.DDL = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionDataType) SetDML(v bool) *ConfigureSubscriptionInstanceRequestSubscriptionDataType {
	s.DML = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionDataType) Validate() error {
	return dara.Validate(s)
}

type ConfigureSubscriptionInstanceRequestSubscriptionInstance struct {
	// The ID of the VPC in which the change tracking instance is deployed.
	//
	// > This parameter is available and required only if the **SubscriptionInstanceNetworkType*	- parameter is set to **vpc**.
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch in the specified VPC.
	//
	// > This parameter is available and required only if the **SubscriptionInstanceNetworkType*	- parameter is set to **vpc**.
	//
	// example:
	//
	// vsw-bp10df3mxae6lpmku****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ConfigureSubscriptionInstanceRequestSubscriptionInstance) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionInstanceRequestSubscriptionInstance) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionInstance) GetVPCId() *string {
	return s.VPCId
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionInstance) SetVPCId(v string) *ConfigureSubscriptionInstanceRequestSubscriptionInstance {
	s.VPCId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionInstance) SetVSwitchId(v string) *ConfigureSubscriptionInstanceRequestSubscriptionInstance {
	s.VSwitchId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionInstance) Validate() error {
	return dara.Validate(s)
}

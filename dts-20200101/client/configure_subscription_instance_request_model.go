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
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because it will be deprecated.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the change tracking instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
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
	// > Specify a descriptive name for easy identification. The name does not need to be unique.
	//
	// example:
	//
	// MySQL订阅
	SubscriptionInstanceName *string `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	// The network type of the change tracking instance. The only valid value is **vpc**, which indicates a virtual private cloud (VPC).
	//
	// > - If you specify this parameter, the change tracking instance is defined as the new version. You must also correctly set the **SubscriptionInstance.VPCId*	- and **SubscriptionInstance.VSwitchID*	- parameters. If you do not specify this parameter, the change tracking instance is defined as the legacy version.
	//
	// > - The legacy version supports change tracking for self-managed MySQL, ApsaraDB RDS for MySQL, and DRDS. The new version supports change tracking for self-managed MySQL, ApsaraDB RDS for MySQL, PolarDB for MySQL, and Oracle.
	//
	// example:
	//
	// vpc
	SubscriptionInstanceNetworkType *string `json:"SubscriptionInstanceNetworkType,omitempty" xml:"SubscriptionInstanceNetworkType,omitempty"`
	// The objects to be subscribed to. The value is a JSON string that supports regular expressions. For more information, see [Subscription object configuration](https://help.aliyun.com/document_detail/141902.html).
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
	// 待订阅的数据库名称。
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// 源数据库的连接地址。
	//
	// > 当源数据库为自建数据库时，本参数才可用且必须传入。
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// 源实例ID。
	//
	// > 源数据库的实例类型为RDS MySQL、PolarDB-X 1.0、PolarDB MySQL时，本参数才可用且必须传入。
	//
	// example:
	//
	// rm-bp1zc3iyqe3qw****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// 源数据库的实例类型，取值：
	//
	// - **RDS**：RDS MySQL。
	//
	// - **PolarDB**：PolarDB MySQL。
	//
	// - **LocalInstance**：有公网IP的自建数据库。
	//
	// - **ECS**：ECS上的自建数据库。
	//
	// - **Express**：通过专线接入的自建数据库。
	//
	// - **CEN**：通过云企业网CEN接入的自建数据库。
	//
	// - **dg**：通过数据库网关接入的自建数据库。
	//
	// > 支持自建数据库的数据库类型为MySQL、Oracle，您需要提前调用[CreateSubscriptionInstance](https://help.aliyun.com/document_detail/49436.html)设置。
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Oracle数据库的SID信息。
	//
	// > 当源数据库为自建Oracle时，且Oracle数据库为非RAC实例时，本参数才可用且必须传入。
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// 源实例所属的阿里云账号ID。
	//
	// > 仅在配置跨阿里云账号的数据订阅时本参数才可用，且必须传入。
	//
	// example:
	//
	// 140692647406****
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// 源实例的数据库账号密码。
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 源数据库的服务端口。
	//
	// > 当源数据库为自建数据库时，本参数才可用且必须传入。
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 源实例的授权角色。当源实例与配置订阅任务所属阿里云账号不同时，需传入该参数，来指定源实例的授权角色，以允许配置订阅任务所属阿里云账号访问源实例的实例信息。
	//
	// > 角色所需的权限及授权方式，请参见[跨阿里云账号数据迁移或同步时如何配置RAM授权](https://help.aliyun.com/document_detail/48468.html)。
	//
	// example:
	//
	// ram-for-dts
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 源实例的数据库账号。
	//
	// > 订阅不同的数据库所需的权限有所差异，详情请参见[DTS数据订阅方案概览](https://help.aliyun.com/document_detail/145715.html)中对应的配置案例。
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
	// 是否订阅DDL类型的数据，取值：
	//
	// - **true**：是，为默认值。
	//
	// - **false**：否。
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DDL *bool `json:"DDL,omitempty" xml:"DDL,omitempty"`
	// 是否订阅DML类型的数据，取值：
	//
	// - **true**：是，为默认值。
	//
	// - **false**：否。
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
	// 订阅实例的专有网络ID。
	//
	// > 当**SubscriptionInstanceNetworkType**取值为**vpc**时，本参数才可用且必须传入。
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// 订阅实例的虚拟交换机ID。
	//
	// > 当**SubscriptionInstanceNetworkType**取值为**vpc**时，本参数才可用且必须传入。
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

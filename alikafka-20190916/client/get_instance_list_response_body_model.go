// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetInstanceListResponseBody
	GetCode() *int32
	SetInstanceList(v *GetInstanceListResponseBodyInstanceList) *GetInstanceListResponseBody
	GetInstanceList() *GetInstanceListResponseBodyInstanceList
	SetMessage(v string) *GetInstanceListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceListResponseBody
	GetSuccess() *bool
}

type GetInstanceListResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instances.
	InstanceList *GetInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// 4B6D821D-7F67-4CAA-9E13-A5A997C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetInstanceListResponseBody) GetInstanceList() *GetInstanceListResponseBodyInstanceList {
	return s.InstanceList
}

func (s *GetInstanceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceListResponseBody) SetCode(v int32) *GetInstanceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceListResponseBody) SetInstanceList(v *GetInstanceListResponseBodyInstanceList) *GetInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *GetInstanceListResponseBody) SetMessage(v string) *GetInstanceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceListResponseBody) SetRequestId(v string) *GetInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceListResponseBody) SetSuccess(v bool) *GetInstanceListResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceList struct {
	InstanceVO []*GetInstanceListResponseBodyInstanceListInstanceVO `json:"InstanceVO,omitempty" xml:"InstanceVO,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceList) GetInstanceVO() []*GetInstanceListResponseBodyInstanceListInstanceVO {
	return s.InstanceVO
}

func (s *GetInstanceListResponseBodyInstanceList) SetInstanceVO(v []*GetInstanceListResponseBodyInstanceListInstanceVO) *GetInstanceListResponseBodyInstanceList {
	s.InstanceVO = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVO struct {
	// The configurations of the deployed ApsaraMQ for Kafka instance.
	//
	// example:
	//
	// {\\"enable.vpc_sasl_ssl\\":\\"false\\",\\"kafka.log.retention.hours\\":\\"66\\",\\"enable.acl\\":\\"false\\",\\"kafka.message.max.bytes\\":\\"6291456\\"}
	AllConfig *string `json:"AllConfig,omitempty" xml:"AllConfig,omitempty"`
	// Indicates whether the flexible group creation feature is enabled.
	//
	// example:
	//
	// true
	AutoCreateGroupEnable *bool `json:"AutoCreateGroupEnable,omitempty" xml:"AutoCreateGroupEnable,omitempty"`
	// Indicates whether the automatic topic creation feature is enabled.
	//
	// example:
	//
	// true
	AutoCreateTopicEnable *bool `json:"AutoCreateTopicEnable,omitempty" xml:"AutoCreateTopicEnable,omitempty"`
	// The ID of the secondary zone.
	//
	// example:
	//
	// cn-hangzhou-a
	BackupZoneId *string `json:"BackupZoneId,omitempty" xml:"BackupZoneId,omitempty"`
	// The parameters that are returned for the ApsaraMQ for Confluent instance.
	ConfluentConfig             *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig             `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
	ConfluentInstanceComponents *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents `json:"ConfluentInstanceComponents,omitempty" xml:"ConfluentInstanceComponents,omitempty" type:"Struct"`
	// The time when the instance was created. Unit: milliseconds.
	//
	// example:
	//
	// 1577961819000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of partitions in a topic that is automatically created.
	//
	// example:
	//
	// 12
	DefaultPartitionNum *int32 `json:"DefaultPartitionNum,omitempty" xml:"DefaultPartitionNum,omitempty"`
	// The type of the network in which the instance is deployed. Valid values:
	//
	// 	- **4**: Internet and VPC
	//
	// 	- **5**: VPC
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size. Unit: GB
	//
	// example:
	//
	// 3600
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type of the instance. Valid values:
	//
	// 	- **0**: ultra disk
	//
	// 	- **1**: standard SSD
	//
	// example:
	//
	// 1
	DiskType *int32 `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The default endpoint of the instance in domain name mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// alikafka-pre-cn-zv**********-1-vpc.alikafka.aliyuncs.com:9092,alikafka-pre-cn-zv**********-2-vpc.alikafka.aliyuncs.com:9092,alikafka-pre-cn-zv**********-3-vpc.alikafka.aliyuncs.com:9092
	DomainEndpoint *string `json:"DomainEndpoint,omitempty" xml:"DomainEndpoint,omitempty"`
	// The maximum Internet traffic in the instance.
	//
	// example:
	//
	// 20
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The default endpoint of the instance in IP address mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// 192.168.XX.XX:9092,192.168.XX.XX:9092,192.168.XX.XX:9092
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// The time when the instance expires. Unit: milliseconds.
	//
	// example:
	//
	// 1893581018000
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// alikafka_pre-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum traffic in the instance.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The maximum read traffic in the instance. Unit: Mbit/s.
	//
	// example:
	//
	// 1000
	IoMaxRead *int32 `json:"IoMaxRead,omitempty" xml:"IoMaxRead,omitempty"`
	// The traffic specification.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The maximum write traffic. Unit: Mbit/s.
	//
	// example:
	//
	// 1000
	IoMaxWrite *int32 `json:"IoMaxWrite,omitempty" xml:"IoMaxWrite,omitempty"`
	// The ID of the key that is used for disk encryption in the region where the instance is deployed.
	//
	// example:
	//
	// 0d24xxxx-da7b-4786-b981-9a164dxxxxxx
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The retention period of messages in the instance. Unit: hours.
	//
	// example:
	//
	// 72
	MsgRetain *int32 `json:"MsgRetain,omitempty" xml:"MsgRetain,omitempty"`
	// The instance name.
	//
	// example:
	//
	// alikafka_post-cn-mp91gnw0****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **0**: the subscription billing method
	//
	// 	- **1**: the pay-as-you-go billing method
	//
	// 	- **3**: the pay-as-you-go billing method for serverless ApsaraMQ for Kafka V3 instances
	//
	// 	- **4**: the pay-as-you-go billing method for ApsaraMQ for Confluent instances
	//
	// example:
	//
	// 1
	PaidType                  *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	RecommendedPartitionCount *int32 `json:"RecommendedPartitionCount,omitempty" xml:"RecommendedPartitionCount,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The traffic reserved for message publishing. Unit: MB/s.
	//
	// >  This parameter is returned only if the instance is a serverless ApsaraMQ for Kafka V3 instance.
	//
	// example:
	//
	// 60
	ReservedPublishCapacity *int32 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// The traffic reserved for message subscription. Unit: MB/s.
	//
	// >  This parameter is returned only if the instance is a serverless ApsaraMQ for Kafka V3 instance.
	//
	// example:
	//
	// 60
	ReservedSubscribeCapacity *int32 `json:"ReservedSubscribeCapacity,omitempty" xml:"ReservedSubscribeCapacity,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Simple Authentication and Security Layer (SASL) endpoint of the instance in domain name mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// alikafka-pre-cn-zv**********-1-vpc.alikafka.aliyuncs.com:9094,alikafka-pre-cn-zv**********-2-vpc.alikafka.aliyuncs.com:9094,alikafka-pre-cn-zv**********-3-vpc.alikafka.aliyuncs.com:9094
	SaslDomainEndpoint *string `json:"SaslDomainEndpoint,omitempty" xml:"SaslDomainEndpoint,omitempty"`
	// The Simple Authentication and Security Layer (SASL) endpoint of the instance in IP address mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// 172.16.3.XX:9094,172.16.3.XX:9094,172.16.3.XX:9094
	SaslEndPoint *string `json:"SaslEndPoint,omitempty" xml:"SaslEndPoint,omitempty"`
	// The security group to which the instance belongs.
	//
	// 	- If the instance is deployed in the ApsaraMQ for Kafka console or by calling the [StartInstance](https://help.aliyun.com/document_detail/157786.html) operation without a security group configured, no value is returned.
	//
	// 	- If the instance is deployed by calling the [StartInstance](https://help.aliyun.com/document_detail/157786.html) operation with a security group configured, the returned value is the configured security group.
	//
	// example:
	//
	// sg-bp13wfx7kz9pkow****
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The instance version. Valid values: v2, v3, and confluent.
	//
	// example:
	//
	// v3
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// >  This parameter is out of date. We recommend that you refer to the ViewInstanceStatusCode parameter.
	//
	// The instance status. Valid values:
	//
	// 	- **0**: pending
	//
	// 	- **1**: preparing hardware resources
	//
	// 	- **2**: initializing
	//
	// 	- **3**: starting
	//
	// 	- **5**: running
	//
	// 	- **6**: migrating
	//
	// 	- **7**: ready for upgrade
	//
	// 	- **8**: upgrading
	//
	// 	- **9**: ready for change
	//
	// 	- **10**: released
	//
	// 	- **11**: changing
	//
	// 	- **15**: expired
	//
	// 	- **30**: scaling
	//
	// example:
	//
	// 5
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The instance edition. Valid values:
	//
	// 	- **professional**: Professional Edition (High Write)
	//
	// 	- **professionalForHighRead**: Professional Edition (High Read)
	//
	// 	- **normal**: Standard Edition
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The SSL endpoint of the instance in domain name mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// alikafka-pre-cn-zv**********-1.alikafka.aliyuncs.com:9093,alikafka-pre-cn-zv**********-2.alikafka.aliyuncs.com:9093,alikafka-pre-cn-zv**********-3.alikafka.aliyuncs.com:9093
	SslDomainEndpoint *string `json:"SslDomainEndpoint,omitempty" xml:"SslDomainEndpoint,omitempty"`
	// The Secure Sockets Layer (SSL) endpoint of the instance in IP address mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// 192.0.XX.XX:9093,198.51.XX.XX:9093,203.0.XX.XX:9093
	SslEndPoint *string `json:"SslEndPoint,omitempty" xml:"SslEndPoint,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-k
	StandardZoneId *string `json:"StandardZoneId,omitempty" xml:"StandardZoneId,omitempty"`
	// The tags.
	Tags *GetInstanceListResponseBodyInstanceListInstanceVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The maximum number of topics on the instance.
	//
	// example:
	//
	// 180
	TopicNumLimit *int32 `json:"TopicNumLimit,omitempty" xml:"TopicNumLimit,omitempty"`
	// The upgrade information about the instance.
	UpgradeServiceDetailInfo *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo `json:"UpgradeServiceDetailInfo,omitempty" xml:"UpgradeServiceDetailInfo,omitempty" type:"Struct"`
	// The number of used groups.
	//
	// example:
	//
	// 10
	UsedGroupCount *int32 `json:"UsedGroupCount,omitempty" xml:"UsedGroupCount,omitempty"`
	// The number of used partitions.
	//
	// example:
	//
	// 25
	UsedPartitionCount *int32 `json:"UsedPartitionCount,omitempty" xml:"UsedPartitionCount,omitempty"`
	// The number of used topics.
	//
	// example:
	//
	// 3
	UsedTopicCount *int32 `json:"UsedTopicCount,omitempty" xml:"UsedTopicCount,omitempty"`
	// The ID of the vSwitch to which the instance belongs.
	//
	// example:
	//
	// vsw-bp1fvuw0ljd7vzmo3****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The vSwitch IDs.
	VSwitchIds *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	// The instance status. The valid values are consistent with the values displayed in the ApsaraMQ for Kafka console. This parameter is used in the new version of ApsaraMQ for Kafka.
	//
	// Valid values:
	//
	// 	- **0**: pending
	//
	// 	- **1**: deploying
	//
	// 	- **2**: running
	//
	// 	- **3**: stopped
	//
	// 	- **4**: expiring
	//
	// 	- **5**: expired
	//
	// 	- **6**: released
	//
	// 	- **7**: upgrading
	//
	// 	- **8**: migrating
	//
	// 	- **21**: stopping
	//
	// 	- **22**: starting
	//
	// 	- **23**: releasing
	//
	// 	- **30**: auto scaling
	//
	// 	- **101**: deployment failed
	//
	// 	- **102**: upgrade failed
	//
	// 	- **103**: migration failed
	//
	// example:
	//
	// 2
	ViewInstanceStatusCode *int32 `json:"ViewInstanceStatusCode,omitempty" xml:"ViewInstanceStatusCode,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-bp1ojac7bv448nifj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The SSL endpoint of the instance in domain name mode. You can use the endpoint to access the instance only in virtual private clouds (VPCs). ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// alikafka-post-cn-******-1-vpc.alikafka.aliyuncs.com:9095,alikafka-post-cn-******-2-vpc.alikafka.aliyuncs.com:9095,alikafka-post-cn-******-3-vpc.alikafka.aliyuncs.com:9095
	VpcSaslDomainEndpoint *string `json:"VpcSaslDomainEndpoint,omitempty" xml:"VpcSaslDomainEndpoint,omitempty"`
	// The Secure Sockets Layer (SSL) endpoint of the instance in IP address mode. You can use the endpoint to access the instance only in virtual private clouds (VPCs). ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// 172.16.3.XX:9095,172.16.3.XX:9095,172.16.3.XX:9095
	VpcSaslEndPoint *string `json:"VpcSaslEndPoint,omitempty" xml:"VpcSaslEndPoint,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// zonei
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVO) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVO) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetAllConfig() *string {
	return s.AllConfig
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetAutoCreateGroupEnable() *bool {
	return s.AutoCreateGroupEnable
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetAutoCreateTopicEnable() *bool {
	return s.AutoCreateTopicEnable
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetBackupZoneId() *string {
	return s.BackupZoneId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetConfluentConfig() *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	return s.ConfluentConfig
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetConfluentInstanceComponents() *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents {
	return s.ConfluentInstanceComponents
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetDefaultPartitionNum() *int32 {
	return s.DefaultPartitionNum
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetDeployType() *int32 {
	return s.DeployType
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetDiskType() *int32 {
	return s.DiskType
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetDomainEndpoint() *string {
	return s.DomainEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetEipMax() *int32 {
	return s.EipMax
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetEndPoint() *string {
	return s.EndPoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetIoMax() *int32 {
	return s.IoMax
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetIoMaxRead() *int32 {
	return s.IoMaxRead
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetIoMaxWrite() *int32 {
	return s.IoMaxWrite
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetMsgRetain() *int32 {
	return s.MsgRetain
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetName() *string {
	return s.Name
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetPaidType() *int32 {
	return s.PaidType
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetRecommendedPartitionCount() *int32 {
	return s.RecommendedPartitionCount
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetReservedPublishCapacity() *int32 {
	return s.ReservedPublishCapacity
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetReservedSubscribeCapacity() *int32 {
	return s.ReservedSubscribeCapacity
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSaslDomainEndpoint() *string {
	return s.SaslDomainEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSaslEndPoint() *string {
	return s.SaslEndPoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSecurityGroup() *string {
	return s.SecurityGroup
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSeries() *string {
	return s.Series
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetServiceStatus() *int32 {
	return s.ServiceStatus
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSpecType() *string {
	return s.SpecType
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSslDomainEndpoint() *string {
	return s.SslDomainEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetSslEndPoint() *string {
	return s.SslEndPoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetStandardZoneId() *string {
	return s.StandardZoneId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetTags() *GetInstanceListResponseBodyInstanceListInstanceVOTags {
	return s.Tags
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetTopicNumLimit() *int32 {
	return s.TopicNumLimit
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetUpgradeServiceDetailInfo() *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo {
	return s.UpgradeServiceDetailInfo
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetUsedGroupCount() *int32 {
	return s.UsedGroupCount
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetUsedPartitionCount() *int32 {
	return s.UsedPartitionCount
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetUsedTopicCount() *int32 {
	return s.UsedTopicCount
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetVSwitchIds() *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds {
	return s.VSwitchIds
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetViewInstanceStatusCode() *int32 {
	return s.ViewInstanceStatusCode
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetVpcSaslDomainEndpoint() *string {
	return s.VpcSaslDomainEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetVpcSaslEndPoint() *string {
	return s.VpcSaslEndPoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetAllConfig(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.AllConfig = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetAutoCreateGroupEnable(v bool) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.AutoCreateGroupEnable = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetAutoCreateTopicEnable(v bool) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.AutoCreateTopicEnable = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetBackupZoneId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.BackupZoneId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetConfluentConfig(v *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ConfluentConfig = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetConfluentInstanceComponents(v *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ConfluentInstanceComponents = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetCreateTime(v int64) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDefaultPartitionNum(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DefaultPartitionNum = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDeployType(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DeployType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDiskSize(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DiskSize = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDiskType(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DiskType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetEipMax(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.EipMax = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.EndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetExpiredTime(v int64) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ExpiredTime = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetInstanceId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMax(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMax = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMaxRead(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMaxRead = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMaxSpec(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMaxSpec = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMaxWrite(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMaxWrite = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetKmsKeyId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.KmsKeyId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetMsgRetain(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.MsgRetain = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetName(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.Name = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetPaidType(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.PaidType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetRecommendedPartitionCount(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.RecommendedPartitionCount = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetRegionId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.RegionId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetReservedPublishCapacity(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ReservedPublishCapacity = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetReservedSubscribeCapacity(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ReservedSubscribeCapacity = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetResourceGroupId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSaslDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SaslDomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSaslEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SaslEndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSecurityGroup(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SecurityGroup = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSeries(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.Series = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetServiceStatus(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ServiceStatus = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSpecType(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SpecType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSslDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SslDomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSslEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SslEndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetStandardZoneId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.StandardZoneId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetTags(v *GetInstanceListResponseBodyInstanceListInstanceVOTags) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.Tags = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetTopicNumLimit(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.TopicNumLimit = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUpgradeServiceDetailInfo(v *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UpgradeServiceDetailInfo = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUsedGroupCount(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UsedGroupCount = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUsedPartitionCount(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UsedPartitionCount = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUsedTopicCount(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UsedTopicCount = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVSwitchId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVSwitchIds(v *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VSwitchIds = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetViewInstanceStatusCode(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ViewInstanceStatusCode = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVpcId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VpcId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVpcSaslDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VpcSaslDomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVpcSaslEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VpcSaslEndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetZoneId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ZoneId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig struct {
	// The number of CPU cores of Connect.
	//
	// example:
	//
	// 4
	ConnectCU *int32 `json:"ConnectCU,omitempty" xml:"ConnectCU,omitempty"`
	// The number of replicas of Connect.
	//
	// example:
	//
	// 2
	ConnectReplica *int32 `json:"ConnectReplica,omitempty" xml:"ConnectReplica,omitempty"`
	// The number of CPU cores of Control Center.
	//
	// example:
	//
	// 4
	ControlCenterCU *int32 `json:"ControlCenterCU,omitempty" xml:"ControlCenterCU,omitempty"`
	// The number of replicas of Control Center.
	//
	// example:
	//
	// 1
	ControlCenterReplica *int32 `json:"ControlCenterReplica,omitempty" xml:"ControlCenterReplica,omitempty"`
	// The disk capacity of Control Center. Unit: GB.
	//
	// example:
	//
	// 300
	ControlCenterStorage *int32 `json:"ControlCenterStorage,omitempty" xml:"ControlCenterStorage,omitempty"`
	// The number of CPU cores of the Kafka broker.
	//
	// example:
	//
	// 4
	KafkaCU *int32 `json:"KafkaCU,omitempty" xml:"KafkaCU,omitempty"`
	// The number of replicas of the Kafka broker.
	//
	// example:
	//
	// 3
	KafkaReplica *int32 `json:"KafkaReplica,omitempty" xml:"KafkaReplica,omitempty"`
	// The number of CPU cores of Kafka Rest Proxy.
	//
	// example:
	//
	// 4
	KafkaRestProxyCU *int32 `json:"KafkaRestProxyCU,omitempty" xml:"KafkaRestProxyCU,omitempty"`
	// The number of replicas of Kafka Rest Proxy.
	//
	// example:
	//
	// 2
	KafkaRestProxyReplica *int32 `json:"KafkaRestProxyReplica,omitempty" xml:"KafkaRestProxyReplica,omitempty"`
	// The disk capacity of the Kafka broker. Unit: GB.
	//
	// example:
	//
	// 800
	KafkaStorage *int32 `json:"KafkaStorage,omitempty" xml:"KafkaStorage,omitempty"`
	// The number of CPU cores of ksqlDB.
	//
	// example:
	//
	// 2
	KsqlCU *int32 `json:"KsqlCU,omitempty" xml:"KsqlCU,omitempty"`
	// The number of replicas of ksqlDB.
	//
	// example:
	//
	// 2
	KsqlReplica *int32 `json:"KsqlReplica,omitempty" xml:"KsqlReplica,omitempty"`
	// The disk capacity of ksqlDB. Unit: GB.
	//
	// example:
	//
	// 100
	KsqlStorage *int32 `json:"KsqlStorage,omitempty" xml:"KsqlStorage,omitempty"`
	// The number of CPU cores of Schema Registry.
	//
	// example:
	//
	// 4
	SchemaRegistryCU *int32 `json:"SchemaRegistryCU,omitempty" xml:"SchemaRegistryCU,omitempty"`
	// The number of replicas of Schema Registry.
	//
	// example:
	//
	// 2
	SchemaRegistryReplica *int32 `json:"SchemaRegistryReplica,omitempty" xml:"SchemaRegistryReplica,omitempty"`
	// The number of CPU cores of ZooKeeper.
	//
	// example:
	//
	// 2
	ZooKeeperCU *int32 `json:"ZooKeeperCU,omitempty" xml:"ZooKeeperCU,omitempty"`
	// The number of replicas of ZooKeeper.
	//
	// example:
	//
	// 3
	ZooKeeperReplica *int32 `json:"ZooKeeperReplica,omitempty" xml:"ZooKeeperReplica,omitempty"`
	// The disk capacity of ZooKeeper. Unit: GB.
	//
	// example:
	//
	// 100
	ZooKeeperStorage *int32 `json:"ZooKeeperStorage,omitempty" xml:"ZooKeeperStorage,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetConnectCU() *int32 {
	return s.ConnectCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetConnectReplica() *int32 {
	return s.ConnectReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetControlCenterCU() *int32 {
	return s.ControlCenterCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetControlCenterReplica() *int32 {
	return s.ControlCenterReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetControlCenterStorage() *int32 {
	return s.ControlCenterStorage
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKafkaCU() *int32 {
	return s.KafkaCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKafkaReplica() *int32 {
	return s.KafkaReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKafkaRestProxyCU() *int32 {
	return s.KafkaRestProxyCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKafkaRestProxyReplica() *int32 {
	return s.KafkaRestProxyReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKafkaStorage() *int32 {
	return s.KafkaStorage
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKsqlCU() *int32 {
	return s.KsqlCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKsqlReplica() *int32 {
	return s.KsqlReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetKsqlStorage() *int32 {
	return s.KsqlStorage
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetSchemaRegistryCU() *int32 {
	return s.SchemaRegistryCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetSchemaRegistryReplica() *int32 {
	return s.SchemaRegistryReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetZooKeeperCU() *int32 {
	return s.ZooKeeperCU
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetZooKeeperReplica() *int32 {
	return s.ZooKeeperReplica
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GetZooKeeperStorage() *int32 {
	return s.ZooKeeperStorage
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetConnectCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ConnectCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetConnectReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ConnectReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetControlCenterCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ControlCenterCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetControlCenterReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ControlCenterReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetControlCenterStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ControlCenterStorage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaRestProxyCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaRestProxyCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaRestProxyReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaRestProxyReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaStorage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKsqlCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKsqlReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KsqlReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKsqlStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KsqlStorage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetSchemaRegistryCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.SchemaRegistryCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetSchemaRegistryReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.SchemaRegistryReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetZooKeeperCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ZooKeeperCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetZooKeeperReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ZooKeeperReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetZooKeeperStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ZooKeeperStorage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents struct {
	ConfluentInstanceComponentVO []*GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO `json:"ConfluentInstanceComponentVO,omitempty" xml:"ConfluentInstanceComponentVO,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) GetConfluentInstanceComponentVO() []*GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO {
	return s.ConfluentInstanceComponentVO
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) SetConfluentInstanceComponentVO(v []*GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents {
	s.ConfluentInstanceComponentVO = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponents) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO struct {
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	DeployModule  *string `json:"DeployModule,omitempty" xml:"DeployModule,omitempty"`
	PubEndpoint   *string `json:"PubEndpoint,omitempty" xml:"PubEndpoint,omitempty"`
	VpcEndpoint   *string `json:"VpcEndpoint,omitempty" xml:"VpcEndpoint,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) GetComponentType() *string {
	return s.ComponentType
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) GetDeployModule() *string {
	return s.DeployModule
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) GetPubEndpoint() *string {
	return s.PubEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) GetVpcEndpoint() *string {
	return s.VpcEndpoint
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) SetComponentType(v string) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO {
	s.ComponentType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) SetDeployModule(v string) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO {
	s.DeployModule = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) SetPubEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO {
	s.PubEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) SetVpcEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO {
	s.VpcEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentInstanceComponentsConfluentInstanceComponentVO) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVOTags struct {
	TagVO []*GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO `json:"TagVO,omitempty" xml:"TagVO,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTags) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTags) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTags) GetTagVO() []*GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO {
	return s.TagVO
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTags) SetTagVO(v []*GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) *GetInstanceListResponseBodyInstanceListInstanceVOTags {
	s.TagVO = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTags) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) GetKey() *string {
	return s.Key
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) GetValue() *string {
	return s.Value
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) SetKey(v string) *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO {
	s.Key = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) SetValue(v string) *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO {
	s.Value = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo struct {
	// The open source Apache Kafka version that corresponds to the instance.
	//
	// example:
	//
	// 2.2.0
	Current2OpenSourceVersion *string `json:"Current2OpenSourceVersion,omitempty" xml:"Current2OpenSourceVersion,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) GetCurrent2OpenSourceVersion() *string {
	return s.Current2OpenSourceVersion
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) SetCurrent2OpenSourceVersion(v string) *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo {
	s.Current2OpenSourceVersion = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds struct {
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) SetVSwitchIds(v []*string) *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds {
	s.VSwitchIds = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudBenchTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v string) *CreateCloudBenchTasksRequest
	GetAmount() *string
	SetBackupId(v string) *CreateCloudBenchTasksRequest
	GetBackupId() *string
	SetBackupTime(v string) *CreateCloudBenchTasksRequest
	GetBackupTime() *string
	SetClientType(v string) *CreateCloudBenchTasksRequest
	GetClientType() *string
	SetDescription(v string) *CreateCloudBenchTasksRequest
	GetDescription() *string
	SetDstConnectionString(v string) *CreateCloudBenchTasksRequest
	GetDstConnectionString() *string
	SetDstInstanceId(v string) *CreateCloudBenchTasksRequest
	GetDstInstanceId() *string
	SetDstPort(v string) *CreateCloudBenchTasksRequest
	GetDstPort() *string
	SetDstSuperAccount(v string) *CreateCloudBenchTasksRequest
	GetDstSuperAccount() *string
	SetDstSuperPassword(v string) *CreateCloudBenchTasksRequest
	GetDstSuperPassword() *string
	SetDstType(v string) *CreateCloudBenchTasksRequest
	GetDstType() *string
	SetDtsJobClass(v string) *CreateCloudBenchTasksRequest
	GetDtsJobClass() *string
	SetDtsJobId(v string) *CreateCloudBenchTasksRequest
	GetDtsJobId() *string
	SetEndState(v string) *CreateCloudBenchTasksRequest
	GetEndState() *string
	SetGatewayVpcId(v string) *CreateCloudBenchTasksRequest
	GetGatewayVpcId() *string
	SetGatewayVpcIp(v string) *CreateCloudBenchTasksRequest
	GetGatewayVpcIp() *string
	SetRate(v string) *CreateCloudBenchTasksRequest
	GetRate() *string
	SetRequestDuration(v string) *CreateCloudBenchTasksRequest
	GetRequestDuration() *string
	SetRequestEndTime(v string) *CreateCloudBenchTasksRequest
	GetRequestEndTime() *string
	SetRequestStartTime(v string) *CreateCloudBenchTasksRequest
	GetRequestStartTime() *string
	SetSmartPressureTime(v string) *CreateCloudBenchTasksRequest
	GetSmartPressureTime() *string
	SetSrcInstanceId(v string) *CreateCloudBenchTasksRequest
	GetSrcInstanceId() *string
	SetSrcPublicIp(v string) *CreateCloudBenchTasksRequest
	GetSrcPublicIp() *string
	SetSrcSuperAccount(v string) *CreateCloudBenchTasksRequest
	GetSrcSuperAccount() *string
	SetSrcSuperPassword(v string) *CreateCloudBenchTasksRequest
	GetSrcSuperPassword() *string
	SetTaskType(v string) *CreateCloudBenchTasksRequest
	GetTaskType() *string
	SetWorkDir(v string) *CreateCloudBenchTasksRequest
	GetWorkDir() *string
}

type CreateCloudBenchTasksRequest struct {
	// The total number of stress testing tasks that you want to create. Valid values: **0*	- to **30**. Default value: **1**.
	//
	// example:
	//
	// 1
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The ID of the backup set. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/26273.html) operation to query the ID of the backup set.
	//
	// example:
	//
	// 229132
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The time when the backup starts. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-04-23T13:22:14Z
	BackupTime *string `json:"BackupTime,omitempty" xml:"BackupTime,omitempty"`
	// The type of the stress testing client. Valid values:
	//
	// 	- **ECS**: indicates that you must create the [DBGateway](https://help.aliyun.com/document_detail/64905.html).
	//
	// 	- **DAS_ECS**: indicates that DAS automatically purchases and deploys an Elastic Compute Service (ECS) instance for stress testing.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The description of the stress testing task.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// test-das-bench-0501
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The endpoint of the destination instance. The specified endpoint must be the endpoint of an ApsaraDB RDS for MySQL instance or a PolarDB for MySQL instance.
	//
	// >  This parameter takes effect only if you set **DstType*	- to **ConnectionString**.
	//
	// example:
	//
	// rm-de21209****.mysql.rds.aliyuncs.com
	DstConnectionString *string `json:"DstConnectionString,omitempty" xml:"DstConnectionString,omitempty"`
	// The ID of the destination instance. The instance must be an ApsaraDB RDS for MySQL instance or a PolarDB for MySQL instance. You can call the [GetInstanceInspections](https://help.aliyun.com/document_detail/202857.html) operation to query the ID.
	//
	// >  This parameter must be specified if you set **DstType*	- to **Instance**.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	DstInstanceId *string `json:"DstInstanceId,omitempty" xml:"DstInstanceId,omitempty"`
	// The port number of the instance that you want to access.
	//
	// >  This parameter takes effect only if you set **DstType*	- to **ConnectionString**.
	//
	// example:
	//
	// 3306
	DstPort *string `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The name of the privileged account for the destination instance.
	//
	// example:
	//
	// root
	DstSuperAccount *string `json:"DstSuperAccount,omitempty" xml:"DstSuperAccount,omitempty"`
	// The password of the privileged account for the destination instance.
	//
	// example:
	//
	// test123
	DstSuperPassword *string `json:"DstSuperPassword,omitempty" xml:"DstSuperPassword,omitempty"`
	// The type of the identifier that is used to indicate the destination instance. Valid values:
	//
	// 	- **Instance**: the instance ID. This is the default value.
	//
	// 	- **ConnectionString**: the endpoint of the instance.
	//
	// example:
	//
	// Instance
	DstType *string `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The specification of the Data Transmission Service (DTS) migration task. You can call the [DescribeCloudbenchTask](https://help.aliyun.com/document_detail/230669.html) operation to query the specification.
	//
	// >  You must migrate the basic data in the source instance to the destination instance before you start a stress testing task. When you create a DTS migration task, you must specify this parameter.
	//
	// example:
	//
	// medium
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The ID of the DTS migration task. You can call the [ConfigureDtsJob](https://help.aliyun.com/document_detail/208399.html) operation to query the ID.
	//
	// >  After a DTS migration task is created in the DTS console, you must specify this parameter.
	//
	// example:
	//
	// 23127
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The state that specifies the last operation that is performed for the stress testing task. Valid values:
	//
	// 	- **WAIT_TARGET**: prepares the destination instance
	//
	// 	- **WAIT_DBGATEWAY**: prepares the DBGateway
	//
	// 	- **WAIT_SQL**: prepares the full SQL statistics
	//
	// 	- **WAIT_LOGIC**: prepares to replay the traffic
	//
	// >  When the state of a stress testing task changes to the state that is specified by the EndState parameter, the stress testing task becomes completed.
	//
	// example:
	//
	// WAIT_TARGET
	EndState *string `json:"EndState,omitempty" xml:"EndState,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the database gateway (DBGateway) is deployed.
	//
	// >  This parameter must be specified if you set **ClientType*	- to **ECS**.
	//
	// example:
	//
	// vpc-t4nsnwvpbc1h76ja4****
	GatewayVpcId *string `json:"GatewayVpcId,omitempty" xml:"GatewayVpcId,omitempty"`
	// The IP address or domain name of the DBGateway.
	//
	// >  This parameter must be specified if you set **ClientType*	- to **ECS**.
	//
	// example:
	//
	// 172.30.XX.XX
	GatewayVpcIp *string `json:"GatewayVpcIp,omitempty" xml:"GatewayVpcIp,omitempty"`
	// The rate at which the traffic captured from the source instance is replayed on the destination instance. The value must be a positive integer. Valid values: **1*	- to **30**. Default value: **1**.
	//
	// example:
	//
	// 1
	Rate *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The duration of the stress testing task for which the traffic is captured from the source instance. Unit: milliseconds.
	//
	// example:
	//
	// 86400000
	RequestDuration *string `json:"RequestDuration,omitempty" xml:"RequestDuration,omitempty"`
	// The time when the stress testing task ends. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1608888296001
	RequestEndTime *string `json:"RequestEndTime,omitempty" xml:"RequestEndTime,omitempty"`
	// The time when the stress testing task starts. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1608888296000
	RequestStartTime *string `json:"RequestStartTime,omitempty" xml:"RequestStartTime,omitempty"`
	// The duration within which the traffic generation stressing test takes effect. Unit: milliseconds.
	//
	// >  This parameter must be specified if you set **TaskType*	- to **smart pressure test**.
	//
	// example:
	//
	// 86400000
	SmartPressureTime *string `json:"SmartPressureTime,omitempty" xml:"SmartPressureTime,omitempty"`
	// The ID of the source instance. The instance must be an ApsaraDB RDS for MySQL instance or a PolarDB for MySQL instance. You can call the [GetInstanceInspections](https://help.aliyun.com/document_detail/202857.html) operation to query the ID.
	//
	// >  This parameter must be specified if you set **DstType*	- to **Instance**.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" xml:"SrcInstanceId,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	SrcPublicIp *string `json:"SrcPublicIp,omitempty" xml:"SrcPublicIp,omitempty"`
	// The name of the privileged account for the source instance. Set the value to **admin**.
	//
	// >  This parameter must be specified if you set **DstType*	- to **Instance**.
	//
	// example:
	//
	// admin
	SrcSuperAccount *string `json:"SrcSuperAccount,omitempty" xml:"SrcSuperAccount,omitempty"`
	// The password of the privileged account for the source instance.
	//
	// >  This parameter must be specified if you set **DstType*	- to **Instance**.
	//
	// example:
	//
	// test123
	SrcSuperPassword *string `json:"SrcSuperPassword,omitempty" xml:"SrcSuperPassword,omitempty"`
	// The type of the stress testing task. Valid values:
	//
	// 	- **pressure test*	- (default): A task of this type replays the traffic that is captured from the source instance on the destination instance at the maximum playback rate that is supported by the destination instance.
	//
	// 	- **smart pressure test**: A task of this type analyzes the traffic that is captured from the source instance over a short period of time and generates traffic on the destination instance for continuous stress testing. The business model based on which the traffic is generated on the destination instance and the traffic distribution are consistent with those on the source instance. Stress testing tasks of this type can help you reduce the amount of time that is consumed to collect data from the source instance and reduce storage costs and performance overheads.
	//
	// This parameter is required.
	//
	// example:
	//
	// pressure test
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The temporary directory generated for stress testing.
	//
	// example:
	//
	// /tmp/bench/
	WorkDir *string `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
}

func (s CreateCloudBenchTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudBenchTasksRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudBenchTasksRequest) GetAmount() *string {
	return s.Amount
}

func (s *CreateCloudBenchTasksRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateCloudBenchTasksRequest) GetBackupTime() *string {
	return s.BackupTime
}

func (s *CreateCloudBenchTasksRequest) GetClientType() *string {
	return s.ClientType
}

func (s *CreateCloudBenchTasksRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCloudBenchTasksRequest) GetDstConnectionString() *string {
	return s.DstConnectionString
}

func (s *CreateCloudBenchTasksRequest) GetDstInstanceId() *string {
	return s.DstInstanceId
}

func (s *CreateCloudBenchTasksRequest) GetDstPort() *string {
	return s.DstPort
}

func (s *CreateCloudBenchTasksRequest) GetDstSuperAccount() *string {
	return s.DstSuperAccount
}

func (s *CreateCloudBenchTasksRequest) GetDstSuperPassword() *string {
	return s.DstSuperPassword
}

func (s *CreateCloudBenchTasksRequest) GetDstType() *string {
	return s.DstType
}

func (s *CreateCloudBenchTasksRequest) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *CreateCloudBenchTasksRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *CreateCloudBenchTasksRequest) GetEndState() *string {
	return s.EndState
}

func (s *CreateCloudBenchTasksRequest) GetGatewayVpcId() *string {
	return s.GatewayVpcId
}

func (s *CreateCloudBenchTasksRequest) GetGatewayVpcIp() *string {
	return s.GatewayVpcIp
}

func (s *CreateCloudBenchTasksRequest) GetRate() *string {
	return s.Rate
}

func (s *CreateCloudBenchTasksRequest) GetRequestDuration() *string {
	return s.RequestDuration
}

func (s *CreateCloudBenchTasksRequest) GetRequestEndTime() *string {
	return s.RequestEndTime
}

func (s *CreateCloudBenchTasksRequest) GetRequestStartTime() *string {
	return s.RequestStartTime
}

func (s *CreateCloudBenchTasksRequest) GetSmartPressureTime() *string {
	return s.SmartPressureTime
}

func (s *CreateCloudBenchTasksRequest) GetSrcInstanceId() *string {
	return s.SrcInstanceId
}

func (s *CreateCloudBenchTasksRequest) GetSrcPublicIp() *string {
	return s.SrcPublicIp
}

func (s *CreateCloudBenchTasksRequest) GetSrcSuperAccount() *string {
	return s.SrcSuperAccount
}

func (s *CreateCloudBenchTasksRequest) GetSrcSuperPassword() *string {
	return s.SrcSuperPassword
}

func (s *CreateCloudBenchTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateCloudBenchTasksRequest) GetWorkDir() *string {
	return s.WorkDir
}

func (s *CreateCloudBenchTasksRequest) SetAmount(v string) *CreateCloudBenchTasksRequest {
	s.Amount = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetBackupId(v string) *CreateCloudBenchTasksRequest {
	s.BackupId = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetBackupTime(v string) *CreateCloudBenchTasksRequest {
	s.BackupTime = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetClientType(v string) *CreateCloudBenchTasksRequest {
	s.ClientType = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDescription(v string) *CreateCloudBenchTasksRequest {
	s.Description = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstConnectionString(v string) *CreateCloudBenchTasksRequest {
	s.DstConnectionString = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstInstanceId(v string) *CreateCloudBenchTasksRequest {
	s.DstInstanceId = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstPort(v string) *CreateCloudBenchTasksRequest {
	s.DstPort = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstSuperAccount(v string) *CreateCloudBenchTasksRequest {
	s.DstSuperAccount = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstSuperPassword(v string) *CreateCloudBenchTasksRequest {
	s.DstSuperPassword = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDstType(v string) *CreateCloudBenchTasksRequest {
	s.DstType = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDtsJobClass(v string) *CreateCloudBenchTasksRequest {
	s.DtsJobClass = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetDtsJobId(v string) *CreateCloudBenchTasksRequest {
	s.DtsJobId = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetEndState(v string) *CreateCloudBenchTasksRequest {
	s.EndState = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetGatewayVpcId(v string) *CreateCloudBenchTasksRequest {
	s.GatewayVpcId = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetGatewayVpcIp(v string) *CreateCloudBenchTasksRequest {
	s.GatewayVpcIp = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetRate(v string) *CreateCloudBenchTasksRequest {
	s.Rate = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetRequestDuration(v string) *CreateCloudBenchTasksRequest {
	s.RequestDuration = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetRequestEndTime(v string) *CreateCloudBenchTasksRequest {
	s.RequestEndTime = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetRequestStartTime(v string) *CreateCloudBenchTasksRequest {
	s.RequestStartTime = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetSmartPressureTime(v string) *CreateCloudBenchTasksRequest {
	s.SmartPressureTime = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetSrcInstanceId(v string) *CreateCloudBenchTasksRequest {
	s.SrcInstanceId = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetSrcPublicIp(v string) *CreateCloudBenchTasksRequest {
	s.SrcPublicIp = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetSrcSuperAccount(v string) *CreateCloudBenchTasksRequest {
	s.SrcSuperAccount = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetSrcSuperPassword(v string) *CreateCloudBenchTasksRequest {
	s.SrcSuperPassword = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetTaskType(v string) *CreateCloudBenchTasksRequest {
	s.TaskType = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) SetWorkDir(v string) *CreateCloudBenchTasksRequest {
	s.WorkDir = &v
	return s
}

func (s *CreateCloudBenchTasksRequest) Validate() error {
	return dara.Validate(s)
}

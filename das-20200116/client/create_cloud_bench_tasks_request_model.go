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
	// The total number of stress testing tasks to create. Valid values: **0*	- to **30**. Default value: **1**.
	//
	// example:
	//
	// 1
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The ID of the backup set. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/26273.html) operation to query the backup list and obtain the ID.
	//
	// example:
	//
	// 229132
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The time of the backup. Format: yyyy-MM-ddTHH:mm:ssZ (UTC time).
	//
	// example:
	//
	// 2021-04-23T13:22:14Z
	BackupTime *string `json:"BackupTime,omitempty" xml:"BackupTime,omitempty"`
	// The type of stress testing machine. Valid values:
	//
	// - **ECS**: You need to prepare a [Database Gateway](https://help.aliyun.com/document_detail/64905.html) yourself.
	//
	// - **DAS_ECS**: An ECS instance that is automatically purchased and deployed by DAS.
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
	// The connection address of the target instance. Only RDS MySQL and PolarDB MySQL instances are supported.
	//
	// > This parameter takes effect when **DstType*	- is set to **ConnectionString**.
	//
	// example:
	//
	// rm-de21209****.mysql.rds.aliyuncs.com
	DstConnectionString *string `json:"DstConnectionString,omitempty" xml:"DstConnectionString,omitempty"`
	// The ID of the target instance. Only RDS MySQL and PolarDB MySQL instances are supported. You can call the [GetInstanceInspections](https://help.aliyun.com/document_detail/202857.html) operation to obtain the ID.
	//
	// > This parameter is required when **DstType*	- is set to **Instance**.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	DstInstanceId *string `json:"DstInstanceId,omitempty" xml:"DstInstanceId,omitempty"`
	// The port of the target instance.
	//
	// > This parameter takes effect when **DstType*	- is set to **ConnectionString**.
	//
	// example:
	//
	// 3306
	DstPort *string `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The privileged account of the target instance.
	//
	// example:
	//
	// root
	DstSuperAccount *string `json:"DstSuperAccount,omitempty" xml:"DstSuperAccount,omitempty"`
	// The password of the privileged account of the target instance.
	//
	// example:
	//
	// test123
	DstSuperPassword *string `json:"DstSuperPassword,omitempty" xml:"DstSuperPassword,omitempty"`
	// The type of the target instance. Valid values:
	//
	// - **Instance*	- (default): instance ID.
	//
	// - **ConnectionString**: connection address of the instance.
	//
	// example:
	//
	// Instance
	DstType *string `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The specification of the DTS migration task. You can call the [DescribeCloudbenchTask](https://help.aliyun.com/document_detail/230669.html) operation to obtain the specification.
	//
	// > The stress testing task needs to migrate the baseline data from the source instance to the target instance. This parameter is required when you create a new DTS task.
	//
	// example:
	//
	// medium
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The ID of the DTS migration task. You can call the [ConfigureDtsJob](https://help.aliyun.com/document_detail/208399.html) operation to obtain the ID.
	//
	// > This parameter is required when a DTS task has been created in the DTS console.
	//
	// example:
	//
	// 23127
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The status after the stress testing task ends. Valid values:
	//
	// - **WAIT_TARGET**: Prepare the target instance for stress testing.
	//
	// - **WAIT_DBGATEWAY**: Prepare the stress testing deployment.
	//
	// - **WAIT_SQL**: Prepare the full SQL statements.
	//
	// - **WAIT_LOGIC**: Prepare to start replaying the traffic.
	//
	// > When the stress testing task completes the status set by EndState, the task directly reaches the completed status.
	//
	// example:
	//
	// WAIT_TARGET
	EndState *string `json:"EndState,omitempty" xml:"EndState,omitempty"`
	// The virtual private cloud (VPC) ID of the Database Gateway.
	//
	// > This parameter is required when **ClientType*	- is set to **ECS**.
	//
	// example:
	//
	// vpc-t4nsnwvpbc1h76ja4****
	GatewayVpcId *string `json:"GatewayVpcId,omitempty" xml:"GatewayVpcId,omitempty"`
	// The IP address or domain name of the Database Gateway.
	//
	// > This parameter is required when **ClientType*	- is set to **ECS**.
	//
	// example:
	//
	// 172.30.XX.XX
	GatewayVpcIp *string `json:"GatewayVpcIp,omitempty" xml:"GatewayVpcIp,omitempty"`
	// The replay speed of the source instance traffic on the target instance. The replay speed must be a positive integer. Valid values: **1*	- to **30**. Default value: **1**.
	//
	// example:
	//
	// 1
	Rate *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The duration of the stress testing task. Unit: milliseconds.
	//
	// example:
	//
	// 86400000
	RequestDuration *string `json:"RequestDuration,omitempty" xml:"RequestDuration,omitempty"`
	// The end time of the stress testing task. The time is in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1608888296001
	RequestEndTime *string `json:"RequestEndTime,omitempty" xml:"RequestEndTime,omitempty"`
	// The start time of the stress testing task. The time is in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1608888296000
	RequestStartTime *string `json:"RequestStartTime,omitempty" xml:"RequestStartTime,omitempty"`
	// The duration of the generated stress testing. Unit: milliseconds.
	//
	// > This parameter is required when **TaskType*	- is set to **smart pressure test**.
	//
	// example:
	//
	// 86400000
	SmartPressureTime *string `json:"SmartPressureTime,omitempty" xml:"SmartPressureTime,omitempty"`
	// The ID of the source instance. Only RDS MySQL and PolarDB MySQL instances are supported. You can call the [GetInstanceInspections](https://help.aliyun.com/document_detail/202857.html) operation to obtain the ID.
	//
	// > This parameter is required when **DstType*	- is set to **Instance**.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" xml:"SrcInstanceId,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// None
	SrcPublicIp *string `json:"SrcPublicIp,omitempty" xml:"SrcPublicIp,omitempty"`
	// The privileged account of the source instance. Value: **admin**.
	//
	// > This parameter is required when **DstType*	- is set to **Instance**.
	//
	// example:
	//
	// admin
	SrcSuperAccount *string `json:"SrcSuperAccount,omitempty" xml:"SrcSuperAccount,omitempty"`
	// The password of the privileged account of the source instance.
	//
	// > This parameter is required when **DstType*	- is set to **Instance**.
	//
	// example:
	//
	// test123
	SrcSuperPassword *string `json:"SrcSuperPassword,omitempty" xml:"SrcSuperPassword,omitempty"`
	// The type of stress testing task. Valid values:
	//
	// - **pressure test*	- (default): Intelligent stress testing, which replays the traffic captured from the source instance on the target instance at the maximum speed supported by the target instance type.
	//
	// - **smart pressure test**: Generated stress testing, which analyzes and learns from the traffic captured from the source instance in a short period of time, generates traffic that is consistent with the business model and traffic distribution of the original traffic for continuous stress testing, reduces the time for collecting data from the source instance, and reduces storage costs and performance overhead.
	//
	// This parameter is required.
	//
	// example:
	//
	// pressure test
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The temporary directory generated by the stress testing.
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

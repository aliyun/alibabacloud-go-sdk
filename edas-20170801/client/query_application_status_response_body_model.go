// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApplicationStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInfo(v *QueryApplicationStatusResponseBodyAppInfo) *QueryApplicationStatusResponseBody
	GetAppInfo() *QueryApplicationStatusResponseBodyAppInfo
	SetCode(v int32) *QueryApplicationStatusResponseBody
	GetCode() *int32
	SetMessage(v string) *QueryApplicationStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryApplicationStatusResponseBody
	GetRequestId() *string
}

type QueryApplicationStatusResponseBody struct {
	// The information about the application.
	AppInfo *QueryApplicationStatusResponseBodyAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryApplicationStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBody) GetAppInfo() *QueryApplicationStatusResponseBodyAppInfo {
	return s.AppInfo
}

func (s *QueryApplicationStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryApplicationStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryApplicationStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryApplicationStatusResponseBody) SetAppInfo(v *QueryApplicationStatusResponseBodyAppInfo) *QueryApplicationStatusResponseBody {
	s.AppInfo = v
	return s
}

func (s *QueryApplicationStatusResponseBody) SetCode(v int32) *QueryApplicationStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryApplicationStatusResponseBody) SetMessage(v string) *QueryApplicationStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryApplicationStatusResponseBody) SetRequestId(v string) *QueryApplicationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryApplicationStatusResponseBody) Validate() error {
	if s.AppInfo != nil {
		if err := s.AppInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryApplicationStatusResponseBodyAppInfo struct {
	// The basic information about the application.
	Application *QueryApplicationStatusResponseBodyAppInfoApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The information about deployment records.
	DeployRecordList *QueryApplicationStatusResponseBodyAppInfoDeployRecordList `json:"DeployRecordList,omitempty" xml:"DeployRecordList,omitempty" type:"Struct"`
	// The information about elastic compute containers (ECCs).
	EccList *QueryApplicationStatusResponseBodyAppInfoEccList `json:"EccList,omitempty" xml:"EccList,omitempty" type:"Struct"`
	// The information about elastic compute units (ECUs).
	EcuList *QueryApplicationStatusResponseBodyAppInfoEcuList `json:"EcuList,omitempty" xml:"EcuList,omitempty" type:"Struct"`
	// The information about the instance groups.
	GroupList *QueryApplicationStatusResponseBodyAppInfoGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Struct"`
}

func (s QueryApplicationStatusResponseBodyAppInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfo) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfo) GetApplication() *QueryApplicationStatusResponseBodyAppInfoApplication {
	return s.Application
}

func (s *QueryApplicationStatusResponseBodyAppInfo) GetDeployRecordList() *QueryApplicationStatusResponseBodyAppInfoDeployRecordList {
	return s.DeployRecordList
}

func (s *QueryApplicationStatusResponseBodyAppInfo) GetEccList() *QueryApplicationStatusResponseBodyAppInfoEccList {
	return s.EccList
}

func (s *QueryApplicationStatusResponseBodyAppInfo) GetEcuList() *QueryApplicationStatusResponseBodyAppInfoEcuList {
	return s.EcuList
}

func (s *QueryApplicationStatusResponseBodyAppInfo) GetGroupList() *QueryApplicationStatusResponseBodyAppInfoGroupList {
	return s.GroupList
}

func (s *QueryApplicationStatusResponseBodyAppInfo) SetApplication(v *QueryApplicationStatusResponseBodyAppInfoApplication) *QueryApplicationStatusResponseBodyAppInfo {
	s.Application = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfo) SetDeployRecordList(v *QueryApplicationStatusResponseBodyAppInfoDeployRecordList) *QueryApplicationStatusResponseBodyAppInfo {
	s.DeployRecordList = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfo) SetEccList(v *QueryApplicationStatusResponseBodyAppInfoEccList) *QueryApplicationStatusResponseBodyAppInfo {
	s.EccList = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfo) SetEcuList(v *QueryApplicationStatusResponseBodyAppInfoEcuList) *QueryApplicationStatusResponseBodyAppInfo {
	s.EcuList = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfo) SetGroupList(v *QueryApplicationStatusResponseBodyAppInfoGroupList) *QueryApplicationStatusResponseBodyAppInfo {
	s.GroupList = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfo) Validate() error {
	if s.Application != nil {
		if err := s.Application.Validate(); err != nil {
			return err
		}
	}
	if s.DeployRecordList != nil {
		if err := s.DeployRecordList.Validate(); err != nil {
			return err
		}
	}
	if s.EccList != nil {
		if err := s.EccList.Validate(); err != nil {
			return err
		}
	}
	if s.EcuList != nil {
		if err := s.EcuList.Validate(); err != nil {
			return err
		}
	}
	if s.GroupList != nil {
		if err := s.GroupList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryApplicationStatusResponseBodyAppInfoApplication struct {
	// The ID of the application.
	//
	// example:
	//
	// 3616cdca-4f92-4413-****-************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The build package number of Enterprise Distributed Application Service (EDAS) Container.
	//
	// example:
	//
	// 0
	BuildPackageId *int32 `json:"BuildPackageId,omitempty" xml:"BuildPackageId,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// 0d247b93-8d62-4e34-****-************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of CPU cores used by the application.
	//
	// example:
	//
	// 0
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the application was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573626207270
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the application is a Docker application.
	//
	// example:
	//
	// false
	Dockerize *bool `json:"Dockerize,omitempty" xml:"Dockerize,omitempty"`
	// The email address of the user who created the application.
	//
	// example:
	//
	// 1234567@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The health check URL.
	//
	// example:
	//
	// “”
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The number of application instances.
	//
	// example:
	//
	// 1
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The time when the application was launched. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 0
	LaunchTime *int64 `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 0
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// EDAS-scaled-cluster:default cluster
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the user who created the application.
	//
	// example:
	//
	// edas_com***_****@******-*****.***
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The mobile number of the user who created the application.
	//
	// example:
	//
	// 1886666****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The port used by the application.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-shenzhen:test
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of application instances that are running.
	//
	// example:
	//
	// 1
	RunningInstanceCount *int32 `json:"RunningInstanceCount,omitempty" xml:"RunningInstanceCount,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// edas_com***_****@******-*****.***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryApplicationStatusResponseBodyAppInfoApplication) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoApplication) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetBuildPackageId() *int32 {
	return s.BuildPackageId
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetDockerize() *bool {
	return s.Dockerize
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetEmail() *string {
	return s.Email
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetHealthCheckUrl() *string {
	return s.HealthCheckUrl
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetLaunchTime() *int64 {
	return s.LaunchTime
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetMemory() *int32 {
	return s.Memory
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetName() *string {
	return s.Name
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetOwner() *string {
	return s.Owner
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetPhone() *string {
	return s.Phone
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetPort() *int32 {
	return s.Port
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetRunningInstanceCount() *int32 {
	return s.RunningInstanceCount
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) GetUserId() *string {
	return s.UserId
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetApplicationId(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.ApplicationId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetBuildPackageId(v int32) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.BuildPackageId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetClusterId(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.ClusterId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetCpu(v int32) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Cpu = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetCreateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.CreateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetDockerize(v bool) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Dockerize = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetEmail(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Email = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetHealthCheckUrl(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.HealthCheckUrl = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetInstanceCount(v int32) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.InstanceCount = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetLaunchTime(v int64) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.LaunchTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetMemory(v int32) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Memory = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetName(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Name = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetOwner(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Owner = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetPhone(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Phone = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetPort(v int32) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Port = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetRegionId(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.RegionId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetRunningInstanceCount(v int32) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.RunningInstanceCount = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetUserId(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.UserId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) Validate() error {
	return dara.Validate(s)
}

type QueryApplicationStatusResponseBodyAppInfoDeployRecordList struct {
	DeployRecord []*QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord `json:"DeployRecord,omitempty" xml:"DeployRecord,omitempty" type:"Repeated"`
}

func (s QueryApplicationStatusResponseBodyAppInfoDeployRecordList) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoDeployRecordList) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordList) GetDeployRecord() []*QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	return s.DeployRecord
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordList) SetDeployRecord(v []*QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) *QueryApplicationStatusResponseBodyAppInfoDeployRecordList {
	s.DeployRecord = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordList) Validate() error {
	if s.DeployRecord != nil {
		for _, item := range s.DeployRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord struct {
	// The time when the deployment record was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573626226691
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the deployment record.
	//
	// example:
	//
	// bbc6c0d5-d792-4907-****-************
	DeployRecordId *string `json:"DeployRecordId,omitempty" xml:"DeployRecordId,omitempty"`
	// The unique ID of the ECC.
	//
	// example:
	//
	// 0cf49a6c-95a8-4aa8-****-************
	EccId *string `json:"EccId,omitempty" xml:"EccId,omitempty"`
	// The unique ID of the ECU.
	//
	// example:
	//
	// 07bd417a-b863-477d-****-************
	EcuId *string `json:"EcuId,omitempty" xml:"EcuId,omitempty"`
	// The MD5 hash value of the deployment package.
	//
	// example:
	//
	// d0db5bcb442e492104d0f00e10a03dd9
	PackageMd5 *string `json:"PackageMd5,omitempty" xml:"PackageMd5,omitempty"`
	// The version of the deployment package that was used to deploy an application in the instance group.
	//
	// example:
	//
	// 441beb18-da42-44dc-****-************
	PackageVersionId *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty"`
}

func (s QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) GetDeployRecordId() *string {
	return s.DeployRecordId
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) GetEccId() *string {
	return s.EccId
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) GetEcuId() *string {
	return s.EcuId
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) GetPackageMd5() *string {
	return s.PackageMd5
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) GetPackageVersionId() *string {
	return s.PackageVersionId
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetCreateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.CreateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetDeployRecordId(v string) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.DeployRecordId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetEccId(v string) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.EccId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetEcuId(v string) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.EcuId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetPackageMd5(v string) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.PackageMd5 = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetPackageVersionId(v string) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.PackageVersionId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) Validate() error {
	return dara.Validate(s)
}

type QueryApplicationStatusResponseBodyAppInfoEccList struct {
	Ecc []*QueryApplicationStatusResponseBodyAppInfoEccListEcc `json:"Ecc,omitempty" xml:"Ecc,omitempty" type:"Repeated"`
}

func (s QueryApplicationStatusResponseBodyAppInfoEccList) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoEccList) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccList) GetEcc() []*QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	return s.Ecc
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccList) SetEcc(v []*QueryApplicationStatusResponseBodyAppInfoEccListEcc) *QueryApplicationStatusResponseBodyAppInfoEccList {
	s.Ecc = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccList) Validate() error {
	if s.Ecc != nil {
		for _, item := range s.Ecc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryApplicationStatusResponseBodyAppInfoEccListEcc struct {
	// The ID of the application.
	//
	// example:
	//
	// 3616cdca-4f92-4413-****-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The status of the application instance. Valid values:
	//
	// 	- 0: AGENT_OFF: indicates that the agent is offline.
	//
	// 	- 1: STOPPED: indicates that the application is stopped.
	//
	// 	- 3: RUNNING_BUT_URL_FAILED: indicates that the health check failed.
	//
	// 	- 7: RUNNING: indicates that the application is running.
	//
	// example:
	//
	// 7
	AppState *int32 `json:"AppState,omitempty" xml:"AppState,omitempty"`
	// The status of the container.
	//
	// example:
	//
	// “”
	ContainerStatus *string `json:"ContainerStatus,omitempty" xml:"ContainerStatus,omitempty"`
	// The time when the ECC was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573626226691
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The unique ID of the ECC.
	//
	// example:
	//
	// 0cf49a6c-95a8-4aa8-****-************
	EccId *string `json:"EccId,omitempty" xml:"EccId,omitempty"`
	// The unique ID of the ECU.
	//
	// example:
	//
	// 07bd417a-b863-477d-****-************
	EcuId *string `json:"EcuId,omitempty" xml:"EcuId,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// 8123db90-880f-486f-****-************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The private IP address of the ECU.
	//
	// example:
	//
	// 172.16.*.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The state of the latest task initiated on the application instance. Valid values:
	//
	// 	- 0: UNKNOWN: indicates that the state of the latest task is unknown.
	//
	// 	- 1: PROCESSING: indicates that the latest task is being processed.
	//
	// 	- 2: SUCCESS: indicates that the latest task is executed.
	//
	// 	- 3: FAILED: indicates that the latest task failed.
	//
	// example:
	//
	// 3
	TaskState *int32 `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// The time when the ECC was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573635952012
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-wz9b246zg************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s QueryApplicationStatusResponseBodyAppInfoEccListEcc) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoEccListEcc) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) GetAppId() *string {
	return s.AppId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) GetAppState() *int32 {
	return s.AppState
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) GetContainerStatus() *string {
	return s.ContainerStatus
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) GetEccId() *string {
	return s.EccId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) GetEcuId() *string {
	return s.EcuId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) GetGroupId() *string {
	return s.GroupId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) GetIp() *string {
	return s.Ip
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) GetTaskState() *int32 {
	return s.TaskState
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) GetVpcId() *string {
	return s.VpcId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetAppId(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.AppId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetAppState(v int32) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.AppState = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetContainerStatus(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.ContainerStatus = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetCreateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.CreateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetEccId(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.EccId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetEcuId(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.EcuId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetGroupId(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.GroupId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetIp(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.Ip = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetTaskState(v int32) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.TaskState = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetUpdateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.UpdateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetVpcId(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.VpcId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) Validate() error {
	return dara.Validate(s)
}

type QueryApplicationStatusResponseBodyAppInfoEcuList struct {
	Ecu []*QueryApplicationStatusResponseBodyAppInfoEcuListEcu `json:"Ecu,omitempty" xml:"Ecu,omitempty" type:"Repeated"`
}

func (s QueryApplicationStatusResponseBodyAppInfoEcuList) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoEcuList) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuList) GetEcu() []*QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	return s.Ecu
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuList) SetEcu(v []*QueryApplicationStatusResponseBodyAppInfoEcuListEcu) *QueryApplicationStatusResponseBodyAppInfoEcuList {
	s.Ecu = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuList) Validate() error {
	if s.Ecu != nil {
		for _, item := range s.Ecu {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryApplicationStatusResponseBodyAppInfoEcuListEcu struct {
	// The number of available CPU cores.
	//
	// example:
	//
	// 0
	AvailableCpu *int32 `json:"AvailableCpu,omitempty" xml:"AvailableCpu,omitempty"`
	// The size of the available memory.
	//
	// example:
	//
	// 0
	AvailableMem *int32 `json:"AvailableMem,omitempty" xml:"AvailableMem,omitempty"`
	// The time when the ECU was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573626207270
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether Docker is installed.
	//
	// example:
	//
	// false
	DockerEnv *bool `json:"DockerEnv,omitempty" xml:"DockerEnv,omitempty"`
	// The unique ID of the ECU. You can run the `dmidecode` command on the ECS instance to query the ECU ID.
	//
	// example:
	//
	// 07bd417a-b863-477d-****-************
	EcuId *string `json:"EcuId,omitempty" xml:"EcuId,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// 8123db90-880f-486f-****-************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The time when the last heartbeat detection was performed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573635952012
	HeartbeatTime *int64 `json:"HeartbeatTime,omitempty" xml:"HeartbeatTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-wz9fp1ljg***********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The private IP address of the ECU.
	//
	// example:
	//
	// 172.16.*.**
	IpAddr *string `json:"IpAddr,omitempty" xml:"IpAddr,omitempty"`
	// The name of the ECU.
	//
	// example:
	//
	// EDAS-scaled-cluster: default cluster
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the ECU is online.
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shen****-*
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the ECU was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573635952012
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user associated with the ECU.
	//
	// example:
	//
	// edas_com***_****@******-*****.***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-wz9b246zg************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-shen****-*
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QueryApplicationStatusResponseBodyAppInfoEcuListEcu) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetAvailableCpu() *int32 {
	return s.AvailableCpu
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetAvailableMem() *int32 {
	return s.AvailableMem
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetDockerEnv() *bool {
	return s.DockerEnv
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetEcuId() *string {
	return s.EcuId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetGroupId() *string {
	return s.GroupId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetHeartbeatTime() *int64 {
	return s.HeartbeatTime
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetIpAddr() *string {
	return s.IpAddr
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetName() *string {
	return s.Name
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetOnline() *bool {
	return s.Online
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetUserId() *string {
	return s.UserId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetVpcId() *string {
	return s.VpcId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GetZoneId() *string {
	return s.ZoneId
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetAvailableCpu(v int32) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.AvailableCpu = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetAvailableMem(v int32) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.AvailableMem = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetCreateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.CreateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetDockerEnv(v bool) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.DockerEnv = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetEcuId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.EcuId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetGroupId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.GroupId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetHeartbeatTime(v int64) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.HeartbeatTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetInstanceId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.InstanceId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetIpAddr(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.IpAddr = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetName(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.Name = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetOnline(v bool) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.Online = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetRegionId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.RegionId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetUpdateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.UpdateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetUserId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.UserId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetVpcId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.VpcId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetZoneId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.ZoneId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) Validate() error {
	return dara.Validate(s)
}

type QueryApplicationStatusResponseBodyAppInfoGroupList struct {
	Group []*QueryApplicationStatusResponseBodyAppInfoGroupListGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s QueryApplicationStatusResponseBodyAppInfoGroupList) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoGroupList) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupList) GetGroup() []*QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	return s.Group
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupList) SetGroup(v []*QueryApplicationStatusResponseBodyAppInfoGroupListGroup) *QueryApplicationStatusResponseBodyAppInfoGroupList {
	s.Group = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupList) Validate() error {
	if s.Group != nil {
		for _, item := range s.Group {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryApplicationStatusResponseBodyAppInfoGroupListGroup struct {
	// The ID of the application.
	//
	// example:
	//
	// 3616cdca-4f92-4413-****-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the change process for application deployment in the instance group.
	//
	// example:
	//
	// changeorder_a**_*******_**
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// 0d247b93-8d62-4e34-****-************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The time when the instance group was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573626155185
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// 8123db90-880f-486f-****-************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the instance group.
	//
	// example:
	//
	// _DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the instance group. Valid values:
	//
	// 	- 0: default group
	//
	// 	- 1: self-managed group
	//
	// 	- 2: canary release group
	//
	// example:
	//
	// 0
	GroupType *int32 `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The version of the deployment package that was used to deploy an application in the instance group.
	//
	// example:
	//
	// 441beb18-da42-44dc-****-************
	PackageVersionId *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty"`
	// The time when the instance group was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573627441388
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryApplicationStatusResponseBodyAppInfoGroupListGroup) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoGroupListGroup) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) GetAppId() *string {
	return s.AppId
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) GetAppVersionId() *string {
	return s.AppVersionId
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) GetGroupType() *int32 {
	return s.GroupType
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) GetPackageVersionId() *string {
	return s.PackageVersionId
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetAppId(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.AppId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetAppVersionId(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.AppVersionId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetClusterId(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.ClusterId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetCreateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.CreateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetGroupId(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.GroupId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetGroupName(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.GroupName = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetGroupType(v int32) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.GroupType = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetPackageVersionId(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.PackageVersionId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetUpdateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.UpdateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) Validate() error {
	return dara.Validate(s)
}

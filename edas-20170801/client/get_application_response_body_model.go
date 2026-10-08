// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody
	GetApplication() *GetApplicationResponseBodyApplication
	SetCode(v int32) *GetApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *GetApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApplicationResponseBody
	GetRequestId() *string
}

type GetApplicationResponseBody struct {
	// The application information.
	Application *GetApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F8DFGED-K98***************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) GetApplication() *GetApplicationResponseBodyApplication {
	return s.Application
}

func (s *GetApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationResponseBody) SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody {
	s.Application = v
	return s
}

func (s *GetApplicationResponseBody) SetCode(v int32) *GetApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplicationResponseBody) SetMessage(v string) *GetApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) Validate() error {
	if s.Application != nil {
		if err := s.Application.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyApplication struct {
	// The application ID.
	//
	// example:
	//
	// cfac****-847e-4325-ad56-b5c2bc54****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The current phase of the Kubernetes application. This helps determine if the application is stable. Configuration operations are prohibited when the application is in an unstable state.
	//
	// - ready: The application is ready and can be changed.
	//
	// - progressing: The application is being changed.
	//
	// - pending: The application change is blocked.
	//
	// - failed: The application change failed.
	//
	// The ready phase is stable. Other phases are unstable.
	//
	// example:
	//
	// ready
	AppPhase *string `json:"AppPhase,omitempty" xml:"AppPhase,omitempty"`
	// The deployment type of the application:
	//
	// - War: The application is deployed from a WAR package.
	//
	// - FatJar: The application is deployed from a JAR package.
	//
	// - Empty: The application is not deployed.
	//
	// example:
	//
	// FatJar
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The ID of the container version.
	//
	// example:
	//
	// 59
	BuildPackageId *int64 `json:"BuildPackageId,omitempty" xml:"BuildPackageId,omitempty"`
	// The ID of the ECS cluster where the application is deployed.
	//
	// example:
	//
	// 5ffc5895-****-b03a-c223c6c3****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the application cluster:
	//
	// - 0: A regular Docker cluster.
	//
	// - 1: A Swarm cluster.
	//
	// - 2: An ECS cluster.
	//
	// - 3: A Kubernetes cluster.
	//
	// - 4: A Pandora application cluster that supports automatic registration.
	//
	// example:
	//
	// 2
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The UNIX timestamp when the application was created.
	//
	// example:
	//
	// 1610550324226
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the application is a Docker application:
	//
	// - false: The application is not a Docker application.
	//
	// - true: The application is a Docker application.
	//
	// example:
	//
	// false
	Dockerize *bool `json:"Dockerize,omitempty" xml:"Dockerize,omitempty"`
	// The email address.
	//
	// example:
	//
	// ****@***.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the port health check is enabled:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// If enabled, EDAS checks if the port is in use during application startup. If the port is in use, the application is considered started.
	//
	// example:
	//
	// false
	EnablePortCheck *bool `json:"EnablePortCheck,omitempty" xml:"EnablePortCheck,omitempty"`
	// Indicates whether the URL health check is enabled:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// If enabled, EDAS probes the specified URL during application startup. If the URL is accessible, the application is considered started.
	//
	// example:
	//
	// false
	EnableUrlCheck *bool `json:"EnableUrlCheck,omitempty" xml:"EnableUrlCheck,omitempty"`
	// The ID of the public-facing SLB instance attached to the application.
	//
	// example:
	//
	// lb-bp1vceck3s3b9xs6x****
	ExtSlbId *string `json:"ExtSlbId,omitempty" xml:"ExtSlbId,omitempty"`
	// The public IP address of the SLB instance attached to the application.
	//
	// example:
	//
	// 47.114.xxx.xx
	ExtSlbIp *string `json:"ExtSlbIp,omitempty" xml:"ExtSlbIp,omitempty"`
	// The name of the public-facing SLB instance attached to the application.
	//
	// example:
	//
	// aa8eee383db084f42aebc4d9f52c****
	ExtSlbName *string `json:"ExtSlbName,omitempty" xml:"ExtSlbName,omitempty"`
	// Indicates whether the current user has management permissions on the application. This parameter is available only in RAM authentication mode.
	//
	// example:
	//
	// true
	HaveManageAccess *string `json:"HaveManageAccess,omitempty" xml:"HaveManageAccess,omitempty"`
	// The health check URL of the application.
	//
	// example:
	//
	// http://127.0.0.1:8080/xyz.html
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The number of instances in the application.
	//
	// example:
	//
	// 1
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The memory size for the application instance, in MB.
	//
	// example:
	//
	// 0
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace to which the application belongs.
	//
	// example:
	//
	// doc-test
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	// The creator of the application.
	//
	// example:
	//
	// ouou@117274586608****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The service port of the application.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region where the application is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekz****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of running application instances.
	//
	// example:
	//
	// 1
	RunningInstanceCount *int32 `json:"RunningInstanceCount,omitempty" xml:"RunningInstanceCount,omitempty"`
	// The ID of the internal-facing SLB instance attached to the application.
	//
	// example:
	//
	// lb-bp****ck3s3b9xs6x****
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// Information about the internal-facing SLB instance attached to the application.
	//
	// example:
	//
	// test
	SlbInfo *string `json:"SlbInfo,omitempty" xml:"SlbInfo,omitempty"`
	// The IP address of the internal-facing SLB instance attached to the application.
	//
	// example:
	//
	// 192.***.***.***
	SlbIp *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	// The name of the internal-facing SLB instance attached to the application.
	//
	// example:
	//
	// test
	SlbName *string `json:"SlbName,omitempty" xml:"SlbName,omitempty"`
	// The port of the internal-facing SLB instance attached to the application.
	//
	// example:
	//
	// 80
	SlbPort *int32 `json:"SlbPort,omitempty" xml:"SlbPort,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// test@dd******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The workload type used to create the application. Supported types are Deployment and StatefulSet. This parameter does not apply to ECS applications.
	//
	// example:
	//
	// StatefulSet
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s GetApplicationResponseBodyApplication) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplication) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationResponseBodyApplication) GetAppPhase() *string {
	return s.AppPhase
}

func (s *GetApplicationResponseBodyApplication) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *GetApplicationResponseBodyApplication) GetBuildPackageId() *int64 {
	return s.BuildPackageId
}

func (s *GetApplicationResponseBodyApplication) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetApplicationResponseBodyApplication) GetClusterType() *string {
	return s.ClusterType
}

func (s *GetApplicationResponseBodyApplication) GetCpu() *int32 {
	return s.Cpu
}

func (s *GetApplicationResponseBodyApplication) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetApplicationResponseBodyApplication) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyApplication) GetDockerize() *bool {
	return s.Dockerize
}

func (s *GetApplicationResponseBodyApplication) GetEmail() *string {
	return s.Email
}

func (s *GetApplicationResponseBodyApplication) GetEnablePortCheck() *bool {
	return s.EnablePortCheck
}

func (s *GetApplicationResponseBodyApplication) GetEnableUrlCheck() *bool {
	return s.EnableUrlCheck
}

func (s *GetApplicationResponseBodyApplication) GetExtSlbId() *string {
	return s.ExtSlbId
}

func (s *GetApplicationResponseBodyApplication) GetExtSlbIp() *string {
	return s.ExtSlbIp
}

func (s *GetApplicationResponseBodyApplication) GetExtSlbName() *string {
	return s.ExtSlbName
}

func (s *GetApplicationResponseBodyApplication) GetHaveManageAccess() *string {
	return s.HaveManageAccess
}

func (s *GetApplicationResponseBodyApplication) GetHealthCheckUrl() *string {
	return s.HealthCheckUrl
}

func (s *GetApplicationResponseBodyApplication) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *GetApplicationResponseBodyApplication) GetMemory() *int32 {
	return s.Memory
}

func (s *GetApplicationResponseBodyApplication) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyApplication) GetNameSpace() *string {
	return s.NameSpace
}

func (s *GetApplicationResponseBodyApplication) GetOwner() *string {
	return s.Owner
}

func (s *GetApplicationResponseBodyApplication) GetPort() *int32 {
	return s.Port
}

func (s *GetApplicationResponseBodyApplication) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApplicationResponseBodyApplication) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetApplicationResponseBodyApplication) GetRunningInstanceCount() *int32 {
	return s.RunningInstanceCount
}

func (s *GetApplicationResponseBodyApplication) GetSlbId() *string {
	return s.SlbId
}

func (s *GetApplicationResponseBodyApplication) GetSlbInfo() *string {
	return s.SlbInfo
}

func (s *GetApplicationResponseBodyApplication) GetSlbIp() *string {
	return s.SlbIp
}

func (s *GetApplicationResponseBodyApplication) GetSlbName() *string {
	return s.SlbName
}

func (s *GetApplicationResponseBodyApplication) GetSlbPort() *int32 {
	return s.SlbPort
}

func (s *GetApplicationResponseBodyApplication) GetUserId() *string {
	return s.UserId
}

func (s *GetApplicationResponseBodyApplication) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *GetApplicationResponseBodyApplication) SetAppId(v string) *GetApplicationResponseBodyApplication {
	s.AppId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAppPhase(v string) *GetApplicationResponseBodyApplication {
	s.AppPhase = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationType(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetBuildPackageId(v int64) *GetApplicationResponseBodyApplication {
	s.BuildPackageId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetClusterId(v string) *GetApplicationResponseBodyApplication {
	s.ClusterId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetClusterType(v string) *GetApplicationResponseBodyApplication {
	s.ClusterType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCpu(v int32) *GetApplicationResponseBodyApplication {
	s.Cpu = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCreateTime(v int64) *GetApplicationResponseBodyApplication {
	s.CreateTime = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetDescription(v string) *GetApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetDockerize(v bool) *GetApplicationResponseBodyApplication {
	s.Dockerize = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetEmail(v string) *GetApplicationResponseBodyApplication {
	s.Email = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetEnablePortCheck(v bool) *GetApplicationResponseBodyApplication {
	s.EnablePortCheck = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetEnableUrlCheck(v bool) *GetApplicationResponseBodyApplication {
	s.EnableUrlCheck = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetExtSlbId(v string) *GetApplicationResponseBodyApplication {
	s.ExtSlbId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetExtSlbIp(v string) *GetApplicationResponseBodyApplication {
	s.ExtSlbIp = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetExtSlbName(v string) *GetApplicationResponseBodyApplication {
	s.ExtSlbName = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetHaveManageAccess(v string) *GetApplicationResponseBodyApplication {
	s.HaveManageAccess = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetHealthCheckUrl(v string) *GetApplicationResponseBodyApplication {
	s.HealthCheckUrl = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetInstanceCount(v int32) *GetApplicationResponseBodyApplication {
	s.InstanceCount = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetMemory(v int32) *GetApplicationResponseBodyApplication {
	s.Memory = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetName(v string) *GetApplicationResponseBodyApplication {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetNameSpace(v string) *GetApplicationResponseBodyApplication {
	s.NameSpace = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetOwner(v string) *GetApplicationResponseBodyApplication {
	s.Owner = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetPort(v int32) *GetApplicationResponseBodyApplication {
	s.Port = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetRegionId(v string) *GetApplicationResponseBodyApplication {
	s.RegionId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetResourceGroupId(v string) *GetApplicationResponseBodyApplication {
	s.ResourceGroupId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetRunningInstanceCount(v int32) *GetApplicationResponseBodyApplication {
	s.RunningInstanceCount = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetSlbId(v string) *GetApplicationResponseBodyApplication {
	s.SlbId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetSlbInfo(v string) *GetApplicationResponseBodyApplication {
	s.SlbInfo = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetSlbIp(v string) *GetApplicationResponseBodyApplication {
	s.SlbIp = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetSlbName(v string) *GetApplicationResponseBodyApplication {
	s.SlbName = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetSlbPort(v int32) *GetApplicationResponseBodyApplication {
	s.SlbPort = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetUserId(v string) *GetApplicationResponseBodyApplication {
	s.UserId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetWorkloadType(v string) *GetApplicationResponseBodyApplication {
	s.WorkloadType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationBaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplcation(v *UpdateApplicationBaseInfoResponseBodyApplcation) *UpdateApplicationBaseInfoResponseBody
	GetApplcation() *UpdateApplicationBaseInfoResponseBodyApplcation
	SetCode(v int32) *UpdateApplicationBaseInfoResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateApplicationBaseInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApplicationBaseInfoResponseBody
	GetRequestId() *string
}

type UpdateApplicationBaseInfoResponseBody struct {
	// The applications that you want to modify.
	Applcation *UpdateApplicationBaseInfoResponseBodyApplcation `json:"Applcation,omitempty" xml:"Applcation,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-**************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationBaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationBaseInfoResponseBody) GetApplcation() *UpdateApplicationBaseInfoResponseBodyApplcation {
	return s.Applcation
}

func (s *UpdateApplicationBaseInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateApplicationBaseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApplicationBaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationBaseInfoResponseBody) SetApplcation(v *UpdateApplicationBaseInfoResponseBodyApplcation) *UpdateApplicationBaseInfoResponseBody {
	s.Applcation = v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBody) SetCode(v int32) *UpdateApplicationBaseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBody) SetMessage(v string) *UpdateApplicationBaseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBody) SetRequestId(v string) *UpdateApplicationBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBody) Validate() error {
	if s.Applcation != nil {
		if err := s.Applcation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationBaseInfoResponseBodyApplcation struct {
	// The ID of the application.
	//
	// example:
	//
	// c627c157-560d-43ff-****-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The deployment type of the application. Valid values:
	//
	// 	- War: The application is deployed by using a WAR package.
	//
	// 	- FatJar: The application is deployed by using a JAR package.
	//
	// 	- Image: The application is deployed by using an image.
	//
	// 	- If this parameter is empty, the application is not deployed.
	//
	// example:
	//
	// FatJar
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The build package number of Enterprise Distributed Application Service (EDAS) Container.
	//
	// example:
	//
	// 58
	BuildPackageId *int64 `json:"BuildPackageId,omitempty" xml:"BuildPackageId,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// d7730a49-629a-47bd-****-f45eb01f****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- 0: normal Docker cluster
	//
	// 	- 1: Swarm cluster
	//
	// 	- 2: ECS cluster
	//
	// 	- 3: self-managed Kubernetes cluster in EDAS
	//
	// 	- 4: cluster in which Pandora automatically registers applications
	//
	// 	- 5: Container Service for Kubernetes (ACK) clusters
	//
	// example:
	//
	// 2
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 0
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the application was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1577259573911
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the application is a Docker application.
	//
	// example:
	//
	// false
	Dockerize *bool `json:"Dockerize,omitempty" xml:"Dockerize,omitempty"`
	// The ID of the Internet-facing SLB instance.
	//
	// example:
	//
	// ace93*******
	ExtSlbId *string `json:"ExtSlbId,omitempty" xml:"ExtSlbId,omitempty"`
	// The IP address of the Internet-facing Server Load Balancer (SLB) instance.
	//
	// example:
	//
	// 39.97.XX.XX
	ExtSlbIp *string `json:"ExtSlbIp,omitempty" xml:"ExtSlbIp,omitempty"`
	// The name of the Internet-facing SLB instance.
	//
	// example:
	//
	// test
	ExtSlbName *string `json:"ExtSlbName,omitempty" xml:"ExtSlbName,omitempty"`
	// The health check URL.
	//
	// example:
	//
	// http://127.0.XX.XX:8080/_etc.html
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The number of application instances.
	//
	// example:
	//
	// 1
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The size of memory configured for an application instance. Unit: MB.
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
	// The owner of the application.
	//
	// example:
	//
	// test@aliyun_XXX.com
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The port used by the application.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing:****
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of running application instances.
	//
	// example:
	//
	// 1
	RunningInstanceCount *int32 `json:"RunningInstanceCount,omitempty" xml:"RunningInstanceCount,omitempty"`
	// The ID of the internal-facing SLB instance.
	//
	// example:
	//
	// a3d4*******
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The IP address of the internal-facing SLB instance.
	//
	// example:
	//
	// 192.168.XX.XX
	SlbIp *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	// The name of the internal-facing SLB instance.
	//
	// example:
	//
	// test
	SlbName *string `json:"SlbName,omitempty" xml:"SlbName,omitempty"`
	// The port used by the internal-facing SLB instance.
	//
	// example:
	//
	// 80
	SlbPort *int32 `json:"SlbPort,omitempty" xml:"SlbPort,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// edas_com***@****.***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateApplicationBaseInfoResponseBodyApplcation) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationBaseInfoResponseBodyApplcation) GoString() string {
	return s.String()
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetAppId() *string {
	return s.AppId
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetBuildPackageId() *int64 {
	return s.BuildPackageId
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetCpu() *int32 {
	return s.Cpu
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetDockerize() *bool {
	return s.Dockerize
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetExtSlbId() *string {
	return s.ExtSlbId
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetExtSlbIp() *string {
	return s.ExtSlbIp
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetExtSlbName() *string {
	return s.ExtSlbName
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetHealthCheckUrl() *string {
	return s.HealthCheckUrl
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetMemory() *int32 {
	return s.Memory
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetOwner() *string {
	return s.Owner
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetPort() *int32 {
	return s.Port
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetRunningInstanceCount() *int32 {
	return s.RunningInstanceCount
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetSlbId() *string {
	return s.SlbId
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetSlbIp() *string {
	return s.SlbIp
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetSlbName() *string {
	return s.SlbName
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetSlbPort() *int32 {
	return s.SlbPort
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) GetUserId() *string {
	return s.UserId
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetAppId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetApplicationType(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ApplicationType = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetBuildPackageId(v int64) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.BuildPackageId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetClusterId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ClusterId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetClusterType(v int32) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ClusterType = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetCpu(v int32) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Cpu = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetCreateTime(v int64) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.CreateTime = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetDescription(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Description = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetDockerize(v bool) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Dockerize = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetExtSlbId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ExtSlbId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetExtSlbIp(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ExtSlbIp = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetExtSlbName(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ExtSlbName = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetHealthCheckUrl(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.HealthCheckUrl = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetInstanceCount(v int32) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.InstanceCount = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetMemory(v int32) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Memory = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetName(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Name = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetOwner(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Owner = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetPort(v int32) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Port = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetRegionId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetRunningInstanceCount(v int32) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.RunningInstanceCount = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetSlbId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.SlbId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetSlbIp(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.SlbIp = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetSlbName(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.SlbName = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetSlbPort(v int32) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.SlbPort = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetUserId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.UserId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) Validate() error {
	return dara.Validate(s)
}

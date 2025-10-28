// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationList(v *ListApplicationResponseBodyApplicationList) *ListApplicationResponseBody
	GetApplicationList() *ListApplicationResponseBodyApplicationList
	SetCode(v int32) *ListApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *ListApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApplicationResponseBody
	GetRequestId() *string
}

type ListApplicationResponseBody struct {
	// The information about applications.
	ApplicationList *ListApplicationResponseBodyApplicationList `json:"ApplicationList,omitempty" xml:"ApplicationList,omitempty" type:"Struct"`
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
	// The request ID.
	//
	// example:
	//
	// 5d6fa0bc-cc3**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBody) GetApplicationList() *ListApplicationResponseBodyApplicationList {
	return s.ApplicationList
}

func (s *ListApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationResponseBody) SetApplicationList(v *ListApplicationResponseBodyApplicationList) *ListApplicationResponseBody {
	s.ApplicationList = v
	return s
}

func (s *ListApplicationResponseBody) SetCode(v int32) *ListApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationResponseBody) SetMessage(v string) *ListApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationResponseBody) SetRequestId(v string) *ListApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationResponseBody) Validate() error {
	if s.ApplicationList != nil {
		if err := s.ApplicationList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationResponseBodyApplicationList struct {
	Application []*ListApplicationResponseBodyApplicationListApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
}

func (s ListApplicationResponseBodyApplicationList) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationResponseBodyApplicationList) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBodyApplicationList) GetApplication() []*ListApplicationResponseBodyApplicationListApplication {
	return s.Application
}

func (s *ListApplicationResponseBodyApplicationList) SetApplication(v []*ListApplicationResponseBodyApplicationListApplication) *ListApplicationResponseBodyApplicationList {
	s.Application = v
	return s
}

func (s *ListApplicationResponseBodyApplicationList) Validate() error {
	if s.Application != nil {
		for _, item := range s.Application {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationResponseBodyApplicationListApplication struct {
	// The ID of the application.
	//
	// example:
	//
	// 00ee517d-dd7d-4d4e-****-************
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
	// c37aec2a-bcca-4ec1-****-************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the cluster in which the application is deployed. Valid values:
	//
	// 	- **2**: Elastic Compute Service (ECS) cluster
	//
	// 	- **3**: self-managed Kubernetes cluster in EDAS
	//
	// 	- **5**: Container Service for Kubernetes (ACK) cluster
	//
	// example:
	//
	// 2
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 1664208000000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The IP address of the Internet-facing SLB instance.
	//
	// example:
	//
	// 100.100.70.***
	ExtSlbIp *string `json:"ExtSlbIp,omitempty" xml:"ExtSlbIp,omitempty"`
	// The listener port of the Internet-facing SLB instance.
	//
	// example:
	//
	// 8080
	ExtSlbListenerPort *int32 `json:"ExtSlbListenerPort,omitempty" xml:"ExtSlbListenerPort,omitempty"`
	// The number of application instances.
	//
	// example:
	//
	// 5
	Instances *int32 `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// example:
	//
	// default
	K8sNamespace *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// doc-test-consumer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the microservices namespace.
	//
	// example:
	//
	// cn-hangzhou:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The service port of the application.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the application.
	//
	// example:
	//
	// cn-beijing:docTes
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek24j4s4b*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of running application instances.
	//
	// example:
	//
	// 0
	RunningInstanceCount *int32 `json:"RunningInstanceCount,omitempty" xml:"RunningInstanceCount,omitempty"`
	// The IP address of the internal-facing Server Load Balancer (SLB) instance.
	//
	// example:
	//
	// 192.168.0.***
	SlbIp *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	// The listener port of the internal-facing SLB instance.
	//
	// example:
	//
	// 8088
	SlbListenerPort *int32 `json:"SlbListenerPort,omitempty" xml:"SlbListenerPort,omitempty"`
	// The port of the internal-facing SLB instance.
	//
	// example:
	//
	// 80
	SlbPort *int32 `json:"SlbPort,omitempty" xml:"SlbPort,omitempty"`
	// The state of the application. Valid values:
	//
	// 	- RUNNING: The application is running.
	//
	// 	- STOPPED: The application is stopped.
	//
	// 	- DEPLOYING: The application is being deployed.
	//
	// 	- DELETING: The application is being deleted.
	//
	// example:
	//
	// RUNNING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListApplicationResponseBodyApplicationListApplication) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationResponseBodyApplicationListApplication) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetBuildPackageId() *int64 {
	return s.BuildPackageId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetExtSlbIp() *string {
	return s.ExtSlbIp
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetExtSlbListenerPort() *int32 {
	return s.ExtSlbListenerPort
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetInstances() *int32 {
	return s.Instances
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetK8sNamespace() *string {
	return s.K8sNamespace
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetName() *string {
	return s.Name
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetPort() *int32 {
	return s.Port
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetRunningInstanceCount() *int32 {
	return s.RunningInstanceCount
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetSlbIp() *string {
	return s.SlbIp
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetSlbListenerPort() *int32 {
	return s.SlbListenerPort
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetSlbPort() *int32 {
	return s.SlbPort
}

func (s *ListApplicationResponseBodyApplicationListApplication) GetState() *string {
	return s.State
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetAppId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.AppId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetApplicationType(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.ApplicationType = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetBuildPackageId(v int64) *ListApplicationResponseBodyApplicationListApplication {
	s.BuildPackageId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetClusterId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.ClusterId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetClusterType(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.ClusterType = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetCreateTime(v int64) *ListApplicationResponseBodyApplicationListApplication {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetExtSlbIp(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.ExtSlbIp = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetExtSlbListenerPort(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.ExtSlbListenerPort = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetInstances(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.Instances = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetK8sNamespace(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.K8sNamespace = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetName(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.Name = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetNamespaceId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.NamespaceId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetPort(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.Port = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetRegionId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.RegionId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetResourceGroupId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetRunningInstanceCount(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.RunningInstanceCount = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetSlbIp(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.SlbIp = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetSlbListenerPort(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.SlbListenerPort = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetSlbPort(v int32) *ListApplicationResponseBodyApplicationListApplication {
	s.SlbPort = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetState(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.State = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) Validate() error {
	return dara.Validate(s)
}

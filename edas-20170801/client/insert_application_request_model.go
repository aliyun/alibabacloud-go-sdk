// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *InsertApplicationRequest
	GetApplicationName() *string
	SetBuildPackId(v int32) *InsertApplicationRequest
	GetBuildPackId() *int32
	SetClusterId(v string) *InsertApplicationRequest
	GetClusterId() *string
	SetComponentIds(v string) *InsertApplicationRequest
	GetComponentIds() *string
	SetCpu(v int32) *InsertApplicationRequest
	GetCpu() *int32
	SetDescription(v string) *InsertApplicationRequest
	GetDescription() *string
	SetEcuInfo(v string) *InsertApplicationRequest
	GetEcuInfo() *string
	SetEnablePortCheck(v bool) *InsertApplicationRequest
	GetEnablePortCheck() *bool
	SetEnableUrlCheck(v bool) *InsertApplicationRequest
	GetEnableUrlCheck() *bool
	SetHealthCheckUrl(v string) *InsertApplicationRequest
	GetHealthCheckUrl() *string
	SetHooks(v string) *InsertApplicationRequest
	GetHooks() *string
	SetJdk(v string) *InsertApplicationRequest
	GetJdk() *string
	SetJvmOptions(v string) *InsertApplicationRequest
	GetJvmOptions() *string
	SetLogicalRegionId(v string) *InsertApplicationRequest
	GetLogicalRegionId() *string
	SetMaxHeapSize(v int32) *InsertApplicationRequest
	GetMaxHeapSize() *int32
	SetMaxPermSize(v int32) *InsertApplicationRequest
	GetMaxPermSize() *int32
	SetMem(v int32) *InsertApplicationRequest
	GetMem() *int32
	SetMinHeapSize(v int32) *InsertApplicationRequest
	GetMinHeapSize() *int32
	SetPackageType(v string) *InsertApplicationRequest
	GetPackageType() *string
	SetReservedPortStr(v string) *InsertApplicationRequest
	GetReservedPortStr() *string
	SetResourceGroupId(v string) *InsertApplicationRequest
	GetResourceGroupId() *string
	SetWebContainer(v string) *InsertApplicationRequest
	GetWebContainer() *string
}

type InsertApplicationRequest struct {
	// The name of the application. The name can contain only digits, letters, hyphens (-), and underscores (_) and must start with a letter. The name can be up to 36 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// hello-edas-test-1
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The build package number of EDAS Container. This parameter is required if you create a High-Speed Service Framework (HSF) application. You can query the build package number by using one of the following methods:
	//
	// 	- Call the ListBuildPack operation. For more information, see [ListBuildPack](https://help.aliyun.com/document_detail/149391.html).
	//
	// 	- Obtain the value in the **Build package number*	- column of the [Release notes for EDAS Container](https://help.aliyun.com/document_detail/92614.html) topic.
	//
	// example:
	//
	// 59
	BuildPackId *int32 `json:"BuildPackId,omitempty" xml:"BuildPackId,omitempty"`
	// The ID of the ECS cluster in which you want to create the application. If you specify an ID, the application is created in the specified ECS cluster. If you leave this parameter empty, the application is created in the default cluster. We recommend that you specify this parameter.
	//
	// example:
	//
	// 13136119-f384-4f50-b76e-xxxxxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the application component. You can call the ListComponents operation to query the component IDs. For more information, see [ListComponents](https://help.aliyun.com/document_detail/97502.html).
	//
	// This parameter is required if the application runs in Apache Tomcat or in a standard Java application runtime environment. The Apache Tomcat application runtime environment is applicable to Dubbo applications that are deployed by using WAR packages. A standard Java application runtime environment is applicable to Spring Boot or Spring Cloud applications that are deployed by using JAR packages.
	//
	// Valid values for common application components:
	//
	// 	- 4: Apache Tomcat 7.0.91
	//
	// 	- 7: Apache Tomcat 8.5.42
	//
	// 	- 5: OpenJDK 1.8.x
	//
	// 	- 6: OpenJDK 1.7.x
	//
	// This parameter is available only for Java SDK 2.57.3 or later, or Python SDK 2.57.3 or later. Assume that you use an SDK that is not provided by EDAS, for example, aliyun-python-sdk-core, aliyun-java-sdk-core, and Alibaba Cloud CLI. In this case, you can directly specify this parameter.
	//
	// example:
	//
	// 7
	ComponentIds *string `json:"ComponentIds,omitempty" xml:"ComponentIds,omitempty"`
	// The number of CPU cores that can be used by the application container in a Swarm cluster. \\*\\*This parameter is deprecated.\\*\\*
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// create by edas pop api
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The value of `ecu_id` of the ECS instance to be added during scale-out. The ECU ID is the unique identity for an ECS instance that is imported to EDAS. Separate multiple values of `ecu_id` with commas (,). You can call the ListScaleOutEcu operation to query the value of `ecu_id`. For more information, see [ListScaleOutEcu](https://help.aliyun.com/document_detail/149371.html).
	//
	// example:
	//
	// 07bd417a-b863-477d-****-************
	EcuInfo *string `json:"EcuInfo,omitempty" xml:"EcuInfo,omitempty"`
	// Specifies whether to enable the port health check. Valid values:
	//
	// 	- **true**: enable the port health check.
	//
	// 	- **false**: does not enable the port health check.
	//
	// example:
	//
	// true
	EnablePortCheck *bool `json:"EnablePortCheck,omitempty" xml:"EnablePortCheck,omitempty"`
	// Specifies whether to enable the URL health check. Valid values:
	//
	// 	- **true**: enables the URL health check.
	//
	// 	- **false**: does not enable the URL health check.
	//
	// example:
	//
	// true
	EnableUrlCheck *bool `json:"EnableUrlCheck,omitempty" xml:"EnableUrlCheck,omitempty"`
	// The health check URL of the application. This parameter is equivalent to the HealthCheckURL parameter.
	//
	// example:
	//
	// http://127.0.0.1:8080/_ehc.html
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The script to mount. Set the value in the JSON format. Example: `[{"ignoreFail":false,"name":"postprepareInstanceEnvironmentOnScaleOut","script":"ls"},{"ignoreFail":true,"name":"postdeleteInstanceDataOnScaleIn","script":""},{"ignoreFail":true,"name":"prestartInstance","script":""},{"ignoreFail":true,"name":"poststartInstance","script":""},{"ignoreFail":true,"name":"prestopInstance","script":""},{"ignoreFail":true,"name":"poststopInstance","script":""}]`
	//
	// example:
	//
	// [{"ignoreFail":false,"name":"postprepareInstanceEnvironmentOnScaleOut","script":"ls"}]
	Hooks *string `json:"Hooks,omitempty" xml:"Hooks,omitempty"`
	// The version of the Java Development Kit (JDK) used to deploy the application. **This parameter is deprecated.
	//
	// example:
	//
	// 8
	Jdk *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	// The custom parameters.
	//
	// example:
	//
	// -Dproperty=value
	JvmOptions *string `json:"JvmOptions,omitempty" xml:"JvmOptions,omitempty"`
	// The ID of the microservices namespace. To query the ID of a microservices namespace, you can choose **Resource Management*	- > **Microservice Namespaces*	- in the left-side navigation pane of the EDAS console or call the ListUserDefineRegion operation. For more information, see [ListUserDefineRegion](https://help.aliyun.com/document_detail/149377.html).
	//
	// 	- This parameter is required if the cluster you specify is not deployed in the default microservices namespace. Otherwise, the message `application regionId is different with cluster regionId!` appears.
	//
	// 	- If the cluster you specify is deployed in the default microservices namespace, you do not need to specify this parameter. Set this parameter to the ID of the microservices namespace in which the cluster you specify is deployed.
	//
	// example:
	//
	// cn-beijing:prod
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	// The maximum size of the heap memory. Unit: MB.
	//
	// example:
	//
	// 1000
	MaxHeapSize *int32 `json:"MaxHeapSize,omitempty" xml:"MaxHeapSize,omitempty"`
	// The size of the permanent generation heap memory. Unit: MB.
	//
	// example:
	//
	// 200
	MaxPermSize *int32 `json:"MaxPermSize,omitempty" xml:"MaxPermSize,omitempty"`
	// The memory size that can be used by the application container in a Swarm cluster. \\*\\*This parameter is deprecated.\\*\\*
	//
	// example:
	//
	// 2048
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The initial size of the heap memory. Unit: MB.
	//
	// example:
	//
	// 500
	MinHeapSize *int32 `json:"MinHeapSize,omitempty" xml:"MinHeapSize,omitempty"`
	// The type of the application deployment package. Valid values: war and jar.
	//
	// example:
	//
	// war
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The reserved port for the application. This parameter is deprecated.
	//
	// example:
	//
	// 8090
	ReservedPortStr *string `json:"ReservedPortStr,omitempty" xml:"ReservedPortStr,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek24j4s4b*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The version of Apache Tomcat. **This parameter is deprecated.
	//
	// example:
	//
	// 4
	WebContainer *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
}

func (s InsertApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertApplicationRequest) GoString() string {
	return s.String()
}

func (s *InsertApplicationRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *InsertApplicationRequest) GetBuildPackId() *int32 {
	return s.BuildPackId
}

func (s *InsertApplicationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InsertApplicationRequest) GetComponentIds() *string {
	return s.ComponentIds
}

func (s *InsertApplicationRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *InsertApplicationRequest) GetDescription() *string {
	return s.Description
}

func (s *InsertApplicationRequest) GetEcuInfo() *string {
	return s.EcuInfo
}

func (s *InsertApplicationRequest) GetEnablePortCheck() *bool {
	return s.EnablePortCheck
}

func (s *InsertApplicationRequest) GetEnableUrlCheck() *bool {
	return s.EnableUrlCheck
}

func (s *InsertApplicationRequest) GetHealthCheckUrl() *string {
	return s.HealthCheckUrl
}

func (s *InsertApplicationRequest) GetHooks() *string {
	return s.Hooks
}

func (s *InsertApplicationRequest) GetJdk() *string {
	return s.Jdk
}

func (s *InsertApplicationRequest) GetJvmOptions() *string {
	return s.JvmOptions
}

func (s *InsertApplicationRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *InsertApplicationRequest) GetMaxHeapSize() *int32 {
	return s.MaxHeapSize
}

func (s *InsertApplicationRequest) GetMaxPermSize() *int32 {
	return s.MaxPermSize
}

func (s *InsertApplicationRequest) GetMem() *int32 {
	return s.Mem
}

func (s *InsertApplicationRequest) GetMinHeapSize() *int32 {
	return s.MinHeapSize
}

func (s *InsertApplicationRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *InsertApplicationRequest) GetReservedPortStr() *string {
	return s.ReservedPortStr
}

func (s *InsertApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *InsertApplicationRequest) GetWebContainer() *string {
	return s.WebContainer
}

func (s *InsertApplicationRequest) SetApplicationName(v string) *InsertApplicationRequest {
	s.ApplicationName = &v
	return s
}

func (s *InsertApplicationRequest) SetBuildPackId(v int32) *InsertApplicationRequest {
	s.BuildPackId = &v
	return s
}

func (s *InsertApplicationRequest) SetClusterId(v string) *InsertApplicationRequest {
	s.ClusterId = &v
	return s
}

func (s *InsertApplicationRequest) SetComponentIds(v string) *InsertApplicationRequest {
	s.ComponentIds = &v
	return s
}

func (s *InsertApplicationRequest) SetCpu(v int32) *InsertApplicationRequest {
	s.Cpu = &v
	return s
}

func (s *InsertApplicationRequest) SetDescription(v string) *InsertApplicationRequest {
	s.Description = &v
	return s
}

func (s *InsertApplicationRequest) SetEcuInfo(v string) *InsertApplicationRequest {
	s.EcuInfo = &v
	return s
}

func (s *InsertApplicationRequest) SetEnablePortCheck(v bool) *InsertApplicationRequest {
	s.EnablePortCheck = &v
	return s
}

func (s *InsertApplicationRequest) SetEnableUrlCheck(v bool) *InsertApplicationRequest {
	s.EnableUrlCheck = &v
	return s
}

func (s *InsertApplicationRequest) SetHealthCheckUrl(v string) *InsertApplicationRequest {
	s.HealthCheckUrl = &v
	return s
}

func (s *InsertApplicationRequest) SetHooks(v string) *InsertApplicationRequest {
	s.Hooks = &v
	return s
}

func (s *InsertApplicationRequest) SetJdk(v string) *InsertApplicationRequest {
	s.Jdk = &v
	return s
}

func (s *InsertApplicationRequest) SetJvmOptions(v string) *InsertApplicationRequest {
	s.JvmOptions = &v
	return s
}

func (s *InsertApplicationRequest) SetLogicalRegionId(v string) *InsertApplicationRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *InsertApplicationRequest) SetMaxHeapSize(v int32) *InsertApplicationRequest {
	s.MaxHeapSize = &v
	return s
}

func (s *InsertApplicationRequest) SetMaxPermSize(v int32) *InsertApplicationRequest {
	s.MaxPermSize = &v
	return s
}

func (s *InsertApplicationRequest) SetMem(v int32) *InsertApplicationRequest {
	s.Mem = &v
	return s
}

func (s *InsertApplicationRequest) SetMinHeapSize(v int32) *InsertApplicationRequest {
	s.MinHeapSize = &v
	return s
}

func (s *InsertApplicationRequest) SetPackageType(v string) *InsertApplicationRequest {
	s.PackageType = &v
	return s
}

func (s *InsertApplicationRequest) SetReservedPortStr(v string) *InsertApplicationRequest {
	s.ReservedPortStr = &v
	return s
}

func (s *InsertApplicationRequest) SetResourceGroupId(v string) *InsertApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *InsertApplicationRequest) SetWebContainer(v string) *InsertApplicationRequest {
	s.WebContainer = &v
	return s
}

func (s *InsertApplicationRequest) Validate() error {
	return dara.Validate(s)
}

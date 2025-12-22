// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeApplicationConfigResponseBody
	GetCode() *string
	SetData(v *DescribeApplicationConfigResponseBodyData) *DescribeApplicationConfigResponseBody
	GetData() *DescribeApplicationConfigResponseBodyData
	SetErrorCode(v string) *DescribeApplicationConfigResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeApplicationConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApplicationConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeApplicationConfigResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeApplicationConfigResponseBody
	GetTraceId() *string
}

type DescribeApplicationConfigResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the application.
	Data *DescribeApplicationConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the configurations of an application were obtained. Valid values:
	//
	// 	- **true**: The configurations were obtained.
	//
	// 	- **false**: The configurations failed to be obtained.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// ac1a0b2215622246421415014e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApplicationConfigResponseBody) GetData() *DescribeApplicationConfigResponseBodyData {
	return s.Data
}

func (s *DescribeApplicationConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeApplicationConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApplicationConfigResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeApplicationConfigResponseBody) SetCode(v string) *DescribeApplicationConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetData(v *DescribeApplicationConfigResponseBodyData) *DescribeApplicationConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetErrorCode(v string) *DescribeApplicationConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetMessage(v string) *DescribeApplicationConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetRequestId(v string) *DescribeApplicationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetSuccess(v bool) *DescribeApplicationConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationConfigResponseBody) SetTraceId(v string) *DescribeApplicationConfigResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationConfigResponseBodyData struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is used to pull images across accounts. For more information, see [Pull images across Alibaba Cloud accounts](https://help.aliyun.com/document_detail/190675.html) and [Grant permissions across Alibaba Cloud accounts by using a RAM role](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The ID of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// cri-xxxxxx
	AcrInstanceId           *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	AgentVersion            *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	AlbIngressReadinessGate *string `json:"AlbIngressReadinessGate,omitempty" xml:"AlbIngressReadinessGate,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// The application description.
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The SAE application type.
	//
	// 	- micro_service
	//
	// 	- web
	//
	// 	- job
	//
	// example:
	//
	// micro_service
	AppSource *string `json:"AppSource,omitempty" xml:"AppSource,omitempty"`
	// Indicates whether an elastic IP address (EIP) is associated with the application instance. Valid values:
	//
	// 	- **true**: The EIP is associated with the application instance.
	//
	// 	- **false**: The EIP is not associated with the application instance.
	//
	// example:
	//
	// true
	AssociateEip *bool `json:"AssociateEip,omitempty" xml:"AssociateEip,omitempty"`
	// The base app ID.
	//
	// example:
	//
	// 8c573618-8d72-4407-baf4-f7b64b******
	BaseAppId *string `json:"BaseAppId,omitempty" xml:"BaseAppId,omitempty"`
	// The interval between batches in a phased release. Unit: seconds.
	//
	// example:
	//
	// 10
	BatchWaitTime *int32 `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 495fc79c-ae61-4600-866d-a09d68******
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CmsServiceId *string `json:"CmsServiceId,omitempty" xml:"CmsServiceId,omitempty"`
	// The command that is used to start the image. The command must be an existing executable object in the container. Example:
	//
	// ```
	//
	// command:
	//
	//       - echo
	//
	//       - abc
	//
	//       - >
	//
	//       - file0
	//
	// ```
	//
	// In this example, the Command parameter is set to `Command="echo", CommandArgs=["abc", ">", "file0"]`.
	//
	// example:
	//
	// echo
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The parameters of the image startup command. The CommandArgs parameter contains the parameters that are required for the **Command*	- parameter. Format:
	//
	// `["a","b"]`
	//
	// In the preceding **Command*	- example, the CommandArgs parameter is set to `CommandArgs=["abc", ">", "file0"]`. The data type of `["abc", ">", "file0"]` must be an array of strings in the JSON format. You do not need to configure this parameter if it does not exist in the Command parameter.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The details of the ConfigMap.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The CPU specifications that are required for each instance. Unit: millicores. You cannot set this parameter to 0. Valid values:
	//
	// 	- **500**
	//
	// 	- **1000**
	//
	// 	- **2000**
	//
	// 	- **4000**
	//
	// 	- **8000**
	//
	// 	- **16000**
	//
	// 	- **32000**
	//
	// example:
	//
	// 1000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The custom mappings between hostnames and IP addresses in the container. Valid values:
	//
	// 	- **hostName**: the domain name or hostname.
	//
	// 	- **ip**: the IP address.
	//
	// example:
	//
	// [{"hostName":"test.host.name","ip":"0.0.0.0"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The type of custom image. Set to empty string if using pre-built image.
	//
	// - internet: public network image.
	//
	// - intranet: private network image.
	//
	// example:
	//
	// internet
	CustomImageNetworkType *string `json:"CustomImageNetworkType,omitempty" xml:"CustomImageNetworkType,omitempty"`
	DeploymentName         *string `json:"DeploymentName,omitempty" xml:"DeploymentName,omitempty"`
	// The disk size. Unit: GB.
	//
	// example:
	//
	// 20
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The version of .NET.
	//
	// - .NET 3.1
	//
	// - .NET 5.0
	//
	// - .NET 6.0
	//
	// - .NET 7.0
	//
	// - .NET 8.0
	//
	// example:
	//
	// .NET 8.0
	Dotnet *string `json:"Dotnet,omitempty" xml:"Dotnet,omitempty"`
	// The version of the container, such as Ali-Tomcat, in which an application developed based on High-speed Service Framework (HSF) is deployed.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string                                                  `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	EmptyDirDesc         []*DescribeApplicationConfigResponseBodyDataEmptyDirDesc `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// Indicates whether access to Application High Availability Service (AHAS) is enabled. Valid values:
	//
	// 	- **true**: Access to AHAS is enabled.
	//
	// 	- **false**: Access to AHAS is disabled.
	//
	// example:
	//
	// true
	EnableAhas *string `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	// Enable CPU Burst.
	//
	// - true: enable
	//
	// - false: disable
	//
	// example:
	//
	// true
	EnableCpuBurst *string `json:"EnableCpuBurst,omitempty" xml:"EnableCpuBurst,omitempty"`
	// Indicates whether canary release rules are enabled. Canary release rules apply only to applications in Spring Cloud and Dubbo frameworks. Valid values:
	//
	// 	- **true**: The canary release rules are enabled.
	//
	// 	- **false**: The canary release rules are disabled.
	//
	// example:
	//
	// false
	EnableGreyTagRoute *bool `json:"EnableGreyTagRoute,omitempty" xml:"EnableGreyTagRoute,omitempty"`
	// Enable idle mode.
	//
	// - true: enable
	//
	// - false: disable
	//
	// example:
	//
	// false
	EnableIdle                  *bool `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	EnableNamespaceAgentVersion *bool `json:"EnableNamespaceAgentVersion,omitempty" xml:"EnableNamespaceAgentVersion,omitempty"`
	// Enable new ARMS feature.
	//
	// - true: enable
	//
	// - false: disable
	//
	// example:
	//
	// false
	EnableNewArms    *bool `json:"EnableNewArms,omitempty" xml:"EnableNewArms,omitempty"`
	EnablePrometheus *bool `json:"EnablePrometheus,omitempty" xml:"EnablePrometheus,omitempty"`
	// The environment variables. Variable description:
	//
	// 	- **name**: the name of the environment variable.
	//
	// 	- **value**: the value or reference of the environment variable.
	//
	// example:
	//
	// [{"name":"TEST_ENV_KEY","value":"TEST_ENV_VAR"}]
	Envs                  *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	GpuCount              *string `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	GpuType               *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	HeadlessPvtzDiscovery *string `json:"HeadlessPvtzDiscovery,omitempty" xml:"HeadlessPvtzDiscovery,omitempty"`
	Html                  *string `json:"Html,omitempty" xml:"Html,omitempty"`
	// The ID of the corresponding secret dictionary.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The URL of the image. This parameter is returned only if the **PackageType*	- parameter is set to **Image**.
	//
	// example:
	//
	// docker.io/library/nginx:1.14.2
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Initialize container configuration.
	InitContainersConfig []*DescribeApplicationConfigResponseBodyDataInitContainersConfig `json:"InitContainersConfig,omitempty" xml:"InitContainersConfig,omitempty" type:"Repeated"`
	IsStateful           *bool                                                            `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The arguments in the JAR package. The arguments are used to start the application container. The default startup command is `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`.
	//
	// example:
	//
	// start
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The option settings in the JAR package. The settings are used to start the application container. The default startup command is `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`.
	//
	// example:
	//
	// -Dtest=true
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The version of the Java development kit (JDK) on which the deployment package of the application depends. The following versions are supported:
	//
	// 	- **Open JDK 8**
	//
	// 	- **Open JDK 7**
	//
	// 	- **Dragonwell 11**
	//
	// 	- **Dragonwell 8**
	//
	// 	- **openjdk-8u191-jdk-alpine3.9**
	//
	// 	- **openjdk-7u201-jdk-alpine3.9**
	//
	// This parameter is not returned if the **PackageType*	- parameter is set to **Image**.
	//
	// example:
	//
	// Open JDK 8
	Jdk *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	// The logging configurations of Message Queue for Apache Kafka. The following parameters are involved:
	//
	// 	- **KafkaConfigs**: the configurations of Message Queue for Apache Kafka.
	//
	// 	- **createTime**: the time when the Message Queue for Apache Kafka instance was created.
	//
	// 	- **kafkaTopic**: the message topic that is used to classify messages.
	//
	// 	- **logDir**: the path in which logs are stored.
	//
	// 	- **logType**: the type of collected logs. Valid values:
	//
	//     	- **file_log**: the file log that is stored in the container. The path of the file logs in the container is returned.
	//
	//     	- **stdout**: the standard output log of the container. You can specify only one stdout value.
	//
	// 	- **kafkaEndpoint**: the endpoint of the Message Queue for Apache Kafka service.
	//
	// 	- **kafkaInstanceId**: the ID of the Message Queue for Apache Kafka instance.
	//
	// 	- **region**: the region where the Message Queue for Apache Kafka instance resides.
	KafkaConfigs *string `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	// The details of the availability check that was performed on the container. If the container fails this health check multiple times, the system disables and restarts the container. You can use one of the following methods to perform the health check:
	//
	// 	- Sample code of the **exec*	- method: `{"exec":{"command":["sh","-c","cat/home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}`
	//
	// 	- Sample code of the **httpGet*	- method: `{"httpGet":{"path":"/","port":18091,"scheme":"HTTP","isContainKeyWord":true,"keyWord":"SAE"},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// 	- Sample code of the **tcpSocket*	- method: `{"tcpSocket":{"port":18091},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// >  You can use only one method to perform the health check.
	//
	// The following parameters are involved:
	//
	// 	- **exec.command**: the health check command.
	//
	// 	- **httpGet.path**: the request path.
	//
	// 	- **httpGet.scheme**: the protocol that is used to perform the health check. Valid values: **HTTP*	- and **HTTPS**.
	//
	// 	- **httpGet.isContainKeyWord**: indicates whether the response contains keywords. Valid values: **true*	- and **false**. If this field is not returned, the advanced settings are not used.
	//
	// 	- **httpGet.keyWord**: the custom keyword. This parameter is available only if the **isContainKeyWord*	- field is returned.
	//
	// 	- **tcpSocket.port**: the port that is used to check the status of TCP connections.
	//
	// 	- **initialDelaySeconds**: the delay of the health check. Default value: 10. Unit: seconds.
	//
	// 	- **periodSeconds**: the interval at which health checks are performed. Default value: 30. Unit: seconds.
	//
	// 	- **timeoutSeconds**: the timeout period of the health check. Default value: 1. Unit: seconds. If you set this parameter to 0 or leave this parameter empty, the timeout period is automatically set to 1 second.
	//
	// example:
	//
	// {"exec":{"command":["curl http://localhost:8080"]},"initialDelaySeconds":20,"timeoutSeconds":3}
	Liveness              *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	LokiConfigs           *string `json:"LokiConfigs,omitempty" xml:"LokiConfigs,omitempty"`
	MaxSurgeInstanceRatio *int32  `json:"MaxSurgeInstanceRatio,omitempty" xml:"MaxSurgeInstanceRatio,omitempty"`
	MaxSurgeInstances     *int32  `json:"MaxSurgeInstances,omitempty" xml:"MaxSurgeInstances,omitempty"`
	// The size of memory required by each instance. Unit: MB. You cannot set this parameter to 0. The values of this parameter correspond to the values of the Cpu parameter:
	//
	// 	- This parameter is set to **1024*	- if the Cpu parameter is set to 500 or 1000.
	//
	// 	- This parameter is set to **2048*	- if the Cpu parameter is set to 500, 1000, or 1000.
	//
	// 	- This parameter is set to **4096*	- if the Cpu parameter is set to 1000, 2000, or 4000.
	//
	// 	- This parameter is set to **8192*	- if the Cpu parameter is set to 2000, 4000, or 8000.
	//
	// 	- This parameter is set to **12288*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **16384*	- if the Cpu parameter is set to 4000, 8000, or 16000.
	//
	// 	- This parameter is set to **24567*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **32768*	- if the Cpu parameter is set to 16000.
	//
	// 	- This parameter is set to **65536*	- if the Cpu parameter is set to 8000, 16000, or 32000.
	//
	// 	- This parameter is set to **131072*	- if the Cpu parameter is set to 32000.
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The Nacos registry. Valid values:
	//
	// 	- **0**: SAE built-in Nacos registry
	//
	// 	- **1**: self-managed Nacos registry
	//
	// 	- **2*	- : MSE Nacos registry
	//
	// example:
	//
	// "0"
	MicroRegistration *string `json:"MicroRegistration,omitempty" xml:"MicroRegistration,omitempty"`
	// The configuration of registration center. Takes effect only the type of registration center is MSE enterprise Nacos.
	//
	// example:
	//
	// {\\"instanceId\\":\\"mse-cn-1ls43******\\",\\"namespace\\":\\"62ee12fb-c279-4da4-be96-21**********\\"}
	MicroRegistrationConfig *string `json:"MicroRegistrationConfig,omitempty" xml:"MicroRegistrationConfig,omitempty"`
	// Configure microservices governance
	//
	// enable: Whether to enable microservices governance
	//
	// - true: Enable
	//
	// - false: Disable
	//
	// mseLosslessRule: Configure lossless online/offline deployment
	//
	// - delayTime: Delay duration (unit: seconds)
	//
	// - enable: Whether to enable lossless deployment. Set to "true" to enable; set to "false" to disable.
	//
	// - notice: Whether to enable notifications. Set to "true" to enable; set to "false" to disable.
	//
	// - warmupTime: Small-traffic warm-up duration (unit: seconds)
	//
	// example:
	//
	// {\\"Enable\\":true,\\"MseLosslessRule\\":{\\"enable\\":true,\\"notice\\":true,\\"delayTime\\":10,\\"warmupTime\\":120,\\"funcType\\":2,\\"aligned\\":false,\\"related\\":false,\\"lossLessDetail\\":false}}
	MicroserviceEngineConfig *string `json:"MicroserviceEngineConfig,omitempty" xml:"MicroserviceEngineConfig,omitempty"`
	// The percentage of the minimum number of available instances. Valid values:
	//
	// 	- **-1**: the default value. This value indicates that the minimum number of available instances is not measured by percentage. If you do not configure this parameter, the default value **-1*	- is used.
	//
	// 	- **0 to 100**: indicates that the minimum number of available instances is calculated by using the following formula: Current number of instances × (Value of MinReadyInstanceRatio × 100%). If the calculated result is not an integer, the result is rounded up to the nearest integer. For example, if the percentage is set to **50**% and five instances are available, the minimum number of available instances is 3.
	//
	// >  If the **MinReadyInstance*	- and **MinReadyInstanceRatio*	- parameters are returned and the value of the **MinReadyInstanceRatio*	- parameter is not **-1**, the value of the **MinReadyInstanceRatio*	- parameter takes effect. If the **MinReadyInstances*	- parameter is set to **5*	- and the **MinReadyInstanceRatio*	- parameter is set to **50**, the value of the **MinReadyInstanceRatio*	- parameter determines the minimum number of available instances.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of available instances. Valid values:
	//
	// 	- If you set the value to **0**, business interruptions occur when the application is updated.
	//
	// 	- If you set the value to **-1**, the minimum number of available instances is automatically set to a system-recommended value. The value is the nearest integer to which the calculated result of the following formula is rounded up: Current number of instances × 25%. For example, if five instances are available, the minimum number of available instances is calculated by using the following formula: 5 × 25% = 1.25. In this case, the minimum number of available instances is 2.
	//
	// >  Make sure that at least one instance is available during application deployment and rollback to prevent business interruptions.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// The details of the mounted NAS file system.
	MountDesc []*DescribeApplicationConfigResponseBodyDataMountDesc `json:"MountDesc,omitempty" xml:"MountDesc,omitempty" type:"Repeated"`
	// The mount target of the NAS file system in the VPC where the application is deployed. If you do not need to modify this configuration during the deployment, configure the **MountHost*	- parameter only in the first request. You do not need to include this parameter in subsequent requests. If you need to remove this configuration, leave the **MountHost*	- parameter empty in the request.
	//
	// example:
	//
	// example.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The ID of the microservice application.
	//
	// example:
	//
	// xxxxxxx@xxxxx
	MseApplicationId *string `json:"MseApplicationId,omitempty" xml:"MseApplicationId,omitempty"`
	// The application name of SAE service registered in MSE.
	//
	// example:
	//
	// cn-shenzhen-alb-demo-5c****
	MseApplicationName *string `json:"MseApplicationName,omitempty" xml:"MseApplicationName,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The configurations for mounting the NAS file system.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// The ID of the NAS file system.
	//
	// example:
	//
	// AKSN89**
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The SAE application edition.
	//
	// - lite: The lightweight edition.
	//
	// - std: The standard edition.
	//
	// - pro: The professional edition.
	//
	// example:
	//
	// pro
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	// The name of the RAM role used to authenticate the user identity.
	//
	// >  You need to create an OpenID Connect (OIDC) identity provider (IdP) and an identity provider (IdP) for role-based single sign-on (SSO) in advance. For more information, see [Creates an OpenID Connect (OIDC) identity provider (IdP)](https://help.aliyun.com/document_detail/2331022.html) and [Creates an identity provider (IdP) for role-based single sign-on (SSO)](https://help.aliyun.com/document_detail/2331016.html).
	//
	// example:
	//
	// sae-test
	OidcRoleName *string `json:"OidcRoleName,omitempty" xml:"OidcRoleName,omitempty"`
	// The AccessKey ID that is used to read data from and write data to Object Storage Service (OSS) buckets.
	//
	// example:
	//
	// xxxxxx
	OssAkId *string `json:"OssAkId,omitempty" xml:"OssAkId,omitempty"`
	// The AccessKey secret that is used to read data from and write data to OSS buckets.
	//
	// example:
	//
	// xxxxxx
	OssAkSecret *string `json:"OssAkSecret,omitempty" xml:"OssAkSecret,omitempty"`
	// The description of the mounted OSS bucket.
	OssMountDescs []*DescribeApplicationConfigResponseBodyDataOssMountDescs `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty" type:"Repeated"`
	// The type of the deployment package. Valid values:
	//
	// 	- If you deploy the application by using a Java Archive (JAR) package, you can set this parameter to **FatJar**, **War**, or **Image**.
	//
	// 	- If you deploy the application by using a PHP package, you can set this parameter to one of the following values:
	//
	//     	- **PhpZip**
	//
	//     	- **IMAGE_PHP_5_4**
	//
	//     	- **IMAGE_PHP_5_4_ALPINE**
	//
	//     	- **IMAGE_PHP_5_5**
	//
	//     	- **IMAGE_PHP_5_5_ALPINE**
	//
	//     	- **IMAGE_PHP_5_6**
	//
	//     	- **IMAGE_PHP_5_6_ALPINE**
	//
	//     	- **IMAGE_PHP_7_0**
	//
	//     	- **IMAGE_PHP_7_0_ALPINE**
	//
	//     	- **IMAGE_PHP_7_1**
	//
	//     	- **IMAGE_PHP_7_1_ALPINE**
	//
	//     	- **IMAGE_PHP_7_2**
	//
	//     	- **IMAGE_PHP_7_2_ALPINE**
	//
	//     	- **IMAGE_PHP_7_3**
	//
	//     	- **IMAGE_PHP_7_3_ALPINE**
	//
	// example:
	//
	// War
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The URL of the deployment package. This parameter is returned only if the **PackageType*	- parameter is set to **FatJar*	- or **War**.
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package. This parameter is returned only if the **PackageType*	- parameter is set to **FatJar*	- or **War**.
	//
	// example:
	//
	// 1.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The version of PHP supporting PHP deployment packages. Image is not supported.
	//
	// example:
	//
	// PHP-FPM 7.0
	Php *string `json:"Php,omitempty" xml:"Php,omitempty"`
	// The path on which the PHP configuration file for application monitoring is mounted. Make sure that the PHP server loads the configuration file.
	//
	// SAE automatically generates the corresponding configuration file. No manual operations are required.
	//
	// example:
	//
	// /usr/local/etc/php/conf.d/arms.ini
	PhpArmsConfigLocation *string `json:"PhpArmsConfigLocation,omitempty" xml:"PhpArmsConfigLocation,omitempty"`
	// The details of the PHP configuration file.
	//
	// example:
	//
	// k1=v1
	PhpConfig *string `json:"PhpConfig,omitempty" xml:"PhpConfig,omitempty"`
	// The path on which the PHP configuration file for application startup is mounted. Make sure that the PHP server uses this configuration file during the startup.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// The script that is run immediately after the container is started. Example: `{"exec":{"command":["cat","/etc/group"]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script that is run before the container is stopped. Example: `{"exec":{"command":["cat","/etc/group"]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The programming language that is used to create the application. Valid values:
	//
	// 	- **java**: Java
	//
	// 	- **php**: PHP
	//
	// 	- **other**: Other programming languages, such as Python, C++, Go, .NET, and Node.js.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// Enable K8s Service discovery and registration.
	//
	// example:
	//
	// {"serviceName":"bwm-poc-sc-gateway-cn-beijing-front","namespaceId":"cn-beijing:front","portAndProtocol":{"18012":"TCP"},"enable":true}
	PvtzDiscovery *string `json:"PvtzDiscovery,omitempty" xml:"PvtzDiscovery,omitempty"`
	// The Python environment. PYTHON 3.9.15 is supported.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// The configurations for installing custom module dependencies. By default, the dependencies defined by the requirements.txt file in the root directory are installed. If no software package is configured, you can specify dependencies based on your business requirements.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// The details of the health check that was performed on the container. If the container fails this health check multiple times, the system disables and restarts the container. Containers that fail health checks cannot receive traffic from Server Load Balancer (SLB) instances. You can use the **exec**, **httpGet**, or **tcpSocket*	- method to perform health checks. For more information, see the description of the **Liveness*	- parameter.
	//
	// >  You can use only one method to perform the health check.
	//
	// example:
	//
	// {"exec":{"command":["curl http://localhost:8080"]},"initialDelaySeconds":20,"timeoutSeconds":5}
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of application instances.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The type of the resource. Set the value to `application`.
	//
	// example:
	//
	// application
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Secret mount description.
	SecretMountDesc []*DescribeApplicationConfigResponseBodyDataSecretMountDesc `json:"SecretMountDesc,omitempty" xml:"SecretMountDesc,omitempty" type:"Repeated"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The canary tag configured for the application.
	ServiceTags map[string]*string `json:"ServiceTags,omitempty" xml:"ServiceTags,omitempty"`
	// The configuration of the Sidecar container.
	SidecarContainersConfig []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfig `json:"SidecarContainersConfig,omitempty" xml:"SidecarContainersConfig,omitempty" type:"Repeated"`
	// The logging configurations of Log Service.
	//
	// 	- To use Log Service resources that are automatically created by SAE, set this parameter to `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// 	- To use custom Log Service resources, set this parameter to `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// The following parameters are involved:
	//
	// 	- **projectName**: the name of the Log Service project.
	//
	// 	- **logDir**: the path in which logs are stored.
	//
	// 	- **logType**: the log type. **stdout**: the standard output log of the container. You can specify only one stdout value for this parameter. If you leave this parameter empty, file logs are collected.
	//
	// 	- **logstoreName**: the name of the Logstore in Log Service.
	//
	// 	- **logtailName**: the name of the Logtail configuration in Log Service. If you do not configure this parameter, a new Logtail configuration is created.
	//
	// If you do not need to modify the logging configurations when you deploy the application, configure the **SlsConfigs*	- parameter only in the first request. You do not need to include this parameter in subsequent requests. If you no longer need to use Log Service, leave the **SlsConfigs*	- parameter empty in the request.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs    *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	SlsLogEnvTags *string `json:"SlsLogEnvTags,omitempty" xml:"SlsLogEnvTags,omitempty"`
	// Enable startup probe.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"/bin/sh\\",\\"-c\\",\\"#!Note: If microservice config is enabled, the application will be automatically injected with the prestop configuration for lossless offline. If you delete this prestop configuration, lossless offline will not be effective.\\\\n echo stop > /tmp/prestop; /home/admin/.tools/curl http://127.0.0.1:54199/offline; sleep 30\\"]}}
	StartupProbe *string `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	// Configuration of K8s Service discovery and registration, and full-chain gray-release feature.
	//
	// example:
	//
	// {\\"enable\\":\\"false\\",\\"namespaceId\\":\\"cn-beijing:test\\",\\"portAndProtocol\\":{\\"2000:TCP\\":\\"18081\\"},\\"portProtocols\\":[{\\"port\\":2000,\\"protocol\\":\\"TCP\\",\\"targetPort\\":18081}],\\"pvtzDiscoveryName\\":\\"cn-beijing-1421801774382676\\",\\"serviceId\\":\\"3513\\",\\"serviceName\\":\\"demo-gray.test\\"}
	SwimlanePvtzDiscovery *string `json:"SwimlanePvtzDiscovery,omitempty" xml:"SwimlanePvtzDiscovery,omitempty"`
	// The details of the tags.
	Tags []*DescribeApplicationConfigResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period for a graceful shutdown. Default value: 30. Unit: seconds. Valid values: 1 to 300.
	//
	// example:
	//
	// 10
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The time zone. Default value: **Asia/Shanghai**.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The Tomcat configuration. If you want to delete the configuration, set this parameter to {} or leave this parameter empty. The following parameters are involved:
	//
	// 	- **port**: the port number. Valid values: 1024 to 65535. The root permissions are required to perform operations on ports whose number is smaller than 1024. Enter a value that ranges from 1025 to 65535 because the container has only the admin permissions. If you do not configure this parameter, the default port number 8080 is used.
	//
	// 	- **contextPath**: the path. Default value: /. This value indicates the root directory.
	//
	// 	- **maxThreads**: the maximum number of connections in the connection pool. Default value: 400.
	//
	// 	- **uriEncoding**: the URI encoding scheme in the Tomcat container. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. If you do not configure this parameter, the default value **ISO-8859-1*	- is used.
	//
	// 	- **useBodyEncoding**: indicates whether to use the encoding scheme that is specified by **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	// The deployment policy. If the minimum number of available instances is 1, the value of the **UpdateStrategy*	- parameter is an empty string (""). If the minimum number of available instances is greater than 1, the following strategies can be configured:
	//
	// 	- The application is deployed on an instance. The remaining instances are automatically classified into two release batches whose interval is set to 1. In this case, the parameter is set to `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}`.
	//
	// 	- The application is deployed on an instance. The remaining instances are manually classified into two release batches. In this case, the parameter is set to `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"manual"},"grayUpdate":{"gray":1}}`.
	//
	// 	- All instances are automatically classified into two release batches. The application is deployed on the instances of the two batches in parallel. In this case, the parameter is set to `{"type":"BatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":0}}`
	//
	// The following parameters are involved:
	//
	// 	- **type**: the type of the release policy. Valid values: **GrayBatchUpdate*	- and **BatchUpdate**.
	//
	// 	- **batchUpdate**: the phased release policy.
	//
	//     	- **batch**: the number of release batches.
	//
	//     	- **releaseType**: the processing method for the batches. Valid values: **auto*	- and **manual**.
	//
	//     	- **batchWaitTime**: the interval between release batches. Unit: seconds.
	//
	// 	- **grayUpdate**: the number of release batches in the phased release after a canary release. This parameter is returned only if the **type*	- parameter is set to **GrayBatchUpdate**.
	//
	// example:
	//
	// {"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}
	UpdateStrategy *string `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-2ze559r1z1bpwqxwp****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2ze0i263cnn311nvj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The option settings in the WAR package. The settings are used to start the application container. The default startup command is `java $JAVA_OPTS $CATALINA_OPTS -Options org.apache.catalina.startup.Bootstrap "$@" start`.
	//
	// example:
	//
	// custom-option
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The version of the Tomcat container on which the deployment package depends. Valid values:
	//
	// 	- **apache-tomcat-7.0.91**
	//
	// 	- **apache-tomcat-8.5.42**
	//
	// This parameter is not returned if the **PackageType*	- parameter is set to **Image**.
	//
	// example:
	//
	// apache-tomcat-7.0.91
	WebContainer *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyData) GetAcrAssumeRoleArn() *string {
	return s.AcrAssumeRoleArn
}

func (s *DescribeApplicationConfigResponseBodyData) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *DescribeApplicationConfigResponseBodyData) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *DescribeApplicationConfigResponseBodyData) GetAlbIngressReadinessGate() *string {
	return s.AlbIngressReadinessGate
}

func (s *DescribeApplicationConfigResponseBodyData) GetAppDescription() *string {
	return s.AppDescription
}

func (s *DescribeApplicationConfigResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationConfigResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeApplicationConfigResponseBodyData) GetAppSource() *string {
	return s.AppSource
}

func (s *DescribeApplicationConfigResponseBodyData) GetAssociateEip() *bool {
	return s.AssociateEip
}

func (s *DescribeApplicationConfigResponseBodyData) GetBaseAppId() *string {
	return s.BaseAppId
}

func (s *DescribeApplicationConfigResponseBodyData) GetBatchWaitTime() *int32 {
	return s.BatchWaitTime
}

func (s *DescribeApplicationConfigResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApplicationConfigResponseBodyData) GetCmsServiceId() *string {
	return s.CmsServiceId
}

func (s *DescribeApplicationConfigResponseBodyData) GetCommand() *string {
	return s.Command
}

func (s *DescribeApplicationConfigResponseBodyData) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *DescribeApplicationConfigResponseBodyData) GetConfigMapMountDesc() []*DescribeApplicationConfigResponseBodyDataConfigMapMountDesc {
	return s.ConfigMapMountDesc
}

func (s *DescribeApplicationConfigResponseBodyData) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeApplicationConfigResponseBodyData) GetCustomHostAlias() *string {
	return s.CustomHostAlias
}

func (s *DescribeApplicationConfigResponseBodyData) GetCustomImageNetworkType() *string {
	return s.CustomImageNetworkType
}

func (s *DescribeApplicationConfigResponseBodyData) GetDeploymentName() *string {
	return s.DeploymentName
}

func (s *DescribeApplicationConfigResponseBodyData) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DescribeApplicationConfigResponseBodyData) GetDotnet() *string {
	return s.Dotnet
}

func (s *DescribeApplicationConfigResponseBodyData) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *DescribeApplicationConfigResponseBodyData) GetEmptyDirDesc() []*DescribeApplicationConfigResponseBodyDataEmptyDirDesc {
	return s.EmptyDirDesc
}

func (s *DescribeApplicationConfigResponseBodyData) GetEnableAhas() *string {
	return s.EnableAhas
}

func (s *DescribeApplicationConfigResponseBodyData) GetEnableCpuBurst() *string {
	return s.EnableCpuBurst
}

func (s *DescribeApplicationConfigResponseBodyData) GetEnableGreyTagRoute() *bool {
	return s.EnableGreyTagRoute
}

func (s *DescribeApplicationConfigResponseBodyData) GetEnableIdle() *bool {
	return s.EnableIdle
}

func (s *DescribeApplicationConfigResponseBodyData) GetEnableNamespaceAgentVersion() *bool {
	return s.EnableNamespaceAgentVersion
}

func (s *DescribeApplicationConfigResponseBodyData) GetEnableNewArms() *bool {
	return s.EnableNewArms
}

func (s *DescribeApplicationConfigResponseBodyData) GetEnablePrometheus() *bool {
	return s.EnablePrometheus
}

func (s *DescribeApplicationConfigResponseBodyData) GetEnvs() *string {
	return s.Envs
}

func (s *DescribeApplicationConfigResponseBodyData) GetGpuCount() *string {
	return s.GpuCount
}

func (s *DescribeApplicationConfigResponseBodyData) GetGpuType() *string {
	return s.GpuType
}

func (s *DescribeApplicationConfigResponseBodyData) GetHeadlessPvtzDiscovery() *string {
	return s.HeadlessPvtzDiscovery
}

func (s *DescribeApplicationConfigResponseBodyData) GetHtml() *string {
	return s.Html
}

func (s *DescribeApplicationConfigResponseBodyData) GetImagePullSecrets() *string {
	return s.ImagePullSecrets
}

func (s *DescribeApplicationConfigResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeApplicationConfigResponseBodyData) GetInitContainersConfig() []*DescribeApplicationConfigResponseBodyDataInitContainersConfig {
	return s.InitContainersConfig
}

func (s *DescribeApplicationConfigResponseBodyData) GetIsStateful() *bool {
	return s.IsStateful
}

func (s *DescribeApplicationConfigResponseBodyData) GetJarStartArgs() *string {
	return s.JarStartArgs
}

func (s *DescribeApplicationConfigResponseBodyData) GetJarStartOptions() *string {
	return s.JarStartOptions
}

func (s *DescribeApplicationConfigResponseBodyData) GetJdk() *string {
	return s.Jdk
}

func (s *DescribeApplicationConfigResponseBodyData) GetKafkaConfigs() *string {
	return s.KafkaConfigs
}

func (s *DescribeApplicationConfigResponseBodyData) GetLiveness() *string {
	return s.Liveness
}

func (s *DescribeApplicationConfigResponseBodyData) GetLokiConfigs() *string {
	return s.LokiConfigs
}

func (s *DescribeApplicationConfigResponseBodyData) GetMaxSurgeInstanceRatio() *int32 {
	return s.MaxSurgeInstanceRatio
}

func (s *DescribeApplicationConfigResponseBodyData) GetMaxSurgeInstances() *int32 {
	return s.MaxSurgeInstances
}

func (s *DescribeApplicationConfigResponseBodyData) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeApplicationConfigResponseBodyData) GetMicroRegistration() *string {
	return s.MicroRegistration
}

func (s *DescribeApplicationConfigResponseBodyData) GetMicroRegistrationConfig() *string {
	return s.MicroRegistrationConfig
}

func (s *DescribeApplicationConfigResponseBodyData) GetMicroserviceEngineConfig() *string {
	return s.MicroserviceEngineConfig
}

func (s *DescribeApplicationConfigResponseBodyData) GetMinReadyInstanceRatio() *int32 {
	return s.MinReadyInstanceRatio
}

func (s *DescribeApplicationConfigResponseBodyData) GetMinReadyInstances() *int32 {
	return s.MinReadyInstances
}

func (s *DescribeApplicationConfigResponseBodyData) GetMountDesc() []*DescribeApplicationConfigResponseBodyDataMountDesc {
	return s.MountDesc
}

func (s *DescribeApplicationConfigResponseBodyData) GetMountHost() *string {
	return s.MountHost
}

func (s *DescribeApplicationConfigResponseBodyData) GetMseApplicationId() *string {
	return s.MseApplicationId
}

func (s *DescribeApplicationConfigResponseBodyData) GetMseApplicationName() *string {
	return s.MseApplicationName
}

func (s *DescribeApplicationConfigResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeApplicationConfigResponseBodyData) GetNasConfigs() *string {
	return s.NasConfigs
}

func (s *DescribeApplicationConfigResponseBodyData) GetNasId() *string {
	return s.NasId
}

func (s *DescribeApplicationConfigResponseBodyData) GetNewSaeVersion() *string {
	return s.NewSaeVersion
}

func (s *DescribeApplicationConfigResponseBodyData) GetOidcRoleName() *string {
	return s.OidcRoleName
}

func (s *DescribeApplicationConfigResponseBodyData) GetOssAkId() *string {
	return s.OssAkId
}

func (s *DescribeApplicationConfigResponseBodyData) GetOssAkSecret() *string {
	return s.OssAkSecret
}

func (s *DescribeApplicationConfigResponseBodyData) GetOssMountDescs() []*DescribeApplicationConfigResponseBodyDataOssMountDescs {
	return s.OssMountDescs
}

func (s *DescribeApplicationConfigResponseBodyData) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeApplicationConfigResponseBodyData) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *DescribeApplicationConfigResponseBodyData) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *DescribeApplicationConfigResponseBodyData) GetPhp() *string {
	return s.Php
}

func (s *DescribeApplicationConfigResponseBodyData) GetPhpArmsConfigLocation() *string {
	return s.PhpArmsConfigLocation
}

func (s *DescribeApplicationConfigResponseBodyData) GetPhpConfig() *string {
	return s.PhpConfig
}

func (s *DescribeApplicationConfigResponseBodyData) GetPhpConfigLocation() *string {
	return s.PhpConfigLocation
}

func (s *DescribeApplicationConfigResponseBodyData) GetPostStart() *string {
	return s.PostStart
}

func (s *DescribeApplicationConfigResponseBodyData) GetPreStop() *string {
	return s.PreStop
}

func (s *DescribeApplicationConfigResponseBodyData) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *DescribeApplicationConfigResponseBodyData) GetPvtzDiscovery() *string {
	return s.PvtzDiscovery
}

func (s *DescribeApplicationConfigResponseBodyData) GetPython() *string {
	return s.Python
}

func (s *DescribeApplicationConfigResponseBodyData) GetPythonModules() *string {
	return s.PythonModules
}

func (s *DescribeApplicationConfigResponseBodyData) GetReadiness() *string {
	return s.Readiness
}

func (s *DescribeApplicationConfigResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationConfigResponseBodyData) GetReplicas() *int32 {
	return s.Replicas
}

func (s *DescribeApplicationConfigResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeApplicationConfigResponseBodyData) GetSecretMountDesc() []*DescribeApplicationConfigResponseBodyDataSecretMountDesc {
	return s.SecretMountDesc
}

func (s *DescribeApplicationConfigResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeApplicationConfigResponseBodyData) GetServiceTags() map[string]*string {
	return s.ServiceTags
}

func (s *DescribeApplicationConfigResponseBodyData) GetSidecarContainersConfig() []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	return s.SidecarContainersConfig
}

func (s *DescribeApplicationConfigResponseBodyData) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *DescribeApplicationConfigResponseBodyData) GetSlsLogEnvTags() *string {
	return s.SlsLogEnvTags
}

func (s *DescribeApplicationConfigResponseBodyData) GetStartupProbe() *string {
	return s.StartupProbe
}

func (s *DescribeApplicationConfigResponseBodyData) GetSwimlanePvtzDiscovery() *string {
	return s.SwimlanePvtzDiscovery
}

func (s *DescribeApplicationConfigResponseBodyData) GetTags() []*DescribeApplicationConfigResponseBodyDataTags {
	return s.Tags
}

func (s *DescribeApplicationConfigResponseBodyData) GetTerminationGracePeriodSeconds() *int32 {
	return s.TerminationGracePeriodSeconds
}

func (s *DescribeApplicationConfigResponseBodyData) GetTimezone() *string {
	return s.Timezone
}

func (s *DescribeApplicationConfigResponseBodyData) GetTomcatConfig() *string {
	return s.TomcatConfig
}

func (s *DescribeApplicationConfigResponseBodyData) GetUpdateStrategy() *string {
	return s.UpdateStrategy
}

func (s *DescribeApplicationConfigResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeApplicationConfigResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeApplicationConfigResponseBodyData) GetWarStartOptions() *string {
	return s.WarStartOptions
}

func (s *DescribeApplicationConfigResponseBodyData) GetWebContainer() *string {
	return s.WebContainer
}

func (s *DescribeApplicationConfigResponseBodyData) SetAcrAssumeRoleArn(v string) *DescribeApplicationConfigResponseBodyData {
	s.AcrAssumeRoleArn = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAcrInstanceId(v string) *DescribeApplicationConfigResponseBodyData {
	s.AcrInstanceId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAgentVersion(v string) *DescribeApplicationConfigResponseBodyData {
	s.AgentVersion = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAlbIngressReadinessGate(v string) *DescribeApplicationConfigResponseBodyData {
	s.AlbIngressReadinessGate = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAppDescription(v string) *DescribeApplicationConfigResponseBodyData {
	s.AppDescription = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAppId(v string) *DescribeApplicationConfigResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAppName(v string) *DescribeApplicationConfigResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAppSource(v string) *DescribeApplicationConfigResponseBodyData {
	s.AppSource = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetAssociateEip(v bool) *DescribeApplicationConfigResponseBodyData {
	s.AssociateEip = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetBaseAppId(v string) *DescribeApplicationConfigResponseBodyData {
	s.BaseAppId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetBatchWaitTime(v int32) *DescribeApplicationConfigResponseBodyData {
	s.BatchWaitTime = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetClusterId(v string) *DescribeApplicationConfigResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetCmsServiceId(v string) *DescribeApplicationConfigResponseBodyData {
	s.CmsServiceId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetCommand(v string) *DescribeApplicationConfigResponseBodyData {
	s.Command = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetCommandArgs(v string) *DescribeApplicationConfigResponseBodyData {
	s.CommandArgs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetConfigMapMountDesc(v []*DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) *DescribeApplicationConfigResponseBodyData {
	s.ConfigMapMountDesc = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetCpu(v int32) *DescribeApplicationConfigResponseBodyData {
	s.Cpu = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetCustomHostAlias(v string) *DescribeApplicationConfigResponseBodyData {
	s.CustomHostAlias = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetCustomImageNetworkType(v string) *DescribeApplicationConfigResponseBodyData {
	s.CustomImageNetworkType = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetDeploymentName(v string) *DescribeApplicationConfigResponseBodyData {
	s.DeploymentName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetDiskSize(v int32) *DescribeApplicationConfigResponseBodyData {
	s.DiskSize = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetDotnet(v string) *DescribeApplicationConfigResponseBodyData {
	s.Dotnet = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEdasContainerVersion(v string) *DescribeApplicationConfigResponseBodyData {
	s.EdasContainerVersion = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEmptyDirDesc(v []*DescribeApplicationConfigResponseBodyDataEmptyDirDesc) *DescribeApplicationConfigResponseBodyData {
	s.EmptyDirDesc = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEnableAhas(v string) *DescribeApplicationConfigResponseBodyData {
	s.EnableAhas = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEnableCpuBurst(v string) *DescribeApplicationConfigResponseBodyData {
	s.EnableCpuBurst = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEnableGreyTagRoute(v bool) *DescribeApplicationConfigResponseBodyData {
	s.EnableGreyTagRoute = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEnableIdle(v bool) *DescribeApplicationConfigResponseBodyData {
	s.EnableIdle = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEnableNamespaceAgentVersion(v bool) *DescribeApplicationConfigResponseBodyData {
	s.EnableNamespaceAgentVersion = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEnableNewArms(v bool) *DescribeApplicationConfigResponseBodyData {
	s.EnableNewArms = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEnablePrometheus(v bool) *DescribeApplicationConfigResponseBodyData {
	s.EnablePrometheus = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetEnvs(v string) *DescribeApplicationConfigResponseBodyData {
	s.Envs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetGpuCount(v string) *DescribeApplicationConfigResponseBodyData {
	s.GpuCount = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetGpuType(v string) *DescribeApplicationConfigResponseBodyData {
	s.GpuType = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetHeadlessPvtzDiscovery(v string) *DescribeApplicationConfigResponseBodyData {
	s.HeadlessPvtzDiscovery = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetHtml(v string) *DescribeApplicationConfigResponseBodyData {
	s.Html = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetImagePullSecrets(v string) *DescribeApplicationConfigResponseBodyData {
	s.ImagePullSecrets = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetImageUrl(v string) *DescribeApplicationConfigResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetInitContainersConfig(v []*DescribeApplicationConfigResponseBodyDataInitContainersConfig) *DescribeApplicationConfigResponseBodyData {
	s.InitContainersConfig = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetIsStateful(v bool) *DescribeApplicationConfigResponseBodyData {
	s.IsStateful = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetJarStartArgs(v string) *DescribeApplicationConfigResponseBodyData {
	s.JarStartArgs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetJarStartOptions(v string) *DescribeApplicationConfigResponseBodyData {
	s.JarStartOptions = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetJdk(v string) *DescribeApplicationConfigResponseBodyData {
	s.Jdk = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetKafkaConfigs(v string) *DescribeApplicationConfigResponseBodyData {
	s.KafkaConfigs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetLiveness(v string) *DescribeApplicationConfigResponseBodyData {
	s.Liveness = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetLokiConfigs(v string) *DescribeApplicationConfigResponseBodyData {
	s.LokiConfigs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMaxSurgeInstanceRatio(v int32) *DescribeApplicationConfigResponseBodyData {
	s.MaxSurgeInstanceRatio = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMaxSurgeInstances(v int32) *DescribeApplicationConfigResponseBodyData {
	s.MaxSurgeInstances = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMemory(v int32) *DescribeApplicationConfigResponseBodyData {
	s.Memory = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMicroRegistration(v string) *DescribeApplicationConfigResponseBodyData {
	s.MicroRegistration = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMicroRegistrationConfig(v string) *DescribeApplicationConfigResponseBodyData {
	s.MicroRegistrationConfig = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMicroserviceEngineConfig(v string) *DescribeApplicationConfigResponseBodyData {
	s.MicroserviceEngineConfig = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMinReadyInstanceRatio(v int32) *DescribeApplicationConfigResponseBodyData {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMinReadyInstances(v int32) *DescribeApplicationConfigResponseBodyData {
	s.MinReadyInstances = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMountDesc(v []*DescribeApplicationConfigResponseBodyDataMountDesc) *DescribeApplicationConfigResponseBodyData {
	s.MountDesc = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMountHost(v string) *DescribeApplicationConfigResponseBodyData {
	s.MountHost = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMseApplicationId(v string) *DescribeApplicationConfigResponseBodyData {
	s.MseApplicationId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetMseApplicationName(v string) *DescribeApplicationConfigResponseBodyData {
	s.MseApplicationName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetNamespaceId(v string) *DescribeApplicationConfigResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetNasConfigs(v string) *DescribeApplicationConfigResponseBodyData {
	s.NasConfigs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetNasId(v string) *DescribeApplicationConfigResponseBodyData {
	s.NasId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetNewSaeVersion(v string) *DescribeApplicationConfigResponseBodyData {
	s.NewSaeVersion = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetOidcRoleName(v string) *DescribeApplicationConfigResponseBodyData {
	s.OidcRoleName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetOssAkId(v string) *DescribeApplicationConfigResponseBodyData {
	s.OssAkId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetOssAkSecret(v string) *DescribeApplicationConfigResponseBodyData {
	s.OssAkSecret = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetOssMountDescs(v []*DescribeApplicationConfigResponseBodyDataOssMountDescs) *DescribeApplicationConfigResponseBodyData {
	s.OssMountDescs = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPackageType(v string) *DescribeApplicationConfigResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPackageUrl(v string) *DescribeApplicationConfigResponseBodyData {
	s.PackageUrl = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPackageVersion(v string) *DescribeApplicationConfigResponseBodyData {
	s.PackageVersion = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPhp(v string) *DescribeApplicationConfigResponseBodyData {
	s.Php = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPhpArmsConfigLocation(v string) *DescribeApplicationConfigResponseBodyData {
	s.PhpArmsConfigLocation = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPhpConfig(v string) *DescribeApplicationConfigResponseBodyData {
	s.PhpConfig = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPhpConfigLocation(v string) *DescribeApplicationConfigResponseBodyData {
	s.PhpConfigLocation = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPostStart(v string) *DescribeApplicationConfigResponseBodyData {
	s.PostStart = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPreStop(v string) *DescribeApplicationConfigResponseBodyData {
	s.PreStop = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetProgrammingLanguage(v string) *DescribeApplicationConfigResponseBodyData {
	s.ProgrammingLanguage = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPvtzDiscovery(v string) *DescribeApplicationConfigResponseBodyData {
	s.PvtzDiscovery = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPython(v string) *DescribeApplicationConfigResponseBodyData {
	s.Python = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetPythonModules(v string) *DescribeApplicationConfigResponseBodyData {
	s.PythonModules = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetReadiness(v string) *DescribeApplicationConfigResponseBodyData {
	s.Readiness = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetRegionId(v string) *DescribeApplicationConfigResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetReplicas(v int32) *DescribeApplicationConfigResponseBodyData {
	s.Replicas = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetResourceType(v string) *DescribeApplicationConfigResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetSecretMountDesc(v []*DescribeApplicationConfigResponseBodyDataSecretMountDesc) *DescribeApplicationConfigResponseBodyData {
	s.SecretMountDesc = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetSecurityGroupId(v string) *DescribeApplicationConfigResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetServiceTags(v map[string]*string) *DescribeApplicationConfigResponseBodyData {
	s.ServiceTags = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetSidecarContainersConfig(v []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) *DescribeApplicationConfigResponseBodyData {
	s.SidecarContainersConfig = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetSlsConfigs(v string) *DescribeApplicationConfigResponseBodyData {
	s.SlsConfigs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetSlsLogEnvTags(v string) *DescribeApplicationConfigResponseBodyData {
	s.SlsLogEnvTags = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetStartupProbe(v string) *DescribeApplicationConfigResponseBodyData {
	s.StartupProbe = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetSwimlanePvtzDiscovery(v string) *DescribeApplicationConfigResponseBodyData {
	s.SwimlanePvtzDiscovery = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetTags(v []*DescribeApplicationConfigResponseBodyDataTags) *DescribeApplicationConfigResponseBodyData {
	s.Tags = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetTerminationGracePeriodSeconds(v int32) *DescribeApplicationConfigResponseBodyData {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetTimezone(v string) *DescribeApplicationConfigResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetTomcatConfig(v string) *DescribeApplicationConfigResponseBodyData {
	s.TomcatConfig = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetUpdateStrategy(v string) *DescribeApplicationConfigResponseBodyData {
	s.UpdateStrategy = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetVSwitchId(v string) *DescribeApplicationConfigResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetVpcId(v string) *DescribeApplicationConfigResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetWarStartOptions(v string) *DescribeApplicationConfigResponseBodyData {
	s.WarStartOptions = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) SetWebContainer(v string) *DescribeApplicationConfigResponseBodyData {
	s.WebContainer = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyData) Validate() error {
	if s.ConfigMapMountDesc != nil {
		for _, item := range s.ConfigMapMountDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EmptyDirDesc != nil {
		for _, item := range s.EmptyDirDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InitContainersConfig != nil {
		for _, item := range s.InitContainersConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MountDesc != nil {
		for _, item := range s.MountDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OssMountDescs != nil {
		for _, item := range s.OssMountDescs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecretMountDesc != nil {
		for _, item := range s.SecretMountDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SidecarContainersConfig != nil {
		for _, item := range s.SidecarContainersConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationConfigResponseBodyDataConfigMapMountDesc struct {
	// The ID of the ConfigMap.
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	// The name of the ConfigMap.
	//
	// example:
	//
	// test
	ConfigMapName *string `json:"ConfigMapName,omitempty" xml:"ConfigMapName,omitempty"`
	// The key-value pair that is stored in the ConfigMap.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The path on which the ConfigMap is mounted.
	//
	// example:
	//
	// /tmp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) GetConfigMapId() *int64 {
	return s.ConfigMapId
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) GetConfigMapName() *string {
	return s.ConfigMapName
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) GetKey() *string {
	return s.Key
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) SetConfigMapId(v int64) *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc {
	s.ConfigMapId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) SetConfigMapName(v string) *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc {
	s.ConfigMapName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) SetKey(v string) *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc {
	s.Key = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataConfigMapMountDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataEmptyDirDesc struct {
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataEmptyDirDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataEmptyDirDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataEmptyDirDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeApplicationConfigResponseBodyDataEmptyDirDesc) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationConfigResponseBodyDataEmptyDirDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataEmptyDirDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataEmptyDirDesc) SetName(v string) *DescribeApplicationConfigResponseBodyDataEmptyDirDesc {
	s.Name = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataEmptyDirDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataInitContainersConfig struct {
	// The command that is used to start the image. The command must be an existing executable object in the container. Sample statements:
	//
	//     command:
	//
	//           - echo
	//
	//           - abc
	//
	//           - >
	//
	//           - file0
	//
	// In this example, the Command parameter is set to `Command="echo", CommandArgs=["abc", ">", "file0"]`.
	//
	// example:
	//
	// /bin/sh
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The parameters of the image startup command. The CommandArgs parameter specifies the parameters that are required for the **Command*	- parameter. You can specify the name in one of the following formats:
	//
	// `["a","b"]`
	//
	// In the preceding example, the CommandArgs parameter is set to `CommandArgs=["abc", ">", "file0"]`. The data type of `["abc", ">", "file0"]` must be an array of strings in the JSON format. This parameter is optional.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The information of ConfigMap.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	EmptyDirDesc       []*DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc       `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// The environment variables. You can configure custom environment variables or reference a ConfigMap. If you want to reference a ConfigMap, you must first create a ConfigMap. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Take note of the following rules:
	//
	// 	- Customize
	//
	//     	- **name**: the name of the environment variable.
	//
	//     	- **value**: the value of the environment variable.
	//
	// 	- Reference ConfigMap
	//
	//     	- **name**: the name of the environment variable. You can reference one or all keys. If you want to reference all keys, specify `sae-sys-configmap-all-<ConfigMap name>`. Example: `sae-sys-configmap-all-test1`.
	//
	//     	- **valueFrom**: the reference of the environment variable. Set the value to `configMapRef`.
	//
	//     	- **configMapId**: the ConfigMap ID.
	//
	//     	- **key**: the key. If you want to reference all keys, do not configure this parameter.
	//
	// example:
	//
	// [{"name":"TEST_ENV_KEY","value":"TEST_ENV_VAR"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The image URL of the initialized container.
	//
	// example:
	//
	// registry.cn-shenzhen.aliyuncs.com/sae-serverless-demo/sae-demo:microservice-java-provider-v1.0
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The name of the initialized container.
	//
	// example:
	//
	// init-container
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataInitContainersConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataInitContainersConfig) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) GetCommand() *string {
	return s.Command
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) GetConfigMapMountDesc() []*DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc {
	return s.ConfigMapMountDesc
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) GetEmptyDirDesc() []*DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc {
	return s.EmptyDirDesc
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) GetEnvs() *string {
	return s.Envs
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) SetCommand(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfig {
	s.Command = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) SetCommandArgs(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfig {
	s.CommandArgs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) SetConfigMapMountDesc(v []*DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) *DescribeApplicationConfigResponseBodyDataInitContainersConfig {
	s.ConfigMapMountDesc = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) SetEmptyDirDesc(v []*DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc) *DescribeApplicationConfigResponseBodyDataInitContainersConfig {
	s.EmptyDirDesc = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) SetEnvs(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfig {
	s.Envs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) SetImageUrl(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfig {
	s.ImageUrl = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) SetName(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfig {
	s.Name = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) Validate() error {
	if s.ConfigMapMountDesc != nil {
		for _, item := range s.ConfigMapMountDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EmptyDirDesc != nil {
		for _, item := range s.EmptyDirDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc struct {
	// ConfigMap ID。
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	// The name of the ConfigMap.
	//
	// example:
	//
	// test
	ConfigMapName *string `json:"ConfigMapName,omitempty" xml:"ConfigMapName,omitempty"`
	// The key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path.
	//
	// example:
	//
	// /tmp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) GetConfigMapId() *int64 {
	return s.ConfigMapId
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) GetConfigMapName() *string {
	return s.ConfigMapName
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) GetKey() *string {
	return s.Key
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) SetConfigMapId(v int64) *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc {
	s.ConfigMapId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) SetConfigMapName(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc {
	s.ConfigMapName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) SetKey(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc {
	s.Key = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc struct {
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc) SetName(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc {
	s.Name = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataMountDesc struct {
	// The path on which the NAS file system is mounted.
	//
	// example:
	//
	// /tmp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The directory in the NAS file system.
	//
	// example:
	//
	// /
	NasPath *string `json:"NasPath,omitempty" xml:"NasPath,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataMountDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataMountDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataMountDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeApplicationConfigResponseBodyDataMountDesc) GetNasPath() *string {
	return s.NasPath
}

func (s *DescribeApplicationConfigResponseBodyDataMountDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataMountDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataMountDesc) SetNasPath(v string) *DescribeApplicationConfigResponseBodyDataMountDesc {
	s.NasPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataMountDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataOssMountDescs struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// oss-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The directory or object in OSS. If the specified directory or object does not exist, an error is returned.
	//
	// example:
	//
	// data/user.data
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	// The path of the container in SAE. The parameter value that you specified overwrites the original value. If the specified path does not exist, SAE automatically creates the path.
	//
	// example:
	//
	// /usr/data/user.data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// Indicates whether the application can use the container path to read data from or write data to resources in the directory of the OSS bucket. Valid values:
	//
	// 	- **true**: The application has the read-only permissions.
	//
	// 	- **false**: The application has the read and write permissions.
	//
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataOssMountDescs) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataOssMountDescs) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) GetBucketPath() *string {
	return s.BucketPath
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) SetBucketName(v string) *DescribeApplicationConfigResponseBodyDataOssMountDescs {
	s.BucketName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) SetBucketPath(v string) *DescribeApplicationConfigResponseBodyDataOssMountDescs {
	s.BucketPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataOssMountDescs {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) SetReadOnly(v bool) *DescribeApplicationConfigResponseBodyDataOssMountDescs {
	s.ReadOnly = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataOssMountDescs) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataSecretMountDesc struct {
	// The key to Base64 encode values.
	//
	// example:
	//
	// task-center
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path.
	//
	// example:
	//
	// /opt/www/runtime/logs
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The secret ID of the instance.
	//
	// example:
	//
	// 520
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// dummy-name-opaque-894
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataSecretMountDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataSecretMountDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataSecretMountDesc) GetKey() *string {
	return s.Key
}

func (s *DescribeApplicationConfigResponseBodyDataSecretMountDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeApplicationConfigResponseBodyDataSecretMountDesc) GetSecretId() *int64 {
	return s.SecretId
}

func (s *DescribeApplicationConfigResponseBodyDataSecretMountDesc) GetSecretName() *string {
	return s.SecretName
}

func (s *DescribeApplicationConfigResponseBodyDataSecretMountDesc) SetKey(v string) *DescribeApplicationConfigResponseBodyDataSecretMountDesc {
	s.Key = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSecretMountDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataSecretMountDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSecretMountDesc) SetSecretId(v int64) *DescribeApplicationConfigResponseBodyDataSecretMountDesc {
	s.SecretId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSecretMountDesc) SetSecretName(v string) *DescribeApplicationConfigResponseBodyDataSecretMountDesc {
	s.SecretName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSecretMountDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataSidecarContainersConfig struct {
	// The ID of Container Registry Enterprise Edition instance. This parameter is required when the **ImageUrl*	- parameter is set to the URL of an image in an ACR Enterprise Edition instance.
	//
	// example:
	//
	// cri-fhzlneorxala66ip
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The command that is used to start the image. The command must be an existing executable object in the container. Sample statements:
	//
	//     command:
	//
	//           - echo
	//
	//           - abc
	//
	//           - >
	//
	//           - file0
	//
	// In this example, the Command parameter is set to `Command="echo", CommandArgs=["abc", ">", "file0"]`.
	//
	// example:
	//
	// /bin/sh
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The parameters of the image startup command. The CommandArgs parameter specifies the parameters that are required for the **Command*	- parameter. You can specify the name in one of the following formats:
	//
	// `["a","b"]`
	//
	// In the preceding example, the CommandArgs parameter is set to `CommandArgs=["abc", ">", "file0"]`. The data type of `["abc", ">", "file0"]` must be an array of strings in the JSON format. This parameter is optional.
	//
	// example:
	//
	// [\\"-c\\",\\"echo \\\\\\"test\\\\\\" > /home/nas/test.log && sleep 10000000s\\"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The description of the **ConfigMap*	- instance mounted to the application. Use configurations created on the Configuration Items page to configure containers. The following table describes the parameters that are used in the preceding statements.
	//
	// 	- **congfigMapId**: the ID of the ConfigMap instance. You can call the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the ID.
	//
	// 	- **key**: the key.
	//
	// > You can use the `sae-sys-configmap-all` key to mount all keys.
	//
	// 	- **mountPath**: the mount path in the container.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// Set the CPU resource limit of the primary container that can be used by Sidecar container.
	//
	// example:
	//
	// 500
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Shared temporary storage mounted to the primary container and the Sidecar container.
	EmptyDirDesc []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// The environment variables. You can configure custom environment variables or reference a ConfigMap. If you want to reference a ConfigMap, you must first create a ConfigMap. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Take note of the following rules:
	//
	// 	- Customize
	//
	//     	- **name**: the name of the environment variable.
	//
	//     	- **value**: the value of the environment variable.
	//
	// 	- Reference ConfigMap
	//
	//     	- **name**: the name of the environment variable. You can reference one or all keys. If you want to reference all keys, specify `sae-sys-configmap-all-<ConfigMap name>`. Example: `sae-sys-configmap-all-test1`.
	//
	//     	- **valueFrom**: the reference of the environment variable. Set the value to `configMapRef`.
	//
	//     	- **configMapId**: the ConfigMap ID.
	//
	//     	- **key**: the key. If you want to reference all keys, do not configure this parameter.
	//
	// example:
	//
	// [{\\"name\\":\\"k1\\",\\"value\\":\\"c8e3a815-e5d3-4adf-abb3-98b106a607c4\\"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The URL of the image.
	//
	// example:
	//
	// registry.cn-beijing.aliyuncs.com/sae-dev-test/nginx:stable
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Set the memory limit of the primary container that can be used by Sidecar container.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The container name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetCommand() *string {
	return s.Command
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetConfigMapMountDesc() []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc {
	return s.ConfigMapMountDesc
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetEmptyDirDesc() []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc {
	return s.EmptyDirDesc
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetEnvs() *string {
	return s.Envs
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetAcrInstanceId(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.AcrInstanceId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetCommand(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.Command = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetCommandArgs(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.CommandArgs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetConfigMapMountDesc(v []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.ConfigMapMountDesc = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetCpu(v int32) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.Cpu = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetEmptyDirDesc(v []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.EmptyDirDesc = v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetEnvs(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.Envs = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetImageUrl(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.ImageUrl = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetMemory(v int32) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.Memory = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetName(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.Name = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) Validate() error {
	if s.ConfigMapMountDesc != nil {
		for _, item := range s.ConfigMapMountDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EmptyDirDesc != nil {
		for _, item := range s.EmptyDirDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc struct {
	// The ConfigMap ID.
	//
	// example:
	//
	// 7361
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	// The ConfigMap name.
	//
	// example:
	//
	// ConfigMap-test
	ConfigMapName *string `json:"ConfigMapName,omitempty" xml:"ConfigMapName,omitempty"`
	// The ConfigMap key
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path.
	//
	// example:
	//
	// /mnt/test
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) GetConfigMapId() *int64 {
	return s.ConfigMapId
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) GetConfigMapName() *string {
	return s.ConfigMapName
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) GetKey() *string {
	return s.Key
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) SetConfigMapId(v int64) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc {
	s.ConfigMapId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) SetConfigMapName(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc {
	s.ConfigMapName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) SetKey(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc {
	s.Key = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc struct {
	// Mount path of the data volume within the container.
	//
	// example:
	//
	// /mnt/cache
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The name of the shared temporary storage.
	//
	// example:
	//
	// sidecar-container
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc) SetName(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc {
	s.Name = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataTags struct {
	// The key of the tag.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *DescribeApplicationConfigResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *DescribeApplicationConfigResponseBodyDataTags) SetKey(v string) *DescribeApplicationConfigResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataTags) SetValue(v string) *DescribeApplicationConfigResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}

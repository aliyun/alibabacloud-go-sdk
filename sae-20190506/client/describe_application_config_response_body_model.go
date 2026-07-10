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
	// The API status or POP error code. Valid values:
	//
	// - **2xx**: success.
	//
	// - **3xx**: redirection.
	//
	// - **4xx**: request error.
	//
	// - **5xx**: server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The application information.
	Data *DescribeApplicationConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- field is not returned.
	//
	// - If the request fails, the **ErrorCode*	- field is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information about the call result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application configuration information is retrieved. Valid values:
	//
	// - **true**: Retrieved.
	//
	// - **false**: Failed to retrieve.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID, which is used to query the details of a call.
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
	// The ARN of the RAM role required for pulling images across accounts. For more information, see [Pull Alibaba Cloud images across accounts](https://help.aliyun.com/document_detail/190675.html) and [Grant cross-account permissions by using RAM roles](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The ACR Enterprise instance ID.
	//
	// example:
	//
	// cri-xxxxxx
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The agent version.
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The ALB gateway ReadinessGate configuration.
	AlbIngressReadinessGate *string `json:"AlbIngressReadinessGate,omitempty" xml:"AlbIngressReadinessGate,omitempty"`
	// The application description.
	//
	// example:
	//
	// Sample application
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The SAE application type.
	//
	// - micro_service
	//
	// - web
	//
	// - job
	//
	// example:
	//
	// micro_service
	AppSource *string `json:"AppSource,omitempty" xml:"AppSource,omitempty"`
	// Specifies whether to associate an EIP. Valid values:
	//
	// - **true**: Associated.
	//
	// - **false**: Not associated.
	//
	// example:
	//
	// true
	AssociateEip *bool `json:"AssociateEip,omitempty" xml:"AssociateEip,omitempty"`
	// The ID of the baseline application.
	//
	// example:
	//
	// 8c573618-8d72-4407-baf4-f7b64b******
	BaseAppId *string `json:"BaseAppId,omitempty" xml:"BaseAppId,omitempty"`
	// The wait time between batches during a phased release, in seconds.
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
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The CloudMonitor service ID.
	CmsServiceId *string `json:"CmsServiceId,omitempty" xml:"CmsServiceId,omitempty"`
	// The image startup command. This command must be an executable object that exists in the container. Example:
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
	// Based on the preceding example, `Command="echo", CommandArgs=["abc", ">", "file0"]`.
	//
	// example:
	//
	// echo
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments for the image startup command. These are the arguments required by the startup command **Command**. Format:
	//
	// `["a","b"]`
	//
	// In the example for the **Command*	- parameter, `CommandArgs=["abc", ">", "file0"]`, where `["abc", ">", "file0"]` must be converted to the String type and the internal format is a JSON array. If this parameter is not required, leave it empty.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The ConfigMap information.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The CPU required by each instance, in millicores. This value cannot be 0. Only the following defined specifications are supported:
	//
	// - **500**
	//
	// - **1000**
	//
	// - **2000**
	//
	// - **4000**
	//
	// - **8000**
	//
	// - **16000**
	//
	// - **32000**
	//
	// example:
	//
	// 1000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The custom host mapping in the container. Valid values:
	//
	// - **hostName**: The domain name or hostname.
	//
	// - **ip**: The IP address.
	//
	// example:
	//
	// [{"hostName":"test.host.name","ip":"0.0.0.0"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The type of the custom image. If the image is not a custom image, set this parameter to an empty string. Valid values:
	//
	// - internet: public image
	//
	// - intranet: internal image
	//
	// example:
	//
	// internet
	CustomImageNetworkType *string `json:"CustomImageNetworkType,omitempty" xml:"CustomImageNetworkType,omitempty"`
	// The instance name of the application in the ASI cluster.
	//
	// example:
	//
	// demo-b0cdce9b-3d1f-4168-b206-b59f348f1a8a
	DeploymentName *string `json:"DeploymentName,omitempty" xml:"DeploymentName,omitempty"`
	// The disk storage size, in GB.
	//
	// example:
	//
	// 20
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The .NET framework version:
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
	// The version of the application runtime environment in the HSF framework, such as the Ali-Tomcat container.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// The shared temporary storage.
	EmptyDirDesc []*DescribeApplicationConfigResponseBodyDataEmptyDirDesc `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// Specifies whether to connect to Application High Availability Service (AHAS). Valid values:
	//
	// - **true**: Connected to AHAS.
	//
	// - **false**: Not connected to AHAS.
	//
	// example:
	//
	// true
	EnableAhas *string `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	// Specifies whether to enable the CPU Burst feature. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// example:
	//
	// true
	EnableCpuBurst *string `json:"EnableCpuBurst,omitempty" xml:"EnableCpuBurst,omitempty"`
	// Specifies whether to enable the traffic canary release rule. This rule applies only to applications that use the Spring Cloud and Dubbo frameworks. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableGreyTagRoute *bool `json:"EnableGreyTagRoute,omitempty" xml:"EnableGreyTagRoute,omitempty"`
	// Specifies whether to enable idle mode. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// false
	EnableIdle *bool `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	// Indicates whether the namespace agent version configuration is reused.
	EnableNamespaceAgentVersion *bool `json:"EnableNamespaceAgentVersion,omitempty" xml:"EnableNamespaceAgentVersion,omitempty"`
	// Specifies whether to enable the new ARMS feature. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// example:
	//
	// false
	EnableNewArms *bool `json:"EnableNewArms,omitempty" xml:"EnableNewArms,omitempty"`
	// Indicates whether Prometheus custom metric collection is enabled.
	EnablePrometheus *bool `json:"EnablePrometheus,omitempty" xml:"EnablePrometheus,omitempty"`
	// The container environment variable parameters. Custom values or references to configuration items are supported. To reference a configuration item, create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//     - **name**: The environment variable name.
	//
	//     - **value**: The environment variable value.
	//
	// - Reference to a configuration item
	//
	//     - **name**: The environment variable name. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<ConfigMap name>`, such as `sae-sys-configmap-all-test1`.
	//
	//     - **valueFrom**: The environment variable reference. Set the value to `configMapRef`.
	//
	//     - **configMapId**: The ConfigMap ID.
	//
	//     - **key**: The key. If all keys are referenced, do not set this field.
	//
	// example:
	//
	// [{"name":"TEST_ENV_KEY","value":"TEST_ENV_VAR"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The number of GPUs.
	GpuCount *string `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The GPU type.
	GpuType               *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	HeadlessPvtzDiscovery *string `json:"HeadlessPvtzDiscovery,omitempty" xml:"HeadlessPvtzDiscovery,omitempty"`
	Html                  *string `json:"Html,omitempty" xml:"Html,omitempty"`
	IdleHour              *string `json:"IdleHour,omitempty" xml:"IdleHour,omitempty"`
	// The corresponding secret ID.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The image URL. This parameter is required when **Package Type*	- is set to **Image**.
	//
	// example:
	//
	// docker.io/library/nginx:1.14.2
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The init container configuration.
	InitContainersConfig []*DescribeApplicationConfigResponseBodyDataInitContainersConfig `json:"InitContainersConfig,omitempty" xml:"InitContainersConfig,omitempty" type:"Repeated"`
	// Indicates whether the application is stateful.
	IsStateful *bool `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The arguments for starting the JAR package application. The default startup command for the application: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// start
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The options for starting the JAR package application. The default startup command for the application: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// -Dtest=true
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The JDK version on which the deployment package depends. Valid values:
	//
	// - **Open JDK 8**
	//
	// - **Open JDK 7**
	//
	// - **Dragonwell 11**
	//
	// - **Dragonwell 8**
	//
	// - **openjdk-8u191-jdk-alpine3.9**
	//
	// - **openjdk-7u201-jdk-alpine3.9**
	//
	// This parameter is not supported when **Package Type*	- is set to **Image**.
	//
	// example:
	//
	// Open JDK 8
	Jdk *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	// The summary of log collection configurations for Kafka. Valid values:
	//
	// - **kafkaEndpoint**: The endpoint of the Kafka API.
	//
	// - **kafkaInstanceId**: The Kafka instance ID.
	//
	// - **kafkaConfigs**: The configuration summary for one or more log entries. For example values and parameter descriptions, see the **kafkaConfigs*	- request parameter in this topic.
	//
	// example:
	//
	// {"kafkaEndpoint":"10.0.X.XXX:XXXX,10.0.X.XXX:XXXX,10.0.X.XXX:XXXX","kafkaInstanceId":"alikafka_pre-cn-7pp2l8kr****","kafkaConfigs":[{"logType":"file_log","logDir":"/tmp/a.log","kafkaTopic":"test2"},{"logType":"stdout","logDir":"","kafkaTopic":"test"}]}
	KafkaConfigs *string `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	// The labels.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The container health check settings. Containers that fail the health check are shut down and recovered. The following methods are supported:
	//
	// - **exec**: For example, `{"exec":{"command":["sh","-c","cat/home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}`
	//
	// - **httpGet**: For example, `{"httpGet":{"path":"/","port":18091,"scheme":"HTTP","isContainKeyWord":true,"keyWord":"SAE"},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// - **tcpSocket**: For example, `{"tcpSocket":{"port":18091},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// > You can select only one method for health checks.
	//
	// Parameter descriptions:
	//
	// - **exec.command**: The health check command.
	//
	// - **httpGet.path**: The access path.
	//
	// - **httpGet.scheme**: **HTTP*	- or **HTTPS**.
	//
	// - **httpGet.isContainKeyWord**: **true*	- indicates that the keyword is included. **false*	- indicates that the keyword is not included. If this field is missing, the advanced feature is not used.
	//
	// - **httpGet.keyWord**: The custom keyword. The **isContainKeyWord*	- field must be present when this parameter is used.
	//
	// - **tcpSocket.port**: The port for TCP connection detection.
	//
	// - **initialDelaySeconds**: The initial delay for the health check. Default value: 10. Unit: seconds.
	//
	// - **periodSeconds**: The health check period. Default value: 30. Unit: seconds.
	//
	// - **timeoutSeconds**: The health check timeout period. Default value: 1. Unit: seconds. If this parameter is set to 0 or is not set, the default timeout period is 1 second.
	//
	// example:
	//
	// {"exec":{"command":["curl http://localhost:8080"]},"initialDelaySeconds":20,"timeoutSeconds":3}
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// LokiConfigs
	LokiConfigs *string `json:"LokiConfigs,omitempty" xml:"LokiConfigs,omitempty"`
	// The Peak Volume instance ratio.
	MaxSurgeInstanceRatio *int32 `json:"MaxSurgeInstanceRatio,omitempty" xml:"MaxSurgeInstanceRatio,omitempty"`
	// The Peak Volume of instances.
	MaxSurgeInstances *int32 `json:"MaxSurgeInstances,omitempty" xml:"MaxSurgeInstances,omitempty"`
	// The memory size required by each instance, in MB. This value cannot be 0. The memory size has a one-to-one mapping with CPU. Only the following defined specifications are supported:
	//
	// - **1024**: Corresponds to 500 millicores and 1000 millicores of CPU.
	//
	// - **2048**: Corresponds to 500, 1000, and 2000 millicores of CPU.
	//
	// - **4096**: Corresponds to 1000, 2000, and 4000 millicores of CPU.
	//
	// - **8192**: Corresponds to 2000, 4000, and 8000 millicores of CPU.
	//
	// - **12288**: Corresponds to 12000 millicores of CPU.
	//
	// - **16384**: Corresponds to 4000, 8000, and 16000 millicores of CPU.
	//
	// - **24576**: Corresponds to 12000 millicores of CPU.
	//
	// - **32768**: Corresponds to 16000 millicores of CPU.
	//
	// - **65536**: Corresponds to 8000, 16000, and 32000 millicores of CPU.
	//
	// - **131072**: Corresponds to 32000 millicores of CPU.
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The Nacos registry. Valid values:
	//
	// - **0**: SAE built-in Nacos.
	//
	// - **1**: Self-managed Nacos.
	//
	// - **2**: MSE commercial edition Nacos.
	//
	// example:
	//
	// "0"
	MicroRegistration *string `json:"MicroRegistration,omitempty" xml:"MicroRegistration,omitempty"`
	// The registry configuration. This parameter takes effect only when the registry type is MSE Nacos Enterprise Edition.
	//
	// example:
	//
	// {\\"instanceId\\":\\"mse-cn-1ls43******\\",\\"namespace\\":\\"62ee12fb-c279-4da4-be96-21**********\\"}
	MicroRegistrationConfig *string `json:"MicroRegistrationConfig,omitempty" xml:"MicroRegistrationConfig,omitempty"`
	// The microservice governance configuration.
	//
	// - Specifies whether to enable microservice governance (enable):
	//
	//    - true: enabled
	//
	//   - false: disabled
	//
	// - Lossless rolling update configuration (mseLosslessRule):
	//
	//   - delayTime: the delay time.
	//
	//   - enable: specifies whether to enable the lossless online feature. true indicates enabled. false indicates disabled.
	//
	//   - notice: specifies whether to enable the notification feature. true indicates enabled. false indicates disabled.
	//
	//   - warmupTime: the warm-up duration for traffic ramping, in seconds.
	//
	// example:
	//
	// {\\"Enable\\":true,\\"MseLosslessRule\\":{\\"enable\\":true,\\"notice\\":true,\\"delayTime\\":10,\\"warmupTime\\":120,\\"funcType\\":2,\\"aligned\\":false,\\"related\\":false,\\"lossLessDetail\\":false}}
	MicroserviceEngineConfig *string `json:"MicroserviceEngineConfig,omitempty" xml:"MicroserviceEngineConfig,omitempty"`
	// The minimum percentage of available instances. Valid values:
	//
	// - **-1**: The default value, which indicates that the percentage is not used. If this parameter is not specified, the system uses **-1*	- by default.
	//
	// - **0~100**: The unit is percentage, rounded up. For example, if set to **50**%, and the current number of instances is 5, the minimum number of available instances is 3.
	//
	// > When both **MinReadyInstance*	- and **MinReadyInstanceRatio*	- are specified and the value of **MinReadyInstanceRatio*	- is not **-1**, the **MinReadyInstanceRatio*	- parameter takes precedence. For example, if **MinReadyInstances*	- is set to **5*	- and **MinReadyInstanceRatio*	- is set to **50**, the system uses **MinReadyInstanceRatio*	- to calculate the minimum number of available instances.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of available instances. Valid values:
	//
	// - If set to **0**, the application interrupts services during the upgrade process.
	//
	// - If set to **-1**, the system-recommended value is used, which is 25% of the current number of instances. If the current number of instances is 5, 5 × 25% = 1.25, which is rounded up to 2.
	//
	// > Set the minimum number of available instances to ≥ 1 for each rolling deployment to avoid service interruptions.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// The mount description information.
	MountDesc []*DescribeApplicationConfigResponseBodyDataMountDesc `json:"MountDesc,omitempty" xml:"MountDesc,omitempty" type:"Repeated"`
	// The mount point of NAS within the application VPC. If the configuration has not changed during deployment, you do not need to set this parameter (that is, the **MountHost*	- field does not need to be included in the request). To clear the NAS configuration, set the value of this field to an empty string in the request (that is, set the value of the **MountHost*	- field to "").
	//
	// example:
	//
	// example.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The application ID on the Microservices Engine (MSE) side.
	//
	// example:
	//
	// xxxxxxx@xxxxx
	MseApplicationId *string `json:"MseApplicationId,omitempty" xml:"MseApplicationId,omitempty"`
	// The application name after the SAE service is registered with MSE.
	//
	// example:
	//
	// cn-shenzhen-alb-demo-5c****
	MseApplicationName *string `json:"MseApplicationName,omitempty" xml:"MseApplicationName,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The NAS mount configurations.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// NAS ID。
	//
	// example:
	//
	// AKSN****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The application version. Valid values:
	//
	// - lite: Lite Edition
	//
	// - std: Standard Edition
	//
	// - pro: Professional Edition
	//
	// example:
	//
	// pro
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	// The RAM role for identity authentication.
	//
	// > Create an OIDC identity provider and an identity provider role in the same region in advance. For more information, see [Create an OIDC identity provider](https://help.aliyun.com/document_detail/2331022.html) and [Create a role for SSO identity provider](https://help.aliyun.com/document_detail/2331016.html).
	//
	// example:
	//
	// sae-test
	OidcRoleName *string `json:"OidcRoleName,omitempty" xml:"OidcRoleName,omitempty"`
	// The AccessKey ID for OSS read/write operations.
	//
	// example:
	//
	// xxxxxx
	OssAkId *string `json:"OssAkId,omitempty" xml:"OssAkId,omitempty"`
	// The AccessKey Secret for OSS read/write operations.
	//
	// example:
	//
	// xxxxxx
	OssAkSecret *string `json:"OssAkSecret,omitempty" xml:"OssAkSecret,omitempty"`
	// The OSS mount description.
	OssMountDescs []*DescribeApplicationConfigResponseBodyDataOssMountDescs `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty" type:"Repeated"`
	// The application package type. Valid values:
	//
	// - When you deploy with Java, **FatJar**, **War**, and **Image*	- are supported.
	//
	// - When you deploy with PHP, the following types are supported:
	//
	//     - **PhpZip**
	//
	//     - **IMAGE_PHP_5_4**
	//
	//     - **IMAGE_PHP_5_4_ALPINE**
	//
	//     - **IMAGE_PHP_5_5**
	//
	//     - **IMAGE_PHP_5_5_ALPINE**
	//
	//     - **IMAGE_PHP_5_6**
	//
	//     - **IMAGE_PHP_5_6_ALPINE**
	//
	//     - **IMAGE_PHP_7_0**
	//
	//     - **IMAGE_PHP_7_0_ALPINE**
	//
	//     - **IMAGE_PHP_7_1**
	//
	//     - **IMAGE_PHP_7_1_ALPINE**
	//
	//     - **IMAGE_PHP_7_2**
	//
	//     - **IMAGE_PHP_7_2_ALPINE**
	//
	//     - **IMAGE_PHP_7_3**
	//
	//     - **IMAGE_PHP_7_3_ALPINE**
	//
	// example:
	//
	// War
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The deployment package URL. If your deployment package is uploaded through SAE, note the following:
	//
	// - This URL cannot be used for direct download. Use the GetPackageVersionAccessableUrl operation to obtain a downloadable URL (valid for 10 minutes).
	//
	// - SAE retains the package for a maximum of 90 days. After 90 days, the URL is no longer returned and the package is no longer available for download.
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package. This parameter is required when **Package Type*	- is set to **FatJar*	- or **War**.
	//
	// example:
	//
	// 1.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The PHP version on which the PHP deployment package depends. Images are not supported.
	//
	// example:
	//
	// PHP-FPM 7.0
	Php *string `json:"Php,omitempty" xml:"Php,omitempty"`
	// The mount path for PHP application monitoring. Make sure that the PHP server loads the configuration file from this path.
	//
	// You do not need to manage the configuration content. SAE automatically renders the correct configuration file.
	//
	// example:
	//
	// /usr/local/etc/php/conf.d/arms.ini
	PhpArmsConfigLocation *string `json:"PhpArmsConfigLocation,omitempty" xml:"PhpArmsConfigLocation,omitempty"`
	// The content of the PHP configuration file.
	//
	// example:
	//
	// k1=v1
	PhpConfig *string `json:"PhpConfig,omitempty" xml:"PhpConfig,omitempty"`
	// The mount path of the PHP application startup configuration. Make sure that the PHP server uses this configuration file for startup.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// The script that is run after the container starts. A script is triggered immediately after the container is created. Format: `{"exec":{"command":["cat","/etc/group"\\]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script that is run before the container stops. A script is triggered before the container is deleted. Format: `{"exec":{"command":["cat","/etc/group"\\]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The programming language of the technology stack used to create the application. Valid values:
	//
	// - **java**: Java.
	//
	// - **php**: PHP.
	//
	// - **other**: Other languages, such as Python, C++, Go, .NET, and Node.js.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// Enables K8s Service-based service registration and discovery.
	//
	// example:
	//
	// {     "serviceName": "bwm-poc-sc-gateway-cn-beijing-front",     "namespaceId": "cn-beijing:front",     "portAndProtocol": {         "18012": "TCP"     },     "portProtocols": [         {             "port": "18012",             "protocol": "TCP"         }     ],     "enable": true }
	PvtzDiscovery *string `json:"PvtzDiscovery,omitempty" xml:"PvtzDiscovery,omitempty"`
	// The Python environment. PYTHON 3.9.15 is supported.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// The custom installation module dependencies. By default, the dependencies defined in the requirements.txt file in the root directory are installed. If no dependencies are configured or custom packages are needed, you can specify the dependencies to install.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string                                                `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	RaspConfig    []*DescribeApplicationConfigResponseBodyDataRaspConfig `json:"RaspConfig,omitempty" xml:"RaspConfig,omitempty" type:"Repeated"`
	// The application startup status check. Containers that fail multiple health checks are shut down and restarted. Containers that do not pass the health check do not receive SLB traffic. The **exec**, **httpGet**, and **tcpSocket*	- methods are supported. For specific examples, see the **Liveness*	- parameter.
	//
	// > You can select only one method for health checks.
	//
	// example:
	//
	// {"exec":{"command":["curl http://localhost:8080"]},"initialDelaySeconds":20,"timeoutSeconds":5}
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The region ID.
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
	// The resource type. Only `application` is supported.
	//
	// example:
	//
	// application
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The Secret mount description.
	SecretMountDesc []*DescribeApplicationConfigResponseBodyDataSecretMountDesc `json:"SecretMountDesc,omitempty" xml:"SecretMountDesc,omitempty" type:"Repeated"`
	// The security group ID.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The canary release tags configured for the application.
	ServiceTags map[string]*string `json:"ServiceTags,omitempty" xml:"ServiceTags,omitempty"`
	// The sidecar container configuration.
	SidecarContainersConfig []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfig `json:"SidecarContainersConfig,omitempty" xml:"SidecarContainersConfig,omitempty" type:"Repeated"`
	// The settings for log collection to Simple Log Service (SLS).
	//
	// - To use SLS resources that are automatically created by Serverless App Engine (SAE): `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - To use custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// Parameter descriptions:
	//
	// - **projectName**: The name of the project in SLS.
	//
	// - **logDir**: The log path.
	//
	// - **logType**: The log type. **stdout*	- indicates container standard output logs. You can set only one stdout entry. If this parameter is not set, file logs are collected.
	//
	// - **logstoreName**: The name of the Logstore in SLS.
	//
	// - **logtailName**: The name of the Logtail in SLS. If this parameter is not specified, a new Logtail is created through automatic creation.
	//
	// If the SLS collection configuration has not changed during multiple deployments, you do not need to set this parameter (that is, the **SlsConfigs*	- field does not need to be included in the request). If you no longer need the SLS collection feature, set the value of this field to an empty string in the request (that is, set the value of the **SlsConfigs*	- field to "").
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// sls log env tags
	SlsLogEnvTags *string `json:"SlsLogEnvTags,omitempty" xml:"SlsLogEnvTags,omitempty"`
	// The startup probe of the application.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"/bin/sh\\",\\"-c\\",\\"#!Note: If microservice config is enabled, the application will be automatically injected with the prestop configuration for lossless offline. If you delete this prestop configuration, lossless offline will not be effective.\\\\n echo stop > /tmp/prestop; /home/admin/.tools/curl http://127.0.0.1:54199/offline; sleep 30\\"]}}
	StartupProbe *string `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	// Configures K8s Service-based service registration and discovery with end-to-end canary release.
	//
	// example:
	//
	// {\\"enable\\":\\"false\\",\\"namespaceId\\":\\"cn-beijing:test\\",\\"portAndProtocol\\":{\\"2000:TCP\\":\\"18081\\"},\\"portProtocols\\":[{\\"port\\":2000,\\"protocol\\":\\"TCP\\",\\"targetPort\\":18081}],\\"pvtzDiscoveryName\\":\\"cn-beijing-1421801774382676\\",\\"serviceId\\":\\"3513\\",\\"serviceName\\":\\"demo-gray.test\\"}
	SwimlanePvtzDiscovery *string `json:"SwimlanePvtzDiscovery,omitempty" xml:"SwimlanePvtzDiscovery,omitempty"`
	// The tag information.
	Tags []*DescribeApplicationConfigResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The graceful shutdown timeout period. Default value: 30. Unit: seconds. Valid values: 1 to 300.
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
	// The Tomcat configuration. Set this parameter to "" or "{}" to delete the configuration:
	//
	// - **port**: The port number. Valid values: 1024 to 65535. Ports smaller than 1024 require root permissions. Because the container is configured with admin permissions, specify a port greater than 1024. Default value: 8080.
	//
	// - **contextPath**: The access path. Default value: root directory "/".
	//
	// - **maxThreads**: The maximum number of connections in the connection pool. Default value: 400.
	//
	// - **uriEncoding**: The encoding format of Tomcat. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. Default value: **ISO-8859-1**.
	//
	// - **useBodyEncoding**: Specifies whether to use **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	// The deployment policy. When the minimum number of available instances is 1, the value of the **UpdateStrategy*	- field is "". When the minimum number of available instances is greater than 1, examples are as follows:
	//
	// - Canary release of 1 instance + 2 subsequent batches + automatic batching + 1-minute batch interval: `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}`
	//
	//
	//
	// - Canary release of 1 instance + 2 subsequent batches + manual batching: `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"manual"},"grayUpdate":{"gray":1}}`
	//
	// - 2 batches + automatic batching + 0-minute batch interval: `{"type":"BatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":0}}`
	//
	// Parameter descriptions:
	//
	// - **type**: The release policy type. Valid values: **GrayBatchUpdate*	- (grayscale batch release) and **BatchUpdate*	- (batch release).
	//
	// - **batchUpdate**: The batch release policy.
	//
	//     - **batch**: The number of release batches.
	//
	//     - **releaseType**: The processing method between batches. Valid values: **auto*	- (automatic) and **manual*	- (manual).
	//
	//     - **batchWaitTime**: The interval between deployments within a batch, in seconds.
	//
	// - **grayUpdate**: The remaining batches after grayscale release. This parameter is required when **type*	- is set to **GrayBatchUpdate**.
	//
	// example:
	//
	// {"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}
	UpdateStrategy *string `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
	// vSwitch ID。
	//
	// example:
	//
	// vsw-2ze559r1z1bpwqxwp****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-2ze0i263cnn311nvj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The startup options for the WAR package application. The default startup command for the application: `java $JAVA_OPTS $CATALINA_OPTS -Options org.apache.catalina.startup.Bootstrap "$@" start`.
	//
	// example:
	//
	// custom-option
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The Tomcat version on which the deployment package depends. Valid values:
	//
	// - **apache-tomcat-7.0.91**
	//
	// - **apache-tomcat-8.5.42**
	//
	// This parameter is not supported when **Package Type*	- is set to **Image**.
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

func (s *DescribeApplicationConfigResponseBodyData) GetIdleHour() *string {
	return s.IdleHour
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

func (s *DescribeApplicationConfigResponseBodyData) GetLabels() map[string]*string {
	return s.Labels
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

func (s *DescribeApplicationConfigResponseBodyData) GetRaspConfig() []*DescribeApplicationConfigResponseBodyDataRaspConfig {
	return s.RaspConfig
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

func (s *DescribeApplicationConfigResponseBodyData) SetIdleHour(v string) *DescribeApplicationConfigResponseBodyData {
	s.IdleHour = &v
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

func (s *DescribeApplicationConfigResponseBodyData) SetLabels(v map[string]*string) *DescribeApplicationConfigResponseBodyData {
	s.Labels = v
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

func (s *DescribeApplicationConfigResponseBodyData) SetRaspConfig(v []*DescribeApplicationConfigResponseBodyDataRaspConfig) *DescribeApplicationConfigResponseBodyData {
	s.RaspConfig = v
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
	if s.RaspConfig != nil {
		for _, item := range s.RaspConfig {
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
	// ConfigMap ID。
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	// The ConfigMap name.
	//
	// example:
	//
	// test
	ConfigMapName *string `json:"ConfigMapName,omitempty" xml:"ConfigMapName,omitempty"`
	// The ConfigMap key-value pair.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The container mount path.
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
	// The mount path.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The temporary storage name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The image startup command. This command must be an executable object that exists in the container. Example:
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
	// Based on the preceding example, `Command="echo", CommandArgs=["abc", ">", "file0"]`.
	//
	// example:
	//
	// /bin/sh
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments for the image startup command. These are the arguments required by the startup command **Command**. Format:
	//
	// `["a","b"]`
	//
	// In the preceding example, `CommandArgs=["abc", ">", "file0"]`, where `["abc", ">", "file0"]` must be converted to the String type and the internal format is a JSON array. If this parameter is not required, leave it empty.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The ConfigMap information.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The shared temporary storage.
	EmptyDirDesc []*DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// The container environment variable parameters. You can customize environment variables or reference ConfigMap instances. To reference a ConfigMap instance, create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//     - **name**: the name of the environment variable.
	//
	//     - **value**: the value of the environment variable. This takes priority over valueFrom.
	//
	// - Reference a ConfigMap instance (valueFrom)
	//
	//     - **name**: the name of the environment variable. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<ConfigMap name>`, such as `sae-sys-configmap-all-test1`.
	//
	//     - **valueFrom**: the reference of the environment variable. Set the value to `configMapRef`.
	//
	//     - **configMapId**: the ID of the ConfigMap instance.
	//
	//     - **key**: the key. Do not set this field if you want to reference all keys.
	//
	// - Reference a secret (valueFrom)
	//
	//     - **name**: the name of the environment variable. You can reference a single key or all keys. To reference all keys, enter `sae-sys-secret-all-<secret name>`, such as `sae-sys-secret-all-test1`.
	//
	//     - **valueFrom**: the reference of the environment variable. Set the value to `secretRef`.
	//
	//     - **secretId**: the ID of the secret.
	//
	//     - **key**: the key. Do not set this field if you want to reference all keys.
	//
	// example:
	//
	// [{"name":"TEST_ENV_KEY","value":"TEST_ENV_VAR"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The image URL used by the init container.
	//
	// [_single.resp.200.props.Data.InitContainersConfig.items.Env
	//
	// example:
	//
	// registry.cn-shenzhen.aliyuncs.com/sae-serverless-demo/sae-demo:microservice-java-provider-v1.0
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The init container name.
	//
	// example:
	//
	// init-container
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Secret mount description.
	SecretMountDesc []*DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc `json:"SecretMountDesc,omitempty" xml:"SecretMountDesc,omitempty" type:"Repeated"`
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

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) GetSecretMountDesc() []*DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc {
	return s.SecretMountDesc
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

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfig) SetSecretMountDesc(v []*DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) *DescribeApplicationConfigResponseBodyDataInitContainersConfig {
	s.SecretMountDesc = v
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
	if s.SecretMountDesc != nil {
		for _, item := range s.SecretMountDesc {
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
	// The ConfigMap name.
	//
	// example:
	//
	// test
	ConfigMapName *string `json:"ConfigMapName,omitempty" xml:"ConfigMapName,omitempty"`
	// The ConfigMap key-value pair.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The container mount path.
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
	// The mount path of the data volume in the container.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The temporary storage name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

type DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc struct {
	// The key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The Secret instance ID.
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The Secret instance name.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) GetKey() *string {
	return s.Key
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) GetSecretId() *int64 {
	return s.SecretId
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) GetSecretName() *string {
	return s.SecretName
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) SetKey(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc {
	s.Key = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) SetSecretId(v int64) *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc {
	s.SecretId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) SetSecretName(v string) *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc {
	s.SecretName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataInitContainersConfigSecretMountDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataMountDesc struct {
	// The container mount path.
	//
	// example:
	//
	// /tmp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The NAS relative file directory.
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
	// The bucket name.
	//
	// example:
	//
	// oss-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The directory or OSS object that you created in OSS. If the OSS mount directory does not exist, an exception is triggered.
	//
	// example:
	//
	// data/user.data
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	// The container path in SAE. If the path already exists, it is overwritten. If the path does not exist, it is created.
	//
	// example:
	//
	// /usr/data/user.data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// Indicates whether the container path has read-only permission to the mounted directory resources. Valid values:
	//
	// - **true**: Read-only permission.
	//
	// - **false**: Read and write permission.
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

type DescribeApplicationConfigResponseBodyDataRaspConfig struct {
	EnableRasp  *bool   `json:"EnableRasp,omitempty" xml:"EnableRasp,omitempty"`
	RaspAppKey  *string `json:"RaspAppKey,omitempty" xml:"RaspAppKey,omitempty"`
	RaspAppName *string `json:"RaspAppName,omitempty" xml:"RaspAppName,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataRaspConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataRaspConfig) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataRaspConfig) GetEnableRasp() *bool {
	return s.EnableRasp
}

func (s *DescribeApplicationConfigResponseBodyDataRaspConfig) GetRaspAppKey() *string {
	return s.RaspAppKey
}

func (s *DescribeApplicationConfigResponseBodyDataRaspConfig) GetRaspAppName() *string {
	return s.RaspAppName
}

func (s *DescribeApplicationConfigResponseBodyDataRaspConfig) SetEnableRasp(v bool) *DescribeApplicationConfigResponseBodyDataRaspConfig {
	s.EnableRasp = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataRaspConfig) SetRaspAppKey(v string) *DescribeApplicationConfigResponseBodyDataRaspConfig {
	s.RaspAppKey = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataRaspConfig) SetRaspAppName(v string) *DescribeApplicationConfigResponseBodyDataRaspConfig {
	s.RaspAppName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataRaspConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataSecretMountDesc struct {
	// The key with Base64-encoded data value.
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
	// The queried Secret instance ID.
	//
	// example:
	//
	// 520
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The Secret instance name.
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
	// The ACR Enterprise instance ID. This parameter is required when **ImageUrl*	- is from ACR Enterprise Edition.
	//
	// example:
	//
	// cri-fhzlneorxala66ip
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The image startup command. This command must be an executable object that exists in the container. Example:
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
	// Based on the preceding example, `Command="echo", CommandArgs=["abc", ">", "file0"]`.
	//
	// example:
	//
	// /bin/sh
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments for the image startup command. These are the arguments required by the startup command **Command**. Format:
	//
	// `["a","b"]`
	//
	// In the preceding example, `CommandArgs=["abc", ">", "file0"]`, where `["abc", ">", "file0"]` must be converted to the String type and the internal format is a JSON array. If this parameter is not required, leave it empty.
	//
	// example:
	//
	// [\\"-c\\",\\"echo \\\\\\"test\\\\\\" > /home/nas/test.log && sleep 10000000s\\"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The ConfigMap mount description. Use the configuration items created on the namespace configuration page to inject configuration information into the container. Parameter descriptions:
	//
	// - **configMapId**: The ConfigMap instance ID. You can obtain this ID by calling the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation.
	//
	// - **key**: The key-value pair.
	//
	// > You can mount all keys by passing the `sae-sys-configmap-all` parameter.
	//
	// - **mountPath**: The mount path.
	//
	// - **ConfigMapName**: The ConfigMap name.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The maximum CPU resources of the primary container that the sidecar container can use.
	//
	// example:
	//
	// 500
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The shared temporary storage. Sets a temporary storage directory and mounts it to the primary container and sidecar container.
	EmptyDirDesc []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// The container environment variable parameters. Custom values or references to configuration items are supported. To reference a configuration item, create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//     - **name**: The environment variable name.
	//
	//     - **value**: The environment variable value. This value takes precedence over valueFrom.
	//
	// - Reference to a configuration item (valueFrom)
	//
	//     - **name**: The environment variable name. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<ConfigMap name>`, such as `sae-sys-configmap-all-test1`.
	//
	//     - **valueFrom**: The environment variable reference. Set the value to `configMapRef`.
	//
	//         - **configMapId**: The ConfigMap ID.
	//
	//         - **key**: The key. If all keys are referenced, do not set this field.
	//
	// example:
	//
	// [{\\"name\\":\\"k1\\",\\"value\\":\\"c8e3a815-e5d3-4adf-abb3-98b106a607c4\\"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The image URL.
	//
	// example:
	//
	// registry.cn-beijing.aliyuncs.com/sae-dev-test/nginx:stable
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The container health check.
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The maximum memory resources of the primary container that the sidecar container can use.
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
	// The script that is run after the container starts.
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script that is run before the container stops.
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The application startup status check.
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The Secret mount description.
	SecretMountDesc []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc `json:"SecretMountDesc,omitempty" xml:"SecretMountDesc,omitempty" type:"Repeated"`
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

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetLiveness() *string {
	return s.Liveness
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetPostStart() *string {
	return s.PostStart
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetPreStop() *string {
	return s.PreStop
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetReadiness() *string {
	return s.Readiness
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetSecretMountDesc() []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc {
	return s.SecretMountDesc
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

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetLiveness(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.Liveness = &v
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

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetPostStart(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.PostStart = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetPreStop(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.PreStop = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetReadiness(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.Readiness = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetSecretMountDesc(v []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.SecretMountDesc = v
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
	if s.SecretMountDesc != nil {
		for _, item := range s.SecretMountDesc {
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
	// The ConfigMap instance ID.
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
	// The ConfigMap key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The container mount path.
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
	// The mount path of the data volume in the container.
	//
	// example:
	//
	// /mnt/cache
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The temporary storage name.
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

type DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc struct {
	// The key with Base64-encoded data value.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The Secret instance ID.
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The Secret instance name.
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) GetKey() *string {
	return s.Key
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) GetSecretId() *int64 {
	return s.SecretId
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) GetSecretName() *string {
	return s.SecretName
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) SetKey(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc {
	s.Key = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) SetMountPath(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) SetSecretId(v int64) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc {
	s.SecretId = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) SetSecretName(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc {
	s.SecretName = &v
	return s
}

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationConfigResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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

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
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Client error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The application information.
	Data *DescribeApplicationConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidApplication.NotFound
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
	// Indicates whether the application configuration information was obtained. Valid values:
	//
	// - **true**: The application configuration information was obtained.
	//
	// - **false**: The application configuration information failed to be obtained.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID, which is used to accurately query call information.
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
	// The ARN of the RAM role required to pull images across accounts. For more information, see [Pull Alibaba Cloud images across accounts](https://help.aliyun.com/document_detail/190675.html) and [Implement cross-account authorization by using RAM roles](https://help.aliyun.com/document_detail/223585.html).
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
	//
	// example:
	//
	// 4.4.2
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The ALB gateway readiness gate configuration.
	//
	// example:
	//
	// default
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
	// The SAE application type. Valid values:
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
	// - **true**: Associate an EIP.
	//
	// - **false**: Do not associate an EIP.
	//
	// example:
	//
	// true
	AssociateEip *bool `json:"AssociateEip,omitempty" xml:"AssociateEip,omitempty"`
	// The baseline application ID.
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
	//
	// example:
	//
	// cms-id
	CmsServiceId *string `json:"CmsServiceId,omitempty" xml:"CmsServiceId,omitempty"`
	// The image startup command. The command must be an executable object that exists in the container. Example:
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
	// The image startup command parameters. The parameters required by the **Command*	- startup command. Format:
	//
	// `["a","b"]`
	//
	// In the example of the **Command*	- parameter, `CommandArgs=["abc", ">", "file0"]`, where `["abc", ">", "file0"]` must be converted to a string in JSON array format. If this parameter is not required, leave it empty.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The ConfigMap information.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The CPU required for each instance, in millicores. The value cannot be 0. Only the following defined specifications are supported:
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
	// The custom image type. If the image is not a custom image, set this parameter to an empty string. Valid values:
	//
	// - internet: Public image.
	//
	// - intranet: Private image.
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
	// The .NET framework version. Valid values:
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
	// The application runtime environment version in the HSF framework, such as the Ali-Tomcat container.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// The shared ephemeral storage.
	EmptyDirDesc []*DescribeApplicationConfigResponseBodyDataEmptyDirDesc `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// Specifies whether to connect to AHAS. Valid values:
	//
	// - **true**: Connect to AHAS.
	//
	// - **false**: Do not connect to AHAS.
	//
	// example:
	//
	// true
	EnableAhas *string `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	// Specifies whether to enable the CPU Burst feature. Valid values:
	//
	// - true: Enable.
	//
	// - false: Disable.
	//
	// example:
	//
	// true
	EnableCpuBurst *string `json:"EnableCpuBurst,omitempty" xml:"EnableCpuBurst,omitempty"`
	// Specifies whether to enable traffic canary release rules. This rule applies only to applications that use the Spring Cloud and Dubbo frameworks. Valid values:
	//
	// - **true**: Enable canary release rules.
	//
	// - **false**: Disable canary release rules.
	//
	// example:
	//
	// false
	EnableGreyTagRoute *bool `json:"EnableGreyTagRoute,omitempty" xml:"EnableGreyTagRoute,omitempty"`
	// Specifies whether to enable the idle mode. Valid values:
	//
	// - true: Enable.
	//
	// - false: Disable.
	//
	// example:
	//
	// false
	EnableIdle *bool `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	// Specifies whether to reuse the namespace agent version configuration.
	EnableNamespaceAgentVersion *bool `json:"EnableNamespaceAgentVersion,omitempty" xml:"EnableNamespaceAgentVersion,omitempty"`
	// Specifies whether to enable the new ARMS features. Valid values:
	//
	// - true: Enable.
	//
	// - false: Disable.
	//
	// example:
	//
	// false
	EnableNewArms *bool `json:"EnableNewArms,omitempty" xml:"EnableNewArms,omitempty"`
	// Specifies whether to enable Prometheus custom metric collection.
	EnablePrometheus *bool `json:"EnablePrometheus,omitempty" xml:"EnablePrometheus,omitempty"`
	// The container environment variable parameters. You can customize the parameters or reference configuration items. To reference configuration items, create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//     - **name**: The environment variable name.
	//
	//     - **value**: The environment variable value.
	//
	// - Reference configuration items
	//
	//     - **name**: The environment variable name. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<Configuration item name>`, for example, `sae-sys-configmap-all-test1`.
	//
	//     - **valueFrom**: The environment variable reference. Set the value to `configMapRef`.
	//
	//     - **configMapId**: The configuration item ID.
	//
	//     - **key**: The key. If all key-value pairs are referenced, do not set this field.
	//
	// example:
	//
	// [{"name":"TEST_ENV_KEY","value":"TEST_ENV_VAR"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 1
	GpuCount *string `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The GPU type.
	//
	// example:
	//
	// PPU810E
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// The headless PrivateZone service discovery configuration for stateful applications. The value is a JSON string.
	//
	// example:
	//
	// {"enable":"true","serviceName":"svc1","namespaceId":"cn-shenzhen"}
	HeadlessPvtzDiscovery *string `json:"HeadlessPvtzDiscovery,omitempty" xml:"HeadlessPvtzDiscovery,omitempty"`
	// The NGINX runtime environment version used by the HTML application.
	//
	// example:
	//
	// nginx 1.28
	Html *string `json:"Html,omitempty" xml:"Html,omitempty"`
	// The application idle time period configuration. The value is a JSON array string that contains startTime and endTime. Specify the time on the hour.
	//
	// example:
	//
	// [{"startTime":"1:00","endTime":"16:00"}]
	IdleHour *string `json:"IdleHour,omitempty" xml:"IdleHour,omitempty"`
	// The ID of the corresponding Secret.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The image address. This parameter is required when **Package Type*	- is set to **Image**.
	//
	// example:
	//
	// docker.io/library/nginx:1.14.2
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The init container configuration.
	InitContainersConfig []*DescribeApplicationConfigResponseBodyDataInitContainersConfig `json:"InitContainersConfig,omitempty" xml:"InitContainersConfig,omitempty" type:"Repeated"`
	// Specifies whether the application is a stateful application.
	IsStateful *bool `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The application startup parameters for the JAR package. Default application startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// start
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The application startup options for the JAR package. Default application startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// -Dtest=true
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The JDK version on which the deployment package depends. Supported versions:
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
	// The summary information of the configuration for collecting logs to Kafka. Valid values:
	//
	// - **kafkaEndpoint**: The service registration address of the Kafka API.
	//
	// - **kafkaInstanceId**: The Kafka instance ID.
	//
	// - **kafkaConfigs**: The summary information of the configuration for one or more logs. For example values and metric descriptions, see the **kafkaConfigs*	- request parameter in this topic.
	//
	// example:
	//
	// {"kafkaEndpoint":"10.0.X.XXX:XXXX,10.0.X.XXX:XXXX,10.0.X.XXX:XXXX","kafkaInstanceId":"alikafka_pre-cn-7pp2l8kr****","kafkaConfigs":[{"logType":"file_log","logDir":"/tmp/a.log","kafkaTopic":"test2"},{"logType":"stdout","logDir":"","kafkaTopic":"test"}]}
	KafkaConfigs *string `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	// The tags.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The container health check. Containers that failed the health check are shutdown and recover. Supported methods:
	//
	// - **exec**: Example: `{"exec":{"command":["sh","-c","cat/home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}`
	//
	// - **httpGet**: Example: `{"httpGet":{"path":"/","port":18091,"scheme":"HTTP","isContainKeyWord":true,"keyWord":"SAE"},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// - **tcpSocket**: Example: `{"tcpSocket":{"port":18091},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// > You can select only one method for the health check.
	//
	// Metric descriptions:
	//
	// - **exec.command**: Settings for the health check command.
	//
	// - **httpGet.path**: The access path.
	//
	// - **httpGet.scheme**: **HTTP*	- or **HTTPS**.
	//
	// - **httpGet.isContainKeyWord**: **true*	- indicates that the keyword is included, and **false*	- indicates that the keyword is not included. If this field is missing, the advanced feature is not used.
	//
	// - **httpGet.keyWord**: The custom keyword. The **isContainKeyWord*	- field is required when this field is used.
	//
	// - **tcpSocket.port**: The port for TCP connection detection.
	//
	// - **initialDelaySeconds**: Settings for the health check delay detection time. Default value: 10. Unit: seconds.
	//
	// - **periodSeconds**: Settings for the health check epoch. Default value: 30. Unit: seconds.
	//
	// - **timeoutSeconds**: Settings for the health check timeout waiting time. Default value: 1. Unit: seconds. If this parameter is set to 0 or not set, the default timeout waiting time is 1 second.
	//
	// example:
	//
	// {"exec":{"command":["curl http://localhost:8080"]},"initialDelaySeconds":20,"timeoutSeconds":3}
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// LokiConfigs
	//
	// example:
	//
	// []
	LokiConfigs *string `json:"LokiConfigs,omitempty" xml:"LokiConfigs,omitempty"`
	// The peak volume ratio of surge instances.
	//
	// example:
	//
	// -1
	MaxSurgeInstanceRatio *int32 `json:"MaxSurgeInstanceRatio,omitempty" xml:"MaxSurgeInstanceRatio,omitempty"`
	// The peak volume of surge instances.
	//
	// example:
	//
	// -1
	MaxSurgeInstances *int32 `json:"MaxSurgeInstances,omitempty" xml:"MaxSurgeInstances,omitempty"`
	// The memory required for each instance, in MB. The value cannot be 0. The memory has a one-to-one mapping with the CPU. Only the following defined specifications are supported:
	//
	// - **1024**: corresponds to 500 millicores and 1000 millicores.
	//
	// - **2048**: corresponds to 500, 1000, and 2000 millicores.
	//
	// - **4096**: corresponds to 1000, 2000, and 4000 millicores.
	//
	// - **8192**: corresponds to 2000, 4000, and 8000 millicores.
	//
	// - **12288**: corresponds to 12000 millicores.
	//
	// - **16384**: corresponds to 4000, 8000, and 16000 millicores.
	//
	// - **24576**: corresponds to 12000 millicores.
	//
	// - **32768**: corresponds to 16000 millicores.
	//
	// - **65536**: corresponds to 8000, 16000, and 32000 millicores.
	//
	// - **131072**: corresponds to 32000 millicores.
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The Nacos registry center. Valid values:
	//
	// - **0**: SAE built-in Nacos.
	//
	// - **1**: User-managed Nacos.
	//
	// - **2**: MSE Nacos Commercial Edition.
	//
	// example:
	//
	// "0"
	MicroRegistration *string `json:"MicroRegistration,omitempty" xml:"MicroRegistration,omitempty"`
	// The registry center configuration. This parameter takes effect only when the registry center type is MSE Nacos Enterprise Edition.
	//
	// example:
	//
	// {\\"instanceId\\":\\"mse-cn-1ls43******\\",\\"namespace\\":\\"62ee12fb-c279-4da4-be96-21**********\\"}
	MicroRegistrationConfig *string `json:"MicroRegistrationConfig,omitempty" xml:"MicroRegistrationConfig,omitempty"`
	// Configures the microservice governance feature.
	//
	// - Whether to enable microservice governance (enable):
	//
	//    - true: Enabled.
	//
	//   - false: Disabled.
	//
	// - Configure lossless online and offline (mseLosslessRule):
	//
	//   - delayTime: The delay time.
	//
	//   - enable: Specifies whether to enable the lossless online feature. Valid values: true (enabled) and false (disabled).
	//
	//   - notice: Specifies whether to enable the notification feature. Valid values: true (enabled) and false (disabled).
	//
	//   - warmupTime: The warm-up duration for small traffic, in seconds.
	//
	// example:
	//
	// {\\"Enable\\":true,\\"MseLosslessRule\\":{\\"enable\\":true,\\"notice\\":true,\\"delayTime\\":10,\\"warmupTime\\":120,\\"funcType\\":2,\\"aligned\\":false,\\"related\\":false,\\"lossLessDetail\\":false}}
	MicroserviceEngineConfig *string `json:"MicroserviceEngineConfig,omitempty" xml:"MicroserviceEngineConfig,omitempty"`
	// The minimum percentage of ready instances. Valid values:
	//
	// - **-1**: The default value, which indicates that the percentage is not used. If you do not specify this parameter, the system defaults to **-1**.
	//
	// - **0 to 100**: The unit is percentage, rounded up. For example, if you set this parameter to **50*	- and the current number of instances is 5, the minimum number of ready instances is 3.
	//
	// > If both **MinReadyInstance*	- and **MinReadyInstanceRatio*	- are specified and the value of **MinReadyInstanceRatio*	- is not **-1**, the **MinReadyInstanceRatio*	- parameter prevails. For example, if **MinReadyInstances*	- is set to **5*	- and **MinReadyInstanceRatio*	- is set to **50**, the system uses **MinReadyInstanceRatio*	- to calculate the minimum number of ready instances.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of ready instances. Valid values:
	//
	// - If you set this parameter to **0**, services are interrupted during the application upgrade.
	//
	// - If you set this parameter to **-1**, the minimum number of ready instances uses the system-recommended value, which is 25% of the existing instances. For example, if the current number of instances is 5, 5 × 25% = 1.25. After rounding up, the minimum number of ready instances is 2.
	//
	// > Set the minimum number of ready instances for each rolling deployment to a value greater than or equal to 1 to ensure service continuity.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// The mount description.
	MountDesc []*DescribeApplicationConfigResponseBodyDataMountDesc `json:"MountDesc,omitempty" xml:"MountDesc,omitempty" type:"Repeated"`
	// The NAS mount target within the application VPC. If the configuration remains unchanged during deployment, you do not need to set this parameter (that is, the **MountHost*	- field is not required in the request). To clear the NAS configuration, set the value of this field to an empty string in the request (that is, set the **MountHost*	- field to "").
	//
	// example:
	//
	// example.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The application ID on the MSE side.
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
	// The NAS mount configuration.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// The NAS ID.
	//
	// example:
	//
	// AKSN****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The application edition. Valid values:
	//
	// - lite: Lite Edition.
	//
	// - std: Standard Edition.
	//
	// - pro: Professional Edition.
	//
	// example:
	//
	// pro
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	// The RAM role for the identity authentication service.
	//
	// > You must create an OIDC identity provider and an identity provider role in the same region in advance. For more information, see [Create an OIDC identity provider](https://help.aliyun.com/document_detail/2331022.html) and [Create a role-based SSO identity provider](https://help.aliyun.com/document_detail/2331016.html).
	//
	// example:
	//
	// sae-test
	OidcRoleName *string `json:"OidcRoleName,omitempty" xml:"OidcRoleName,omitempty"`
	// The AccessKey ID used to read and write OSS data.
	//
	// example:
	//
	// xxxxxx
	OssAkId *string `json:"OssAkId,omitempty" xml:"OssAkId,omitempty"`
	// The AccessKey Secret used to read and write OSS data.
	//
	// example:
	//
	// xxxxxx
	OssAkSecret *string `json:"OssAkSecret,omitempty" xml:"OssAkSecret,omitempty"`
	// The OSS mount description.
	OssMountDescs []*DescribeApplicationConfigResponseBodyDataOssMountDescs `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty" type:"Repeated"`
	// The application package type. Valid values:
	//
	// - When you deploy a Java application, valid values are **FatJar**, **War**, and **Image**.
	//
	// - When you deploy a PHP application, valid values are:
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
	// The deployment package URL. If your deployment package is uploaded through SAE, note the following points:
	//
	// - The URL cannot be used to download the package directly. You must call the GetPackageVersionAccessableUrl operation to obtain a downloadable URL, which is valid for 10 minutes.
	//
	// - SAE retains the package for a maximum of 90 days. After 90 days, the URL is not returned and the package cannot be downloaded.
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version number of the deployment package. This parameter is required when **PackageType*	- is set to **FatJar*	- or **War**.
	//
	// example:
	//
	// 1.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The PHP version required by the PHP deployment package. Images are not supported.
	//
	// example:
	//
	// PHP-FPM 7.0
	Php *string `json:"Php,omitempty" xml:"Php,omitempty"`
	// The mount path for PHP application monitoring. You must ensure that the PHP server loads the configuration file from this path.
	//
	// You do not need to focus on the configuration content because SAE automatically renders the correct configuration file.
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
	// The mount path for the PHP application startup configuration. You must ensure that the PHP server uses this configuration file to start.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// The script to execute after the container starts. A script is executed immediately after the container is created. Format: `{"exec":{"command":["cat","/etc/group"]}}`.
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script to execute before the container stops. A script is executed before the container is deleted. Format: `{"exec":{"command":["cat","/etc/group"]}}`.
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The programming language used to create the application. Valid values:
	//
	// - **java**: Java.
	//
	// - **php**: PHP.
	//
	// - **other**: Multiple languages, such as Python, C++, Go, .NET, and Node.js.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// Specifies whether to enable Kubernetes Service registration and discovery.
	//
	// example:
	//
	// {     "serviceName": "bwm-poc-sc-gateway-cn-beijing-front",     "namespaceId": "cn-beijing:front",     "portAndProtocol": {         "18012": "TCP"     },     "portProtocols": [         {             "port": "18012",             "protocol": "TCP"         }     ],     "enable": true }
	PvtzDiscovery *string `json:"PvtzDiscovery,omitempty" xml:"PvtzDiscovery,omitempty"`
	// The Python environment. Python 3.9.15 is supported.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// The module dependencies for custom installation. By default, the dependencies defined in the requirements.txt file in the root folder are installed. If no dependencies are configured or custom packages are used, you can specify the dependencies to install.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// The RASP configuration items.
	RaspConfig []*DescribeApplicationConfigResponseBodyDataRaspConfig `json:"RaspConfig,omitempty" xml:"RaspConfig,omitempty" type:"Repeated"`
	// The application startup status check. Containers that fail the health check multiple times are stopped and restarted. Traffic from SLB is not routed to containers that fail the health check. The **exec**, **httpGet**, and **tcpSocket*	- methods are supported. For specific examples, see the **Liveness*	- parameter.
	//
	// > You can select only one method for the health check.
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
	// The configuration for collecting logs to Simple Log Service (SLS).
	//
	// - To use SLS resources with automatic creation by Serverless App Engine (SAE): `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - To use custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// Metric description:
	//
	// - **projectName**: The name of the SLS project.
	//
	// - **logDir**: The log path character string.
	//
	// - **logType**: The log type. **stdout*	- indicates container standard output logs. You can set only one stdout log. If the Settings are not specified, file logs are collected.
	//
	// - **logstoreName**: The name of the SLS Logstore.
	//
	// - **logtailName**: The name of the SLS Logtail. If this parameter is not specified, a new Logtail is created.
	//
	// If the SLS collection configuration remains unchanged during multiple deployments, you do not need to set this parameter. That is, you do not need to include the **SlsConfigs*	- field in the request. If you no longer need the SLS collection feature, you must set the value of this field to an empty string in the request. That is, set the **SlsConfigs*	- field to "".
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// sls log env tags
	//
	// example:
	//
	// []
	SlsLogEnvTags *string `json:"SlsLogEnvTags,omitempty" xml:"SlsLogEnvTags,omitempty"`
	// The application startup probe.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"/bin/sh\\",\\"-c\\",\\"#!Note: If microservice config is enabled, the application will be automatically injected with the prestop configuration for lossless offline. If you delete this prestop configuration, lossless offline will not be effective.\\\\n echo stop > /tmp/prestop; /home/admin/.tools/curl http://127.0.0.1:54199/offline; sleep 30\\"]}}
	StartupProbe *string `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	// The configuration for service registration and discovery and end-to-end canary release based on Kubernetes Services.
	//
	// example:
	//
	// {\\"enable\\":\\"false\\",\\"namespaceId\\":\\"cn-beijing:test\\",\\"portAndProtocol\\":{\\"2000:TCP\\":\\"18081\\"},\\"portProtocols\\":[{\\"port\\":2000,\\"protocol\\":\\"TCP\\",\\"targetPort\\":18081}],\\"pvtzDiscoveryName\\":\\"cn-beijing-1421801774382676\\",\\"serviceId\\":\\"3513\\",\\"serviceName\\":\\"demo-gray.test\\"}
	SwimlanePvtzDiscovery *string `json:"SwimlanePvtzDiscovery,omitempty" xml:"SwimlanePvtzDiscovery,omitempty"`
	// The tag information.
	Tags []*DescribeApplicationConfigResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period for graceful shutdown. Default value: 30. Unit: seconds. Valid values: 1 to 300.
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
	// The Tomcat file configuration. Set the value to "" or "{}" to delete the configuration. Valid values:
	//
	// - **port**: The port number. Valid values: 1024 to 65535. Ports smaller than 1024 require root permissions. Because the container is configured with admin permissions, specify a port greater than 1024. Default value: 8080.
	//
	// - **contextPath**: The access path. Default value: the root directory "/".
	//
	// - **maxThreads**: The maximum number of connections in the connection pool. Default value: 400.
	//
	// - **uriEncoding**: The Tomcat encoding format. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. Default value: **ISO-8859-1**.
	//
	// - **useBodyEncoding**: Specifies whether to use **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	// The deployment policy. When the minimum number of surviving instances is 1, the value of the **UpdateStrategy*	- field is "". When the minimum number of surviving instances is greater than 1, examples are as follows:
	//
	// - Grayscale 1 instance, followed by 2 batches, automatic batching, and a batch interval of 1 minute: `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}`
	//
	//
	//
	// - Grayscale 1 instance, followed by 2 batches, and manual batching: `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"manual"},"grayUpdate":{"gray":1}}`
	//
	// - 2 batches, automatic batching, and a batch interval of 0 minutes: `{"type":"BatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":0}}`
	//
	// Metric description:
	//
	// - **type**: The publish policy type. Valid values: **GrayBatchUpdate*	- for grayscale publish and **BatchUpdate*	- for batch publish.
	//
	// - **batchUpdate**: The batch publish policy.
	//
	//     - **batch**: The number of publish batches.
	//
	//     - **releaseType**: The processing method between batches. Valid values: **auto*	- for automatic and **manual*	- for manual.
	//
	//     - **batchWaitTime**: The deployment interval between batches. Unit: seconds.
	//
	// - **grayUpdate**: The remaining batches after the grayscale publish. This parameter is required when **type*	- is set to **GrayBatchUpdate**.
	//
	// example:
	//
	// {"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}
	UpdateStrategy *string `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2ze559r1z1bpwqxwp****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2ze0i263cnn311nvj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The startup options for WAR package applications. Default application startup command: `java $JAVA_OPTS $CATALINA_OPTS -Options org.apache.catalina.startup.Bootstrap "$@" start`.
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
	// The ConfigMap ID.
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
	//
	// example:
	//
	// /usr/local/tomcat/webapps
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The ephemeral storage name.
	//
	// example:
	//
	// workdir
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
	// The image startup command. The command must be an executable object that exists in the container. Example:
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
	// The image startup command parameters. The parameters required by the **Command*	- startup command. Format:
	//
	// `["a","b"]`
	//
	// In the preceding example, `CommandArgs=["abc", ">", "file0"]`, where `["abc", ">", "file0"]` must be converted to a string in JSON array format. If this parameter is not required, leave it empty.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The ConfigMap information.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The shared ephemeral storage.
	EmptyDirDesc []*DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// The container environment variable parameters. You can customize the parameters or reference configuration items. To reference configuration items, create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//     - **name**: The environment variable name.
	//
	//     - **value**: The environment variable value. The priority is higher than valueFrom.
	//
	// - Reference configuration items (valueFrom)
	//
	//     - **name**: The environment variable name. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<Configuration item name>`, for example, `sae-sys-configmap-all-test1`.
	//
	//     - **valueFrom**: The environment variable reference. Set the value to `configMapRef`.
	//
	//     - **configMapId**: The configuration item ID.
	//
	//     - **key**: The key. If all key-value pairs are referenced, do not set this field.
	//
	// - Reference Secrets (valueFrom)
	//
	//     - **name**: The environment variable name. You can reference a single key or all keys. To reference all keys, enter `sae-sys-secret-all-<Secret name>`, for example, `sae-sys-secret-all-test1`.
	//
	//     - **valueFrom**: The environment variable reference. Set the value to `secretRef`.
	//
	//     - **secretId**: The Secret ID.
	//
	//     - **key**: The key. If all key-value pairs are referenced, do not set this field.
	//
	// example:
	//
	// [{"name":"TEST_ENV_KEY","value":"TEST_ENV_VAR"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The image address used by the init container.
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
	// The ConfigMap ID.
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
	//
	// example:
	//
	// /usr/local/tomcat/webapps
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The ephemeral storage name.
	//
	// example:
	//
	// workdir
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
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path.
	//
	// example:
	//
	// /tmp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The Secret instance ID.
	//
	// example:
	//
	// 16
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The Secret instance name.
	//
	// example:
	//
	// registry-auth
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
	// The relative NAS file directory.
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
	// Specifies whether the container path has read-only permissions for the mounted directory resources. Valid values:
	//
	// - **true**: Read-only permissions.
	//
	// - **false**: Read and write permissions.
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
	// Specifies whether to enable RASP.
	EnableRasp *bool `json:"EnableRasp,omitempty" xml:"EnableRasp,omitempty"`
	// The RASP application key.
	//
	// example:
	//
	// key-1
	RaspAppKey *string `json:"RaspAppKey,omitempty" xml:"RaspAppKey,omitempty"`
	// The RASP application name.
	//
	// example:
	//
	// sae-test
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
	// The key whose data value is Base64-encoded.
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
	// The instance ID of the queried Secret.
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
	// The instance ID of the Container Registry Enterprise instance. This parameter is required when **ImageUrl*	- is set to a Container Registry Enterprise instance.
	//
	// example:
	//
	// cri-fhzlneorxala66ip
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The image startup command. The command must be an executable object that exists in the container. Example:
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
	// The image startup command parameters. The parameters required by the **Command*	- startup command. Format:
	//
	// `["a","b"]`
	//
	// In the preceding example, `CommandArgs=["abc", ">", "file0"]`, where `["abc", ">", "file0"]` must be converted to a string in JSON array format. If this parameter is not required, leave it empty.
	//
	// example:
	//
	// [\\"-c\\",\\"echo \\\\\\"test\\\\\\" > /home/nas/test.log && sleep 10000000s\\"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The ConfigMap mount description. Uses the ConfigMap created on the namespace configuration item page to inject configuration information into the container. The metric description is as follows:
	//
	// - **configMapId**: The ConfigMap instance ID. You can invoke the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the ID.
	//
	// - **key**: The key value.
	//
	// > You can pass the `sae-sys-configmap-all` parameter to mount all keys.
	//
	// - **mountPath**: The mount path.
	//
	// - **ConfigMapName**: The ConfigMap name.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The maximum CPU resources that the sidecar container can use from the main container.
	//
	// example:
	//
	// 500
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The shared ephemeral storage. Sets the ephemeral storage directory and mounts it to the main container and the sidecar container.
	EmptyDirDesc []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// The container environment variable parameters. You can define custom variables or reference configuration items. To reference configuration items, you must first create a ConfigMap instance. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//     - **name**: The name of the environment variable.
	//
	//     - **value**: The value of the environment variable. This parameter has a higher priority than valueFrom.
	//
	// - Reference configuration items (valueFrom)
	//
	//     - **name**: The name of the environment variable. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<ConfigMap name>`, such as `sae-sys-configmap-all-test1`.
	//
	//     - **valueFrom**: The reference of the environment variable. Set the value to `configMapRef`.
	//
	//         - **configMapId**: The ID of the ConfigMap.
	//
	//         - **key**: The key. If you reference all key-value pairs, do not set this field.
	//
	// example:
	//
	// [{\\"name\\":\\"k1\\",\\"value\\":\\"c8e3a815-e5d3-4adf-abb3-98b106a607c4\\"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The image address.
	//
	// example:
	//
	// registry.cn-beijing.aliyuncs.com/sae-dev-test/nginx:stable
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The container health check.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The upper limit of memory resources that the sidecar container can use from the main container.
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
	// The script to execute after the container starts.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script to execute before the container stops.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The application startup status check.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The Secret mount description.
	SecretMountDesc []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigSecretMountDesc `json:"SecretMountDesc,omitempty" xml:"SecretMountDesc,omitempty" type:"Repeated"`
	// The configuration for collecting sidecar logs to Simple Log Service (SLS). The structure is aligned with SlsConfigs.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
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

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) GetSlsConfigs() *string {
	return s.SlsConfigs
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

func (s *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig) SetSlsConfigs(v string) *DescribeApplicationConfigResponseBodyDataSidecarContainersConfig {
	s.SlsConfigs = &v
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
	// The ephemeral storage name.
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
	// The key whose data value is Base64-encoded.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path.
	//
	// example:
	//
	// /tmp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The Secret instance ID.
	//
	// example:
	//
	// 16
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The Secret instance name.
	//
	// example:
	//
	// registry-auth
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

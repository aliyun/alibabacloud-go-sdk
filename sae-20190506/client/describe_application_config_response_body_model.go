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
	// - **2xx**: The request is successful.
	//
	// - **3xx**: The request is redirected.
	//
	// - **4xx**: The request is invalid.
	//
	// - **5xx**: A server error occurs.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the application.
	Data *DescribeApplicationConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is not returned if the request is successful.
	//
	// - If the request fails, this parameter is returned. For more information, see the "Error codes" section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information that is returned.
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
	// Indicates whether the application configuration was retrieved. Valid values:
	//
	// - **true**: The configuration was retrieved.
	//
	// - **false**: The configuration failed to be retrieved.
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
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is required to pull images across accounts. For more information, see [Pull images across Alibaba Cloud accounts](https://help.aliyun.com/document_detail/190675.html) and [Use a RAM role to grant permissions across Alibaba Cloud accounts](https://help.aliyun.com/document_detail/223585.html).
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
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The agent version.
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The configuration of the Application Load Balancer (ALB) gateway readiness gate.
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
	// The type of the SAE application.
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
	// Specifies whether to bind an elastic IP address (EIP). Valid values:
	//
	// - **true**: Bind an EIP.
	//
	// - **false**: Do not bind an EIP.
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
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The Cloud Monitor service ID.
	CmsServiceId *string `json:"CmsServiceId,omitempty" xml:"CmsServiceId,omitempty"`
	// The startup command of the image. The command must be an executable object that exists in the container. Example:
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
	// The arguments of the image startup command. The arguments are required by the **Command*	- parameter. Format:
	//
	// `["a","b"]`
	//
	// In the example of the **Command*	- parameter, `CommandArgs=["abc", ">", "file0"]`. The value `["abc", ">", "file0"]` must be converted into a string in the JSON array format. If this parameter is not required, you do not need to specify it.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The configurations of the ConfigMap.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The CPU required by each instance. Unit: millicores. The value cannot be 0. The following specifications are supported:
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
	// The custom mapping between a domain name and an IP address in the container. Valid values:
	//
	// - **hostName**: The domain name or hostname.
	//
	// - **ip**: The IP address.
	//
	// example:
	//
	// [{"hostName":"test.host.name","ip":"0.0.0.0"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The type of the custom image. If you do not use a custom image, set this parameter to an empty string. Valid values:
	//
	// - internet: a public image
	//
	// - intranet: a private image
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
	// The disk storage size. Unit: GB.
	//
	// example:
	//
	// 20
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The version of the .NET framework:
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
	// The version of the application runtime environment in the High-Speed Service Framework (HSF), such as an Ali-Tomcat container.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// The shared temporary storage.
	EmptyDirDesc []*DescribeApplicationConfigResponseBodyDataEmptyDirDesc `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// Specifies whether to enable Application High Availability Service (AHAS). Valid values:
	//
	// - **true**: Enable AHAS.
	//
	// - **false**: Disable AHAS.
	//
	// example:
	//
	// true
	EnableAhas *string `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	// Specifies whether to enable the CPU burst feature:
	//
	// - true: Enable
	//
	// - false: Disable
	//
	// example:
	//
	// true
	EnableCpuBurst *string `json:"EnableCpuBurst,omitempty" xml:"EnableCpuBurst,omitempty"`
	// Specifies whether to enable the canary release rule. This rule is applicable only to applications that use the Spring Cloud and Dubbo frameworks. Valid values:
	//
	// - **true**: Enable the canary release rule.
	//
	// - **false**: Disable the canary release rule.
	//
	// example:
	//
	// false
	EnableGreyTagRoute *bool `json:"EnableGreyTagRoute,omitempty" xml:"EnableGreyTagRoute,omitempty"`
	// Specifies whether to enable the idle mode:
	//
	// - true: Enable
	//
	// - false: Disable
	//
	// example:
	//
	// false
	EnableIdle *bool `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	// Specifies whether to reuse the agent version configuration of the namespace.
	EnableNamespaceAgentVersion *bool `json:"EnableNamespaceAgentVersion,omitempty" xml:"EnableNamespaceAgentVersion,omitempty"`
	// Specifies whether to enable the new ARMS feature:
	//
	// - true: Enable
	//
	// - false: Disable
	//
	// example:
	//
	// false
	EnableNewArms *bool `json:"EnableNewArms,omitempty" xml:"EnableNewArms,omitempty"`
	// Specifies whether to enable custom metric collection for Prometheus.
	EnablePrometheus *bool `json:"EnablePrometheus,omitempty" xml:"EnablePrometheus,omitempty"`
	// The environment variables for the container. You can customize environment variables or reference a ConfigMap. To reference a ConfigMap, you must first create a ConfigMap. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//   - **name**: The name of the environment variable.
	//
	//   - **value**: The value of the environment variable.
	//
	// - Reference a configuration item
	//
	//   - **name**: The name of the environment variable. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<ConfigMap name>`, for example, `sae-sys-configmap-all-test1`.
	//
	//   - **valueFrom**: The reference of the environment variable. Set the value to `configMapRef`.
	//
	//   - **configMapId**: The ID of the ConfigMap.
	//
	//   - **key**: The key. If you want to reference all keys, do not specify this parameter.
	//
	// example:
	//
	// [{"name":"TEST_ENV_KEY","value":"TEST_ENV_VAR"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The number of GPUs.
	GpuCount *string `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The GPU card type.
	GpuType               *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	HeadlessPvtzDiscovery *string `json:"HeadlessPvtzDiscovery,omitempty" xml:"HeadlessPvtzDiscovery,omitempty"`
	Html                  *string `json:"Html,omitempty" xml:"Html,omitempty"`
	IdleHour              *string `json:"IdleHour,omitempty" xml:"IdleHour,omitempty"`
	// The ID of the secret.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The URL of the image. This parameter is required when **Package Type*	- is set to **Image**.
	//
	// example:
	//
	// docker.io/library/nginx:1.14.2
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The configurations of the init container.
	InitContainersConfig []*DescribeApplicationConfigResponseBodyDataInitContainersConfig `json:"InitContainersConfig,omitempty" xml:"InitContainersConfig,omitempty" type:"Repeated"`
	// Specifies whether the application is a stateful application.
	IsStateful *bool `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The arguments for starting the JAR package. The default startup command is: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// start
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The options for starting the JAR package. The default startup command is: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// -Dtest=true
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The version of the Java Development Kit (JDK) that the deployment package requires. The following versions are supported:
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
	// The configurations for collecting logs to Kafka. Valid values:
	//
	// - **kafkaEndpoint**: The endpoint of the Kafka API.
	//
	// - **kafkaInstanceId**: The ID of the Kafka instance.
	//
	// - **kafkaConfigs**: The configurations of one or more logs. For more information about the example and parameters, see the **kafkaConfigs*	- request parameter in this topic.
	//
	// example:
	//
	// {"kafkaEndpoint":"10.0.X.XXX:XXXX,10.0.X.XXX:XXXX,10.0.X.XXX:XXXX","kafkaInstanceId":"alikafka_pre-cn-7pp2l8kr****","kafkaConfigs":[{"logType":"file_log","logDir":"/tmp/a.log","kafkaTopic":"test2"},{"logType":"stdout","logDir":"","kafkaTopic":"test"}]}
	KafkaConfigs *string `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	// The labels.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The liveness probe of the container. A container that fails the health check is shut down and restored. The following methods are supported:
	//
	// - **exec**: example: `{"exec":{"command":["sh","-c","cat/home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}`
	//
	// - **httpGet**: example:`{"httpGet":{"path":"/","port":18091,"scheme":"HTTP","isContainKeyWord":true,"keyWord":"SAE"},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// - **tcpSocket**: example:`{"tcpSocket":{"port":18091},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// > You can use only one method for a health check.
	//
	// The parameters are described as follows:
	//
	// - **exec.command**: The command that is used for the health check.
	//
	// - **httpGet.path**: The request path.
	//
	// - **httpGet.scheme**: **HTTP*	- or **HTTPS**.
	//
	// - **httpGet.isContainKeyWord**: Specifies whether the response must contain a keyword. A value of **true*	- indicates that the response must contain the keyword. A value of **false*	- indicates that the response does not need to contain the keyword. If you do not specify this parameter, this advanced feature is not used.
	//
	// - **httpGet.keyWord**: The custom keyword. This parameter is required if you set the **isContainKeyWord*	- parameter.
	//
	// - **tcpSocket.port**: The port that is used for the TCP connection check.
	//
	// - **initialDelaySeconds**: The delay for the health check. Default value: 10. Unit: seconds.
	//
	// - **periodSeconds**: The interval for the health check. Default value: 30. Unit: seconds.
	//
	// - **timeoutSeconds**: The timeout period for the health check. Default value: 1. Unit: seconds. If you set this parameter to 0 or do not specify this parameter, the default value is used.
	//
	// example:
	//
	// {"exec":{"command":["curl http://localhost:8080"]},"initialDelaySeconds":20,"timeoutSeconds":3}
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The Loki configurations.
	LokiConfigs *string `json:"LokiConfigs,omitempty" xml:"LokiConfigs,omitempty"`
	// The maximum surge instance ratio.
	MaxSurgeInstanceRatio *int32 `json:"MaxSurgeInstanceRatio,omitempty" xml:"MaxSurgeInstanceRatio,omitempty"`
	// The maximum number of surge instances.
	MaxSurgeInstances *int32 `json:"MaxSurgeInstances,omitempty" xml:"MaxSurgeInstances,omitempty"`
	// The memory required by each instance. Unit: MB. The value cannot be 0. This parameter corresponds to the \\`Cpu\\` parameter. The following specifications are supported:
	//
	// - **1024**: corresponds to 500 millicores and 1,000 millicores.
	//
	// - **2048**: corresponds to 500, 1,000, and 2,000 millicores.
	//
	// - **4096**: corresponds to 1,000, 2,000, and 4,000 millicores.
	//
	// - **8192**: corresponds to 2,000, 4,000, and 8,000 millicores.
	//
	// - **12288**: corresponds to 12,000 millicores.
	//
	// - **16384**: corresponds to 4,000, 8,000, and 16,000 millicores.
	//
	// - **24576**: corresponds to 12,000 millicores.
	//
	// - **32768**: corresponds to 16,000 millicores.
	//
	// - **65536**: corresponds to 8,000, 16,000, and 32,000 millicores.
	//
	// - **131072**: corresponds to 32,000 millicores.
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The Nacos registry. Valid values:
	//
	// - **0**: SAE built-in Nacos.
	//
	// - **1**: User-created Nacos.
	//
	// - **2**: MSE Nacos.
	//
	// example:
	//
	// "0"
	MicroRegistration *string `json:"MicroRegistration,omitempty" xml:"MicroRegistration,omitempty"`
	// The configuration of the registry. This parameter takes effect only when the registry is MSE Nacos Enterprise Edition.
	//
	// example:
	//
	// {\\"instanceId\\":\\"mse-cn-1ls43******\\",\\"namespace\\":\\"62ee12fb-c279-4da4-be96-21**********\\"}
	MicroRegistrationConfig *string `json:"MicroRegistrationConfig,omitempty" xml:"MicroRegistrationConfig,omitempty"`
	// The configurations of microservice governance.
	//
	// - Specifies whether to enable microservice governance (enable):
	//
	//   - true: Enable
	//
	//   - false: Disable
	//
	// - The configuration of graceful start and shutdown (mseLosslessRule):
	//
	//   - delayTime: The delay time.
	//
	//   - enable: Specifies whether to enable graceful start. true indicates that graceful start is enabled. false indicates that graceful start is not enabled.
	//
	//   - notice: Specifies whether to enable notifications. true indicates that notifications are enabled. false indicates that notifications are not enabled.
	//
	//   - warmupTime: The warm-up duration for a small amount of traffic. Unit: seconds.
	//
	// example:
	//
	// {\\"Enable\\":true,\\"MseLosslessRule\\":{\\"enable\\":true,\\"notice\\":true,\\"delayTime\\":10,\\"warmupTime\\":120,\\"funcType\\":2,\\"aligned\\":false,\\"related\\":false,\\"lossLessDetail\\":false}}
	MicroserviceEngineConfig *string `json:"MicroserviceEngineConfig,omitempty" xml:"MicroserviceEngineConfig,omitempty"`
	// The percentage of the minimum number of ready instances. Valid values:
	//
	// - -1: The default value. This value indicates that the percentage is not used. If you do not specify this parameter, the system uses the default value **-1**.
	//
	// - **0 to 100**: The percentage of the minimum number of ready instances. The value is rounded up. For example, if you set this parameter to **50**%, and you have five instances, the minimum number of ready instances is 3.
	//
	// > If you specify both \\`MinReadyInstances\\` and **MinReadyInstanceRatio**, and the value of **MinReadyInstanceRatio*	- is not **-1**, the value of **MinReadyInstanceRatio*	- takes precedence. For example, if **MinReadyInstances*	- is set to **5*	- and **MinReadyInstanceRatio*	- is set to **50**, the minimum number of ready instances is calculated based on the value of **MinReadyInstanceRatio**.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of ready instances. Valid values:
	//
	// - If you set this parameter to **0**, the application may be interrupted during an upgrade.
	//
	// - If you set this parameter to -1, the system uses a recommended value, which is 25% of the total number of instances. For example, if you have five instances, the minimum number of ready instances is 2 after the value is rounded up.
	//
	// > We recommend that you set the minimum number of ready instances to a value of 1 or greater to ensure that the application is not interrupted during a rolling update.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// The mount description.
	MountDesc []*DescribeApplicationConfigResponseBodyDataMountDesc `json:"MountDesc,omitempty" xml:"MountDesc,omitempty" type:"Repeated"`
	// The mount target of the Apsara File Storage NAS file system in the application VPC. If you do not change the NAS configuration during a deployment, you do not need to specify this parameter. If you want to clear the NAS configuration, set this parameter to an empty string ("").
	//
	// example:
	//
	// example.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The ID of the application in Microservices Engine (MSE).
	//
	// example:
	//
	// xxxxxxx@xxxxx
	MseApplicationId *string `json:"MseApplicationId,omitempty" xml:"MseApplicationId,omitempty"`
	// The name of the application after the SAE service is registered with MSE.
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
	// The configurations for mounting a NAS file system.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// The ID of the NAS file system.
	//
	// example:
	//
	// AKSN****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The application version.
	//
	// - lite: Lightweight Edition
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
	// > You must create an OpenID Connect (OIDC) identity provider (IdP) and a RAM role for the IdP in the same region beforehand. For more information, see [Create an OIDC IdP](https://help.aliyun.com/document_detail/2331022.html) and [Create a RAM role for a trusted IdP](https://help.aliyun.com/document_detail/2331016.html).
	//
	// example:
	//
	// sae-test
	OidcRoleName *string `json:"OidcRoleName,omitempty" xml:"OidcRoleName,omitempty"`
	// The AccessKey ID that is used to read data from and write data to Object Storage Service (OSS).
	//
	// example:
	//
	// xxxxxx
	OssAkId *string `json:"OssAkId,omitempty" xml:"OssAkId,omitempty"`
	// The AccessKey secret that is used to read data from and write data to OSS.
	//
	// example:
	//
	// xxxxxx
	OssAkSecret *string `json:"OssAkSecret,omitempty" xml:"OssAkSecret,omitempty"`
	// The description of the OSS mount.
	OssMountDescs []*DescribeApplicationConfigResponseBodyDataOssMountDescs `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty" type:"Repeated"`
	// The type of the application package. Valid values:
	//
	// - If you deploy the application using Java, valid values are **FatJar**, **War**, and **Image**.
	//
	// - If you deploy the application using PHP, the following values are supported:
	//
	//   - **PhpZip**
	//
	//   - **IMAGE_PHP_5_4**
	//
	//   - **IMAGE_PHP_5_4_ALPINE**
	//
	//   - **IMAGE_PHP_5_5**
	//
	//   - **IMAGE_PHP_5_5_ALPINE**
	//
	//   - **IMAGE_PHP_5_6**
	//
	//   - **IMAGE_PHP_5_6_ALPINE**
	//
	//   - **IMAGE_PHP_7_0**
	//
	//   - **IMAGE_PHP_7_0_ALPINE**
	//
	//   - **IMAGE_PHP_7_1**
	//
	//   - **IMAGE_PHP_7_1_ALPINE**
	//
	//   - **IMAGE_PHP_7_2**
	//
	//   - **IMAGE_PHP_7_2_ALPINE**
	//
	//   - **IMAGE_PHP_7_3**
	//
	//   - **IMAGE_PHP_7_3_ALPINE**
	//
	// example:
	//
	// War
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The URL of the deployment package. If you upload the deployment package using SAE, note the following:
	//
	// - You cannot download the package from this URL. Call the \\`GetPackageVersionAccessableUrl\\` operation to obtain a download URL that is valid for 10 minutes.
	//
	// - SAE stores the package for a maximum of 90 days. After 90 days, the URL is not returned and you cannot download the package.
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package. This parameter is required when **Package Type*	- is set to **FatJar*	- or **War**.
	//
	// example:
	//
	// 1.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The PHP version required for the deployment package. This parameter is not supported for images.
	//
	// example:
	//
	// PHP-FPM 7.0
	Php *string `json:"Php,omitempty" xml:"Php,omitempty"`
	// The mount path of the Application Real-Time Monitoring Service (ARMS) configuration file for a PHP application. Make sure that the PHP server loads the configuration file from this path.
	//
	// SAE automatically renders the correct configuration file. You do not need to manage the content of the configuration file.
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
	// The mount path of the PHP application startup configuration file. Make sure that the PHP server uses this configuration file to start.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// The script that runs after the container starts. The script runs immediately after the container is created. Example: `{"exec":{"command":["cat","/etc/group"]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script that runs before the container is stopped. The script runs before the container is deleted. Example: `{"exec":{"command":["cat","/etc/group"]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The programming language of the application. Valid values:
	//
	// - **java**: Java
	//
	// - **php**: PHP
	//
	// - **other**: other languages, such as Python, C++, Go, .NET, and Node.js.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// Enables service registration and discovery for a Kubernetes Service.
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
	// The dependencies for custom installation modules. By default, the dependencies that are defined in the requirements.txt file in the root directory are installed. If no software package is configured or a custom software package is used, you can specify the dependencies to be installed.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// The readiness probe of the application. A container that fails the health check multiple times is shut down and restarted. A container that fails the health check does not receive traffic from a Server Load Balancer (SLB) instance. You can perform the health check using the **exec**, **httpGet**, or **tcpSocket*	- method. For more information, see the **Liveness*	- parameter.
	//
	// > You can use only one method for a health check.
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
	// The description of the secret that you want to mount.
	SecretMountDesc []*DescribeApplicationConfigResponseBodyDataSecretMountDesc `json:"SecretMountDesc,omitempty" xml:"SecretMountDesc,omitempty" type:"Repeated"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The canary release tags of the application configuration.
	ServiceTags map[string]*string `json:"ServiceTags,omitempty" xml:"ServiceTags,omitempty"`
	// The configurations of the sidecar container.
	SidecarContainersConfig []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfig `json:"SidecarContainersConfig,omitempty" xml:"SidecarContainersConfig,omitempty" type:"Repeated"`
	// The configurations for collecting logs to Simple Log Service (SLS).
	//
	// - To use an SLS resource that is automatically created by SAE: `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - To use a custom SLS resource: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// The parameters are described as follows:
	//
	// - **projectName**: The name of the SLS project.
	//
	// - **logDir**: The log path.
	//
	// - **logType**: The log type. **stdout*	- indicates the standard output log of the container. You can specify only one log of this type. If you do not specify this parameter, file logs are collected.
	//
	// - **logstoreName**: The name of the Logstore in SLS.
	//
	// - **logtailName**: The name of the Logtail configuration in SLS. If you do not specify this parameter, a new Logtail configuration is created.
	//
	// If you do not change the log collection configuration during a deployment, you do not need to specify this parameter. If you no longer need to use the log collection feature, set this parameter to an empty string ("") in the request.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// The environment tags for SLS logs.
	SlsLogEnvTags *string `json:"SlsLogEnvTags,omitempty" xml:"SlsLogEnvTags,omitempty"`
	// The startup probe of the application.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"/bin/sh\\",\\"-c\\",\\"#!Note: If microservice config is enabled, the application will be automatically injected with the prestop configuration for lossless offline. If you delete this prestop configuration, lossless offline will not be effective.\\\\n echo stop > /tmp/prestop; /home/admin/.tools/curl http://127.0.0.1:54199/offline; sleep 30\\"]}}
	StartupProbe *string `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	// The configuration for service registration and discovery based on a Kubernetes Service and for end-to-end canary release.
	//
	// example:
	//
	// {\\"enable\\":\\"false\\",\\"namespaceId\\":\\"cn-beijing:test\\",\\"portAndProtocol\\":{\\"2000:TCP\\":\\"18081\\"},\\"portProtocols\\":[{\\"port\\":2000,\\"protocol\\":\\"TCP\\",\\"targetPort\\":18081}],\\"pvtzDiscoveryName\\":\\"cn-beijing-1421801774382676\\",\\"serviceId\\":\\"3513\\",\\"serviceName\\":\\"demo-gray.test\\"}
	SwimlanePvtzDiscovery *string `json:"SwimlanePvtzDiscovery,omitempty" xml:"SwimlanePvtzDiscovery,omitempty"`
	// The tags.
	Tags []*DescribeApplicationConfigResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period for a graceful shutdown. Default value: 30. Unit: seconds. The value can range from 1 to 300.
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
	// The Tomcat configuration. To delete the configuration, set this parameter to "" or "{}".
	//
	// - **port**: The port number. The port number can range from 1024 to 65535. A port number smaller than 1024 requires the root permission to be operated. Because the container is configured with the administrator permission, specify a port number that is greater than 1024. If you do not configure this parameter, the default port 8080 is used.
	//
	// - **contextPath**: The access path. The default value is the root directory "/".
	//
	// - **maxThreads**: The maximum number of connections in the connection pool. Default value: 400.
	//
	// - **uriEncoding**: The URI encoding scheme of Tomcat. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. If you do not set this parameter, the default value **ISO-8859-1*	- is used.
	//
	// - **useBodyEncoding**: Specifies whether to use **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	// The deployment policy. If the minimum number of ready instances is 1, the value of the **UpdateStrategy*	- parameter is "". If the minimum number of ready instances is greater than 1, see the following examples:
	//
	// - Canary release of one instance, phased release in two batches, automatic batching, and a 1-minute interval between batches: `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}`
	//
	// - Canary release of one instance and phased release in two batches with manual batching: `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"manual"},"grayUpdate":{"gray":1}}`
	//
	// - Phased release in two batches, automatic batching, and a 0-minute interval between batches: `{"type":"BatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":0}}`
	//
	// The parameters are described as follows:
	//
	// - **type**: The type of the release policy. Valid values: **GrayBatchUpdate*	- (canary release) and **BatchUpdate*	- (phased release).
	//
	// - **batchUpdate**: The phased release policy.
	//
	//   - **batch**: The number of release batches.
	//
	//   - **releaseType**: The processing method for batches. Valid values: **auto*	- and **manual**.
	//
	//   - **batchWaitTime**: The interval between batches. Unit: seconds.
	//
	// - **grayUpdate**: The number of batches for the remaining instances after the canary release. This parameter is required when **type*	- is set to **GrayBatchUpdate**.
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
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2ze0i263cnn311nvj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The options for starting the WAR package. The default startup command is: `java $JAVA_OPTS $CATALINA_OPTS -Options org.apache.catalina.startup.Bootstrap "$@" start`.
	//
	// example:
	//
	// custom-option
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The Tomcat version that the deployment package requires. The following versions are supported:
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
	// The key of the key-value pair.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path of the container.
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
	// The name of the temporary storage.
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
	// The startup command of the image. The command must be an executable object that exists in the container. Example:
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
	// The arguments of the image startup command. The arguments are required by the preceding startup command **Command**. Format:
	//
	// `["a","b"]`
	//
	// In the preceding example, `CommandArgs=["abc", ">", "file0"]`. The value `["abc", ">", "file0"]` must be converted into a string in the JSON array format. If this parameter is not required, you do not need to specify it.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The configurations of the ConfigMap.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataInitContainersConfigConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The shared temporary storage.
	EmptyDirDesc []*DescribeApplicationConfigResponseBodyDataInitContainersConfigEmptyDirDesc `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// The environment variables of the container. You can customize environment variables or reference a ConfigMap. To reference a ConfigMap, you must first create a ConfigMap instance. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//   - **name**: The name of the environment variable.
	//
	//   - **value**: The value of the environment variable. This parameter takes precedence over valueFrom.
	//
	// - Reference a configuration item (valueFrom)
	//
	//   - **name**: The name of the environment variable. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<ConfigMap name>`, for example, `sae-sys-configmap-all-test1`.
	//
	//   - **valueFrom**: The reference of the environment variable. Set the value to `configMapRef`.
	//
	//   - **configMapId**: The ID of the ConfigMap.
	//
	//   - **key**: The key. If you want to reference all keys, do not specify this parameter.
	//
	// - Reference a secret (valueFrom)
	//
	//   - **name**: The name of the environment variable. You can reference a single key or all keys. To reference all keys, enter `sae-sys-secret-all-<Secret name>`, for example, `sae-sys-secret-all-test1`.
	//
	//   - **valueFrom**: The reference of the environment variable. Set the value to `secretRef`.
	//
	//   - **secretId**: The ID of the secret.
	//
	//   - **key**: The key. If you want to reference all keys, do not specify this parameter.
	//
	// example:
	//
	// [{"name":"TEST_ENV_KEY","value":"TEST_ENV_VAR"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The URL of the image that is used by the init container.
	//
	// example:
	//
	// registry.cn-shenzhen.aliyuncs.com/sae-serverless-demo/sae-demo:microservice-java-provider-v1.0
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The name of the init container.
	//
	// example:
	//
	// init-container
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the secret that you want to mount.
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
	// The key of the key-value pair.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path of the container.
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
	// The path on which the volume is mounted in the container.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The name of the temporary storage.
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
	// The ID of the secret instance.
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The name of the secret instance.
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
	// The mount path of the container.
	//
	// example:
	//
	// /tmp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The path of the NAS file system.
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
	// The directory or object that you created in OSS. An error occurs if the mount directory does not exist.
	//
	// example:
	//
	// data/user.data
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	// The path of the container in SAE. If the path exists, the path is overwritten. If the path does not exist, a new path is created.
	//
	// example:
	//
	// /usr/data/user.data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// Specifies whether the container has the read-only permission on the mount directory resources. Valid values:
	//
	// - **true**: The read-only permission.
	//
	// - **false**: The read and write permissions.
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
	// The key of the data value that is encoded in Base64.
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
	// The ID of the queried secret instance.
	//
	// example:
	//
	// 520
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The name of the secret instance.
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
	// The ID of the Container Registry Enterprise Edition instance. This parameter is required if **ImageUrl*	- is set to an image in Container Registry Enterprise Edition.
	//
	// example:
	//
	// cri-fhzlneorxala66ip
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The startup command of the image. The command must be an executable object that exists in the container. Example:
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
	// The arguments of the image startup command. The arguments are required by the preceding startup command **Command**. Format:
	//
	// `["a","b"]`
	//
	// In the preceding example, `CommandArgs=["abc", ">", "file0"]`. The value `["abc", ">", "file0"]` must be converted into a string in the JSON array format. If this parameter is not required, you do not need to specify it.
	//
	// example:
	//
	// [\\"-c\\",\\"echo \\\\\\"test\\\\\\" > /home/nas/test.log && sleep 10000000s\\"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The description of the ConfigMap that you want to mount. Use the configuration item that you created on the Namespace Configuration Items page to inject configuration information into the container. The parameters are described as follows:
	//
	// - **configMapId**: The ID of the ConfigMap instance. You can call the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the ID.
	//
	// - **key**: The key.
	//
	// > You can pass the `sae-sys-configmap-all` parameter to mount all keys.
	//
	// - **mountPath**: The mount path.
	//
	// - **ConfigMapName**: The name of the ConfigMap.
	ConfigMapMountDesc []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The maximum CPU resources that the sidecar container can use from the main container.
	//
	// example:
	//
	// 500
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The shared temporary storage. Set a temporary storage directory and mount it to the main container and the sidecar container.
	EmptyDirDesc []*DescribeApplicationConfigResponseBodyDataSidecarContainersConfigEmptyDirDesc `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty" type:"Repeated"`
	// The environment variables of the container. You can customize environment variables or reference a ConfigMap. To reference a ConfigMap, you must first create a ConfigMap instance. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//   - **name**: The name of the environment variable.
	//
	//   - **value**: The value of the environment variable. This parameter takes precedence over valueFrom.
	//
	// - Reference a configuration item (valueFrom)
	//
	//   - **name**: The name of the environment variable. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<ConfigMap name>`, for example, `sae-sys-configmap-all-test1`.
	//
	//   - **valueFrom**: The reference of the environment variable. Set the value to `configMapRef`.
	//
	//     - **configMapId**: The ID of the ConfigMap.
	//
	//     - **key**: The key. If you want to reference all keys, do not specify this parameter.
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
	// The health check on the container.
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The maximum memory resources that the sidecar container can use from the main container.
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
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	PreStop   *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The check on the application startup status.
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The description of the secret that you want to mount.
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
	// The ID of the ConfigMap instance.
	//
	// example:
	//
	// 7361
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	// The name of the ConfigMap.
	//
	// example:
	//
	// ConfigMap-test
	ConfigMapName *string `json:"ConfigMapName,omitempty" xml:"ConfigMapName,omitempty"`
	// The key of the ConfigMap.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path of the container.
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
	// The path on which the volume is mounted in the container.
	//
	// example:
	//
	// /mnt/cache
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The name of the temporary storage.
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
	// The key of the data value that is encoded in Base64.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path.
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The ID of the secret instance.
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The name of the secret instance.
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

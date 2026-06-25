// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrAssumeRoleArn(v string) *DeployApplicationRequest
	GetAcrAssumeRoleArn() *string
	SetAcrInstanceId(v string) *DeployApplicationRequest
	GetAcrInstanceId() *string
	SetAgentVersion(v string) *DeployApplicationRequest
	GetAgentVersion() *string
	SetAlbIngressReadinessGate(v string) *DeployApplicationRequest
	GetAlbIngressReadinessGate() *string
	SetAppId(v string) *DeployApplicationRequest
	GetAppId() *string
	SetAssociateEip(v bool) *DeployApplicationRequest
	GetAssociateEip() *bool
	SetAutoEnableApplicationScalingRule(v bool) *DeployApplicationRequest
	GetAutoEnableApplicationScalingRule() *bool
	SetBatchWaitTime(v int32) *DeployApplicationRequest
	GetBatchWaitTime() *int32
	SetChangeOrderDesc(v string) *DeployApplicationRequest
	GetChangeOrderDesc() *string
	SetCommand(v string) *DeployApplicationRequest
	GetCommand() *string
	SetCommandArgs(v string) *DeployApplicationRequest
	GetCommandArgs() *string
	SetConfigMapMountDesc(v string) *DeployApplicationRequest
	GetConfigMapMountDesc() *string
	SetCpu(v int32) *DeployApplicationRequest
	GetCpu() *int32
	SetCustomHostAlias(v string) *DeployApplicationRequest
	GetCustomHostAlias() *string
	SetCustomImageNetworkType(v string) *DeployApplicationRequest
	GetCustomImageNetworkType() *string
	SetDeploy(v string) *DeployApplicationRequest
	GetDeploy() *string
	SetDotnet(v string) *DeployApplicationRequest
	GetDotnet() *string
	SetEdasContainerVersion(v string) *DeployApplicationRequest
	GetEdasContainerVersion() *string
	SetEmptyDirDesc(v string) *DeployApplicationRequest
	GetEmptyDirDesc() *string
	SetEnableAhas(v string) *DeployApplicationRequest
	GetEnableAhas() *string
	SetEnableCpuBurst(v bool) *DeployApplicationRequest
	GetEnableCpuBurst() *bool
	SetEnableGreyTagRoute(v bool) *DeployApplicationRequest
	GetEnableGreyTagRoute() *bool
	SetEnableNamespaceAgentVersion(v bool) *DeployApplicationRequest
	GetEnableNamespaceAgentVersion() *bool
	SetEnableNewArms(v bool) *DeployApplicationRequest
	GetEnableNewArms() *bool
	SetEnablePrometheus(v bool) *DeployApplicationRequest
	GetEnablePrometheus() *bool
	SetEnableSidecarResourceIsolated(v bool) *DeployApplicationRequest
	GetEnableSidecarResourceIsolated() *bool
	SetEnvs(v string) *DeployApplicationRequest
	GetEnvs() *string
	SetGpuConfig(v string) *DeployApplicationRequest
	GetGpuConfig() *string
	SetHtml(v string) *DeployApplicationRequest
	GetHtml() *string
	SetImagePullSecrets(v string) *DeployApplicationRequest
	GetImagePullSecrets() *string
	SetImageUrl(v string) *DeployApplicationRequest
	GetImageUrl() *string
	SetInitContainersConfig(v []*InitContainerConfig) *DeployApplicationRequest
	GetInitContainersConfig() []*InitContainerConfig
	SetJarStartArgs(v string) *DeployApplicationRequest
	GetJarStartArgs() *string
	SetJarStartOptions(v string) *DeployApplicationRequest
	GetJarStartOptions() *string
	SetJdk(v string) *DeployApplicationRequest
	GetJdk() *string
	SetKafkaConfigs(v string) *DeployApplicationRequest
	GetKafkaConfigs() *string
	SetLabels(v map[string]*string) *DeployApplicationRequest
	GetLabels() map[string]*string
	SetLiveness(v string) *DeployApplicationRequest
	GetLiveness() *string
	SetLokiConfigs(v string) *DeployApplicationRequest
	GetLokiConfigs() *string
	SetMaxSurgeInstanceRatio(v int32) *DeployApplicationRequest
	GetMaxSurgeInstanceRatio() *int32
	SetMaxSurgeInstances(v int32) *DeployApplicationRequest
	GetMaxSurgeInstances() *int32
	SetMemory(v int32) *DeployApplicationRequest
	GetMemory() *int32
	SetMicroRegistration(v string) *DeployApplicationRequest
	GetMicroRegistration() *string
	SetMicroRegistrationConfig(v string) *DeployApplicationRequest
	GetMicroRegistrationConfig() *string
	SetMicroserviceEngineConfig(v string) *DeployApplicationRequest
	GetMicroserviceEngineConfig() *string
	SetMinReadyInstanceRatio(v int32) *DeployApplicationRequest
	GetMinReadyInstanceRatio() *int32
	SetMinReadyInstances(v int32) *DeployApplicationRequest
	GetMinReadyInstances() *int32
	SetMountDesc(v string) *DeployApplicationRequest
	GetMountDesc() *string
	SetMountHost(v string) *DeployApplicationRequest
	GetMountHost() *string
	SetNasConfigs(v string) *DeployApplicationRequest
	GetNasConfigs() *string
	SetNasId(v string) *DeployApplicationRequest
	GetNasId() *string
	SetNewSaeVersion(v string) *DeployApplicationRequest
	GetNewSaeVersion() *string
	SetOidcRoleName(v string) *DeployApplicationRequest
	GetOidcRoleName() *string
	SetOssAkId(v string) *DeployApplicationRequest
	GetOssAkId() *string
	SetOssAkSecret(v string) *DeployApplicationRequest
	GetOssAkSecret() *string
	SetOssMountDescs(v string) *DeployApplicationRequest
	GetOssMountDescs() *string
	SetPackageType(v string) *DeployApplicationRequest
	GetPackageType() *string
	SetPackageUrl(v string) *DeployApplicationRequest
	GetPackageUrl() *string
	SetPackageVersion(v string) *DeployApplicationRequest
	GetPackageVersion() *string
	SetPhp(v string) *DeployApplicationRequest
	GetPhp() *string
	SetPhpArmsConfigLocation(v string) *DeployApplicationRequest
	GetPhpArmsConfigLocation() *string
	SetPhpConfig(v string) *DeployApplicationRequest
	GetPhpConfig() *string
	SetPhpConfigLocation(v string) *DeployApplicationRequest
	GetPhpConfigLocation() *string
	SetPostStart(v string) *DeployApplicationRequest
	GetPostStart() *string
	SetPreStop(v string) *DeployApplicationRequest
	GetPreStop() *string
	SetPvtzDiscoverySvc(v string) *DeployApplicationRequest
	GetPvtzDiscoverySvc() *string
	SetPython(v string) *DeployApplicationRequest
	GetPython() *string
	SetPythonModules(v string) *DeployApplicationRequest
	GetPythonModules() *string
	SetReadiness(v string) *DeployApplicationRequest
	GetReadiness() *string
	SetReplicas(v int32) *DeployApplicationRequest
	GetReplicas() *int32
	SetSecretMountDesc(v string) *DeployApplicationRequest
	GetSecretMountDesc() *string
	SetSecurityGroupId(v string) *DeployApplicationRequest
	GetSecurityGroupId() *string
	SetServiceTags(v string) *DeployApplicationRequest
	GetServiceTags() *string
	SetSidecarContainersConfig(v []*SidecarContainerConfig) *DeployApplicationRequest
	GetSidecarContainersConfig() []*SidecarContainerConfig
	SetSlsConfigs(v string) *DeployApplicationRequest
	GetSlsConfigs() *string
	SetSlsLogEnvTags(v string) *DeployApplicationRequest
	GetSlsLogEnvTags() *string
	SetStartupProbe(v string) *DeployApplicationRequest
	GetStartupProbe() *string
	SetSwimlanePvtzDiscoverySvc(v string) *DeployApplicationRequest
	GetSwimlanePvtzDiscoverySvc() *string
	SetTerminationGracePeriodSeconds(v int32) *DeployApplicationRequest
	GetTerminationGracePeriodSeconds() *int32
	SetTimezone(v string) *DeployApplicationRequest
	GetTimezone() *string
	SetTomcatConfig(v string) *DeployApplicationRequest
	GetTomcatConfig() *string
	SetUpdateStrategy(v string) *DeployApplicationRequest
	GetUpdateStrategy() *string
	SetVSwitchId(v string) *DeployApplicationRequest
	GetVSwitchId() *string
	SetWarStartOptions(v string) *DeployApplicationRequest
	GetWarStartOptions() *string
	SetWebContainer(v string) *DeployApplicationRequest
	GetWebContainer() *string
}

type DeployApplicationRequest struct {
	// The ARN of the RAM role required to pull images across accounts. For more information, see [Authorize cross-account image pulls using RAM roles](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The Container Registry Enterprise Edition instance ID. Required when **ImageUrl*	- is from Container Registry Enterprise Edition.
	//
	// example:
	//
	// cri-xxxxxx
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The AliyunAgent version.
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
	// The ID of the application to deploy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Whether to associate an EIP. Values:
	//
	// - **true**: Associate.
	//
	// - **false**: Do not associate.
	//
	// example:
	//
	// true
	AssociateEip *bool `json:"AssociateEip,omitempty" xml:"AssociateEip,omitempty"`
	// Whether to automatically enable application Auto Scaling rules. Values:
	//
	// - **true**: Enable.
	//
	// - **false**: Disable.
	//
	// example:
	//
	// true
	AutoEnableApplicationScalingRule *bool `json:"AutoEnableApplicationScalingRule,omitempty" xml:"AutoEnableApplicationScalingRule,omitempty"`
	// The wait time between batches, in seconds.
	//
	// example:
	//
	// 10
	BatchWaitTime *int32 `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	// The description of the release task.
	//
	// example:
	//
	// Start application
	ChangeOrderDesc *string `json:"ChangeOrderDesc,omitempty" xml:"ChangeOrderDesc,omitempty"`
	// The startup command for your image. This command must be an executable object inside the container. Example:
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
	// In this example, Command="echo" and `CommandArgs=["abc", ">", "file0"]`.
	//
	// example:
	//
	// echo
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments for the startup command **Command**. Format:
	//
	// `["a","b"]`
	//
	// In the earlier example, `CommandArgs=["abc", ">", "file0"]`. The value `["abc", ">", "file0"]` must be converted to a string in JSON array format. Leave this field empty if no arguments are needed.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The mount description for a **ConfigMap**. Use configuration items created on the namespace configuration page to inject configuration into your container. Parameters:
	//
	// - **configMapId**: The ID of the ConfigMap instance. Get it by calling the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) API.
	//
	// - **key**: The key.
	//
	// > You can mount all keys by passing `sae-sys-configmap-all`.
	//
	// - **mountPath**: The mount path.
	//
	// example:
	//
	// [{"configMapId":16,"key":"test","mountPath":"/tmp"}]
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	// The CPU required per instance, in milliCPU. Cannot be zero. Supported fixed specifications:
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
	// - **12000**
	//
	// - **16000**
	//
	// - **32000**
	//
	// example:
	//
	// 1000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Custom host mappings inside your container. Values:
	//
	// - **hostName**: A domain name or hostname.
	//
	// - **ip**: An IP address.
	//
	// example:
	//
	// [{"hostName":"samplehost","ip":"127.0.0.1"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The custom image type. Set to an empty string for non-custom images:
	//
	// - internet: Public network image
	//
	// - intranet: Private network image
	//
	// example:
	//
	// internet
	CustomImageNetworkType *string `json:"CustomImageNetworkType,omitempty" xml:"CustomImageNetworkType,omitempty"`
	// This parameter applies only to stopped applications. If you call **DeployApplication*	- on a running application, it deploys immediately.
	//
	// - **true**: Default. Deploys immediately, applies the new configuration, and starts instances.
	//
	// - **false**: Applies the new configuration only. Does not start application instances.
	//
	// example:
	//
	// true
	Deploy *string `json:"Deploy,omitempty" xml:"Deploy,omitempty"`
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
	// .NET 3.1
	Dotnet *string `json:"Dotnet,omitempty" xml:"Dotnet,omitempty"`
	// The version of the application runtime environment for HSF applications, such as Ali-Tomcat containers.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// The configuration for shared temporary storage.
	//
	// example:
	//
	// [{\\"name\\":\\"workdir\\",\\"mountPath\\":\\"/usr/local/tomcat/webapps\\"}]
	EmptyDirDesc *string `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty"`
	// Whether to integrate with AHAS. Values:
	//
	// - **true**: Integrate with AHAS.
	//
	// - **false**: Do not integrate with AHAS.
	//
	// example:
	//
	// false
	EnableAhas *string `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	// Whether to enable CPU Burst:
	//
	// - true: Enable.
	//
	// - false: Do not enable.
	//
	// example:
	//
	// true
	EnableCpuBurst *bool `json:"EnableCpuBurst,omitempty" xml:"EnableCpuBurst,omitempty"`
	// Whether to enable traffic canary rules. These rules apply only to Spring Cloud and Dubbo applications. Values:
	//
	// - **true**: Enable canary rules.
	//
	// - **false**: Disable canary rules.
	//
	// example:
	//
	// false
	EnableGreyTagRoute *bool `json:"EnableGreyTagRoute,omitempty" xml:"EnableGreyTagRoute,omitempty"`
	// Whether to reuse the namespace Agent version configuration.
	//
	// example:
	//
	// true
	EnableNamespaceAgentVersion *bool `json:"EnableNamespaceAgentVersion,omitempty" xml:"EnableNamespaceAgentVersion,omitempty"`
	// Whether to enable the new ARMS feature:
	//
	// - true: Enable.
	//
	// - false: Do not enable.
	//
	// example:
	//
	// true
	EnableNewArms *bool `json:"EnableNewArms,omitempty" xml:"EnableNewArms,omitempty"`
	// Whether to enable Prometheus custom metric collection.
	//
	// example:
	//
	// false
	EnablePrometheus *bool `json:"EnablePrometheus,omitempty" xml:"EnablePrometheus,omitempty"`
	// Whether to isolate sidecar resources:
	//
	// - true: Isolate.
	//
	// - false: Do not isolate.
	//
	// example:
	//
	// true
	EnableSidecarResourceIsolated *bool `json:"EnableSidecarResourceIsolated,omitempty" xml:"EnableSidecarResourceIsolated,omitempty"`
	// The environment variables for your container. You can define custom variables or reference configuration items. To reference a configuration item, first create a ConfigMap instance. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Values:
	//
	// - Custom configuration
	//
	//   - **name**: The name of the environment variable.
	//
	//   - **value**: The value of the environment variable. Takes precedence over valueFrom.
	//
	// - Reference a configuration item (valueFrom)
	//
	//   - **name**: The name of the environment variable. You can reference a single key or all keys. To reference all keys, use `sae-sys-configmap-all-<configmap-name>`, for example `sae-sys-configmap-all-test1`.
	//
	//   - **valueFrom**: The reference type. Set to `configMapRef`.
	//
	//   - **configMapId**: The ID of the ConfigMap instance.
	//
	//   - **key**: The key. Omit this field if you reference all keys.
	//
	// - Reference a secret (valueFrom)
	//
	//   - **name**: The name of the environment variable. You can reference a single key or all keys. To reference all keys, use `sae-sys-secret-all-<secret-name>`, for example `sae-sys-secret-all-test1`.
	//
	//   - **valueFrom**: The reference type. Set to `secretRef`.
	//
	//   - **secretId**: The ID of the secret.
	//
	//   - **key**: The key. Omit this field if you reference all keys.
	//
	// example:
	//
	// [ { "name": "sae-sys-configmap-all-hello", "valueFrom": { "configMapRef": { "configMapId": 100, "key": "" } } }, { "name": "hello", "valueFrom": { "configMapRef": { "configMapId": 101, "key": "php-fpm" } } }, { "name": "sae-sys-secret-all-hello", "valueFrom": { “secretRef": { “secretId": 100, "key": "" } } }, { "name": “password”, "valueFrom": { “secretRef": { “secretId": 101, "key": “password” } } }, { "name": "envtmp", "value": "newenv" } ]
	Envs      *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	GpuConfig *string `json:"GpuConfig,omitempty" xml:"GpuConfig,omitempty"`
	// The Nginx version:
	//
	// - nginx 1.20
	//
	// - nginx 1.22
	//
	// - nginx 1.24
	//
	// - nginx 1.26
	//
	// - nginx 1.28
	//
	// example:
	//
	// nginx 1.28
	Html *string `json:"Html,omitempty" xml:"Html,omitempty"`
	// The ID of the corresponding secret.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The registry address of your image. Required when **Package Type*	- is **Image**.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/sae_test/ali_sae_test:0.0.1
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The initialization container configuration.
	InitContainersConfig []*InitContainerConfig `json:"InitContainersConfig,omitempty" xml:"InitContainersConfig,omitempty" type:"Repeated"`
	// Startup arguments for your JAR package. Default startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// -Xms4G -Xmx4G
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// Startup options for your JAR package. Default startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// custom-option
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The JDK version that your deployment package depends on. Supported versions include the following:
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
	// This parameter is not supported when **Package Type*	- is **Image**.
	//
	// example:
	//
	// Open JDK 8
	Jdk *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	// The configuration for collecting logs to Kafka. Values:
	//
	// - **kafkaEndpoint**: The endpoint for the Kafka API.
	//
	// - **kafkaInstanceId**: The Kafka instance ID.
	//
	// - **kafkaConfigs**: The configuration for one or more log entries. For examples and details, see the \\*\\*kafkaConfigs\\*\\	- request parameter in this topic.
	//
	// example:
	//
	// {"kafkaEndpoint":"10.0.X.XXX:XXXX,10.0.X.XXX:XXXX,10.0.X.XXX:XXXX","kafkaInstanceId":"alikafka_pre-cn-7pp2l8kr****","kafkaConfigs":[{"logType":"file_log","logDir":"/tmp/a.log","kafkaTopic":"test2"},{"logType":"stdout","logDir":"","kafkaTopic":"test"}]}
	KafkaConfigs *string            `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	Labels       map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Health checks for your container. Containers that fail health checks are terminated and restarted. Supported methods:
	//
	// - **exec**: For example, `{"exec":{"command":["sh","-c","cat/home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}`
	//
	// - **httpGet**: For example, `{"httpGet":{"path":"/","port":18091,"scheme":"HTTP","isContainKeyWord":true,"keyWord":"SAE"},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// - **tcpSocket**: For example, `{"tcpSocket":{"port":18091},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// > You can select only one health check method.
	//
	// Parameters:
	//
	// - **exec.command**: The health check command.
	//
	// - **httpGet.path**: The path to access.
	//
	// - **httpGet.scheme**: **HTTP*	- or **HTTPS**.
	//
	// - **httpGet.isContainKeyWord**: **true*	- means the response contains a keyword. **false*	- means it does not. If omitted, advanced features are disabled.
	//
	// - **httpGet.keyWord**: Your custom keyword. Include **isContainKeyWord*	- when using this field.
	//
	// - **tcpSocket.port**: The port for TCP connection checks.
	//
	// - **initialDelaySeconds**: The delay before the first health check, in seconds. Default is 10.
	//
	// - **periodSeconds**: The interval between health checks, in seconds. Default is 30.
	//
	// - **timeoutSeconds**: The timeout for each health check, in seconds. Default is 1. If set to 0 or omitted, the default is 1 second.
	//
	// example:
	//
	// {"exec":{"command":["sleep","5s"]},"initialDelaySeconds":10,"timeoutSeconds":11}
	Liveness    *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	LokiConfigs *string `json:"LokiConfigs,omitempty" xml:"LokiConfigs,omitempty"`
	// The maximum number of surge instances as a percentage of total instances. Values:
	//
	// If the minimum available instances is 100%, the maximum surge cannot be set to 0. If set to -1, the system uses its recommended value: 30% of your current instance count. For example, with 10 instances, 10 × 30% = 3.
	//
	// example:
	//
	// -1
	MaxSurgeInstanceRatio *int32 `json:"MaxSurgeInstanceRatio,omitempty" xml:"MaxSurgeInstanceRatio,omitempty"`
	// The maximum number of surge instances during a rolling update. Values:
	//
	// If the minimum available instances is 100%, the maximum surge cannot be set to 0. If set to -1, the system uses its recommended value: 30% of your current instance count. For example, with 10 instances, 10 × 30% = 3.
	//
	// example:
	//
	// -1
	MaxSurgeInstances *int32 `json:"MaxSurgeInstances,omitempty" xml:"MaxSurgeInstances,omitempty"`
	// The memory required per instance, in MB. Cannot be zero. Memory and CPU are paired. Supported fixed specifications:
	//
	// - **1024**: Pairs with 500 and 1000 milliCPU.
	//
	// - **2048**: Pairs with 500, 1000, and 2000 milliCPU.
	//
	// - **4096**: Pairs with 1000, 2000, and 4000 milliCPU.
	//
	// - **8192**: Pairs with 2000, 4000, and 8000 milliCPU.
	//
	// - **12288**: Pairs with 12000 milliCPU.
	//
	// - **16384**: Pairs with 4000, 8000, and 16000 milliCPU.
	//
	// - **24576**: Pairs with 12000 milliCPU.
	//
	// - **32768**: Pairs with 16000 milliCPU.
	//
	// - **65536**: Pairs with 8000, 16000, and 32000 milliCPU.
	//
	// - **131072**: Pairs with 32000 milliCPU.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Select a Nacos registry center. Values:
	//
	// - **0**: Built-in Nacos in SAE.
	//
	// - **1**: Self-managed Nacos.
	//
	// - **2**: MSE Nacos Commercial Edition.
	//
	// > If you select built-in Nacos in SAE, you cannot retrieve its configuration.
	//
	// example:
	//
	// "0"
	MicroRegistration *string `json:"MicroRegistration,omitempty" xml:"MicroRegistration,omitempty"`
	// The registry configuration. Applies only when the registry type is MSE Nacos Enterprise Edition.
	//
	// example:
	//
	// {\\"instanceId\\":\\"mse-cn-zvp2bh6h70r\\",\\"namespace\\":\\"4c0aa74f-57cb-423c-b6af-5d9f2d0e3dbd\\"}
	MicroRegistrationConfig *string `json:"MicroRegistrationConfig,omitempty" xml:"MicroRegistrationConfig,omitempty"`
	// Configure microservice governance features.
	//
	// - Enable microservice governance (enable):
	//
	//   - true: Enable
	//
	//   - false: Disable
	//
	// - Configure graceful start and shutdown (mseLosslessRule):
	//
	//   - delayTime: Delay time
	//
	//   - enable: Whether to enable graceful start. true enables it. false disables it.
	//
	//   - notice: Whether to enable notifications. true enables them. false disables them.
	//
	//   - warmupTime: Warm-up duration for small traffic, in seconds.
	//
	// example:
	//
	// {"enable": true,"mseLosslessRule": {"delayTime": 0,"enable": false,"notice": false,"warmupTime": 120}}
	MicroserviceEngineConfig *string `json:"MicroserviceEngineConfig,omitempty" xml:"MicroserviceEngineConfig,omitempty"`
	// The minimum number of available instances as a percentage of total instances. Values:
	//
	// - **-1**: Use the default value. No percentage is applied.
	//
	// - **0–100**: Percentage value. Rounded up. For example, if set to **50**% and you have 5 instances, the minimum is 3.
	//
	// > If both **MinReadyInstances*	- and **MinReadyInstanceRatio*	- are provided, and **MinReadyInstanceRatio*	- is not **-1**, then **MinReadyInstanceRatio*	- takes precedence. For example, if **MinReadyInstances*	- is **5*	- and **MinReadyInstanceRatio*	- is **50**, the system calculates the minimum based on 50%.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of instances that remain available during a rolling update. Values:
	//
	// - If set to **0**, your application experiences downtime during updates.
	//
	// - If set to -1, the system uses its recommended value: 25% of your current instance count. For example, with 5 instances, 5 × 25% = 1.25, rounded up to 2.
	//
	// > We recommend setting this value to at least 1 to avoid service interruptions.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// We recommend using **NasConfigs*	- instead of this field. The NAS mount description. If your NAS configuration remains unchanged, omit this parameter. To clear your NAS configuration, set this field to an empty string.
	//
	// example:
	//
	// [{mountPath: "/tmp", nasPath: "/"}]
	MountDesc *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	// We recommend using **NasConfigs*	- instead of this field. The mount target of the NAS in your application\\"s VPC. If your NAS configuration remains unchanged, omit this parameter. To clear your NAS configuration, set this field to an empty string.
	//
	// example:
	//
	// 10d3b4bc9****.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The configuration for mounting NAS. Values:
	//
	// - **mountPath**: The mount path in the container.
	//
	// - **readOnly**: Set to **false*	- for read and write permissions.
	//
	// - **nasId**: The NAS ID.
	//
	// - **mountDomain**: The mount target address. For more information, see [DescribeMountTargets](https://help.aliyun.com/document_detail/62626.html).
	//
	// - **nasPath**: The relative directory in NAS.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// We recommend using **NasConfigs*	- instead of this field. The ID of the NAS file system. If your NAS configuration remains unchanged, omit this parameter. To clear your NAS configuration, set this field to an empty string.
	//
	// example:
	//
	// 10d3b4****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The application version:
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
	// > Create an OIDC identity provider and an associated role in the same region before using this parameter. For more information, see [Create an OIDC identity provider](https://help.aliyun.com/document_detail/2331022.html) and [Create a role for SSO identity providers](https://help.aliyun.com/document_detail/2331016.html).
	//
	// example:
	//
	// sae-test
	OidcRoleName *string `json:"OidcRoleName,omitempty" xml:"OidcRoleName,omitempty"`
	// The AccessKey ID for OSS read and write operations.
	//
	// example:
	//
	// xxxxxx
	OssAkId *string `json:"OssAkId,omitempty" xml:"OssAkId,omitempty"`
	// The AccessKey secret for OSS read and write operations.
	//
	// example:
	//
	// xxxxxx
	OssAkSecret *string `json:"OssAkSecret,omitempty" xml:"OssAkSecret,omitempty"`
	// The OSS mount description. Parameters:
	//
	// - **bucketName**: The name of the bucket.
	//
	// - **bucketPath**: The directory or object in OSS. If the directory does not exist, an error occurs.
	//
	// - **mountPath**: The path in your SAE container. If the path exists, it is overwritten. If it does not exist, it is created.
	//
	// - **readOnly**: Whether the container path has read-only access to the mounted resource. Values:
	//
	//   - **true**: Read-only.
	//
	//   - **false**: Read and write.
	//
	// example:
	//
	// [{"bucketName": "oss-bucket", "bucketPath": "data/user.data", "mountPath": "/usr/data/user.data", "readOnly": true}]
	OssMountDescs *string `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty"`
	// The type of your application package. Values:
	//
	// - For Java applications: **FatJar**, **War**, and **Image**.
	//
	// - For PHP applications:
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
	// - For Python applications: **PythonZip*	- and **Image**.
	//
	// example:
	//
	// FatJar
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The URL of your deployment package. Required when **Package Type*	- is **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// http://myoss.oss-cn-hangzhou.aliyuncs.com/my-buc/2019-06-30/****.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version number of your deployment package. Required when **Package Type*	- is **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// 1.0.1
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The PHP version that your PHP deployment package depends on. Not supported for images.
	//
	// example:
	//
	// PHP-FPM 7.0
	Php *string `json:"Php,omitempty" xml:"Php,omitempty"`
	// The mount path for PHP application monitoring. Ensure your PHP server loads the configuration file at this path. You do not need to manage the configuration content. SAE renders the correct configuration automatically.
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
	// The mount path for the PHP startup configuration. Ensure your PHP server uses this configuration file to start.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// A script that runs after your container starts. It executes immediately after the container is created. Format: `{"exec":{"command":["sh","-c","echo hello"]}}`
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// A script that runs before your container stops. It executes just before the container is deleted. Format: `{"exec":{"command":["sh","-c","echo hello"]}}`
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// Enable K8s Service registration and discovery. Values:
	//
	// - **portProtocols**: Port and protocol. Port range is [1,65535]. Protocols supported: **TCP*	- and **UDP**.
	//
	// - portAndProtocol: Port and protocol. Port range is [1,65535]. Protocols supported: **TCP*	- and **UDP**. **portProtocols*	- takes precedence. If both are set, only **portProtocols*	- applies.
	//
	// - **enable**: Enable K8s Service registration and discovery.
	//
	// example:
	//
	// {"portProtocols":[{"port":18012,"protocol":"TCP"}],"portAndProtocol":{"18012":"TCP"},"enable":true}
	PvtzDiscoverySvc *string `json:"PvtzDiscoverySvc,omitempty" xml:"PvtzDiscoverySvc,omitempty"`
	// The Python runtime environment. Supported: **PYTHON 3.9.15**.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// Custom module dependencies. By default, dependencies defined in requirements.txt in the root directory are installed. If no configuration or custom packages exist, specify the dependencies to install.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// Startup status checks for your application. Containers that repeatedly fail readiness checks are terminated and restarted. Containers that fail readiness checks receive no SLB traffic. Supports **exec**, **httpGet**, and **tcpSocket**. For examples, see the **Liveness*	- parameter.
	//
	// > You can select only one health check method.
	//
	// example:
	//
	// {"exec":{"command":["sleep","6s"]},"initialDelaySeconds":15,"timeoutSeconds":12}
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The mount description for a **Secret**. Use secrets created on the namespace secrets page to inject sensitive information into your container. Parameters:
	//
	// - **secretId**: The ID of the secret instance. Get it by calling the ListSecrets API.
	//
	// - **key**: The key.
	//
	// > You can mount all keys by passing `sae-sys-secret-all`.
	//
	// - **mountPath**: The mount path.
	//
	// example:
	//
	// [{“secretId":10,”key":"test","mountPath":"/tmp"}]
	SecretMountDesc *string `json:"SecretMountDesc,omitempty" xml:"SecretMountDesc,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The canary tags configured for your application.
	//
	// example:
	//
	// {\\"alicloud.service.tag\\":\\"g1\\"}
	ServiceTags *string `json:"ServiceTags,omitempty" xml:"ServiceTags,omitempty"`
	// Container configuration information.
	SidecarContainersConfig []*SidecarContainerConfig `json:"SidecarContainersConfig,omitempty" xml:"SidecarContainersConfig,omitempty" type:"Repeated"`
	// The configuration for collecting logs to Simple Log Service (SLS).
	//
	// - Using SAE-managed SLS resources: `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - Using custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// Parameters:
	//
	// - **projectName**: The name of the SLS project.
	//
	// - **logDir**: The log file path.
	//
	// - logType: The log type. **stdout*	- means standard output logs from the container. Only one **stdout*	- entry is allowed. If omitted, file logs are collected.
	//
	// - **logstoreName**: The name of the SLS Logstore.
	//
	// - **logtailName**: The name of the SLS Logtail. If omitted, a new Logtail is created.
	//
	// If your SLS collection configuration remains unchanged across deployments, omit this parameter. To disable SLS collection, set this field to an empty string.
	//
	// > Projects automatically created by SAE are deleted when the application is deleted. Do not select these projects when choosing an existing project.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// The SLS log tags.
	SlsLogEnvTags *string `json:"SlsLogEnvTags,omitempty" xml:"SlsLogEnvTags,omitempty"`
	// Enable application startup probing.
	//
	// - Success: The application starts successfully. If you configure Liveness and Readiness checks, they run after startup.
	//
	// - Failure: The application fails to start. SAE reports an error and restarts the container automatically.
	//
	// > Description
	//
	// >
	//
	// > - Supports exec, httpGet, and tcpSocket. For examples, see the Liveness parameter.
	//
	// >
	//
	// > - You can select only one health check method.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	StartupProbe *string `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	// Configures service discovery and end-to-end canary release based on a Kubernetes Service:
	//
	// - enable: Specifies whether to enable the end-to-end canary release feature.
	//
	//   - true: Enables the feature.
	//
	//   - false: Disables the feature.
	//
	// - namespaceId: The namespace ID.
	//
	// - portAndProtocol: The listening port and protocol. The format is {"\\<port>:\\<protocol>":"\\<target_port>"}.
	//
	// - portProtocols: A list of ports and protocols for the service.
	//
	//   - port: The port number.
	//
	//   - protocol: The protocol.
	//
	//   - targetPort: The container port.
	//
	// - pvtzDiscoveryName: The service discovery name.
	//
	// - serviceId: The service ID.
	//
	// - serviceName: The service name.
	//
	// example:
	//
	// {\\"enable\\":\\"false\\",\\"namespaceId\\":\\"cn-beijing:test\\",\\"portAndProtocol\\":{\\"2000:TCP\\":\\"18081\\"},\\"portProtocols\\":[{\\"port\\":2000,\\"protocol\\":\\"TCP\\",\\"targetPort\\":18081}],\\"pvtzDiscoveryName\\":\\"cn-beijing-1421801774382676\\",\\"serviceId\\":\\"3513\\",\\"serviceName\\":\\"demo-gray.test\\"}
	SwimlanePvtzDiscoverySvc *string `json:"SwimlanePvtzDiscoverySvc,omitempty" xml:"SwimlanePvtzDiscoverySvc,omitempty"`
	// The graceful shutdown timeout, in seconds. Default is 30. Valid values: 1–300.
	//
	// example:
	//
	// 10
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The time zone. Default is **Asia/Shanghai**.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The Tomcat configuration. Set to an empty string or {} to delete the configuration. Values:
	//
	// - **port**: Port range is 1024–65535. Ports below 1024 require root privileges. Because containers run with admin privileges, use ports above 1024. Default is 8080.
	//
	// - **contextPath**: The access path. Default is the root directory /.
	//
	// - **maxThreads**: The size of the connection pool. Default is 400.
	//
	// - uriEncoding: The encoding format for Tomcat. Options include **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. Default is **ISO-8859-1**.
	//
	// - **useBodyEncodingForUri**: Whether to use body encoding for URLs. Default is **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	// The release strategy. When MinReadyInstances equals 1, set UpdateStrategy to an empty string. When **MinReadyInstances*	- is greater than 1, examples include the following:
	//
	// - Canary release with 1 instance, followed by 2 automatic batches with a 1-minute interval: `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}`
	//
	// - Canary release with 1 instance, followed by 2 manual batches: `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"manual"},"grayUpdate":{"gray":1}}`
	//
	// - Two automatic batches with a 0-minute interval: `{"type":"BatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":0}}`
	//
	// Parameters:
	//
	// - **type**: The release strategy type. Options are **GrayBatchUpdate*	- (canary release) or **BatchUpdate*	- (phased release).
	//
	// - **batchUpdate**: The phased release strategy.
	//
	//   - **batch**: The number of batches.
	//
	//   - **releaseType**: How batches are processed. Options are **auto*	- (automatic) or **manual*	- (manual).
	//
	//   - **batchWaitTime**: The wait time between batches, in minutes.
	//
	// - **grayUpdate**: The number of canary instances. Required when **type*	- is **GrayBatchUpdate**.
	//
	// example:
	//
	// {"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}
	UpdateStrategy *string `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
	// The virtual switch where your application instance elastic network interfaces reside. This switch must be in the specified VPC.
	//
	// example:
	//
	// vsw-bp12mw1f8k3jgygk9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The startup command for your WAR package. Configure it the same way as the startup command for images. For more information, see [Set the startup command](https://help.aliyun.com/document_detail/96677.html).
	//
	// example:
	//
	// CATALINA_OPTS=\\"$CATALINA_OPTS $Options\\" catalina.sh run
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The Tomcat version that your deployment package depends on. Supported versions include the following:
	//
	// - **apache-tomcat-7.0.91**
	//
	// - **apache-tomcat-8.5.42**
	//
	// This parameter is not supported when **Package Type*	- is **Image**.
	//
	// example:
	//
	// apache-tomcat-7.0.91
	WebContainer *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
}

func (s DeployApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeployApplicationRequest) GetAcrAssumeRoleArn() *string {
	return s.AcrAssumeRoleArn
}

func (s *DeployApplicationRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *DeployApplicationRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *DeployApplicationRequest) GetAlbIngressReadinessGate() *string {
	return s.AlbIngressReadinessGate
}

func (s *DeployApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeployApplicationRequest) GetAssociateEip() *bool {
	return s.AssociateEip
}

func (s *DeployApplicationRequest) GetAutoEnableApplicationScalingRule() *bool {
	return s.AutoEnableApplicationScalingRule
}

func (s *DeployApplicationRequest) GetBatchWaitTime() *int32 {
	return s.BatchWaitTime
}

func (s *DeployApplicationRequest) GetChangeOrderDesc() *string {
	return s.ChangeOrderDesc
}

func (s *DeployApplicationRequest) GetCommand() *string {
	return s.Command
}

func (s *DeployApplicationRequest) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *DeployApplicationRequest) GetConfigMapMountDesc() *string {
	return s.ConfigMapMountDesc
}

func (s *DeployApplicationRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *DeployApplicationRequest) GetCustomHostAlias() *string {
	return s.CustomHostAlias
}

func (s *DeployApplicationRequest) GetCustomImageNetworkType() *string {
	return s.CustomImageNetworkType
}

func (s *DeployApplicationRequest) GetDeploy() *string {
	return s.Deploy
}

func (s *DeployApplicationRequest) GetDotnet() *string {
	return s.Dotnet
}

func (s *DeployApplicationRequest) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *DeployApplicationRequest) GetEmptyDirDesc() *string {
	return s.EmptyDirDesc
}

func (s *DeployApplicationRequest) GetEnableAhas() *string {
	return s.EnableAhas
}

func (s *DeployApplicationRequest) GetEnableCpuBurst() *bool {
	return s.EnableCpuBurst
}

func (s *DeployApplicationRequest) GetEnableGreyTagRoute() *bool {
	return s.EnableGreyTagRoute
}

func (s *DeployApplicationRequest) GetEnableNamespaceAgentVersion() *bool {
	return s.EnableNamespaceAgentVersion
}

func (s *DeployApplicationRequest) GetEnableNewArms() *bool {
	return s.EnableNewArms
}

func (s *DeployApplicationRequest) GetEnablePrometheus() *bool {
	return s.EnablePrometheus
}

func (s *DeployApplicationRequest) GetEnableSidecarResourceIsolated() *bool {
	return s.EnableSidecarResourceIsolated
}

func (s *DeployApplicationRequest) GetEnvs() *string {
	return s.Envs
}

func (s *DeployApplicationRequest) GetGpuConfig() *string {
	return s.GpuConfig
}

func (s *DeployApplicationRequest) GetHtml() *string {
	return s.Html
}

func (s *DeployApplicationRequest) GetImagePullSecrets() *string {
	return s.ImagePullSecrets
}

func (s *DeployApplicationRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DeployApplicationRequest) GetInitContainersConfig() []*InitContainerConfig {
	return s.InitContainersConfig
}

func (s *DeployApplicationRequest) GetJarStartArgs() *string {
	return s.JarStartArgs
}

func (s *DeployApplicationRequest) GetJarStartOptions() *string {
	return s.JarStartOptions
}

func (s *DeployApplicationRequest) GetJdk() *string {
	return s.Jdk
}

func (s *DeployApplicationRequest) GetKafkaConfigs() *string {
	return s.KafkaConfigs
}

func (s *DeployApplicationRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *DeployApplicationRequest) GetLiveness() *string {
	return s.Liveness
}

func (s *DeployApplicationRequest) GetLokiConfigs() *string {
	return s.LokiConfigs
}

func (s *DeployApplicationRequest) GetMaxSurgeInstanceRatio() *int32 {
	return s.MaxSurgeInstanceRatio
}

func (s *DeployApplicationRequest) GetMaxSurgeInstances() *int32 {
	return s.MaxSurgeInstances
}

func (s *DeployApplicationRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *DeployApplicationRequest) GetMicroRegistration() *string {
	return s.MicroRegistration
}

func (s *DeployApplicationRequest) GetMicroRegistrationConfig() *string {
	return s.MicroRegistrationConfig
}

func (s *DeployApplicationRequest) GetMicroserviceEngineConfig() *string {
	return s.MicroserviceEngineConfig
}

func (s *DeployApplicationRequest) GetMinReadyInstanceRatio() *int32 {
	return s.MinReadyInstanceRatio
}

func (s *DeployApplicationRequest) GetMinReadyInstances() *int32 {
	return s.MinReadyInstances
}

func (s *DeployApplicationRequest) GetMountDesc() *string {
	return s.MountDesc
}

func (s *DeployApplicationRequest) GetMountHost() *string {
	return s.MountHost
}

func (s *DeployApplicationRequest) GetNasConfigs() *string {
	return s.NasConfigs
}

func (s *DeployApplicationRequest) GetNasId() *string {
	return s.NasId
}

func (s *DeployApplicationRequest) GetNewSaeVersion() *string {
	return s.NewSaeVersion
}

func (s *DeployApplicationRequest) GetOidcRoleName() *string {
	return s.OidcRoleName
}

func (s *DeployApplicationRequest) GetOssAkId() *string {
	return s.OssAkId
}

func (s *DeployApplicationRequest) GetOssAkSecret() *string {
	return s.OssAkSecret
}

func (s *DeployApplicationRequest) GetOssMountDescs() *string {
	return s.OssMountDescs
}

func (s *DeployApplicationRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *DeployApplicationRequest) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *DeployApplicationRequest) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *DeployApplicationRequest) GetPhp() *string {
	return s.Php
}

func (s *DeployApplicationRequest) GetPhpArmsConfigLocation() *string {
	return s.PhpArmsConfigLocation
}

func (s *DeployApplicationRequest) GetPhpConfig() *string {
	return s.PhpConfig
}

func (s *DeployApplicationRequest) GetPhpConfigLocation() *string {
	return s.PhpConfigLocation
}

func (s *DeployApplicationRequest) GetPostStart() *string {
	return s.PostStart
}

func (s *DeployApplicationRequest) GetPreStop() *string {
	return s.PreStop
}

func (s *DeployApplicationRequest) GetPvtzDiscoverySvc() *string {
	return s.PvtzDiscoverySvc
}

func (s *DeployApplicationRequest) GetPython() *string {
	return s.Python
}

func (s *DeployApplicationRequest) GetPythonModules() *string {
	return s.PythonModules
}

func (s *DeployApplicationRequest) GetReadiness() *string {
	return s.Readiness
}

func (s *DeployApplicationRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *DeployApplicationRequest) GetSecretMountDesc() *string {
	return s.SecretMountDesc
}

func (s *DeployApplicationRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeployApplicationRequest) GetServiceTags() *string {
	return s.ServiceTags
}

func (s *DeployApplicationRequest) GetSidecarContainersConfig() []*SidecarContainerConfig {
	return s.SidecarContainersConfig
}

func (s *DeployApplicationRequest) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *DeployApplicationRequest) GetSlsLogEnvTags() *string {
	return s.SlsLogEnvTags
}

func (s *DeployApplicationRequest) GetStartupProbe() *string {
	return s.StartupProbe
}

func (s *DeployApplicationRequest) GetSwimlanePvtzDiscoverySvc() *string {
	return s.SwimlanePvtzDiscoverySvc
}

func (s *DeployApplicationRequest) GetTerminationGracePeriodSeconds() *int32 {
	return s.TerminationGracePeriodSeconds
}

func (s *DeployApplicationRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *DeployApplicationRequest) GetTomcatConfig() *string {
	return s.TomcatConfig
}

func (s *DeployApplicationRequest) GetUpdateStrategy() *string {
	return s.UpdateStrategy
}

func (s *DeployApplicationRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DeployApplicationRequest) GetWarStartOptions() *string {
	return s.WarStartOptions
}

func (s *DeployApplicationRequest) GetWebContainer() *string {
	return s.WebContainer
}

func (s *DeployApplicationRequest) SetAcrAssumeRoleArn(v string) *DeployApplicationRequest {
	s.AcrAssumeRoleArn = &v
	return s
}

func (s *DeployApplicationRequest) SetAcrInstanceId(v string) *DeployApplicationRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *DeployApplicationRequest) SetAgentVersion(v string) *DeployApplicationRequest {
	s.AgentVersion = &v
	return s
}

func (s *DeployApplicationRequest) SetAlbIngressReadinessGate(v string) *DeployApplicationRequest {
	s.AlbIngressReadinessGate = &v
	return s
}

func (s *DeployApplicationRequest) SetAppId(v string) *DeployApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeployApplicationRequest) SetAssociateEip(v bool) *DeployApplicationRequest {
	s.AssociateEip = &v
	return s
}

func (s *DeployApplicationRequest) SetAutoEnableApplicationScalingRule(v bool) *DeployApplicationRequest {
	s.AutoEnableApplicationScalingRule = &v
	return s
}

func (s *DeployApplicationRequest) SetBatchWaitTime(v int32) *DeployApplicationRequest {
	s.BatchWaitTime = &v
	return s
}

func (s *DeployApplicationRequest) SetChangeOrderDesc(v string) *DeployApplicationRequest {
	s.ChangeOrderDesc = &v
	return s
}

func (s *DeployApplicationRequest) SetCommand(v string) *DeployApplicationRequest {
	s.Command = &v
	return s
}

func (s *DeployApplicationRequest) SetCommandArgs(v string) *DeployApplicationRequest {
	s.CommandArgs = &v
	return s
}

func (s *DeployApplicationRequest) SetConfigMapMountDesc(v string) *DeployApplicationRequest {
	s.ConfigMapMountDesc = &v
	return s
}

func (s *DeployApplicationRequest) SetCpu(v int32) *DeployApplicationRequest {
	s.Cpu = &v
	return s
}

func (s *DeployApplicationRequest) SetCustomHostAlias(v string) *DeployApplicationRequest {
	s.CustomHostAlias = &v
	return s
}

func (s *DeployApplicationRequest) SetCustomImageNetworkType(v string) *DeployApplicationRequest {
	s.CustomImageNetworkType = &v
	return s
}

func (s *DeployApplicationRequest) SetDeploy(v string) *DeployApplicationRequest {
	s.Deploy = &v
	return s
}

func (s *DeployApplicationRequest) SetDotnet(v string) *DeployApplicationRequest {
	s.Dotnet = &v
	return s
}

func (s *DeployApplicationRequest) SetEdasContainerVersion(v string) *DeployApplicationRequest {
	s.EdasContainerVersion = &v
	return s
}

func (s *DeployApplicationRequest) SetEmptyDirDesc(v string) *DeployApplicationRequest {
	s.EmptyDirDesc = &v
	return s
}

func (s *DeployApplicationRequest) SetEnableAhas(v string) *DeployApplicationRequest {
	s.EnableAhas = &v
	return s
}

func (s *DeployApplicationRequest) SetEnableCpuBurst(v bool) *DeployApplicationRequest {
	s.EnableCpuBurst = &v
	return s
}

func (s *DeployApplicationRequest) SetEnableGreyTagRoute(v bool) *DeployApplicationRequest {
	s.EnableGreyTagRoute = &v
	return s
}

func (s *DeployApplicationRequest) SetEnableNamespaceAgentVersion(v bool) *DeployApplicationRequest {
	s.EnableNamespaceAgentVersion = &v
	return s
}

func (s *DeployApplicationRequest) SetEnableNewArms(v bool) *DeployApplicationRequest {
	s.EnableNewArms = &v
	return s
}

func (s *DeployApplicationRequest) SetEnablePrometheus(v bool) *DeployApplicationRequest {
	s.EnablePrometheus = &v
	return s
}

func (s *DeployApplicationRequest) SetEnableSidecarResourceIsolated(v bool) *DeployApplicationRequest {
	s.EnableSidecarResourceIsolated = &v
	return s
}

func (s *DeployApplicationRequest) SetEnvs(v string) *DeployApplicationRequest {
	s.Envs = &v
	return s
}

func (s *DeployApplicationRequest) SetGpuConfig(v string) *DeployApplicationRequest {
	s.GpuConfig = &v
	return s
}

func (s *DeployApplicationRequest) SetHtml(v string) *DeployApplicationRequest {
	s.Html = &v
	return s
}

func (s *DeployApplicationRequest) SetImagePullSecrets(v string) *DeployApplicationRequest {
	s.ImagePullSecrets = &v
	return s
}

func (s *DeployApplicationRequest) SetImageUrl(v string) *DeployApplicationRequest {
	s.ImageUrl = &v
	return s
}

func (s *DeployApplicationRequest) SetInitContainersConfig(v []*InitContainerConfig) *DeployApplicationRequest {
	s.InitContainersConfig = v
	return s
}

func (s *DeployApplicationRequest) SetJarStartArgs(v string) *DeployApplicationRequest {
	s.JarStartArgs = &v
	return s
}

func (s *DeployApplicationRequest) SetJarStartOptions(v string) *DeployApplicationRequest {
	s.JarStartOptions = &v
	return s
}

func (s *DeployApplicationRequest) SetJdk(v string) *DeployApplicationRequest {
	s.Jdk = &v
	return s
}

func (s *DeployApplicationRequest) SetKafkaConfigs(v string) *DeployApplicationRequest {
	s.KafkaConfigs = &v
	return s
}

func (s *DeployApplicationRequest) SetLabels(v map[string]*string) *DeployApplicationRequest {
	s.Labels = v
	return s
}

func (s *DeployApplicationRequest) SetLiveness(v string) *DeployApplicationRequest {
	s.Liveness = &v
	return s
}

func (s *DeployApplicationRequest) SetLokiConfigs(v string) *DeployApplicationRequest {
	s.LokiConfigs = &v
	return s
}

func (s *DeployApplicationRequest) SetMaxSurgeInstanceRatio(v int32) *DeployApplicationRequest {
	s.MaxSurgeInstanceRatio = &v
	return s
}

func (s *DeployApplicationRequest) SetMaxSurgeInstances(v int32) *DeployApplicationRequest {
	s.MaxSurgeInstances = &v
	return s
}

func (s *DeployApplicationRequest) SetMemory(v int32) *DeployApplicationRequest {
	s.Memory = &v
	return s
}

func (s *DeployApplicationRequest) SetMicroRegistration(v string) *DeployApplicationRequest {
	s.MicroRegistration = &v
	return s
}

func (s *DeployApplicationRequest) SetMicroRegistrationConfig(v string) *DeployApplicationRequest {
	s.MicroRegistrationConfig = &v
	return s
}

func (s *DeployApplicationRequest) SetMicroserviceEngineConfig(v string) *DeployApplicationRequest {
	s.MicroserviceEngineConfig = &v
	return s
}

func (s *DeployApplicationRequest) SetMinReadyInstanceRatio(v int32) *DeployApplicationRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *DeployApplicationRequest) SetMinReadyInstances(v int32) *DeployApplicationRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *DeployApplicationRequest) SetMountDesc(v string) *DeployApplicationRequest {
	s.MountDesc = &v
	return s
}

func (s *DeployApplicationRequest) SetMountHost(v string) *DeployApplicationRequest {
	s.MountHost = &v
	return s
}

func (s *DeployApplicationRequest) SetNasConfigs(v string) *DeployApplicationRequest {
	s.NasConfigs = &v
	return s
}

func (s *DeployApplicationRequest) SetNasId(v string) *DeployApplicationRequest {
	s.NasId = &v
	return s
}

func (s *DeployApplicationRequest) SetNewSaeVersion(v string) *DeployApplicationRequest {
	s.NewSaeVersion = &v
	return s
}

func (s *DeployApplicationRequest) SetOidcRoleName(v string) *DeployApplicationRequest {
	s.OidcRoleName = &v
	return s
}

func (s *DeployApplicationRequest) SetOssAkId(v string) *DeployApplicationRequest {
	s.OssAkId = &v
	return s
}

func (s *DeployApplicationRequest) SetOssAkSecret(v string) *DeployApplicationRequest {
	s.OssAkSecret = &v
	return s
}

func (s *DeployApplicationRequest) SetOssMountDescs(v string) *DeployApplicationRequest {
	s.OssMountDescs = &v
	return s
}

func (s *DeployApplicationRequest) SetPackageType(v string) *DeployApplicationRequest {
	s.PackageType = &v
	return s
}

func (s *DeployApplicationRequest) SetPackageUrl(v string) *DeployApplicationRequest {
	s.PackageUrl = &v
	return s
}

func (s *DeployApplicationRequest) SetPackageVersion(v string) *DeployApplicationRequest {
	s.PackageVersion = &v
	return s
}

func (s *DeployApplicationRequest) SetPhp(v string) *DeployApplicationRequest {
	s.Php = &v
	return s
}

func (s *DeployApplicationRequest) SetPhpArmsConfigLocation(v string) *DeployApplicationRequest {
	s.PhpArmsConfigLocation = &v
	return s
}

func (s *DeployApplicationRequest) SetPhpConfig(v string) *DeployApplicationRequest {
	s.PhpConfig = &v
	return s
}

func (s *DeployApplicationRequest) SetPhpConfigLocation(v string) *DeployApplicationRequest {
	s.PhpConfigLocation = &v
	return s
}

func (s *DeployApplicationRequest) SetPostStart(v string) *DeployApplicationRequest {
	s.PostStart = &v
	return s
}

func (s *DeployApplicationRequest) SetPreStop(v string) *DeployApplicationRequest {
	s.PreStop = &v
	return s
}

func (s *DeployApplicationRequest) SetPvtzDiscoverySvc(v string) *DeployApplicationRequest {
	s.PvtzDiscoverySvc = &v
	return s
}

func (s *DeployApplicationRequest) SetPython(v string) *DeployApplicationRequest {
	s.Python = &v
	return s
}

func (s *DeployApplicationRequest) SetPythonModules(v string) *DeployApplicationRequest {
	s.PythonModules = &v
	return s
}

func (s *DeployApplicationRequest) SetReadiness(v string) *DeployApplicationRequest {
	s.Readiness = &v
	return s
}

func (s *DeployApplicationRequest) SetReplicas(v int32) *DeployApplicationRequest {
	s.Replicas = &v
	return s
}

func (s *DeployApplicationRequest) SetSecretMountDesc(v string) *DeployApplicationRequest {
	s.SecretMountDesc = &v
	return s
}

func (s *DeployApplicationRequest) SetSecurityGroupId(v string) *DeployApplicationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployApplicationRequest) SetServiceTags(v string) *DeployApplicationRequest {
	s.ServiceTags = &v
	return s
}

func (s *DeployApplicationRequest) SetSidecarContainersConfig(v []*SidecarContainerConfig) *DeployApplicationRequest {
	s.SidecarContainersConfig = v
	return s
}

func (s *DeployApplicationRequest) SetSlsConfigs(v string) *DeployApplicationRequest {
	s.SlsConfigs = &v
	return s
}

func (s *DeployApplicationRequest) SetSlsLogEnvTags(v string) *DeployApplicationRequest {
	s.SlsLogEnvTags = &v
	return s
}

func (s *DeployApplicationRequest) SetStartupProbe(v string) *DeployApplicationRequest {
	s.StartupProbe = &v
	return s
}

func (s *DeployApplicationRequest) SetSwimlanePvtzDiscoverySvc(v string) *DeployApplicationRequest {
	s.SwimlanePvtzDiscoverySvc = &v
	return s
}

func (s *DeployApplicationRequest) SetTerminationGracePeriodSeconds(v int32) *DeployApplicationRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *DeployApplicationRequest) SetTimezone(v string) *DeployApplicationRequest {
	s.Timezone = &v
	return s
}

func (s *DeployApplicationRequest) SetTomcatConfig(v string) *DeployApplicationRequest {
	s.TomcatConfig = &v
	return s
}

func (s *DeployApplicationRequest) SetUpdateStrategy(v string) *DeployApplicationRequest {
	s.UpdateStrategy = &v
	return s
}

func (s *DeployApplicationRequest) SetVSwitchId(v string) *DeployApplicationRequest {
	s.VSwitchId = &v
	return s
}

func (s *DeployApplicationRequest) SetWarStartOptions(v string) *DeployApplicationRequest {
	s.WarStartOptions = &v
	return s
}

func (s *DeployApplicationRequest) SetWebContainer(v string) *DeployApplicationRequest {
	s.WebContainer = &v
	return s
}

func (s *DeployApplicationRequest) Validate() error {
	if s.InitContainersConfig != nil {
		for _, item := range s.InitContainersConfig {
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
	return nil
}

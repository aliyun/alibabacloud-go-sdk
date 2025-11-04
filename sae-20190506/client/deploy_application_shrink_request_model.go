// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrAssumeRoleArn(v string) *DeployApplicationShrinkRequest
	GetAcrAssumeRoleArn() *string
	SetAcrInstanceId(v string) *DeployApplicationShrinkRequest
	GetAcrInstanceId() *string
	SetAgentVersion(v string) *DeployApplicationShrinkRequest
	GetAgentVersion() *string
	SetAlbIngressReadinessGate(v string) *DeployApplicationShrinkRequest
	GetAlbIngressReadinessGate() *string
	SetAppId(v string) *DeployApplicationShrinkRequest
	GetAppId() *string
	SetAssociateEip(v bool) *DeployApplicationShrinkRequest
	GetAssociateEip() *bool
	SetAutoEnableApplicationScalingRule(v bool) *DeployApplicationShrinkRequest
	GetAutoEnableApplicationScalingRule() *bool
	SetBatchWaitTime(v int32) *DeployApplicationShrinkRequest
	GetBatchWaitTime() *int32
	SetChangeOrderDesc(v string) *DeployApplicationShrinkRequest
	GetChangeOrderDesc() *string
	SetCommand(v string) *DeployApplicationShrinkRequest
	GetCommand() *string
	SetCommandArgs(v string) *DeployApplicationShrinkRequest
	GetCommandArgs() *string
	SetConfigMapMountDesc(v string) *DeployApplicationShrinkRequest
	GetConfigMapMountDesc() *string
	SetCpu(v int32) *DeployApplicationShrinkRequest
	GetCpu() *int32
	SetCustomHostAlias(v string) *DeployApplicationShrinkRequest
	GetCustomHostAlias() *string
	SetCustomImageNetworkType(v string) *DeployApplicationShrinkRequest
	GetCustomImageNetworkType() *string
	SetDeploy(v string) *DeployApplicationShrinkRequest
	GetDeploy() *string
	SetDotnet(v string) *DeployApplicationShrinkRequest
	GetDotnet() *string
	SetEdasContainerVersion(v string) *DeployApplicationShrinkRequest
	GetEdasContainerVersion() *string
	SetEmptyDirDesc(v string) *DeployApplicationShrinkRequest
	GetEmptyDirDesc() *string
	SetEnableAhas(v string) *DeployApplicationShrinkRequest
	GetEnableAhas() *string
	SetEnableCpuBurst(v bool) *DeployApplicationShrinkRequest
	GetEnableCpuBurst() *bool
	SetEnableGreyTagRoute(v bool) *DeployApplicationShrinkRequest
	GetEnableGreyTagRoute() *bool
	SetEnableNamespaceAgentVersion(v bool) *DeployApplicationShrinkRequest
	GetEnableNamespaceAgentVersion() *bool
	SetEnableNewArms(v bool) *DeployApplicationShrinkRequest
	GetEnableNewArms() *bool
	SetEnablePrometheus(v bool) *DeployApplicationShrinkRequest
	GetEnablePrometheus() *bool
	SetEnableSidecarResourceIsolated(v bool) *DeployApplicationShrinkRequest
	GetEnableSidecarResourceIsolated() *bool
	SetEnvs(v string) *DeployApplicationShrinkRequest
	GetEnvs() *string
	SetGpuConfig(v string) *DeployApplicationShrinkRequest
	GetGpuConfig() *string
	SetHtml(v string) *DeployApplicationShrinkRequest
	GetHtml() *string
	SetImagePullSecrets(v string) *DeployApplicationShrinkRequest
	GetImagePullSecrets() *string
	SetImageUrl(v string) *DeployApplicationShrinkRequest
	GetImageUrl() *string
	SetInitContainersConfigShrink(v string) *DeployApplicationShrinkRequest
	GetInitContainersConfigShrink() *string
	SetJarStartArgs(v string) *DeployApplicationShrinkRequest
	GetJarStartArgs() *string
	SetJarStartOptions(v string) *DeployApplicationShrinkRequest
	GetJarStartOptions() *string
	SetJdk(v string) *DeployApplicationShrinkRequest
	GetJdk() *string
	SetKafkaConfigs(v string) *DeployApplicationShrinkRequest
	GetKafkaConfigs() *string
	SetLiveness(v string) *DeployApplicationShrinkRequest
	GetLiveness() *string
	SetMaxSurgeInstanceRatio(v int32) *DeployApplicationShrinkRequest
	GetMaxSurgeInstanceRatio() *int32
	SetMaxSurgeInstances(v int32) *DeployApplicationShrinkRequest
	GetMaxSurgeInstances() *int32
	SetMemory(v int32) *DeployApplicationShrinkRequest
	GetMemory() *int32
	SetMicroRegistration(v string) *DeployApplicationShrinkRequest
	GetMicroRegistration() *string
	SetMicroRegistrationConfig(v string) *DeployApplicationShrinkRequest
	GetMicroRegistrationConfig() *string
	SetMicroserviceEngineConfig(v string) *DeployApplicationShrinkRequest
	GetMicroserviceEngineConfig() *string
	SetMinReadyInstanceRatio(v int32) *DeployApplicationShrinkRequest
	GetMinReadyInstanceRatio() *int32
	SetMinReadyInstances(v int32) *DeployApplicationShrinkRequest
	GetMinReadyInstances() *int32
	SetMountDesc(v string) *DeployApplicationShrinkRequest
	GetMountDesc() *string
	SetMountHost(v string) *DeployApplicationShrinkRequest
	GetMountHost() *string
	SetNasConfigs(v string) *DeployApplicationShrinkRequest
	GetNasConfigs() *string
	SetNasId(v string) *DeployApplicationShrinkRequest
	GetNasId() *string
	SetNewSaeVersion(v string) *DeployApplicationShrinkRequest
	GetNewSaeVersion() *string
	SetOidcRoleName(v string) *DeployApplicationShrinkRequest
	GetOidcRoleName() *string
	SetOssAkId(v string) *DeployApplicationShrinkRequest
	GetOssAkId() *string
	SetOssAkSecret(v string) *DeployApplicationShrinkRequest
	GetOssAkSecret() *string
	SetOssMountDescs(v string) *DeployApplicationShrinkRequest
	GetOssMountDescs() *string
	SetPackageType(v string) *DeployApplicationShrinkRequest
	GetPackageType() *string
	SetPackageUrl(v string) *DeployApplicationShrinkRequest
	GetPackageUrl() *string
	SetPackageVersion(v string) *DeployApplicationShrinkRequest
	GetPackageVersion() *string
	SetPhp(v string) *DeployApplicationShrinkRequest
	GetPhp() *string
	SetPhpArmsConfigLocation(v string) *DeployApplicationShrinkRequest
	GetPhpArmsConfigLocation() *string
	SetPhpConfig(v string) *DeployApplicationShrinkRequest
	GetPhpConfig() *string
	SetPhpConfigLocation(v string) *DeployApplicationShrinkRequest
	GetPhpConfigLocation() *string
	SetPostStart(v string) *DeployApplicationShrinkRequest
	GetPostStart() *string
	SetPreStop(v string) *DeployApplicationShrinkRequest
	GetPreStop() *string
	SetPvtzDiscoverySvc(v string) *DeployApplicationShrinkRequest
	GetPvtzDiscoverySvc() *string
	SetPython(v string) *DeployApplicationShrinkRequest
	GetPython() *string
	SetPythonModules(v string) *DeployApplicationShrinkRequest
	GetPythonModules() *string
	SetReadiness(v string) *DeployApplicationShrinkRequest
	GetReadiness() *string
	SetReplicas(v int32) *DeployApplicationShrinkRequest
	GetReplicas() *int32
	SetSecretMountDesc(v string) *DeployApplicationShrinkRequest
	GetSecretMountDesc() *string
	SetSecurityGroupId(v string) *DeployApplicationShrinkRequest
	GetSecurityGroupId() *string
	SetServiceTags(v string) *DeployApplicationShrinkRequest
	GetServiceTags() *string
	SetSidecarContainersConfigShrink(v string) *DeployApplicationShrinkRequest
	GetSidecarContainersConfigShrink() *string
	SetSlsConfigs(v string) *DeployApplicationShrinkRequest
	GetSlsConfigs() *string
	SetStartupProbe(v string) *DeployApplicationShrinkRequest
	GetStartupProbe() *string
	SetSwimlanePvtzDiscoverySvc(v string) *DeployApplicationShrinkRequest
	GetSwimlanePvtzDiscoverySvc() *string
	SetTerminationGracePeriodSeconds(v int32) *DeployApplicationShrinkRequest
	GetTerminationGracePeriodSeconds() *int32
	SetTimezone(v string) *DeployApplicationShrinkRequest
	GetTimezone() *string
	SetTomcatConfig(v string) *DeployApplicationShrinkRequest
	GetTomcatConfig() *string
	SetUpdateStrategy(v string) *DeployApplicationShrinkRequest
	GetUpdateStrategy() *string
	SetVSwitchId(v string) *DeployApplicationShrinkRequest
	GetVSwitchId() *string
	SetWarStartOptions(v string) *DeployApplicationShrinkRequest
	GetWarStartOptions() *string
	SetWebContainer(v string) *DeployApplicationShrinkRequest
	GetWebContainer() *string
}

type DeployApplicationShrinkRequest struct {
	// The Alibaba Cloud Resource Name (ARN) required for a RAM role to obtain images across accounts. For more information, see [Grant permissions across Alibaba Cloud accounts by using a RAM role](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The ID of Container Registry Enterprise Edition instance N. This parameter is required when the **ImageUrl*	- parameter is set to the URL of an image in an ACR Enterprise Edition instance.
	//
	// example:
	//
	// cri-xxxxxx
	AcrInstanceId           *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	AgentVersion            *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	AlbIngressReadinessGate *string `json:"AlbIngressReadinessGate,omitempty" xml:"AlbIngressReadinessGate,omitempty"`
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to associate an EIP with the node pool. Take note of the following rules:
	//
	// 	- **true**: The EIP is associated with the application instance.
	//
	// 	- **false**: The EIP is not associated with the application instance.
	//
	// example:
	//
	// true
	AssociateEip *bool `json:"AssociateEip,omitempty" xml:"AssociateEip,omitempty"`
	// Specifies whether to automatically enable an auto scaling policy for the application. Take note of the following rules:
	//
	// 	- **true**: turns on Logon-free Sharing
	//
	// 	- **false**: turns off Logon-free Sharing
	//
	// example:
	//
	// true
	AutoEnableApplicationScalingRule *bool `json:"AutoEnableApplicationScalingRule,omitempty" xml:"AutoEnableApplicationScalingRule,omitempty"`
	// The interval between batches during a batch release. Unit: minutes.
	//
	// example:
	//
	// 10
	BatchWaitTime *int32 `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	// The description of the change order.
	//
	// example:
	//
	// Start the application
	ChangeOrderDesc *string `json:"ChangeOrderDesc,omitempty" xml:"ChangeOrderDesc,omitempty"`
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
	// echo
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
	// The description of the **ConfigMap*	- instance mounted to the application. Use configurations created on the Configuration Items page to configure containers. The following table describes the parameters that are used in the preceding statements.
	//
	// 	- **congfigMapId**: the ID of the ConfigMap instance. You can call the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the ID.
	//
	// 	- **key**: the key.
	//
	// > You can use `sae-sys-configmap-all` to mount all keys.
	//
	// 	- **mountPath**: the mount path in the container.
	//
	// example:
	//
	// [{"configMapId":16,"key":"test","mountPath":"/tmp"}]
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	// The CPU specifications that are required for each instance. Unit: millicores. This parameter cannot be set to 0. Valid values:
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
	// 	- **12000**
	//
	// 	- **16000**
	//
	// 	- **32000**
	//
	// example:
	//
	// 1000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The custom mappings between hostnames and IP addresses in the container. Take note of the following rules:
	//
	// 	- **hostName**: the domain name or hostname.
	//
	// 	- **ip**: the IP address.
	//
	// example:
	//
	// [{"hostName":"samplehost","ip":"127.0.0.1"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// Custom image type. To it to empty string to use pre-built image.
	//
	// - internet: Public network image
	//
	// - intranet: Private network image
	//
	// example:
	//
	// internet
	CustomImageNetworkType *string `json:"CustomImageNetworkType,omitempty" xml:"CustomImageNetworkType,omitempty"`
	// This parameter takes effect only for applications that are in the Stopped state. If you call the **DeployApplication*	- operation to manage a running application, the application is immediately redeployed.
	//
	// 	- **true*	- (default): specifies that the system immediately deploys the application, enables new configurations, and pulls application instances.
	//
	// 	- **false**: specifies that the system only enables the new configurations.
	//
	// example:
	//
	// true
	Deploy *string `json:"Deploy,omitempty" xml:"Deploy,omitempty"`
	// The version of .NET
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
	// The version of the container, such as Ali-Tomcat, in which an application developed based on High-speed Service Framework (HSF) is deployed.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	EmptyDirDesc         *string `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty"`
	// Indicates whether access to Application High Availability Service (AHAS) is enabled. Take note of the following rules:
	//
	// 	- **true**: Access to AHAS is enabled.
	//
	// 	- **false**: Access to AHAS is disabled.
	//
	// example:
	//
	// false
	EnableAhas *string `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	// Enable CPU Burst.
	//
	// true: enable
	//
	// false: disable
	//
	// example:
	//
	// true
	EnableCpuBurst *bool `json:"EnableCpuBurst,omitempty" xml:"EnableCpuBurst,omitempty"`
	// Indicates whether canary release rules are enabled. Canary release rules apply only to applications in Spring Cloud and Dubbo frameworks. Take note of the following rules:
	//
	// 	- **true**: The canary release rules are enabled.
	//
	// 	- **false**: The canary release rules are disabled.
	//
	// example:
	//
	// false
	EnableGreyTagRoute          *bool `json:"EnableGreyTagRoute,omitempty" xml:"EnableGreyTagRoute,omitempty"`
	EnableNamespaceAgentVersion *bool `json:"EnableNamespaceAgentVersion,omitempty" xml:"EnableNamespaceAgentVersion,omitempty"`
	// Enable new ARMS features.
	//
	// - true: enable
	//
	// - false: disable
	//
	// example:
	//
	// true
	EnableNewArms    *bool `json:"EnableNewArms,omitempty" xml:"EnableNewArms,omitempty"`
	EnablePrometheus *bool `json:"EnablePrometheus,omitempty" xml:"EnablePrometheus,omitempty"`
	// Enable Sidecar resource isolation.
	//
	// true: enable
	//
	// false: disable
	//
	// example:
	//
	// true
	EnableSidecarResourceIsolated *bool `json:"EnableSidecarResourceIsolated,omitempty" xml:"EnableSidecarResourceIsolated,omitempty"`
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
	// 	- Reference secret dictionary
	//
	//     	- **name**: the name of the environment variable. You can reference one or all keys. If you want to reference all keys, specify `sae-sys-secret-all-<Secret dictionary name>`. Example: `sae-sys-secret-all-test1`.
	//
	//     	- **valueFrom**: the reference of the environment variable. Set the value to `secretRef`.
	//
	//     	- **secretId**: the secret dictionary ID.
	//
	//     	- **key**: the key. If you want to reference all keys, do not configure this parameter.
	//
	// example:
	//
	// [{"name":"envtmp","value":"0"}]
	Envs      *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	GpuConfig *string `json:"GpuConfig,omitempty" xml:"GpuConfig,omitempty"`
	Html      *string `json:"Html,omitempty" xml:"Html,omitempty"`
	// The ID of the corresponding Secret.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The URL of the image. This parameter is returned only if the **PackageType*	- parameter is set to **Image**.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/sae_test/ali_sae_test:0.0.1
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Initialize container configuration.
	InitContainersConfigShrink *string `json:"InitContainersConfig,omitempty" xml:"InitContainersConfig,omitempty"`
	// The arguments in the JAR package. The arguments are used to start the application container. The default startup command is `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`.
	//
	// example:
	//
	// -Xms4G -Xmx4G
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The option settings in the JAR package. The settings are used to start the application container. The default startup command for application deployment is `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`.
	//
	// example:
	//
	// custom-option
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
	// The logging configurations of Message Queue for Apache Kafka. Take note of the following rules:
	//
	// 	- **kafkaEndpoint**: the endpoint of the Message Queue for Apache Kafka API.
	//
	// 	- **kafkaInstanceId**: the ID of the Message Queue for Apache Kafka instance.
	//
	// 	- **kafkaConfigs**: One or more logging configurations of Message Queue for Apache Kafka. For information about sample values and parameters, see the request parameter **KafkaLogfileConfig*	- in this topic.
	//
	// example:
	//
	// {"kafkaEndpoint":"10.0.X.XXX:XXXX,10.0.X.XXX:XXXX,10.0.X.XXX:XXXX\\","kafkaInstanceId":"alikafka_pre-cn-7pp2l8kr****","kafkaConfigs":[{"logType":"file_log","logDir":"/tmp/a.log","kafkaTopic":"test2"},{"logType":"stdout","logDir":"","kafkaTopic":"test"}]}
	KafkaConfigs *string `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	// The details of the availability check that was performed on the container. If the container fails this health check multiple times, the system disables and restarts the container. You can use one of the following methods to perform the health check:
	//
	// 	- Example of **exec**: `{"exec":{"command":["sh","-c","cat/home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}`
	//
	// 	- Sample code of the **httpGet*	- method: `{"httpGet":{"path":"/","port":18091,"scheme":"HTTP","isContainKeyWord":true,"keyWord":"SAE"},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// 	- Sample code of the **tcpSocket*	- method: `{"tcpSocket":{"port":18091},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// > You can use only one method to perform the health check.
	//
	// The following table describes the parameters that are used in the preceding statements.
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
	// {"exec":{"command":["sleep","5s"]},"initialDelaySeconds":10,"timeoutSeconds":11}
	Liveness              *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	MaxSurgeInstanceRatio *int32  `json:"MaxSurgeInstanceRatio,omitempty" xml:"MaxSurgeInstanceRatio,omitempty"`
	MaxSurgeInstances     *int32  `json:"MaxSurgeInstances,omitempty" xml:"MaxSurgeInstances,omitempty"`
	// The memory size that is required by each instance. Unit: MB. This parameter cannot be set to 0. The values of this parameter correspond to the values of the Cpu parameter:
	//
	// 	- This parameter is set to **1024*	- if the Cpu parameter is set to 500 or 1000.
	//
	// 	- This parameter is set to **2048*	- if the Cpu parameter is set to 500, 1000, or 2000.
	//
	// 	- This parameter is set to **4096*	- if the Cpu parameter is set to 1000, 2000, or 4000.
	//
	// 	- This parameter is set to **8192*	- if the Cpu parameter is set to 2000, 4000, or 8,000.
	//
	// 	- This parameter is set to **12288*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **16384*	- if the Cpu parameter is set to 4000, 8000, or 16000.
	//
	// 	- This parameter is set to **24576*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **32768*	- if the Cpu parameter is set to 16000.
	//
	// 	- This parameter is set to **65536*	- if the Cpu parameter is set to 8000, 16000, or 32000.
	//
	// 	- This parameter is set to **131072*	- if the Cpu parameter is set to 32000.
	//
	// example:
	//
	// 1024
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
	// Select the edition of Nacos.
	//
	// - 0: SAE built-in Nacos. Unable to get the configuration of SAE built-in Nacos.
	//
	// - 1: Self-built Nacos from users.
	//
	// - 2: MSE enterprise Nacos.
	//
	// example:
	//
	// {\\"instanceId\\":\\"mse-cn-zvp2bh6h70r\\",\\"namespace\\":\\"4c0aa74f-57cb-423c-b6af-5d9f2d0e3dbd\\"}
	MicroRegistrationConfig *string `json:"MicroRegistrationConfig,omitempty" xml:"MicroRegistrationConfig,omitempty"`
	// Configure Microservices Governance
	//
	// Whether to enable microservices governance (enable):
	//
	// - true: Enable
	//
	// - false: Disable
	//
	// Configure lossless online/offline deployment (mseLosslessRule):
	//
	// delayTime: Delay duration (unit: seconds)
	//
	// enable: Whether to enable lossless deployment
	//
	// - true: Enable
	//
	// - false: Disable
	//
	// notice: Whether to enable notifications
	//
	// - true: Enable
	//
	// - false: Disable
	//
	// warmupTime: Small-traffic warm-up duration (unit: seconds)
	//
	// example:
	//
	// {"enable": true,"mseLosslessRule": {"delayTime": 0,"enable": false,"notice": false,"warmupTime": 120}}
	MicroserviceEngineConfig *string `json:"MicroserviceEngineConfig,omitempty" xml:"MicroserviceEngineConfig,omitempty"`
	// The percentage of the minimum number of available instances. Take note of the following rules:
	//
	// 	- If you set the value to **-1**, the minimum number of available instances is not determined based on this parameter. Default value: -1.
	//
	// 	- If you set the value to a number **from 0 to 100**, the minimum number of available instances is calculated by using the following formula: Current number of instances × (Value of MinReadyInstanceRatio × 100%). The value is the nearest integer rounded up from the calculated result. For example, if the percentage is set to **50**% and five instances are available, the minimum number of available instances is 3.
	//
	// > When both **MinReadyInstance*	- and **MinReadyInstanceRatio*	- are specified and **MinReadyInstanceRatio*	- is set to a number from 0 to 100, the value of **MinReadyInstanceRatio***	- takes precedence. For example, if **MinReadyInstances*	- is set to **5, and **MinReadyInstanceRatio*	- is set to **50**, the minimum number of available instances is set to the nearest integer rounded up from the calculated result of the following formula: Current number of instances × **50%**.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of available instances. Special values:
	//
	// 	- If you set the value to **0**, business interruptions occur when the application is updated.
	//
	// 	- If you set the value to \\*\\*-1\\*\\*, the minimum number of available instances is automatically set to a system-recommended value. The value is the nearest integer to which the calculated result of the following formula is rounded up: Current number of instances × 25%. For example, if five instances are available, the minimum number of available instances is calculated by using the following formula: 5 × 25% = 1.25. In this case, the minimum number of available instances is 2.
	//
	// > Make sure that at least one instance is available during application deployment and rollback to prevent business interruptions.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// The configurations for mounting the NAS file system. After the application is created, you may want to call other operations to manage the application. If you do not want to change the NAS configurations in these subsequent operations, you can omit the **MountDesc*	- parameter in the requests. If you want to unmount the NAS file system, you must set the **MountDesc*	- values in the subsequent requests to an empty string ("").
	//
	// example:
	//
	// [{mountPath: "/tmp", nasPath: "/"}]
	MountDesc *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	// The mount target of the NAS file system in the VPC where the application is deployed. If you do not need to modify this configuration during the deployment, configure the **MountHost*	- parameter only in the first request. You do not need to include this parameter in subsequent requests. If you need to remove this configuration, leave the **MountHost*	- parameter empty in the request.
	//
	// example:
	//
	// 10d3b4bc9****.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The configurations of mounting the NAS file system. Take note of the following rules:
	//
	// 	- **mountPath**: the mount path of the container.
	//
	// 	- **readOnly**: If you set the value to **false**, the application has the read and write permissions.
	//
	// 	- **nasId**: the ID of the NAS file system.
	//
	// 	- **mountDomain**: the domain name of the mount target. For more information, see [DescribeMountTargets](https://help.aliyun.com/document_detail/62626.html).
	//
	// 	- **nasPath**: the directory in the NAS file system.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// The ID of the File Storage NAS file system. After the application is created, you may want to call other operations to manage the application. If you do not want to change the NAS configurations in these subsequent operations, you can omit the **NasId*	- parameter in the requests. If you want to unmount the NAS file system, you must set the **NasId*	- values in the subsequent requests to an empty string ("").
	//
	// example:
	//
	// 10d3b4****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// SAE edition.
	//
	// - lite: the lightweight edition.
	//
	// - std: the standard edition.
	//
	// - pro: the professional edition.
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
	// Information of the Object Storage Service (OSS) bucket mounted to the application. The following table describes the parameters that are used in the preceding statements.
	//
	// 	- **bucketName**: the name of the OSS bucket.
	//
	// 	- **bucketPath**: the directory or object in OSS. If the specified directory or object does not exist, an error is returned.
	//
	// 	- **mountPath**: the directory of the container in SAE. If the path already exists, the newly specified path overwrites the previous one. If the path does not exist, it is created.
	//
	// 	- **readOnly**: specifies whether to only allow the container path to read data from the OSS directory. Valid values:
	//
	//     	- **true**: The container path only has read permission on the OSS directory.
	//
	//     	- **false**: The application has read and write permissions.
	//
	// example:
	//
	// [{"bucketName": "oss-bucket", "bucketPath": "data/user.data", "mountPath": "/usr/data/user.data", "readOnly": true}]
	OssMountDescs *string `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty"`
	// The package type.
	//
	// When using Java, FatJar, War and Image are supported.
	//
	// When using Python, PythonZip and Image are supported.
	//
	// When using PHP, the followings are supported:
	//
	// - PhpZip
	//
	// - IMAGE_PHP_5_4
	//
	// - IMAGE_PHP_5_4_ALPINE
	//
	// - IMAGE_PHP_5_5
	//
	// - IMAGE_PHP_5_5_ALPINE
	//
	// - IMAGE_PHP_5_6
	//
	// - IMAGE_PHP_5_6_ALPINE
	//
	// - IMAGE_PHP_7_0
	//
	// - IMAGE_PHP_7_0_ALPINE
	//
	// - IMAGE_PHP_7_1
	//
	// - IMAGE_PHP_7_1_ALPINE
	//
	// - IMAGE_PHP_7_2
	//
	// - IMAGE_PHP_7_2_ALPINE
	//
	// - IMAGE_PHP_7_3
	//
	// - IMAGE_PHP_7_3_ALPINE
	//
	// example:
	//
	// FatJar
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The address of the deployment package. This parameter is required when the **PackageType*	- parameter is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// http://myoss.oss-cn-hangzhou.aliyuncs.com/my-buc/2019-06-30/****.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package. This parameter is required when the **PackageType*	- parameter is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// 1.0.1
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The dependent PHP version of PHP package. Image is not supported.
	//
	// example:
	//
	// PHP-FPM 7.0
	Php *string `json:"Php,omitempty" xml:"Php,omitempty"`
	// The path on which the PHP configuration file for application monitoring is mounted. Make sure that the PHP server loads the configuration file. SAE automatically generates the corresponding configuration file. No manual operations are required.
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
	// The script that is run immediately after the container is started. Example: `{"exec":{"command":["sh","-c","echo hello"\\]}}`
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script that is run before the container is stopped. Example: `{"exec":{"command":["sh","-c","echo hello"\\]}}`
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The configurations of Kubernetes Service-based service registration and discovery. Take note of the following rules:
	//
	// 	- **serviceName**: the name of the Alibaba Cloud service. Format: `<Custom content>-<Namespace ID>`. `-<Namespace ID>` is automatically specified based on the namespace in which an application resides and cannot be changed. For example, if you select the default namespace in the China (Beijing) region, `-cn-beijing-default` is automatically specified.
	//
	// 	- **namespaceId**: the namespace ID.
	//
	// 	- **portAndProtocol**: the port number and protocol. Valid values of the port number: 1 to 65535. Valid values of the protocol: **TCP*	- and **UDP**.
	//
	// 	- **enable**: enables the Kubernetes Service-based registration and discovery feature.
	//
	// example:
	//
	// {"serviceName":"bwm-poc-sc-gateway-cn-beijing-front","namespaceId":"cn-beijing:front","portAndProtocol":{"18012":"TCP"},"enable":true}
	PvtzDiscoverySvc *string `json:"PvtzDiscoverySvc,omitempty" xml:"PvtzDiscoverySvc,omitempty"`
	// The Python environment. Set the value to **PYTHON 3.9.15**.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// The configurations for installing custom module dependencies. By default, the dependencies defined by the requirements.txt file in the root directory are installed. If the package does not contain this file and you do not configure custom dependencies in the package, specify the dependencies that you want to install in the text box.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// The details of the health check that was performed on the container. If the container fails this health check multiple times, the system disables and restarts the container. Containers that fail health checks cannot receive traffic from Server Load Balancer (SLB) instances. You can use the **exec**, **httpGet**, or **tcpSocket*	- method to perform health checks. For more information, see the description of the **Liveness*	- parameter.
	//
	// > You can use only one method to perform the health check.
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
	// Secret Mount Description
	//
	// Use the secret dictionaries created in the Namespace Secret Dictionary page to inject information into containers. Parameter descriptions are as follows:
	//
	// - secretId: Secret instance ID. Obtain via the ListSecrets interface.
	//
	// - key: Key-value pair. Note: Set the parameter sae-sys-secret-all to mount all keys.
	//
	// - mountPath: Mount path.
	//
	// example:
	//
	// [{“secretId":10,”key":"test","mountPath":"/tmp"}]
	SecretMountDesc *string `json:"SecretMountDesc,omitempty" xml:"SecretMountDesc,omitempty"`
	// Security group ID.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The gray-release tag of the application.
	//
	// example:
	//
	// {\\"alicloud.service.tag\\":\\"g1\\"}
	ServiceTags *string `json:"ServiceTags,omitempty" xml:"ServiceTags,omitempty"`
	// The configuration of the container.
	SidecarContainersConfigShrink *string `json:"SidecarContainersConfig,omitempty" xml:"SidecarContainersConfig,omitempty"`
	// The logging configurations of Log Service.
	//
	// 	- To use Log Service resources that are automatically created by SAE, set this parameter to `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// 	- To use custom Log Service resources, set this parameter to `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// The following table describes the parameters that are used in the preceding statements.
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
	// > A Log Service project that is automatically created by SAE when you create an application is deleted when the application is deleted. Therefore, when you create an application, you cannot select a Log Service project that is automatically created by SAE for log collection.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// Check Failure: Indicates that the application failed to start. The system will report the exception and automatically restart it.
	//
	// Note:
	//
	// Supports exec, httpGet, and tcpSocket methods. For specific examples, see Liveness Parameters.
	//
	// Only one method can be selected for health checks.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	StartupProbe *string `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	// Configure K8s Service-based Service Registration/Discovery and Full-Chain Grayscale Capabilities
	//
	// - enable: Whether to enable full-link grayscale based on K8s Service (set to "true" to enable; set to "false" to disable).
	//
	// - namespaceId: Namespace ID
	//
	// - portAndProtocol: Listener port and protocol. Format: {"Port:Protocol Type":"Container Port"}
	//
	// - portProtocols: Define service ports and protocols
	//
	// port: Port
	//
	// protocol: Protocol
	//
	// targetPort: Container port
	//
	// - pvtzDiscoveryName: Service discovery name
	//
	// - serviceId: Service ID
	//
	// - serviceName: Service name
	//
	// example:
	//
	// {\\"enable\\":\\"false\\",\\"namespaceId\\":\\"cn-beijing:test\\",\\"portAndProtocol\\":{\\"2000:TCP\\":\\"18081\\"},\\"portProtocols\\":[{\\"port\\":2000,\\"protocol\\":\\"TCP\\",\\"targetPort\\":18081}],\\"pvtzDiscoveryName\\":\\"cn-beijing-1421801774382676\\",\\"serviceId\\":\\"3513\\",\\"serviceName\\":\\"demo-gray.test\\"}
	SwimlanePvtzDiscoverySvc *string `json:"SwimlanePvtzDiscoverySvc,omitempty" xml:"SwimlanePvtzDiscoverySvc,omitempty"`
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
	// The Tomcat configuration. If you want to cancel this configuration, set this parameter to "" or "{}". The following variables are included in the configuration: Take note of the following rules:
	//
	// 	- **port**: the port number. The port number ranges from 1024 to 65535. Though the admin permissions are configured for the container, the root permissions are required to perform operations on ports whose number is smaller than 1024. Enter a value that ranges from 1025 to 65535 because the container has only the admin permissions. If you do not specify this parameter, the default port number 8080 is used.
	//
	// 	- **contextPath**: the path. Default value: /. This value indicates the root directory.
	//
	// 	- **maxThreads**: the maximum number of connections in the connection pool. Default value: 400.
	//
	// 	- **uriEncoding**: the URI encoding scheme in the Tomcat container. Valid values: UTF-8, ISO-8859-1, GBK, and GB2312.***********	- If you do not specify this parameter, the default value **ISO-8859-1*	- is used.
	//
	// 	- **useBodyEncoding**: specifies whether to use the encoding scheme specified in the request body for URI query parameters. Default value: true.
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
	// The following table describes the parameters that are used in the preceding statements.
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
	// The ID of the vSwitch, where the EIP of the application instances resides. The vSwitch must reside in the VPC above.
	//
	// example:
	//
	// vsw-bp12mw1f8k3jgygk9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The startup command of the WAR package. For information about how to configure the startup command, see [Configure startup commands](https://help.aliyun.com/document_detail/96677.html).
	//
	// example:
	//
	// CATALINA_OPTS=\\"$CATALINA_OPTS $Options\\" catalina.sh run
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

func (s DeployApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeployApplicationShrinkRequest) GetAcrAssumeRoleArn() *string {
	return s.AcrAssumeRoleArn
}

func (s *DeployApplicationShrinkRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *DeployApplicationShrinkRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *DeployApplicationShrinkRequest) GetAlbIngressReadinessGate() *string {
	return s.AlbIngressReadinessGate
}

func (s *DeployApplicationShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeployApplicationShrinkRequest) GetAssociateEip() *bool {
	return s.AssociateEip
}

func (s *DeployApplicationShrinkRequest) GetAutoEnableApplicationScalingRule() *bool {
	return s.AutoEnableApplicationScalingRule
}

func (s *DeployApplicationShrinkRequest) GetBatchWaitTime() *int32 {
	return s.BatchWaitTime
}

func (s *DeployApplicationShrinkRequest) GetChangeOrderDesc() *string {
	return s.ChangeOrderDesc
}

func (s *DeployApplicationShrinkRequest) GetCommand() *string {
	return s.Command
}

func (s *DeployApplicationShrinkRequest) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *DeployApplicationShrinkRequest) GetConfigMapMountDesc() *string {
	return s.ConfigMapMountDesc
}

func (s *DeployApplicationShrinkRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *DeployApplicationShrinkRequest) GetCustomHostAlias() *string {
	return s.CustomHostAlias
}

func (s *DeployApplicationShrinkRequest) GetCustomImageNetworkType() *string {
	return s.CustomImageNetworkType
}

func (s *DeployApplicationShrinkRequest) GetDeploy() *string {
	return s.Deploy
}

func (s *DeployApplicationShrinkRequest) GetDotnet() *string {
	return s.Dotnet
}

func (s *DeployApplicationShrinkRequest) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *DeployApplicationShrinkRequest) GetEmptyDirDesc() *string {
	return s.EmptyDirDesc
}

func (s *DeployApplicationShrinkRequest) GetEnableAhas() *string {
	return s.EnableAhas
}

func (s *DeployApplicationShrinkRequest) GetEnableCpuBurst() *bool {
	return s.EnableCpuBurst
}

func (s *DeployApplicationShrinkRequest) GetEnableGreyTagRoute() *bool {
	return s.EnableGreyTagRoute
}

func (s *DeployApplicationShrinkRequest) GetEnableNamespaceAgentVersion() *bool {
	return s.EnableNamespaceAgentVersion
}

func (s *DeployApplicationShrinkRequest) GetEnableNewArms() *bool {
	return s.EnableNewArms
}

func (s *DeployApplicationShrinkRequest) GetEnablePrometheus() *bool {
	return s.EnablePrometheus
}

func (s *DeployApplicationShrinkRequest) GetEnableSidecarResourceIsolated() *bool {
	return s.EnableSidecarResourceIsolated
}

func (s *DeployApplicationShrinkRequest) GetEnvs() *string {
	return s.Envs
}

func (s *DeployApplicationShrinkRequest) GetGpuConfig() *string {
	return s.GpuConfig
}

func (s *DeployApplicationShrinkRequest) GetHtml() *string {
	return s.Html
}

func (s *DeployApplicationShrinkRequest) GetImagePullSecrets() *string {
	return s.ImagePullSecrets
}

func (s *DeployApplicationShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DeployApplicationShrinkRequest) GetInitContainersConfigShrink() *string {
	return s.InitContainersConfigShrink
}

func (s *DeployApplicationShrinkRequest) GetJarStartArgs() *string {
	return s.JarStartArgs
}

func (s *DeployApplicationShrinkRequest) GetJarStartOptions() *string {
	return s.JarStartOptions
}

func (s *DeployApplicationShrinkRequest) GetJdk() *string {
	return s.Jdk
}

func (s *DeployApplicationShrinkRequest) GetKafkaConfigs() *string {
	return s.KafkaConfigs
}

func (s *DeployApplicationShrinkRequest) GetLiveness() *string {
	return s.Liveness
}

func (s *DeployApplicationShrinkRequest) GetMaxSurgeInstanceRatio() *int32 {
	return s.MaxSurgeInstanceRatio
}

func (s *DeployApplicationShrinkRequest) GetMaxSurgeInstances() *int32 {
	return s.MaxSurgeInstances
}

func (s *DeployApplicationShrinkRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *DeployApplicationShrinkRequest) GetMicroRegistration() *string {
	return s.MicroRegistration
}

func (s *DeployApplicationShrinkRequest) GetMicroRegistrationConfig() *string {
	return s.MicroRegistrationConfig
}

func (s *DeployApplicationShrinkRequest) GetMicroserviceEngineConfig() *string {
	return s.MicroserviceEngineConfig
}

func (s *DeployApplicationShrinkRequest) GetMinReadyInstanceRatio() *int32 {
	return s.MinReadyInstanceRatio
}

func (s *DeployApplicationShrinkRequest) GetMinReadyInstances() *int32 {
	return s.MinReadyInstances
}

func (s *DeployApplicationShrinkRequest) GetMountDesc() *string {
	return s.MountDesc
}

func (s *DeployApplicationShrinkRequest) GetMountHost() *string {
	return s.MountHost
}

func (s *DeployApplicationShrinkRequest) GetNasConfigs() *string {
	return s.NasConfigs
}

func (s *DeployApplicationShrinkRequest) GetNasId() *string {
	return s.NasId
}

func (s *DeployApplicationShrinkRequest) GetNewSaeVersion() *string {
	return s.NewSaeVersion
}

func (s *DeployApplicationShrinkRequest) GetOidcRoleName() *string {
	return s.OidcRoleName
}

func (s *DeployApplicationShrinkRequest) GetOssAkId() *string {
	return s.OssAkId
}

func (s *DeployApplicationShrinkRequest) GetOssAkSecret() *string {
	return s.OssAkSecret
}

func (s *DeployApplicationShrinkRequest) GetOssMountDescs() *string {
	return s.OssMountDescs
}

func (s *DeployApplicationShrinkRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *DeployApplicationShrinkRequest) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *DeployApplicationShrinkRequest) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *DeployApplicationShrinkRequest) GetPhp() *string {
	return s.Php
}

func (s *DeployApplicationShrinkRequest) GetPhpArmsConfigLocation() *string {
	return s.PhpArmsConfigLocation
}

func (s *DeployApplicationShrinkRequest) GetPhpConfig() *string {
	return s.PhpConfig
}

func (s *DeployApplicationShrinkRequest) GetPhpConfigLocation() *string {
	return s.PhpConfigLocation
}

func (s *DeployApplicationShrinkRequest) GetPostStart() *string {
	return s.PostStart
}

func (s *DeployApplicationShrinkRequest) GetPreStop() *string {
	return s.PreStop
}

func (s *DeployApplicationShrinkRequest) GetPvtzDiscoverySvc() *string {
	return s.PvtzDiscoverySvc
}

func (s *DeployApplicationShrinkRequest) GetPython() *string {
	return s.Python
}

func (s *DeployApplicationShrinkRequest) GetPythonModules() *string {
	return s.PythonModules
}

func (s *DeployApplicationShrinkRequest) GetReadiness() *string {
	return s.Readiness
}

func (s *DeployApplicationShrinkRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *DeployApplicationShrinkRequest) GetSecretMountDesc() *string {
	return s.SecretMountDesc
}

func (s *DeployApplicationShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeployApplicationShrinkRequest) GetServiceTags() *string {
	return s.ServiceTags
}

func (s *DeployApplicationShrinkRequest) GetSidecarContainersConfigShrink() *string {
	return s.SidecarContainersConfigShrink
}

func (s *DeployApplicationShrinkRequest) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *DeployApplicationShrinkRequest) GetStartupProbe() *string {
	return s.StartupProbe
}

func (s *DeployApplicationShrinkRequest) GetSwimlanePvtzDiscoverySvc() *string {
	return s.SwimlanePvtzDiscoverySvc
}

func (s *DeployApplicationShrinkRequest) GetTerminationGracePeriodSeconds() *int32 {
	return s.TerminationGracePeriodSeconds
}

func (s *DeployApplicationShrinkRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *DeployApplicationShrinkRequest) GetTomcatConfig() *string {
	return s.TomcatConfig
}

func (s *DeployApplicationShrinkRequest) GetUpdateStrategy() *string {
	return s.UpdateStrategy
}

func (s *DeployApplicationShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DeployApplicationShrinkRequest) GetWarStartOptions() *string {
	return s.WarStartOptions
}

func (s *DeployApplicationShrinkRequest) GetWebContainer() *string {
	return s.WebContainer
}

func (s *DeployApplicationShrinkRequest) SetAcrAssumeRoleArn(v string) *DeployApplicationShrinkRequest {
	s.AcrAssumeRoleArn = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetAcrInstanceId(v string) *DeployApplicationShrinkRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetAgentVersion(v string) *DeployApplicationShrinkRequest {
	s.AgentVersion = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetAlbIngressReadinessGate(v string) *DeployApplicationShrinkRequest {
	s.AlbIngressReadinessGate = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetAppId(v string) *DeployApplicationShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetAssociateEip(v bool) *DeployApplicationShrinkRequest {
	s.AssociateEip = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetAutoEnableApplicationScalingRule(v bool) *DeployApplicationShrinkRequest {
	s.AutoEnableApplicationScalingRule = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetBatchWaitTime(v int32) *DeployApplicationShrinkRequest {
	s.BatchWaitTime = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetChangeOrderDesc(v string) *DeployApplicationShrinkRequest {
	s.ChangeOrderDesc = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetCommand(v string) *DeployApplicationShrinkRequest {
	s.Command = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetCommandArgs(v string) *DeployApplicationShrinkRequest {
	s.CommandArgs = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetConfigMapMountDesc(v string) *DeployApplicationShrinkRequest {
	s.ConfigMapMountDesc = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetCpu(v int32) *DeployApplicationShrinkRequest {
	s.Cpu = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetCustomHostAlias(v string) *DeployApplicationShrinkRequest {
	s.CustomHostAlias = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetCustomImageNetworkType(v string) *DeployApplicationShrinkRequest {
	s.CustomImageNetworkType = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetDeploy(v string) *DeployApplicationShrinkRequest {
	s.Deploy = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetDotnet(v string) *DeployApplicationShrinkRequest {
	s.Dotnet = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetEdasContainerVersion(v string) *DeployApplicationShrinkRequest {
	s.EdasContainerVersion = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetEmptyDirDesc(v string) *DeployApplicationShrinkRequest {
	s.EmptyDirDesc = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetEnableAhas(v string) *DeployApplicationShrinkRequest {
	s.EnableAhas = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetEnableCpuBurst(v bool) *DeployApplicationShrinkRequest {
	s.EnableCpuBurst = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetEnableGreyTagRoute(v bool) *DeployApplicationShrinkRequest {
	s.EnableGreyTagRoute = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetEnableNamespaceAgentVersion(v bool) *DeployApplicationShrinkRequest {
	s.EnableNamespaceAgentVersion = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetEnableNewArms(v bool) *DeployApplicationShrinkRequest {
	s.EnableNewArms = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetEnablePrometheus(v bool) *DeployApplicationShrinkRequest {
	s.EnablePrometheus = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetEnableSidecarResourceIsolated(v bool) *DeployApplicationShrinkRequest {
	s.EnableSidecarResourceIsolated = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetEnvs(v string) *DeployApplicationShrinkRequest {
	s.Envs = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetGpuConfig(v string) *DeployApplicationShrinkRequest {
	s.GpuConfig = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetHtml(v string) *DeployApplicationShrinkRequest {
	s.Html = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetImagePullSecrets(v string) *DeployApplicationShrinkRequest {
	s.ImagePullSecrets = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetImageUrl(v string) *DeployApplicationShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetInitContainersConfigShrink(v string) *DeployApplicationShrinkRequest {
	s.InitContainersConfigShrink = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetJarStartArgs(v string) *DeployApplicationShrinkRequest {
	s.JarStartArgs = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetJarStartOptions(v string) *DeployApplicationShrinkRequest {
	s.JarStartOptions = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetJdk(v string) *DeployApplicationShrinkRequest {
	s.Jdk = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetKafkaConfigs(v string) *DeployApplicationShrinkRequest {
	s.KafkaConfigs = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetLiveness(v string) *DeployApplicationShrinkRequest {
	s.Liveness = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetMaxSurgeInstanceRatio(v int32) *DeployApplicationShrinkRequest {
	s.MaxSurgeInstanceRatio = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetMaxSurgeInstances(v int32) *DeployApplicationShrinkRequest {
	s.MaxSurgeInstances = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetMemory(v int32) *DeployApplicationShrinkRequest {
	s.Memory = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetMicroRegistration(v string) *DeployApplicationShrinkRequest {
	s.MicroRegistration = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetMicroRegistrationConfig(v string) *DeployApplicationShrinkRequest {
	s.MicroRegistrationConfig = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetMicroserviceEngineConfig(v string) *DeployApplicationShrinkRequest {
	s.MicroserviceEngineConfig = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetMinReadyInstanceRatio(v int32) *DeployApplicationShrinkRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetMinReadyInstances(v int32) *DeployApplicationShrinkRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetMountDesc(v string) *DeployApplicationShrinkRequest {
	s.MountDesc = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetMountHost(v string) *DeployApplicationShrinkRequest {
	s.MountHost = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetNasConfigs(v string) *DeployApplicationShrinkRequest {
	s.NasConfigs = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetNasId(v string) *DeployApplicationShrinkRequest {
	s.NasId = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetNewSaeVersion(v string) *DeployApplicationShrinkRequest {
	s.NewSaeVersion = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetOidcRoleName(v string) *DeployApplicationShrinkRequest {
	s.OidcRoleName = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetOssAkId(v string) *DeployApplicationShrinkRequest {
	s.OssAkId = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetOssAkSecret(v string) *DeployApplicationShrinkRequest {
	s.OssAkSecret = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetOssMountDescs(v string) *DeployApplicationShrinkRequest {
	s.OssMountDescs = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPackageType(v string) *DeployApplicationShrinkRequest {
	s.PackageType = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPackageUrl(v string) *DeployApplicationShrinkRequest {
	s.PackageUrl = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPackageVersion(v string) *DeployApplicationShrinkRequest {
	s.PackageVersion = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPhp(v string) *DeployApplicationShrinkRequest {
	s.Php = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPhpArmsConfigLocation(v string) *DeployApplicationShrinkRequest {
	s.PhpArmsConfigLocation = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPhpConfig(v string) *DeployApplicationShrinkRequest {
	s.PhpConfig = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPhpConfigLocation(v string) *DeployApplicationShrinkRequest {
	s.PhpConfigLocation = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPostStart(v string) *DeployApplicationShrinkRequest {
	s.PostStart = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPreStop(v string) *DeployApplicationShrinkRequest {
	s.PreStop = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPvtzDiscoverySvc(v string) *DeployApplicationShrinkRequest {
	s.PvtzDiscoverySvc = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPython(v string) *DeployApplicationShrinkRequest {
	s.Python = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetPythonModules(v string) *DeployApplicationShrinkRequest {
	s.PythonModules = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetReadiness(v string) *DeployApplicationShrinkRequest {
	s.Readiness = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetReplicas(v int32) *DeployApplicationShrinkRequest {
	s.Replicas = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetSecretMountDesc(v string) *DeployApplicationShrinkRequest {
	s.SecretMountDesc = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetSecurityGroupId(v string) *DeployApplicationShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetServiceTags(v string) *DeployApplicationShrinkRequest {
	s.ServiceTags = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetSidecarContainersConfigShrink(v string) *DeployApplicationShrinkRequest {
	s.SidecarContainersConfigShrink = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetSlsConfigs(v string) *DeployApplicationShrinkRequest {
	s.SlsConfigs = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetStartupProbe(v string) *DeployApplicationShrinkRequest {
	s.StartupProbe = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetSwimlanePvtzDiscoverySvc(v string) *DeployApplicationShrinkRequest {
	s.SwimlanePvtzDiscoverySvc = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetTerminationGracePeriodSeconds(v int32) *DeployApplicationShrinkRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetTimezone(v string) *DeployApplicationShrinkRequest {
	s.Timezone = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetTomcatConfig(v string) *DeployApplicationShrinkRequest {
	s.TomcatConfig = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetUpdateStrategy(v string) *DeployApplicationShrinkRequest {
	s.UpdateStrategy = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetVSwitchId(v string) *DeployApplicationShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetWarStartOptions(v string) *DeployApplicationShrinkRequest {
	s.WarStartOptions = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetWebContainer(v string) *DeployApplicationShrinkRequest {
	s.WebContainer = &v
	return s
}

func (s *DeployApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}

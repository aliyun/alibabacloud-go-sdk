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
	SetLabelsShrink(v string) *DeployApplicationShrinkRequest
	GetLabelsShrink() *string
	SetLiveness(v string) *DeployApplicationShrinkRequest
	GetLiveness() *string
	SetLokiConfigs(v string) *DeployApplicationShrinkRequest
	GetLokiConfigs() *string
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
	SetRaspConfigShrink(v string) *DeployApplicationShrinkRequest
	GetRaspConfigShrink() *string
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
	SetSlsLogEnvTags(v string) *DeployApplicationShrinkRequest
	GetSlsLogEnvTags() *string
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
	// The ARN of the RAM role required for cross-account image pulling. For more information, see [Grant permissions across Alibaba Cloud accounts by using a RAM role](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The Container Registry Enterprise instance ID. This parameter is required when **ImageUrl*	- is set to a Container Registry Enterprise instance image.
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
	// The ALB gateway ReadinessGate configuration.
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
	// Specifies whether to associate an elastic IP address (EIP). Valid values:
	//
	// - **true**: associates an EIP.
	//
	// - **false**: does not associate an EIP.
	//
	// example:
	//
	// true
	AssociateEip *bool `json:"AssociateEip,omitempty" xml:"AssociateEip,omitempty"`
	// Specifies whether to automatically enable the application elastic scaling policy. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	AutoEnableApplicationScalingRule *bool `json:"AutoEnableApplicationScalingRule,omitempty" xml:"AutoEnableApplicationScalingRule,omitempty"`
	// The interval between deployment batches. Unit: seconds.
	//
	// example:
	//
	// 10
	BatchWaitTime *int32 `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	// The description of the change order.
	//
	// example:
	//
	// Start application
	ChangeOrderDesc *string `json:"ChangeOrderDesc,omitempty" xml:"ChangeOrderDesc,omitempty"`
	// The image startup command. The command must be an executable object in the container. Example:
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
	// The arguments of the image startup command. These are the arguments required by the startup command specified in **Command**. Format:
	//
	// `["a","b"]`
	//
	// In the preceding example, `CommandArgs=["abc", ">", "file0"]`, where `["abc", ">", "file0"]` must be converted to a String type, and the internal format is a JSON array. If this parameter is not required, leave it empty.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The **ConfigMap*	- mount description. Use a ConfigMap created on the namespace configuration items page to inject configuration information into the container. Settings:
	//
	// - **configMapId**: the ConfigMap instance ID. You can obtain it by calling the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation.
	//
	// - **key**: the key.
	//
	// > You can mount all keys by passing the `sae-sys-configmap-all` parameter.
	//
	// - **mountPath**: the mount path.
	//
	// example:
	//
	// [{"configMapId":16,"key":"test","mountPath":"/tmp"}]
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	// The CPU resources required for each instance. Unit: millicores. This parameter cannot be set to 0. Only the following defined specifications are supported:
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
	// The custom host mapping in the container. Valid values:
	//
	// - **hostName**: the domain name or hostname.
	//
	// - **ip**: the IP address.
	//
	// example:
	//
	// [{"hostName":"samplehost","ip":"127.0.0.1"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The custom image type. If the image is not a custom image, set this parameter to an empty string:
	//
	// - internet: public image.
	//
	// - intranet: private image.
	//
	// example:
	//
	// internet
	CustomImageNetworkType *string `json:"CustomImageNetworkType,omitempty" xml:"CustomImageNetworkType,omitempty"`
	// This parameter takes effect only for applications in the stopped state. If you call the **DeployApplication*	- operation for a running application, the application is immediately redeployed.
	//
	// - **true**: default value. Deploys immediately, applies the new deployment configuration, and starts instances.
	//
	// - **false**: applies the new deployment configuration only, without starting application instances.
	//
	// example:
	//
	// true
	Deploy *string `json:"Deploy,omitempty" xml:"Deploy,omitempty"`
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
	// .NET 3.1
	Dotnet *string `json:"Dotnet,omitempty" xml:"Dotnet,omitempty"`
	// The version of the application runtime environment in the HSF framework, such as the Ali-Tomcat container.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// The shared ephemeral storage configuration.
	//
	// example:
	//
	// [{\\"name\\":\\"workdir\\",\\"mountPath\\":\\"/usr/local/tomcat/webapps\\"}]
	EmptyDirDesc *string `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty"`
	// Specifies whether to enable Application High Availability Service (AHAS). Valid values:
	//
	// - **true**: enables AHAS.
	//
	// - **false**: does not enable AHAS.
	//
	// example:
	//
	// false
	EnableAhas *string `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	// Specifies whether to enable the CPU Burst feature:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// example:
	//
	// true
	EnableCpuBurst *bool `json:"EnableCpuBurst,omitempty" xml:"EnableCpuBurst,omitempty"`
	// Specifies whether to enable traffic canary release rules. This rule applies only to applications that use the Spring Cloud or Dubbo framework. Valid values:
	//
	// - **true**: enables canary release rules.
	//
	// - **false**: disables canary release rules.
	//
	// example:
	//
	// false
	EnableGreyTagRoute *bool `json:"EnableGreyTagRoute,omitempty" xml:"EnableGreyTagRoute,omitempty"`
	// Specifies whether to reuse the namespace Agent version configuration.
	//
	// example:
	//
	// true
	EnableNamespaceAgentVersion *bool `json:"EnableNamespaceAgentVersion,omitempty" xml:"EnableNamespaceAgentVersion,omitempty"`
	// Specifies whether to enable the new ARMS feature:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// example:
	//
	// true
	EnableNewArms *bool `json:"EnableNewArms,omitempty" xml:"EnableNewArms,omitempty"`
	// Specifies whether to enable custom Prometheus metric collection.
	//
	// example:
	//
	// false
	EnablePrometheus *bool `json:"EnablePrometheus,omitempty" xml:"EnablePrometheus,omitempty"`
	// Specifies whether to enable sidecar resource isolation:
	//
	// - true: enables isolation.
	//
	// - false: does not enable isolation.
	//
	// example:
	//
	// true
	EnableSidecarResourceIsolated *bool `json:"EnableSidecarResourceIsolated,omitempty" xml:"EnableSidecarResourceIsolated,omitempty"`
	// The container environment variable parameters. You can customize environment variables or reference ConfigMap items. To reference a ConfigMap item, create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//     - **name**: the name of the environment variable.
	//
	//     - **value**: the value of the environment variable. This takes priority over valueFrom.
	//
	// - Reference a ConfigMap item (valueFrom)
	//
	//     - **name**: the name of the environment variable. To reference all keys, enter `sae-sys-configmap-all-<ConfigMap name>`, such as `sae-sys-configmap-all-test1`.
	//
	//     - **valueFrom**: the environment variable reference. Set the value to `configMapRef`.
	//
	//     - **configMapId**: the ConfigMap ID.
	//
	//     - **key**: the key. Do not set this field if you want to reference all keys.
	//
	// - Reference a secret (valueFrom)
	//
	//     - **name**: the name of the environment variable. To reference all keys, enter `sae-sys-secret-all-<secret name>`, such as `sae-sys-secret-all-test1`.
	//
	//     - **valueFrom**: the environment variable reference. Set the value to `secretRef`.
	//
	//     - **secretId**: the secret ID.
	//
	//     - **key**: the key. Do not set this field if you want to reference all keys.
	//
	// example:
	//
	// [ { "name": "sae-sys-configmap-all-hello", "valueFrom": { "configMapRef": { "configMapId": 100, "key": "" } } }, { "name": "hello", "valueFrom": { "configMapRef": { "configMapId": 101, "key": "php-fpm" } } }, { "name": "sae-sys-secret-all-hello", "valueFrom": { “secretRef": { “secretId": 100, "key": "" } } }, { "name": “password”, "valueFrom": { “secretRef": { “secretId": 101, "key": “password” } } }, { "name": "envtmp", "value": "newenv" } ]
	Envs      *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	GpuConfig *string `json:"GpuConfig,omitempty" xml:"GpuConfig,omitempty"`
	// The Nginx version.
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
	// registry.cn-hangzhou.aliyuncs.com/sae_test/ali_sae_test:0.0.1
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The init container configuration.
	InitContainersConfigShrink *string `json:"InitContainersConfig,omitempty" xml:"InitContainersConfig,omitempty"`
	// The startup arguments for a JAR package-based application. The default startup command is `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`.
	//
	// example:
	//
	// -Xms4G -Xmx4G
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The startup options for a JAR package-based application. The default startup command is `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`.
	//
	// example:
	//
	// custom-option
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
	// - **kafkaEndpoint**: the endpoint of the Kafka API.
	//
	// - **kafkaInstanceId**: the Kafka instance ID.
	//
	// - **kafkaConfigs**: the configuration summary for one or more log entries. For example values and parameter descriptions, see the **kafkaConfigs*	- request parameter in this topic.
	//
	// example:
	//
	// {"kafkaEndpoint":"10.0.X.XXX:XXXX,10.0.X.XXX:XXXX,10.0.X.XXX:XXXX","kafkaInstanceId":"alikafka_pre-cn-7pp2l8kr****","kafkaConfigs":[{"logType":"file_log","logDir":"/tmp/a.log","kafkaTopic":"test2"},{"logType":"stdout","logDir":"","kafkaTopic":"test"}]}
	KafkaConfigs *string `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The container health check. Containers that fail the health check are shutdown and recovered. The following methods are supported:
	//
	// - **exec**: for example, `{"exec":{"command":["sh","-c","cat/home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}`
	//
	// - **httpGet**: for example, `{"httpGet":{"path":"/","port":18091,"scheme":"HTTP","isContainKeyWord":true,"keyWord":"SAE"},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// - **tcpSocket**: for example, `{"tcpSocket":{"port":18091},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// > Only one method can be selected for health checks.
	//
	// Settings:
	//
	// - **exec.command**: sets the health check command.
	//
	// - **httpGet.path**: the access path.
	//
	// - **httpGet.scheme**: **HTTP*	- or **HTTPS**.
	//
	// - **httpGet.isContainKeyWord**: **true*	- indicates that the keyword is included, **false*	- indicates that the keyword is not included, and the absence of this field indicates that the advanced feature is not used.
	//
	// - **httpGet.keyWord**: the custom keyword. The **isContainKeyWord*	- field must be present when this parameter is used.
	//
	// - **tcpSocket.port**: the port for TCP connection detection.
	//
	// - **initialDelaySeconds**: sets the initial delay for the health check. Default value: 10. Unit: seconds.
	//
	// - **periodSeconds**: sets the health check epoch. Default value: 30. Unit: seconds.
	//
	// - **timeoutSeconds**: sets the health check timeout period. Default value: 1. Unit: seconds. If this parameter is set to 0 or is not set, the default timeout period is 1 second.
	//
	// example:
	//
	// {"exec":{"command":["sleep","5s"]},"initialDelaySeconds":10,"timeoutSeconds":11}
	Liveness    *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	LokiConfigs *string `json:"LokiConfigs,omitempty" xml:"LokiConfigs,omitempty"`
	// The maximum surge instance percentage. Valid values:
	//
	// If the minimum number of available instances is 100%, the maximum surge cannot be set to 0. If this parameter is set to **-1**, the system-recommended value of 30% is used, which is 30% of the current number of instances. For example, if the current number of instances is 10, the value is 10 × 30% = 3.
	//
	// example:
	//
	// -1
	MaxSurgeInstanceRatio *int32 `json:"MaxSurgeInstanceRatio,omitempty" xml:"MaxSurgeInstanceRatio,omitempty"`
	// The maximum number of surge instances. Valid values:
	//
	// If the minimum number of available instances is 100%, the maximum surge cannot be set to 0.
	//
	// If this parameter is set to **-1**, the system-recommended value of 30% is used, which is 30% of the current number of instances. For example, if the current number of instances is 10, the value is 10 × 30% = 3.
	//
	// example:
	//
	// -1
	MaxSurgeInstances *int32 `json:"MaxSurgeInstances,omitempty" xml:"MaxSurgeInstances,omitempty"`
	// The memory required for each instance. Unit: MB. This parameter cannot be set to 0. The memory has a one-to-one mapping with CPU. Only the following defined specifications are supported:
	//
	// - **1024**: corresponds to 500 and 1000 millicores of CPU.
	//
	// - **2048**: corresponds to 500, 1000, and 2000 millicores of CPU.
	//
	// - **4096**: corresponds to 1000, 2000, and 4000 millicores of CPU.
	//
	// - **8192**: corresponds to 2000, 4000, and 8000 millicores of CPU.
	//
	// - **12288**: corresponds to 12000 millicores of CPU.
	//
	// - **16384**: corresponds to 4000, 8000, and 16000 millicores of CPU.
	//
	// - **24576**: corresponds to 12000 millicores of CPU.
	//
	// - **32768**: corresponds to 16000 millicores of CPU.
	//
	// - **65536**: corresponds to 8000, 16000, and 32000 millicores of CPU.
	//
	// - **131072**: corresponds to 32000 millicores of CPU.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Specifies the Nacos registry. Valid values:
	//
	// - **0**: Serverless App Engine built-in Nacos.
	//
	// - **1**: self-managed Nacos.
	//
	// - **2**: Microservices Engine (MSE) commercial edition Nacos.
	//
	// > If you select Serverless App Engine built-in Nacos, you cannot obtain the configuration of the built-in Nacos.
	//
	// example:
	//
	// "0"
	MicroRegistration *string `json:"MicroRegistration,omitempty" xml:"MicroRegistration,omitempty"`
	// The registry configuration, which takes effect only when the registry type is MSE Nacos Enterprise Edition.
	//
	// example:
	//
	// {\\"instanceId\\":\\"mse-cn-zvp2bh6h70r\\",\\"namespace\\":\\"4c0aa74f-57cb-423c-b6af-5d9f2d0e3dbd\\"}
	MicroRegistrationConfig *string `json:"MicroRegistrationConfig,omitempty" xml:"MicroRegistrationConfig,omitempty"`
	// Configures microservice governance.
	//
	// - Specifies whether to enable microservice governance (enable):
	//
	//    - true: enabled.
	//
	//   - false: disabled.
	//
	// - Configures lossless online/offline (mseLosslessRule):
	//
	//   - delayTime: the delay time.
	//
	//   - enable: specifies whether to enable the lossless online feature. true indicates enabled, and false indicates disabled.
	//
	//   - notice: specifies whether to enable the notification feature. true indicates enabled, and false indicates disabled.
	//
	//   - warmupTime: the warm-up duration for low-traffic scenarios. Unit: seconds.
	//
	// example:
	//
	// {"enable": true,"mseLosslessRule": {"delayTime": 0,"enable": false,"notice": false,"warmupTime": 120}}
	MicroserviceEngineConfig *string `json:"MicroserviceEngineConfig,omitempty" xml:"MicroserviceEngineConfig,omitempty"`
	// The minimum percentage of available instances. Valid values:
	//
	//  - **-1**: the initialization value, which indicates that the percentage is not used.
	//
	//  - **0~100**: the unit is percentage, rounded up. For example, if this parameter is set to **50**% and the current number of instances is 5, the minimum number of available instances is 3.
	//
	// > When both **MinReadyInstance*	- and **MinReadyInstanceRatio*	- are specified and the value of **MinReadyInstanceRatio*	- is not **-1**, the **MinReadyInstanceRatio*	- parameter takes precedence. For example, if **MinReadyInstances*	- is set to **5*	- and **MinReadyInstanceRatio*	- is set to **50**, the value **50*	- is used to calculate the minimum number of available instances.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of available instances. Valid values:
	//
	// - If this parameter is set to **0**, service is interrupted during the upgrade.
	//
	// - If this parameter is set to **-1**, the system-recommended value is used, which is 25% of the current number of instances. If the current number of instances is 5, 5 × 25% = 1.25, which is rounded up to 2.
	//
	// > The minimum number of available instances for each rolling deployment should be ≥ 1 to prevent service interruption.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// We do not recommend that you configure this field. Configure **NasConfigs*	- instead. The NAS mount description. If the configuration has not changed during deployment, you do not need to set this parameter (that is, the **MountDesc*	- field does not need to be included in the request). To clear the NAS configuration, set the value of this field to an empty string (that is, set the value of the **MountDesc*	- field to "" in the request).
	//
	// example:
	//
	// [{mountPath: "/tmp", nasPath: "/"}]
	MountDesc *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	// We do not recommend that you configure this field. Configure **NasConfigs*	- instead. The mount point of the NAS file system in the VPC of the application. If the configuration has not changed during deployment, you do not need to set this parameter (that is, the **MountHost*	- field does not need to be included in the request). To clear the NAS configuration, set the value of this field to an empty string (that is, set the value of the **MountHost*	- field to "" in the request).
	//
	// example:
	//
	// 10d3b4bc9****.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The NAS mount configuration. Valid values:
	//
	// - **mountPath**: the container mount path.
	//
	// - **readOnly**: set to **false*	- for read and write permission.
	//
	// - **nasId**: the NAS ID.
	//
	// - **mountDomain**: the container mount point address. For more information, see [DescribeMountTargets](https://help.aliyun.com/document_detail/62626.html).
	//
	// - **nasPath**: the NAS relative file directory.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// We do not recommend that you configure this field. Configure **NasConfigs*	- instead. The ID of the NAS file system. If the configuration has not changed during deployment, you do not need to set this parameter (that is, the **NasId*	- field does not need to be included in the request). To clear the NAS configuration, set the value of this field to an empty string (that is, set the value of the **NasId*	- field to "" in the request).
	//
	// example:
	//
	// 10d3b4****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The application version:
	//
	// - lite: lite edition.
	//
	// - std: standard edition.
	//
	// - pro: professional edition.
	//
	// example:
	//
	// pro
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	// The RAM role for identity authentication.
	//
	//
	// > Create an OIDC identity provider and an identity provider role in the same region in advance. For more information, see [Create an OIDC identity provider](https://help.aliyun.com/document_detail/2331022.html) and [Create a role for SSO identity provider](https://help.aliyun.com/document_detail/2331016.html).
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
	// The AccessKey Secret for OSS read and write operations.
	//
	// example:
	//
	// xxxxxx
	OssAkSecret *string `json:"OssAkSecret,omitempty" xml:"OssAkSecret,omitempty"`
	// The OSS mount description. Parameter settings:
	//
	// - **bucketName**: the bucket name.
	//
	// - **bucketPath**: the folder or object that you created in OSS. If the OSS mount folder does not exist, an exception is triggered.
	//
	// - **mountPath**: the container path in Serverless App Engine. If the path already exists, it is an overwrite relationship. If the path does not exist, it is created.
	//
	// - **readOnly**: specifies whether the container path has read-only permission on the mounted folder resources. Valid values:
	//
	//     - **true**: read-only permission.
	//
	//     - **false**: read and write permission.
	//
	// example:
	//
	// [{"bucketName": "oss-bucket", "bucketPath": "data/user.data", "mountPath": "/usr/data/user.data", "readOnly": true}]
	OssMountDescs *string `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty"`
	// The application package type. Valid values:
	//
	// - When you deploy with Java, the following types are supported: **FatJar**, **War**, and **Image**.
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
	// - When you deploy with Python, the following types are supported: **PythonZip*	- and **Image**.
	//
	// example:
	//
	// FatJar
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The URL of the deployment package. This parameter is required when **Package Type*	- is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// http://myoss.oss-cn-hangzhou.aliyuncs.com/my-buc/2019-06-30/****.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version number of the deployment package. This parameter is required when **Package Type*	- is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// 1.0.1
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The PHP version on which the deployment package depends. Not supported for images.
	//
	// example:
	//
	// PHP-FPM 7.0
	Php *string `json:"Php,omitempty" xml:"Php,omitempty"`
	// The mount path for PHP application monitoring. Ensure that the PHP server loads the configuration file from this path. You do not need to manage the configuration content because Serverless App Engine automatically renders the correct configuration file.
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
	// The mount path for the PHP application startup configuration. Ensure that the PHP server uses this configuration file for startup.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// The script that is run after the container starts. A script is triggered immediately after the container is created. Format: `{"exec":{"command":["sh","-c","echo hello"\\]}}`.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script that is run before the container stops. A script is triggered before the container is deleted. Format: `{"exec":{"command":["sh","-c","echo hello"\\]}}`.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// Enables K8s Service-based service registration and discovery. Valid values:
	//
	// - **portProtocols**: the port and protocol. Valid port values: [1,65535]. Valid protocol values: **TCP*	- and **UDP**.
	//
	// - **portAndProtocol**: the port and protocol. Valid port values: [1,65535]. Valid protocol values: **TCP*	- and **UDP**. **portProtocols is recommended. If portProtocols is set, only portProtocols takes effect**.
	//
	// - **enable**: enables K8s Service-based service registration and discovery.
	//
	// example:
	//
	// {"portProtocols":[{"port":18012,"protocol":"TCP"}],"portAndProtocol":{"18012":"TCP"},"enable":true}
	PvtzDiscoverySvc *string `json:"PvtzDiscoverySvc,omitempty" xml:"PvtzDiscoverySvc,omitempty"`
	// The Python environment. **PYTHON 3.9.15*	- is supported.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// The custom installation module dependencies. By default, the dependencies defined in the requirements.txt file in the root folder are installed. If no dependencies are configured or custom packages are needed, you can specify the dependencies to install.
	//
	// example:
	//
	// Flask==2.0
	PythonModules    *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	RaspConfigShrink *string `json:"RaspConfig,omitempty" xml:"RaspConfig,omitempty"`
	// The application startup status check. Containers that fail multiple health checks are shut down and restarted. Containers that do not pass the health check do not receive SLB traffic. The **exec**, **httpGet**, and **tcpSocket*	- methods are supported. For specific examples, see the **Liveness*	- parameter.
	//
	// > Only one method can be selected for health checks.
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
	// The **Secret*	- mount description. Use a secret created on the namespace secrets page to inject sensitive information into the container. Settings:
	//
	// - **secretId**: the secret instance ID. You can obtain it by calling the ListSecrets operation.
	//
	// - **key**: the key.
	//
	// > You can mount all keys by passing the `sae-sys-secret-all` parameter.
	//
	// - **mountPath**: the mount path.
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
	// The canary release tags configured for the application.
	//
	// example:
	//
	// {\\"alicloud.service.tag\\":\\"g1\\"}
	ServiceTags *string `json:"ServiceTags,omitempty" xml:"ServiceTags,omitempty"`
	// The sidecar container configuration.
	SidecarContainersConfigShrink *string `json:"SidecarContainersConfig,omitempty" xml:"SidecarContainersConfig,omitempty"`
	// The configuration for log collection to Simple Log Service.
	//
	// - Use SLS resources automatically created by Serverless App Engine: `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - Use custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// Settings:
	//
	// - **projectName**: the name of the project in Simple Log Service.
	//
	// - **logDir**: the log path.
	//
	// - **logType**: the log type. **stdout*	- indicates container standard output logs, and only one entry can be set. If this parameter is not set, file logs are collected.
	//
	// - **logstoreName**: the name of the Logstore in Simple Log Service.
	//
	// - **logtailName**: the name of the Logtail in Simple Log Service. If this parameter is not specified, a new Logtail is created by automatic creation.
	//
	// If the SLS collection configuration has not changed during multiple deployments, you do not need to set this parameter (that is, the **SlsConfigs*	- field does not need to be included in the request). If you no longer need the SLS collection feature, set the value of this field to an empty string (that is, set the value of the **SlsConfigs*	- field to "" in the request).
	//
	// > A project that is automatically created with the application is deleted when the application is deleted. Therefore, do not select a project that is automatically created by Serverless App Engine when selecting an existing project.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// sls log tags
	SlsLogEnvTags *string `json:"SlsLogEnvTags,omitempty" xml:"SlsLogEnvTags,omitempty"`
	// Enables the application startup probe.
	//
	// - Check succeeded: indicates that the application started successfully. If you configured Liveness and Readiness checks, they are performed after the application starts successfully.
	//
	// - Check failed: indicates that the application failed to start. An exception is reported and the application is automatically restarted.
	//
	// > This is the description content.
	//
	// > - The exec, httpGet, and tcpSocket methods are supported. For specific examples, see the Liveness parameter.
	//
	// > - Only one method can be selected for health checks.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	StartupProbe *string `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	// Configures K8s Service-based service registration and discovery with end-to-end canary release:
	//
	// - enable: specifies whether to enable end-to-end canary release based on K8s Service.
	//
	//   - true: enabled.
	//
	//   - false: disabled.
	//
	// - namespaceId: the namespace ID.
	//
	// - portAndProtocol: the listening port and protocol. Format: {"port:protocol type":"container port"}.
	//
	// - portProtocols: defines the service port and protocol.
	//
	//   - port: the port.
	//
	//   - protocol: the protocol.
	//
	//   - targetPort: the container port.
	//
	// - pvtzDiscoveryName: the service discovery name.
	//
	// - serviceId: the service ID.
	//
	// - serviceName: the service name.
	//
	// example:
	//
	// {\\"enable\\":\\"false\\",\\"namespaceId\\":\\"cn-beijing:test\\",\\"portAndProtocol\\":{\\"2000:TCP\\":\\"18081\\"},\\"portProtocols\\":[{\\"port\\":2000,\\"protocol\\":\\"TCP\\",\\"targetPort\\":18081}],\\"pvtzDiscoveryName\\":\\"cn-beijing-1421801774382676\\",\\"serviceId\\":\\"3513\\",\\"serviceName\\":\\"demo-gray.test\\"}
	SwimlanePvtzDiscoverySvc *string `json:"SwimlanePvtzDiscoverySvc,omitempty" xml:"SwimlanePvtzDiscoverySvc,omitempty"`
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
	// The Tomcat configuration. Set this parameter to "" or "{}" to delete the configuration. Valid values:
	//
	// - **port**: the port number. Valid values: 1024 to 65535. Ports smaller than 1024 require root permissions. Because the container is configured with admin permissions, specify a port greater than 1024. Default value: 8080.
	//
	// - **contextPath**: the access path. Default value: root directory "/".
	//
	// - **maxThreads**: the connection pool size. Default value: 400.
	//
	// - **uriEncoding**: the encoding format of Tomcat. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. Default value: **ISO-8859-1**.
	//
	// - **useBodyEncodingForUri**: specifies whether to use **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	// The deployment policy. When the minimum number of available instances is 1, the value of the **UpdateStrategy*	- field is "". When the minimum number of available instances is greater than 1, examples are as follows:
	//
	//
	//
	//  - Grayscale 1 instance + 2 subsequent batches + automatic batching + 1-minute batch interval: `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}`
	//
	//   - Grayscale 1 instance + 2 subsequent batches + manual batching: `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"manual"},"grayUpdate":{"gray":1}}`
	//
	//   - 2 batches + automatic batching + 0-minute batch interval: `{"type":"BatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":0}}`
	//
	// Settings:
	//
	//   - **type**: the publish policy type. Valid values: **GrayBatchUpdate*	- (grayscale publish) and **BatchUpdate*	- (batch publish).
	//
	//   - **batchUpdate**: the batch publish policy.
	//
	//     - **batch**: the number of publish batches.
	//
	//     - **releaseType**: the processing method between batches. Valid values: **auto*	- (automatic) and **manual*	- (manual).
	//
	//     - **batchWaitTime**: the interval between batches. Unit: minutes.
	//
	//   - **grayUpdate**: the number of grayscale instances. This parameter is required when **type*	- is set to **GrayBatchUpdate**.
	//
	// example:
	//
	// {"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}
	UpdateStrategy *string `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
	// The vSwitch where the network interface controllers (NICs) of the application instance reside. The vSwitch must be in the specified VPC.
	//
	// example:
	//
	// vsw-bp12mw1f8k3jgygk9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The startup command for deploying a WAR package-based application. The configuration procedure is the same as that for the startup command of an image-based deployment. For more information, see [Configure a startup command](https://help.aliyun.com/document_detail/96677.html).
	//
	// example:
	//
	// CATALINA_OPTS=\\"$CATALINA_OPTS $Options\\" catalina.sh run
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

func (s *DeployApplicationShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *DeployApplicationShrinkRequest) GetLiveness() *string {
	return s.Liveness
}

func (s *DeployApplicationShrinkRequest) GetLokiConfigs() *string {
	return s.LokiConfigs
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

func (s *DeployApplicationShrinkRequest) GetRaspConfigShrink() *string {
	return s.RaspConfigShrink
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

func (s *DeployApplicationShrinkRequest) GetSlsLogEnvTags() *string {
	return s.SlsLogEnvTags
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

func (s *DeployApplicationShrinkRequest) SetLabelsShrink(v string) *DeployApplicationShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetLiveness(v string) *DeployApplicationShrinkRequest {
	s.Liveness = &v
	return s
}

func (s *DeployApplicationShrinkRequest) SetLokiConfigs(v string) *DeployApplicationShrinkRequest {
	s.LokiConfigs = &v
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

func (s *DeployApplicationShrinkRequest) SetRaspConfigShrink(v string) *DeployApplicationShrinkRequest {
	s.RaspConfigShrink = &v
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

func (s *DeployApplicationShrinkRequest) SetSlsLogEnvTags(v string) *DeployApplicationShrinkRequest {
	s.SlsLogEnvTags = &v
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

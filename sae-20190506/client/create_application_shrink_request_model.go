// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrAssumeRoleArn(v string) *CreateApplicationShrinkRequest
	GetAcrAssumeRoleArn() *string
	SetAcrInstanceId(v string) *CreateApplicationShrinkRequest
	GetAcrInstanceId() *string
	SetAgentVersion(v string) *CreateApplicationShrinkRequest
	GetAgentVersion() *string
	SetAppDescription(v string) *CreateApplicationShrinkRequest
	GetAppDescription() *string
	SetAppName(v string) *CreateApplicationShrinkRequest
	GetAppName() *string
	SetAppSource(v string) *CreateApplicationShrinkRequest
	GetAppSource() *string
	SetAssociateEip(v bool) *CreateApplicationShrinkRequest
	GetAssociateEip() *bool
	SetAutoConfig(v bool) *CreateApplicationShrinkRequest
	GetAutoConfig() *bool
	SetBaseAppId(v string) *CreateApplicationShrinkRequest
	GetBaseAppId() *string
	SetCommand(v string) *CreateApplicationShrinkRequest
	GetCommand() *string
	SetCommandArgs(v string) *CreateApplicationShrinkRequest
	GetCommandArgs() *string
	SetConfigMapMountDesc(v string) *CreateApplicationShrinkRequest
	GetConfigMapMountDesc() *string
	SetCpu(v int32) *CreateApplicationShrinkRequest
	GetCpu() *int32
	SetCustomHostAlias(v string) *CreateApplicationShrinkRequest
	GetCustomHostAlias() *string
	SetCustomImageNetworkType(v string) *CreateApplicationShrinkRequest
	GetCustomImageNetworkType() *string
	SetDeploy(v bool) *CreateApplicationShrinkRequest
	GetDeploy() *bool
	SetDiskSize(v int32) *CreateApplicationShrinkRequest
	GetDiskSize() *int32
	SetDotnet(v string) *CreateApplicationShrinkRequest
	GetDotnet() *string
	SetEdasContainerVersion(v string) *CreateApplicationShrinkRequest
	GetEdasContainerVersion() *string
	SetEmptyDirDesc(v string) *CreateApplicationShrinkRequest
	GetEmptyDirDesc() *string
	SetEnableCpuBurst(v bool) *CreateApplicationShrinkRequest
	GetEnableCpuBurst() *bool
	SetEnableEbpf(v string) *CreateApplicationShrinkRequest
	GetEnableEbpf() *string
	SetEnableNamespaceAgentVersion(v bool) *CreateApplicationShrinkRequest
	GetEnableNamespaceAgentVersion() *bool
	SetEnableNamespaceSlsConfig(v bool) *CreateApplicationShrinkRequest
	GetEnableNamespaceSlsConfig() *bool
	SetEnableNewArms(v bool) *CreateApplicationShrinkRequest
	GetEnableNewArms() *bool
	SetEnablePrometheus(v bool) *CreateApplicationShrinkRequest
	GetEnablePrometheus() *bool
	SetEnableSidecarResourceIsolated(v bool) *CreateApplicationShrinkRequest
	GetEnableSidecarResourceIsolated() *bool
	SetEnvs(v string) *CreateApplicationShrinkRequest
	GetEnvs() *string
	SetGpuConfig(v string) *CreateApplicationShrinkRequest
	GetGpuConfig() *string
	SetHeadlessPvtzDiscoverySvc(v string) *CreateApplicationShrinkRequest
	GetHeadlessPvtzDiscoverySvc() *string
	SetHtml(v string) *CreateApplicationShrinkRequest
	GetHtml() *string
	SetImagePullSecrets(v string) *CreateApplicationShrinkRequest
	GetImagePullSecrets() *string
	SetImageUrl(v string) *CreateApplicationShrinkRequest
	GetImageUrl() *string
	SetInitContainersConfigShrink(v string) *CreateApplicationShrinkRequest
	GetInitContainersConfigShrink() *string
	SetIsStateful(v bool) *CreateApplicationShrinkRequest
	GetIsStateful() *bool
	SetJarStartArgs(v string) *CreateApplicationShrinkRequest
	GetJarStartArgs() *string
	SetJarStartOptions(v string) *CreateApplicationShrinkRequest
	GetJarStartOptions() *string
	SetJdk(v string) *CreateApplicationShrinkRequest
	GetJdk() *string
	SetKafkaConfigs(v string) *CreateApplicationShrinkRequest
	GetKafkaConfigs() *string
	SetLabelsShrink(v string) *CreateApplicationShrinkRequest
	GetLabelsShrink() *string
	SetLiveness(v string) *CreateApplicationShrinkRequest
	GetLiveness() *string
	SetLokiConfigs(v string) *CreateApplicationShrinkRequest
	GetLokiConfigs() *string
	SetMemory(v int32) *CreateApplicationShrinkRequest
	GetMemory() *int32
	SetMicroRegistration(v string) *CreateApplicationShrinkRequest
	GetMicroRegistration() *string
	SetMicroRegistrationConfig(v string) *CreateApplicationShrinkRequest
	GetMicroRegistrationConfig() *string
	SetMicroserviceEngineConfig(v string) *CreateApplicationShrinkRequest
	GetMicroserviceEngineConfig() *string
	SetMountDesc(v string) *CreateApplicationShrinkRequest
	GetMountDesc() *string
	SetMountHost(v string) *CreateApplicationShrinkRequest
	GetMountHost() *string
	SetNamespaceId(v string) *CreateApplicationShrinkRequest
	GetNamespaceId() *string
	SetNasConfigs(v string) *CreateApplicationShrinkRequest
	GetNasConfigs() *string
	SetNasId(v string) *CreateApplicationShrinkRequest
	GetNasId() *string
	SetNewSaeVersion(v string) *CreateApplicationShrinkRequest
	GetNewSaeVersion() *string
	SetOidcRoleName(v string) *CreateApplicationShrinkRequest
	GetOidcRoleName() *string
	SetOssAkId(v string) *CreateApplicationShrinkRequest
	GetOssAkId() *string
	SetOssAkSecret(v string) *CreateApplicationShrinkRequest
	GetOssAkSecret() *string
	SetOssMountDescs(v string) *CreateApplicationShrinkRequest
	GetOssMountDescs() *string
	SetPackageType(v string) *CreateApplicationShrinkRequest
	GetPackageType() *string
	SetPackageUrl(v string) *CreateApplicationShrinkRequest
	GetPackageUrl() *string
	SetPackageVersion(v string) *CreateApplicationShrinkRequest
	GetPackageVersion() *string
	SetPhp(v string) *CreateApplicationShrinkRequest
	GetPhp() *string
	SetPhpArmsConfigLocation(v string) *CreateApplicationShrinkRequest
	GetPhpArmsConfigLocation() *string
	SetPhpConfig(v string) *CreateApplicationShrinkRequest
	GetPhpConfig() *string
	SetPhpConfigLocation(v string) *CreateApplicationShrinkRequest
	GetPhpConfigLocation() *string
	SetPostStart(v string) *CreateApplicationShrinkRequest
	GetPostStart() *string
	SetPreStop(v string) *CreateApplicationShrinkRequest
	GetPreStop() *string
	SetProgrammingLanguage(v string) *CreateApplicationShrinkRequest
	GetProgrammingLanguage() *string
	SetPvtzDiscoverySvc(v string) *CreateApplicationShrinkRequest
	GetPvtzDiscoverySvc() *string
	SetPython(v string) *CreateApplicationShrinkRequest
	GetPython() *string
	SetPythonModules(v string) *CreateApplicationShrinkRequest
	GetPythonModules() *string
	SetRaspConfigShrink(v string) *CreateApplicationShrinkRequest
	GetRaspConfigShrink() *string
	SetReadiness(v string) *CreateApplicationShrinkRequest
	GetReadiness() *string
	SetReplicas(v int32) *CreateApplicationShrinkRequest
	GetReplicas() *int32
	SetResourceType(v string) *CreateApplicationShrinkRequest
	GetResourceType() *string
	SetSaeVersion(v string) *CreateApplicationShrinkRequest
	GetSaeVersion() *string
	SetSecretMountDesc(v string) *CreateApplicationShrinkRequest
	GetSecretMountDesc() *string
	SetSecurityGroupId(v string) *CreateApplicationShrinkRequest
	GetSecurityGroupId() *string
	SetServiceTags(v string) *CreateApplicationShrinkRequest
	GetServiceTags() *string
	SetSidecarContainersConfigShrink(v string) *CreateApplicationShrinkRequest
	GetSidecarContainersConfigShrink() *string
	SetSlsConfigs(v string) *CreateApplicationShrinkRequest
	GetSlsConfigs() *string
	SetSlsLogEnvTags(v string) *CreateApplicationShrinkRequest
	GetSlsLogEnvTags() *string
	SetStartupProbe(v string) *CreateApplicationShrinkRequest
	GetStartupProbe() *string
	SetTerminationGracePeriodSeconds(v int32) *CreateApplicationShrinkRequest
	GetTerminationGracePeriodSeconds() *int32
	SetTimezone(v string) *CreateApplicationShrinkRequest
	GetTimezone() *string
	SetTomcatConfig(v string) *CreateApplicationShrinkRequest
	GetTomcatConfig() *string
	SetVSwitchId(v string) *CreateApplicationShrinkRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateApplicationShrinkRequest
	GetVpcId() *string
	SetWarStartOptions(v string) *CreateApplicationShrinkRequest
	GetWarStartOptions() *string
	SetWebContainer(v string) *CreateApplicationShrinkRequest
	GetWebContainer() *string
}

type CreateApplicationShrinkRequest struct {
	// The ARN of the RAM role required for cross-account image pulling. For more information, see [Grant permissions across Alibaba Cloud accounts by using a RAM role](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The instance ID of the Container Registry Enterprise instance. This parameter is required when **ImageUrl*	- is set to a Container Registry Enterprise Edition image.
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
	// The application description. The description can be up to 1024 characters in length.
	//
	// example:
	//
	// This is a test description.
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The application name. The name can contain digits, letters, and hyphens (-). The name must start with a letter and cannot end with a hyphen (-). The name can be up to 36 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Set this parameter to micro_service to create a microservice application.
	//
	// example:
	//
	// micro_service
	AppSource *string `json:"AppSource,omitempty" xml:"AppSource,omitempty"`
	// Specifies whether to associate an EIP. Valid values:
	//
	// - **true**: associate an EIP.
	//
	// - **false**: do not associate an EIP.
	//
	// example:
	//
	// true
	AssociateEip *bool `json:"AssociateEip,omitempty" xml:"AssociateEip,omitempty"`
	// Specifies whether to automatically configure the network environment. Valid values:
	//
	// - **true**: SAE automatically configures the network environment when the application is created. The values of **NamespaceId**, **VpcId**, **vSwitchId**, and **SecurityGroupId*	- are ignored.
	//
	// - **false**: SAE manually configures the network environment when the application is created.
	//
	// > If this parameter is set to **true**, any other **NamespaceId*	- value that is passed is ignored.
	//
	// example:
	//
	// true
	AutoConfig *bool `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty"`
	// The base application ID.
	//
	// example:
	//
	// ee99cce6-1c8e-4bfa-96c3-3e2fa9de8a41
	BaseAppId *string `json:"BaseAppId,omitempty" xml:"BaseAppId,omitempty"`
	// The command that is used to start the image. The command must be an executable object in the container. Example:
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
	// In the preceding example, `Command="echo", CommandArgs=["abc", ">", "file0"]`.
	//
	//
	// 	Notice: This parameter is required when PackageType is set to DotnetZip.
	//
	// example:
	//
	// echo
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments of the image startup command. These are the arguments required by the startup command specified in **Command**. Format:
	//
	// `["a","b"]`
	//
	// In the preceding example, `CommandArgs=["abc", ">", "file0"]`, where `["abc", ">", "file0"]` must be converted to the String type. The internal format is a JSON array. If this parameter is not required, leave it empty.
	//
	// 	Notice: This parameter is required when PackageType is set to DotnetZip.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The **ConfigMap*	- mount description. Use a ConfigMap created on the namespace configuration items page to inject configuration information into the container. Parameter description:
	//
	// - **configMapId**: the ConfigMap instance ID. You can obtain the ID by invoking the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation.
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
	// The CPU specifications required for each instance, in millicores. This parameter cannot be set to 0. Only the following defined specifications are supported:
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
	// - **hostName**: the domain name or hostname.
	//
	// - **ip**: the IP address.
	//
	// example:
	//
	// [{"hostName":"samplehost","ip":"127.0.0.1"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The custom image type. Set this parameter to an empty string if the image is not a custom image:
	//
	// - internet: public image
	//
	// - intranet: private image
	//
	// example:
	//
	// internet
	CustomImageNetworkType *string `json:"CustomImageNetworkType,omitempty" xml:"CustomImageNetworkType,omitempty"`
	// Specifies whether to immediately deploy the application. Valid values:
	//
	// - **true**: default value. The application is deployed immediately.
	//
	// - **false**: the application is deployed later.
	//
	// example:
	//
	// true
	Deploy *bool `json:"Deploy,omitempty" xml:"Deploy,omitempty"`
	// The disk storage size, in GB.
	//
	// example:
	//
	// 50
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
	// Specifies whether to enable the CPU Burst feature:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	EnableCpuBurst *bool `json:"EnableCpuBurst,omitempty" xml:"EnableCpuBurst,omitempty"`
	// Specifies whether to enable application monitoring for non-Java applications based on eBPF technology. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled. This is the default value.
	//
	// example:
	//
	// false
	EnableEbpf *string `json:"EnableEbpf,omitempty" xml:"EnableEbpf,omitempty"`
	// Specifies whether to reuse the namespace agent version configuration.
	//
	// example:
	//
	// true
	EnableNamespaceAgentVersion *bool `json:"EnableNamespaceAgentVersion,omitempty" xml:"EnableNamespaceAgentVersion,omitempty"`
	// Specifies whether to reuse the namespace SLS log configuration.
	//
	// example:
	//
	// true
	EnableNamespaceSlsConfig *bool `json:"EnableNamespaceSlsConfig,omitempty" xml:"EnableNamespaceSlsConfig,omitempty"`
	// Specifies whether to enable the new ARMS feature:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// false
	EnableNewArms *bool `json:"EnableNewArms,omitempty" xml:"EnableNewArms,omitempty"`
	// Specifies whether to enable Prometheus custom metric collection.
	//
	// example:
	//
	// false
	EnablePrometheus *bool `json:"EnablePrometheus,omitempty" xml:"EnablePrometheus,omitempty"`
	// Specifies whether to enable sidecar resource isolation:
	//
	// - true: Isolated.
	//
	// - false: Not isolated.
	//
	// example:
	//
	// true
	EnableSidecarResourceIsolated *bool `json:"EnableSidecarResourceIsolated,omitempty" xml:"EnableSidecarResourceIsolated,omitempty"`
	// The container environment variable parameters. You can customize environment variables or reference a ConfigMap. To reference a ConfigMap, create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//     - **name**: the name of the environment variable.
	//
	//     - **value**: the value of the environment variable. This takes priority over valueFrom.
	//
	// - Reference a ConfigMap (valueFrom)
	//
	//     - **name**: the name of the environment variable. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<ConfigMap name>`, such as `sae-sys-configmap-all-test1`.
	//
	//     - **valueFrom**: the environment variable reference. Set the value to `configMapRef`.
	//
	//         - **configMapId**: the ConfigMap ID.
	//
	//         - **key**: the key. If you reference all keys, do not set this field.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "name": "sae-sys-configmap-all-hello",
	//
	//         "valueFrom": {
	//
	//             "configMapRef": {
	//
	//                 "configMapId": 100,
	//
	//                 "key": ""
	//
	//             }
	//
	//         }
	//
	//     },
	//
	//     {
	//
	//         "name": "hello",
	//
	//         "valueFrom": {
	//
	//             "configMapRef": {
	//
	//                 "configMapId": 101,
	//
	//                 "key": "php-fpm"
	//
	//             }
	//
	//         }
	//
	//     },
	//
	//     {
	//
	//         "name": "envtmp",
	//
	//         "value": "newenv"
	//
	//     }
	//
	// ]
	Envs      *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	GpuConfig *string `json:"GpuConfig,omitempty" xml:"GpuConfig,omitempty"`
	// The K8s Headless Service-based service registration and discovery.
	//
	// - serviceName: the service name.
	//
	// - namespaceId: the namespace ID.
	//
	// example:
	//
	// {\\"serviceName\\":\\"leaf-test-headless\\",\\"namespaceId\\":\\"cn-zhangjiakou:prod\\"}
	HeadlessPvtzDiscoverySvc *string `json:"HeadlessPvtzDiscoverySvc,omitempty" xml:"HeadlessPvtzDiscoverySvc,omitempty"`
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
	// The image address. This parameter is required when **Package Type*	- is set to **Image**.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/sae_test/ali_sae_test:0.0.1
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The init container configuration.
	InitContainersConfigShrink *string `json:"InitContainersConfig,omitempty" xml:"InitContainersConfig,omitempty"`
	// Specifies whether the application is stateful.
	IsStateful *bool `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The arguments for starting the JAR package application. The default startup command for the application: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// custom-args
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The options for starting the JAR package application. The default startup command for the application: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// -Xms4G -Xmx4G
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
	// The summary of configurations for log collection to Kafka. Valid values:
	//
	// - **kafkaEndpoint**: the service registration address of the Kafka API.
	//
	// - **kafkaInstanceId**: the Kafka instance ID.
	//
	// - **kafkaConfigs**: the summary of configurations for one or more log entries. For more information about the valid values, see the **kafkaConfigs*	- request parameter in this topic.
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
	// > You can use only one method for health checks.
	//
	// Parameter description:
	//
	// - **exec.command**: the health check command.
	//
	// - **httpGet.path**: the access path.
	//
	// - **httpGet.scheme**: **HTTP*	- or **HTTPS**.
	//
	// - **httpGet.isContainKeyWord**: **true*	- indicates that the keyword is included. **false*	- indicates that the keyword is not included. If this field is missing, the advanced feature is not used.
	//
	// - **httpGet.keyWord**: the custom keyword. The **isContainKeyWord*	- field must be present when this field is used.
	//
	// - **tcpSocket.port**: the port for TCP connection detection.
	//
	// - **initialDelaySeconds**: the health check delay detection time. Default value: 10. Unit: seconds.
	//
	// - **periodSeconds**: the health check period. Default value: 30. Unit: seconds.
	//
	// - **timeoutSeconds**: the health check timeout period. Default value: 1. Unit: seconds. If this parameter is set to 0 or is not set, the default timeout period is 1 second.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	Liveness    *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	LokiConfigs *string `json:"LokiConfigs,omitempty" xml:"LokiConfigs,omitempty"`
	// The memory required for each instance, in MB. This parameter cannot be set to 0. The memory has a one-to-one mapping with CPU. Only the following defined specifications are supported:
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
	// - **0**: SAE built-in Nacos.
	//
	// - **1**: self-managed Nacos.
	//
	// - **2**: MSE commercial edition Nacos.
	//
	// example:
	//
	// "0"
	MicroRegistration *string `json:"MicroRegistration,omitempty" xml:"MicroRegistration,omitempty"`
	// The registry configuration.
	//
	// example:
	//
	// {\\"instanceId\\":\\"mse-cn-zvp2bh6h70r\\",\\"namespace\\":\\"4c0aa74f-57cb-423c-b6af-5d9f2d0e3dbd\\"}
	MicroRegistrationConfig *string `json:"MicroRegistrationConfig,omitempty" xml:"MicroRegistrationConfig,omitempty"`
	// Configures the microservice governance feature.
	//
	// - Specifies whether to enable microservice governance (enable):
	//
	//    - true: Enabled.
	//
	//   - false: Disabled.
	//
	// - Configures lossless online/offline (mseLosslessRule):
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
	// {"enable": true,"mseLosslessRule": {"delayTime": 0,"enable": false,"notice": false,"warmupTime": 120}}
	MicroserviceEngineConfig *string `json:"MicroserviceEngineConfig,omitempty" xml:"MicroserviceEngineConfig,omitempty"`
	// We recommend that you do not set this parameter. Set **NasConfigs*	- instead. The NAS mount description. If the configuration does not change during deployment, you do not need to set this parameter (that is, the **MountDesc*	- field does not need to be included in the request). To clear the NAS configuration, set the value of this field to an empty string (that is, set the value of the **MountDesc*	- field to "" in the request).
	//
	// example:
	//
	// [{mountPath: "/tmp", nasPath: "/"}]
	MountDesc *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	// We recommend that you do not set this parameter. Set **NasConfigs*	- instead. The mount target of the NAS file system in the VPC of the application. If the configuration does not change during deployment, you do not need to set this parameter (that is, the **MountHost*	- field does not need to be included in the request). To clear the NAS configuration, set the value of this field to an empty string (that is, set the value of the **MountHost*	- field to "" in the request).
	//
	// example:
	//
	// example.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The SAE namespace ID. Only namespaces whose names contain lowercase letters and hyphens (-) are supported. The name must start with a letter.
	//
	// You can obtain namespaces by calling the [DescribeNamespaceList](https://help.aliyun.com/document_detail/126547.html) operation.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The NAS mount configuration. Valid values:
	//
	// - **mountPath**: the container mount path.
	//
	// - **readOnly**: set to **false*	- to grant read and write permission.
	//
	// - **nasId**: the NAS ID.
	//
	// - **mountDomain**: the container mount target address. For more information, see [DescribeMountTargets](https://help.aliyun.com/document_detail/62626.html).
	//
	// - **nasPath**: the NAS relative file directory.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// We recommend that you do not set this parameter. Set **NasConfigs*	- instead. The ID of the mounted NAS file system. The NAS file system must be in the same region as the cluster. The NAS file system must have available mount target creation quota, or its mount target must already be on a vSwitch in the VPC. If this parameter is left empty and the **mountDescs*	- field exists, a NAS file system is automatically purchased and mounted to a vSwitch in the VPC.
	//
	// If the configuration does not change during deployment, you do not need to set this parameter (that is, the **NASId*	- field does not need to be included in the request). To clear the NAS configuration, set the value of this field to an empty string (that is, set the value of the **NASId*	- field to "" in the request).
	//
	// example:
	//
	// KSAK****
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
	// Specifies the RAM role for identity authentication.
	//
	// > Create an OIDC identity provider and an identity provider role in the same region in advance. For more information, see <props="china">[CreateOIDCProvider](https://www.alibabacloud.com/help/en/ram/developer-reference/api-ims-2019-08-15-createoidcprovider) and [CreateSAMLProvider](https://www.alibabacloud.com/help/en/ram/developer-reference/api-ims-2019-08-15-createsamlprovider)<props="intl">[CreateOIDCProvider](https://www.alibabacloud.com/help/zh/ram/developer-reference/api-ims-2019-08-15-createoidcprovider) and [CreateSAMLProvider](https://www.alibabacloud.com/help/zh/ram/developer-reference/api-ims-2019-08-15-createsamlprovider).
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
	// The OSS mount description. Parameter description:
	//
	// - **bucketName**: the bucket name.
	//
	// - **bucketPath**: the folder or object that you created in OSS. If the OSS mount folder does not exist, an exception is triggered.
	//
	// - **mountPath**: the container path in SAE. If the path already exists, it is an overwrite relationship. If the path does not exist, it is created.
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
	// The type of the application deployment package. Valid values:
	//
	// - If you use Java for deployment, **FatJar**, **War**, and **Image*	- are supported.
	//
	// - If you use PHP for deployment, the following types are supported:
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
	// - If you use Python for deployment, **PythonZip*	- and **Image*	- are supported.
	//
	// - If you use .NET Core for deployment, **DotnetZip*	- and **Image*	- are supported.
	//
	//   >
	//
	//   > When DotnetZip is selected, Dotnet specifies the version of the .NET Core runtime. .NET 3.1, .NET 5.0, .NET 6.0, .NET 7.0, and .NET 8.0 are supported. The Dotnet, Command, and CommandArgs parameters are required.
	//
	// This parameter is required.
	//
	// example:
	//
	// FatJar
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The address of the deployment package. This parameter is required when **Package Type*	- is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// http://myoss.oss-cn-****.aliyuncs.com/my-buc/2019-06-30/****.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package. This parameter is required when **Package Type*	- is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// 1.0.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The PHP version on which the deployment package depends. Not supported for images.
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
	// The mount path for the PHP application startup configuration. Make sure that the PHP server uses this configuration file to start.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// The script that is run after the container is started. A script is triggered and run immediately after the container is created. Format: `{"exec":{"command":["cat","/etc/group"]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script that is run before the container is stopped. A script is triggered and run before the container is deleted. Format: `{"exec":{"command":["cat","/etc/group"]}}`
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
	// - **python**: Python.
	//
	// - **dotnet**: .NET Core.
	//
	// - **other**: multiple languages, such as C++, Go, and Node.js.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// Enables K8s Service-based service registration and discovery. Valid values:
	//
	// - **serviceName**: the service name. Format: `custom name-namespace ID`. The suffix `-namespace ID` cannot be customized and must be set based on the namespace of the application. For example, if you select the default namespace in the China (Beijing) region, the suffix is `-cn-beijing-default`.
	//
	// - **namespaceId**: the namespace ID.
	//
	// - **portProtocols**: the port and protocol. Valid port values: [1,65535]. Valid protocol values: **TCP*	- and **UDP**.
	//
	// - **portAndProtocol**: the port and protocol. Valid port values: [1,65535]. Valid protocol values: **TCP*	- and **UDP**. **portProtocols is recommended. If portProtocols is set, only portProtocols takes effect**.
	//
	// - **enable**: enables K8s Service-based service registration and discovery.
	//
	// example:
	//
	// {"serviceName":"bwm-poc-sc-gateway-cn-beijing-front","namespaceId":"cn-beijing:front","portAndProtocol":{"18012":"TCP"},"enable":true,"portProtocols":[{"port":18012,"protocol":"TCP"}]}
	PvtzDiscoverySvc *string `json:"PvtzDiscoverySvc,omitempty" xml:"PvtzDiscoverySvc,omitempty"`
	// The Python environment. **PYTHON 3.9.15*	- is supported.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// The custom installation module dependencies. By default, the dependencies defined in the requirements.txt file in the root folder are installed. If the file is not configured or you need custom packages, specify the dependencies to install.
	//
	// example:
	//
	// Flask==2.0
	PythonModules    *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	RaspConfigShrink *string `json:"RaspConfig,omitempty" xml:"RaspConfig,omitempty"`
	// The application startup status check. Containers that fail multiple health checks are shut down and restarted. Containers that do not pass the health check do not receive SLB traffic. The **exec**, **httpGet**, and **tcpSocket*	- methods are supported. For specific examples, see the **Liveness*	- parameter.
	//
	// > You can use only one method for health checks.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The initial number of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The resource type. Valid values: NULL (default), default, and haiguang (Hygon server).
	//
	// example:
	//
	// NULL
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The SAE version. Valid values:
	//
	// - **v1**
	//
	// - **v2**
	//
	// example:
	//
	// v1
	SaeVersion *string `json:"SaeVersion,omitempty" xml:"SaeVersion,omitempty"`
	// The **Secret*	- mount description. Use a secret created on the namespace secrets page to inject sensitive information into the container. Parameter description:
	//
	// - **secretId**: the secret instance ID. You can obtain the ID by calling the ListSecrets operation.
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
	// The configurations for log collection to Simple Log Service.
	//
	// - Use SLS resources that are automatically created by SAE: `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - Use custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// Parameter description:
	//
	// - **projectName**: the Project name in Simple Log Service.
	//
	// - **logDir**: the log path.
	//
	// - **logType**: the log type. **stdout*	- indicates container standard output logs. You can configure only one entry for this type. If this parameter is not set, file logs are collected.
	//
	// - **logstoreName**: the Logstore name in Simple Log Service.
	//
	// - **logtailName**: the Logtail name in Simple Log Service. If this parameter is not specified, a new Logtail is created.
	//
	// If the SLS collection configuration does not change during multiple deployments, you do not need to set this parameter (that is, the **SlsConfigs*	- field does not need to be included in the request). If you no longer need the SLS collection feature, set the value of this field to an empty string (that is, set the value of the **SlsConfigs*	- field to "" in the request).
	//
	// > Projects that are automatically created with the application are deleted when the application is deleted. Therefore, do not select a project that is automatically created by SAE when you select an existing project.
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
	// >
	//
	// > - The exec, httpGet, and tcpSocket methods are supported. For specific examples, see the Liveness parameter.
	//
	// > - You can use only one method for health checks.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	StartupProbe *string `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	// The timeout period for graceful shutdown. Default value: 30. Unit: seconds. Valid values: 1 to 300.
	//
	// example:
	//
	// 30
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The time zone. Default value: **Asia/Shanghai**.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The Tomcat configuration. Set this parameter to "" or "{}" to delete the configuration:
	//
	// - **port**: the port number. Valid values: 1024 to 65535. Ports less than 1024 require root permissions. Because the container is configured with admin permissions, specify a port greater than 1024. Default value: 8080.
	//
	// - **contextPath**: the access path. Default value: root directory "/".
	//
	// - **maxThreads**: the maximum number of connections in the connection pool. Default value: 400.
	//
	// - **uriEncoding**: the encoding format of Tomcat. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. Default value: **ISO-8859-1**.
	//
	// - **useBodyEncodingForUri**: specifies whether to use **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	// The vSwitch where the elastic network interface controller (NIC) of the application instance resides. The vSwitch must be in the specified VPC. The vSwitch also has a binding relationship with the SAE namespace. If you leave this parameter empty, the vSwitch attached to the namespace is used by default.
	//
	// example:
	//
	// vsw-bp12mw1f8k3jgygk9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC that corresponds to the SAE namespace. In SAE, a namespace can correspond to only one VPC, and the mapping cannot be modified. The binding relationship is established when the first SAE application is created in the namespace. Multiple namespaces can correspond to the same VPC. If you leave this parameter empty, the VPC bound to the namespace is used by default.
	//
	// example:
	//
	// vpc-bp1aevy8sofi8mh1q****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The startup command for deploying a WAR package application. The configuration procedure is the same as that for the startup command of an image deployment. For more information, see [Configure a startup command](https://help.aliyun.com/document_detail/96677.html).
	//
	// example:
	//
	// CATALINA_OPTS=\\"$CATALINA_OPTS $Options\\" catalina.sh run
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The version of Tomcat on which the WebContainer deployment package depends. Valid values:
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

func (s CreateApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationShrinkRequest) GetAcrAssumeRoleArn() *string {
	return s.AcrAssumeRoleArn
}

func (s *CreateApplicationShrinkRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CreateApplicationShrinkRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *CreateApplicationShrinkRequest) GetAppDescription() *string {
	return s.AppDescription
}

func (s *CreateApplicationShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateApplicationShrinkRequest) GetAppSource() *string {
	return s.AppSource
}

func (s *CreateApplicationShrinkRequest) GetAssociateEip() *bool {
	return s.AssociateEip
}

func (s *CreateApplicationShrinkRequest) GetAutoConfig() *bool {
	return s.AutoConfig
}

func (s *CreateApplicationShrinkRequest) GetBaseAppId() *string {
	return s.BaseAppId
}

func (s *CreateApplicationShrinkRequest) GetCommand() *string {
	return s.Command
}

func (s *CreateApplicationShrinkRequest) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *CreateApplicationShrinkRequest) GetConfigMapMountDesc() *string {
	return s.ConfigMapMountDesc
}

func (s *CreateApplicationShrinkRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateApplicationShrinkRequest) GetCustomHostAlias() *string {
	return s.CustomHostAlias
}

func (s *CreateApplicationShrinkRequest) GetCustomImageNetworkType() *string {
	return s.CustomImageNetworkType
}

func (s *CreateApplicationShrinkRequest) GetDeploy() *bool {
	return s.Deploy
}

func (s *CreateApplicationShrinkRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreateApplicationShrinkRequest) GetDotnet() *string {
	return s.Dotnet
}

func (s *CreateApplicationShrinkRequest) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *CreateApplicationShrinkRequest) GetEmptyDirDesc() *string {
	return s.EmptyDirDesc
}

func (s *CreateApplicationShrinkRequest) GetEnableCpuBurst() *bool {
	return s.EnableCpuBurst
}

func (s *CreateApplicationShrinkRequest) GetEnableEbpf() *string {
	return s.EnableEbpf
}

func (s *CreateApplicationShrinkRequest) GetEnableNamespaceAgentVersion() *bool {
	return s.EnableNamespaceAgentVersion
}

func (s *CreateApplicationShrinkRequest) GetEnableNamespaceSlsConfig() *bool {
	return s.EnableNamespaceSlsConfig
}

func (s *CreateApplicationShrinkRequest) GetEnableNewArms() *bool {
	return s.EnableNewArms
}

func (s *CreateApplicationShrinkRequest) GetEnablePrometheus() *bool {
	return s.EnablePrometheus
}

func (s *CreateApplicationShrinkRequest) GetEnableSidecarResourceIsolated() *bool {
	return s.EnableSidecarResourceIsolated
}

func (s *CreateApplicationShrinkRequest) GetEnvs() *string {
	return s.Envs
}

func (s *CreateApplicationShrinkRequest) GetGpuConfig() *string {
	return s.GpuConfig
}

func (s *CreateApplicationShrinkRequest) GetHeadlessPvtzDiscoverySvc() *string {
	return s.HeadlessPvtzDiscoverySvc
}

func (s *CreateApplicationShrinkRequest) GetHtml() *string {
	return s.Html
}

func (s *CreateApplicationShrinkRequest) GetImagePullSecrets() *string {
	return s.ImagePullSecrets
}

func (s *CreateApplicationShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateApplicationShrinkRequest) GetInitContainersConfigShrink() *string {
	return s.InitContainersConfigShrink
}

func (s *CreateApplicationShrinkRequest) GetIsStateful() *bool {
	return s.IsStateful
}

func (s *CreateApplicationShrinkRequest) GetJarStartArgs() *string {
	return s.JarStartArgs
}

func (s *CreateApplicationShrinkRequest) GetJarStartOptions() *string {
	return s.JarStartOptions
}

func (s *CreateApplicationShrinkRequest) GetJdk() *string {
	return s.Jdk
}

func (s *CreateApplicationShrinkRequest) GetKafkaConfigs() *string {
	return s.KafkaConfigs
}

func (s *CreateApplicationShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *CreateApplicationShrinkRequest) GetLiveness() *string {
	return s.Liveness
}

func (s *CreateApplicationShrinkRequest) GetLokiConfigs() *string {
	return s.LokiConfigs
}

func (s *CreateApplicationShrinkRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateApplicationShrinkRequest) GetMicroRegistration() *string {
	return s.MicroRegistration
}

func (s *CreateApplicationShrinkRequest) GetMicroRegistrationConfig() *string {
	return s.MicroRegistrationConfig
}

func (s *CreateApplicationShrinkRequest) GetMicroserviceEngineConfig() *string {
	return s.MicroserviceEngineConfig
}

func (s *CreateApplicationShrinkRequest) GetMountDesc() *string {
	return s.MountDesc
}

func (s *CreateApplicationShrinkRequest) GetMountHost() *string {
	return s.MountHost
}

func (s *CreateApplicationShrinkRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateApplicationShrinkRequest) GetNasConfigs() *string {
	return s.NasConfigs
}

func (s *CreateApplicationShrinkRequest) GetNasId() *string {
	return s.NasId
}

func (s *CreateApplicationShrinkRequest) GetNewSaeVersion() *string {
	return s.NewSaeVersion
}

func (s *CreateApplicationShrinkRequest) GetOidcRoleName() *string {
	return s.OidcRoleName
}

func (s *CreateApplicationShrinkRequest) GetOssAkId() *string {
	return s.OssAkId
}

func (s *CreateApplicationShrinkRequest) GetOssAkSecret() *string {
	return s.OssAkSecret
}

func (s *CreateApplicationShrinkRequest) GetOssMountDescs() *string {
	return s.OssMountDescs
}

func (s *CreateApplicationShrinkRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *CreateApplicationShrinkRequest) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *CreateApplicationShrinkRequest) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *CreateApplicationShrinkRequest) GetPhp() *string {
	return s.Php
}

func (s *CreateApplicationShrinkRequest) GetPhpArmsConfigLocation() *string {
	return s.PhpArmsConfigLocation
}

func (s *CreateApplicationShrinkRequest) GetPhpConfig() *string {
	return s.PhpConfig
}

func (s *CreateApplicationShrinkRequest) GetPhpConfigLocation() *string {
	return s.PhpConfigLocation
}

func (s *CreateApplicationShrinkRequest) GetPostStart() *string {
	return s.PostStart
}

func (s *CreateApplicationShrinkRequest) GetPreStop() *string {
	return s.PreStop
}

func (s *CreateApplicationShrinkRequest) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *CreateApplicationShrinkRequest) GetPvtzDiscoverySvc() *string {
	return s.PvtzDiscoverySvc
}

func (s *CreateApplicationShrinkRequest) GetPython() *string {
	return s.Python
}

func (s *CreateApplicationShrinkRequest) GetPythonModules() *string {
	return s.PythonModules
}

func (s *CreateApplicationShrinkRequest) GetRaspConfigShrink() *string {
	return s.RaspConfigShrink
}

func (s *CreateApplicationShrinkRequest) GetReadiness() *string {
	return s.Readiness
}

func (s *CreateApplicationShrinkRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *CreateApplicationShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateApplicationShrinkRequest) GetSaeVersion() *string {
	return s.SaeVersion
}

func (s *CreateApplicationShrinkRequest) GetSecretMountDesc() *string {
	return s.SecretMountDesc
}

func (s *CreateApplicationShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateApplicationShrinkRequest) GetServiceTags() *string {
	return s.ServiceTags
}

func (s *CreateApplicationShrinkRequest) GetSidecarContainersConfigShrink() *string {
	return s.SidecarContainersConfigShrink
}

func (s *CreateApplicationShrinkRequest) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *CreateApplicationShrinkRequest) GetSlsLogEnvTags() *string {
	return s.SlsLogEnvTags
}

func (s *CreateApplicationShrinkRequest) GetStartupProbe() *string {
	return s.StartupProbe
}

func (s *CreateApplicationShrinkRequest) GetTerminationGracePeriodSeconds() *int32 {
	return s.TerminationGracePeriodSeconds
}

func (s *CreateApplicationShrinkRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateApplicationShrinkRequest) GetTomcatConfig() *string {
	return s.TomcatConfig
}

func (s *CreateApplicationShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateApplicationShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateApplicationShrinkRequest) GetWarStartOptions() *string {
	return s.WarStartOptions
}

func (s *CreateApplicationShrinkRequest) GetWebContainer() *string {
	return s.WebContainer
}

func (s *CreateApplicationShrinkRequest) SetAcrAssumeRoleArn(v string) *CreateApplicationShrinkRequest {
	s.AcrAssumeRoleArn = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAcrInstanceId(v string) *CreateApplicationShrinkRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAgentVersion(v string) *CreateApplicationShrinkRequest {
	s.AgentVersion = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAppDescription(v string) *CreateApplicationShrinkRequest {
	s.AppDescription = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAppName(v string) *CreateApplicationShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAppSource(v string) *CreateApplicationShrinkRequest {
	s.AppSource = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAssociateEip(v bool) *CreateApplicationShrinkRequest {
	s.AssociateEip = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAutoConfig(v bool) *CreateApplicationShrinkRequest {
	s.AutoConfig = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetBaseAppId(v string) *CreateApplicationShrinkRequest {
	s.BaseAppId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetCommand(v string) *CreateApplicationShrinkRequest {
	s.Command = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetCommandArgs(v string) *CreateApplicationShrinkRequest {
	s.CommandArgs = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetConfigMapMountDesc(v string) *CreateApplicationShrinkRequest {
	s.ConfigMapMountDesc = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetCpu(v int32) *CreateApplicationShrinkRequest {
	s.Cpu = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetCustomHostAlias(v string) *CreateApplicationShrinkRequest {
	s.CustomHostAlias = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetCustomImageNetworkType(v string) *CreateApplicationShrinkRequest {
	s.CustomImageNetworkType = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDeploy(v bool) *CreateApplicationShrinkRequest {
	s.Deploy = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDiskSize(v int32) *CreateApplicationShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDotnet(v string) *CreateApplicationShrinkRequest {
	s.Dotnet = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEdasContainerVersion(v string) *CreateApplicationShrinkRequest {
	s.EdasContainerVersion = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEmptyDirDesc(v string) *CreateApplicationShrinkRequest {
	s.EmptyDirDesc = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEnableCpuBurst(v bool) *CreateApplicationShrinkRequest {
	s.EnableCpuBurst = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEnableEbpf(v string) *CreateApplicationShrinkRequest {
	s.EnableEbpf = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEnableNamespaceAgentVersion(v bool) *CreateApplicationShrinkRequest {
	s.EnableNamespaceAgentVersion = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEnableNamespaceSlsConfig(v bool) *CreateApplicationShrinkRequest {
	s.EnableNamespaceSlsConfig = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEnableNewArms(v bool) *CreateApplicationShrinkRequest {
	s.EnableNewArms = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEnablePrometheus(v bool) *CreateApplicationShrinkRequest {
	s.EnablePrometheus = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEnableSidecarResourceIsolated(v bool) *CreateApplicationShrinkRequest {
	s.EnableSidecarResourceIsolated = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEnvs(v string) *CreateApplicationShrinkRequest {
	s.Envs = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetGpuConfig(v string) *CreateApplicationShrinkRequest {
	s.GpuConfig = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetHeadlessPvtzDiscoverySvc(v string) *CreateApplicationShrinkRequest {
	s.HeadlessPvtzDiscoverySvc = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetHtml(v string) *CreateApplicationShrinkRequest {
	s.Html = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetImagePullSecrets(v string) *CreateApplicationShrinkRequest {
	s.ImagePullSecrets = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetImageUrl(v string) *CreateApplicationShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetInitContainersConfigShrink(v string) *CreateApplicationShrinkRequest {
	s.InitContainersConfigShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetIsStateful(v bool) *CreateApplicationShrinkRequest {
	s.IsStateful = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetJarStartArgs(v string) *CreateApplicationShrinkRequest {
	s.JarStartArgs = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetJarStartOptions(v string) *CreateApplicationShrinkRequest {
	s.JarStartOptions = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetJdk(v string) *CreateApplicationShrinkRequest {
	s.Jdk = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetKafkaConfigs(v string) *CreateApplicationShrinkRequest {
	s.KafkaConfigs = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetLabelsShrink(v string) *CreateApplicationShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetLiveness(v string) *CreateApplicationShrinkRequest {
	s.Liveness = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetLokiConfigs(v string) *CreateApplicationShrinkRequest {
	s.LokiConfigs = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetMemory(v int32) *CreateApplicationShrinkRequest {
	s.Memory = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetMicroRegistration(v string) *CreateApplicationShrinkRequest {
	s.MicroRegistration = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetMicroRegistrationConfig(v string) *CreateApplicationShrinkRequest {
	s.MicroRegistrationConfig = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetMicroserviceEngineConfig(v string) *CreateApplicationShrinkRequest {
	s.MicroserviceEngineConfig = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetMountDesc(v string) *CreateApplicationShrinkRequest {
	s.MountDesc = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetMountHost(v string) *CreateApplicationShrinkRequest {
	s.MountHost = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetNamespaceId(v string) *CreateApplicationShrinkRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetNasConfigs(v string) *CreateApplicationShrinkRequest {
	s.NasConfigs = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetNasId(v string) *CreateApplicationShrinkRequest {
	s.NasId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetNewSaeVersion(v string) *CreateApplicationShrinkRequest {
	s.NewSaeVersion = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetOidcRoleName(v string) *CreateApplicationShrinkRequest {
	s.OidcRoleName = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetOssAkId(v string) *CreateApplicationShrinkRequest {
	s.OssAkId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetOssAkSecret(v string) *CreateApplicationShrinkRequest {
	s.OssAkSecret = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetOssMountDescs(v string) *CreateApplicationShrinkRequest {
	s.OssMountDescs = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPackageType(v string) *CreateApplicationShrinkRequest {
	s.PackageType = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPackageUrl(v string) *CreateApplicationShrinkRequest {
	s.PackageUrl = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPackageVersion(v string) *CreateApplicationShrinkRequest {
	s.PackageVersion = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPhp(v string) *CreateApplicationShrinkRequest {
	s.Php = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPhpArmsConfigLocation(v string) *CreateApplicationShrinkRequest {
	s.PhpArmsConfigLocation = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPhpConfig(v string) *CreateApplicationShrinkRequest {
	s.PhpConfig = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPhpConfigLocation(v string) *CreateApplicationShrinkRequest {
	s.PhpConfigLocation = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPostStart(v string) *CreateApplicationShrinkRequest {
	s.PostStart = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPreStop(v string) *CreateApplicationShrinkRequest {
	s.PreStop = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetProgrammingLanguage(v string) *CreateApplicationShrinkRequest {
	s.ProgrammingLanguage = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPvtzDiscoverySvc(v string) *CreateApplicationShrinkRequest {
	s.PvtzDiscoverySvc = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPython(v string) *CreateApplicationShrinkRequest {
	s.Python = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPythonModules(v string) *CreateApplicationShrinkRequest {
	s.PythonModules = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetRaspConfigShrink(v string) *CreateApplicationShrinkRequest {
	s.RaspConfigShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetReadiness(v string) *CreateApplicationShrinkRequest {
	s.Readiness = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetReplicas(v int32) *CreateApplicationShrinkRequest {
	s.Replicas = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetResourceType(v string) *CreateApplicationShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSaeVersion(v string) *CreateApplicationShrinkRequest {
	s.SaeVersion = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSecretMountDesc(v string) *CreateApplicationShrinkRequest {
	s.SecretMountDesc = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSecurityGroupId(v string) *CreateApplicationShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetServiceTags(v string) *CreateApplicationShrinkRequest {
	s.ServiceTags = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSidecarContainersConfigShrink(v string) *CreateApplicationShrinkRequest {
	s.SidecarContainersConfigShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSlsConfigs(v string) *CreateApplicationShrinkRequest {
	s.SlsConfigs = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetSlsLogEnvTags(v string) *CreateApplicationShrinkRequest {
	s.SlsLogEnvTags = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetStartupProbe(v string) *CreateApplicationShrinkRequest {
	s.StartupProbe = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetTerminationGracePeriodSeconds(v int32) *CreateApplicationShrinkRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetTimezone(v string) *CreateApplicationShrinkRequest {
	s.Timezone = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetTomcatConfig(v string) *CreateApplicationShrinkRequest {
	s.TomcatConfig = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetVSwitchId(v string) *CreateApplicationShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetVpcId(v string) *CreateApplicationShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetWarStartOptions(v string) *CreateApplicationShrinkRequest {
	s.WarStartOptions = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetWebContainer(v string) *CreateApplicationShrinkRequest {
	s.WebContainer = &v
	return s
}

func (s *CreateApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}

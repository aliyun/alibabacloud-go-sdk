// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrAssumeRoleArn(v string) *CreateApplicationRequest
	GetAcrAssumeRoleArn() *string
	SetAcrInstanceId(v string) *CreateApplicationRequest
	GetAcrInstanceId() *string
	SetAgentVersion(v string) *CreateApplicationRequest
	GetAgentVersion() *string
	SetAppDescription(v string) *CreateApplicationRequest
	GetAppDescription() *string
	SetAppName(v string) *CreateApplicationRequest
	GetAppName() *string
	SetAppSource(v string) *CreateApplicationRequest
	GetAppSource() *string
	SetAssociateEip(v bool) *CreateApplicationRequest
	GetAssociateEip() *bool
	SetAutoConfig(v bool) *CreateApplicationRequest
	GetAutoConfig() *bool
	SetBaseAppId(v string) *CreateApplicationRequest
	GetBaseAppId() *string
	SetCommand(v string) *CreateApplicationRequest
	GetCommand() *string
	SetCommandArgs(v string) *CreateApplicationRequest
	GetCommandArgs() *string
	SetConfigMapMountDesc(v string) *CreateApplicationRequest
	GetConfigMapMountDesc() *string
	SetCpu(v int32) *CreateApplicationRequest
	GetCpu() *int32
	SetCustomHostAlias(v string) *CreateApplicationRequest
	GetCustomHostAlias() *string
	SetCustomImageNetworkType(v string) *CreateApplicationRequest
	GetCustomImageNetworkType() *string
	SetDeploy(v bool) *CreateApplicationRequest
	GetDeploy() *bool
	SetDiskSize(v int32) *CreateApplicationRequest
	GetDiskSize() *int32
	SetDotnet(v string) *CreateApplicationRequest
	GetDotnet() *string
	SetEdasContainerVersion(v string) *CreateApplicationRequest
	GetEdasContainerVersion() *string
	SetEmptyDirDesc(v string) *CreateApplicationRequest
	GetEmptyDirDesc() *string
	SetEnableCpuBurst(v bool) *CreateApplicationRequest
	GetEnableCpuBurst() *bool
	SetEnableEbpf(v string) *CreateApplicationRequest
	GetEnableEbpf() *string
	SetEnableNamespaceAgentVersion(v bool) *CreateApplicationRequest
	GetEnableNamespaceAgentVersion() *bool
	SetEnableNamespaceSlsConfig(v bool) *CreateApplicationRequest
	GetEnableNamespaceSlsConfig() *bool
	SetEnableNewArms(v bool) *CreateApplicationRequest
	GetEnableNewArms() *bool
	SetEnablePrometheus(v bool) *CreateApplicationRequest
	GetEnablePrometheus() *bool
	SetEnableSidecarResourceIsolated(v bool) *CreateApplicationRequest
	GetEnableSidecarResourceIsolated() *bool
	SetEnvs(v string) *CreateApplicationRequest
	GetEnvs() *string
	SetGpuConfig(v string) *CreateApplicationRequest
	GetGpuConfig() *string
	SetHeadlessPvtzDiscoverySvc(v string) *CreateApplicationRequest
	GetHeadlessPvtzDiscoverySvc() *string
	SetHtml(v string) *CreateApplicationRequest
	GetHtml() *string
	SetImagePullSecrets(v string) *CreateApplicationRequest
	GetImagePullSecrets() *string
	SetImageUrl(v string) *CreateApplicationRequest
	GetImageUrl() *string
	SetInitContainersConfig(v []*InitContainerConfig) *CreateApplicationRequest
	GetInitContainersConfig() []*InitContainerConfig
	SetIsStateful(v bool) *CreateApplicationRequest
	GetIsStateful() *bool
	SetJarStartArgs(v string) *CreateApplicationRequest
	GetJarStartArgs() *string
	SetJarStartOptions(v string) *CreateApplicationRequest
	GetJarStartOptions() *string
	SetJdk(v string) *CreateApplicationRequest
	GetJdk() *string
	SetKafkaConfigs(v string) *CreateApplicationRequest
	GetKafkaConfigs() *string
	SetLabels(v map[string]*string) *CreateApplicationRequest
	GetLabels() map[string]*string
	SetLiveness(v string) *CreateApplicationRequest
	GetLiveness() *string
	SetLokiConfigs(v string) *CreateApplicationRequest
	GetLokiConfigs() *string
	SetMemory(v int32) *CreateApplicationRequest
	GetMemory() *int32
	SetMicroRegistration(v string) *CreateApplicationRequest
	GetMicroRegistration() *string
	SetMicroRegistrationConfig(v string) *CreateApplicationRequest
	GetMicroRegistrationConfig() *string
	SetMicroserviceEngineConfig(v string) *CreateApplicationRequest
	GetMicroserviceEngineConfig() *string
	SetMountDesc(v string) *CreateApplicationRequest
	GetMountDesc() *string
	SetMountHost(v string) *CreateApplicationRequest
	GetMountHost() *string
	SetNamespaceId(v string) *CreateApplicationRequest
	GetNamespaceId() *string
	SetNasConfigs(v string) *CreateApplicationRequest
	GetNasConfigs() *string
	SetNasId(v string) *CreateApplicationRequest
	GetNasId() *string
	SetNewSaeVersion(v string) *CreateApplicationRequest
	GetNewSaeVersion() *string
	SetOidcRoleName(v string) *CreateApplicationRequest
	GetOidcRoleName() *string
	SetOssAkId(v string) *CreateApplicationRequest
	GetOssAkId() *string
	SetOssAkSecret(v string) *CreateApplicationRequest
	GetOssAkSecret() *string
	SetOssMountDescs(v string) *CreateApplicationRequest
	GetOssMountDescs() *string
	SetPackageType(v string) *CreateApplicationRequest
	GetPackageType() *string
	SetPackageUrl(v string) *CreateApplicationRequest
	GetPackageUrl() *string
	SetPackageVersion(v string) *CreateApplicationRequest
	GetPackageVersion() *string
	SetPhp(v string) *CreateApplicationRequest
	GetPhp() *string
	SetPhpArmsConfigLocation(v string) *CreateApplicationRequest
	GetPhpArmsConfigLocation() *string
	SetPhpConfig(v string) *CreateApplicationRequest
	GetPhpConfig() *string
	SetPhpConfigLocation(v string) *CreateApplicationRequest
	GetPhpConfigLocation() *string
	SetPostStart(v string) *CreateApplicationRequest
	GetPostStart() *string
	SetPreStop(v string) *CreateApplicationRequest
	GetPreStop() *string
	SetProgrammingLanguage(v string) *CreateApplicationRequest
	GetProgrammingLanguage() *string
	SetPvtzDiscoverySvc(v string) *CreateApplicationRequest
	GetPvtzDiscoverySvc() *string
	SetPython(v string) *CreateApplicationRequest
	GetPython() *string
	SetPythonModules(v string) *CreateApplicationRequest
	GetPythonModules() *string
	SetRaspConfig(v *CreateApplicationRequestRaspConfig) *CreateApplicationRequest
	GetRaspConfig() *CreateApplicationRequestRaspConfig
	SetReadiness(v string) *CreateApplicationRequest
	GetReadiness() *string
	SetReplicas(v int32) *CreateApplicationRequest
	GetReplicas() *int32
	SetResourceType(v string) *CreateApplicationRequest
	GetResourceType() *string
	SetSaeVersion(v string) *CreateApplicationRequest
	GetSaeVersion() *string
	SetSecretMountDesc(v string) *CreateApplicationRequest
	GetSecretMountDesc() *string
	SetSecurityGroupId(v string) *CreateApplicationRequest
	GetSecurityGroupId() *string
	SetServiceTags(v string) *CreateApplicationRequest
	GetServiceTags() *string
	SetSidecarContainersConfig(v []*SidecarContainerConfig) *CreateApplicationRequest
	GetSidecarContainersConfig() []*SidecarContainerConfig
	SetSlsConfigs(v string) *CreateApplicationRequest
	GetSlsConfigs() *string
	SetSlsLogEnvTags(v string) *CreateApplicationRequest
	GetSlsLogEnvTags() *string
	SetStartupProbe(v string) *CreateApplicationRequest
	GetStartupProbe() *string
	SetTerminationGracePeriodSeconds(v int32) *CreateApplicationRequest
	GetTerminationGracePeriodSeconds() *int32
	SetTimezone(v string) *CreateApplicationRequest
	GetTimezone() *string
	SetTomcatConfig(v string) *CreateApplicationRequest
	GetTomcatConfig() *string
	SetVSwitchId(v string) *CreateApplicationRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateApplicationRequest
	GetVpcId() *string
	SetWarStartOptions(v string) *CreateApplicationRequest
	GetWarStartOptions() *string
	SetWebContainer(v string) *CreateApplicationRequest
	GetWebContainer() *string
}

type CreateApplicationRequest struct {
	// The ARN of the RAM role required to pull images across accounts. For more information, see [Use RAM roles to authorize access across Alibaba Cloud accounts](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The Container Registry ACR Enterprise instance ID. This parameter is required when **ImageUrl*	- is set to a Container Registry Enterprise instance image.
	//
	// example:
	//
	// cri-xxxxxx
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The Aliyun Agent version.
	//
	// example:
	//
	// 4.4.2
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The description of the application. The description can be up to 1,024 characters in length.
	//
	// example:
	//
	// This is a test description.
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The name of the application. The name can contain digits, letters, and hyphens (-). It must start with a letter and cannot end with a hyphen (-). The name can be up to 36 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The source of the application. A value of micro_service indicates a microservice application.
	//
	// example:
	//
	// micro_service
	AppSource *string `json:"AppSource,omitempty" xml:"AppSource,omitempty"`
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
	// Specifies whether to automatically configure the network environment. Valid values:
	//
	// - **true**: SAE automatically configures the network environment when the application is created. The values of **NamespaceId**, **VpcId**, **vSwitchId**, and **SecurityGroupId*	- are ignored.
	//
	// - **false**: SAE manually configures the network environment when the application is created.
	//
	// > If you set this parameter to **true**, other **NamespaceId*	- values passed are ignored.
	//
	// example:
	//
	// true
	AutoConfig *bool `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty"`
	// The ID of the base application.
	//
	// example:
	//
	// ee99cce6-1c8e-4bfa-96c3-3e2fa9de8a41
	BaseAppId *string `json:"BaseAppId,omitempty" xml:"BaseAppId,omitempty"`
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
	// Based on the preceding example, Command="echo", CommandArgs=["abc", ">", "file0"].
	//
	//
	// 	Notice: This parameter is required when PackageType is set to DotnetZip.</notice>
	//
	// example:
	//
	// echo
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments for the image startup command, which are required by the startup command **Command**. The format is as follows:
	//
	// `["a","b"]`
	//
	// In the preceding example, `CommandArgs=["abc", ">", "file0"]`, where `["abc", ">", "file0"]` must be converted to a string in JSON array format. If no arguments are required, leave this parameter empty.
	//
	// <notice>This parameter is required when PackageType is set to DotnetZip.</notice>
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The mount description of the **ConfigMap**. Uses the configuration item created on the configuration item page of the namespace to inject configuration information into the container. Parameter description:
	//
	// - **configMapId**: instance ID of the ConfigMap instance. You can invoke the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the instance ID.
	//
	// - **key**: the key value.
	//
	// > You can mount all keys by passing the sae-sys-configmap-all parameter.
	//
	// - **mountPath**: the mount path.
	//
	// example:
	//
	// [{"configMapId":16,"key":"test","mountPath":"/tmp"}]
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	// The number of CPU cores required for each instance. Unit: millicore. The value cannot be 0. Only the following defined specifications are supported:
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
	// The custom image type. If the image is not a custom image, set this parameter to an empty string:
	//
	// - internet: public network image.
	//
	// - intranet: private network image.
	//
	// example:
	//
	// internet
	CustomImageNetworkType *string `json:"CustomImageNetworkType,omitempty" xml:"CustomImageNetworkType,omitempty"`
	// Specifies whether to deploy the application immediately. Valid values:
	//
	// - **true**: deploys the application immediately. This is the default value.
	//
	// - **false**: deploys the application later.
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
	// The version number of the .NET framework:
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
	// The version of the application runtime environment in the High-Speed Service Framework (HSF), such as the Ali-Tomcat container.
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
	// - true: enables the feature.
	//
	// - false: disables the feature.
	//
	// example:
	//
	// true
	EnableCpuBurst *bool `json:"EnableCpuBurst,omitempty" xml:"EnableCpuBurst,omitempty"`
	// Specifies whether to enable application monitoring for non-Java applications based on eBPF technology. Valid values:
	//
	// - **true**: enables the feature.
	//
	// - **false**: disables the feature. This is the default value.
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
	// Specifies whether to reuse the namespace Simple Log Service (SLS) log configuration.
	//
	// example:
	//
	// true
	EnableNamespaceSlsConfig *bool `json:"EnableNamespaceSlsConfig,omitempty" xml:"EnableNamespaceSlsConfig,omitempty"`
	// Specifies whether to enable the new Application Real-Time Monitoring Service (ARMS) features:
	//
	// - true: enables the features.
	//
	// - false: disables the features.
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
	// - true: enables isolation.
	//
	// - false: disables isolation.
	//
	// example:
	//
	// true
	EnableSidecarResourceIsolated *bool `json:"EnableSidecarResourceIsolated,omitempty" xml:"EnableSidecarResourceIsolated,omitempty"`
	// The environment variables of the container. You can customize environment variables or reference ConfigMaps. To reference a ConfigMap, you must create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//     - **name**: the name of the environment variable.
	//
	//     - **value**: the value of the environment variable. The priority of this parameter is higher than that of valueFrom.
	//
	// - Reference a ConfigMap (valueFrom)
	//
	//     - **name**: the name of the environment variable. You can reference a single key or all keys. To reference all keys, enter sae-sys-configmap-all-<ConfigMap name>, for example, sae-sys-configmap-all-test1.
	//
	//     - **valueFrom**: the reference of the environment variable. Set the value to configMapRef.
	//
	//         - **configMapId**: the ID of the ConfigMap.
	//
	//         - **key**: the key. If you reference all key-value pairs, do not set this field.
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
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The GPU configuration.
	//
	// example:
	//
	// {\\n\\"gpuType\\":\\"PPU810E\\",\\n\\"gpuCount\\":\\"1\\"\\n}
	GpuConfig *string `json:"GpuConfig,omitempty" xml:"GpuConfig,omitempty"`
	// The Kubernetes headless service registration and discovery configuration.
	//
	// - serviceName: The service name.
	//
	// - namespaceId: The namespace ID.
	//
	// example:
	//
	// {\\"serviceName\\":\\"leaf-test-headless\\",\\"namespaceId\\":\\"cn-zhangjiakou:prod\\"}
	HeadlessPvtzDiscoverySvc *string `json:"HeadlessPvtzDiscoverySvc,omitempty" xml:"HeadlessPvtzDiscoverySvc,omitempty"`
	// The NGINX version. Valid values:
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
	// The ID of the corresponding Secret.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The URL of the image. This parameter is required when **PackageType*	- is set to **Image**.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/sae_test/ali_sae_test:0.0.1
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The init container configuration.
	InitContainersConfig []*InitContainerConfig `json:"InitContainersConfig,omitempty" xml:"InitContainersConfig,omitempty" type:"Repeated"`
	// Specifies whether the application is stateful.
	//
	// example:
	//
	// true
	IsStateful *bool `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The startup arguments for the JAR package. Default startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// custom-args
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The startup options for the JAR package. Default startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
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
	// This parameter is not supported when **PackageType*	- is set to **Image**.
	//
	// example:
	//
	// Open JDK 8
	Jdk *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	// The summary configuration for collecting logs to Kafka. Valid values:
	//
	// - **kafkaEndpoint**: the service registration address of the Kafka API.
	//
	// - **kafkaInstanceId**: the Kafka instance ID.
	//
	// - **kafkaConfigs**: the summary configuration for one or more logs. For value description, see the **kafkaConfigs*	- request parameter in this topic.
	//
	// example:
	//
	// {"kafkaEndpoint":"10.0.X.XXX:XXXX,10.0.X.XXX:XXXX,10.0.X.XXX:XXXX","kafkaInstanceId":"alikafka_pre-cn-7pp2l8kr****","kafkaConfigs":[{"logType":"file_log","logDir":"/tmp/a.log","kafkaTopic":"test2"},{"logType":"stdout","logDir":"","kafkaTopic":"test"}]}
	KafkaConfigs *string `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	// The application labels. Both keys and values are strings.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The liveness probe of the container. A container that failed the health check undergoes shutdown and recovers. The following methods are supported:
	//
	// - **exec**: for example, {"exec":{"command":["sh","-c","cat/home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	//
	// - **httpGet**: for example, {"httpGet":{"path":"/","port":18091,"scheme":"HTTP","isContainKeyWord":true,"keyWord":"SAE"},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}
	//
	// - **tcpSocket**: for example, {"tcpSocket":{"port":18091},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}
	//
	// > You can select only one method for the health check.
	//
	// Parameter description:
	//
	// - **exec.command**: Settings for the health check command.
	//
	// - **httpGet.path**: the access path.
	//
	// - **httpGet.scheme**: **HTTP*	- or **HTTPS**.
	//
	// - **httpGet.isContainKeyWord**: **true*	- indicates that the keyword is included, and **false*	- indicates that the keyword is not included. If this field is missing, the advanced feature is not used.
	//
	// - **httpGet.keyWord**: the custom keyword. The **isContainKeyWord*	- field is required when this field is used.
	//
	// - **tcpSocket.port**: the port for TCP connection detection.
	//
	// - **initialDelaySeconds**: Settings for the delay before the health check starts. Default value: 10. Unit: seconds.
	//
	// - **periodSeconds**: Settings for the health check epoch. Default value: 30. Unit: seconds.
	//
	// - **timeoutSeconds**: Settings for the timeout period of the health check. Default value: 1. Unit: seconds. If this parameter is set to 0 or not set, the default timeout period is 1 second.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// LokiConfigs
	//
	// example:
	//
	// []
	LokiConfigs *string `json:"LokiConfigs,omitempty" xml:"LokiConfigs,omitempty"`
	// The memory size required for each instance. Unit: MB. The value cannot be 0. The memory size has a one-to-one correspondence with the number of CPU cores. Only the following defined specifications are supported:
	//
	// - **1024**: corresponds to 500 and 1,000 millicores.
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
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The Nacos registry. Valid values:
	//
	// - **0**: the built-in Nacos of SAE.
	//
	// - **1**: a self-managed Nacos.
	//
	// - **2**: Microservices Engine (MSE) Nacos.
	//
	// example:
	//
	// "0"
	MicroRegistration *string `json:"MicroRegistration,omitempty" xml:"MicroRegistration,omitempty"`
	// The configuration of the registry.
	//
	// example:
	//
	// {\\"instanceId\\":\\"mse-cn-zvp2bh6h70r\\",\\"namespace\\":\\"4c0aa74f-57cb-423c-b6af-5d9f2d0e3dbd\\"}
	MicroRegistrationConfig *string `json:"MicroRegistrationConfig,omitempty" xml:"MicroRegistrationConfig,omitempty"`
	// Configures microservice governance features.
	//
	// - enable: Specifies whether to enable microservice governance. Valid values:
	//
	//   - true: Enabled.
	//
	//   - false: Disabled.
	//
	// - mseLosslessRule: Configures lossless online and offline rules.
	//
	//   - delayTime: The delay time.
	//
	//   - enable: Specifies whether to enable the lossless online feature. Valid values: true (enabled) and false (disabled).
	//
	//   - notice: Specifies whether to enable the notification feature. Valid values: true (enabled) and false (disabled).
	//
	//   - warmupTime: The warm-up duration for low traffic, in seconds.
	//
	// example:
	//
	// {"enable": true,"mseLosslessRule": {"delayTime": 0,"enable": false,"notice": false,"warmupTime": 120}}
	MicroserviceEngineConfig *string `json:"MicroserviceEngineConfig,omitempty" xml:"MicroserviceEngineConfig,omitempty"`
	// Do not configure this field. Configure **NasConfigs*	- instead. The mount description of the NAS file system. If the configuration remains unchanged during deployment, you do not need to set this parameter. That is, you do not need to include the **MountDesc*	- field in the request. If you need to clear the NAS configuration, set the value of this field to an empty string in the request. That is, set the **MountDesc*	- field to "".
	//
	// example:
	//
	// [{mountPath: "/tmp", nasPath: "/"}]
	MountDesc *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	// Do not configure this field. Configure **NasConfigs*	- instead. The mount target of the NAS file system in the VPC of the application. If the configuration remains unchanged during deployment, you do not need to set this parameter. That is, you do not need to include the **MountHost*	- field in the request. If you need to clear the NAS configuration, set the value of this field to an empty string in the request. That is, set the **MountHost*	- field to "".
	//
	// example:
	//
	// example.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The ID of the Serverless App Engine (SAE) namespace. Only namespaces whose names contain lowercase letters and hyphens (-) are supported. The name must start with a letter.
	//
	// You can obtain the namespace by calling the [DescribeNamespaceList](https://help.aliyun.com/document_detail/126547.html) operation.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The configuration for mounting the NAS file system. Valid values:
	//
	// - **mountPath**: the mount path of the container.
	//
	// - **readOnly**: A value of **false*	- indicates read and write permissions.
	//
	// - **nasId**: the ID of the NAS file system.
	//
	// - **mountDomain**: the mount target address of the container. For more information, see [DescribeMountTargets](https://help.aliyun.com/document_detail/62626.html).
	//
	// - **nasPath**: the relative file directory of the NAS file system.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// Do not configure this field. Configure **NasConfigs*	- instead. The ID of the mounted Apsara File Storage NAS file system. The NAS file system must reside in the same region as the cluster. It must have an available quota for creating mount targets, or its mount targets must already exist on a vSwitch in the VPC. If you do not specify this parameter and the **mountDescs*	- field exists, a NAS file system is automatically purchased and mounted to a vSwitch in the VPC by default.
	//
	// If the configuration remains unchanged during deployment, you do not need to set this parameter. That is, you do not need to include the **NASId*	- field in the request. If you need to clear the NAS configuration, set the value of this field to an empty string in the request. That is, set the **NASId*	- field to "".
	//
	// example:
	//
	// KSAK****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The application version. Valid values:
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
	// > You must create an OIDC identity provider and an identity provider role in the same region in advance. For more information, see <props="china">[Create an OIDC identity provider](https://www.alibabacloud.com/help/en/ram/developer-reference/api-ims-2019-08-15-createoidcprovider) and [Create a role-based SSO identity provider](https://www.alibabacloud.com/help/en/ram/developer-reference/api-ims-2019-08-15-createsamlprovider)<props="intl">[Create an OIDC identity provider](https://www.alibabacloud.com/help/zh/ram/developer-reference/api-ims-2019-08-15-createoidcprovider) and [Create a role-based SSO identity provider](https://www.alibabacloud.com/help/zh/ram/developer-reference/api-ims-2019-08-15-createsamlprovider).
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
	// The OSS mount description. Parameter description:
	//
	// - **bucketName**: the name of the bucket.
	//
	// - **bucketPath**: the folder or object that you created in OSS. If the OSS mount folder does not exist, an exception is triggered.
	//
	// - **mountPath**: the container path in SAE. If the path already exists, it will overwrite the existing content. If the path does not exist, it is created.
	//
	// - **readOnly**: specifies whether the container path has read permissions on the mounted folder resources. Valid values:
	//
	//     - **true**: read-only permissions.
	//
	//     - **false**: read and write permissions.
	//
	// example:
	//
	// [{"bucketName": "oss-bucket", "bucketPath": "data/user.data", "mountPath": "/usr/data/user.data", "readOnly": true}]
	OssMountDescs *string `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty"`
	// The type of the application package. Valid values:
	//
	// - When you deploy the application by using Java, **FatJar**, **War**, and **Image*	- are supported.
	//
	// - When you deploy the application by using PHP, the following types are supported:
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
	// - When you deploy the application by using Python, **PythonZip*	- and **Image*	- are supported.
	//
	// - When you deploy the application by using .NET Core, **DotnetZip*	- and **Image*	- are supported.
	//
	//   >
	//
	//   > If you select DotnetZip, Dotnet specifies the version number of the .NET Core environment. .NET 3.1, .NET 5.0, .NET 6.0, .NET 7.0, and .NET 8.0 are supported. In this case, the Dotnet, Command, and CommandArgs parameters are required.
	//
	// This parameter is required.
	//
	// example:
	//
	// FatJar
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The URL of the deployment package. This parameter is required when **PackageType*	- is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// http://myoss.oss-cn-****.aliyuncs.com/my-buc/2019-06-30/****.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version number of the deployment package. This parameter is required when **PackageType*	- is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// 1.0.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The PHP version on which the PHP deployment package depends. Images are not supported.
	//
	// example:
	//
	// PHP-FPM 7.0
	Php *string `json:"Php,omitempty" xml:"Php,omitempty"`
	// The mount path for PHP application monitoring. You must ensure that the PHP server loads the configuration file in this path.
	//
	// You do not need to pay attention to the configuration content. SAE automatically renders the correct configuration file.
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
	// The script to execute after the container starts. A script is triggered and executed immediately after the container is created. Format: {"exec":{"command":["cat","/etc/group"]}}
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script to execute before the container stops. A script is triggered and executed before the container is deleted. Format: {"exec":{"command":["cat","/etc/group"]}}
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
	// Enables Kubernetes Service registration and discovery. Valid values:
	//
	// - **serviceName**: the service name. The format is custom-namespace ID. The suffix -namespace ID cannot be customized and must be filled in based on the namespace where the application resides. For example, if you select the default namespace in the China (Beijing) region, the value is -cn-beijing-default.
	//
	// - **namespaceId**: the ID of the namespace.
	//
	// - **portProtocols**: the port and protocol. The port value ranges from 1 to 65535. The protocol supports **TCP*	- and **UDP**.
	//
	// - **portAndProtocol**: the port and protocol. The port value ranges from 1 to 65535. The protocol supports **TCP*	- and **UDP**. **We recommend that you use portProtocols. If portProtocols Settings are configured, only portProtocols takes effect**.
	//
	// - **enable**: enables Kubernetes Service registration and discovery.
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
	// The custom installation module dependencies. By default, the dependencies defined in the requirements.txt file in the root folder are installed. If no dependencies are configured or custom software packages are used, you can specify the dependencies to install.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// The application runtime security protection (RASP) configuration.
	RaspConfig *CreateApplicationRequestRaspConfig `json:"RaspConfig,omitempty" xml:"RaspConfig,omitempty" type:"Struct"`
	// The readiness probe of the application. A container that fails multiple health checks is stopped and restarted. Traffic from Server Load Balancer (SLB) is not routed to containers that fail the health check. The **exec**, **httpGet**, and **tcpSocket*	- methods are supported. For specific examples, see the **Liveness*	- parameter.
	//
	// > You can select only one method for the health check.
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
	// The resource type. Valid values: NULL (default), default (default), and haiguang (Hygon servers).
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
	// The mount description of the **Secret**. Uses the Secret created on the Secret page of the namespace to inject confidential information into the container. Parameter description:
	//
	// - **secretId**: instance ID of the Secret instance. You can obtain the instance ID by invoking the ListSecrets operation.
	//
	// - **key**: the key value.
	//
	// > You can mount all keys by passing the sae-sys-secret-all parameter.
	//
	// - **mountPath**: the mount path.
	//
	// example:
	//
	// [{“secretId":10,”key":"test","mountPath":"/tmp"}]
	SecretMountDesc *string `json:"SecretMountDesc,omitempty" xml:"SecretMountDesc,omitempty"`
	// The ID of the security group.
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
	// The container configuration.
	SidecarContainersConfig []*SidecarContainerConfig `json:"SidecarContainersConfig,omitempty" xml:"SidecarContainersConfig,omitempty" type:"Repeated"`
	// The configuration for collecting logs to Simple Log Service (SLS).
	//
	// - Use SLS resources from automatic creation by SAE: [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}].
	//
	// - Use custom SLS resources: [{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}].
	//
	// Parameter description:
	//
	// - **projectName**: the name of the SLS project.
	//
	// - **logDir**: the log path.
	//
	// - **logType**: the log type. **stdout*	- indicates container standard output logs. Only one stdout log Settings can be configured. If Settings are not specified, file logs are collected.
	//
	// - **logstoreName**: the name of the SLS Logstore.
	//
	// - **logtailName**: the name of the SLS Logtail. If this parameter is not specified, a new Logtail is created.
	//
	// If the SLS collection configuration remains unchanged during multiple deployments, you do not need to set this parameter. That is, you do not need to include the **SlsConfigs*	- field in the request. If you no longer need the SLS collection feature, you must set the value of this field to an empty string in the request. That is, set the **SlsConfigs*	- field to "".
	//
	// > Projects from automatic creation with the application are deleted when the application is deleted. Therefore, when selecting an existing project, you cannot select a project from automatic creation by SAE.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// sls log tags
	//
	// example:
	//
	// []
	SlsLogEnvTags *string `json:"SlsLogEnvTags,omitempty" xml:"SlsLogEnvTags,omitempty"`
	// Enables application startup probes.
	//
	// - Success: Indicates that the application is started. If you configure liveness and readiness probes, the liveness and readiness probes are executed after the application is started.
	//
	// - Failure: Indicates that the application fails to start. An exception is reported and the application is automatically restarted.
	//
	// >
	//
	// > - The exec, httpGet, and tcpSocket methods are supported. For specific examples, refer to the Liveness parameter.
	//
	// > - You can select only one method for health checks.
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
	// The Tomcat file configuration. Set the value to "" or "{}" to delete the configuration:
	//
	// - **port**: the port range is 1024 to 65535. Ports smaller than 1024 require root permissions. Because the container is configured with admin permissions, specify a port greater than 1024. If this parameter is not configured, the default value is 8080.
	//
	// - **contextPath**: the access path. The default value is the root directory "/".
	//
	// - **maxThreads**: the maximum number of connections in the connection pool. The default value is 400.
	//
	// - **uriEncoding**: the encoding format of Tomcat. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. If this parameter is not set, the default value is **ISO-8859-1**.
	//
	// - **useBodyEncodingForUri**: specifies whether to use **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	// The ID of the vSwitch to which the elastic network interface controllers (NICs) of the application instance belong. The vSwitch must reside in the preceding VPC. The vSwitch is also attached to the SAE namespace. If you do not specify this parameter, the ID of the vSwitch attached to the namespace is used by default.
	//
	// example:
	//
	// vsw-bp12mw1f8k3jgygk9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) that corresponds to the SAE namespace. In SAE, a namespace can correspond to only one VPC, and the correspondence cannot be modified. The binding relationship is established when you create the first SAE application in the namespace. Multiple namespaces can correspond to one VPC. If you do not specify this parameter, the ID of the VPC bound to the namespace is used by default.
	//
	// example:
	//
	// vpc-bp1aevy8sofi8mh1q****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The startup command for the application deployed by using a WAR package. The procedure is the same as that for configuring the startup command for image deployment. For more information, see [Set a startup command](https://help.aliyun.com/document_detail/96677.html).
	//
	// example:
	//
	// CATALINA_OPTS=\\"$CATALINA_OPTS $Options\\" catalina.sh run
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The Tomcat version on which the WebContainer deployment package depends. Valid values:
	//
	// - **apache-tomcat-7.0.91**
	//
	// - **apache-tomcat-8.5.42**
	//
	// This parameter is not supported when **PackageType*	- is set to **Image**.
	//
	// example:
	//
	// apache-tomcat-7.0.91
	WebContainer *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) GetAcrAssumeRoleArn() *string {
	return s.AcrAssumeRoleArn
}

func (s *CreateApplicationRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CreateApplicationRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *CreateApplicationRequest) GetAppDescription() *string {
	return s.AppDescription
}

func (s *CreateApplicationRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateApplicationRequest) GetAppSource() *string {
	return s.AppSource
}

func (s *CreateApplicationRequest) GetAssociateEip() *bool {
	return s.AssociateEip
}

func (s *CreateApplicationRequest) GetAutoConfig() *bool {
	return s.AutoConfig
}

func (s *CreateApplicationRequest) GetBaseAppId() *string {
	return s.BaseAppId
}

func (s *CreateApplicationRequest) GetCommand() *string {
	return s.Command
}

func (s *CreateApplicationRequest) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *CreateApplicationRequest) GetConfigMapMountDesc() *string {
	return s.ConfigMapMountDesc
}

func (s *CreateApplicationRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateApplicationRequest) GetCustomHostAlias() *string {
	return s.CustomHostAlias
}

func (s *CreateApplicationRequest) GetCustomImageNetworkType() *string {
	return s.CustomImageNetworkType
}

func (s *CreateApplicationRequest) GetDeploy() *bool {
	return s.Deploy
}

func (s *CreateApplicationRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreateApplicationRequest) GetDotnet() *string {
	return s.Dotnet
}

func (s *CreateApplicationRequest) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *CreateApplicationRequest) GetEmptyDirDesc() *string {
	return s.EmptyDirDesc
}

func (s *CreateApplicationRequest) GetEnableCpuBurst() *bool {
	return s.EnableCpuBurst
}

func (s *CreateApplicationRequest) GetEnableEbpf() *string {
	return s.EnableEbpf
}

func (s *CreateApplicationRequest) GetEnableNamespaceAgentVersion() *bool {
	return s.EnableNamespaceAgentVersion
}

func (s *CreateApplicationRequest) GetEnableNamespaceSlsConfig() *bool {
	return s.EnableNamespaceSlsConfig
}

func (s *CreateApplicationRequest) GetEnableNewArms() *bool {
	return s.EnableNewArms
}

func (s *CreateApplicationRequest) GetEnablePrometheus() *bool {
	return s.EnablePrometheus
}

func (s *CreateApplicationRequest) GetEnableSidecarResourceIsolated() *bool {
	return s.EnableSidecarResourceIsolated
}

func (s *CreateApplicationRequest) GetEnvs() *string {
	return s.Envs
}

func (s *CreateApplicationRequest) GetGpuConfig() *string {
	return s.GpuConfig
}

func (s *CreateApplicationRequest) GetHeadlessPvtzDiscoverySvc() *string {
	return s.HeadlessPvtzDiscoverySvc
}

func (s *CreateApplicationRequest) GetHtml() *string {
	return s.Html
}

func (s *CreateApplicationRequest) GetImagePullSecrets() *string {
	return s.ImagePullSecrets
}

func (s *CreateApplicationRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateApplicationRequest) GetInitContainersConfig() []*InitContainerConfig {
	return s.InitContainersConfig
}

func (s *CreateApplicationRequest) GetIsStateful() *bool {
	return s.IsStateful
}

func (s *CreateApplicationRequest) GetJarStartArgs() *string {
	return s.JarStartArgs
}

func (s *CreateApplicationRequest) GetJarStartOptions() *string {
	return s.JarStartOptions
}

func (s *CreateApplicationRequest) GetJdk() *string {
	return s.Jdk
}

func (s *CreateApplicationRequest) GetKafkaConfigs() *string {
	return s.KafkaConfigs
}

func (s *CreateApplicationRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *CreateApplicationRequest) GetLiveness() *string {
	return s.Liveness
}

func (s *CreateApplicationRequest) GetLokiConfigs() *string {
	return s.LokiConfigs
}

func (s *CreateApplicationRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateApplicationRequest) GetMicroRegistration() *string {
	return s.MicroRegistration
}

func (s *CreateApplicationRequest) GetMicroRegistrationConfig() *string {
	return s.MicroRegistrationConfig
}

func (s *CreateApplicationRequest) GetMicroserviceEngineConfig() *string {
	return s.MicroserviceEngineConfig
}

func (s *CreateApplicationRequest) GetMountDesc() *string {
	return s.MountDesc
}

func (s *CreateApplicationRequest) GetMountHost() *string {
	return s.MountHost
}

func (s *CreateApplicationRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateApplicationRequest) GetNasConfigs() *string {
	return s.NasConfigs
}

func (s *CreateApplicationRequest) GetNasId() *string {
	return s.NasId
}

func (s *CreateApplicationRequest) GetNewSaeVersion() *string {
	return s.NewSaeVersion
}

func (s *CreateApplicationRequest) GetOidcRoleName() *string {
	return s.OidcRoleName
}

func (s *CreateApplicationRequest) GetOssAkId() *string {
	return s.OssAkId
}

func (s *CreateApplicationRequest) GetOssAkSecret() *string {
	return s.OssAkSecret
}

func (s *CreateApplicationRequest) GetOssMountDescs() *string {
	return s.OssMountDescs
}

func (s *CreateApplicationRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *CreateApplicationRequest) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *CreateApplicationRequest) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *CreateApplicationRequest) GetPhp() *string {
	return s.Php
}

func (s *CreateApplicationRequest) GetPhpArmsConfigLocation() *string {
	return s.PhpArmsConfigLocation
}

func (s *CreateApplicationRequest) GetPhpConfig() *string {
	return s.PhpConfig
}

func (s *CreateApplicationRequest) GetPhpConfigLocation() *string {
	return s.PhpConfigLocation
}

func (s *CreateApplicationRequest) GetPostStart() *string {
	return s.PostStart
}

func (s *CreateApplicationRequest) GetPreStop() *string {
	return s.PreStop
}

func (s *CreateApplicationRequest) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *CreateApplicationRequest) GetPvtzDiscoverySvc() *string {
	return s.PvtzDiscoverySvc
}

func (s *CreateApplicationRequest) GetPython() *string {
	return s.Python
}

func (s *CreateApplicationRequest) GetPythonModules() *string {
	return s.PythonModules
}

func (s *CreateApplicationRequest) GetRaspConfig() *CreateApplicationRequestRaspConfig {
	return s.RaspConfig
}

func (s *CreateApplicationRequest) GetReadiness() *string {
	return s.Readiness
}

func (s *CreateApplicationRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *CreateApplicationRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateApplicationRequest) GetSaeVersion() *string {
	return s.SaeVersion
}

func (s *CreateApplicationRequest) GetSecretMountDesc() *string {
	return s.SecretMountDesc
}

func (s *CreateApplicationRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateApplicationRequest) GetServiceTags() *string {
	return s.ServiceTags
}

func (s *CreateApplicationRequest) GetSidecarContainersConfig() []*SidecarContainerConfig {
	return s.SidecarContainersConfig
}

func (s *CreateApplicationRequest) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *CreateApplicationRequest) GetSlsLogEnvTags() *string {
	return s.SlsLogEnvTags
}

func (s *CreateApplicationRequest) GetStartupProbe() *string {
	return s.StartupProbe
}

func (s *CreateApplicationRequest) GetTerminationGracePeriodSeconds() *int32 {
	return s.TerminationGracePeriodSeconds
}

func (s *CreateApplicationRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateApplicationRequest) GetTomcatConfig() *string {
	return s.TomcatConfig
}

func (s *CreateApplicationRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateApplicationRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateApplicationRequest) GetWarStartOptions() *string {
	return s.WarStartOptions
}

func (s *CreateApplicationRequest) GetWebContainer() *string {
	return s.WebContainer
}

func (s *CreateApplicationRequest) SetAcrAssumeRoleArn(v string) *CreateApplicationRequest {
	s.AcrAssumeRoleArn = &v
	return s
}

func (s *CreateApplicationRequest) SetAcrInstanceId(v string) *CreateApplicationRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateApplicationRequest) SetAgentVersion(v string) *CreateApplicationRequest {
	s.AgentVersion = &v
	return s
}

func (s *CreateApplicationRequest) SetAppDescription(v string) *CreateApplicationRequest {
	s.AppDescription = &v
	return s
}

func (s *CreateApplicationRequest) SetAppName(v string) *CreateApplicationRequest {
	s.AppName = &v
	return s
}

func (s *CreateApplicationRequest) SetAppSource(v string) *CreateApplicationRequest {
	s.AppSource = &v
	return s
}

func (s *CreateApplicationRequest) SetAssociateEip(v bool) *CreateApplicationRequest {
	s.AssociateEip = &v
	return s
}

func (s *CreateApplicationRequest) SetAutoConfig(v bool) *CreateApplicationRequest {
	s.AutoConfig = &v
	return s
}

func (s *CreateApplicationRequest) SetBaseAppId(v string) *CreateApplicationRequest {
	s.BaseAppId = &v
	return s
}

func (s *CreateApplicationRequest) SetCommand(v string) *CreateApplicationRequest {
	s.Command = &v
	return s
}

func (s *CreateApplicationRequest) SetCommandArgs(v string) *CreateApplicationRequest {
	s.CommandArgs = &v
	return s
}

func (s *CreateApplicationRequest) SetConfigMapMountDesc(v string) *CreateApplicationRequest {
	s.ConfigMapMountDesc = &v
	return s
}

func (s *CreateApplicationRequest) SetCpu(v int32) *CreateApplicationRequest {
	s.Cpu = &v
	return s
}

func (s *CreateApplicationRequest) SetCustomHostAlias(v string) *CreateApplicationRequest {
	s.CustomHostAlias = &v
	return s
}

func (s *CreateApplicationRequest) SetCustomImageNetworkType(v string) *CreateApplicationRequest {
	s.CustomImageNetworkType = &v
	return s
}

func (s *CreateApplicationRequest) SetDeploy(v bool) *CreateApplicationRequest {
	s.Deploy = &v
	return s
}

func (s *CreateApplicationRequest) SetDiskSize(v int32) *CreateApplicationRequest {
	s.DiskSize = &v
	return s
}

func (s *CreateApplicationRequest) SetDotnet(v string) *CreateApplicationRequest {
	s.Dotnet = &v
	return s
}

func (s *CreateApplicationRequest) SetEdasContainerVersion(v string) *CreateApplicationRequest {
	s.EdasContainerVersion = &v
	return s
}

func (s *CreateApplicationRequest) SetEmptyDirDesc(v string) *CreateApplicationRequest {
	s.EmptyDirDesc = &v
	return s
}

func (s *CreateApplicationRequest) SetEnableCpuBurst(v bool) *CreateApplicationRequest {
	s.EnableCpuBurst = &v
	return s
}

func (s *CreateApplicationRequest) SetEnableEbpf(v string) *CreateApplicationRequest {
	s.EnableEbpf = &v
	return s
}

func (s *CreateApplicationRequest) SetEnableNamespaceAgentVersion(v bool) *CreateApplicationRequest {
	s.EnableNamespaceAgentVersion = &v
	return s
}

func (s *CreateApplicationRequest) SetEnableNamespaceSlsConfig(v bool) *CreateApplicationRequest {
	s.EnableNamespaceSlsConfig = &v
	return s
}

func (s *CreateApplicationRequest) SetEnableNewArms(v bool) *CreateApplicationRequest {
	s.EnableNewArms = &v
	return s
}

func (s *CreateApplicationRequest) SetEnablePrometheus(v bool) *CreateApplicationRequest {
	s.EnablePrometheus = &v
	return s
}

func (s *CreateApplicationRequest) SetEnableSidecarResourceIsolated(v bool) *CreateApplicationRequest {
	s.EnableSidecarResourceIsolated = &v
	return s
}

func (s *CreateApplicationRequest) SetEnvs(v string) *CreateApplicationRequest {
	s.Envs = &v
	return s
}

func (s *CreateApplicationRequest) SetGpuConfig(v string) *CreateApplicationRequest {
	s.GpuConfig = &v
	return s
}

func (s *CreateApplicationRequest) SetHeadlessPvtzDiscoverySvc(v string) *CreateApplicationRequest {
	s.HeadlessPvtzDiscoverySvc = &v
	return s
}

func (s *CreateApplicationRequest) SetHtml(v string) *CreateApplicationRequest {
	s.Html = &v
	return s
}

func (s *CreateApplicationRequest) SetImagePullSecrets(v string) *CreateApplicationRequest {
	s.ImagePullSecrets = &v
	return s
}

func (s *CreateApplicationRequest) SetImageUrl(v string) *CreateApplicationRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateApplicationRequest) SetInitContainersConfig(v []*InitContainerConfig) *CreateApplicationRequest {
	s.InitContainersConfig = v
	return s
}

func (s *CreateApplicationRequest) SetIsStateful(v bool) *CreateApplicationRequest {
	s.IsStateful = &v
	return s
}

func (s *CreateApplicationRequest) SetJarStartArgs(v string) *CreateApplicationRequest {
	s.JarStartArgs = &v
	return s
}

func (s *CreateApplicationRequest) SetJarStartOptions(v string) *CreateApplicationRequest {
	s.JarStartOptions = &v
	return s
}

func (s *CreateApplicationRequest) SetJdk(v string) *CreateApplicationRequest {
	s.Jdk = &v
	return s
}

func (s *CreateApplicationRequest) SetKafkaConfigs(v string) *CreateApplicationRequest {
	s.KafkaConfigs = &v
	return s
}

func (s *CreateApplicationRequest) SetLabels(v map[string]*string) *CreateApplicationRequest {
	s.Labels = v
	return s
}

func (s *CreateApplicationRequest) SetLiveness(v string) *CreateApplicationRequest {
	s.Liveness = &v
	return s
}

func (s *CreateApplicationRequest) SetLokiConfigs(v string) *CreateApplicationRequest {
	s.LokiConfigs = &v
	return s
}

func (s *CreateApplicationRequest) SetMemory(v int32) *CreateApplicationRequest {
	s.Memory = &v
	return s
}

func (s *CreateApplicationRequest) SetMicroRegistration(v string) *CreateApplicationRequest {
	s.MicroRegistration = &v
	return s
}

func (s *CreateApplicationRequest) SetMicroRegistrationConfig(v string) *CreateApplicationRequest {
	s.MicroRegistrationConfig = &v
	return s
}

func (s *CreateApplicationRequest) SetMicroserviceEngineConfig(v string) *CreateApplicationRequest {
	s.MicroserviceEngineConfig = &v
	return s
}

func (s *CreateApplicationRequest) SetMountDesc(v string) *CreateApplicationRequest {
	s.MountDesc = &v
	return s
}

func (s *CreateApplicationRequest) SetMountHost(v string) *CreateApplicationRequest {
	s.MountHost = &v
	return s
}

func (s *CreateApplicationRequest) SetNamespaceId(v string) *CreateApplicationRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateApplicationRequest) SetNasConfigs(v string) *CreateApplicationRequest {
	s.NasConfigs = &v
	return s
}

func (s *CreateApplicationRequest) SetNasId(v string) *CreateApplicationRequest {
	s.NasId = &v
	return s
}

func (s *CreateApplicationRequest) SetNewSaeVersion(v string) *CreateApplicationRequest {
	s.NewSaeVersion = &v
	return s
}

func (s *CreateApplicationRequest) SetOidcRoleName(v string) *CreateApplicationRequest {
	s.OidcRoleName = &v
	return s
}

func (s *CreateApplicationRequest) SetOssAkId(v string) *CreateApplicationRequest {
	s.OssAkId = &v
	return s
}

func (s *CreateApplicationRequest) SetOssAkSecret(v string) *CreateApplicationRequest {
	s.OssAkSecret = &v
	return s
}

func (s *CreateApplicationRequest) SetOssMountDescs(v string) *CreateApplicationRequest {
	s.OssMountDescs = &v
	return s
}

func (s *CreateApplicationRequest) SetPackageType(v string) *CreateApplicationRequest {
	s.PackageType = &v
	return s
}

func (s *CreateApplicationRequest) SetPackageUrl(v string) *CreateApplicationRequest {
	s.PackageUrl = &v
	return s
}

func (s *CreateApplicationRequest) SetPackageVersion(v string) *CreateApplicationRequest {
	s.PackageVersion = &v
	return s
}

func (s *CreateApplicationRequest) SetPhp(v string) *CreateApplicationRequest {
	s.Php = &v
	return s
}

func (s *CreateApplicationRequest) SetPhpArmsConfigLocation(v string) *CreateApplicationRequest {
	s.PhpArmsConfigLocation = &v
	return s
}

func (s *CreateApplicationRequest) SetPhpConfig(v string) *CreateApplicationRequest {
	s.PhpConfig = &v
	return s
}

func (s *CreateApplicationRequest) SetPhpConfigLocation(v string) *CreateApplicationRequest {
	s.PhpConfigLocation = &v
	return s
}

func (s *CreateApplicationRequest) SetPostStart(v string) *CreateApplicationRequest {
	s.PostStart = &v
	return s
}

func (s *CreateApplicationRequest) SetPreStop(v string) *CreateApplicationRequest {
	s.PreStop = &v
	return s
}

func (s *CreateApplicationRequest) SetProgrammingLanguage(v string) *CreateApplicationRequest {
	s.ProgrammingLanguage = &v
	return s
}

func (s *CreateApplicationRequest) SetPvtzDiscoverySvc(v string) *CreateApplicationRequest {
	s.PvtzDiscoverySvc = &v
	return s
}

func (s *CreateApplicationRequest) SetPython(v string) *CreateApplicationRequest {
	s.Python = &v
	return s
}

func (s *CreateApplicationRequest) SetPythonModules(v string) *CreateApplicationRequest {
	s.PythonModules = &v
	return s
}

func (s *CreateApplicationRequest) SetRaspConfig(v *CreateApplicationRequestRaspConfig) *CreateApplicationRequest {
	s.RaspConfig = v
	return s
}

func (s *CreateApplicationRequest) SetReadiness(v string) *CreateApplicationRequest {
	s.Readiness = &v
	return s
}

func (s *CreateApplicationRequest) SetReplicas(v int32) *CreateApplicationRequest {
	s.Replicas = &v
	return s
}

func (s *CreateApplicationRequest) SetResourceType(v string) *CreateApplicationRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateApplicationRequest) SetSaeVersion(v string) *CreateApplicationRequest {
	s.SaeVersion = &v
	return s
}

func (s *CreateApplicationRequest) SetSecretMountDesc(v string) *CreateApplicationRequest {
	s.SecretMountDesc = &v
	return s
}

func (s *CreateApplicationRequest) SetSecurityGroupId(v string) *CreateApplicationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateApplicationRequest) SetServiceTags(v string) *CreateApplicationRequest {
	s.ServiceTags = &v
	return s
}

func (s *CreateApplicationRequest) SetSidecarContainersConfig(v []*SidecarContainerConfig) *CreateApplicationRequest {
	s.SidecarContainersConfig = v
	return s
}

func (s *CreateApplicationRequest) SetSlsConfigs(v string) *CreateApplicationRequest {
	s.SlsConfigs = &v
	return s
}

func (s *CreateApplicationRequest) SetSlsLogEnvTags(v string) *CreateApplicationRequest {
	s.SlsLogEnvTags = &v
	return s
}

func (s *CreateApplicationRequest) SetStartupProbe(v string) *CreateApplicationRequest {
	s.StartupProbe = &v
	return s
}

func (s *CreateApplicationRequest) SetTerminationGracePeriodSeconds(v int32) *CreateApplicationRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *CreateApplicationRequest) SetTimezone(v string) *CreateApplicationRequest {
	s.Timezone = &v
	return s
}

func (s *CreateApplicationRequest) SetTomcatConfig(v string) *CreateApplicationRequest {
	s.TomcatConfig = &v
	return s
}

func (s *CreateApplicationRequest) SetVSwitchId(v string) *CreateApplicationRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateApplicationRequest) SetVpcId(v string) *CreateApplicationRequest {
	s.VpcId = &v
	return s
}

func (s *CreateApplicationRequest) SetWarStartOptions(v string) *CreateApplicationRequest {
	s.WarStartOptions = &v
	return s
}

func (s *CreateApplicationRequest) SetWebContainer(v string) *CreateApplicationRequest {
	s.WebContainer = &v
	return s
}

func (s *CreateApplicationRequest) Validate() error {
	if s.InitContainersConfig != nil {
		for _, item := range s.InitContainersConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RaspConfig != nil {
		if err := s.RaspConfig.Validate(); err != nil {
			return err
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

type CreateApplicationRequestRaspConfig struct {
	// Specifies whether the application runtime security protection (RASP) is enabled.
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

func (s CreateApplicationRequestRaspConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestRaspConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestRaspConfig) GetEnableRasp() *bool {
	return s.EnableRasp
}

func (s *CreateApplicationRequestRaspConfig) GetRaspAppKey() *string {
	return s.RaspAppKey
}

func (s *CreateApplicationRequestRaspConfig) GetRaspAppName() *string {
	return s.RaspAppName
}

func (s *CreateApplicationRequestRaspConfig) SetEnableRasp(v bool) *CreateApplicationRequestRaspConfig {
	s.EnableRasp = &v
	return s
}

func (s *CreateApplicationRequestRaspConfig) SetRaspAppKey(v string) *CreateApplicationRequestRaspConfig {
	s.RaspAppKey = &v
	return s
}

func (s *CreateApplicationRequestRaspConfig) SetRaspAppName(v string) *CreateApplicationRequestRaspConfig {
	s.RaspAppName = &v
	return s
}

func (s *CreateApplicationRequestRaspConfig) Validate() error {
	return dara.Validate(s)
}

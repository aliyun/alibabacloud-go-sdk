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
	// The ARN of the RAM role required to pull images across Alibaba Cloud accounts. For more information, see [Authorize cross-account access using a RAM role](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The Container Registry Enterprise Edition (ACR Enterprise Edition) instance ID. This parameter is required when **ImageUrl*	- is a Container Registry Enterprise Edition image.
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
	// The application description. It cannot exceed 1024 characters.
	//
	// example:
	//
	// This is a test description.
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The application name. It can contain digits, letters, and hyphens (-). It must start with a letter and cannot end with a hyphen (-). The name cannot exceed 36 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Select micro_service for a microservice application.
	//
	// example:
	//
	// micro_service
	AppSource *string `json:"AppSource,omitempty" xml:"AppSource,omitempty"`
	// Whether to bind an Elastic IP address (EIP). Valid values:
	//
	// - **true**: Bind.
	//
	// - **false**: Do not bind.
	//
	// example:
	//
	// true
	AssociateEip *bool `json:"AssociateEip,omitempty" xml:"AssociateEip,omitempty"`
	// Whether to automatically configure the network environment. Valid values:
	//
	// - **true**: SAE automatically configures the network environment when creating an application. The values of **NamespaceId**, **VpcId**, **vSwitchId**, and **SecurityGroupId*	- are ignored.
	//
	// - **false**: SAE manually configures the network environment when creating an application.
	//
	// > If you select **true**, other **NamespaceId*	- values passed are ignored.
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
	// The image start command. This command must be an executable object that exists in the container. Example:
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
	// Based on the example, Command="echo" and `CommandArgs=["abc", ">", "file0"]`.
	//
	// 	Notice:
	//
	// This option is required when PackageType is DotnetZip.
	//
	// example:
	//
	// echo
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The image start command parameters. These are the parameters required by the **Command*	- parameter. Format:
	//
	// `["a","b"]`
	//
	// In the example, `CommandArgs=["abc", ">", "file0"]`. Convert `["abc", ">", "file0"]` to a string type, with the format as a JSON array. If this parameter is not needed, do not specify it.
	//
	// 	Notice: This option is required when PackageType is DotnetZip.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The **ConfigMap*	- mount description. Use configuration items created on the namespace configuration item page to inject configuration information into the container. Parameter description:
	//
	// - **configMapId**: The ConfigMap instance ID. Obtain it by calling the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) API operation.
	//
	// - **key**: The key value.
	//
	// > You can mount all keys by passing the `sae-sys-configmap-all` parameter.
	//
	// - **mountPath**: The mount path.
	//
	// example:
	//
	// [{"configMapId":16,"key":"test","mountPath":"/tmp"}]
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	// The CPU required for each instance, in millicores. It cannot be 0. Currently, only the following defined specifications are supported:
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
	// Custom Host mapping within the container. Valid values:
	//
	// - **hostName**: The domain name or hostname.
	//
	// - **ip**: The IP address.
	//
	// example:
	//
	// [{"hostName":"samplehost","ip":"127.0.0.1"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The custom image type. If it is not a custom image, set it to an empty string:
	//
	// - internet: Public network image.
	//
	// - intranet: Private network image.
	//
	// example:
	//
	// internet
	CustomImageNetworkType *string `json:"CustomImageNetworkType,omitempty" xml:"CustomImageNetworkType,omitempty"`
	// Whether to deploy immediately. Valid values:
	//
	// - **true**: Default value. Deploy immediately.
	//
	// - **false**: Deploy later.
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
	// The application runtime environment version in the HSF framework, such as the Ali-Tomcat container.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// Shared temporary storage configuration.
	//
	// example:
	//
	// [{\\"name\\":\\"workdir\\",\\"mountPath\\":\\"/usr/local/tomcat/webapps\\"}]
	EmptyDirDesc *string `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty"`
	// Whether to enable the CPU Burst feature:
	//
	// - true: Enable.
	//
	// - false: Do not enable.
	//
	// example:
	//
	// true
	EnableCpuBurst *bool `json:"EnableCpuBurst,omitempty" xml:"EnableCpuBurst,omitempty"`
	// Enable application monitoring for non-Java applications based on eBPF technology. Valid values:
	//
	// - **true**: Enable.
	//
	// - **false**: Disable. Default value.
	//
	// example:
	//
	// false
	EnableEbpf *string `json:"EnableEbpf,omitempty" xml:"EnableEbpf,omitempty"`
	// Whether to reuse the namespace Agent version configuration.
	//
	// example:
	//
	// true
	EnableNamespaceAgentVersion *bool `json:"EnableNamespaceAgentVersion,omitempty" xml:"EnableNamespaceAgentVersion,omitempty"`
	// Whether to reuse the namespace SLS log configuration.
	//
	// example:
	//
	// true
	EnableNamespaceSlsConfig *bool `json:"EnableNamespaceSlsConfig,omitempty" xml:"EnableNamespaceSlsConfig,omitempty"`
	// Whether to enable new ARMS features:
	//
	// - true: Enable.
	//
	// - false: Do not enable.
	//
	// example:
	//
	// false
	EnableNewArms *bool `json:"EnableNewArms,omitempty" xml:"EnableNewArms,omitempty"`
	// Whether to enable Prometheus custom metric collection.
	//
	// example:
	//
	// false
	EnablePrometheus *bool `json:"EnablePrometheus,omitempty" xml:"EnablePrometheus,omitempty"`
	// Whether to enable Sidecar resource isolation:
	//
	// - true: Enable isolation.
	//
	// - false: Do not enable isolation.
	//
	// example:
	//
	// true
	EnableSidecarResourceIsolated *bool `json:"EnableSidecarResourceIsolated,omitempty" xml:"EnableSidecarResourceIsolated,omitempty"`
	// Container environment variable parameters. Support custom configurations or referencing configuration items. To reference a configuration item, create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//   - **name**: The environment variable name.
	//
	//   - **value**: The environment variable value. This has a higher priority than valueFrom.
	//
	// - Reference configuration item (valueFrom)
	//
	//   - **name**: The environment variable name. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<configuration item name>`, for example, `sae-sys-configmap-all-test1`.
	//
	//   - **valueFrom**: The environment variable reference. Set this to `configMapRef`.
	//
	//     - **configMapId**: The configuration item ID.
	//
	//     - **key**: The key. If you reference all key-values, do not set this field.
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
	// K8s Headless Service service discovery.
	//
	// - serviceName: The service name.
	//
	// - namespaceId: The namespace ID.
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
	// The ID of the corresponding secret.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The image address. This parameter is required when **Package Type*	- is **Image**.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/sae_test/ali_sae_test:0.0.1
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Initialization container configuration.
	InitContainersConfigShrink *string `json:"InitContainersConfig,omitempty" xml:"InitContainersConfig,omitempty"`
	// Whether it is a stateful application.
	IsStateful *bool `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// JAR package startup parameters for the application. The application\\"s default start command is: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// custom-args
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// JAR package startup options for the application. The application\\"s default start command is: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// -Xms4G -Xmx4G
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The JDK version that the deployment package depends on. Supported versions:
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
	// The summary configuration for collecting logs to Kafka. Valid values:
	//
	// - **kafkaEndpoint**: The service registration address for the Kafka API.
	//
	// - **kafkaInstanceId**: The Kafka instance ID.
	//
	// - **kafkaConfigs**: The summary configuration for single or multiple logs. For valid values, see the **kafkaConfigs*	- request parameter in this topic.
	//
	// example:
	//
	// {"kafkaEndpoint":"10.0.X.XXX:XXXX,10.0.X.XXX:XXXX,10.0.X.XXX:XXXX","kafkaInstanceId":"alikafka_pre-cn-7pp2l8kr****","kafkaConfigs":[{"logType":"file_log","logDir":"/tmp/a.log","kafkaTopic":"test2"},{"logType":"stdout","logDir":"","kafkaTopic":"test"}]}
	KafkaConfigs *string `json:"KafkaConfigs,omitempty" xml:"KafkaConfigs,omitempty"`
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Container health check. Containers that fail the health check are shut down and recovered. Supported methods:
	//
	// - **exec**: For example, `{"exec":{"command":["sh","-c","cat/home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}`
	//
	// - **httpGet**: For example, `{"httpGet":{"path":"/","port":18091,"scheme":"HTTP","isContainKeyWord":true,"keyWord":"SAE"},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// - **tcpSocket**: For example, `{"tcpSocket":{"port":18091},"initialDelaySeconds":11,"periodSeconds":10,"timeoutSeconds":1}`
	//
	// > Select only one method for the health check.
	//
	// Parameter description:
	//
	// - **exec.command**: Set the health check command.
	//
	// - **httpGet.path**: The access path.
	//
	// - **httpGet.scheme**: **HTTP*	- or **HTTPS**.
	//
	// - **httpGet.isContainKeyWord**: **true*	- means the keyword is included, **false*	- means the keyword is not included. If this field is missing, advanced features are not used.
	//
	// - **httpGet.keyWord**: The custom keyword. Do not omit the **isContainKeyWord*	- field when using it.
	//
	// - **tcpSocket.port**: The port for TCP connection detection.
	//
	// - **initialDelaySeconds**: Set the health check delay detection time. Default is 10 seconds.
	//
	// - **periodSeconds**: Set the health check period. Default is 30 seconds.
	//
	// - **timeoutSeconds**: Set the health check timeout duration. Default is 1 second. If you set it to 0 or do not set it, the default timeout is 1 second.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	Liveness    *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	LokiConfigs *string `json:"LokiConfigs,omitempty" xml:"LokiConfigs,omitempty"`
	// The memory required for each instance, in MB. It cannot be 0. It has a one-to-one correspondence with CPU. Currently, only the following defined specifications are supported:
	//
	// - **1024**: Corresponds to 500 millicores and 1000 millicores CPU.
	//
	// - **2048**: Corresponds to 500, 1000 millicores, and 2000 millicores CPU.
	//
	// - **4096**: Corresponds to 1000, 2000 millicores, and 4000 millicores CPU.
	//
	// - **8192**: Corresponds to 2000, 4000 millicores, and 8000 millicores CPU.
	//
	// - **12288**: Corresponds to 12000 millicores CPU.
	//
	// - **16384**: Corresponds to 4000, 8000 millicores, and 16000 millicores CPU.
	//
	// - **24576**: Corresponds to 12000 millicores CPU.
	//
	// - **32768**: Corresponds to 16000 millicores CPU.
	//
	// - **65536**: Corresponds to 8000, 16000, and 32000 millicores CPU.
	//
	// - **131072**: Corresponds to 32000 millicores CPU.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Select the Nacos registry. Valid values:
	//
	// - **0**: SAE built-in Nacos.
	//
	// - **1**: User-managed Nacos.
	//
	// - **2**: MSE Professional Edition Nacos.
	//
	// example:
	//
	// "0"
	MicroRegistration *string `json:"MicroRegistration,omitempty" xml:"MicroRegistration,omitempty"`
	// The registry configuration information.
	//
	// example:
	//
	// {\\"instanceId\\":\\"mse-cn-zvp2bh6h70r\\",\\"namespace\\":\\"4c0aa74f-57cb-423c-b6af-5d9f2d0e3dbd\\"}
	MicroRegistrationConfig *string `json:"MicroRegistrationConfig,omitempty" xml:"MicroRegistrationConfig,omitempty"`
	// Configure microservice administration features.
	//
	// - Whether to enable microservice administration (enable):
	//
	//   - true: Enable.
	//
	//   - false: Do not enable.
	//
	// - Configure graceful start and graceful shutdown (mseLosslessRule):
	//
	//   - delayTime: The delay time.
	//
	//   - enable: Whether to enable the graceful start feature. true means enabled, false means not enabled.
	//
	//   - notice: Whether to enable the notification feature. true means enabled, false means enabled.
	//
	//   - warmupTime: The duration of traffic prefetch, in seconds.
	//
	// example:
	//
	// {"enable": true,"mseLosslessRule": {"delayTime": 0,"enable": false,"notice": false,"warmupTime": 120}}
	MicroserviceEngineConfig *string `json:"MicroserviceEngineConfig,omitempty" xml:"MicroserviceEngineConfig,omitempty"`
	// Do not configure this field; configure **NasConfigs*	- instead. The NAS mount description. If the configuration has not changed during deployment, you do not need to set this parameter (that is, the request does not need to include the **MountDesc*	- field). To clear the NAS configuration, set the value of this field to an empty string in the request (that is, the value of the **MountDesc*	- field in the request is "").
	//
	// example:
	//
	// [{mountPath: "/tmp", nasPath: "/"}]
	MountDesc *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	// Do not configure this field; configure **NasConfigs*	- instead. The NAS mount target within the application VPC. If the configuration has not changed during deployment, you do not need to set this parameter (that is, the request does not need to include the **MountHost*	- field). To clear the NAS configuration, set the value of this field to an empty string in the request (that is, the value of the **MountHost*	- field in the request is "").
	//
	// example:
	//
	// example.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The SAE namespace ID. Only namespaces with names consisting of lowercase letters and hyphens (-) are supported. The name must start with a letter. Obtain the namespace by calling the [DescribeNamespaceList](https://help.aliyun.com/document_detail/126547.html) API operation.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The configuration for mounting NAS. Valid values:
	//
	// - **mountPath**: The container mount path.
	//
	// - **readOnly**: If the value is **false**, it indicates read and write permission.
	//
	// - **nasId**: The NAS ID.
	//
	// - **mountDomain**: The container mount target address. For more information, see [DescribeMountTargets](https://help.aliyun.com/document_detail/62626.html).
	//
	// - **nasPath**: The relative file directory of NAS.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// Do not configure this field; configure **NasConfigs*	- instead. The ID of the mounted NAS. It must be in the same region as the cluster. It must have available mount target creation quotas, or its mount target must already be on a vSwitch within the VPC. If you do not specify this parameter and the **mountDescs*	- field exists, the system automatically purchases a NAS and mounts it to a vSwitch within the VPC by default.
	//
	// If the configuration has not changed during deployment, you do not need to set this parameter (that is, the request does not need to include the **NASId*	- field). To clear the NAS configuration, set the value of this field to an empty string in the request (that is, the value of the **NASId*	- field in the request is "").
	//
	// example:
	//
	// KSAK****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The application version:
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
	// Set the identity authentication service RAM role.
	//
	// > Create an OpenID Connect (OIDC) identity provider and an identity provider role in the same region beforehand. For more information, see<props="china">[Create an OIDC identity provider](https://help.aliyun.com/zh/ram/developer-reference/api-ims-2019-08-15-createoidcprovider?spm=a2c4g.11186623.help-menu-28625.d_4_1_0_3_2_7.7f0443efmdpxa3) and[Create a role SSO identity provider](https://help.aliyun.com/zh/ram/developer-reference/api-ims-2019-08-15-createsamlprovider?spm=a2c4g.11186623.help-menu-28625.d_4_1_0_3_2_2.632244b1s8QbQt)<props="intl">[Create an OIDC identity provider](https://www.alibabacloud.com/help/zh/ram/developer-reference/api-ims-2019-08-15-createoidcprovider) and[Create a role SSO identity provider](https://www.alibabacloud.com/help/zh/ram/developer-reference/api-ims-2019-08-15-createsamlprovider).
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
	// OSS mount description. Parameter description:
	//
	// - **bucketName**: The Bucket name.
	//
	// - **bucketPath**: The directory or OSS object you created in OSS. If the OSS mount directory does not exist, an exception is triggered.
	//
	// - **mountPath**: The container path in SAE. If the path exists, it is overwritten. If the path does not exist, it is created.
	//
	// - **readOnly**: Whether the container path has read permission for the mounted directory resource. Valid values:
	//
	//   - **true**: Read-only permission.
	//
	//   - **false**: Read and write permission.
	//
	// example:
	//
	// [{"bucketName": "oss-bucket", "bucketPath": "data/user.data", "mountPath": "/usr/data/user.data", "readOnly": true}]
	OssMountDescs *string `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty"`
	// The application package type. Valid values:
	//
	// - If you deploy with Java, supported types are **FatJar**, **War**, and **Image**.
	//
	// - If you deploy with PHP, supported types are:
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
	// - If you deploy with Python, supported types are **PythonZip*	- and **Image**.
	//
	// - If you deploy with .NET Core, supported types are **DotnetZip*	- and **Image**.
	//
	//   > When you select DotnetZip, Dotnet is the version number of the .NET Core environment. Supported versions are .NET 3.1, .NET 5.0, .NET 6.0, .NET 7.0, and .NET 8.0. The Dotnet, Command, and CommandArgs options are required.
	//
	// This parameter is required.
	//
	// example:
	//
	// FatJar
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The URL of the deployment package. This parameter is required when **Package Type*	- is **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// http://myoss.oss-cn-****.aliyuncs.com/my-buc/2019-06-30/****.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version number of the deployment package. This parameter is required when **Package Type*	- is **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// 1.0.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The PHP version that the PHP deployment package depends on. Images do not support this.
	//
	// example:
	//
	// PHP-FPM 7.0
	Php *string `json:"Php,omitempty" xml:"Php,omitempty"`
	// The mount path for PHP application monitoring. Ensure that the PHP server loads the configuration file from this path. You do not need to focus on the configuration content; SAE automatically renders the correct configuration file.
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
	// The mount path for PHP application startup configuration. Ensure that the PHP server uses this configuration file to start.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// The script to execute after the container starts. A script is triggered immediately after the container is created. Format: `{"exec":{"command":["cat","/etc/group"]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script to execute before the container stops. A script is triggered before the container is deleted. Format: `{"exec":{"command":["cat","/etc/group"]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The technology stack language for creating the application. Valid values:
	//
	// - **java**: Java language.
	//
	// - **php**: PHP language.
	//
	// - **python**: Python language.
	//
	// - **dotnet**: .NET Core language.
	//
	// - **other**: Multi-language, such as C++, Go, and Node.js.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// Enable K8s Service service discovery. Valid values:
	//
	// - **serviceName**: The service name. Format: `custom-namespace ID`. The suffix `-namespace ID` cannot be customized; specify it based on the application\\"s namespace. For example, if you select the default namespace in the China (Beijing) region, it is `-cn-beijing-default`.
	//
	// - **namespaceId**: The namespace ID.
	//
	// - **portProtocols**: The port and protocol. The port range is [1, 65535]. Supported protocols are **TCP*	- and **UDP**.
	//
	// - portAndProtocol: The port and protocol. The port range is [1, 65535]. Supported protocols are TCP and **UDP**. **portProtocols*	- is recommended. If **portProtocols*	- is set, only **portProtocols*	- takes effect.
	//
	// - **enable**: Enable K8s Service service discovery.
	//
	// example:
	//
	// {"serviceName":"bwm-poc-sc-gateway-cn-beijing-front","namespaceId":"cn-beijing:front","portAndProtocol":{"18012":"TCP"},"enable":true,"portProtocols":[{"port":18012,"protocol":"TCP"}]}
	PvtzDiscoverySvc *string `json:"PvtzDiscoverySvc,omitempty" xml:"PvtzDiscoverySvc,omitempty"`
	// The Python environment. Supports **PYTHON 3.9.15**.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// Custom installation of module dependencies. By default, the system installs dependencies defined in requirements.txt in the root directory. If you do not configure or customize packages, you can specify the dependencies to install.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// Application startup status check. Containers that fail multiple health checks are shut down and restarted. Containers that do not pass the health check will not receive SLB traffic. Supported methods are **exec**, **httpGet**, and **tcpSocket**. For examples, see the **Liveness*	- parameter.
	//
	// > Select only one method for the health check.
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
	// The resource type. Supports NULL (default), default, and haiguang (Haiguang server) types.
	//
	// example:
	//
	// NULL
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The SAE version. Supported versions:
	//
	// - **v1**
	//
	// - **v2**
	//
	// example:
	//
	// v1
	SaeVersion *string `json:"SaeVersion,omitempty" xml:"SaeVersion,omitempty"`
	// The **Secret*	- mount description. Use secrets created on the namespace secret page to inject secret information into the container. Parameter description:
	//
	// - **secretId**: The secret instance ID. Obtain it by calling the ListSecrets API operation.
	//
	// - **key**: The key value.
	//
	// > You can mount all keys by passing the `sae-sys-secret-all` parameter.
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
	// The grayscale tags for application configuration.
	//
	// example:
	//
	// {\\"alicloud.service.tag\\":\\"g1\\"}
	ServiceTags *string `json:"ServiceTags,omitempty" xml:"ServiceTags,omitempty"`
	// Container configuration information.
	SidecarContainersConfigShrink *string `json:"SidecarContainersConfig,omitempty" xml:"SidecarContainersConfig,omitempty"`
	// The configuration for collecting logs to Simple Log Service (SLS).
	//
	// - Use SLS resources automatically created by SAE: `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - Use custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// Parameter description:
	//
	// - **projectName**: The name of the Project on SLS.
	//
	// - **logDir**: The log path.
	//
	// - **logType**: The log type. **stdout*	- indicates container standard output logs; you can set only one such entry. If you do not set this, the system collects file logs.
	//
	// - **logstoreName**: The name of the Logstore on SLS.
	//
	// - **logtailName**: The name of the Logtail on SLS. If you do not specify this, the system creates a new Logtail.
	//
	// If the SLS collection configuration has not changed during multiple deployments, you do not need to set this parameter (that is, the request does not need to include the **SlsConfigs*	- field). If you no longer need the SLS collection feature, set the value of this field to an empty string in the request (that is, the value of the **SlsConfigs*	- field in the request is "").
	//
	// > Projects automatically created with an application are deleted when the application is deleted. Therefore, when selecting an existing Project, do not select a Project automatically created by SAE.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// SLS log tags.
	SlsLogEnvTags *string `json:"SlsLogEnvTags,omitempty" xml:"SlsLogEnvTags,omitempty"`
	// Enable application startup probes.
	//
	// - Successful check: Indicates that the application started successfully. If you configured Liveness and Readiness checks, the system performs Liveness and Readiness checks after the application starts successfully.
	//
	// - Failed check: Indicates that the application failed to start. The system reports an exception and automatically restarts the application.
	//
	// > 	- Supported methods are exec, httpGet, and tcpSocket. For examples, see the Liveness parameter.
	//
	// >
	//
	// > 	- Select only one method for the health check.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	StartupProbe *string `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	// The graceful shutdown timeout duration. Default is 30 seconds. Valid values are 1 to 300.
	//
	// example:
	//
	// 30
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The time zone. Default is **Asia/Shanghai**.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// Tomcat file configuration. Set to "" or "{}" to delete the configuration:
	//
	// - **port**: The port range is 1024 to 65535. Ports less than 1024 require root permissions to operate. Because the container is configured with Admin permissions, specify a port greater than 1024. If you do not configure this, the default is 8080.
	//
	// - **contextPath**: The access path. Default is the root directory "/".
	//
	// - **maxThreads**: Configure the connection pool size. Default is 400.
	//
	// - uriEncoding: The encoding format for Tomcat, including **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. If you do not set this, the default is **ISO-8859-1**.
	//
	// - **useBodyEncodingForUri**: Whether to use **BodyEncoding for URL**. Default is **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	// The virtual switch (vSwitch) where the application instance\\"s Elastic Network Interface (ENI) is located. This vSwitch must be within the specified VPC. This vSwitch also has a binding relationship with the SAE namespace. If you do not specify this parameter, the system uses the vSwitch ID bound to the namespace by default.
	//
	// example:
	//
	// vsw-bp12mw1f8k3jgygk9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC corresponding to the SAE namespace. In SAE, a namespace can only correspond to one VPC, and you cannot change it. The first time you create an SAE application in a namespace, a binding relationship forms. Multiple namespaces can correspond to one VPC. If you do not specify this parameter, the system uses the VPC ID bound to the namespace by default.
	//
	// example:
	//
	// vpc-bp1aevy8sofi8mh1q****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Set the startup command for WAR package deployed applications. The procedure is the same as configuring the startup command for image deployments. For more information, see [Set the startup command](https://help.aliyun.com/document_detail/96677.html).
	//
	// example:
	//
	// CATALINA_OPTS=\\"$CATALINA_OPTS $Options\\" catalina.sh run
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The Tomcat version that the WebContainer deployment package depends on. Supported versions:
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

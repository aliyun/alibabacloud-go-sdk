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
	SetLiveness(v string) *CreateApplicationRequest
	GetLiveness() *string
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
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	AgentVersion  *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The description of the template. The description cannot exceed 1,024 characters in length.
	//
	// example:
	//
	// This is a test description.
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The name of the application. The name can contain digits, letters, and hyphens (-). The name must start with a letter and cannot end with a hyphen (-). It cannot exceed 36 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Select micro_service, which is the application.
	//
	// example:
	//
	// micro_service
	AppSource *string `json:"AppSource,omitempty" xml:"AppSource,omitempty"`
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
	// Specifies whether to automatically configure the network environment. Valid values:
	//
	// 	- **true**: SAE automatically configures the network environment when you create the application. If you set this parameter to true, the values of the **NamespaceId**, **VpcId**, **vSwitchId**, and **SecurityGroupId*	- parameters are ignored.
	//
	// 	- **false**: SAE configures the network environment based on your settings when you create the application.
	//
	// >  If you select **true**, other **NamespaceId*	- will be ignored.
	//
	// example:
	//
	// true
	AutoConfig *bool `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty"`
	// The ID of the basic application.
	//
	// example:
	//
	// ee99cce6-1c8e-4bfa-96c3-3e2fa9de8a41
	BaseAppId *string `json:"BaseAppId,omitempty" xml:"BaseAppId,omitempty"`
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
	// Whether to deploy now.
	//
	// 	- **true*	- (default): Deploy now.
	//
	// 	- **false**: Deploy later.
	//
	// example:
	//
	// true
	Deploy *bool `json:"Deploy,omitempty" xml:"Deploy,omitempty"`
	// The disk size. Unit: GB.
	//
	// example:
	//
	// 50
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// . NET Framework version number:
	//
	// 	- .NET 3.1
	//
	// 	- .NET 5.0
	//
	// 	- .NET 6.0
	//
	// 	- .NET 7.0
	//
	// 	- .NET 8.0
	//
	// example:
	//
	// .NET 3.1
	Dotnet *string `json:"Dotnet,omitempty" xml:"Dotnet,omitempty"`
	// The version of the container in HSF.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	EmptyDirDesc         *string `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty"`
	// Enable CPU Burst.
	//
	// - true: enable
	//
	// - false: disable
	//
	// example:
	//
	// true
	EnableCpuBurst *bool `json:"EnableCpuBurst,omitempty" xml:"EnableCpuBurst,omitempty"`
	// Enable application monitoring for non-Java applications based on eBPF technology. The value options are as follows:
	//
	// - true: Enable.
	//
	// - false: Disable (default).
	//
	// example:
	//
	// false
	EnableEbpf                  *string `json:"EnableEbpf,omitempty" xml:"EnableEbpf,omitempty"`
	EnableNamespaceAgentVersion *bool   `json:"EnableNamespaceAgentVersion,omitempty" xml:"EnableNamespaceAgentVersion,omitempty"`
	EnableNamespaceSlsConfig    *bool   `json:"EnableNamespaceSlsConfig,omitempty" xml:"EnableNamespaceSlsConfig,omitempty"`
	// Indicates whether to enable the new ARMS feature:
	//
	// 	- true: enables this parameter.
	//
	// 	- false: disables this parameter.
	//
	// example:
	//
	// false
	EnableNewArms    *bool `json:"EnableNewArms,omitempty" xml:"EnableNewArms,omitempty"`
	EnablePrometheus *bool `json:"EnablePrometheus,omitempty" xml:"EnablePrometheus,omitempty"`
	// Enable Sidecar resource isolation.
	//
	// - true: enable
	//
	// - false: disable
	//
	// example:
	//
	// true
	EnableSidecarResourceIsolated *bool `json:"EnableSidecarResourceIsolated,omitempty" xml:"EnableSidecarResourceIsolated,omitempty"`
	// The environment variables. You can configure custom environment variables or reference a ConfigMap. Before you can reference a ConfigMap, you must create a ConfigMap. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// 	- Custom configuration
	//
	//     	- **name**: the name of the environment variable.
	//
	//     	- **value**: the value of the environment variable. The priority of the custom configuration is higher than valueFrom.
	//
	// 	- Reference a ConfigMap (valueFrom)
	//
	//     	- **name**: the name of the environment variable. You can reference one or all keys. To reference all keys, specify `sae-sys-configmap-all-<ConfigMap name>`. Example: `sae-sys-configmap-all-test1`.
	//
	//     	- **valueFrom**: the reference of the environment variable. Valid value: `configMapRef`.
	//
	//     	- **configMapId**: the ID of the ConfigMap.
	//
	//     	- **key**: the key. If you want to reference all key values, you do not need to configure this parameter.
	//
	// example:
	//
	// [{"name":"envtmp","value":"0"}]
	Envs                     *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	GpuConfig                *string `json:"GpuConfig,omitempty" xml:"GpuConfig,omitempty"`
	HeadlessPvtzDiscoverySvc *string `json:"HeadlessPvtzDiscoverySvc,omitempty" xml:"HeadlessPvtzDiscoverySvc,omitempty"`
	Html                     *string `json:"Html,omitempty" xml:"Html,omitempty"`
	// The ID of the corresponding Secret.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The URL of the image. This parameter is required if you set the `PackageType` parameter to `Image`.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/sae_test/ali_sae_test:0.0.1
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Initialize container configuration.
	InitContainersConfig []*InitContainerConfig `json:"InitContainersConfig,omitempty" xml:"InitContainersConfig,omitempty" type:"Repeated"`
	IsStateful           *bool                  `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The arguments in the JAR package. The arguments are used to start the application container. The default startup command is `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`.
	//
	// example:
	//
	// custom-args
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The option settings in the JAR package. The settings are used to start the application container. The default startup command for application deployment is `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`.
	//
	// example:
	//
	// -Xms4G -Xmx4G
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
	// Container health check. If the container fails this check, it will be revoked and relaunch again. Use one of the following methods to perform the health check:
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
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
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
	// 	- **2*	- : MSE enterprise edition Nacos registry
	//
	// example:
	//
	// "0"
	MicroRegistration *string `json:"MicroRegistration,omitempty" xml:"MicroRegistration,omitempty"`
	// The Registry configurations.
	//
	// example:
	//
	// {\\"instanceId\\":\\"mse-cn-zvp2bh6h70r\\",\\"namespace\\":\\"4c0aa74f-57cb-423c-b6af-5d9f2d0e3dbd\\"}
	MicroRegistrationConfig *string `json:"MicroRegistrationConfig,omitempty" xml:"MicroRegistrationConfig,omitempty"`
	// Configure microservices governance
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
	// It is not recommended to configure this field; configuring NasConfigs instead. This field specifies the NAS mount description. When deploying, if the configuration has not changed, you do not need to set this parameter (i.e., the MountDesc field does not need to be included in the request). If you need to clear the NAS configuration, set the value of this field to an empty string in the request (i.e., set the value of the MountDesc field to "").
	//
	// example:
	//
	// [{mountPath: "/tmp", nasPath: "/"}]
	MountDesc *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	// It is not recommended to configure this field; configuring NasConfigs instead. This field specifies the NAS mount point within the application\\"s VPC. When deploying, if the configuration has not changed, you do not need to set this parameter (i.e., the MountHost field does not need to be included in the request). If you need to clear the NAS configuration, set the value of this field to an empty string in the request (i.e., set the value of the MountHost field to "").
	//
	// example:
	//
	// example.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// SAE namespace ID. Only namespaces consisting of lowercase letters and hyphens (-) are supported, and the name must start with a letter.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
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
	// It is not recommended to configure this field; configuring NasConfigs instead. The ID of the mounted NAS must be in the same region as the cluster. The NAS must have available mount point quota or its mount point must already be on a switch within the VPC. If this field is not specified and the mountDescs field exists, a NAS will be automatically purchased and mounted to a switch within the VPC by default.
	//
	// When deploying, if the configuration has not changed, you do not need to set this parameter (i.e., the NASId field does not need to be included in the request). If you need to clear the NAS configuration, set the value of this field to an empty string in the request (i.e., set the value of the NASId field to "").
	//
	// example:
	//
	// KSAK****
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
	// The Accesskey ID that the OSS reads and writes from.
	//
	// example:
	//
	// xxxxxx
	OssAkId *string `json:"OssAkId,omitempty" xml:"OssAkId,omitempty"`
	// The AccessKey Secret that the OSS reads and writes from.
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
	// The type of the deployment package. Take note of the following rules:
	//
	// 	- If you deploy the application by using a Java Archive (JAR) package, you can set this parameter to **FatJar**, **War**, or **Image**.
	//
	// 	- If you deploy the application by using a PHP package, you can set this parameter to one of the following values:
	//
	// **PhpZip*	- **IMAGE_PHP_5_4*	- **IMAGE_PHP_5_4_ALPINE*	- **IMAGE_PHP_5_5*	- **IMAGE_PHP_5_5_ALPINE*	- **IMAGE_PHP_5_6*	- **IMAGE_PHP_5_6_ALPINE*	- **IMAGE_PHP_7_0*	- **IMAGE_PHP_7_0_ALPINE*	- **IMAGE_PHP_7_1*	- **IMAGE_PHP_7_1_ALPINE*	- **IMAGE_PHP_7_2*	- **IMAGE_PHP_7_2_ALPINE*	- **IMAGE_PHP_7_3*	- **IMAGE_PHP_7_3_ALPINE**
	//
	// 	- If you deploy the application by using a **Python*	- package, you can set this parameter to **PythonZip*	- or **Image**:
	//
	// This parameter is required.
	//
	// example:
	//
	// FatJar
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The address of the deployment package. This parameter is required if you set **PackageType*	- to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// http://myoss.oss-cn-****.aliyuncs.com/my-buc/2019-06-30/****.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package. This parameter is required when the **PackageType*	- parameter is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// 1.0.0
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
	// Control whether to run a script after the container is initialized. Example: {"exec":{"command":["cat","/etc/group"]}}
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// To controle whether to run a script before the container stops. Example: {"exec":{"command":["cat","/etc/group"]}}
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The programming language for the application’s technology stack. The value options are as follows:
	//
	// - java: Java language
	//
	// - php: PHP language
	//
	// - python: Python language
	//
	// - dotnet: .NET Core language
	//
	// - other: Multi-language, such as C++, Go, Node.js, etc.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
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
	// Check the launch status of the container. Containers that fail health checks more than once will not receive traffic from Server Load Balancer (SLB) instances any loner. You can use the **exec**, **httpGet**, or **tcpSocket*	- method to perform health checks. For more information, see the description of the **Liveness*	- parameter.
	//
	// > You can use only one method to perform the health check.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The number of instances when initialized.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The resource type. Supports NULL (default) and haiguang (haiguang server).
	//
	// example:
	//
	// UNLL
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The SAE version. Supported versions:
	//
	// 	- **v1**
	//
	// 	- **v2**
	//
	// example:
	//
	// v1
	SaeVersion *string `json:"SaeVersion,omitempty" xml:"SaeVersion,omitempty"`
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
	// The canary tag configured for the application.
	//
	// example:
	//
	// {\\"alicloud.service.tag\\":\\"g1\\"}
	ServiceTags *string `json:"ServiceTags,omitempty" xml:"ServiceTags,omitempty"`
	// The configuration of the container.
	SidecarContainersConfig []*SidecarContainerConfig `json:"SidecarContainersConfig,omitempty" xml:"SidecarContainersConfig,omitempty" type:"Repeated"`
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
	SlsConfigs    *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	SlsLogEnvTags *string `json:"SlsLogEnvTags,omitempty" xml:"SlsLogEnvTags,omitempty"`
	// Enable application startup probe.
	//
	// Check succeeded: Indicates that the application has started successfully. If you have configured Liveness and Readiness checks, they will be performed after the application startup is successful.
	//
	// Check failed: Indicates that the application failed to start; an exception will be reported and the application will be automatically restarted.
	//
	// > - exec, httpGet, and tcpSocket methods are supported. For specific examples, see the Liveness parameter documentation.
	//
	// > - Only one health check method can be selected.
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds":2}
	StartupProbe *string `json:"StartupProbe,omitempty" xml:"StartupProbe,omitempty"`
	// The timeout period for a graceful shutdown. Default value: 30. Unit: seconds. Valid values: 1 to 300.
	//
	// example:
	//
	// 30
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// Time zone. Default to time zone of Asia/Shanghai.
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
	// The vSwitch to which the elastic network interface (ENI) of the application instance is connected. The vSwitch must be located in the VPC specified by the VpcId parameter. The SAE namespace is bound with this vSwitch. The default value is the ID of the vSwitch that is bound to the namespace.
	//
	// example:
	//
	// vsw-bp12mw1f8k3jgygk9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) that corresponds to the SAE namespace. In SAE, once correspondence is configured between a namespace and a VPC, the namespace cannot correspond to other VPCs. When the SAE application is created within the namespace, the application is bound with the VPC. Multiple namespaces can correspond to the same VPC. The default value is the ID of the VPC that is bound to the namespace.
	//
	// example:
	//
	// vpc-bp1aevy8sofi8mh1q****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *CreateApplicationRequest) GetLiveness() *string {
	return s.Liveness
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

func (s *CreateApplicationRequest) SetLiveness(v string) *CreateApplicationRequest {
	s.Liveness = &v
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrAssumeRoleArn(v string) *CreateJobRequest
	GetAcrAssumeRoleArn() *string
	SetAcrInstanceId(v string) *CreateJobRequest
	GetAcrInstanceId() *string
	SetAppDescription(v string) *CreateJobRequest
	GetAppDescription() *string
	SetAppName(v string) *CreateJobRequest
	GetAppName() *string
	SetAutoConfig(v bool) *CreateJobRequest
	GetAutoConfig() *bool
	SetBackoffLimit(v int64) *CreateJobRequest
	GetBackoffLimit() *int64
	SetBestEffortType(v string) *CreateJobRequest
	GetBestEffortType() *string
	SetCommand(v string) *CreateJobRequest
	GetCommand() *string
	SetCommandArgs(v string) *CreateJobRequest
	GetCommandArgs() *string
	SetConcurrencyPolicy(v string) *CreateJobRequest
	GetConcurrencyPolicy() *string
	SetConfigMapMountDesc(v string) *CreateJobRequest
	GetConfigMapMountDesc() *string
	SetCpu(v int32) *CreateJobRequest
	GetCpu() *int32
	SetCustomHostAlias(v string) *CreateJobRequest
	GetCustomHostAlias() *string
	SetEdasContainerVersion(v string) *CreateJobRequest
	GetEdasContainerVersion() *string
	SetEnableImageAccl(v bool) *CreateJobRequest
	GetEnableImageAccl() *bool
	SetEnvs(v string) *CreateJobRequest
	GetEnvs() *string
	SetImagePullSecrets(v string) *CreateJobRequest
	GetImagePullSecrets() *string
	SetImageUrl(v string) *CreateJobRequest
	GetImageUrl() *string
	SetJarStartArgs(v string) *CreateJobRequest
	GetJarStartArgs() *string
	SetJarStartOptions(v string) *CreateJobRequest
	GetJarStartOptions() *string
	SetJdk(v string) *CreateJobRequest
	GetJdk() *string
	SetMemory(v int32) *CreateJobRequest
	GetMemory() *int32
	SetMountDesc(v string) *CreateJobRequest
	GetMountDesc() *string
	SetMountHost(v string) *CreateJobRequest
	GetMountHost() *string
	SetNamespaceId(v string) *CreateJobRequest
	GetNamespaceId() *string
	SetNasConfigs(v string) *CreateJobRequest
	GetNasConfigs() *string
	SetNasId(v string) *CreateJobRequest
	GetNasId() *string
	SetOssAkId(v string) *CreateJobRequest
	GetOssAkId() *string
	SetOssAkSecret(v string) *CreateJobRequest
	GetOssAkSecret() *string
	SetOssMountDescs(v string) *CreateJobRequest
	GetOssMountDescs() *string
	SetPackageType(v string) *CreateJobRequest
	GetPackageType() *string
	SetPackageUrl(v string) *CreateJobRequest
	GetPackageUrl() *string
	SetPackageVersion(v string) *CreateJobRequest
	GetPackageVersion() *string
	SetPhpConfig(v string) *CreateJobRequest
	GetPhpConfig() *string
	SetPhpConfigLocation(v string) *CreateJobRequest
	GetPhpConfigLocation() *string
	SetPostStart(v string) *CreateJobRequest
	GetPostStart() *string
	SetPreStop(v string) *CreateJobRequest
	GetPreStop() *string
	SetProgrammingLanguage(v string) *CreateJobRequest
	GetProgrammingLanguage() *string
	SetPython(v string) *CreateJobRequest
	GetPython() *string
	SetPythonModules(v string) *CreateJobRequest
	GetPythonModules() *string
	SetRefAppId(v string) *CreateJobRequest
	GetRefAppId() *string
	SetReplicas(v int32) *CreateJobRequest
	GetReplicas() *int32
	SetSecurityGroupId(v string) *CreateJobRequest
	GetSecurityGroupId() *string
	SetSlice(v bool) *CreateJobRequest
	GetSlice() *bool
	SetSliceEnvs(v string) *CreateJobRequest
	GetSliceEnvs() *string
	SetSlsConfigs(v string) *CreateJobRequest
	GetSlsConfigs() *string
	SetTerminationGracePeriodSeconds(v int32) *CreateJobRequest
	GetTerminationGracePeriodSeconds() *int32
	SetTimeout(v int64) *CreateJobRequest
	GetTimeout() *int64
	SetTimezone(v string) *CreateJobRequest
	GetTimezone() *string
	SetTomcatConfig(v string) *CreateJobRequest
	GetTomcatConfig() *string
	SetTriggerConfig(v string) *CreateJobRequest
	GetTriggerConfig() *string
	SetVSwitchId(v string) *CreateJobRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateJobRequest
	GetVpcId() *string
	SetWarStartOptions(v string) *CreateJobRequest
	GetWarStartOptions() *string
	SetWebContainer(v string) *CreateJobRequest
	GetWebContainer() *string
	SetWorkload(v string) *CreateJobRequest
	GetWorkload() *string
}

type CreateJobRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is required to pull images across accounts. For more information, see [Grant permissions to pull images across Alibaba Cloud accounts by using a RAM role](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The ID of the Container Registry (ACR) Enterprise Edition instance. This parameter is required when **ImageUrl*	- points to an image in an ACR Enterprise Edition instance.
	//
	// example:
	//
	// cri-xxxxxx
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The description of the job template. It cannot exceed 1,024 characters.
	//
	// example:
	//
	// This is a test description.
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The name of the job template. The name can contain letters, digits, and hyphens (-). It must start with a letter and be no longer than 36 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether to automatically configure the network environment. Valid values:
	//
	// - **true**: SAE automatically configures the network environment when you create the job template. The values of **NamespaceId**, **VpcId**, **vSwitchId**, and **SecurityGroupId*	- are ignored.
	//
	// - **false**: You must manually configure the network environment.
	//
	// example:
	//
	// false
	AutoConfig *bool `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty"`
	// The maximum number of retries for a task before it is marked as failed.
	//
	// example:
	//
	// 3
	BackoffLimit *int64 `json:"BackoffLimit,omitempty" xml:"BackoffLimit,omitempty"`
	// The BestEffort policy.
	BestEffortType *string `json:"BestEffortType,omitempty" xml:"BestEffortType,omitempty"`
	// The entrypoint command for the container. The command must be an executable inside the container. Example:
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
	// For the preceding example, `Command="echo", CommandArgs=["abc", ">", "file0"]`.
	//
	// example:
	//
	// echo
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// Arguments for the entrypoint command (**Command**). The format is as follows:
	//
	// `["a","b"]`
	//
	// In the example for the `Command` parameter, the value for `CommandArgs` is `["abc", ">", "file0"]`. This value must be a string that contains a JSON array. If the command takes no arguments, you can omit this parameter.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The concurrency policy for the task. Valid values:
	//
	// - **Forbid**: Prohibits concurrent runs. A new task is not created if the previous one is not complete.
	//
	// - **Allow**: Allows concurrent running.
	//
	// - **Replace**: If a previous task is still running when the next one is scheduled, the new task replaces the old one.
	//
	// example:
	//
	// Allow
	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" xml:"ConcurrencyPolicy,omitempty"`
	// The **ConfigMap*	- mount description. Use a ConfigMap created in the namespace to inject configurations into the container. The parameters are described as follows:
	//
	// - **configMapId**: The ID of the ConfigMap. You can call the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain this ID.
	//
	// - **key**: The key.
	//
	// > You can pass the `sae-sys-configmap-all` parameter to mount all keys.
	//
	// - **mountPath**: The mount path in the container.
	//
	// example:
	//
	// [{"configMapId":16,"key":"test","mountPath":"/tmp"}]
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	// The CPU required for each instance, in millicores. This value cannot be 0. Only the following fixed specifications are currently supported:
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
	// The host alias that maps a hostname to an IP address in the container. The parameters are described as follows:
	//
	// - **hostName**: The domain name or hostname.
	//
	// - **ip**: The IP address.
	//
	// example:
	//
	// [{"hostName":"samplehost","ip":"127.0.0.1"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The version of the HSF runtime environment for the task, such as an Ali-Tomcat container.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// Specifies whether to enable image acceleration. Valid values:
	//
	// - **true**: Enables image acceleration.
	//
	// - **false**: Disables image acceleration.
	//
	// example:
	//
	// false
	EnableImageAccl *bool `json:"EnableImageAccl,omitempty" xml:"EnableImageAccl,omitempty"`
	// Environment variables to set in the container. To reference variables, the ConfigMap must already exist. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). The value can be configured in one of the following ways:
	//
	// - Specify custom variables
	//
	//   - **name**: The name of the environment variable.
	//
	//   - **value**: The value of the environment variable.
	//
	// - Reference a ConfigMap
	//
	//   - **name**: The name of the environment variable. You can reference a single key or all keys. To reference all keys, enter a value in the `sae-sys-configmap-all-<ConfigMap name>` format. Example: `sae-sys-configmap-all-test1`.
	//
	//   - **valueFrom**: The source of the environment variable. Set the value to `configMapRef`.
	//
	//   - **configMapId**: The ID of the ConfigMap.
	//
	//   - **key**: The key to reference. If you want to reference all key-value pairs, do not specify this parameter.
	//
	// example:
	//
	// [{"name":"envtmp","value":"0"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The ID of the secret used to pull the image.
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
	// The startup arguments for the JAR package. The default startup command is: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// -Xms4G -Xmx4G
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The startup options for the JAR package. The default startup command is: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// custom-option
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The JDK version that the deployment package requires. The following versions are supported:
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
	// The memory required for each instance, in MB. This value cannot be 0. CPU and memory specifications are coupled. The following specifications are currently supported:
	//
	// - **1024**: corresponds to 500 or 1,000 millicores of CPU.
	//
	// - **2048**: corresponds to 500, 1,000, or 2,000 millicores of CPU.
	//
	// - **4096**: corresponds to 1,000, 2,000, or 4,000 millicores of CPU.
	//
	// - **8192**: corresponds to 2,000, 4,000, or 8,000 millicores of CPU.
	//
	// - **12288**: corresponds to 12,000 millicores of CPU.
	//
	// - **16384**: corresponds to 4,000, 8,000, or 16,000 millicores of CPU.
	//
	// - **24576**: corresponds to 12,000 millicores of CPU.
	//
	// - **32768**: corresponds to 16,000 millicores of CPU.
	//
	// - **65536**: corresponds to 8,000, 16,000, or 32,000 millicores of CPU.
	//
	// - **131072**: corresponds to 32,000 millicores of CPU.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The NAS mount description. If this configuration does not change in subsequent deployments, you can omit this parameter. To clear the NAS configuration, set this parameter to an empty string ("").
	//
	// example:
	//
	// [{mountPath: "/tmp", nasPath: "/"}]
	MountDesc *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	// The NAS mount target in the VPC of the job template. If this configuration does not change in subsequent deployments, you can omit this parameter. To clear the NAS configuration, set this parameter to an empty string ("").
	//
	// example:
	//
	// 10d3b4bc9****.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The ID of the SAE namespace. The namespace name can contain only lowercase letters and hyphens (-), and must start with a letter.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The configurations for mounting a NAS file system.
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// The ID of the NAS file system. If this configuration does not change in subsequent deployments, you can omit this parameter. To clear the NAS configuration, set this parameter to an empty string ("").
	//
	// example:
	//
	// 10d3b4****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The AccessKey ID for reading from and writing to OSS.
	//
	// example:
	//
	// xxxxxx
	OssAkId *string `json:"OssAkId,omitempty" xml:"OssAkId,omitempty"`
	// The AccessKey secret for reading from and writing to OSS.
	//
	// example:
	//
	// xxxxxx
	OssAkSecret *string `json:"OssAkSecret,omitempty" xml:"OssAkSecret,omitempty"`
	// The description of the Object Storage Service (OSS) mount. The parameters are described as follows:
	//
	// - **bucketName**: The name of the bucket.
	//
	// - **bucketPath**: The directory or object in OSS. If the specified directory or object does not exist, an exception is thrown.
	//
	// - **mountPath**: The path in the SAE container. If the path exists, it is overwritten. If the path does not exist, it is created.
	//
	// - **readOnly**: Specifies whether the container has read-only access to the resources in the mount directory. Valid values:
	//
	//   - **true**: read-only permission.
	//
	//   - **false**: read and write permissions.
	//
	// example:
	//
	// [{"bucketName": "oss-bucket", "bucketPath": "data/user.data", "mountPath": "/usr/data/user.data", "readOnly": true}]
	OssMountDescs *string `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty"`
	// The type of the deployment package. Valid values:
	//
	// - For Java applications, valid values are **FatJar**, **War**, and **Image**.
	//
	// - For PHP applications, the valid values are:
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
	// - For **Python*	- applications, valid values are **PythonZip*	- and **Image**.
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
	// http://myoss.oss-cn-hangzhou.aliyuncs.com/my-buc/2019-06-30/****.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package. This parameter is required when **PackageType*	- is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// 1.0.1
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The content of the PHP configuration file.
	//
	// example:
	//
	// k1=v1
	PhpConfig *string `json:"PhpConfig,omitempty" xml:"PhpConfig,omitempty"`
	// The mount path of the startup configuration file for a PHP task. You must make sure that the PHP server uses this configuration file on startup.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// A PostStart hook. This script runs immediately after the container is created. The value must be a JSON string, for example: `{"exec":{"command":["sh","-c","echo hello"]}}`
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// A PreStop hook. This script runs immediately before the container is stopped. The value must be a JSON string, for example: `{"exec":{"command":["sh","-c","echo hello"]}}`
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The programming language. Valid values: **java**, **php**, **python**, and **shell**.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// The Python environment. **PYTHON 3.9.15*	- is supported.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// Python dependencies to install by using pip. If you do not set this parameter, SAE installs dependencies from the \\"requirements.txt\\" file in the root directory of your project.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// The ID of the referenced job.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	RefAppId *string `json:"RefAppId,omitempty" xml:"RefAppId,omitempty"`
	// The number of concurrent task instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Specifies whether to enable task sharding.
	//
	// example:
	//
	// true
	Slice *bool `json:"Slice,omitempty" xml:"Slice,omitempty"`
	// The parameters for task sharding.
	//
	// example:
	//
	// [0,1,2]
	SliceEnvs *string `json:"SliceEnvs,omitempty" xml:"SliceEnvs,omitempty"`
	// The configuration for collecting logs to Simple Log Service (SLS).
	//
	// - To use SLS resources that are automatically created by SAE: `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - To use your own SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// The parameters are described as follows:
	//
	// - **projectName**: The name of the SLS Project.
	//
	// - **logDir**: The path of the log file.
	//
	// - **logType**: The log type. **stdout*	- indicates the standard output of the container. You can specify only one standard output. If you do not set this parameter, file logs are collected.
	//
	// - **logstoreName**: The name of the Logstore in SLS.
	//
	// - **logtailName**: The name of the Logtail in SLS. If you do not specify this parameter, a new Logtail is created.
	//
	// If the log collection configuration does not change during subsequent deployments, you do not need to set this parameter (the request does not need to include the **SlsConfigs*	- field). If you no longer need to use the log collection feature, set the value of this parameter to an empty string ("") in your request.
	//
	// > SAE deletes a project that it automatically created when you delete the corresponding job template. Therefore, if you specify an existing project, do not use one that was automatically created by SAE.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// The graceful shutdown timeout, in seconds. The value must be an integer from 1 to 300. Default: 30.
	//
	// example:
	//
	// 10
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The task timeout, in seconds.
	//
	// example:
	//
	// 3600
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The time zone. Default value: **Asia/Shanghai**.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The Tomcat configuration. To delete the configuration, set this parameter to `""` or `{}`. The parameters are described as follows:
	//
	// - **port**: The port number. The valid range is 1024 to 65535. Ports below 1024 require root permissions. Because the container is configured with administrator permissions, specify a port number greater than 1024. If this parameter is not configured, the default port 8080 is used.
	//
	// - **contextPath**: The context path. Default value: /.
	//
	// - **maxThreads**: The maximum number of threads in the connection pool. Default value: 400.
	//
	// - **uriEncoding**: The URI encoding scheme for Tomcat. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. If this parameter is not set, the default value **ISO-8859-1*	- is used.
	//
	// - **useBodyEncodingForUri**: Specifies whether to use the encoding specified in `request.getCharacterEncoding()` to decode the request URI. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig  *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	TriggerConfig *string `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
	// The ID of the vSwitch for the elastic network interface of the task instance. The vSwitch must be located in the specified VPC. The vSwitch is also bound to the SAE namespace. If you do not specify this parameter, the ID of the vSwitch that is bound to the namespace is used by default.
	//
	// example:
	//
	// vsw-bp12mw1f8k3jgygk9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC for the SAE namespace. In SAE, a namespace can be bound to only one VPC, and this binding cannot be changed. The binding is established when you create the first SAE job template in the namespace. A single VPC can be associated with multiple namespaces. If you do not specify this parameter, the ID of the VPC that is bound to the namespace is used by default.
	//
	// example:
	//
	// vpc-bp1aevy8sofi8mh1q****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The startup command for a WAR package deployment. The configuration steps are the same as for an image-based deployment. For more information, see [Set a startup command](https://help.aliyun.com/document_detail/96677.html).
	//
	// example:
	//
	// CATALINA_OPTS=\\"$CATALINA_OPTS $Options\\" catalina.sh run
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The Tomcat version that the deployment package requires. The following versions are supported:
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
	// The workload. Set the value to `job`.
	//
	// This parameter is required.
	//
	// example:
	//
	// job
	Workload *string `json:"Workload,omitempty" xml:"Workload,omitempty"`
}

func (s CreateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) GetAcrAssumeRoleArn() *string {
	return s.AcrAssumeRoleArn
}

func (s *CreateJobRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CreateJobRequest) GetAppDescription() *string {
	return s.AppDescription
}

func (s *CreateJobRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateJobRequest) GetAutoConfig() *bool {
	return s.AutoConfig
}

func (s *CreateJobRequest) GetBackoffLimit() *int64 {
	return s.BackoffLimit
}

func (s *CreateJobRequest) GetBestEffortType() *string {
	return s.BestEffortType
}

func (s *CreateJobRequest) GetCommand() *string {
	return s.Command
}

func (s *CreateJobRequest) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *CreateJobRequest) GetConcurrencyPolicy() *string {
	return s.ConcurrencyPolicy
}

func (s *CreateJobRequest) GetConfigMapMountDesc() *string {
	return s.ConfigMapMountDesc
}

func (s *CreateJobRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateJobRequest) GetCustomHostAlias() *string {
	return s.CustomHostAlias
}

func (s *CreateJobRequest) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *CreateJobRequest) GetEnableImageAccl() *bool {
	return s.EnableImageAccl
}

func (s *CreateJobRequest) GetEnvs() *string {
	return s.Envs
}

func (s *CreateJobRequest) GetImagePullSecrets() *string {
	return s.ImagePullSecrets
}

func (s *CreateJobRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateJobRequest) GetJarStartArgs() *string {
	return s.JarStartArgs
}

func (s *CreateJobRequest) GetJarStartOptions() *string {
	return s.JarStartOptions
}

func (s *CreateJobRequest) GetJdk() *string {
	return s.Jdk
}

func (s *CreateJobRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateJobRequest) GetMountDesc() *string {
	return s.MountDesc
}

func (s *CreateJobRequest) GetMountHost() *string {
	return s.MountHost
}

func (s *CreateJobRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateJobRequest) GetNasConfigs() *string {
	return s.NasConfigs
}

func (s *CreateJobRequest) GetNasId() *string {
	return s.NasId
}

func (s *CreateJobRequest) GetOssAkId() *string {
	return s.OssAkId
}

func (s *CreateJobRequest) GetOssAkSecret() *string {
	return s.OssAkSecret
}

func (s *CreateJobRequest) GetOssMountDescs() *string {
	return s.OssMountDescs
}

func (s *CreateJobRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *CreateJobRequest) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *CreateJobRequest) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *CreateJobRequest) GetPhpConfig() *string {
	return s.PhpConfig
}

func (s *CreateJobRequest) GetPhpConfigLocation() *string {
	return s.PhpConfigLocation
}

func (s *CreateJobRequest) GetPostStart() *string {
	return s.PostStart
}

func (s *CreateJobRequest) GetPreStop() *string {
	return s.PreStop
}

func (s *CreateJobRequest) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *CreateJobRequest) GetPython() *string {
	return s.Python
}

func (s *CreateJobRequest) GetPythonModules() *string {
	return s.PythonModules
}

func (s *CreateJobRequest) GetRefAppId() *string {
	return s.RefAppId
}

func (s *CreateJobRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *CreateJobRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateJobRequest) GetSlice() *bool {
	return s.Slice
}

func (s *CreateJobRequest) GetSliceEnvs() *string {
	return s.SliceEnvs
}

func (s *CreateJobRequest) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *CreateJobRequest) GetTerminationGracePeriodSeconds() *int32 {
	return s.TerminationGracePeriodSeconds
}

func (s *CreateJobRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CreateJobRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateJobRequest) GetTomcatConfig() *string {
	return s.TomcatConfig
}

func (s *CreateJobRequest) GetTriggerConfig() *string {
	return s.TriggerConfig
}

func (s *CreateJobRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateJobRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateJobRequest) GetWarStartOptions() *string {
	return s.WarStartOptions
}

func (s *CreateJobRequest) GetWebContainer() *string {
	return s.WebContainer
}

func (s *CreateJobRequest) GetWorkload() *string {
	return s.Workload
}

func (s *CreateJobRequest) SetAcrAssumeRoleArn(v string) *CreateJobRequest {
	s.AcrAssumeRoleArn = &v
	return s
}

func (s *CreateJobRequest) SetAcrInstanceId(v string) *CreateJobRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateJobRequest) SetAppDescription(v string) *CreateJobRequest {
	s.AppDescription = &v
	return s
}

func (s *CreateJobRequest) SetAppName(v string) *CreateJobRequest {
	s.AppName = &v
	return s
}

func (s *CreateJobRequest) SetAutoConfig(v bool) *CreateJobRequest {
	s.AutoConfig = &v
	return s
}

func (s *CreateJobRequest) SetBackoffLimit(v int64) *CreateJobRequest {
	s.BackoffLimit = &v
	return s
}

func (s *CreateJobRequest) SetBestEffortType(v string) *CreateJobRequest {
	s.BestEffortType = &v
	return s
}

func (s *CreateJobRequest) SetCommand(v string) *CreateJobRequest {
	s.Command = &v
	return s
}

func (s *CreateJobRequest) SetCommandArgs(v string) *CreateJobRequest {
	s.CommandArgs = &v
	return s
}

func (s *CreateJobRequest) SetConcurrencyPolicy(v string) *CreateJobRequest {
	s.ConcurrencyPolicy = &v
	return s
}

func (s *CreateJobRequest) SetConfigMapMountDesc(v string) *CreateJobRequest {
	s.ConfigMapMountDesc = &v
	return s
}

func (s *CreateJobRequest) SetCpu(v int32) *CreateJobRequest {
	s.Cpu = &v
	return s
}

func (s *CreateJobRequest) SetCustomHostAlias(v string) *CreateJobRequest {
	s.CustomHostAlias = &v
	return s
}

func (s *CreateJobRequest) SetEdasContainerVersion(v string) *CreateJobRequest {
	s.EdasContainerVersion = &v
	return s
}

func (s *CreateJobRequest) SetEnableImageAccl(v bool) *CreateJobRequest {
	s.EnableImageAccl = &v
	return s
}

func (s *CreateJobRequest) SetEnvs(v string) *CreateJobRequest {
	s.Envs = &v
	return s
}

func (s *CreateJobRequest) SetImagePullSecrets(v string) *CreateJobRequest {
	s.ImagePullSecrets = &v
	return s
}

func (s *CreateJobRequest) SetImageUrl(v string) *CreateJobRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateJobRequest) SetJarStartArgs(v string) *CreateJobRequest {
	s.JarStartArgs = &v
	return s
}

func (s *CreateJobRequest) SetJarStartOptions(v string) *CreateJobRequest {
	s.JarStartOptions = &v
	return s
}

func (s *CreateJobRequest) SetJdk(v string) *CreateJobRequest {
	s.Jdk = &v
	return s
}

func (s *CreateJobRequest) SetMemory(v int32) *CreateJobRequest {
	s.Memory = &v
	return s
}

func (s *CreateJobRequest) SetMountDesc(v string) *CreateJobRequest {
	s.MountDesc = &v
	return s
}

func (s *CreateJobRequest) SetMountHost(v string) *CreateJobRequest {
	s.MountHost = &v
	return s
}

func (s *CreateJobRequest) SetNamespaceId(v string) *CreateJobRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateJobRequest) SetNasConfigs(v string) *CreateJobRequest {
	s.NasConfigs = &v
	return s
}

func (s *CreateJobRequest) SetNasId(v string) *CreateJobRequest {
	s.NasId = &v
	return s
}

func (s *CreateJobRequest) SetOssAkId(v string) *CreateJobRequest {
	s.OssAkId = &v
	return s
}

func (s *CreateJobRequest) SetOssAkSecret(v string) *CreateJobRequest {
	s.OssAkSecret = &v
	return s
}

func (s *CreateJobRequest) SetOssMountDescs(v string) *CreateJobRequest {
	s.OssMountDescs = &v
	return s
}

func (s *CreateJobRequest) SetPackageType(v string) *CreateJobRequest {
	s.PackageType = &v
	return s
}

func (s *CreateJobRequest) SetPackageUrl(v string) *CreateJobRequest {
	s.PackageUrl = &v
	return s
}

func (s *CreateJobRequest) SetPackageVersion(v string) *CreateJobRequest {
	s.PackageVersion = &v
	return s
}

func (s *CreateJobRequest) SetPhpConfig(v string) *CreateJobRequest {
	s.PhpConfig = &v
	return s
}

func (s *CreateJobRequest) SetPhpConfigLocation(v string) *CreateJobRequest {
	s.PhpConfigLocation = &v
	return s
}

func (s *CreateJobRequest) SetPostStart(v string) *CreateJobRequest {
	s.PostStart = &v
	return s
}

func (s *CreateJobRequest) SetPreStop(v string) *CreateJobRequest {
	s.PreStop = &v
	return s
}

func (s *CreateJobRequest) SetProgrammingLanguage(v string) *CreateJobRequest {
	s.ProgrammingLanguage = &v
	return s
}

func (s *CreateJobRequest) SetPython(v string) *CreateJobRequest {
	s.Python = &v
	return s
}

func (s *CreateJobRequest) SetPythonModules(v string) *CreateJobRequest {
	s.PythonModules = &v
	return s
}

func (s *CreateJobRequest) SetRefAppId(v string) *CreateJobRequest {
	s.RefAppId = &v
	return s
}

func (s *CreateJobRequest) SetReplicas(v int32) *CreateJobRequest {
	s.Replicas = &v
	return s
}

func (s *CreateJobRequest) SetSecurityGroupId(v string) *CreateJobRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateJobRequest) SetSlice(v bool) *CreateJobRequest {
	s.Slice = &v
	return s
}

func (s *CreateJobRequest) SetSliceEnvs(v string) *CreateJobRequest {
	s.SliceEnvs = &v
	return s
}

func (s *CreateJobRequest) SetSlsConfigs(v string) *CreateJobRequest {
	s.SlsConfigs = &v
	return s
}

func (s *CreateJobRequest) SetTerminationGracePeriodSeconds(v int32) *CreateJobRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *CreateJobRequest) SetTimeout(v int64) *CreateJobRequest {
	s.Timeout = &v
	return s
}

func (s *CreateJobRequest) SetTimezone(v string) *CreateJobRequest {
	s.Timezone = &v
	return s
}

func (s *CreateJobRequest) SetTomcatConfig(v string) *CreateJobRequest {
	s.TomcatConfig = &v
	return s
}

func (s *CreateJobRequest) SetTriggerConfig(v string) *CreateJobRequest {
	s.TriggerConfig = &v
	return s
}

func (s *CreateJobRequest) SetVSwitchId(v string) *CreateJobRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateJobRequest) SetVpcId(v string) *CreateJobRequest {
	s.VpcId = &v
	return s
}

func (s *CreateJobRequest) SetWarStartOptions(v string) *CreateJobRequest {
	s.WarStartOptions = &v
	return s
}

func (s *CreateJobRequest) SetWebContainer(v string) *CreateJobRequest {
	s.WebContainer = &v
	return s
}

func (s *CreateJobRequest) SetWorkload(v string) *CreateJobRequest {
	s.Workload = &v
	return s
}

func (s *CreateJobRequest) Validate() error {
	return dara.Validate(s)
}

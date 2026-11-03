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
	// The Alibaba Cloud Resource Name (ARN) of the Resource Access Management (RAM) role required to pull images across accounts. For more information, see [Use a RAM role to authorize access across Alibaba Cloud accounts](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The instance ID of the ACR Enterprise instance. This parameter is required when **ImageUrl*	- is set to an ACR Enterprise instance image.
	//
	// example:
	//
	// cri-xxxxxx
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The description of the job template. The description can be up to 1,024 characters in length.
	//
	// example:
	//
	// This is a test description.
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The name of the job template. The name can contain digits, letters, and hyphens (-). It must start with a letter and can be up to 36 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether to automatically configure the network environment. Valid values:
	//
	// - **true**: SAE automatically configures the network environment when creating the job template. The values of **NamespaceId**, **VpcId**, **vSwitchId**, and **SecurityGroupId*	- are ignored.
	//
	// - **false**: You manually configure the network environment when creating the job template.
	//
	// example:
	//
	// false
	AutoConfig *bool `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty"`
	// The number of retries for the job.
	//
	// example:
	//
	// 3
	BackoffLimit *int64 `json:"BackoffLimit,omitempty" xml:"BackoffLimit,omitempty"`
	// The BestEffort policy.
	//
	// example:
	//
	// besteffort
	BestEffortType *string `json:"BestEffortType,omitempty" xml:"BestEffortType,omitempty"`
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
	// Based on the preceding example, Command is set to echo and CommandArgs is set to ["abc", ">", "file0"].
	//
	// example:
	//
	// echo
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments for the image startup command. These are the arguments required by the **Command*	- parameter. Format:
	//
	// `["a","b"]`
	//
	// In the preceding example, CommandArgs is set to ["abc", ">", "file0"]. The value ["abc", ">", "file0"] must be converted to a string in JSON array format. If this parameter is not required, you do not need to specify it.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The concurrency policy of the job. Valid values:
	//
	// - **Forbid**: forbids concurrent execution. A new job is not created if the previous job is not completed.
	//
	// - **Allow**: allows concurrent execution of jobs.
	//
	// - **Replace**: replaces the previous job with a new job if the previous job is not completed when the creation time of the new job arrives.
	//
	// example:
	//
	// Allow
	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" xml:"ConcurrencyPolicy,omitempty"`
	// The mount description of the **ConfigMap**. Uses the configuration items created on the namespace configuration item page to inject configuration information into the container. The parameter descriptions are as follows:
	//
	// - **congfigMapId**: the instance ID of the ConfigMap instance. You can invoke the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the instance ID.
	//
	// - **key**: the key value.
	//
	// > You can pass the sae-sys-configmap-all parameter to mount all keys.
	//
	// - **mountPath**: the mount path.
	//
	// example:
	//
	// [{"configMapId":16,"key":"test","mountPath":"/tmp"}]
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	// The number of CPU cores required for each instance. Unit: millicores. The value cannot be 0. Only the following defined specifications are supported:
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
	// The runtime environment version of the job in the High-Speed Service Framework (HSF), such as the Ali-Tomcat container.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// Specifies whether to enable image acceleration. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// false
	EnableImageAccl *bool `json:"EnableImageAccl,omitempty" xml:"EnableImageAccl,omitempty"`
	// The environment variables for the container. You can customize environment variables or reference ConfigMaps. To reference a ConfigMap, you must create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//     - **name**: the name of the environment variable.
	//
	//     - **value**: the value of the environment variable.
	//
	// - Reference a ConfigMap
	//
	//     - **name**: the name of the environment variable. You can reference a single key or all keys. To reference all keys, enter sae-sys-configmap-all-<ConfigMap name>, for example, sae-sys-configmap-all-test1.
	//
	//     - **valueFrom**: the reference of the environment variable. Set the value to configMapRef.
	//
	//     - **configMapId**: the ID of the ConfigMap.
	//
	//     - **key**: the key. If you reference all key-value pairs, do not set this field.
	//
	// example:
	//
	// [{"name":"envtmp","value":"0"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The ID of the secret.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The image address. This parameter is required when **PackageType*	- is set to **Image**.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/sae_test/ali_sae_test:0.0.1
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The startup arguments for the JAR package. Default job startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// -Xms4G -Xmx4G
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The startup options for the JAR package. Default job startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// custom-option
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The Java Development Kit (JDK) version on which the deployment package depends. Valid values:
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
	// The memory size required for each instance. Unit: MB. The value cannot be 0. The memory size has a one-to-one mapping with the number of CPU cores. Only the following defined specifications are supported:
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
	// The description of the NAS mount. If the configuration remains unchanged during deployment, you do not need to set this parameter, which means the **MountDesc*	- field is not required in the request. To clear the NAS configuration, you must set this field to an empty string in the request, which means the value of the **MountDesc*	- field is "".
	//
	// example:
	//
	// [{mountPath: "/tmp", nasPath: "/"}]
	MountDesc *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	// The mount target of the NAS file system in the VPC of the job template. If the configuration remains unchanged during deployment, you do not need to set this parameter, which means the **MountHost*	- field is not required in the request. To clear the NAS configuration, you must set this field to an empty string in the request, which means the value of the **MountHost*	- field is "".
	//
	// example:
	//
	// 10d3b4bc9****.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The ID of the Serverless App Engine (SAE) namespace. Only namespaces whose names contain lowercase letters and hyphens (-) are supported. The name must start with a letter.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The configuration for mounting the NAS file system.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// The ID of the Apsara File Storage NAS file system. If the configuration remains unchanged during deployment, you do not need to set this parameter, which means the **NasId*	- field is not required in the request. To clear the NAS configuration, you must set this field to an empty string in the request, which means the value of the **NasId*	- field is "".
	//
	// example:
	//
	// 10d3b4****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
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
	// The OSS mount description. The parameter descriptions are as follows:
	//
	// - **bucketName**: the name of the bucket.
	//
	// - **bucketPath**: the folder or object that you created in OSS. If the OSS mount folder does not exist, an exception is triggered.
	//
	// - **mountPath**: the container path in SAE. If the path already exists, it overwrites the existing content. If the path does not exist, it is created.
	//
	// - **readOnly**: specifies whether the container path has read permissions on the mount folder resources. Valid values:
	//
	//     - **true**: read-only permissions.
	//
	//     - **false**: read and write permissions.
	//
	// example:
	//
	// [{"bucketName": "oss-bucket", "bucketPath": "data/user.data", "mountPath": "/usr/data/user.data", "readOnly": true}]
	OssMountDescs *string `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty"`
	// The type of the job package. Valid values:
	//
	// - If you use Java for deployment, valid values are **FatJar**, **War**, and **Image**.
	//
	// - If you use PHP for deployment, valid values are:
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
	// - If you use **Python*	- for deployment, valid values are **PythonZip*	- and **Image**.
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
	// The mount path for the PHP job startup configuration. You must ensure that the PHP server uses this configuration file to start.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// The script to execute after the container starts. A script is triggered and executed immediately after the container is created. Format: `{"exec":{"command":["sh","-c","echo hello"]}}`
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script to execute before the container stops. A script is triggered and executed before the container is deleted. Format: `{"exec":{"command":["sh","-c","echo hello"]}}`
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
	// The Python environment. Valid value: **PYTHON 3.9.15**.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// The module dependencies for custom installation. By default, the dependencies defined in the requirements.txt file in the root folder are installed. If no dependencies are configured or custom packages are used, you can specify the dependencies for custom installation.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// The ID of the referenced job target.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	RefAppId *string `json:"RefAppId,omitempty" xml:"RefAppId,omitempty"`
	// The number of concurrent instances for the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Specifies whether to enable job sharding.
	//
	// example:
	//
	// true
	Slice *bool `json:"Slice,omitempty" xml:"Slice,omitempty"`
	// The sharding parameters of the job.
	//
	// example:
	//
	// [0,1,2]
	SliceEnvs *string `json:"SliceEnvs,omitempty" xml:"SliceEnvs,omitempty"`
	// The configuration for collecting logs to Simple Log Service (SLS).
	//
	// - To use SLS resources automatically created by SAE: `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - To use custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// The parameter descriptions are as follows:
	//
	// - **projectName**: the name of the SLS project.
	//
	// - **logDir**: the log path.
	//
	// - **logType**: the log type. **stdout*	- indicates container standard output logs. You can configure only one such log source. If you do not configure this parameter, file logs are collected.
	//
	// - **logstoreName**: the name of the SLS Logstore.
	//
	// - **logtailName**: the name of the SLS Logtail. If you do not specify this parameter, a new Logtail is created.
	//
	// If the SLS collection configuration remains unchanged during multiple deployments, you do not need to set this parameter (that is, the **SlsConfigs*	- field is not required in the request). If you no longer need the SLS collection feature, you must set this field to an empty string (a string with zero characters) in the request (that is, the value of the **SlsConfigs*	- field is "").
	//
	// > The project automatically created with the job template is deleted when the job template is deleted. Therefore, when you select an existing project, you cannot select a project automatically created by SAE.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// The timeout period for graceful shutdown. Default value: 30. Unit: seconds. Valid values: 1 to 300.
	//
	// example:
	//
	// 10
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The timeout period of the job. Unit: seconds.
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
	// The Tomcat configuration. Set the value to "" or "{}" to delete the configuration. Valid values:
	//
	// - **port**: the port range is 1024 to 65535. Ports smaller than 1024 require root permissions. Because the container is configured with admin permissions, specify a port greater than 1024. Default value: 8080.
	//
	// - **contextPath**: the access path. Default value: the root directory "/".
	//
	// - **maxThreads**: the maximum number of connections in the connection pool. Default value: 400.
	//
	// - **uriEncoding**: the encoding format of Tomcat. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. Default value: **ISO-8859-1**.
	//
	// - **useBodyEncoding**: specifies whether to use **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig  *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	TriggerConfig *string `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
	// The ID of the vSwitch to which the elastic network interface controller (NIC) of the job instance belongs. The vSwitch must reside in the preceding VPC. The vSwitch is also associated with the SAE namespace. If you do not specify this parameter, the ID of the vSwitch that is associated with the namespace is used by default.
	//
	// example:
	//
	// vsw-bp12mw1f8k3jgygk9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) that is associated with the SAE namespace. In SAE, a namespace can be associated with only one VPC, and the association cannot be modified. The association is established when you create the first SAE job template in the namespace. Multiple namespaces can be associated with the same VPC. If you do not specify this parameter, the ID of the VPC that is associated with the namespace is used by default.
	//
	// example:
	//
	// vpc-bp1aevy8sofi8mh1q****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The startup command for the WAR package deployment job. The procedure is the same as configuring the startup command for image deployment. For more information, see [Configure a startup command](https://help.aliyun.com/document_detail/96677.html).
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
	// This parameter is not supported when **PackageType*	- is set to **Image**.
	//
	// example:
	//
	// apache-tomcat-7.0.91
	WebContainer *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
	// The workload type. Set the value to `job`.
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

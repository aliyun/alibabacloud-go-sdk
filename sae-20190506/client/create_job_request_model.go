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
	// The description of the template. The description cannot exceed 1,024 characters in length.
	//
	// example:
	//
	// This is a test description.
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The name of the job template. The name can contain digits, letters, and hyphens (-). The name must start with a letter and cannot exceed 36 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Specifies whether to automatically configure the network environment. Take note of the following rules:
	//
	// 	- **true**: The network environment is automatically configured by SAE when the application is created. In this case, the values of the **NamespaceId**, **VpcId**, **vSwitchId**, and **SecurityGroupId*	- parameters are ignored.
	//
	// 	- **false**: The network environment is manually configured based on your settings when the application is created.
	//
	// example:
	//
	// false
	AutoConfig *bool `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty"`
	// The number of times the job is retried.
	//
	// example:
	//
	// 3
	BackoffLimit *int64 `json:"BackoffLimit,omitempty" xml:"BackoffLimit,omitempty"`
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
	// The concurrency policy of the job. Take note of the following rules:
	//
	// 	- **Forbid**: Prohibits concurrent running. If the previous job is not completed, no new job is created.
	//
	// 	- **Allow**: Allows concurrent running.
	//
	// 	- **Replace**: If the previous job is not completed when the time to create a new job is reached, the new job replaces the previous job.
	//
	// example:
	//
	// Allow
	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" xml:"ConcurrencyPolicy,omitempty"`
	// The description of the **ConfigMap*	- instance mounted to the application. Use configurations created on the Configuration Items page to configure containers. The following table describes the parameters that are used in the preceding statements.
	//
	// 	- **congfigMapId**: the ID of the ConfigMap instance. You can call the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the ID.
	//
	// 	- **key**: the key.
	//
	// > You can use the `sae-sys-configmap-all` key to mount all keys.
	//
	// 	- **mountPath**: the mount path in the container.
	//
	// example:
	//
	// [{"configMapId":16,"key":"test","mountPath":"/tmp"}]
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	// The CPU specifications that are required for each instance. Unit: millicores. You cannot set this parameter to 0. Valid values:
	//
	// 	- 500
	//
	// 	- 1000
	//
	// 	- 2000
	//
	// 	- 4000
	//
	// 	- 8000
	//
	// 	- 16000
	//
	// 	- 32000
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
	// The version of the container, such as Ali-Tomcat, in which an application developed based on High-speed Service Framework (HSF) is deployed.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// example:
	//
	// false
	EnableImageAccl *bool `json:"EnableImageAccl,omitempty" xml:"EnableImageAccl,omitempty"`
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
	// example:
	//
	// [{"name":"envtmp","value":"0"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
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
	// The size of memory required by each instance. Unit: MB. You cannot set this parameter to 0. The values of this parameter correspond to the values of the Cpu parameter:
	//
	// 	- Set the value to 1024 when Cpu is set to 500 or 1000.
	//
	// 	- Set the value to 2048 when Cpu is set to 500, 1000 or 2000.
	//
	// 	- Set the value to 4096 when Cpu is set to 1000, 2000, or 4000.
	//
	// 	- Set the value to 8192 when Cpu is set to 2000, 4000, or 8000.
	//
	// 	- Set the value to 12288 when Cpu is set to 12000.
	//
	// 	- Set the value to 16384 when Cpu is set to 4000, 8000, or 16000.
	//
	// 	- Set the value to 24576 when Cpu is set to 12000.
	//
	// 	- Set the value to 32768 when Cpu is set to 16000.
	//
	// 	- Set the value to 65536 when Cpu is set to 8000, 16000, or 32000.
	//
	// 	- Set the value to 131072 when Cpu is set to 32000.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
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
	// The ID of the Serverless App Engine (SAE) namespace. The ID can contain only lowercase letters and hyphens (-). It must start with a lowercase letter.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The ID of the Apsara File Storage NAS file system. After the application is created, you may want to call other operations to manage the application. If you do not want to change the NAS configurations in these subsequent operations, you can omit the **NasId*	- parameter in the requests. If you want to unmount the NAS file system, you must set the **NasId*	- values in the subsequent requests to an empty string ("").
	//
	// example:
	//
	// 10d3b4****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
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
	// http://myoss.oss-cn-hangzhou.aliyuncs.com/my-buc/2019-06-30/****.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package. This parameter is required if you set **PackageType*	- to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// 1.0.1
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
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
	// The programming language. Valid values: **java**, **php**, **python**, and **shell**.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
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
	// The ID of the job that you reference.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	RefAppId *string `json:"RefAppId,omitempty" xml:"RefAppId,omitempty"`
	// The number of concurrent instances.
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
	// The parameters of job sharding.
	//
	// example:
	//
	// [0,1,2]
	SliceEnvs *string `json:"SliceEnvs,omitempty" xml:"SliceEnvs,omitempty"`
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
	// The timeout period for a graceful shutdown. Default value: 30. Unit: seconds. Valid values: 1 to 300.
	//
	// example:
	//
	// 10
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The timeout period. Unit: seconds.
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
	TomcatConfig  *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	TriggerConfig *string `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
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
	// Set the value to `job`.
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

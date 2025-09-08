// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrAssumeRoleArn(v string) *UpdateJobRequest
	GetAcrAssumeRoleArn() *string
	SetAcrInstanceId(v string) *UpdateJobRequest
	GetAcrInstanceId() *string
	SetAppId(v string) *UpdateJobRequest
	GetAppId() *string
	SetBackoffLimit(v int64) *UpdateJobRequest
	GetBackoffLimit() *int64
	SetBestEffortType(v string) *UpdateJobRequest
	GetBestEffortType() *string
	SetCommand(v string) *UpdateJobRequest
	GetCommand() *string
	SetCommandArgs(v string) *UpdateJobRequest
	GetCommandArgs() *string
	SetConcurrencyPolicy(v string) *UpdateJobRequest
	GetConcurrencyPolicy() *string
	SetConfigMapMountDesc(v string) *UpdateJobRequest
	GetConfigMapMountDesc() *string
	SetCustomHostAlias(v string) *UpdateJobRequest
	GetCustomHostAlias() *string
	SetEdasContainerVersion(v string) *UpdateJobRequest
	GetEdasContainerVersion() *string
	SetEnableImageAccl(v bool) *UpdateJobRequest
	GetEnableImageAccl() *bool
	SetEnvs(v string) *UpdateJobRequest
	GetEnvs() *string
	SetImagePullSecrets(v string) *UpdateJobRequest
	GetImagePullSecrets() *string
	SetImageUrl(v string) *UpdateJobRequest
	GetImageUrl() *string
	SetJarStartArgs(v string) *UpdateJobRequest
	GetJarStartArgs() *string
	SetJarStartOptions(v string) *UpdateJobRequest
	GetJarStartOptions() *string
	SetJdk(v string) *UpdateJobRequest
	GetJdk() *string
	SetMountDesc(v string) *UpdateJobRequest
	GetMountDesc() *string
	SetMountHost(v string) *UpdateJobRequest
	GetMountHost() *string
	SetNasConfigs(v string) *UpdateJobRequest
	GetNasConfigs() *string
	SetNasId(v string) *UpdateJobRequest
	GetNasId() *string
	SetOssAkId(v string) *UpdateJobRequest
	GetOssAkId() *string
	SetOssAkSecret(v string) *UpdateJobRequest
	GetOssAkSecret() *string
	SetOssMountDescs(v string) *UpdateJobRequest
	GetOssMountDescs() *string
	SetPackageUrl(v string) *UpdateJobRequest
	GetPackageUrl() *string
	SetPackageVersion(v string) *UpdateJobRequest
	GetPackageVersion() *string
	SetPhp(v string) *UpdateJobRequest
	GetPhp() *string
	SetPhpConfig(v string) *UpdateJobRequest
	GetPhpConfig() *string
	SetPhpConfigLocation(v string) *UpdateJobRequest
	GetPhpConfigLocation() *string
	SetPostStart(v string) *UpdateJobRequest
	GetPostStart() *string
	SetPreStop(v string) *UpdateJobRequest
	GetPreStop() *string
	SetProgrammingLanguage(v string) *UpdateJobRequest
	GetProgrammingLanguage() *string
	SetPython(v string) *UpdateJobRequest
	GetPython() *string
	SetPythonModules(v string) *UpdateJobRequest
	GetPythonModules() *string
	SetRefAppId(v string) *UpdateJobRequest
	GetRefAppId() *string
	SetReplicas(v string) *UpdateJobRequest
	GetReplicas() *string
	SetSlice(v bool) *UpdateJobRequest
	GetSlice() *bool
	SetSliceEnvs(v string) *UpdateJobRequest
	GetSliceEnvs() *string
	SetSlsConfigs(v string) *UpdateJobRequest
	GetSlsConfigs() *string
	SetTerminationGracePeriodSeconds(v int32) *UpdateJobRequest
	GetTerminationGracePeriodSeconds() *int32
	SetTimeout(v int64) *UpdateJobRequest
	GetTimeout() *int64
	SetTimezone(v string) *UpdateJobRequest
	GetTimezone() *string
	SetTomcatConfig(v string) *UpdateJobRequest
	GetTomcatConfig() *string
	SetTriggerConfig(v string) *UpdateJobRequest
	GetTriggerConfig() *string
	SetWarStartOptions(v string) *UpdateJobRequest
	GetWarStartOptions() *string
	SetWebContainer(v string) *UpdateJobRequest
	GetWebContainer() *string
}

type UpdateJobRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is used to pull images across accounts. For more information, see [Grant permissions across Alibaba Cloud accounts by using a RAM role](https://help.aliyun.com/document_detail/223585.html).
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
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The number of times the job is retried.
	//
	// example:
	//
	// 3
	BackoffLimit   *int64  `json:"BackoffLimit,omitempty" xml:"BackoffLimit,omitempty"`
	BestEffortType *string `json:"BestEffortType,omitempty" xml:"BestEffortType,omitempty"`
	// The command that is used to start the image. The command must be an existing executable object in the container. Example:
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
	// The parameters of the image startup command. The CommandArgs parameter specifies the parameters that are required for the **Command*	- parameter. The name must meet the following format requirements:
	//
	// `["a","b"]`
	//
	// In the preceding example, the CommandArgs parameter is set to `CommandArgs=["abc", ">", "file0"]`. The data type of `["abc", ">", "file0"]` must be an array of strings in the JSON format. This parameter is optional.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The concurrency policy of the job. Valid values:
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
	// The description of the **ConfigMap*	- instance mounted to the application. Use configurations created on the Configuration Items page to configure containers. The following parameters are involved:
	//
	// 	- **congfigMapId**: the ID of the ConfigMap instance. You can call the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the ID.
	//
	// 	- **key**: the key.
	//
	// > You can use the `sae-sys-configmap-all` key to mount all keys.
	//
	// 	- **mountPath**: the mount path.
	//
	// example:
	//
	// [{"configMapId":16,"key":"test","mountPath":"/tmp"}]
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	// The custom mappings between hostnames and IP addresses in the container. Valid values:
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
	EnableImageAccl      *bool   `json:"EnableImageAccl,omitempty" xml:"EnableImageAccl,omitempty"`
	// The environment variables. You can configure custom environment variables or reference a ConfigMap. If you want to reference a ConfigMap, you must first create a ConfigMap. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// 	- Configure custom environment variables
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
	// The configurations for mounting the NAS file system. If you do not need to modify the NAS configurations when you deploy the application, configure **MountDesc*	- only in the first request. If you no longer need to use NAS, leave **MountDesc*	- empty in the request.
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
	MountHost  *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// The ID of the Apsara File Storage NAS file system. If you do not need to modify the NAS configurations when you deploy the application, configure **NasId*	- only in the first request. If you no longer need to use NAS, leave **NasId*	- empty in the request.
	//
	// example:
	//
	// 10d3b4****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The AccessKey ID that is used to read data from and write data to OSS.
	//
	// example:
	//
	// xxxxxx
	OssAkId *string `json:"OssAkId,omitempty" xml:"OssAkId,omitempty"`
	// The AccessKey secret that is used to read data from and write data to OSS.
	//
	// example:
	//
	// xxxxxx
	OssAkSecret *string `json:"OssAkSecret,omitempty" xml:"OssAkSecret,omitempty"`
	// The information about the mounted Object Storage Service (OSS) bucket. The following parameters are involved:
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
	// The ID of Container Registry Enterprise Edition instance N.
	//
	// example:
	//
	// cri-xxxxxx
	Php *string `json:"Php,omitempty" xml:"Php,omitempty"`
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
	// The script to be run after the container is started. Example: `{"exec":{"command":["sh","-c","echo hello"\\]}}`
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
	// example:
	//
	// 3
	Replicas *string `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
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
	// The configurations of Log Service.
	//
	// 	- To use Log Service resources that are automatically created by SAE, set this parameter to `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// 	- To use custom Log Service resources, set this parameter to `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// The following parameters are involved:
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
	// The Tomcat configuration. If you want to delete the configuration, set this parameter to {} or leave this parameter empty. Valid values:
	//
	// 	- **port**: the port number. The port number ranges from 1024 to 65535. Though the admin permissions are configured for the container, the root permissions are required to perform operations on ports whose number is smaller than 1024. Enter a value that ranges from 1025 to 65535 because the container has only the admin permissions. If you do not specify this parameter, the default port number 8080 is used.
	//
	// 	- **contextPath**: the path. Default value: /. This value indicates the root directory.
	//
	// 	- **maxThreads**: the maximum number of connections in the connection pool. Default value: 400.
	//
	// 	- **uriEncoding**: the URI encoding scheme in the Tomcat container. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and GB2312. If you do not specify this parameter, the default value **ISO-8859-1*	- is used.
	//
	// 	- **useBodyEncoding**: specifies whether to use the encoding scheme that is specified by **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig  *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	TriggerConfig *string `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
	// The startup command of the WAR package. For information about how to configure the startup command, see [Configure startup commands](https://help.aliyun.com/document_detail/96677.html).
	//
	// example:
	//
	// CATALINA_OPTS=\\"$CATALINA_OPTS $Options\\" catalina.sh run
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The version of the Tomcat container on which the deployment package depends. The following versions are supported:
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

func (s UpdateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) GetAcrAssumeRoleArn() *string {
	return s.AcrAssumeRoleArn
}

func (s *UpdateJobRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *UpdateJobRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateJobRequest) GetBackoffLimit() *int64 {
	return s.BackoffLimit
}

func (s *UpdateJobRequest) GetBestEffortType() *string {
	return s.BestEffortType
}

func (s *UpdateJobRequest) GetCommand() *string {
	return s.Command
}

func (s *UpdateJobRequest) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *UpdateJobRequest) GetConcurrencyPolicy() *string {
	return s.ConcurrencyPolicy
}

func (s *UpdateJobRequest) GetConfigMapMountDesc() *string {
	return s.ConfigMapMountDesc
}

func (s *UpdateJobRequest) GetCustomHostAlias() *string {
	return s.CustomHostAlias
}

func (s *UpdateJobRequest) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *UpdateJobRequest) GetEnableImageAccl() *bool {
	return s.EnableImageAccl
}

func (s *UpdateJobRequest) GetEnvs() *string {
	return s.Envs
}

func (s *UpdateJobRequest) GetImagePullSecrets() *string {
	return s.ImagePullSecrets
}

func (s *UpdateJobRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *UpdateJobRequest) GetJarStartArgs() *string {
	return s.JarStartArgs
}

func (s *UpdateJobRequest) GetJarStartOptions() *string {
	return s.JarStartOptions
}

func (s *UpdateJobRequest) GetJdk() *string {
	return s.Jdk
}

func (s *UpdateJobRequest) GetMountDesc() *string {
	return s.MountDesc
}

func (s *UpdateJobRequest) GetMountHost() *string {
	return s.MountHost
}

func (s *UpdateJobRequest) GetNasConfigs() *string {
	return s.NasConfigs
}

func (s *UpdateJobRequest) GetNasId() *string {
	return s.NasId
}

func (s *UpdateJobRequest) GetOssAkId() *string {
	return s.OssAkId
}

func (s *UpdateJobRequest) GetOssAkSecret() *string {
	return s.OssAkSecret
}

func (s *UpdateJobRequest) GetOssMountDescs() *string {
	return s.OssMountDescs
}

func (s *UpdateJobRequest) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *UpdateJobRequest) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *UpdateJobRequest) GetPhp() *string {
	return s.Php
}

func (s *UpdateJobRequest) GetPhpConfig() *string {
	return s.PhpConfig
}

func (s *UpdateJobRequest) GetPhpConfigLocation() *string {
	return s.PhpConfigLocation
}

func (s *UpdateJobRequest) GetPostStart() *string {
	return s.PostStart
}

func (s *UpdateJobRequest) GetPreStop() *string {
	return s.PreStop
}

func (s *UpdateJobRequest) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *UpdateJobRequest) GetPython() *string {
	return s.Python
}

func (s *UpdateJobRequest) GetPythonModules() *string {
	return s.PythonModules
}

func (s *UpdateJobRequest) GetRefAppId() *string {
	return s.RefAppId
}

func (s *UpdateJobRequest) GetReplicas() *string {
	return s.Replicas
}

func (s *UpdateJobRequest) GetSlice() *bool {
	return s.Slice
}

func (s *UpdateJobRequest) GetSliceEnvs() *string {
	return s.SliceEnvs
}

func (s *UpdateJobRequest) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *UpdateJobRequest) GetTerminationGracePeriodSeconds() *int32 {
	return s.TerminationGracePeriodSeconds
}

func (s *UpdateJobRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *UpdateJobRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdateJobRequest) GetTomcatConfig() *string {
	return s.TomcatConfig
}

func (s *UpdateJobRequest) GetTriggerConfig() *string {
	return s.TriggerConfig
}

func (s *UpdateJobRequest) GetWarStartOptions() *string {
	return s.WarStartOptions
}

func (s *UpdateJobRequest) GetWebContainer() *string {
	return s.WebContainer
}

func (s *UpdateJobRequest) SetAcrAssumeRoleArn(v string) *UpdateJobRequest {
	s.AcrAssumeRoleArn = &v
	return s
}

func (s *UpdateJobRequest) SetAcrInstanceId(v string) *UpdateJobRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *UpdateJobRequest) SetAppId(v string) *UpdateJobRequest {
	s.AppId = &v
	return s
}

func (s *UpdateJobRequest) SetBackoffLimit(v int64) *UpdateJobRequest {
	s.BackoffLimit = &v
	return s
}

func (s *UpdateJobRequest) SetBestEffortType(v string) *UpdateJobRequest {
	s.BestEffortType = &v
	return s
}

func (s *UpdateJobRequest) SetCommand(v string) *UpdateJobRequest {
	s.Command = &v
	return s
}

func (s *UpdateJobRequest) SetCommandArgs(v string) *UpdateJobRequest {
	s.CommandArgs = &v
	return s
}

func (s *UpdateJobRequest) SetConcurrencyPolicy(v string) *UpdateJobRequest {
	s.ConcurrencyPolicy = &v
	return s
}

func (s *UpdateJobRequest) SetConfigMapMountDesc(v string) *UpdateJobRequest {
	s.ConfigMapMountDesc = &v
	return s
}

func (s *UpdateJobRequest) SetCustomHostAlias(v string) *UpdateJobRequest {
	s.CustomHostAlias = &v
	return s
}

func (s *UpdateJobRequest) SetEdasContainerVersion(v string) *UpdateJobRequest {
	s.EdasContainerVersion = &v
	return s
}

func (s *UpdateJobRequest) SetEnableImageAccl(v bool) *UpdateJobRequest {
	s.EnableImageAccl = &v
	return s
}

func (s *UpdateJobRequest) SetEnvs(v string) *UpdateJobRequest {
	s.Envs = &v
	return s
}

func (s *UpdateJobRequest) SetImagePullSecrets(v string) *UpdateJobRequest {
	s.ImagePullSecrets = &v
	return s
}

func (s *UpdateJobRequest) SetImageUrl(v string) *UpdateJobRequest {
	s.ImageUrl = &v
	return s
}

func (s *UpdateJobRequest) SetJarStartArgs(v string) *UpdateJobRequest {
	s.JarStartArgs = &v
	return s
}

func (s *UpdateJobRequest) SetJarStartOptions(v string) *UpdateJobRequest {
	s.JarStartOptions = &v
	return s
}

func (s *UpdateJobRequest) SetJdk(v string) *UpdateJobRequest {
	s.Jdk = &v
	return s
}

func (s *UpdateJobRequest) SetMountDesc(v string) *UpdateJobRequest {
	s.MountDesc = &v
	return s
}

func (s *UpdateJobRequest) SetMountHost(v string) *UpdateJobRequest {
	s.MountHost = &v
	return s
}

func (s *UpdateJobRequest) SetNasConfigs(v string) *UpdateJobRequest {
	s.NasConfigs = &v
	return s
}

func (s *UpdateJobRequest) SetNasId(v string) *UpdateJobRequest {
	s.NasId = &v
	return s
}

func (s *UpdateJobRequest) SetOssAkId(v string) *UpdateJobRequest {
	s.OssAkId = &v
	return s
}

func (s *UpdateJobRequest) SetOssAkSecret(v string) *UpdateJobRequest {
	s.OssAkSecret = &v
	return s
}

func (s *UpdateJobRequest) SetOssMountDescs(v string) *UpdateJobRequest {
	s.OssMountDescs = &v
	return s
}

func (s *UpdateJobRequest) SetPackageUrl(v string) *UpdateJobRequest {
	s.PackageUrl = &v
	return s
}

func (s *UpdateJobRequest) SetPackageVersion(v string) *UpdateJobRequest {
	s.PackageVersion = &v
	return s
}

func (s *UpdateJobRequest) SetPhp(v string) *UpdateJobRequest {
	s.Php = &v
	return s
}

func (s *UpdateJobRequest) SetPhpConfig(v string) *UpdateJobRequest {
	s.PhpConfig = &v
	return s
}

func (s *UpdateJobRequest) SetPhpConfigLocation(v string) *UpdateJobRequest {
	s.PhpConfigLocation = &v
	return s
}

func (s *UpdateJobRequest) SetPostStart(v string) *UpdateJobRequest {
	s.PostStart = &v
	return s
}

func (s *UpdateJobRequest) SetPreStop(v string) *UpdateJobRequest {
	s.PreStop = &v
	return s
}

func (s *UpdateJobRequest) SetProgrammingLanguage(v string) *UpdateJobRequest {
	s.ProgrammingLanguage = &v
	return s
}

func (s *UpdateJobRequest) SetPython(v string) *UpdateJobRequest {
	s.Python = &v
	return s
}

func (s *UpdateJobRequest) SetPythonModules(v string) *UpdateJobRequest {
	s.PythonModules = &v
	return s
}

func (s *UpdateJobRequest) SetRefAppId(v string) *UpdateJobRequest {
	s.RefAppId = &v
	return s
}

func (s *UpdateJobRequest) SetReplicas(v string) *UpdateJobRequest {
	s.Replicas = &v
	return s
}

func (s *UpdateJobRequest) SetSlice(v bool) *UpdateJobRequest {
	s.Slice = &v
	return s
}

func (s *UpdateJobRequest) SetSliceEnvs(v string) *UpdateJobRequest {
	s.SliceEnvs = &v
	return s
}

func (s *UpdateJobRequest) SetSlsConfigs(v string) *UpdateJobRequest {
	s.SlsConfigs = &v
	return s
}

func (s *UpdateJobRequest) SetTerminationGracePeriodSeconds(v int32) *UpdateJobRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *UpdateJobRequest) SetTimeout(v int64) *UpdateJobRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateJobRequest) SetTimezone(v string) *UpdateJobRequest {
	s.Timezone = &v
	return s
}

func (s *UpdateJobRequest) SetTomcatConfig(v string) *UpdateJobRequest {
	s.TomcatConfig = &v
	return s
}

func (s *UpdateJobRequest) SetTriggerConfig(v string) *UpdateJobRequest {
	s.TriggerConfig = &v
	return s
}

func (s *UpdateJobRequest) SetWarStartOptions(v string) *UpdateJobRequest {
	s.WarStartOptions = &v
	return s
}

func (s *UpdateJobRequest) SetWebContainer(v string) *UpdateJobRequest {
	s.WebContainer = &v
	return s
}

func (s *UpdateJobRequest) Validate() error {
	return dara.Validate(s)
}

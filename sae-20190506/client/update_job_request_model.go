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
	// The ARN of the RAM role required to pull images across accounts. For more information, see [Implement cross-account authorization by using RAM roles](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The ACR Enterprise instance ID. This parameter is required when **ImageUrl*	- is a Container Registry Enterprise instance.
	//
	// example:
	//
	// cri-xxxxxx
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The ID of the job template to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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
	// Based on the preceding example, `Command="echo"` and `CommandArgs=["abc", ">", "file0"]`.
	//
	// example:
	//
	// echo
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments for the image startup command. These are the arguments required by the **Command*	- parameter. Format:
	//
	// `["a","b"]`
	//
	// In the preceding example, `CommandArgs=["abc", ">", "file0"]`. The `["abc", ">", "file0"]` part must be converted to a string in JSON array format. If this parameter is not required, leave it empty.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The concurrency policy of the job. Valid values:
	//
	// - **Forbid**: Forbids concurrent execution. A new job is not created if the previous job is not completed.
	//
	// - **Allow**: Allows concurrent execution of jobs.
	//
	// - **Replace**: Replaces the previous job with a new job if the previous job is not completed when the new job is scheduled to be created.
	//
	// example:
	//
	// Allow
	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" xml:"ConcurrencyPolicy,omitempty"`
	// The mount description of the **ConfigMap**. Use the ConfigMaps created on the namespace configuration item page to inject configuration information into the container. Parameter descriptions:
	//
	// - **congfigMapId**: The instance ID of the ConfigMap instance. You can invoke the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the instance ID.
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
	// The custom host mappings in the container. Valid values:
	//
	// - **hostName**: The domain name or hostname.
	//
	// - **ip**: The IP address.
	//
	// example:
	//
	// [{"hostName":"samplehost","ip":"127.0.0.1"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The runtime environment version for the job in the HSF framework, such as the Ali-Tomcat container.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// Specifies whether to enable image acceleration.
	EnableImageAccl *bool `json:"EnableImageAccl,omitempty" xml:"EnableImageAccl,omitempty"`
	// The environment variables for the container. You can specify custom environment variables or reference ConfigMaps. To reference a ConfigMap, you must create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration
	//
	//     - **name**: The name of the environment variable.
	//
	//     - **value**: The value of the environment variable.
	//
	// - Reference a ConfigMap
	//
	//     - **name**: The name of the environment variable. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<ConfigMap name>`, for example, `sae-sys-configmap-all-test1`.
	//
	//     - **valueFrom**: The reference of the environment variable. Set the value to `configMapRef`.
	//
	//     - **configMapId**: The ID of the ConfigMap.
	//
	//     - **key**: The key. If you reference all key-value pairs, do not set this field.
	//
	// example:
	//
	// [{"name":"envtmp","value":"0"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The ID of the corresponding secret.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The URL of the image. This parameter is required when **Package Type*	- is set to **Image**.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/sae_test/ali_sae_test:0.0.1
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The startup arguments for the JAR package job. Default job startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// -Xms4G -Xmx4G
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The startup options for the JAR package job. Default job startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// custom-option
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The JDK version required by the deployment package. Valid values:
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
	// The description of the NAS mount. If the configuration remains unchanged during deployment, you do not need to set this parameter. That is, you do not need to include the **MountDesc*	- field in the request. To clear the NAS configuration, you must set this field to an empty string in the request. That is, set the value of **MountDesc*	- to "".
	//
	// example:
	//
	// [{mountPath: "/tmp", nasPath: "/"}]
	MountDesc *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	// The mount target of the NAS file system in the VPC of the job template. If the configuration remains unchanged during deployment, you do not need to set this parameter. That is, you do not need to include the **MountHost*	- field in the request. To clear the NAS configuration, you must set this field to an empty string in the request. That is, set the value of **MountHost*	- to "".
	//
	// example:
	//
	// 10d3b4bc9****.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The configuration for mounting the NAS file system.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// The ID of the Apsara File Storage NAS file system. If the configuration remains unchanged during deployment, you do not need to set this parameter. That is, you do not need to include the **NasId*	- field in the request. To clear the NAS configuration, you must set this field to an empty string in the request. That is, set the value of **NasId*	- to "".
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
	// The OSS mount description. Parameter descriptions:
	//
	// - **bucketName**: The name of the bucket.
	//
	// - **bucketPath**: The folder or object that you created in OSS. If the OSS mount folder does not exist, an exception is triggered.
	//
	// - **mountPath**: The container path in SAE. If the path already exists, it overwrites the existing content. If the path does not exist, it is created.
	//
	// - **readOnly**: Specifies whether the container path has read permissions on the mount folder resources. Valid values:
	//
	//     - **true**: Read-only permissions.
	//
	//     - **false**: Read and write permissions.
	//
	// example:
	//
	// [{"bucketName": "oss-bucket", "bucketPath": "data/user.data", "mountPath": "/usr/data/user.data", "readOnly": true}]
	OssMountDescs *string `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty"`
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
	// The ACR Enterprise instance ID.
	//
	// example:
	//
	// cri-xxxxxx
	Php *string `json:"Php,omitempty" xml:"Php,omitempty"`
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
	// The script to execute after the container starts. A script is executed immediately after the container is created. Format: `{"exec":{"command":["sh","-c","echo hello"]}}`
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script to execute before the container stops. A script is executed before the container is deleted. Format: `{"exec":{"command":["sh","-c","echo hello"]}}`
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
	// The Python environment. Supported version: **PYTHON 3.9.15**.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// The custom installation module dependencies. By default, the dependencies defined in the requirements.txt file in the root folder are installed. If the file is not configured or you use a custom software package, you can specify the dependencies to install.
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
	// The number of concurrent instances for the job.
	//
	// example:
	//
	// 3
	Replicas *string `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// Specifies whether to enable job slicing.
	//
	// example:
	//
	// true
	Slice *bool `json:"Slice,omitempty" xml:"Slice,omitempty"`
	// The job slicing parameters.
	//
	// example:
	//
	// [0,1,2]
	SliceEnvs *string `json:"SliceEnvs,omitempty" xml:"SliceEnvs,omitempty"`
	// The configuration for collecting logs to Simple Log Service (SLS).
	//
	// - To use SLS resources automatically created by Serverless App Engine (SAE): `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - To use custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// Parameter descriptions:
	//
	// - **projectName**: The name of the SLS project.
	//
	// - **logDir**: The log path.
	//
	// - **logType**: The log type. **stdout*	- indicates container standard output logs. Only one entry is allowed. If no entry is specified, file logs are collected.
	//
	// - **logstoreName**: The name of the SLS Logstore.
	//
	// - **logtailName**: The name of the SLS Logtail. If this parameter is not specified, a new Logtail is created.
	//
	// If the SLS collection configuration remains unchanged during multiple deployments, you do not need to configure this parameter. That is, you do not need to include the **SlsConfigs*	- field in the request. If you no longer need the SLS collection feature, you must set the value of this field to an empty string in the request. That is, set the value of **SlsConfigs*	- to "".
	//
	// > The project automatically created with the job template is deleted when the job template is deleted. Therefore, when selecting an existing project, you cannot select a project automatically created by SAE.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// The timeout period for graceful shutdown. Unit: seconds. Default value: 30. Valid values: 1 to 300.
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
	// The Tomcat file configuration. Set the value to "" or "{}" to delete the configuration. Valid values:
	//
	// - **port**: The port number. Valid values: 1024 to 65535. Ports smaller than 1024 require root permissions. Because the container is configured with admin permissions, specify a port greater than 1024. Default value: 8080.
	//
	// - **contextPath**: The access path. Default value: the root directory "/".
	//
	// - **maxThreads**: The maximum number of connections in the connection pool. Default value: 400.
	//
	// - **uriEncoding**: The encoding format of Tomcat. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. Default value: **ISO-8859-1**.
	//
	// - **useBodyEncoding**: Specifies whether to use **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig  *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	TriggerConfig *string `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
	// The startup command for the WAR package deployment job. The procedure is the same as configuring the startup command for image deployment. For more information, see [Configure a startup command](https://help.aliyun.com/document_detail/96677.html).
	//
	// example:
	//
	// CATALINA_OPTS=\\"$CATALINA_OPTS $Options\\" catalina.sh run
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The Tomcat version required by the deployment package. Valid values:
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

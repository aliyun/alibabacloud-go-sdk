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
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is required to pull images across accounts. For more information, see [Grant permissions across Alibaba Cloud accounts by using a RAM role](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The ID of the Container Registry Enterprise Edition instance. This parameter is required if **ImageUrl*	- is set to an image in a Container Registry Enterprise Edition instance.
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
	// In this example, `Command="echo" and CommandArgs=["abc", ">", "file0"]`.
	//
	// example:
	//
	// echo
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments of the image startup **Command**. The value must be a JSON array that is converted to a string. Format:
	//
	// `["a","b"]`
	//
	// In the preceding example, `CommandArgs=["abc", ">", "file0"]`. The `["abc", ">", "file0"]` array is converted to a string. This parameter is optional.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The policy of running concurrent jobs. Valid values:
	//
	// - **Forbid**: A new job is not created if the previous job is not completed.
	//
	// - **Allow**: Concurrent jobs are allowed.
	//
	// - **Replace**: When the time to create a new job is reached, the new job replaces the previous job if the previous job is not completed.
	//
	// example:
	//
	// Allow
	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" xml:"ConcurrencyPolicy,omitempty"`
	// The description of the **ConfigMap*	- instance that is mounted to the container. You can use the ConfigMap instance created on the Namespace Configurations page to inject configurations into the container. The value is a JSON string. The following fields are supported:
	//
	// - **configMapId**: The ID of the ConfigMap instance. You can call the [ListNamespacedConfigMaps](https://help.aliyun.com/document_detail/176917.html) operation to obtain the ID.
	//
	// - **key**: The key of the key-value pair.
	//
	// > You can pass the `sae-sys-configmap-all` parameter to mount all key-value pairs.
	//
	// - **mountPath**: The mount path.
	//
	// example:
	//
	// [{"configMapId":16,"key":"test","mountPath":"/tmp"}]
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	// The custom mapping between a hostname and an IP address in the container. The value is a JSON string. The following fields are supported:
	//
	// - **hostName**: the domain name or hostname.
	//
	// - **ip**: the IP address.
	//
	// example:
	//
	// [{"hostName":"samplehost","ip":"127.0.0.1"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The version of the application runtime environment in High-speed Service Framework (HSF), such as an Ali-Tomcat container.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// Specifies whether to enable image acceleration.
	EnableImageAccl *bool `json:"EnableImageAccl,omitempty" xml:"EnableImageAccl,omitempty"`
	// The environment variables of the container. You can customize environment variables or reference variables from a ConfigMap. To reference a ConfigMap, you must create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). The value is a JSON string. The following fields are supported:
	//
	// - Custom variables
	//
	//   - **name**: the name of the environment variable.
	//
	//   - **value**: the value of the environment variable.
	//
	// - Reference variables from a ConfigMap
	//
	//   - **name**: The name of the environment variable. You can reference a single key-value pair or all key-value pairs. To reference all key-value pairs, set the value to `sae-sys-configmap-all-<ConfigMap name>`. Example: `sae-sys-configmap-all-test1`.
	//
	//   - **valueFrom**: the reference of the environment variable. Set the value to `configMapRef`.
	//
	//   - **configMapId**: the ID of the ConfigMap.
	//
	//   - **key**: The key of the key-value pair. If you want to reference all key-value pairs, do not configure this field.
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
	// The URL of the image. This parameter is required if **Package Type*	- is set to **Image**.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/sae_test/ali_sae_test:0.0.1
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The arguments of the JAR package to start the application. The default startup command of the application is: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// -Xms4G -Xmx4G
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The options of the JAR package to start the application. The default startup command of the application is: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// custom-option
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The Java Development Kit (JDK) version that the deployment package depends on. The following versions are supported:
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
	// The description of the NAS mount. If the configurations are not changed during a deployment, you do not need to configure this parameter. To clear the NAS configurations, set the value of this parameter to an empty string (`""`) in the request.
	//
	// example:
	//
	// [{mountPath: "/tmp", nasPath: "/"}]
	MountDesc *string `json:"MountDesc,omitempty" xml:"MountDesc,omitempty"`
	// The mount target of the NAS file system in the virtual private cloud (VPC) where the job template is located. If the configurations are not changed during a deployment, you do not need to configure this parameter. To clear the NAS configurations, set the value of this parameter to an empty string (`""`).
	//
	// example:
	//
	// 10d3b4bc9****.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The configurations of mounting a NAS file system.
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// The ID of the Apsara File Storage NAS file system. If the configurations are not changed during a deployment, you do not need to configure this parameter. To clear the NAS configurations, set the value of this parameter to an empty string (`""`).
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
	// The description of the OSS mount. The value is a JSON string. The following parameters are supported:
	//
	// - **bucketName**: the name of the bucket.
	//
	// - **bucketPath**: the directory or object that you created in OSS. An exception occurs if the specified OSS mount directory does not exist.
	//
	// - **mountPath**: The path in the SAE container. If the path exists, the new path overwrites the existing one. If the path does not exist, a new path is created.
	//
	// - **readOnly**: specifies whether a container has the read-only permission on the resources in the mount directory.
	//
	//   - **true**: The container has the read-only permission.
	//
	//   - **false**: The container has the read and write permissions.
	//
	// example:
	//
	// [{"bucketName": "oss-bucket", "bucketPath": "data/user.data", "mountPath": "/usr/data/user.data", "readOnly": true}]
	OssMountDescs *string `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty"`
	// The URL of the deployment package. This parameter is required if **Package Type*	- is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// http://myoss.oss-cn-hangzhou.aliyuncs.com/my-buc/2019-06-30/****.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package. This parameter is required if **Package Type*	- is set to **FatJar**, **War**, or **PythonZip**.
	//
	// example:
	//
	// 1.0.1
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The ID of the Container Registry Enterprise Edition instance.
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
	// The path on which the PHP application startup configuration file is mounted. Make sure that the PHP server uses this configuration file to start the application.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// The script that is executed after the container is started. Example: `{"exec":{"command":["sh","-c","echo hello"]}}`
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script that is executed before the container is stopped. Example: `{"exec":{"command":["sh","-c","echo hello"]}}`
	//
	// example:
	//
	// {"exec":{"command":["sh","-c","echo hello"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The programming language. Supported values: **java**, **php**, **python**, and **shell**.
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
	// The custom module dependencies. By default, the dependencies that are defined in the requirements.txt file in the root directory of the package are installed. If you do not configure this parameter or the package does not have a requirements.txt file, you can specify the dependencies that you want to install.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// The ID of the referenced application.
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
	// Enables job sharding.
	//
	// example:
	//
	// true
	Slice *bool `json:"Slice,omitempty" xml:"Slice,omitempty"`
	// The parameters for job sharding.
	//
	// example:
	//
	// [0,1,2]
	SliceEnvs *string `json:"SliceEnvs,omitempty" xml:"SliceEnvs,omitempty"`
	// The configurations of collecting logs to Log Service.
	//
	// - Use the Log Service resources that are automatically created by SAE: `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - Use a custom Log Service resource: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// The following fields are supported:
	//
	// - **projectName**: The name of the Log Service project.
	//
	// - **logDir**: The log path.
	//
	// - **logType**: The log type. **stdout*	- indicates the standard output log of the container. You can specify only one standard output. If you do not configure this field, file logs are collected.
	//
	// - **logstoreName**: The name of the Logstore in Log Service.
	//
	// - **logtailName**: The name of the Logtail. If you do not specify this parameter, a new Logtail is created.
	//
	// If the SLS configuration is not changed during a deployment, you do not need to configure this parameter. To stop using the log collection feature, set the value of this parameter to an empty string (`""`).
	//
	// > Projects that are automatically created with a job template are deleted when the job template is deleted. Therefore, when you select an existing project, do not select a project that is automatically created by SAE.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// The graceful timeout period. Default value: 30. Unit: seconds. Valid values: 1 to 300.
	//
	// example:
	//
	// 10
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The timeout period for the job. Unit: seconds.
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
	// The configurations of the Tomcat file. If you set this parameter to "" or "{}", the configurations are deleted. The value is a JSON string. The following fields are supported:
	//
	// - **port**: The port number. Valid values: 1024 to 65535. The root permission is required to perform operations on ports whose number is smaller than 1024. The container is configured with the administrator permission. Therefore, specify a port whose number is greater than 1024. If you do not configure this field, the default port 8080 is used.
	//
	// - **contextPath**: The context path. Default value: /.
	//
	// - **maxThreads**: The maximum number of connections in the connection pool. Default value: 400.
	//
	// - **uriEncoding**: The URI encoding scheme in Tomcat. Supported values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. If you do not set this parameter, the default value **ISO-8859-1*	- is used.
	//
	// - **useBodyEncodingForUri**: Specifies whether to use **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig  *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	TriggerConfig *string `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
	// The startup command for the application that is deployed in a WAR package. The procedure is the same as that for configuring the startup command for an image. For more information, see [Set a startup command](https://help.aliyun.com/document_detail/96677.html).
	//
	// example:
	//
	// CATALINA_OPTS=\\"$CATALINA_OPTS $Options\\" catalina.sh run
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The Tomcat version that the deployment package depends on. The following versions are supported:
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeJobResponseBody
	GetCode() *string
	SetData(v *DescribeJobResponseBodyData) *DescribeJobResponseBody
	GetData() *DescribeJobResponseBodyData
	SetErrorCode(v string) *DescribeJobResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeJobResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeJobResponseBody
	GetTraceId() *string
}

type DescribeJobResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information of the job template.
	Data *DescribeJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned. Take note of the following rules:
	//
	// 	- If the call is successful, **ErrorCode*	- is not returned.
	//
	// 	- If the call fails, **ErrorCode*	- is returned. For more information, see the "**Error codes**" section in this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the configurations of the job template were obtained. Valid values:
	//
	// 	- **true**: The configurations were obtained.
	//
	// 	- **false**: The configurations failed to be obtained.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// ac1a0b2215622246421415014e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeJobResponseBody) GetData() *DescribeJobResponseBodyData {
	return s.Data
}

func (s *DescribeJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeJobResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeJobResponseBody) SetCode(v string) *DescribeJobResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobResponseBody) SetData(v *DescribeJobResponseBodyData) *DescribeJobResponseBody {
	s.Data = v
	return s
}

func (s *DescribeJobResponseBody) SetErrorCode(v string) *DescribeJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeJobResponseBody) SetMessage(v string) *DescribeJobResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeJobResponseBody) SetRequestId(v string) *DescribeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobResponseBody) SetSuccess(v bool) *DescribeJobResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobResponseBody) SetTraceId(v string) *DescribeJobResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeJobResponseBodyData struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is used to pull images across accounts. For more information, see [Pull images across Alibaba Cloud accounts](https://help.aliyun.com/document_detail/190675.html) and [Grant permissions across Alibaba Cloud accounts by using a RAM role](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The ID of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// cri-xxxxxx
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The description of the job template.
	//
	// example:
	//
	// Sample application
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The ID of the job template.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the job template.
	//
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of times that the job was retried.
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
	// The arguments of the image startup command. This parameter contains the arguments that are required for **Command**. Format:
	//
	// `["a","b"]`
	//
	// In the preceding **Command*	- example, the CommandArgs parameter is set to `CommandArgs=["abc", ">", "file0"]`. The data type of `["abc", ">", "file0"]` must be an array of strings in the JSON format. If this parameter does not exist in the Command parameter, you do not need to configure it.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The concurrency policy of the job. Valid values:
	//
	// 	- **Forbid**: Concurrent running is prohibited. If the previous job is not completed, no new job is created.
	//
	// 	- **Allow**: Concurrent running is allowed.
	//
	// 	- **Replace**: If the previous job is not completed when the time to create a new job is reached, the new job replaces the previous job.
	//
	// example:
	//
	// Allow
	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" xml:"ConcurrencyPolicy,omitempty"`
	// The details of the ConfigMap.
	ConfigMapMountDesc []*DescribeJobResponseBodyDataConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The CPU specifications required for each instance. Unit: millicore. This parameter cannot be set to 0. Valid values:
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
	// 	- **16000**
	//
	// 	- **32000**
	//
	// example:
	//
	// 1000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The custom mapping between the hostname and IP address in the container. Valid values:
	//
	// 	- **hostName**: the domain name or hostname.
	//
	// 	- **ip**: the IP address.
	//
	// example:
	//
	// [{"hostName":"test.host.name","ip":"0.0.0.0"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The version of the container, such as Ali-Tomcat, in which a job that is developed based on High-speed Service Framework (HSF) is deployed.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// The environment variables. You can configure custom environment variables or reference a ConfigMap. If you want to reference a ConfigMap, you must first create a ConfigMap. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// 	- Custom configuration
	//
	//     	- **name**: the name of the environment variable.
	//
	//     	- **value**: the value of the environment variable.
	//
	// 	- Reference a ConfigMap
	//
	//     	- **name**: the name of the environment variable. You can reference one or all keys. To reference all keys, specify `sae-sys-configmap-all-<ConfigMap name>`. Example: `sae-sys-configmap-all-test1`.
	//
	//     	- **valueFrom**: the reference of the environment variable. Set the value to `configMapRef`.
	//
	//     	- **configMapId**: the ID of the ConfigMap.
	//
	//     	- **key**: the key. If you want to reference all keys, you do not need to configure this parameter.
	//
	// example:
	//
	// [{"name":"TEST_ENV_KEY","value":"TEST_ENV_VAR"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The ID of the corresponding secret.
	//
	// example:
	//
	// 10
	ImagePullSecrets *string `json:"ImagePullSecrets,omitempty" xml:"ImagePullSecrets,omitempty"`
	// The URL of the image. This parameter is returned only if **PackageType*	- is set to **Image**.
	//
	// example:
	//
	// docker.io/library/nginx:1.14.2
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The arguments in the JAR package. The arguments are used to start the job. The default startup command is `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`.
	//
	// example:
	//
	// start
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The option settings in the JAR package. The settings are used to start the job. The default startup command is `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`.
	//
	// example:
	//
	// -Dtest=true
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The version of the Java Development Kit (JDK) on which the deployment package of the application depends. The following versions are supported:
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
	// This parameter is not returned if **PackageType*	- is set to **Image**.
	//
	// example:
	//
	// Open JDK 8
	Jdk *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	// The size of memory that is required by each instance. Unit: MB. This parameter cannot be set to 0. The values of this parameter correspond to the values of the Cpu parameter:
	//
	// 	- This parameter is set to **1024*	- if the Cpu parameter is set to 500 or 1000.
	//
	// 	- This parameter is set to **2048*	- if the Cpu parameter is set to 500, 1000, or 2000.
	//
	// 	- This parameter is set to **4096*	- if the Cpu parameter is set to 1000, 2000, or 4000.
	//
	// 	- This parameter is set to **8192*	- if the Cpu parameter is set to 2000, 4000, or 8000.
	//
	// 	- This parameter is set to **12288*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **16384*	- if the Cpu parameter is set to 4000, 8000, or 16000.
	//
	// 	- This parameter is set to **24567*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **32768*	- if the Cpu parameter is set to 16000.
	//
	// 	- This parameter is set to **65536*	- if the Cpu parameter is set to 8000, 16000, or 32000.
	//
	// 	- This parameter is set to **131072*	- if the Cpu parameter is set to 32000.
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The details of the mounted NAS file system.
	MountDesc []*DescribeJobResponseBodyDataMountDesc `json:"MountDesc,omitempty" xml:"MountDesc,omitempty" type:"Repeated"`
	// The mount target of the Apsara File Storage NAS (NAS) file system in the virtual private cloud (VPC) where the job template is deployed. If you do not need to modify the NAS configurations when you deploy the job template, configure the **MountHost*	- parameter only in the first request. You do not need to include this parameter in subsequent requests. If you no longer need to use NAS, leave the **MountHost*	- parameter empty in the request.
	//
	// example:
	//
	// example.com
	MountHost *string `json:"MountHost,omitempty" xml:"MountHost,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The configurations for mounting the NAS file system.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// The ID of the NAS file system.
	//
	// example:
	//
	// AKSN89**
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The AccessKey ID that is used to read data from and write data to Object Storage Service (OSS).
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
	// The description of mounted OSS buckets.
	OssMountDescs []*DescribeJobResponseBodyDataOssMountDescs `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty" type:"Repeated"`
	// The type of the deployment package. Valid values:
	//
	// 	- If you deploy a Java job template, you can set this parameter to **FatJar**, **War**, or **Image**.
	//
	// 	- If you deploy a PHP job template, the following types are available:
	//
	//     	- **PhpZip**
	//
	//     	- **IMAGE_PHP_5_4**
	//
	//     	- **IMAGE_PHP_5_4_ALPINE**
	//
	//     	- **IMAGE_PHP_5_5**
	//
	//     	- **IMAGE_PHP_5_5_ALPINE**
	//
	//     	- **IMAGE_PHP_5_6**
	//
	//     	- **IMAGE_PHP_5_6_ALPINE**
	//
	//     	- **IMAGE_PHP_7_0**
	//
	//     	- **IMAGE_PHP_7_0_ALPINE**
	//
	//     	- **IMAGE_PHP_7_1**
	//
	//     	- **IMAGE_PHP_7_1_ALPINE**
	//
	//     	- **IMAGE_PHP_7_2**
	//
	//     	- **IMAGE_PHP_7_2_ALPINE**
	//
	//     	- **IMAGE_PHP_7_3**
	//
	//     	- **IMAGE_PHP_7_3_ALPINE**
	//
	// 	- If you deploy a Python job template, you can set this parameter to **PythonZip*	- or **Image**.
	//
	// example:
	//
	// War
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The URL of the deployment package. This parameter is returned only if **PackageType*	- is set to **FatJar*	- or **War**.
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package. This parameter is required only if **PackageType*	- is set to **FatJar*	- or **War**.
	//
	// example:
	//
	// 1.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The details of the PHP configuration file.
	//
	// example:
	//
	// k1=v1
	PhpConfig *string `json:"PhpConfig,omitempty" xml:"PhpConfig,omitempty"`
	// The path on which the PHP configuration file for job startup is mounted. Make sure that the PHP server uses this configuration file during the startup.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// The script that is run immediately after the container is started. Example: `{"exec":{"command":["cat","/etc/group"\\]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script that is run before the container is stopped. Example: `{"exec":{"command":["cat","/etc/group"\\]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The programming language in which the job template is created. Valid values:
	//
	// 	- **java**: Java
	//
	// 	- **php**: PHP
	//
	// 	- **python**: Python
	//
	// 	- **other**: other programming languages, such as C++, Go, .NET, and Node.js
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// The Internet request URLs of one-time jobs.
	PublicWebHookUrls []*string `json:"PublicWebHookUrls,omitempty" xml:"PublicWebHookUrls,omitempty" type:"Repeated"`
	// The Python environment. PYTHON 3.9.15 is supported.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// The configurations for installing custom module dependencies. By default, the dependencies defined by the requirements.txt file in the root directory are installed. If no software package is configured, you can specify dependencies based on your business requirements.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// The ID of the job template that you reference.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	RefAppId *string `json:"RefAppId,omitempty" xml:"RefAppId,omitempty"`
	// The IDs of the referenced job templates.
	RefedAppIds []*string `json:"RefedAppIds,omitempty" xml:"RefedAppIds,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of job instances.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Indicates whether job sharding is enabled.
	//
	// example:
	//
	// true
	Slice *bool `json:"Slice,omitempty" xml:"Slice,omitempty"`
	// The parameters of job sharding.
	//
	// example:
	//
	// SliceEnvs
	SliceEnvs *string `json:"SliceEnvs,omitempty" xml:"SliceEnvs,omitempty"`
	// The logging configurations of Log Service.
	//
	// 	- To use Log Service resources that are automatically created by SAE, set this parameter to `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// 	- To use custom Log Service resources, set this parameter to `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// Parameter description:
	//
	// 	- **projectName**: the name of the Log Service project.
	//
	// 	- **logDir**: the path in which logs are stored.
	//
	// 	- **logType**: the log type. **stdout**: the standard output (stdout) log of the container. Only one stdout value for this parameter can be specified. If this parameter is not configured, file logs are collected.
	//
	// 	- **logstoreName**: the name of the Logstore in Log Service.
	//
	// 	- **logtailName**: the name of the Logtail in Log Service. If this parameter is not configured, a new Logtail is created.
	//
	// If you do not need to modify the logging configurations when you deploy the application, configure **SlsConfigs*	- only in the first request. If you no longer need to use Log Service, leave **SlsConfigs*	- empty in the request.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// Indicates whether the job template is suspended.
	//
	// example:
	//
	// false
	Suspend *bool `json:"Suspend,omitempty" xml:"Suspend,omitempty"`
	// The tags.
	Tags []*DescribeJobResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period for a graceful shutdown. Default value: 30. Unit: seconds. Valid values: 1 to 300.
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
	// The Tomcat configuration. If you want to delete the configuration, set this parameter to {} or leave this parameter empty. Parameter description:
	//
	// 	- **port**: the port number. Valid values: 1024 to 65535. The root permissions are required to perform operations on ports whose number is smaller than 1024. Enter a value that ranges from 1025 to 65535 because the container has only the admin permissions. If this parameter is not configured, the default value 8080 is used.
	//
	// 	- **contextPath**: the path. Default value: /. The value indicates the root directory.
	//
	// 	- **maxThreads**: the maximum number of connections in the connection pool. Default value: 400.
	//
	// 	- **uriEncoding**: the URI encoding scheme in the Tomcat container. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. If this parameter is not configured, the default value **ISO-8859-1*	- is used.
	//
	// 	- **useBodyEncoding**: indicates whether to use the encoding scheme that is specified by **BodyEncoding for URL**. Default value: **true**.
	//
	// example:
	//
	// {"port":8080,"contextPath":"/","maxThreads":400,"uriEncoding":"ISO-8859-1","useBodyEncodingForUri":true}
	TomcatConfig  *string `json:"TomcatConfig,omitempty" xml:"TomcatConfig,omitempty"`
	TriggerConfig *string `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2ze559r1z1bpwqxwp****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2ze0i263cnn311nvj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The internal request URLs for one-time jobs.
	VpcWebHookUrls []*string `json:"VpcWebHookUrls,omitempty" xml:"VpcWebHookUrls,omitempty" type:"Repeated"`
	// The option settings in the WAR package. The settings are used to start the job. The default startup command is `java $JAVA_OPTS $CATALINA_OPTS -Options org.apache.catalina.startup.Bootstrap "$@" start`.
	//
	// example:
	//
	// custom-option
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The version of the Tomcat container on which the deployment package depends. The following versions are supported:
	//
	// 	- **apache-tomcat-7.0.91**
	//
	// 	- **apache-tomcat-8.5.42**
	//
	// This parameter is not returned if **PackageType*	- is set to **Image**.
	//
	// example:
	//
	// apache-tomcat-7.0.91
	WebContainer *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
}

func (s DescribeJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyData) GetAcrAssumeRoleArn() *string {
	return s.AcrAssumeRoleArn
}

func (s *DescribeJobResponseBodyData) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *DescribeJobResponseBodyData) GetAppDescription() *string {
	return s.AppDescription
}

func (s *DescribeJobResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *DescribeJobResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeJobResponseBodyData) GetBackoffLimit() *int64 {
	return s.BackoffLimit
}

func (s *DescribeJobResponseBodyData) GetBestEffortType() *string {
	return s.BestEffortType
}

func (s *DescribeJobResponseBodyData) GetCommand() *string {
	return s.Command
}

func (s *DescribeJobResponseBodyData) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *DescribeJobResponseBodyData) GetConcurrencyPolicy() *string {
	return s.ConcurrencyPolicy
}

func (s *DescribeJobResponseBodyData) GetConfigMapMountDesc() []*DescribeJobResponseBodyDataConfigMapMountDesc {
	return s.ConfigMapMountDesc
}

func (s *DescribeJobResponseBodyData) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeJobResponseBodyData) GetCustomHostAlias() *string {
	return s.CustomHostAlias
}

func (s *DescribeJobResponseBodyData) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *DescribeJobResponseBodyData) GetEnvs() *string {
	return s.Envs
}

func (s *DescribeJobResponseBodyData) GetImagePullSecrets() *string {
	return s.ImagePullSecrets
}

func (s *DescribeJobResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeJobResponseBodyData) GetJarStartArgs() *string {
	return s.JarStartArgs
}

func (s *DescribeJobResponseBodyData) GetJarStartOptions() *string {
	return s.JarStartOptions
}

func (s *DescribeJobResponseBodyData) GetJdk() *string {
	return s.Jdk
}

func (s *DescribeJobResponseBodyData) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeJobResponseBodyData) GetMountDesc() []*DescribeJobResponseBodyDataMountDesc {
	return s.MountDesc
}

func (s *DescribeJobResponseBodyData) GetMountHost() *string {
	return s.MountHost
}

func (s *DescribeJobResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeJobResponseBodyData) GetNasConfigs() *string {
	return s.NasConfigs
}

func (s *DescribeJobResponseBodyData) GetNasId() *string {
	return s.NasId
}

func (s *DescribeJobResponseBodyData) GetOssAkId() *string {
	return s.OssAkId
}

func (s *DescribeJobResponseBodyData) GetOssAkSecret() *string {
	return s.OssAkSecret
}

func (s *DescribeJobResponseBodyData) GetOssMountDescs() []*DescribeJobResponseBodyDataOssMountDescs {
	return s.OssMountDescs
}

func (s *DescribeJobResponseBodyData) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeJobResponseBodyData) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *DescribeJobResponseBodyData) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *DescribeJobResponseBodyData) GetPhpConfig() *string {
	return s.PhpConfig
}

func (s *DescribeJobResponseBodyData) GetPhpConfigLocation() *string {
	return s.PhpConfigLocation
}

func (s *DescribeJobResponseBodyData) GetPostStart() *string {
	return s.PostStart
}

func (s *DescribeJobResponseBodyData) GetPreStop() *string {
	return s.PreStop
}

func (s *DescribeJobResponseBodyData) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *DescribeJobResponseBodyData) GetPublicWebHookUrls() []*string {
	return s.PublicWebHookUrls
}

func (s *DescribeJobResponseBodyData) GetPython() *string {
	return s.Python
}

func (s *DescribeJobResponseBodyData) GetPythonModules() *string {
	return s.PythonModules
}

func (s *DescribeJobResponseBodyData) GetRefAppId() *string {
	return s.RefAppId
}

func (s *DescribeJobResponseBodyData) GetRefedAppIds() []*string {
	return s.RefedAppIds
}

func (s *DescribeJobResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeJobResponseBodyData) GetReplicas() *int32 {
	return s.Replicas
}

func (s *DescribeJobResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeJobResponseBodyData) GetSlice() *bool {
	return s.Slice
}

func (s *DescribeJobResponseBodyData) GetSliceEnvs() *string {
	return s.SliceEnvs
}

func (s *DescribeJobResponseBodyData) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *DescribeJobResponseBodyData) GetSuspend() *bool {
	return s.Suspend
}

func (s *DescribeJobResponseBodyData) GetTags() []*DescribeJobResponseBodyDataTags {
	return s.Tags
}

func (s *DescribeJobResponseBodyData) GetTerminationGracePeriodSeconds() *int32 {
	return s.TerminationGracePeriodSeconds
}

func (s *DescribeJobResponseBodyData) GetTimeout() *int64 {
	return s.Timeout
}

func (s *DescribeJobResponseBodyData) GetTimezone() *string {
	return s.Timezone
}

func (s *DescribeJobResponseBodyData) GetTomcatConfig() *string {
	return s.TomcatConfig
}

func (s *DescribeJobResponseBodyData) GetTriggerConfig() *string {
	return s.TriggerConfig
}

func (s *DescribeJobResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeJobResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeJobResponseBodyData) GetVpcWebHookUrls() []*string {
	return s.VpcWebHookUrls
}

func (s *DescribeJobResponseBodyData) GetWarStartOptions() *string {
	return s.WarStartOptions
}

func (s *DescribeJobResponseBodyData) GetWebContainer() *string {
	return s.WebContainer
}

func (s *DescribeJobResponseBodyData) SetAcrAssumeRoleArn(v string) *DescribeJobResponseBodyData {
	s.AcrAssumeRoleArn = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetAcrInstanceId(v string) *DescribeJobResponseBodyData {
	s.AcrInstanceId = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetAppDescription(v string) *DescribeJobResponseBodyData {
	s.AppDescription = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetAppId(v string) *DescribeJobResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetAppName(v string) *DescribeJobResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetBackoffLimit(v int64) *DescribeJobResponseBodyData {
	s.BackoffLimit = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetBestEffortType(v string) *DescribeJobResponseBodyData {
	s.BestEffortType = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetCommand(v string) *DescribeJobResponseBodyData {
	s.Command = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetCommandArgs(v string) *DescribeJobResponseBodyData {
	s.CommandArgs = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetConcurrencyPolicy(v string) *DescribeJobResponseBodyData {
	s.ConcurrencyPolicy = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetConfigMapMountDesc(v []*DescribeJobResponseBodyDataConfigMapMountDesc) *DescribeJobResponseBodyData {
	s.ConfigMapMountDesc = v
	return s
}

func (s *DescribeJobResponseBodyData) SetCpu(v int32) *DescribeJobResponseBodyData {
	s.Cpu = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetCustomHostAlias(v string) *DescribeJobResponseBodyData {
	s.CustomHostAlias = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetEdasContainerVersion(v string) *DescribeJobResponseBodyData {
	s.EdasContainerVersion = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetEnvs(v string) *DescribeJobResponseBodyData {
	s.Envs = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetImagePullSecrets(v string) *DescribeJobResponseBodyData {
	s.ImagePullSecrets = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetImageUrl(v string) *DescribeJobResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetJarStartArgs(v string) *DescribeJobResponseBodyData {
	s.JarStartArgs = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetJarStartOptions(v string) *DescribeJobResponseBodyData {
	s.JarStartOptions = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetJdk(v string) *DescribeJobResponseBodyData {
	s.Jdk = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetMemory(v int32) *DescribeJobResponseBodyData {
	s.Memory = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetMountDesc(v []*DescribeJobResponseBodyDataMountDesc) *DescribeJobResponseBodyData {
	s.MountDesc = v
	return s
}

func (s *DescribeJobResponseBodyData) SetMountHost(v string) *DescribeJobResponseBodyData {
	s.MountHost = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetNamespaceId(v string) *DescribeJobResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetNasConfigs(v string) *DescribeJobResponseBodyData {
	s.NasConfigs = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetNasId(v string) *DescribeJobResponseBodyData {
	s.NasId = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetOssAkId(v string) *DescribeJobResponseBodyData {
	s.OssAkId = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetOssAkSecret(v string) *DescribeJobResponseBodyData {
	s.OssAkSecret = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetOssMountDescs(v []*DescribeJobResponseBodyDataOssMountDescs) *DescribeJobResponseBodyData {
	s.OssMountDescs = v
	return s
}

func (s *DescribeJobResponseBodyData) SetPackageType(v string) *DescribeJobResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetPackageUrl(v string) *DescribeJobResponseBodyData {
	s.PackageUrl = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetPackageVersion(v string) *DescribeJobResponseBodyData {
	s.PackageVersion = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetPhpConfig(v string) *DescribeJobResponseBodyData {
	s.PhpConfig = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetPhpConfigLocation(v string) *DescribeJobResponseBodyData {
	s.PhpConfigLocation = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetPostStart(v string) *DescribeJobResponseBodyData {
	s.PostStart = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetPreStop(v string) *DescribeJobResponseBodyData {
	s.PreStop = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetProgrammingLanguage(v string) *DescribeJobResponseBodyData {
	s.ProgrammingLanguage = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetPublicWebHookUrls(v []*string) *DescribeJobResponseBodyData {
	s.PublicWebHookUrls = v
	return s
}

func (s *DescribeJobResponseBodyData) SetPython(v string) *DescribeJobResponseBodyData {
	s.Python = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetPythonModules(v string) *DescribeJobResponseBodyData {
	s.PythonModules = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetRefAppId(v string) *DescribeJobResponseBodyData {
	s.RefAppId = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetRefedAppIds(v []*string) *DescribeJobResponseBodyData {
	s.RefedAppIds = v
	return s
}

func (s *DescribeJobResponseBodyData) SetRegionId(v string) *DescribeJobResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetReplicas(v int32) *DescribeJobResponseBodyData {
	s.Replicas = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetSecurityGroupId(v string) *DescribeJobResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetSlice(v bool) *DescribeJobResponseBodyData {
	s.Slice = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetSliceEnvs(v string) *DescribeJobResponseBodyData {
	s.SliceEnvs = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetSlsConfigs(v string) *DescribeJobResponseBodyData {
	s.SlsConfigs = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetSuspend(v bool) *DescribeJobResponseBodyData {
	s.Suspend = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetTags(v []*DescribeJobResponseBodyDataTags) *DescribeJobResponseBodyData {
	s.Tags = v
	return s
}

func (s *DescribeJobResponseBodyData) SetTerminationGracePeriodSeconds(v int32) *DescribeJobResponseBodyData {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetTimeout(v int64) *DescribeJobResponseBodyData {
	s.Timeout = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetTimezone(v string) *DescribeJobResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetTomcatConfig(v string) *DescribeJobResponseBodyData {
	s.TomcatConfig = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetTriggerConfig(v string) *DescribeJobResponseBodyData {
	s.TriggerConfig = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetVSwitchId(v string) *DescribeJobResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetVpcId(v string) *DescribeJobResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetVpcWebHookUrls(v []*string) *DescribeJobResponseBodyData {
	s.VpcWebHookUrls = v
	return s
}

func (s *DescribeJobResponseBodyData) SetWarStartOptions(v string) *DescribeJobResponseBodyData {
	s.WarStartOptions = &v
	return s
}

func (s *DescribeJobResponseBodyData) SetWebContainer(v string) *DescribeJobResponseBodyData {
	s.WebContainer = &v
	return s
}

func (s *DescribeJobResponseBodyData) Validate() error {
	if s.ConfigMapMountDesc != nil {
		for _, item := range s.ConfigMapMountDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MountDesc != nil {
		for _, item := range s.MountDesc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OssMountDescs != nil {
		for _, item := range s.OssMountDescs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJobResponseBodyDataConfigMapMountDesc struct {
	// The ConfigMap ID.
	//
	// example:
	//
	// 1
	ConfigMapId *int64 `json:"ConfigMapId,omitempty" xml:"ConfigMapId,omitempty"`
	// The ConfigMap name.
	//
	// example:
	//
	// test
	ConfigMapName *string `json:"ConfigMapName,omitempty" xml:"ConfigMapName,omitempty"`
	// The key-value pair that is stored in the ConfigMap.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The path on which the ConfigMap is mounted.
	//
	// example:
	//
	// /tmp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s DescribeJobResponseBodyDataConfigMapMountDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyDataConfigMapMountDesc) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyDataConfigMapMountDesc) GetConfigMapId() *int64 {
	return s.ConfigMapId
}

func (s *DescribeJobResponseBodyDataConfigMapMountDesc) GetConfigMapName() *string {
	return s.ConfigMapName
}

func (s *DescribeJobResponseBodyDataConfigMapMountDesc) GetKey() *string {
	return s.Key
}

func (s *DescribeJobResponseBodyDataConfigMapMountDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeJobResponseBodyDataConfigMapMountDesc) SetConfigMapId(v int64) *DescribeJobResponseBodyDataConfigMapMountDesc {
	s.ConfigMapId = &v
	return s
}

func (s *DescribeJobResponseBodyDataConfigMapMountDesc) SetConfigMapName(v string) *DescribeJobResponseBodyDataConfigMapMountDesc {
	s.ConfigMapName = &v
	return s
}

func (s *DescribeJobResponseBodyDataConfigMapMountDesc) SetKey(v string) *DescribeJobResponseBodyDataConfigMapMountDesc {
	s.Key = &v
	return s
}

func (s *DescribeJobResponseBodyDataConfigMapMountDesc) SetMountPath(v string) *DescribeJobResponseBodyDataConfigMapMountDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeJobResponseBodyDataConfigMapMountDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeJobResponseBodyDataMountDesc struct {
	// The path on which the NAS file system is mounted.
	//
	// example:
	//
	// /tmp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The directory in the NAS file system.
	//
	// example:
	//
	// /
	NasPath *string `json:"NasPath,omitempty" xml:"NasPath,omitempty"`
}

func (s DescribeJobResponseBodyDataMountDesc) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyDataMountDesc) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyDataMountDesc) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeJobResponseBodyDataMountDesc) GetNasPath() *string {
	return s.NasPath
}

func (s *DescribeJobResponseBodyDataMountDesc) SetMountPath(v string) *DescribeJobResponseBodyDataMountDesc {
	s.MountPath = &v
	return s
}

func (s *DescribeJobResponseBodyDataMountDesc) SetNasPath(v string) *DescribeJobResponseBodyDataMountDesc {
	s.NasPath = &v
	return s
}

func (s *DescribeJobResponseBodyDataMountDesc) Validate() error {
	return dara.Validate(s)
}

type DescribeJobResponseBodyDataOssMountDescs struct {
	// The name of the bucket.
	//
	// example:
	//
	// oss-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The directory or object in OSS. If the specified directory or object does not exist, an error is returned.
	//
	// example:
	//
	// data/user.data
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	// The path of the container in SAE. The parameter value that you specified overwrites the original value. If the specified path does not exist, SAE automatically creates the path.
	//
	// example:
	//
	// /usr/data/user.data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// Indicates whether the job template can use the container directory to read data from or write data to resources in the directory of the OSS bucket. Valid values:
	//
	// 	- **true**: The job template has the read-only permissions.
	//
	// 	- **false**: The job template has the read and write permissions.
	//
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s DescribeJobResponseBodyDataOssMountDescs) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyDataOssMountDescs) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyDataOssMountDescs) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeJobResponseBodyDataOssMountDescs) GetBucketPath() *string {
	return s.BucketPath
}

func (s *DescribeJobResponseBodyDataOssMountDescs) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeJobResponseBodyDataOssMountDescs) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DescribeJobResponseBodyDataOssMountDescs) SetBucketName(v string) *DescribeJobResponseBodyDataOssMountDescs {
	s.BucketName = &v
	return s
}

func (s *DescribeJobResponseBodyDataOssMountDescs) SetBucketPath(v string) *DescribeJobResponseBodyDataOssMountDescs {
	s.BucketPath = &v
	return s
}

func (s *DescribeJobResponseBodyDataOssMountDescs) SetMountPath(v string) *DescribeJobResponseBodyDataOssMountDescs {
	s.MountPath = &v
	return s
}

func (s *DescribeJobResponseBodyDataOssMountDescs) SetReadOnly(v bool) *DescribeJobResponseBodyDataOssMountDescs {
	s.ReadOnly = &v
	return s
}

func (s *DescribeJobResponseBodyDataOssMountDescs) Validate() error {
	return dara.Validate(s)
}

type DescribeJobResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeJobResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *DescribeJobResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *DescribeJobResponseBodyDataTags) SetKey(v string) *DescribeJobResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *DescribeJobResponseBodyDataTags) SetValue(v string) *DescribeJobResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *DescribeJobResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}

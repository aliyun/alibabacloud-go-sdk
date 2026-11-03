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
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task template information.
	Data *DescribeJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidAppId.NotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information about the call result.
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
	// Indicates whether the task template configuration was retrieved successfully. Valid values:
	//
	// - **true**: Retrieved successfully.
	//
	// - **false**: Failed to retrieve.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID, which is used to accurately query call information.
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
	// The ARN of the RAM role required to pull images across accounts. For more information, see [Pull Alibaba Cloud images across accounts](https://help.aliyun.com/document_detail/190675.html) and [Implement cross-account authorization by using RAM roles](https://help.aliyun.com/document_detail/223585.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	AcrAssumeRoleArn *string `json:"AcrAssumeRoleArn,omitempty" xml:"AcrAssumeRoleArn,omitempty"`
	// The instance ID of the ACR Enterprise instance.
	//
	// example:
	//
	// cri-xxxxxx
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The description of the task template.
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
	// The name of the task template.
	//
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of task retries.
	//
	// example:
	//
	// 3
	BackoffLimit *int64 `json:"BackoffLimit,omitempty" xml:"BackoffLimit,omitempty"`
	// The best-effort policy.
	//
	// example:
	//
	// besteffort
	BestEffortType *string `json:"BestEffortType,omitempty" xml:"BestEffortType,omitempty"`
	// The image startup command. The command must be an executable object that exists in the container. Example:
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
	// The parameters for the image startup command. These are the parameters required by the **Command*	- startup command. Format:
	//
	// `["a","b"]`
	//
	// In the example of the **Command*	- parameter, `CommandArgs=["abc", ">", "file0"]`. The `["abc", ">", "file0"]` part must be converted to a string in JSON array format. If this parameter is not required, you do not need to specify it.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The task concurrency policy. Valid values:
	//
	// - **Forbid**: forbids concurrent execution. A new task is not created if the previous task is not completed.
	//
	// - **Allow**: allows concurrent execution of tasks.
	//
	// - **Replace**: replaces the previous task with a new task if the previous task is not completed when the creation time of the new task arrives.
	//
	// example:
	//
	// Allow
	ConcurrencyPolicy *string `json:"ConcurrencyPolicy,omitempty" xml:"ConcurrencyPolicy,omitempty"`
	// The ConfigMap information.
	ConfigMapMountDesc []*DescribeJobResponseBodyDataConfigMapMountDesc `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty" type:"Repeated"`
	// The CPU required for each instance. Unit: millicores. The value cannot be 0. Only the following defined specifications are supported:
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
	// [{"hostName":"test.host.name","ip":"0.0.0.0"}]
	CustomHostAlias *string `json:"CustomHostAlias,omitempty" xml:"CustomHostAlias,omitempty"`
	// The runtime environment version of the task in the HSF framework, such as the Ali-Tomcat container.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// The container environment variable parameters. You can customize environment variables or reference ConfigMaps. To reference a ConfigMap, create a ConfigMap instance first. For more information, see [CreateConfigMap](https://help.aliyun.com/document_detail/176914.html). Valid values:
	//
	// - Custom configuration:
	//
	//     - **name**: the name of the environment variable.
	//
	//     - **value**: the value of the environment variable.
	//
	// - Reference ConfigMap:
	//
	//     - **name**: the name of the environment variable. You can reference a single key or all keys. To reference all keys, enter `sae-sys-configmap-all-<ConfigMap name>`, for example, `sae-sys-configmap-all-test1`.
	//
	//     - **valueFrom**: the environment variable reference. Set the value to `configMapRef`.
	//
	//     - **configMapId**: the ConfigMap ID.
	//
	//     - **key**: the key. If you reference all key-value pairs, do not set this field.
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
	// The image address. This parameter is required when **PackageType*	- is set to **Image**.
	//
	// example:
	//
	// docker.io/library/nginx:1.14.2
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The startup parameters for the JAR package task. Default task startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// start
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The startup options for the JAR package task. Default task startup command: `$JAVA_HOME/bin/java $JarStartOptions -jar $CATALINA_OPTS "$package_path" $JarStartArgs`
	//
	// example:
	//
	// -Dtest=true
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
	// This parameter is not supported when **PackageType*	- is set to **Image**.
	//
	// example:
	//
	// Open JDK 8
	Jdk *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	// The memory required for each instance. Unit: MB. The value cannot be 0. The memory has a one-to-one mapping with the CPU. Only the following defined specifications are supported:
	//
	// - **1024**: corresponds to 500 millicores and 1000 millicores.
	//
	// - **2048**: corresponds to 500, 1000, and 2000 millicores.
	//
	// - **4096**: corresponds to 1000, 2000, and 4000 millicores.
	//
	// - **8192**: corresponds to 2000, 4000, and 8000 millicores.
	//
	// - **12288**: corresponds to 12000 millicores.
	//
	// - **16384**: corresponds to 4000, 8000, and 16000 millicores.
	//
	// - **24576**: corresponds to 12000 millicores.
	//
	// - **32768**: corresponds to 16000 millicores.
	//
	// - **65536**: corresponds to 8000, 16000, and 32000 millicores.
	//
	// - **131072**: corresponds to 32000 millicores.
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The mount description.
	MountDesc []*DescribeJobResponseBodyDataMountDesc `json:"MountDesc,omitempty" xml:"MountDesc,omitempty" type:"Repeated"`
	// The mount target of the NAS file system in the VPC of the task template. If the configuration remains unchanged during deployment, you do not need to set this parameter (that is, the **MountHost*	- field is not required in the request). If you need to clear the NAS configuration, you must set the value of this field to an empty string (that is, set **MountHost*	- to "") in the request.
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
	// The configuration for mounting NAS.
	//
	// example:
	//
	// [{"mountPath":"/test1","readOnly":false,"nasId":"nasId1","mountDomain":"nasId1.cn-shenzhen.nas.aliyuncs.com","nasPath":"/test1"},{"nasId":"nasId2","mountDomain":"nasId2.cn-shenzhen.nas.aliyuncs.com","readOnly":false,"nasPath":"/test2","mountPath":"/test2"}]
	NasConfigs *string `json:"NasConfigs,omitempty" xml:"NasConfigs,omitempty"`
	// The NAS ID.
	//
	// example:
	//
	// AKSN89**
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
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
	// The OSS mount description.
	OssMountDescs []*DescribeJobResponseBodyDataOssMountDescs `json:"OssMountDescs,omitempty" xml:"OssMountDescs,omitempty" type:"Repeated"`
	// The task package type. Valid values:
	//
	// - If you deploy the task by using Java, valid values are **FatJar**, **War**, and **Image**.
	//
	// - If you deploy the task by using PHP, valid values are:
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
	// - If you deploy the task by using Python, valid values are **PythonZip*	- and **Image**.
	//
	// example:
	//
	// War
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The deployment package URL. This parameter is required when **PackageType*	- is set to **FatJar*	- or **War**.
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version number of the deployment package. This parameter is required when **PackageType*	- is set to **FatJar*	- or **War**.
	//
	// example:
	//
	// 1.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The content of the PHP configuration file.
	//
	// example:
	//
	// k1=v1
	PhpConfig *string `json:"PhpConfig,omitempty" xml:"PhpConfig,omitempty"`
	// The mount path of the PHP task startup configuration. You must ensure that the PHP server uses this configuration file to start.
	//
	// example:
	//
	// /usr/local/etc/php/php.ini
	PhpConfigLocation *string `json:"PhpConfigLocation,omitempty" xml:"PhpConfigLocation,omitempty"`
	// The script to execute after the container starts. A script is triggered and executed immediately after the container is created. Format: `{"exec":{"command":["cat","/etc/group"]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script to execute before the container stops. A script is triggered and executed before the container is deleted. Format: `{"exec":{"command":["cat","/etc/group"]}}`
	//
	// example:
	//
	// {"exec":{"command":["cat","/etc/group"]}}
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The programming language of the technology stack used to create the task template. Valid values:
	//
	// - **java**: Java.
	//
	// - **php**: PHP.
	//
	// - **python**: Python.
	//
	// - **other**: Multiple languages, such as C++, Go, .NET, and Node.js.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// The list of public request URLs for the one-time task.
	PublicWebHookUrls []*string `json:"PublicWebHookUrls,omitempty" xml:"PublicWebHookUrls,omitempty" type:"Repeated"`
	// The Python environment. Python 3.9.15 is supported.
	//
	// example:
	//
	// PYTHON 3.9.15
	Python *string `json:"Python,omitempty" xml:"Python,omitempty"`
	// The module dependencies for custom installation. By default, the dependencies defined in the requirements.txt file in the root folder are installed. If no custom package is configured, you can specify the dependencies to install.
	//
	// example:
	//
	// Flask==2.0
	PythonModules *string `json:"PythonModules,omitempty" xml:"PythonModules,omitempty"`
	// The ID of the referenced task template.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	RefAppId *string `json:"RefAppId,omitempty" xml:"RefAppId,omitempty"`
	// The list of IDs of the referenced task targets.
	RefedAppIds []*string `json:"RefedAppIds,omitempty" xml:"RefedAppIds,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of task instances.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Specifies whether to enable task slicing.
	//
	// example:
	//
	// true
	Slice *bool `json:"Slice,omitempty" xml:"Slice,omitempty"`
	// The task slicing parameters.
	//
	// example:
	//
	// SliceEnvs
	SliceEnvs *string `json:"SliceEnvs,omitempty" xml:"SliceEnvs,omitempty"`
	// The configuration for collecting logs to Simple Log Service (SLS).
	//
	// - To use SLS resources automatically created by Serverless App Engine (SAE): `[{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]`.
	//
	// - To use custom SLS resources: `[{"projectName":"test-sls","logType":"stdout","logDir":"","logstoreName":"sae","logtailName":""},{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]`.
	//
	// Parameter descriptions:
	//
	// - **projectName**: the name of the SLS project.
	//
	// - **logDir**: the log path.
	//
	// - **logType**: the log type. **stdout*	- indicates container standard output logs. Only one **stdout*	- log can be set. If this parameter is not set, file logs are collected.
	//
	// - **logstoreName**: the name of the SLS Logstore.
	//
	// - **logtailName**: the name of the SLS Logtail. If this parameter is not specified, a new Logtail is created.
	//
	// If the SLS collection configuration remains unchanged during multiple deployments, you do not need to set this parameter (that is, the **SlsConfigs*	- field is not required in the request). If you no longer need the SLS collection feature, you must set the value of this field to an empty string (that is, set **SlsConfigs*	- to "") in the request.
	//
	// example:
	//
	// [{"logDir":"","logType":"stdout"},{"logDir":"/tmp/a.log"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// Specifies whether to suspend the task template.
	//
	// example:
	//
	// false
	Suspend *bool `json:"Suspend,omitempty" xml:"Suspend,omitempty"`
	// The tag information.
	Tags []*DescribeJobResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The graceful shutdown timeout period. Default value: 30. Unit: seconds. Valid values: 1 to 300.
	//
	// example:
	//
	// 10
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The task timeout period. Unit: seconds.
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
	// The Tomcat file configuration. Set this parameter to "" or "{}" to delete the configuration:
	//
	// - **port**: the port range is 1024 to 65535. Ports smaller than 1024 require root permissions. Because the container is configured with admin permissions, specify a port greater than 1024. Default value: 8080.
	//
	// - **contextPath**: the access path. Default value: the root directory "/".
	//
	// - **maxThreads**: the maximum number of connections in the connection pool. Default value: 400.
	//
	// - **uriEncoding**: the Tomcat encoding format. Valid values: **UTF-8**, **ISO-8859-1**, **GBK**, and **GB2312**. Default value: **ISO-8859-1**.
	//
	// - **useBodyEncoding**: specifies whether to use **BodyEncoding for URL**. Default value: **true**.
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
	// The VPC ID.
	//
	// example:
	//
	// vpc-2ze0i263cnn311nvj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The list of internal request URLs for one-time tasks.
	VpcWebHookUrls []*string `json:"VpcWebHookUrls,omitempty" xml:"VpcWebHookUrls,omitempty" type:"Repeated"`
	// The startup options for the WAR package task. Default task startup command: `java $JAVA_OPTS $CATALINA_OPTS -Options org.apache.catalina.startup.Bootstrap "$@" start`.
	//
	// example:
	//
	// custom-option
	WarStartOptions *string `json:"WarStartOptions,omitempty" xml:"WarStartOptions,omitempty"`
	// The Tomcat version required by the deployment package. Valid values:
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
	// The name of the ConfigMap.
	//
	// example:
	//
	// test
	ConfigMapName *string `json:"ConfigMapName,omitempty" xml:"ConfigMapName,omitempty"`
	// The key of the ConfigMap.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The mount path of the container.
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
	// The mount path of the container.
	//
	// example:
	//
	// /tmp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The relative file directory of the NAS file system.
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
	// The bucket name.
	//
	// example:
	//
	// oss-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The directory or OSS object that you created in OSS. If the OSS mount directory does not exist, an exception is triggered.
	//
	// example:
	//
	// data/user.data
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	// The container path in SAE. If the path already exists, the existing path is overwritten. If the path does not exist, a new path is created.
	//
	// example:
	//
	// /usr/data/user.data
	MountPath *string `json:"mountPath,omitempty" xml:"mountPath,omitempty"`
	// Specifies whether the container path has read permissions on the resources in the mount directory. Valid values:
	//
	// - **true**: read-only permissions.
	//
	// - **false**: read and write permissions.
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

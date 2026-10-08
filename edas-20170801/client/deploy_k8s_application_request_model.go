// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployK8sApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotations(v string) *DeployK8sApplicationRequest
	GetAnnotations() *string
	SetAppId(v string) *DeployK8sApplicationRequest
	GetAppId() *string
	SetArgs(v string) *DeployK8sApplicationRequest
	GetArgs() *string
	SetBatchTimeout(v int32) *DeployK8sApplicationRequest
	GetBatchTimeout() *int32
	SetBatchWaitTime(v int32) *DeployK8sApplicationRequest
	GetBatchWaitTime() *int32
	SetBuildPackId(v string) *DeployK8sApplicationRequest
	GetBuildPackId() *string
	SetCanaryRuleId(v string) *DeployK8sApplicationRequest
	GetCanaryRuleId() *string
	SetChangeOrderDesc(v string) *DeployK8sApplicationRequest
	GetChangeOrderDesc() *string
	SetCommand(v string) *DeployK8sApplicationRequest
	GetCommand() *string
	SetConfigMountDescs(v string) *DeployK8sApplicationRequest
	GetConfigMountDescs() *string
	SetCpuLimit(v int32) *DeployK8sApplicationRequest
	GetCpuLimit() *int32
	SetCpuRequest(v int32) *DeployK8sApplicationRequest
	GetCpuRequest() *int32
	SetCustomAffinity(v string) *DeployK8sApplicationRequest
	GetCustomAffinity() *string
	SetCustomAgentVersion(v string) *DeployK8sApplicationRequest
	GetCustomAgentVersion() *string
	SetCustomTolerations(v string) *DeployK8sApplicationRequest
	GetCustomTolerations() *string
	SetDeployAcrossNodes(v string) *DeployK8sApplicationRequest
	GetDeployAcrossNodes() *string
	SetDeployAcrossZones(v string) *DeployK8sApplicationRequest
	GetDeployAcrossZones() *string
	SetEdasContainerVersion(v string) *DeployK8sApplicationRequest
	GetEdasContainerVersion() *string
	SetEmptyDirs(v string) *DeployK8sApplicationRequest
	GetEmptyDirs() *string
	SetEnableAhas(v bool) *DeployK8sApplicationRequest
	GetEnableAhas() *bool
	SetEnableEmptyPushReject(v bool) *DeployK8sApplicationRequest
	GetEnableEmptyPushReject() *bool
	SetEnableLosslessRule(v bool) *DeployK8sApplicationRequest
	GetEnableLosslessRule() *bool
	SetEnvFroms(v string) *DeployK8sApplicationRequest
	GetEnvFroms() *string
	SetEnvs(v string) *DeployK8sApplicationRequest
	GetEnvs() *string
	SetImage(v string) *DeployK8sApplicationRequest
	GetImage() *string
	SetImagePlatforms(v string) *DeployK8sApplicationRequest
	GetImagePlatforms() *string
	SetImageTag(v string) *DeployK8sApplicationRequest
	GetImageTag() *string
	SetInitContainers(v string) *DeployK8sApplicationRequest
	GetInitContainers() *string
	SetJDK(v string) *DeployK8sApplicationRequest
	GetJDK() *string
	SetJavaStartUpConfig(v string) *DeployK8sApplicationRequest
	GetJavaStartUpConfig() *string
	SetLabels(v string) *DeployK8sApplicationRequest
	GetLabels() *string
	SetLimitEphemeralStorage(v int32) *DeployK8sApplicationRequest
	GetLimitEphemeralStorage() *int32
	SetLiveness(v string) *DeployK8sApplicationRequest
	GetLiveness() *string
	SetLocalVolume(v string) *DeployK8sApplicationRequest
	GetLocalVolume() *string
	SetLosslessRuleAligned(v bool) *DeployK8sApplicationRequest
	GetLosslessRuleAligned() *bool
	SetLosslessRuleDelayTime(v int32) *DeployK8sApplicationRequest
	GetLosslessRuleDelayTime() *int32
	SetLosslessRuleFuncType(v int32) *DeployK8sApplicationRequest
	GetLosslessRuleFuncType() *int32
	SetLosslessRuleRelated(v bool) *DeployK8sApplicationRequest
	GetLosslessRuleRelated() *bool
	SetLosslessRuleWarmupTime(v int32) *DeployK8sApplicationRequest
	GetLosslessRuleWarmupTime() *int32
	SetMcpuLimit(v int32) *DeployK8sApplicationRequest
	GetMcpuLimit() *int32
	SetMcpuRequest(v int32) *DeployK8sApplicationRequest
	GetMcpuRequest() *int32
	SetMemoryLimit(v int32) *DeployK8sApplicationRequest
	GetMemoryLimit() *int32
	SetMemoryRequest(v int32) *DeployK8sApplicationRequest
	GetMemoryRequest() *int32
	SetMountDescs(v string) *DeployK8sApplicationRequest
	GetMountDescs() *string
	SetNasId(v string) *DeployK8sApplicationRequest
	GetNasId() *string
	SetPackageUrl(v string) *DeployK8sApplicationRequest
	GetPackageUrl() *string
	SetPackageVersion(v string) *DeployK8sApplicationRequest
	GetPackageVersion() *string
	SetPackageVersionId(v string) *DeployK8sApplicationRequest
	GetPackageVersionId() *string
	SetPostStart(v string) *DeployK8sApplicationRequest
	GetPostStart() *string
	SetPreStop(v string) *DeployK8sApplicationRequest
	GetPreStop() *string
	SetPvcMountDescs(v string) *DeployK8sApplicationRequest
	GetPvcMountDescs() *string
	SetReadiness(v string) *DeployK8sApplicationRequest
	GetReadiness() *string
	SetReplicas(v int32) *DeployK8sApplicationRequest
	GetReplicas() *int32
	SetRequestsEphemeralStorage(v int32) *DeployK8sApplicationRequest
	GetRequestsEphemeralStorage() *int32
	SetRuntimeClassName(v string) *DeployK8sApplicationRequest
	GetRuntimeClassName() *string
	SetSecurityContext(v string) *DeployK8sApplicationRequest
	GetSecurityContext() *string
	SetSidecars(v string) *DeployK8sApplicationRequest
	GetSidecars() *string
	SetSlsConfigs(v string) *DeployK8sApplicationRequest
	GetSlsConfigs() *string
	SetStartup(v string) *DeployK8sApplicationRequest
	GetStartup() *string
	SetStorageType(v string) *DeployK8sApplicationRequest
	GetStorageType() *string
	SetTerminateGracePeriod(v int32) *DeployK8sApplicationRequest
	GetTerminateGracePeriod() *int32
	SetTrafficControlStrategy(v string) *DeployK8sApplicationRequest
	GetTrafficControlStrategy() *string
	SetUpdateStrategy(v string) *DeployK8sApplicationRequest
	GetUpdateStrategy() *string
	SetUriEncoding(v string) *DeployK8sApplicationRequest
	GetUriEncoding() *string
	SetUseBodyEncoding(v bool) *DeployK8sApplicationRequest
	GetUseBodyEncoding() *bool
	SetUserBaseImageUrl(v string) *DeployK8sApplicationRequest
	GetUserBaseImageUrl() *string
	SetVolumesStr(v string) *DeployK8sApplicationRequest
	GetVolumesStr() *string
	SetWebContainer(v string) *DeployK8sApplicationRequest
	GetWebContainer() *string
	SetWebContainerConfig(v string) *DeployK8sApplicationRequest
	GetWebContainerConfig() *string
}

type DeployK8sApplicationRequest struct {
	// The annotations for the application pod.
	//
	// example:
	//
	// {"annotation-name-1":"annotation-value-1","annotation-name-2":"annotation-value-2"}
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The application ID. Obtain the ID by calling the ListApplication operation. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// e83acea6-****-47e1-96ae-c0e953772cdc
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The arguments for the container startup command. The value must be a JSON array of strings, such as `["Argument 1", "Argument 2"]`. To clear the arguments, set the parameter to an empty JSON array `"[]"`.
	//
	// example:
	//
	// ["args1","args2"]
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The timeout period for a single batch release. Unit: seconds.
	//
	// example:
	//
	// 60
	BatchTimeout *int32 `json:"BatchTimeout,omitempty" xml:"BatchTimeout,omitempty"`
	// The minimum interval for a phased release of pods. For more information, see [minReadySeconds](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#min-ready-seconds).
	//
	// example:
	//
	// 0
	BatchWaitTime *int32 `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	// The build package number for EDAS Container:
	//
	// - If you do not need to change the EDAS Container version during deployment, you can leave this parameter unset.
	//
	// - To update the EDAS Container version of the target application during this deployment, you must set this parameter.
	//
	// You can obtain the number in two ways:
	//
	// - Call the ListBuildPack operation to query the list of container versions. For more information, see [ListBuildPack](https://help.aliyun.com/document_detail/423222.html).
	//
	// - Obtain it from the **Build Package Number*	- column in the [Version guide](https://help.aliyun.com/document_detail/92614.html) table. For example, `59` indicates `EDAS Container 3.5.8`.
	//
	// example:
	//
	// 59
	BuildPackId *string `json:"BuildPackId,omitempty" xml:"BuildPackId,omitempty"`
	// The ID of the canary release rule policy.
	//
	// example:
	//
	// a8daf22e-****-968c7ff2ea34
	CanaryRuleId *string `json:"CanaryRuleId,omitempty" xml:"CanaryRuleId,omitempty"`
	// The description of the change record.
	//
	// example:
	//
	// Upgrade
	ChangeOrderDesc *string `json:"ChangeOrderDesc,omitempty" xml:"ChangeOrderDesc,omitempty"`
	// The container startup command.
	//
	// > To clear this configuration, set the parameter to an empty string `""`.
	//
	// example:
	//
	// ls
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// Configures Kubernetes ConfigMap and Secret mounts. This lets you mount a ConfigMap or Secret to a specified container directory. The parameters for \\`ConfigMountDescs\\` are as follows:
	//
	// - \\`name\\`: The name of the ConfigMap or Secret.
	//
	// - \\`type\\`: The configuration type. \\`ConfigMap\\` and \\`Secret\\` are supported.
	//
	// - \\`mountPath\\`: The mount path. An absolute path in the container that starts with a forward slash (/).
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "name": "nginx-config",
	//
	//             "type": "ConfigMap",
	//
	//             "mountPath": "/etc/nginx"
	//
	//       },
	//
	//       {
	//
	//             "name": "tls-secret",
	//
	//             "type": "Secret",
	//
	//             "mountPath": "/etc/ssh"
	//
	//       }
	//
	// ]
	ConfigMountDescs *string `json:"ConfigMountDescs,omitempty" xml:"ConfigMountDescs,omitempty"`
	// The CPU limit for the application instance during runtime. Unit: cores. A value of 0 means no limit.
	//
	// example:
	//
	// 1
	CpuLimit *int32 `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// The CPU quota to request for the application instance during runtime. Setting this parameter is recommended.
	//
	// Unit: cores. A value of 0 means no limit.
	//
	// > If you set this parameter, also set the CpuLimit parameter. The value of CpuRequest must be less than or equal to the value of CpuLimit.
	//
	// example:
	//
	// 0
	CpuRequest *int32 `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	// The pod affinity configuration. This takes effect only when both \\`DeployAcrossNodes\\` and \\`DeployAcrossZones\\` are \\`false\\`.
	//
	// example:
	//
	// {"nodeAffinity":{"requiredDuringSchedulingIgnoredDuringExecution":{"nodeSelectorTerms":[{"matchExpressions":[{"key":"beta.kubernetes.io/arch","operator":"NotIn","values":["arm64","arm32"]}]}]},"preferredDuringSchedulingIgnoredDuringExecution":[{"weight":5,"preference":{"matchExpressions":[{"key":"kubernetes.io/os","operator":"In","values":["linux"]}]}}]},"podAffinity":{"requiredDuringSchedulingIgnoredDuringExecution":[{"namespaces":["default"],"topologyKey":"kubernetes.io/hostname","labelSelector":{"matchExpressions":[{"key":"edas.oam.acname","operator":"NotIn","values":["edas-test-app"]}]}}]},"podAntiAffinity":{"preferredDuringSchedulingIgnoredDuringExecution":[{"podAffinityTerm":{"namespaces":["default"],"topologyKey":"failure-domain.beta.kubernetes.io/zone","labelSelector":{"matchExpressions":[{"key":"edas.oam.acname","operator":"In","values":["edas-test-app-2"]}]}},"weight":15}]}}
	CustomAffinity *string `json:"CustomAffinity,omitempty" xml:"CustomAffinity,omitempty"`
	// Sets the version of the custom Application Real-Time Monitoring Service (ARMS) agent to mount to the application.
	//
	// > This feature is available only to whitelisted users. To use this feature, submit a ticket to be added to the whitelist.
	//
	// example:
	//
	// 3.1.4
	CustomAgentVersion *string `json:"CustomAgentVersion,omitempty" xml:"CustomAgentVersion,omitempty"`
	// The pod scheduling toleration configuration. This takes effect only when both \\`DeployAcrossNodes\\` and \\`DeployAcrossZones\\` are \\`false\\`.
	//
	// example:
	//
	// [{"key":"edas-taint-key2","operator":"Exists","effect":"NoExecute","tolerationSeconds":50},{"key":"edas-taint-key","operator":"Equal","value":"edas-taint-value","effect":"PreferNoSchedule"}]
	CustomTolerations *string `json:"CustomTolerations,omitempty" xml:"CustomTolerations,omitempty"`
	// Specifies whether to distribute application instances across multiple nodes. \\`true\\` indicates yes, and other values indicate no.
	//
	// example:
	//
	// true
	DeployAcrossNodes *string `json:"DeployAcrossNodes,omitempty" xml:"DeployAcrossNodes,omitempty"`
	// Specifies whether to distribute application instances across multiple zones. \\`true\\` indicates yes, and other values indicate no.
	//
	// example:
	//
	// true
	DeployAcrossZones *string `json:"DeployAcrossZones,omitempty" xml:"DeployAcrossZones,omitempty"`
	// The EDAS Container version on which the deployment package depends. This parameter applies to HSF applications deployed using WAR packages. It is not supported for image-based deployments.
	//
	// example:
	//
	// 3.5.9
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// Configures Kubernetes \\`emptyDir\\` mounts. This lets you mount an \\`emptyDir\\` volume to a specified container directory. The parameters for \\`EmptyDirs\\` are as follows:
	//
	// - \\`mountPath\\`: The container mount path. This is required.
	//
	// - \\`readOnly\\`: Specifies whether the volume is read-only. Optional. \\`true\\` for read-only, \\`false\\` for read-write. The default is \\`false\\`.
	//
	// - \\`subPathExpr\\`: The subdirectory expression. Optional.
	//
	// example:
	//
	// [{"mountPath":"/app-log","subPathExpr":"$(POD_IP)"},{"readOnly":true,"mountPath":"/etc/nginx"}]
	EmptyDirs *string `json:"EmptyDirs,omitempty" xml:"EmptyDirs,omitempty"`
	// Specifies whether to connect to Application High Availability Service (AHAS).
	//
	// example:
	//
	// true
	EnableAhas *bool `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	// Specifies whether to enable empty push protection:
	//
	// - \\`true\\`: Enable empty push protection.
	//
	// - \\`false\\`: Do not enable empty push protection.
	//
	// example:
	//
	// false
	EnableEmptyPushReject *bool `json:"EnableEmptyPushReject,omitempty" xml:"EnableEmptyPushReject,omitempty"`
	// Specifies whether to enable the graceful start rule:
	//
	// - \\`true\\`: Enable the graceful start rule.
	//
	// - \\`false\\`: Do not enable the graceful start rule.
	//
	// example:
	//
	// true
	EnableLosslessRule *bool `json:"EnableLosslessRule,omitempty" xml:"EnableLosslessRule,omitempty"`
	// Configures environment variables of the Kubernetes \\`EnvFrom\\` type. This mounts a specified ConfigMap or Secret to a directory. Each key corresponds to a file in the directory, and the file content is the value of the key.
	//
	// The parameters for \\`EnvFroms\\` are as follows.
	//
	// - \\`configMapRef\\`: A reference to a ConfigMap. This field includes the following parameter:
	//
	//   - \\`name\\`: The name of the ConfigMap.
	//
	// - \\`secretRef\\`: A reference to a Secret. This field includes the following parameter:
	//
	//   - \\`name\\`: The name of the Secret.
	//
	// example:
	//
	// [{"name":"appname","valueFrom":{"configMapKeyRef":{"name":"appconf","key":"name"}}}]
	EnvFroms *string `json:"EnvFroms,omitempty" xml:"EnvFroms,omitempty"`
	// The environment variables for the deployment. The value must be a JSON array of objects. Three types of environment variables are supported: regular, Kubernetes ConfigMap, and Kubernetes Secret. The format for a regular environment variable is as follows:
	//
	// `{"name":"x", "value": "y"}`
	//
	// A ConfigMap environment variable injects the value of a specified key from a ConfigMap into the container\\"s environment variables. The format is as follows:
	//
	// `{ "name": "x2", "valueFrom": { "configMapKeyRef": { "name": "my-config", "key": "y2" } } }`
	//
	// A Secret environment variable injects the value of a specified key from a Secret into the container\\"s environment variables. The format is as follows:
	//
	// `{ "name": "x3", "valueFrom": { "secretKeyRef": { "name": "my-secret", "key": "y3" } } }`
	//
	// > To clear this configuration, set the parameter to an empty JSON array \\`[]\\`.
	//
	// example:
	//
	// [{"name":"x1","value":"y1"},{"name":"x2","valueFrom":{"configMapKeyRef":{"name":"my-config","key":"y2"}}},{"name":"x3","valueFrom":{"secretKeyRef":{"name":"my-secret","key":"y3"}}}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The full URL of the image. This parameter overwrites the ImageTag parameter.
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The target platform architecture for the image. This is valid when deploying with a WAR or JAR file. Examples:
	//
	// - To specify the x86-64 architecture: \\`linux/amd64\\`
	//
	// - To specify the ARM 64 architecture: \\`linux/arm64\\`
	//
	// - To build a dual-architecture image: \\`linux/amd64,linux/arm64\\`
	//
	// - If you do not enter a value, the default architecture is used.
	//
	// example:
	//
	// linux/arm64,linux/amd64
	ImagePlatforms *string `json:"ImagePlatforms,omitempty" xml:"ImagePlatforms,omitempty"`
	// The image tag.
	//
	// example:
	//
	// latest
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// Sets an init container for the application pod. The container configuration is in YAML format. The value is the base64-encoded YAML configuration of the init container.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "yamlEncoded": "Y29tbWFuZDoKICAtIHNsZWVwCiAgLSAnNjAnCmltYWdlOiAnYnVzeWJveDpsYXRlc3QnCm5hbWU6IGluaXQtYnVzeWJveAo="
	//
	//       }
	//
	// ]
	InitContainers *string `json:"InitContainers,omitempty" xml:"InitContainers,omitempty"`
	// The JDK version on which the deployment package depends. Valid values: Open JDK 7, Open JDK 8, or Custom OpenJDK. This parameter is not supported for image-based deployments. If you use Custom OpenJDK, you must also configure the \\`UserBaseImageUrl\\` field.
	//
	// example:
	//
	// Open JDK 8
	JDK *string `json:"JDK,omitempty" xml:"JDK,omitempty"`
	// The Java startup parameters. You can configure memory, application, garbage collection (GC) policy, tools, service registration and discovery, and custom settings. Correctly configuring these parameters helps reduce GC overhead, shorten server response time, and improve throughput. The parameter is a JSON string. \\`original\\` is the configuration value, and \\`startup\\` is the startup parameter. The system automatically concatenates all \\`startup\\` values as the Java startup parameters for the application. Set to `""` or `"{}"` to delete the configuration.
	//
	// example:
	//
	// {"InitialHeapSize":{"original":512,"startup":"-Xms512m"},"MaxHeapSize":{"original":1024,"startup":"-Xmx1024m"}}
	JavaStartUpConfig *string `json:"JavaStartUpConfig,omitempty" xml:"JavaStartUpConfig,omitempty"`
	// The labels for the application pod.
	//
	// example:
	//
	// {"label-name-1":"label-value-1","label-name-2":"label-value-2"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The upper limit of the temporary storage resource requirement. Unit: GB. A value of 0 means no limit.
	//
	// example:
	//
	// 4
	LimitEphemeralStorage *int32 `json:"LimitEphemeralStorage,omitempty" xml:"LimitEphemeralStorage,omitempty"`
	// The liveness probe for the container. Example: `{"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}`. To delete this configuration, set the parameter to `""` or `{}`. If you do not set this parameter, the configuration is ignored.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The configuration for mounting a host file to a container. Example: `[{"type":"","nodePath":"/localfiles","mountPath":"/app/files"},{"type":"Directory","nodePath":"/mnt","mountPath":"/app/storage"}]`. In this example, \\`nodePath\\` is the host path, \\`mountPath\\` is the path in the container, and \\`type\\` is the mount type.
	//
	// example:
	//
	// [{"type":"","nodePath":"/localfiles","mountPath":"/app/files"},{"type":"Directory","nodePath":"/mnt","mountPath":"/app/storage"}]
	LocalVolume *string `json:"LocalVolume,omitempty" xml:"LocalVolume,omitempty"`
	// Specifies whether to enable the graceful rolling deployment mode to complete service registration before the readiness probe succeeds:
	//
	// - \\`true\\`: This switch provides a health check for the application on port 55199 and the \\`/health\\` path without intrusion. When service registration is complete, the interface returns 200. Otherwise, it returns 500.
	//
	// > If \\`LosslessRuleRelated\\` is also set to \\`true\\`, this interface checks whether service prefetch is complete.
	//
	// - \\`false\\`: Does not provide an interface for the application to check if service registration is complete.
	//
	// example:
	//
	// false
	LosslessRuleAligned *bool `json:"LosslessRuleAligned,omitempty" xml:"LosslessRuleAligned,omitempty"`
	// The service registration latency. Unit: seconds. The value ranges from 0 to 86400.
	//
	// example:
	//
	// 0
	LosslessRuleDelayTime *int32 `json:"LosslessRuleDelayTime,omitempty" xml:"LosslessRuleDelayTime,omitempty"`
	// The service prefetch curve. The value ranges from 0 to 20. The default is 2, which is suitable for general prefetch scenarios. This indicates that the traffic receiving curve of the service provider follows a quadratic curve during the prefetch period.
	//
	// example:
	//
	// 2
	LosslessRuleFuncType *int32 `json:"LosslessRuleFuncType,omitempty" xml:"LosslessRuleFuncType,omitempty"`
	// Specifies whether to enable the graceful rolling deployment mode to complete service prefetch before the readiness probe succeeds:
	//
	// - \\`true\\`: This switch provides a health check for the application on port 55199 and the \\`/health\\` path without intrusion. When service prefetch is complete, the interface returns 200. Otherwise, it returns 500.
	//
	// - \\`false\\`: Does not provide an interface for the application to check if service prefetch is complete.
	//
	// example:
	//
	// false
	LosslessRuleRelated *bool `json:"LosslessRuleRelated,omitempty" xml:"LosslessRuleRelated,omitempty"`
	// The service prefetch duration. Unit: seconds. The value ranges from 0 to 86400.
	//
	// example:
	//
	// 120
	LosslessRuleWarmupTime *int32 `json:"LosslessRuleWarmupTime,omitempty" xml:"LosslessRuleWarmupTime,omitempty"`
	// The maximum CPU that can be used. Unit: cores. A value of 0 means no limit.
	//
	// example:
	//
	// 0
	McpuLimit *int32 `json:"McpuLimit,omitempty" xml:"McpuLimit,omitempty"`
	// The minimum CPU resource requirement. Unit: cores. A value of 0 means no limit.
	//
	// > If you set this parameter, you must also set the \\`CpuLimit\\` parameter. The value must be less than or equal to the value of \\`CpuLimit\\`.
	//
	// example:
	//
	// 4
	McpuRequest *int32 `json:"McpuRequest,omitempty" xml:"McpuRequest,omitempty"`
	// The memory limit for the application instance during runtime. Unit: MB. A value of 0 means no limit.
	//
	// example:
	//
	// 0
	MemoryLimit *int32 `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// The memory quota to request for the application instance during runtime. Setting this parameter is recommended. Unit: MB. A value of 0 means no request.
	//
	// > If you set this parameter, also set the MemoryLimit parameter. The value of MemoryRequest must be less than or equal to the value of MemoryLimit.
	//
	// example:
	//
	// 0
	MemoryRequest *int32 `json:"MemoryRequest,omitempty" xml:"MemoryRequest,omitempty"`
	// The mount configurations, which are a serialized JSON string. Example: `[{"nasPath": "/k8s","mountPath": "/mnt"},{"nasPath": "/files","mountPath": "/app/files"}]`. In this example, \\`nasPath\\` is the file storage path and \\`mountPath\\` is the path in the container to which the file system is mounted.
	//
	// example:
	//
	// [{"nasPath": "/k8s","mountPath": "/mnt"},{"nasPath": "/files","mountPath": "/app/files"}]
	MountDescs *string `json:"MountDescs,omitempty" xml:"MountDescs,omitempty"`
	// The ID of the Apsara File Storage NAS (NAS) file system to mount. The NAS file system must be in the same region as the cluster. It must have an available mount target quota, or its mount target must be on a vSwitch in the VPC. If you do not set this parameter but the \\`mountDescs\\` field exists, a NAS file system is automatically purchased and mounted to a vSwitch in the VPC by default.
	//
	// example:
	//
	// dfs23****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The URL of the deployment package. Configure this parameter for applications deployed using a FatJar or WAR package.
	//
	// > The Java or Python SDK for EDAS POP API must be version 2.44.0 or later.
	//
	// example:
	//
	// https://e***.oss-cn-beijing.aliyuncs.com/s***-1.0-SNAPSHOT-spring-boot.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version number of the deployment package. This parameter is required for WAR and FatJar packages. You can define the meaning of the version number.
	//
	// > The Java or Python SDK for EDAS POP API must be version 2.44.0 or later.
	//
	// example:
	//
	// 20200720
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The ID of the deployment package version.
	//
	// example:
	//
	// 2bcc********
	PackageVersionId *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty"`
	// The script to execute after the container starts. Example: `{"exec":{"command":["cat","/etc/group"]}}`. To delete this configuration, set the parameter to `{}`. If you do not set this parameter, the configuration is ignored.
	//
	// example:
	//
	// {
	//
	//     "exec":{
	//
	//         "command":[
	//
	//             "ls",
	//
	//             "/"
	//
	//         ]
	//
	//     }
	//
	// }
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script to execute before stopping the container. Example: `{"tcpSocket":{"host":"", "port":8080}}`.
	//
	// To delete this configuration, set the parameter to `{}`. If you do not set this parameter, the configuration is ignored.
	//
	// example:
	//
	// {
	//
	//     "exec":{
	//
	//         "command":[
	//
	//             "ls",
	//
	//             "/"
	//
	//         ]
	//
	//     }
	//
	// }
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// Configures Kubernetes PersistentVolumeClaim (PVC) mounts. This lets you mount a Kubernetes PVC volume to a specified container directory. The parameters for \\`PvcMountDescs\\` are as follows:
	//
	// - \\`pvcName\\`: The name of the PVC volume. The PVC volume must already exist and be in the Bound state.
	//
	// - \\`mountPaths\\`: A list of mount directories. You can configure multiple mount directories. Each mount directory supports the following two parameters:
	//
	//   - \\`mountPath\\`: The mount path. An absolute path in the container that starts with a forward slash (/).
	//
	//   - \\`readOnly\\`: The mount mode. \\`true\\` for read-only, \\`false\\` for read-write. The default is \\`false\\`.
	//
	// example:
	//
	// [{"pvcName":"nas-pvc-1","mountPaths":[{"mountPath":"/usr/share/nginx/data"},{"mountPath":"/usr/share/nginx/html","readOnly":true}]}]
	PvcMountDescs *string `json:"PvcMountDescs,omitempty" xml:"PvcMountDescs,omitempty"`
	// The readiness probe for the container. If the probe fails, traffic from the Kubernetes service is not routed to the container. Example: `{"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"httpGet": {"path": "/consumer","port": 8080,"scheme": "HTTP","httpHeaders": [{"name": "test","value": "testvalue"}]}}`. To delete this configuration, set the parameter to `""` or `{}`. If you do not set this parameter, the configuration is ignored.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"httpGet": {"path": "/consumer","port": 8080,"scheme": "HTTP","httpHeaders": [{"name": "test","value": "testvalue"}]}}
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The number of application instances. The minimum value is 0.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The minimum temporary storage resource requirement. Unit: GB. A value of 0 means no limit.
	//
	// example:
	//
	// 2
	RequestsEphemeralStorage *int32 `json:"RequestsEphemeralStorage,omitempty" xml:"RequestsEphemeralStorage,omitempty"`
	// The container runtime type:
	//
	// - \\`runc\\`: regular container runtime.
	//
	// - \\`runv\\`: sandboxed container.
	//
	// This parameter applies only to clusters that use sandboxed containers.
	//
	// example:
	//
	// runc
	RuntimeClassName *string `json:"RuntimeClassName,omitempty" xml:"RuntimeClassName,omitempty"`
	// Sets the \\`SecurityContext\\` property for the application pod container. The value is the base64-encoded YAML configuration of the \\`SecurityContext\\`.
	//
	// example:
	//
	// {"yamlEncoded":"cnVuQXNVc2VyOiAwCnJ1bkFzR3JvdXA6IDA="}
	SecurityContext *string `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty"`
	// Sets a sidecar container for the application pod. The container configuration is in YAML format. The value is the base64-encoded YAML configuration of the sidecar container.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "yamlEncoded": "Y29tbWFuZDoKICAtIHRhaWwKICAtICctZicKICAtIC9kZXYvbnVsbAppbWFnZTogJ2J1c3lib3g6bGF0ZXN0JwpuYW1lOiBidXN5Ym94Cg=="
	//
	//       }
	//
	// ]
	Sidecars *string `json:"Sidecars,omitempty" xml:"Sidecars,omitempty"`
	// The Logstore configuration. Set to `""` or `"{}"` to delete the configuration:
	//
	// - \\`Configs\\`:
	//
	//   - \\`type\\`: The collection type. \\`file\\` for file type, \\`stdout\\` for standard output type.
	//
	//   - \\`Logstore\\`: The name of the Logstore. Make sure the Logstore name is unique within the same cluster. The name must follow these rules:
	//
	//     - It can only contain lowercase letters, numbers, hyphens (-), and underscores (_).
	//
	//     - It must start and end with a lowercase letter or a number.
	//
	//     - The name must be 3 to 63 characters long. If left empty, the system generates a name automatically.
	//
	//   - \\`LogDir\\`: If the type is standard output, the collection path is \\`stdout.log\\`. If the type is file, this is the path of the file to collect. Wildcards are supported. The collection path must match the regular expression: `^/(.+)/(.*)^/$`.
	//
	// example:
	//
	// [{"logstore":"thisisanotherfilelog","type":"file","logDir":"/var/log/*"},{"logstore":"","type":"stdout","logDir":"stdout.log"},{"logstore":"thisisafilelog","type":"file","logDir":"/tmp/log/*"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// The startup probe can be used to perform liveness checks on slow-starting containers to prevent them from being killed before they are up and running. Example: {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"httpGet": {"path": "/consumer","port": 8080,"scheme": "HTTP","httpHeaders": [{"name": "test","value": "testvalue"}]}}.
	//
	// To delete this configuration, set the parameter to "" or {}. If you do not set this parameter, the configuration is ignored.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}
	Startup *string `json:"Startup,omitempty" xml:"Startup,omitempty"`
	// The storage type of the NAS file system. Valid values:
	//
	// - General-purpose NAS: \\`Capacity\\` and \\`Performance\\`
	//
	// - Extreme NAS: \\`standard\\` and \\`advance\\`
	//
	// Currently, only the \\`Performance\\` type is supported.
	//
	// example:
	//
	// Performance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The graceful stop timeout period for the application. Unit: seconds.
	//
	// example:
	//
	// 120
	TerminateGracePeriod *int32 `json:"TerminateGracePeriod,omitempty" xml:"TerminateGracePeriod,omitempty"`
	// The traffic control policy for phased release.
	//
	// example:
	//
	// {"http":{"rules":[{"conditionType":"percent","percent":10}]}}
	TrafficControlStrategy *string `json:"TrafficControlStrategy,omitempty" xml:"TrafficControlStrategy,omitempty"`
	// The phased release policy.
	//
	// - Example 1: Phased release with one canary instance, followed by two batches, automatic batching, and a 1-minute interval.
	//
	//   `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}`
	//
	// - Example 2: Phased release with one canary instance, followed by two batches and manual batching.
	//
	//   `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"manual"},"grayUpdate":{"gray":1}}`
	//
	// - Example 3: Phased release in two batches, with automatic batching and a 0-minute interval.
	//
	//   `{"type":"BatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":0}}`
	//
	// example:
	//
	// {"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}
	UpdateStrategy *string `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
	// The URI encoding format. Supported formats: ISO-8859-1, GBK, GB2312, and UTF-8.
	//
	// > If you do not set this parameter in the application configuration, the default Tomcat value is used.
	//
	// example:
	//
	// GBK
	UriEncoding *string `json:"UriEncoding,omitempty" xml:"UriEncoding,omitempty"`
	// Specifies whether to enable \\`useBodyEncodingForURI\\`.
	//
	// > If you do not set this parameter in the application configuration, the default value \\`false\\` is used.
	//
	// example:
	//
	// false
	UseBodyEncoding *bool `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
	// When using a custom JDK runtime, you must configure the base image address. This address must be publicly accessible. The EDAS server pulls this image to build the application image.
	//
	// example:
	//
	// openjdk:8u302
	UserBaseImageUrl *string `json:"UserBaseImageUrl,omitempty" xml:"UserBaseImageUrl,omitempty"`
	// The data volumes.
	//
	// example:
	//
	// test
	VolumesStr *string `json:"VolumesStr,omitempty" xml:"VolumesStr,omitempty"`
	// The Tomcat version on which the deployment package depends. This parameter applies to Spring Cloud and Dubbo applications deployed using WAR packages. It is not supported for image-based deployments.
	//
	// example:
	//
	// apache-tomcat-7.0.91
	WebContainer *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
	// The Tomcat container configuration. Set to `""` or `"{}"` to delete the configuration:
	//
	// - \\`useDefaultConfig\\`: Specifies whether to use a custom configuration. If \\`true\\`, the custom configuration is not used. If \\`false\\`, the custom configuration is used. If you do not use a custom configuration, the following parameter settings do not take effect.
	//
	// - \\`contextInputType\\`: The access path of the application.
	//
	//   - \\`war\\`: You do not need to enter a custom path. The access path is the name of the WAR package.
	//
	//   - \\`root\\`: You do not need to enter a custom path. The access path is \\`/\\`.
	//
	//   - \\`custom\\`: You need to enter a custom path in the \\`contextPath\\` parameter below.
	//
	// - \\`contextPath\\`: The custom path. This parameter is required only when \\`contextInputType\\` is set to \\`custom\\`.
	//
	// - \\`httpPort\\`: The port number. The valid range is 1024 to 65535. Ports smaller than 1024 require root permissions. Because the container is configured with administrator permissions, specify a port number greater than 1024. If you do not configure this, the default port is 8080.
	//
	// - \\`maxThreads\\`: The size of the connection pool. The default value is 400.
	//
	//   > This configuration greatly affects application performance. Configure it under professional guidance.
	//
	// - \\`uriEncoding\\`: The encoding format for Tomcat. Valid values: UTF-8, ISO-8859-1, GBK, and GB2312. If you do not set this, the default is ISO-8859-1.
	//
	// - \\`useBodyEncoding\\`: Specifies whether to use BodyEncoding for URLs.
	//
	// - \\`useAdvancedServerXml\\`: Specifies whether to use advanced configuration to customize the \\`server.xml\\` file. If the preceding parameter types and values do not meet your needs, you can use the advanced settings to directly edit the Tomcat \\`Server.xml\\` file.
	//
	// - \\`serverXml\\`: The content of the custom \\`server.xml\\` text file in the advanced configuration. This takes effect when \\`useAdvancedServerXml\\` is \\`true\\`.
	//
	// example:
	//
	// {"useDefaultConfig":false,"contextInputType":"custom","contextPath":"hello","httpPort":8088,"maxThreads":400,"uriEncoding":"UTF-8","useBodyEncoding":true,"useAdvancedServerXml":false}
	WebContainerConfig *string `json:"WebContainerConfig,omitempty" xml:"WebContainerConfig,omitempty"`
}

func (s DeployK8sApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeployK8sApplicationRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *DeployK8sApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeployK8sApplicationRequest) GetArgs() *string {
	return s.Args
}

func (s *DeployK8sApplicationRequest) GetBatchTimeout() *int32 {
	return s.BatchTimeout
}

func (s *DeployK8sApplicationRequest) GetBatchWaitTime() *int32 {
	return s.BatchWaitTime
}

func (s *DeployK8sApplicationRequest) GetBuildPackId() *string {
	return s.BuildPackId
}

func (s *DeployK8sApplicationRequest) GetCanaryRuleId() *string {
	return s.CanaryRuleId
}

func (s *DeployK8sApplicationRequest) GetChangeOrderDesc() *string {
	return s.ChangeOrderDesc
}

func (s *DeployK8sApplicationRequest) GetCommand() *string {
	return s.Command
}

func (s *DeployK8sApplicationRequest) GetConfigMountDescs() *string {
	return s.ConfigMountDescs
}

func (s *DeployK8sApplicationRequest) GetCpuLimit() *int32 {
	return s.CpuLimit
}

func (s *DeployK8sApplicationRequest) GetCpuRequest() *int32 {
	return s.CpuRequest
}

func (s *DeployK8sApplicationRequest) GetCustomAffinity() *string {
	return s.CustomAffinity
}

func (s *DeployK8sApplicationRequest) GetCustomAgentVersion() *string {
	return s.CustomAgentVersion
}

func (s *DeployK8sApplicationRequest) GetCustomTolerations() *string {
	return s.CustomTolerations
}

func (s *DeployK8sApplicationRequest) GetDeployAcrossNodes() *string {
	return s.DeployAcrossNodes
}

func (s *DeployK8sApplicationRequest) GetDeployAcrossZones() *string {
	return s.DeployAcrossZones
}

func (s *DeployK8sApplicationRequest) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *DeployK8sApplicationRequest) GetEmptyDirs() *string {
	return s.EmptyDirs
}

func (s *DeployK8sApplicationRequest) GetEnableAhas() *bool {
	return s.EnableAhas
}

func (s *DeployK8sApplicationRequest) GetEnableEmptyPushReject() *bool {
	return s.EnableEmptyPushReject
}

func (s *DeployK8sApplicationRequest) GetEnableLosslessRule() *bool {
	return s.EnableLosslessRule
}

func (s *DeployK8sApplicationRequest) GetEnvFroms() *string {
	return s.EnvFroms
}

func (s *DeployK8sApplicationRequest) GetEnvs() *string {
	return s.Envs
}

func (s *DeployK8sApplicationRequest) GetImage() *string {
	return s.Image
}

func (s *DeployK8sApplicationRequest) GetImagePlatforms() *string {
	return s.ImagePlatforms
}

func (s *DeployK8sApplicationRequest) GetImageTag() *string {
	return s.ImageTag
}

func (s *DeployK8sApplicationRequest) GetInitContainers() *string {
	return s.InitContainers
}

func (s *DeployK8sApplicationRequest) GetJDK() *string {
	return s.JDK
}

func (s *DeployK8sApplicationRequest) GetJavaStartUpConfig() *string {
	return s.JavaStartUpConfig
}

func (s *DeployK8sApplicationRequest) GetLabels() *string {
	return s.Labels
}

func (s *DeployK8sApplicationRequest) GetLimitEphemeralStorage() *int32 {
	return s.LimitEphemeralStorage
}

func (s *DeployK8sApplicationRequest) GetLiveness() *string {
	return s.Liveness
}

func (s *DeployK8sApplicationRequest) GetLocalVolume() *string {
	return s.LocalVolume
}

func (s *DeployK8sApplicationRequest) GetLosslessRuleAligned() *bool {
	return s.LosslessRuleAligned
}

func (s *DeployK8sApplicationRequest) GetLosslessRuleDelayTime() *int32 {
	return s.LosslessRuleDelayTime
}

func (s *DeployK8sApplicationRequest) GetLosslessRuleFuncType() *int32 {
	return s.LosslessRuleFuncType
}

func (s *DeployK8sApplicationRequest) GetLosslessRuleRelated() *bool {
	return s.LosslessRuleRelated
}

func (s *DeployK8sApplicationRequest) GetLosslessRuleWarmupTime() *int32 {
	return s.LosslessRuleWarmupTime
}

func (s *DeployK8sApplicationRequest) GetMcpuLimit() *int32 {
	return s.McpuLimit
}

func (s *DeployK8sApplicationRequest) GetMcpuRequest() *int32 {
	return s.McpuRequest
}

func (s *DeployK8sApplicationRequest) GetMemoryLimit() *int32 {
	return s.MemoryLimit
}

func (s *DeployK8sApplicationRequest) GetMemoryRequest() *int32 {
	return s.MemoryRequest
}

func (s *DeployK8sApplicationRequest) GetMountDescs() *string {
	return s.MountDescs
}

func (s *DeployK8sApplicationRequest) GetNasId() *string {
	return s.NasId
}

func (s *DeployK8sApplicationRequest) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *DeployK8sApplicationRequest) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *DeployK8sApplicationRequest) GetPackageVersionId() *string {
	return s.PackageVersionId
}

func (s *DeployK8sApplicationRequest) GetPostStart() *string {
	return s.PostStart
}

func (s *DeployK8sApplicationRequest) GetPreStop() *string {
	return s.PreStop
}

func (s *DeployK8sApplicationRequest) GetPvcMountDescs() *string {
	return s.PvcMountDescs
}

func (s *DeployK8sApplicationRequest) GetReadiness() *string {
	return s.Readiness
}

func (s *DeployK8sApplicationRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *DeployK8sApplicationRequest) GetRequestsEphemeralStorage() *int32 {
	return s.RequestsEphemeralStorage
}

func (s *DeployK8sApplicationRequest) GetRuntimeClassName() *string {
	return s.RuntimeClassName
}

func (s *DeployK8sApplicationRequest) GetSecurityContext() *string {
	return s.SecurityContext
}

func (s *DeployK8sApplicationRequest) GetSidecars() *string {
	return s.Sidecars
}

func (s *DeployK8sApplicationRequest) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *DeployK8sApplicationRequest) GetStartup() *string {
	return s.Startup
}

func (s *DeployK8sApplicationRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DeployK8sApplicationRequest) GetTerminateGracePeriod() *int32 {
	return s.TerminateGracePeriod
}

func (s *DeployK8sApplicationRequest) GetTrafficControlStrategy() *string {
	return s.TrafficControlStrategy
}

func (s *DeployK8sApplicationRequest) GetUpdateStrategy() *string {
	return s.UpdateStrategy
}

func (s *DeployK8sApplicationRequest) GetUriEncoding() *string {
	return s.UriEncoding
}

func (s *DeployK8sApplicationRequest) GetUseBodyEncoding() *bool {
	return s.UseBodyEncoding
}

func (s *DeployK8sApplicationRequest) GetUserBaseImageUrl() *string {
	return s.UserBaseImageUrl
}

func (s *DeployK8sApplicationRequest) GetVolumesStr() *string {
	return s.VolumesStr
}

func (s *DeployK8sApplicationRequest) GetWebContainer() *string {
	return s.WebContainer
}

func (s *DeployK8sApplicationRequest) GetWebContainerConfig() *string {
	return s.WebContainerConfig
}

func (s *DeployK8sApplicationRequest) SetAnnotations(v string) *DeployK8sApplicationRequest {
	s.Annotations = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetAppId(v string) *DeployK8sApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetArgs(v string) *DeployK8sApplicationRequest {
	s.Args = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetBatchTimeout(v int32) *DeployK8sApplicationRequest {
	s.BatchTimeout = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetBatchWaitTime(v int32) *DeployK8sApplicationRequest {
	s.BatchWaitTime = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetBuildPackId(v string) *DeployK8sApplicationRequest {
	s.BuildPackId = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCanaryRuleId(v string) *DeployK8sApplicationRequest {
	s.CanaryRuleId = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetChangeOrderDesc(v string) *DeployK8sApplicationRequest {
	s.ChangeOrderDesc = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCommand(v string) *DeployK8sApplicationRequest {
	s.Command = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetConfigMountDescs(v string) *DeployK8sApplicationRequest {
	s.ConfigMountDescs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCpuLimit(v int32) *DeployK8sApplicationRequest {
	s.CpuLimit = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCpuRequest(v int32) *DeployK8sApplicationRequest {
	s.CpuRequest = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCustomAffinity(v string) *DeployK8sApplicationRequest {
	s.CustomAffinity = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCustomAgentVersion(v string) *DeployK8sApplicationRequest {
	s.CustomAgentVersion = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCustomTolerations(v string) *DeployK8sApplicationRequest {
	s.CustomTolerations = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetDeployAcrossNodes(v string) *DeployK8sApplicationRequest {
	s.DeployAcrossNodes = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetDeployAcrossZones(v string) *DeployK8sApplicationRequest {
	s.DeployAcrossZones = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEdasContainerVersion(v string) *DeployK8sApplicationRequest {
	s.EdasContainerVersion = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEmptyDirs(v string) *DeployK8sApplicationRequest {
	s.EmptyDirs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEnableAhas(v bool) *DeployK8sApplicationRequest {
	s.EnableAhas = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEnableEmptyPushReject(v bool) *DeployK8sApplicationRequest {
	s.EnableEmptyPushReject = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEnableLosslessRule(v bool) *DeployK8sApplicationRequest {
	s.EnableLosslessRule = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEnvFroms(v string) *DeployK8sApplicationRequest {
	s.EnvFroms = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEnvs(v string) *DeployK8sApplicationRequest {
	s.Envs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetImage(v string) *DeployK8sApplicationRequest {
	s.Image = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetImagePlatforms(v string) *DeployK8sApplicationRequest {
	s.ImagePlatforms = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetImageTag(v string) *DeployK8sApplicationRequest {
	s.ImageTag = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetInitContainers(v string) *DeployK8sApplicationRequest {
	s.InitContainers = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetJDK(v string) *DeployK8sApplicationRequest {
	s.JDK = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetJavaStartUpConfig(v string) *DeployK8sApplicationRequest {
	s.JavaStartUpConfig = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLabels(v string) *DeployK8sApplicationRequest {
	s.Labels = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLimitEphemeralStorage(v int32) *DeployK8sApplicationRequest {
	s.LimitEphemeralStorage = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLiveness(v string) *DeployK8sApplicationRequest {
	s.Liveness = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLocalVolume(v string) *DeployK8sApplicationRequest {
	s.LocalVolume = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLosslessRuleAligned(v bool) *DeployK8sApplicationRequest {
	s.LosslessRuleAligned = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLosslessRuleDelayTime(v int32) *DeployK8sApplicationRequest {
	s.LosslessRuleDelayTime = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLosslessRuleFuncType(v int32) *DeployK8sApplicationRequest {
	s.LosslessRuleFuncType = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLosslessRuleRelated(v bool) *DeployK8sApplicationRequest {
	s.LosslessRuleRelated = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLosslessRuleWarmupTime(v int32) *DeployK8sApplicationRequest {
	s.LosslessRuleWarmupTime = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetMcpuLimit(v int32) *DeployK8sApplicationRequest {
	s.McpuLimit = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetMcpuRequest(v int32) *DeployK8sApplicationRequest {
	s.McpuRequest = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetMemoryLimit(v int32) *DeployK8sApplicationRequest {
	s.MemoryLimit = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetMemoryRequest(v int32) *DeployK8sApplicationRequest {
	s.MemoryRequest = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetMountDescs(v string) *DeployK8sApplicationRequest {
	s.MountDescs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetNasId(v string) *DeployK8sApplicationRequest {
	s.NasId = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPackageUrl(v string) *DeployK8sApplicationRequest {
	s.PackageUrl = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPackageVersion(v string) *DeployK8sApplicationRequest {
	s.PackageVersion = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPackageVersionId(v string) *DeployK8sApplicationRequest {
	s.PackageVersionId = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPostStart(v string) *DeployK8sApplicationRequest {
	s.PostStart = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPreStop(v string) *DeployK8sApplicationRequest {
	s.PreStop = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPvcMountDescs(v string) *DeployK8sApplicationRequest {
	s.PvcMountDescs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetReadiness(v string) *DeployK8sApplicationRequest {
	s.Readiness = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetReplicas(v int32) *DeployK8sApplicationRequest {
	s.Replicas = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetRequestsEphemeralStorage(v int32) *DeployK8sApplicationRequest {
	s.RequestsEphemeralStorage = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetRuntimeClassName(v string) *DeployK8sApplicationRequest {
	s.RuntimeClassName = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetSecurityContext(v string) *DeployK8sApplicationRequest {
	s.SecurityContext = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetSidecars(v string) *DeployK8sApplicationRequest {
	s.Sidecars = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetSlsConfigs(v string) *DeployK8sApplicationRequest {
	s.SlsConfigs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetStartup(v string) *DeployK8sApplicationRequest {
	s.Startup = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetStorageType(v string) *DeployK8sApplicationRequest {
	s.StorageType = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetTerminateGracePeriod(v int32) *DeployK8sApplicationRequest {
	s.TerminateGracePeriod = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetTrafficControlStrategy(v string) *DeployK8sApplicationRequest {
	s.TrafficControlStrategy = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetUpdateStrategy(v string) *DeployK8sApplicationRequest {
	s.UpdateStrategy = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetUriEncoding(v string) *DeployK8sApplicationRequest {
	s.UriEncoding = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetUseBodyEncoding(v bool) *DeployK8sApplicationRequest {
	s.UseBodyEncoding = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetUserBaseImageUrl(v string) *DeployK8sApplicationRequest {
	s.UserBaseImageUrl = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetVolumesStr(v string) *DeployK8sApplicationRequest {
	s.VolumesStr = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetWebContainer(v string) *DeployK8sApplicationRequest {
	s.WebContainer = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetWebContainerConfig(v string) *DeployK8sApplicationRequest {
	s.WebContainerConfig = &v
	return s
}

func (s *DeployK8sApplicationRequest) Validate() error {
	return dara.Validate(s)
}
